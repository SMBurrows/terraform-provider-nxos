// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosIPv4AccessListPolicyEgressInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosIPv4AccessListPolicyEgressInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_access_list_policy_egress_interface.test", "interface_id", "eth1/10"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4AccessListPolicyEgressInterfaceConfig = `

resource "nxos_ipv4_access_list_policy_egress_interface" "test" {
  interface_id = "eth1/10"
}

data "nxos_ipv4_access_list_policy_egress_interface" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_ipv4_access_list_policy_egress_interface.test]
}
`