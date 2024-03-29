---
page_title: "kopicloud_security_group_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_security_group_list (Data Source)"  
---

# kopicloud_security_group_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Active Directory Security Groups

## Example Usage

Get the List of AD Security Groups:
```
data "kopicloud_security_group_list" "test_security" { }
```

Returns the List of AD Security Groups:
```
output "OUTPUT_kopicloud_security_groups_list" {
  description = "All Existing Security Groups"
  value       = data.kopicloud_security_group_list.test_security
}
```

## Schema

### Optional

- `name` (String) AD Group Name
- `ou_path` (String) AD OU Path (Distinguished name)

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
