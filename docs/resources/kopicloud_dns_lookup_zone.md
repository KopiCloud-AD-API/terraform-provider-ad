---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "kopicloud_dns_lookup_zone Resource - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD Provider - kopicloud_dns_lookup_zone (Resource)"
---

# kopicloud_dns_lookup_zone (Resource)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

Create or Delete a DNS Lookup Zone in the Microsoft DNS

## Example Usage

Create a DNS Lookup Zone:
```
resource "kopicloud_dns_lookup_zone" "test_lookup" {
  zone_name = "kopicloud.org"
}
```

Returns Created DNS Lookup Zone:
```
output "OUTPUT_dns_lookup_zone" {
  description = "Created DNS Lookup Zone"
  value       = resource.kopicloud_dns_lookup_zone.test_lookup
}
```

## Notes 

Running this resource with `terraform apply` will create a DNS Lookup Zone in the Microsoft DNS and running `terraform destroy` will remove the DNS Lookup Zone from the DNS

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `zone_name` (String) DNS Zone Name

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) Single DNS Zone (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `distinguished_name` (String) DNS Distinguished Name
- `type` (String) DNS Type, possible values are `ForwardDNSZone` or `ReverseDNSZone`
- `zone_name` (String) DNS Zone Name
- `zone_type` (String) DNS Zone Type, possible values are `Primary`, `Secondary` or `Stub Zone`
