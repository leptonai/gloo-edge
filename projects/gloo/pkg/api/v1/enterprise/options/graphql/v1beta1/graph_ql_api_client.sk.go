// Code generated by solo-kit. DO NOT EDIT.

package v1beta1

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/errors"
)

type GraphQLApiWatcher interface {
	// watch namespace-scoped GraphqlApis
	Watch(namespace string, opts clients.WatchOpts) (<-chan GraphQLApiList, <-chan error, error)
}

type GraphQLApiClient interface {
	BaseClient() clients.ResourceClient
	Register() error
	Read(namespace, name string, opts clients.ReadOpts) (*GraphQLApi, error)
	Write(resource *GraphQLApi, opts clients.WriteOpts) (*GraphQLApi, error)
	Delete(namespace, name string, opts clients.DeleteOpts) error
	List(namespace string, opts clients.ListOpts) (GraphQLApiList, error)
	GraphQLApiWatcher
}

type graphQLApiClient struct {
	rc clients.ResourceClient
}

func NewGraphQLApiClient(ctx context.Context, rcFactory factory.ResourceClientFactory) (GraphQLApiClient, error) {
	return NewGraphQLApiClientWithToken(ctx, rcFactory, "")
}

func NewGraphQLApiClientWithToken(ctx context.Context, rcFactory factory.ResourceClientFactory, token string) (GraphQLApiClient, error) {
	rc, err := rcFactory.NewResourceClient(ctx, factory.NewResourceClientParams{
		ResourceType: &GraphQLApi{},
		Token:        token,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "creating base GraphQLApi resource client")
	}
	return NewGraphQLApiClientWithBase(rc), nil
}

func NewGraphQLApiClientWithBase(rc clients.ResourceClient) GraphQLApiClient {
	return &graphQLApiClient{
		rc: rc,
	}
}

func (client *graphQLApiClient) BaseClient() clients.ResourceClient {
	return client.rc
}

func (client *graphQLApiClient) Register() error {
	return client.rc.Register()
}

func (client *graphQLApiClient) Read(namespace, name string, opts clients.ReadOpts) (*GraphQLApi, error) {
	opts = opts.WithDefaults()

	resource, err := client.rc.Read(namespace, name, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*GraphQLApi), nil
}

func (client *graphQLApiClient) Write(graphQLApi *GraphQLApi, opts clients.WriteOpts) (*GraphQLApi, error) {
	opts = opts.WithDefaults()
	resource, err := client.rc.Write(graphQLApi, opts)
	if err != nil {
		return nil, err
	}
	return resource.(*GraphQLApi), nil
}

func (client *graphQLApiClient) Delete(namespace, name string, opts clients.DeleteOpts) error {
	opts = opts.WithDefaults()

	return client.rc.Delete(namespace, name, opts)
}

func (client *graphQLApiClient) List(namespace string, opts clients.ListOpts) (GraphQLApiList, error) {
	opts = opts.WithDefaults()

	resourceList, err := client.rc.List(namespace, opts)
	if err != nil {
		return nil, err
	}
	return convertToGraphQLApi(resourceList), nil
}

func (client *graphQLApiClient) Watch(namespace string, opts clients.WatchOpts) (<-chan GraphQLApiList, <-chan error, error) {
	opts = opts.WithDefaults()

	resourcesChan, errs, initErr := client.rc.Watch(namespace, opts)
	if initErr != nil {
		return nil, nil, initErr
	}
	graphqlApisChan := make(chan GraphQLApiList)
	go func() {
		for {
			select {
			case resourceList := <-resourcesChan:
				select {
				case graphqlApisChan <- convertToGraphQLApi(resourceList):
				case <-opts.Ctx.Done():
					close(graphqlApisChan)
					return
				}
			case <-opts.Ctx.Done():
				close(graphqlApisChan)
				return
			}
		}
	}()
	return graphqlApisChan, errs, nil
}

func convertToGraphQLApi(resources resources.ResourceList) GraphQLApiList {
	var graphQLApiList GraphQLApiList
	for _, resource := range resources {
		graphQLApiList = append(graphQLApiList, resource.(*GraphQLApi))
	}
	return graphQLApiList
}
