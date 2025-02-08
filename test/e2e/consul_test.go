package e2e_test

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"google.golang.org/protobuf/types/known/wrapperspb"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/hashicorp/consul/api"
	"github.com/rotisserie/eris"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	consulplugin "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/consul"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	consul2 "github.com/solo-io/gloo/projects/gloo/pkg/plugins/consul"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams/consul"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/gloo/test/v1helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Consul e2e", func() {

	var (
		ctx                  context.Context
		cancel               context.CancelFunc
		consulInstance       *services.ConsulInstance
		err                  error
		serviceTagsAllowlist []string
		consulWatcher        consul.ConsulWatcher
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		defaults.HttpPort = services.NextBindPort()
		defaults.HttpsPort = services.NextBindPort()

		// Start Consul
		consulInstance, err = consulFactory.NewConsulInstance()
		Expect(err).NotTo(HaveOccurred())
		err = consulInstance.Run()
		Expect(err).NotTo(HaveOccurred())

		// init consul client
		client, err := api.NewClient(api.DefaultConfig())
		Expect(err).NotTo(HaveOccurred())

		serviceTagsAllowlist = []string{"1", "2"}

		// Start Gloo
		consulWatcher, err = consul.NewConsulWatcher(client, nil, serviceTagsAllowlist)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		if consulInstance != nil {
			err = consulInstance.Clean()
			Expect(err).NotTo(HaveOccurred())
		}
		cancel()
	})

	Context("with envoy and gloo", func() {

		var (
			testClients      services.TestClients
			envoyInstance    *services.EnvoyInstance
			envoyPort        uint32
			svc1, svc2, svc3 *v1helpers.TestUpstream
			ro               *services.RunOptions
		)

		const writeNamespace = defaults.GlooSystem

		queryService := func() (string, error) {
			response, err := http.Get(fmt.Sprintf("http://localhost:%d", envoyPort))
			if err != nil {
				return "", err
			}
			//noinspection GoUnhandledErrorResult
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return "", err
			}
			if response.StatusCode != 200 {
				return "", eris.Errorf("bad status code: %v (%v)", response.StatusCode, string(body))
			}
			return string(body), nil
		}

		BeforeEach(func() {
			ro = &services.RunOptions{
				NsToWrite: writeNamespace,
				NsToWatch: []string{"default", writeNamespace},
				WhatToRun: services.What{
					DisableGateway: true,
					DisableUds:     true,
					DisableFds:     true,
				},
				Settings: &gloov1.Settings{
					ConsulDiscovery: &gloov1.Settings_ConsulUpstreamDiscoveryConfiguration{
						ServiceTagsAllowlist: serviceTagsAllowlist,
					},
				},
				ConsulClient:     consulWatcher,
				ConsulDnsAddress: consul2.DefaultDnsAddress,
			}
			testClients = services.RunGlooGatewayUdsFds(ctx, ro)
		})

		JustBeforeEach(func() {
			// Start Envoy
			envoyPort = defaults.HttpPort
			envoyInstance, err = envoyFactory.NewEnvoyInstance()
			Expect(err).NotTo(HaveOccurred())
			envoyInstance.RestXdsPort = uint32(testClients.RestXdsPort)
			err = envoyInstance.RunWithRoleAndRestXds(writeNamespace+"~"+gatewaydefaults.GatewayProxyName, testClients.GlooPort, testClients.RestXdsPort)
			Expect(err).NotTo(HaveOccurred())

			// Run two simple web applications locally
			svc1 = v1helpers.NewTestHttpUpstreamWithReply(ctx, envoyInstance.LocalAddr(), "svc-1")
			svc2 = v1helpers.NewTestHttpUpstreamWithReply(ctx, envoyInstance.LocalAddr(), "svc-2")
			svc3 = v1helpers.NewTestHttpUpstreamWithReply(ctx, envoyInstance.LocalAddr(), "svc-3")

			// Register services with consul
			err = consulInstance.RegisterService("my-svc", "my-svc-1", envoyInstance.GlooAddr, []string{"svc", "1"}, svc1.Port)
			Expect(err).NotTo(HaveOccurred())
			err = consulInstance.RegisterService("my-svc", "my-svc-2", envoyInstance.GlooAddr, []string{"svc", "2"}, svc2.Port)
			Expect(err).NotTo(HaveOccurred())
			//we should not discover this service as it will be filtered out
			err = consulInstance.RegisterService("my-svc-1", "my-svc-3", envoyInstance.GlooAddr, []string{"svc", "3"}, svc3.Port)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			envoyInstance.Clean()
		})

		It("works as expected", func() {
			_, err := testClients.ProxyClient.Write(getProxyWithConsulRoute(writeNamespace, envoyPort), clients.WriteOpts{Ctx: ctx})
			Expect(err).NotTo(HaveOccurred())

			// Wait for proxy to be accepted
			helpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
				return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
			})

			By("requests only go to service with tag '1'")

			// Wait for the endpoints to be registered
			Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc1.C, err
				}
				return svc1.C, nil
			}, "20s", "0.2s").Should(Receive())
			// Service 2 does not match the tags on the route, so we should get only requests from service 1
			Consistently(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc1.C, err
				}
				return svc1.C, nil
			}, "2s", "0.2s").Should(Receive())

			err = consulInstance.RegisterService("my-svc", "my-svc-2", envoyInstance.GlooAddr, []string{"svc", "1"}, svc2.Port)
			Expect(err).NotTo(HaveOccurred())

			By("requests are load balanced between the two services")

			// svc2 first to ensure we also still route to svc1 after registering svc2
			Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc2.C, err
				}
				return svc2.C, nil
			}, "10s", "0.2s").Should(Receive())
			Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc1.C, err
				}
				return svc1.C, nil
			}, "10s", "0.2s").Should(Receive())

		})

		It("resolves consul services with hostname addresses (as opposed to IPs addresses)", func() {
			err = consulInstance.RegisterService("my-svc", "my-svc-1", "my-svc.service.dc1.consul", []string{"svc", "1"}, svc1.Port)
			Expect(err).NotTo(HaveOccurred())

			_, err := testClients.ProxyClient.Write(getProxyWithConsulRoute(writeNamespace, envoyPort), clients.WriteOpts{Ctx: ctx})
			Expect(err).NotTo(HaveOccurred())

			// Wait for proxy to be accepted
			helpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
				return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
			})

			// Wait for endpoints to be discovered
			Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc1.C, err
				}
				return svc1.C, nil
			}, "20s", "0.2s").Should(Receive())

			By("requests only go to service with tag '1'")

			// Service 2 does not match the tags on the route, so we should get only requests from service 1
			Consistently(func() (<-chan *v1helpers.ReceivedRequest, error) {
				_, err := queryService()
				if err != nil {
					return svc1.C, err
				}
				return svc1.C, nil
			}, "2s", "0.2s").Should(Receive())
		})

		Context("test eds only updates", func() {

			runTest := func() {
				_, err := testClients.ProxyClient.Write(getProxyWithConsulRoute(writeNamespace, envoyPort), clients.WriteOpts{Ctx: ctx})
				Expect(err).NotTo(HaveOccurred())

				// Wait for proxy to be accepted
				helpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
					return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
				})

				By("requests only go to endpoints behind test upstream 1")

				// Wait for the endpoints to be registered
				Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
					_, err := queryService()
					if err != nil {
						return svc1.C, err
					}
					return svc1.C, nil
				}, "20s", "0.2s").Should(Receive())
				// Service 2 does not match the tags on the route, so we should get only requests from service 1 with test upstream 1 endpoint
				Consistently(func() (<-chan *v1helpers.ReceivedRequest, error) {
					_, err := queryService()
					if err != nil {
						return svc1.C, err
					}
					return svc1.C, nil
				}, "2s", "0.2s").Should(Receive())

				// update service one to point to test upstream 2 port
				err = consulInstance.RegisterService("my-svc", "my-svc-1", envoyInstance.GlooAddr, []string{"svc", "1"}, svc2.Port)
				Expect(err).NotTo(HaveOccurred())

				By("requests only go to endpoints behind test upstream 2")

				// ensure EDS picked up this endpoint-only change
				Eventually(func() (<-chan *v1helpers.ReceivedRequest, error) {
					_, err := queryService()
					if err != nil {
						return svc2.C, err
					}
					return svc2.C, nil
				}, "20s", "0.2s").Should(Receive())
				// test upstream 1 endpoint is now stale; should only get requests to endpoints for test upstream 2 for svc1
				Consistently(func() (<-chan *v1helpers.ReceivedRequest, error) {
					_, err := queryService()
					if err != nil {
						return svc2.C, err
					}
					return svc2.C, nil
				}, "2s", "0.2s").Should(Receive())
			}

			Context("non-blocking EDS queries", func() {
				It("works as expected", func() {
					runTest()
				})
			})

			Context("blocking EDS queries", func() {

				BeforeEach(func() {
					ro.Settings.ConsulDiscovery.EdsBlockingQueries = &wrapperspb.BoolValue{Value: true}
				})

				It("works as expected", func() {
					runTest()
				})
			})

		})

	})

	// This test was written to prove that the consul golang client behaves differently than the consul CLI, and thus
	// that our `refreshSpecs()` usage in consul eds.go is correct and does not miss updates (which also allows
	// us to make performance optimizations at scale, since our current implementation has a lot more cache hits).
	It("fires service watch even if catalog service is the only update", func() {
		svcsChan, errChan := consulWatcher.WatchServices(ctx, []string{"dc1"}, consulplugin.ConsulConsistencyModes_DefaultMode, nil)

		// use select instead of eventually for easier debugging.
		select {
		case err := <-errChan:
			Expect(err).NotTo(HaveOccurred())
			Fail("err chan closed prematurely")
		case svcsReceived := <-svcsChan:
			// the default consul svc in dc1 does not show up in our watch because of service tag filtering
			Expect(svcsReceived).To(HaveLen(0))
		case <-time.After(5 * time.Second):
			Fail("timeout waiting for services")
		}

		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				// happy path, continue
				ExpectWithOffset(1, svcsReceived).To(HaveLen(0))
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())

		// now that consul has a chance to update raft index with all internal updates; we should see no more updates
		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				ExpectWithOffset(1, svcsReceived).To(HaveLen(0)) // we actually expect len(0) if anything; this is just here to get a nice output / diff before we fail regardless
				Fail("did not expect to receive empty services")
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())

		// add a single service.
		// this will fire a watch via consul CLI for both:
		// - consul watch -type=service -service my-svc ./echo.sh
		// - consul watch -type=services ./echo.sh
		//
		// as the service is new
		//
		// echo.sh is just a shell script with the contents: `echo "watch fired!"` so we can check that the watch fired.
		err = consulInstance.RegisterLiveService("my-svc", "my-svc-1", "127.0.0.1", []string{"svc", "1"}, 80)
		Expect(err).NotTo(HaveOccurred())

		// use select instead of eventually for easier debugging.
		select {
		case err := <-errChan:
			Expect(err).NotTo(HaveOccurred())
			Fail("err chan closed prematurely")
		case svcsReceived := <-svcsChan:
			// the default consul svc in dc1 does not show up in our watch because of service tag filtering
			Expect(svcsReceived).To(HaveLen(1))
			Expect(svcsReceived[0].Name).To(Equal("my-svc"))
			Expect(svcsReceived[0].Tags).To(ConsistOf([]string{"svc", "1"}))
			Expect(svcsReceived[0].DataCenters).To(ConsistOf([]string{"dc1"}))
		case <-time.After(5 * time.Second):
			Fail("timeout waiting for services")
		}

		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				// happy path, continue
				ExpectWithOffset(1, svcsReceived).To(HaveLen(1))
				ExpectWithOffset(1, svcsReceived[0].Name).To(Equal("my-svc"))
				ExpectWithOffset(1, svcsReceived[0].Tags).To(ConsistOf([]string{"svc", "1"}))
				ExpectWithOffset(1, svcsReceived[0].DataCenters).To(ConsistOf([]string{"dc1"}))
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())

		// now that consul has a chance to update raft index with all internal updates; we should see no more updates
		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				ExpectWithOffset(1, svcsReceived).To(HaveLen(0)) // we actually expect len(1) if anything; this is just here to get a nice output / diff before we fail regardless
				Fail("did not expect to receive services")
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())

		// update an existing service.
		// this will fire a watch via consul CLI only for:
		// - consul watch -type=service -service my-svc ./echo.sh
		//
		// as the service is new
		//
		// However this WILL fire an update on our golang client watch for services (last index increments)
		// and thus we can depend on this to signal that we should query again for all catalog services in eds.go
		//
		// It appears the golang client lastIndex mirrors the raft index (per `consul info`), and while dated,
		// this behavior still seems to hold: https://github.com/hashicorp/consul/issues/1244#issuecomment-141146851
		//
		// In the event this behavior changes (this test fails on newer consul versions),
		// we may need to move completely to the `eds_blocking_queries` as true despite the performance implications
		// for correctness.
		err = consulInstance.RegisterLiveService("my-svc", "my-svc-1", "127.0.0.1", []string{"svc", "1"}, 81)
		Expect(err).NotTo(HaveOccurred())

		// this is where golang client differs from cli! CLI registers watch on svc but not svcs; but we get for both
		// use select instead of eventually for easier debugging.
		select {
		case err := <-errChan:
			Expect(err).NotTo(HaveOccurred())
			Fail("err chan closed prematurely")
		case svcsReceived := <-svcsChan:
			// the default consul svc in dc1 does not show up in our watch
			Expect(svcsReceived).To(HaveLen(1))
			Expect(svcsReceived[0].Name).To(Equal("my-svc"))
			Expect(svcsReceived[0].Tags).To(ConsistOf([]string{"svc", "1"}))
			Expect(svcsReceived[0].DataCenters).To(ConsistOf([]string{"dc1"}))
		case <-time.After(5 * time.Second):
			Fail("timeout waiting for services")
		}

		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				// happy path, continue
				ExpectWithOffset(1, svcsReceived).To(HaveLen(1))
				ExpectWithOffset(1, svcsReceived[0].Name).To(Equal("my-svc"))
				ExpectWithOffset(1, svcsReceived[0].Tags).To(ConsistOf([]string{"svc", "1"}))
				ExpectWithOffset(1, svcsReceived[0].DataCenters).To(ConsistOf([]string{"dc1"}))
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())

		// now that consul has a chance to update raft index with all internal updates; we should see no more updates
		Consistently(func() error {
			select {
			case err := <-errChan:
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				return errors.New("err chan closed prematurely")
			case svcsReceived := <-svcsChan:
				ExpectWithOffset(1, svcsReceived).To(HaveLen(0)) // we actually expect len(1) if anything; this is just here to get a nice output / diff before we fail regardless
				Fail("did not expect to receive services")
			case <-time.After(100 * time.Millisecond):
				// happy path, continue
			}
			return nil
		}, "2s", "0.2s").Should(Succeed())
	})

})

func getProxyWithConsulRoute(ns string, bindPort uint32) *gloov1.Proxy {
	return &gloov1.Proxy{
		Metadata: &core.Metadata{
			Name:      gatewaydefaults.GatewayProxyName,
			Namespace: ns,
		},
		Listeners: []*gloov1.Listener{{
			Name:        "listener",
			BindAddress: "::",
			BindPort:    bindPort,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: []*gloov1.VirtualHost{{
						Name:    "vh-1",
						Domains: []string{"*"},
						Routes: []*gloov1.Route{{
							Action: &gloov1.Route_RouteAction{
								RouteAction: &gloov1.RouteAction{
									Destination: &gloov1.RouteAction_Single{
										Single: &gloov1.Destination{
											DestinationType: &gloov1.Destination_Consul{
												Consul: &gloov1.ConsulServiceDestination{
													ServiceName: "my-svc",
													Tags:        []string{"svc", "1"},
												},
											},
										},
									},
								},
							},
						}},
					}},
				},
			},
		}},
	}
}
