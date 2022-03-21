// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosFeatureOSPF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:testAccNxosFeatureOSPFConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
				),
			},
			{
				Config:testAccNxosFeatureOSPFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_feature_ospf.test", "admin_state", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_feature_ospf.test",
				ImportState:   true,
				ImportStateId: "sys/fm/ospf",
			},
		},
	})
}

func testAccNxosFeatureOSPFConfig_minimum() string {
	return `
	resource "nxos_feature_ospf" "test" {
	}
	`
}

func testAccNxosFeatureOSPFConfig_all() string {
	return `
	resource "nxos_feature_ospf" "test" {
		admin_state = "enabled"
	}
	`
}
