---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_pim_vrf Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the PIM VRF configuration.
  API Documentation: pimDom https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:Dom/
---

# nxos_pim_vrf (Resource)

This resource can manage the PIM VRF configuration.

- API Documentation: [pimDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:Dom/)

## Example Usage

```terraform
resource "nxos_pim_vrf" "example" {
  name        = "default"
  admin_state = "enabled"
  bfd         = true
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) VRF name.

### Optional

- **admin_state** (String) Administrative state.
  - Default value: `enabled`
- **bfd** (Boolean) BFD.
  - Default value: `false`

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_pim_vrf.example "sys/pim/inst/dom-[default]"
```