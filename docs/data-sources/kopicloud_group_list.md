---
page_title: "kopicloud_group_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_group_list (Data Source)"
---

# kopicloud_group_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List all AD Groups

## Example Usage

Get All AD Groups:
```
data "kopicloud_group_list" "test_all" { }
```

Returns All AD Groups:
```
output "OUTPUT_kopicloud_all_groups" {
  description = "All Existing AD Groups"
  value       = data.kopicloud_group_list.test_all
}
```

## Schema

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
- `ou_path` (String) = AD Group OU Path (Distinguished Name)
- `scope` (String) AD Group Scope
- `type` (String) AD Group Type 
