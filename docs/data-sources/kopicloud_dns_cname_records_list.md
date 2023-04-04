---
page_title: "kopicloud_dns_cname_records_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_dns_cname_records_list (Data Source)"
---

# kopicloud_dns_cname_records_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List DNS CNAME Records

## Example Usage

List All DNS CNAME Records:
```
data "kopicloud_dns_cname_records_list" "test_cname_all" { }
```

Returns All DNS CNAME Records:
```
output "OUTPUT_dns_cname_records_list_all" {
  description = "List ALL existing DNS CNAME records"
  value       = data.kopicloud_dns_cname_records_list.test_cname_all
}
```

----

Filter DNS CNAME Records with the Zone Name:
```
data "kopicloud_dns_cname_records_list" "test_cname_zone_name" {
  zone_name = "kopicloud.local"
}
```

Returns all DNS CNAME Records that matches the Zone Name:
```
output "OUTPUT_dns_cname_records_list_zone_name" {
  description = "List existing DNS CNAME records in a Zone"
  value       = data.kopicloud_dns_cname_records_list.test_cname_zone_name
}
```

----

Filter DNS CNAME Records with an Alias:
```
data "kopicloud_dns_cname_records_list" "test_cname_alias" {
  hostname_alias = "computer70_alias"
}
```

Returns all DNS CNAME Records that matches the Alias:
```
output "OUTPUT_dns_cname_records_list_ip_address" {
  description = "List existing DNS CNAME Records with the Alias"
  value       = data.kopicloud_dns_cname_records_list.test_cname_alias
}
```

----

Filter DNS CNAME Records with a Computer Hostname:
```
data "kopicloud_dns_cname_records_list" "test_cname_hostname" {
  hostname = "computer33"
}
```

Returns all DNS CNAME Records that matches the Hostname:
```
output "OUTPUT_dns_cname_records_list_hostname" {
  description = "List Existing DNS CNAME Records"
  value       = data.kopicloud_dns_cname_records_list.test_cname_hostname
}
```

## Schema

### Optional

- `hostname` (String) Computer Hostname
- `hostname_alias` (String) Computer Hostname Alias
- `zone_name` (String) DNS Zone Name

### Read-Only

- `id` (String) The ID of this resource
- `result` (List of Objects) List of DNS CNAME Records (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `data` (String) DNS Hostname Alias
- `name` (String) DNS Name
- `timestamp` (String) Timestamp of the Record
- `type` (String) DNS Type
- `zone` (String) DNS Zone Name
