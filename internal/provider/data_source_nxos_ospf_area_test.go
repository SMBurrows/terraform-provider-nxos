// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosOSPFArea(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFConfig_all(),
			},
			{
				Config: testAccNxosOSPFConfig_all()+testAccNxosOSPFInstanceConfig_all(),
			},
			{
				Config: testAccNxosOSPFConfig_all()+testAccNxosOSPFInstanceConfig_all()+testAccNxosOSPFVRFConfig_all(),
			},
			{
				Config: testAccNxosOSPFConfig_all()+testAccNxosOSPFInstanceConfig_all()+testAccNxosOSPFVRFConfig_all()+testAccNxosOSPFAreaConfig_all(),
			},
			{
				Config: testAccDataSourceNxosOSPFAreaConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "area_id", "0.0.0.10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "authentication_type", "none"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "cost", "10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "type", "stub"),
				),
			},
		},
	})
}

const testAccDataSourceNxosOSPFAreaConfig = `
data "nxos_ospf_area" "test" {
  instance_name = "OSPF1"
  vrf_name = "default"
  area_id = "0.0.0.10"
}
`