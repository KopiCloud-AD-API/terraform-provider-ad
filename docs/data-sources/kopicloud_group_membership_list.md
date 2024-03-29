---
page_title: "kopicloud_group_membership_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_group_membership_list (Data Source)"
---

# kopicloud_group_membership_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Active Directory Groups Associated with an Active Directory Username

## Example Usage

List AD Group Membership for an AD User:
```
data "kopicloud_group_membership_list" "test" {
  user_name  = "guillermo"
}
```

Show AD Group Membership:
```
output "OUTPUT_kopicloud_group_membership_list" {
  description = "Show Group Membership"
  value       = data.kopicloud_group_membership_list.test
}
```

## Schema

### Required

- `user_name` (String) AD Username

### Optional

- `group_name` (String) AD Group Name

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) List of AD Groups (see [below for nested schema](#nestedatt--result))

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
