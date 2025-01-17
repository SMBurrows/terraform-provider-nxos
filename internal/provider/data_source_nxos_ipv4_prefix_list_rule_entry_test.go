// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosIPv4PrefixListRuleEntry(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosIPv4PrefixListRuleEntryConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "order", "10"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "action", "permit"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "criteria", "exact"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "prefix", "192.168.1.0/24"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "from_range", "0"),
					resource.TestCheckResourceAttr("data.nxos_ipv4_prefix_list_rule_entry.test", "to_range", "24"),
				),
			},
		},
	})
}

const testAccDataSourceNxosIPv4PrefixListRuleEntryConfig = `

resource "nxos_ipv4_prefix_list_rule_entry" "test" {
  rule_name = "RULE1"
  order = 10
  action = "permit"
  criteria = "exact"
  prefix = "192.168.1.0/24"
  from_range = 0
  to_range = 24
}

data "nxos_ipv4_prefix_list_rule_entry" "test" {
  rule_name = "RULE1"
  order = 10
  depends_on = [nxos_ipv4_prefix_list_rule_entry.test]
}
`
