---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_route_map_rule Resource - terraform-provider-nxos"
subcategory: "Routing"
description: |-
  This resource can manage a Route-Map Rule configuration.
  API Documentation: rtmapRule https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Rule/
  Child resources
  nxosroutemapruleentry https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/route_map_rule_entry
---

# nxos_route_map_rule (Resource)

This resource can manage a Route-Map Rule configuration.

- API Documentation: [rtmapRule](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Rule/)

### Child resources

- [nxos_route_map_rule_entry](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/route_map_rule_entry)

## Example Usage

```terraform
resource "nxos_route_map_rule" "example" {
  name = "RULE1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Route-Map Rule name.

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_route_map_rule.example "sys/rpm/rtmap-[RULE1]"
```
