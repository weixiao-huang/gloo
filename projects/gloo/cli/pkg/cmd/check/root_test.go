package check_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/constants"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/helpers"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/testutils"
)
var _ = Describe("Root", func() {

	Context("with a good vs", func() {
		BeforeEach(func() {
			helpers.UseMemoryClients()
			//
			//// create a settings object
			vsClient := helpers.MustVirtualServiceClient()
			vsvc := &gatewayv1.VirtualService{
				Metadata: core.Metadata{
					Name:      "vs",
					Namespace: "gloo-system",
				},
				VirtualHost: &gatewayv1.VirtualHost{
					Routes: []*gatewayv1.Route{{}},
				},
			}


			vsvc, _ = vsClient.Write(vsvc, clients.WriteOpts{})
			//Expect(err).NotTo(HaveOccurred())
		})

		It("ignore this test", func() {

			//vs, _ := helpers.MustVirtualServiceClient().Read("gloo-system", "vs", clients.ReadOpts{})
			//print(vs.Metadata.Namespace)

			// testutils.Glooctl & testutils.GlooctlOut just spin up a new process; not what we want
			output := testutils.Glooctl("check")
			fmt.Println(output)
		})

		FIt("glooctl test", func() {
			// trying to get in memory clients to mock glooctl completely for check to run
			helpers.UseMemoryClients()
			client := helpers.MustKubeClient()
			client.CoreV1().Namespaces().Create(&v1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: defaults.GlooSystem,
				},
			})
			ns, err := client.CoreV1().Namespaces().Get(defaults.GlooSystem, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			Expect(ns.Name).To(Equal(defaults.GlooSystem))
			if err != nil {
				fmt.Printf("gloo-system NS was not created succesfully - ns - %v, err - %v\n", ns, err)
			}
			helmClient := install.DefaultHelmClient()
			installer := install.NewInstaller(helmClient)
			mode := install.Gloo
			installErr := installer.Install(&install.InstallerConfig{
				InstallCliArgs: &options.Install{
					HelmInstall: options.HelmInstall{
						CreateNamespace:   false,
						Namespace:         defaults.GlooSystem,
						HelmReleaseName:   constants.GlooReleaseName,
					},
				},
				Mode: mode,
			})
			if installErr != nil {
				fmt.Printf("install Err - %v\n", installErr)
			}
			// ns, err := client.CoreV1().Namespaces().Get(defaults.GlooSystem, metav1.GetOptions{})
			// fmt.Printf("ns - %v, err - %v\n", ns, err)
			// ok, err := check.CheckResources(&options.Options{
			// 	Metadata: core.Metadata{Namespace: defaults.GlooSystem},
			// })
			// fmt.Printf("ok? %v, err? %v\n", ok, err)
		})
	})

	})
