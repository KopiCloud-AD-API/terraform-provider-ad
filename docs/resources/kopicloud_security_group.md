---
page_title: "kopicloud_security_group Resource - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD Provider - kopicloud_security_group (Resource)"
---

# kopicloud_security_group (Resource)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

Create, Update, or Delete an Active Directory Security Group

## Example Usage

Create an AD Security Group:
```
resource "kopicloud_security_group" "test_security_global" {
  name        = "kopicloud-europe-security-group"
  scope       = "Global"
  description = "This is a very cool Global security group"
  email       = "europe.security@kopicloud.com"
  ou_path     = "CN=Users,DC=kopicloud,DC=local"
}
```

Returns Created AD Security Group:
```
output "OUTPUT_security_group" {
  description = "Created AD Security Group"
  value       = resource.kopicloud_security_group.test_security_global
}
```

## Notes:

Running this resource with `terraform apply` will create or update the AD group and running `terraform destroy` will remove this AD Group from the Active Directory

## Schema

### Required

- `description` (String) AD Group Description
- `email` (String) AD Group Email Address
- `name` (String) AD Group Name

### Optional

- `ou_path` (String) AD OU Path (Distinguished Name)
- `scope` (String) AD Group Scope, possible values are `Global`, `Universal` or `Domain_Local`. Default is `Global`

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) Single AD Group (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `description` (String) AD Group Description
- `email` (String) AD Group Email Address
- `guid` (String) AD Group GUID
- `name` (String) AD Group Name
- `ou_path` (String) AD Group OU Path (Distinguished Name)
- `scope` (String) AD Group Scope
- `type` (String) AD Group Type 
