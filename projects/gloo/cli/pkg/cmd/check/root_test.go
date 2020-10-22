package check_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

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

		FIt("temporary test", func() {

			//vs, _ := helpers.MustVirtualServiceClient().Read("gloo-system", "vs", clients.ReadOpts{})
			//print(vs.Metadata.Namespace)

			output := testutils.Glooctl("check")
			fmt.Println(output)
		})
	})

	})
