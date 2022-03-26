// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosDefaultQOSPolicyInterfaceIn(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosDefaultQOSPolicyInterfaceInConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_default_qos_policy_interface_in.test", "interface_id", "eth1/10"),
				),
			},
			{
				Config: testAccNxosDefaultQOSPolicyInterfaceInConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_default_qos_policy_interface_in.test", "interface_id", "eth1/10"),
				),
			},
			{
				ResourceName:  "nxos_default_qos_policy_interface_in.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/dflt/policy/in/intf-[eth1/10]",
			},
		},
	})
}

func testAccNxosDefaultQOSPolicyInterfaceInConfig_minimum() string {
	return `
	resource "nxos_default_qos_policy_interface_in" "test" {
		interface_id = "eth1/10"
	}
	`
}

func testAccNxosDefaultQOSPolicyInterfaceInConfig_all() string {
	return `
	resource "nxos_default_qos_policy_interface_in" "test" {
		interface_id = "eth1/10"
	}
	`
}
