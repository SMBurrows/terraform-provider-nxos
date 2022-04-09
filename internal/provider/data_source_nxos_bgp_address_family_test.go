// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosBGPAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosBGPAddressFamilyPrerequisitesConfig + testAccDataSourceNxosBGPAddressFamilyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_bgp_address_family.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("data.nxos_bgp_address_family.test", "critical_nexthop_timeout", "1800"),
					resource.TestCheckResourceAttr("data.nxos_bgp_address_family.test", "non_critical_nexthop_timeout", "1800"),
				),
			},
		},
	})
}

const testAccDataSourceNxosBGPAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosBGPAddressFamilyConfig = `

resource "nxos_bgp_address_family" "test" {
  vrf = "default"
  address_family = "ipv4-ucast"
  critical_nexthop_timeout = 1800
  non_critical_nexthop_timeout = 1800
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_bgp_address_family" "test" {
  vrf = "default"
  address_family = "ipv4-ucast"
  depends_on = [nxos_bgp_address_family.test]
}
`
