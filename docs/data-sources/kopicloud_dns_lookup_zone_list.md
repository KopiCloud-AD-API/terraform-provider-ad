---
page_title: "kopicloud_dns_lookup_zone_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_dns_lookup_zone_list (Data Source)"
---

# kopicloud_dns_lookup_zone_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List DNS Lookup Zones

## Example Usage

List All DNS Lookup Zones:
```
data "kopicloud_dns_lookup_zone_list" "test_lookup_all" { }
```

Returns List of DNS Lookup Zones:
```
output "OUTPUT_dns_lookup_zone_list" {
  description = "List of DNS Lookup Zones"
  value       = data.kopicloud_dns_lookup_zone_list.test_lookup_all
}
```

## Schema

### Optional

- `zone_name` (String) DNS Zone Name

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) List of DNS Zones (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `distinguished_name` (String) DNS Distinguished Name
- `type` (String) DNS Type, possible values are `ForwardDNSZone` or `ReverseDNSZone`
- `zone_name` (String) DNS Zone Name
- `zone_type` (String) DNS Zone Type, possible values are `Primary`, `Secondary` or `Stub Zone`

