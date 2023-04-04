---

page_title: "kopicloud_dns_aaaa_records_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD API - kopicloud_dns_aaaa_records_list (Data Source)"
---

# kopicloud_dns_aaaa_records_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List DNS AAAA Records

## Example Usage

List All DNS AAAA Records:
```
data "kopicloud_dns_aaaa_records_list" "test_aaaa_all" {}
```

Returns All DNS AAAA Records:
```
output "OUTPUT_dns_aaaa_records_list_all" {
  description = "List ALL existing DNS AAAA Records"
  value       = data.kopicloud_dns_aaaa_records_list.test_aaaa_all
}
```

----

Filter DNS AAAA Records with the Zone Name:
```
data "kopicloud_dns_aaaa_records_list" "test_aaaa_zone_name" {
  zone_name = "kopicloud.local"
}
```

Returns all DNS AAAA Records that matches the Zone Name:
```
output "OUTPUT_dns_aaaa_records_list_zone_name" {
  description = "List existing DNS AAAA Records filtered by Zone Name"
  value       = data.kopicloud_dns_aaaa_records_list.test_aaaa_zone_name
}
```

----

Filter DNS AAAA Records with an IPv6 Address:
```
data "kopicloud_dns_aaaa_records_list" "test_aaaa_ip" {
  ipv6_address = "2001:db8:3333:4444:5555:6666:7777:8888"
}
```

Returns all DNS AAAA Records that matches the IPv6 Address:
```
output "OUTPUT_dns_aaaa_records_list_ip_address" {
  description = "List existing DNS AAAA Records filtered by IPv6 Address"
  value       = data.kopicloud_dns_aaaa_records_list.test_aaaa_ip
}
```

----

Filter DNS AAAA Records with a Hostname:
```
data "kopicloud_dns_aaaa_records_list" "test_aaaa_hostname" {
  hostname = "computer70v6"
}
```

Returns all DNS AAAA Records that matches the Hostname:
```
output "OUTPUT_dns_aaaa_records_list_hostname" {
  description = "List existing DNS AAAA Records filtered by Hostname"
  value       = data.kopicloud_dns_aaaa_records_list.test_aaaa_hostname
}
```

## Schema

### Optional

- `hostname` (String) Computer Hostname
- `ipv6_address` (String) IPv6 Address
- `zone_name` (String) DNS Zone Name

### Read-Only

- `id` (String) The ID of this resource
- `result` (List of Objects) List of DNS AAAA Records (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `data` (String) IPv6 Address
- `name` (String) Computer Hostname
- `timestamp` (String) Timestamp of the Record
- `type` (String) DNS Type
- `zone` (String) DNS Zone
