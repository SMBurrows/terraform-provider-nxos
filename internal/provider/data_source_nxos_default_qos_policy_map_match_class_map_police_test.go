// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPolice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapConfig_all(),
			},
			{
				Config: testAccNxosDefaultQOSClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapConfig_all() + testAccNxosDefaultQOSPolicyMapMatchClassMapPoliceConfig_all(),
			},
			{
				Config: testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPoliceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "bc_rate", "200"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "bc_unit", "mbytes"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "be_rate", "200"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "be_unit", "mbytes"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "cir_rate", "10000"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "cir_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_action", "transmit"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "conform_set_qos_group", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_action", "transmit"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "exceed_set_qos_group", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "pir_rate", "10000"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "pir_unit", "mbps"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_action", "drop"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_cos", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_dscp", "0"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_precedence", "routine"),
					resource.TestCheckResourceAttr("data.nxos_default_qos_policy_map_match_class_map_police.test", "violate_set_qos_group", "0"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDefaultQOSPolicyMapMatchClassMapPoliceConfig = `
data "nxos_default_qos_policy_map_match_class_map_police" "test" {
  policy_map_name = "PM1"
  class_map_name = "Voice"
}
`