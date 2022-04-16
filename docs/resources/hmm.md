---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_hmm Resource - terraform-provider-nxos"
subcategory: "HMM"
description: |-
  This resource can manage the Host Mobility Manager (HMM) Entity configuration.
  API Documentation: hmmEntity https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:Entity/
  Child resources
  nxoshmminstance https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/hmm_instance
---

# nxos_hmm (Resource)

This resource can manage the Host Mobility Manager (HMM) Entity configuration.

- API Documentation: [hmmEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:Entity/)

### Child resources

- [nxos_hmm_instance](https://registry.terraform.io/providers/netascode/nxos/latest/docs/resources/hmm_instance)

## Example Usage

```terraform
resource "nxos_hmm" "example" {
  admin_state = "enabled"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `admin_state` (String) Administrative state.
  - Choices: `enabled`, `disabled`
  - Default value: `enabled`
- `device` (String) A device name from the provider configuration.

### Read-Only

- `id` (String) The distinguished name of the object.

## Import

Import is supported using the following syntax:

```shell
terraform import nxos_hmm.example "sys/hmm"
```