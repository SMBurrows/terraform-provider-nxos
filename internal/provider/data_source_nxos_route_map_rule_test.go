// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosRouteMapRule(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosRouteMapRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_route_map_rule.test", "name", "RULE1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosRouteMapRuleConfig = `

resource "nxos_route_map_rule" "test" {
  name = "RULE1"
}

data "nxos_route_map_rule" "test" {
  name = "RULE1"
  depends_on = [nxos_route_map_rule.test]
}
`
