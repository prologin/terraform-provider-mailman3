---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "mailman3_domain Resource - terraform-provider-mailman3"
subcategory: ""
description: |-
  This resource can be used to manage mailman domains.
---

# mailman3_domain (Resource)

This resource can be used to manage mailman domains.

## Example Usage

```terraform
resource "mailman3_domain" "my_domain" {
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


