---
page_title: "kopicloud_computer_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_computer_list (Data Source)"
---

# kopicloud_computer_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Active Directory Computers

## Example Usage

List All AD Computers:
```
data "kopicloud_computer_list" "test_all" { }
```

Returns All AD Computers:
```
output "OUTPUT_all_computers_list" {
  description = "List ALL Existing Computers"
  value       = data.kopicloud_computer_list.test_all
}
```

----

List All AD Computers Inside an OU:
```
data "kopicloud_computer_list" "test" {
  ou_path = "OU=Domain Controllers,DC=kopicloud,DC=local"
}
```

Returns All AD Computers Inside an OU:
```
output "OUTPUT_all_computers_list_inside_ou" {
  description = "All Existing Computers Inside an OU"
  value       = data.kopicloud_computer_list.test
}
```

## Schema

### Optional

- `ou_path` (String) AD OU Path (Distinguished Name)

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) List of AD Computers (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `computer_name` (String) AD Computer Name
- `created` (String) Creation Date
- `description` (String) Computer Description
- `dns_name` (String) DNS Name
- `operating_system` (String) Operating System
- `path` (String) AD OU Path (Distinguished Name) 
- `sid` (String) Computer SID
