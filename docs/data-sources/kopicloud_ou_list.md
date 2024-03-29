---
page_title: "kopicloud_ou_list Data Source - KopiCloud AD Provider"
subcategory: ""
description:   "KopiCloud AD API - kopicloud_ou_list (Data Source)"
---

# kopicloud_ou_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Active Directory Organization Units (OUs)

## Example Usage

List AD OUs:

```
data "kopicloud_ou_list" "test" { }
```

Returns List of AD OUs:

```
output "OUTPUT_list_ou" {
  description = "List of Existing AD OUs"
  value       = data.kopicloud_ou_list.test
}
```

## Schema

### Optional

- `ou_name` (String) Name of the AD OU
- `ou_path` (String) Path of the AD OU (Distinguished Name)

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) List of AD OUs (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `description` (String) The Description of the AD OU
- `guid` (String) The GUID of the AD OU
- `name` (String) Name of the AD OU
- `path` (String) Path of the AD OU (Distinguished Name)
- `protected` (Boolean) Protect the AD OU from Accidental Deletion
