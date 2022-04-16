// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosHMM(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosHMMPrerequisitesConfig + testAccDataSourceNxosHMMConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_hmm.test", "admin_state", "enabled"),
				),
			},
		},
	})
}

const testAccDataSourceNxosHMMPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/hmm"
  class_name = "fmHmm"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  content = {
      adminSt = "enabled"
  }
}

`

const testAccDataSourceNxosHMMConfig = `

resource "nxos_hmm" "test" {
  admin_state = "enabled"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_hmm" "test" {
  depends_on = [nxos_hmm.test]
}
`