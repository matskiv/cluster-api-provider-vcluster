package e2e

import (
	"math/rand"
	"testing"
	"time"

	"github.com/loft-sh/cluster-api-provider-vcluster/api/v1alpha1"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	apiregistrationv1 "k8s.io/kube-aggregator/pkg/apis/apiregistration/v1"

	// Enable cloud provider auth
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	// Register tests
	// _ "github.com/loft-sh/vcluster/test/e2e/kubeconfig"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	// API extensions are not in the above scheme set,
	// and must thus be added separately.
	_ = apiextensionsv1beta1.AddToScheme(scheme)
	_ = apiextensionsv1.AddToScheme(scheme)
	_ = apiregistrationv1.AddToScheme(scheme)
	_ = v1alpha1.AddToScheme(scheme)
}

// TestRunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
// If a "report directory" is specified, one or more JUnit test reports will be
// generated in this directory, and cluster logs will also be saved.
// This function is called on each Ginkgo node in parallel mode.
func TestRunE2ETests(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	gomega.RegisterFailHandler(ginkgo.Fail)

	// Setup?

	var _ = ginkgo.AfterSuite(func() {
		// Cleanup?
	})

	ginkgo.RunSpecs(t, "cluster-api-provider-vcluster e2e suite")
}
