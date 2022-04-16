---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nxos_hmm Data Source - terraform-provider-nxos"
subcategory: "HMM"
description: |-
  This data source can read the Host Mobility Manager (HMM) Entity configuration.
  API Documentation: hmmEntity https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:Entity/
---

# nxos_hmm (Data Source)

This data source can read the Host Mobility Manager (HMM) Entity configuration.

- API Documentation: [hmmEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:Entity/)

## Example Usage

```terraform
data "nxos_hmm" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `device` (String) A device name from the provider configuration.

### Read-Only

- `admin_state` (String) Administrative state.
- `id` (String) The distinguished name of the object.

