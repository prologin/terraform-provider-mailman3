---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mailman3_domain Data Source - terraform-provider-mailman3"
subcategory: ""
description: |-
  This data source can be used to retrieve information about a domain.
---

# mailman3_domain (Data Source)

This data source can be used to retrieve information about a domain.

## Example Usage

```terraform
data "mailman3_domain" "my_domain" {
  mail_host = "ml.example.org"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `mail_host` (String) The mail host for this domain.

### Optional

- `alias_domain` (String)
- `description` (String)

### Read-Only

- `id` (String) The ID of this resource.


