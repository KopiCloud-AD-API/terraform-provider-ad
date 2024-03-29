---
page_title: "kopicloud_distribution_group_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_distribution_group_list (Data Source)"
---

# kopicloud_distribution_group_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Active Directory Distribution Groups

## Example Usage

Get the List of AD Distribution Groups:
```
data "kopicloud_distribution_group_list" "test_distribution" { }
```

Returns the List of AD Distribution Groups:
```
output "OUTPUT_kopicloud_distribution_groups_list" {
  description = "All Existing Distribution Groups"
  value       = data.kopicloud_distribution_group_list.test_distribution
}
```

## Schema

### Optional

- `name` (String) AD Group Name
- `ou_path` (String) AD OU Path (Distinguished Name)

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objectss) List of AD Groups (see [below for nested schema](#nestedatt--result))

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
