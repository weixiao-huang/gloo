package check_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/check"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Root", func() {
	BeforeEach(func() {
		helpers.UseMemoryClients()
		client := helpers.MustKubeClient()
		client.CoreV1().Namespaces().Create(&corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: defaults.GlooSystem,
			},
		})

		appName := "default"
		client.AppsV1().Deployments("gloo-system").Create(&appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      appName,
				Namespace: "gloo-system",
			},
			Spec: appsv1.DeploymentSpec{},
		})

		// For GetSettings
		helpers.MustNamespacedSettingsClient("gloo-system").Write(&v1.Settings{
			Metadata: core.Metadata{
				Name:      "default",
				Namespace: "gloo-system",
			},
		}, clients.WriteOpts{})
	})

	Context("With a gloo-system namespace", func() {

		It("returns OK connection with default args", func() {
			//opts := &options.Options{
			//	Metadata: core.Metadata{
			//		Namespace: "gloo-system",
			//	},
			//}
			err := check.CheckConnection("gloo-system")
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns connection err when namespace doesn't exist", func() {
			err := check.CheckConnection("my-namespace")
			Expect(err.Error()).To(BeEquivalentTo("Could not communicate with kubernetes cluster: namespaces \"my-namespace\" not found"))
		})
	})

	Context("With a good kube client", func() {

		It("returns OK for all checks", func() {

			options := &options.Options{
				Metadata: core.Metadata{
					Namespace: "gloo-system",
				},
			}

			err := check.CheckConnection(options.Metadata.Namespace)
			Expect(err).NotTo(HaveOccurred())

			deployments, _, err := check.GetAndCheckDeployments(options)
			Expect(err).NotTo(HaveOccurred())

			_, err = check.CheckPods(options)
			Expect(err).NotTo(HaveOccurred())

			settings, err := check.GetSettings(options)
			Expect(err).NotTo(HaveOccurred())

			ns, err := check.GetNamespaces(settings)
			Expect(err).NotTo(HaveOccurred())

			knownUpstreams, _, err := check.CheckUpstreams(ns)
			Expect(err).NotTo(HaveOccurred())

			knownAuthConfigs, _, err := check.CheckAuthConfigs(ns)
			Expect(err).NotTo(HaveOccurred())

			knownRateLimitConfigs, _, err := check.CheckRateLimitConfigs(ns)
			Expect(err).NotTo(HaveOccurred())

			_, err = check.CheckVirtualServices(ns, knownUpstreams, knownAuthConfigs, knownRateLimitConfigs)
			Expect(err).NotTo(HaveOccurred())

			_, err = check.CheckGateways(ns)
			Expect(err).NotTo(HaveOccurred())

			_, err = check.CheckProxies(options.Top.Ctx, ns, options.Metadata.Namespace, deployments)
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
