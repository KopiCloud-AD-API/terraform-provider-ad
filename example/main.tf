terraform {
  required_providers {
    kopicloud = {
      source  = "gitlab.com/kopicloud/kopicloud-ad-tf-provider"
    }
  }
}

provider "kopicloud" {
  host = "https://labapi.kopicloud-ad-api.com"
  token   = "Basic b3NjYXI6UzFsdjNyQnVsbDN0"
}

# data "kopicloud_computers" "test" {
#   oupath= "OU=Domain Controllers,DC=kopicloud,DC=local"
#   recursive=true
# }

# # Returns all computers
# output "computers_list" {
#   description = "Existing computers"
#   value = data.kopicloud_computers.test
# }


data "kopicloud_computer_list" "test" {}

# Returns all computers
output "all_computers_list" {
  description = "Existing computers"
  value = data.kopicloud_computer_list.test.result
}

data "kopicloud_computer_list" "by_ou" {
  ou_path = "OU=Domain Controllers,DC=kopicloud,DC=local"
}

# Returns all computers
output "ou_computers_list" {
  description = "Existing computers"
  value = data.kopicloud_computer_list.by_ou.result
}

data "kopicloud_computer_list" "by_name" {
  ad_computer_name = "BIG-WOPER"
}

# Returns all computers
output "by_name_computers_list" {
  description = "Existing computers"
  value = data.kopicloud_computer_list.by_name.result
}


# data "kopicloud_all_ad_groups" "test" {}

# # Returns all ADGroups
# output "kopicloud_all_ad_groups" {
#   description = "Existing AD Groups"
#   value = data.kopicloud_all_ad_groups.test
# }

# data "kopicloud_ad_group_membership" "test" {
#   username = "guillermo"
# }

# # Returns all ADGroups for a User
# output "kopicloud_ad_group_membership" {
#   description = "AD Groups for user"
#   value = data.kopicloud_ad_group_membership.test
# }


# data "kopicloud_all_ad_users" "test" {}

# # Returns all ADUsers
# output "kopicloud_all_ad_users" {
#   description = "Existing AD Users"
#   value = data.kopicloud_all_ad_users.test
# }

# data "kopicloud_all_ous" "test" {}

# # Returns all ous
# output "all_ous_list" {
#   description = "Existing OUs"
#   value = data.kopicloud_all_ous.test
# }


data "kopicloud_dns_a_records_list" "by_ip" {
  ip_address = "10.20.1.241"
}

# Returns all DNS A Records
output "by_ip_dns_a_records" {
  description = "Existing DNS A Records"
  value = data.kopicloud_dns_a_records_list.by_ip
}

# data "kopicloud_dns_a_records_list" "by_zone" {
#   zone_name = "kopicloud.local"
# }

# # Returns all DNS A Records
# output "by_zone_dns_a_records" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.by_zone
# }


# data "kopicloud_dns_a_records_list" "one" {
#   hostname = "computer1975"
#   ip_address = "10.20.1.241"
#   zone_name = "kopicloud.local"
# }

# # Returns all DNS A Records
# output "some_dns_a_records" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.one
# }

# data "kopicloud_dns_a_records_list" "all" {}

# # Returns all DNS A Records
# output "all_dns_a_records" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.all
# }

data "kopicloud_dns_aaaa_records_list" "test" {}

# Returns all DNS AAAA Records
output "all_dns_aaaa_records" {
  description = "Existing DNS AAAA Records"
  value = data.kopicloud_dns_aaaa_records_list.test.result
}

data "kopicloud_dns_cname_records_list" "test" {}

# Returns all DNS CNAME Records
output "all_dns_cname_records" {
  description = "Existing DNS CName Records"
  value = data.kopicloud_dns_cname_records_list.test.result
}

# data "kopicloud_all_dns_zones" "test" {}

# # Returns all DNS CNAME Records
# output "all_dns_zones" {
#   description = "Existing DNS Zones"
#   value = data.kopicloud_all_dns_zones.test
# }

resource "kopicloud_computer" "test" {
  ad_computer_name = "BIG-WOPER"
  ou_path= "OU=Domain Controllers,DC=kopicloud,DC=local"    
}

output "computer" {
  description = "Created Computer"
  value = resource.kopicloud_computer.test
}


# resource "kopicloud_dns_a_record" "test" {
#   hostname = "computer1975"
#   ip_address = "10.20.1.241"
#   zone_name = "kopicloud.local"
# }

# output "dns_a_record" {
#   description = "Created DNS A Record"
#   value = resource.kopicloud_dns_a_record.test
# }

# resource "kopicloud_dns_aaaa_record" "test" {
#   hostname = "computer-AAAA-5"
#   ipv6_address = "2001:db8:3333:4444:5555:6666:7777:8889"
#   zone_name = "kopicloud.local"
# }

# output "dns_aaaa_record" {
#   description = "Created DNS AAAA Record"
#   value = resource.kopicloud_dns_aaaa_record.test
# }

# resource "kopicloud_dns_cname_record" "test" {
#   hostname = "computer999"
#   hostname_alias = "computer70_alias"
#   zone_name = "kopicloud.local"
# }

# output "dns_cname_record" {
#   description = "Created DNS CName Record"
#   value = resource.kopicloud_dns_cname_record.test.result
# }

# resource "kopicloud_dns_lookup_zone" "test" {
#   zone_name = "my_lookup_zone_3"
# }

# output "dns_lookup_zone" {
#   description = "Created DNS Lookup zone"
#   value = resource.kopicloud_dns_lookup_zone.test
# }

# resource "kopicloud_dns_reverse_lookup_zone" "test" {
#   zone_name = "my_reverse_lookup_zone_4"
# }

# output "dns_reverse_lookup_zone" {
#   description = "Created DNS Reverse Lookup zone"
#   value = resource.kopicloud_dns_reverse_lookup_zone.test
# }

# data "kopicloud_dns_a_records_list" "test_hostname" { 
#   hostname = "tito2" 
# }

# output "kopicloud_dns_a_records_list" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.test_hostname.result
# }
