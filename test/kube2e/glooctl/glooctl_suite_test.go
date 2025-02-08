package glooctl_test

import (
	"context"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kube2e"
	"github.com/solo-io/go-utils/log"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/k8s-utils/testutils/helper"
	skhelpers "github.com/solo-io/solo-kit/test/helpers"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGlooctl(t *testing.T) {
	if os.Getenv("KUBE2E_TESTS") != "glooctl" {
		log.Warnf("This test is disabled. " +
			"To enable, set KUBE2E_TESTS to 'glooctl' in your env.")
		return
	}
	helpers.RegisterGlooDebugLogPrintHandlerAndClearLogs()
	skhelpers.RegisterCommonFailHandlers()
	skhelpers.SetupLog()
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "glooctl Suite", []Reporter{junitReporter})
}

var (
	ctx               context.Context
	cancel            context.CancelFunc
	namespace         = defaults.GlooSystem
	testHelper        *helper.SoloTestHelper
	resourceClientset *kube2e.KubeResourceClientSet
)

var _ = BeforeSuite(StartTestHelper)
var _ = AfterSuite(TearDownTestHelper)

func StartTestHelper() {
	ctx, cancel = context.WithCancel(context.Background())

	var err error
	testHelper, err = kube2e.GetTestHelper(ctx, namespace)
	Expect(err).NotTo(HaveOccurred())
	// Register additional fail handlers
	skhelpers.RegisterPreFailHandler(helpers.KubeDumpOnFail(GinkgoWriter, "istio-system", testHelper.InstallNamespace))

	// Define helm overrides
	valuesOverrideFile, cleanupFunc := getHelmValuesOverrideFile()
	defer cleanupFunc()

	// Install Gloo
	err = testHelper.InstallGloo(ctx, helper.GATEWAY, 5*time.Minute, helper.ExtraArgs("--values", valuesOverrideFile))
	Expect(err).NotTo(HaveOccurred())

	// Check that everything is OK
	kube2e.GlooctlCheckEventuallyHealthy(1, testHelper, "90s")

	// Create KubeResourceClientSet
	cfg, err := kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())

	resourceClientset, err = kube2e.NewKubeResourceClientSet(ctx, cfg)
	Expect(err).NotTo(HaveOccurred())
}

func getHelmValuesOverrideFile() (filename string, cleanup func()) {
	values, err := ioutil.TempFile("", "values-*.yaml")
	Expect(err).NotTo(HaveOccurred())

	// disabling panic threshold
	// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/panic_threshold.html
	_, err = values.Write([]byte(`
gatewayProxies:
  gatewayProxy:
    healthyPanicThreshold: 0
`))
	Expect(err).NotTo(HaveOccurred())

	err = values.Close()
	Expect(err).NotTo(HaveOccurred())

	return values.Name(), func() { _ = os.Remove(values.Name()) }
}

func TearDownTestHelper() {
	if os.Getenv("TEAR_DOWN") == "true" {
		Expect(testHelper).ToNot(BeNil())
		err := testHelper.UninstallGloo()
		Expect(err).NotTo(HaveOccurred())
		_, err = kube2e.MustKubeClient().CoreV1().Namespaces().Get(ctx, testHelper.InstallNamespace, metav1.GetOptions{})
		Expect(apierrors.IsNotFound(err)).To(BeTrue())
	}
	cancel()
}
