// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMConfig_all(),
			},
			{
				Config: testAccDataSourceNxosPIMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMConfig = `
data "nxos_pim" "test" {
}
`