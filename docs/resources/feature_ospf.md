---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_feature_ospf Resource - terraform-provider-nxos"
subcategory: ""
description: |-
  This resource can manage the OSPF feature configuration.
  API Documentation: fmOspf https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospf/
---

# nxos_feature_ospf (Resource)

This resource can manage the OSPF feature configuration.

- API Documentation: [fmOspf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospf/)

## Example Usage

```terraform
resource "nxos_feature_ospf" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **admin_state** (String) Administrative state.
  - Default value: `disabled`

### Read-Only

- **id** (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_feature_ospf.example "sys/fm/ospf"
```