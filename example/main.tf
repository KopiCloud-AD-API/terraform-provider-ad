terraform {
  required_providers {
    kopicloud = {
      source = "kopicloud-ad-api/ad"
      # source  = "github.com/KopiCloud-AD-API/ad"
    }
  }
}

provider "kopicloud" {
  host  = "https://labapi.kopicloud-ad-api.com"
  token = var.token
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


# data "kopicloud_computer_list" "test" {}

# # Returns all computers
# output "all_computers_list" {
#   description = "Existing computers"
#   value = data.kopicloud_computer_list.test.result
# }

# data "kopicloud_ou_list" "test" {}

# # Returns all computers
# output "all_ous_list" {
#   description = "Existing OUs"
#   value = data.kopicloud_ou_list.test.result
# }


# data "kopicloud_computer_list" "by_ou" {
#   ou_path = "OU=Domain Controllers,DC=kopicloud,DC=local"
# }

# # Returns all computers
# output "ou_computers_list" {
#   description = "Existing computers"
#   value = data.kopicloud_computer_list.by_ou.result
# }

# data "kopicloud_computer_list" "by_name" {
#   ad_computer_name = "BIG-WOPER"
# }

# # Returns all computers
# output "by_name_computers_list" {
#   description = "Existing computers"
#   value = data.kopicloud_computer_list.by_name.result
# }


# data "kopicloud_group_list" "test" {}

# # Returns all ADGroups
# output "kopicloud_all_groups" {
#   description = "Existing Groups"
#   value = data.kopicloud_group_list.test
# }

# data "kopicloud_security_group_list" "test" {}

# # Returns all ADGroups
# output "kopicloud_all_security_groups" {
#   description = "Existing Security Groups"
#   value = data.kopicloud_security_group_list.test
# }

# data "kopicloud_distribution_group_list" "test" {}

# # Returns all ADGroups
# output "kopicloud_all_distribution_groups" {
#   description = "Existing Distribution Groups"
#   value = data.kopicloud_distribution_group_list.test
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


# data "kopicloud_dns_a_records_list" "by_ip" {
#   ip_address = "10.20.1.241"
# }

# # Returns all DNS A Records
# output "by_ip_dns_a_records" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.by_ip
# }

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

# data "kopicloud_dns_aaaa_records_list" "test" {}

# # Returns all DNS AAAA Records
# output "all_dns_aaaa_records" {
#   description = "Existing DNS AAAA Records"
#   value = data.kopicloud_dns_aaaa_records_list.test.result
# }

# data "kopicloud_dns_cname_records_list" "test" {}

# # Returns all DNS CNAME Records
# output "all_dns_cname_records" {
#   description = "Existing DNS CName Records"
#   value = data.kopicloud_dns_cname_records_list.test.result
# }

# data "kopicloud_all_dns_zones" "test" {}

# # Returns all DNS CNAME Records
# output "all_dns_zones" {
#   description = "Existing DNS Zones"
#   value = data.kopicloud_all_dns_zones.test
# }

# resource "kopicloud_computer" "test" {
#   ad_computer_name = "BIG-WOPER"
#   ou_path= "OU=Domain Controllers,DC=kopicloud,DC=local"
# }

# output "computer" {
#   description = "Created Computer"
#   value = resource.kopicloud_computer.test
# }


# resource "kopicloud_distribution_group" "test" {
#   name      = "my-new-distro-group"
#   scope     = "Global"
#   description  = "This is a very cool distribution group"
#   email = "distribution@kopicloud.com"
#   ou_path      = "OU=Domain Controllers,DC=kopicloud,DC=local"
# }

# output "distribution_group" {
#   description = "Created Security Group"
#   value = resource.kopicloud_distribution_group.test
# }

# resource "kopicloud_ou" "test" {
#   ou_name      = "kopicloud-13"
#   ou_path      = "OU=Domain Controllers,DC=kopicloud,DC=local"
#   description  = "This is a very cool organization with improved description"
#   protected    = false
# }

# output "ou" {
#   description = "Created OU"
#   value = resource.kopicloud_ou.test
# }

# resource "kopicloud_dns_a_record" "test" {
#   hostname = "computer1978"
#   ip_address = "10.20.1.241"
#   zone_name = "kopicloud.local"
# }

# output "dns_a_record" {
#   description = "Created DNS A Record"
#   value = resource.kopicloud_dns_a_record.test
# }

# resource "kopicloud_dns_aaaa_record" "test" {
#   hostname = "computer-AAAA-6"
#   ipv6_address = "2001:db8:3333:4444:5555:6666:7777:8889"
#   zone_name = "kopicloud.local"
# }

# output "dns_aaaa_record" {
#   description = "Created DNS AAAA Record"
#   value = resource.kopicloud_dns_aaaa_record.test
# }

# resource "kopicloud_dns_cname_record" "test" {
#   hostname = "computer1969"
#   hostname_alias = "computer1969  _alias"
#   zone_name = "kopicloud.local"
# }

# output "dns_cname_record" {
#   description = "Created DNS CName Record"
#   value = resource.kopicloud_dns_cname_record.test.result
# }

# resource "kopicloud_dns_lookup_zone" "test" {
#   zone_name = "my-zone-returns"
# }

# output "dns_lookup_zone" {
#   description = "Created DNS Lookup zone"
#   value = resource.kopicloud_dns_lookup_zone.test
# }

# resource "kopicloud_dns_reverse_lookup_zone" "test" {
#   network_id = "192.169.50.0/24"
# }

# output "dns_reverse_lookup_zone" {
#   description = "Created DNS Reverse Lookup zone"
#   value = resource.kopicloud_dns_reverse_lookup_zone.test
# }

# data "kopicloud_dns_lookup_zone_list" "test" {}

# # Returns all computers
# output "all_dns_looup_zones_list" {
#   description = "Existing Lookup Zones"
#   value = data.kopicloud_dns_lookup_zone_list.test.result
# }

# data "kopicloud_dns_reverse_lookup_zone_list" "test" {}

# # Returns all computers
# output "all_dns_reverse_looup_zones_list" {
#   description = "Existing Reverse Lookup Zones"
#   value = data.kopicloud_dns_reverse_lookup_zone_list.test.result
# }


# data "kopicloud_dns_zone_list" "test" {}

# # Returns all computers
# output "all_dns_zones_list" {
#   description = "Existing Reverse Lookup Zones"
#   value = data.kopicloud_dns_zone_list.test.result
# }


# data "kopicloud_dns_a_records_list" "test_hostname" {
#   hostname = "tito2"
# }

# output "kopicloud_dns_a_records_list" {
#   description = "Existing DNS A Records"
#   value = data.kopicloud_dns_a_records_list.test_hostname.result
# }

# data "kopicloud_group_membership_list" "test" {
#   user_name = "guillermo"
# }

# # Returns all ADGroups for a User
# output "kopicloud_group_membership_list" {
#   description = "AD Groups for user"
#   value = data.kopicloud_group_membership_list.test.result
# }

# resource "kopicloud_computer_cleanup" "test" {
#   days = 1
#   ou_path = "OU=Domain Controllers,DC=kopicloud,DC=local"
#   recursive = true
# }

# output "computer_cleanup" {
#   description = "Computers To CleanUp"
#   value = resource.kopicloud_computer_cleanup.test
# }

# data "kopicloud_user_list" "test" {
#   show_fields = "firstName,lastName,displayName"
# }

# # Returns all ADUsers
# output "kopicloud_all_users" {
#   description = "Existing AD Users"
#   value = data.kopicloud_user_list.test
# }

resource "kopicloud_user" "test" {
  username     = "oforero15"
  password     = "vMD48X6Vdyj49j8%"
  first_name   = "Oscar"
  last_name    = "Forero"
  initials     = "OMFC"
  company = "Kopicloud Inc."
}

output "user" {
  description = "Created User"
  value = resource.kopicloud_user.test.result
}

# resource "kopicloud_user_password_reset" "test" {
#   username     = "oforero11"
#   new_password = "vMD48X6Vdyj69j8999%"
# }

# output "user_password_reset" {
#   description = "Password Resetted for User"
#   value = resource.kopicloud_user_password_reset.test.result
# }

# # Create a Global Distribution Group
# resource "kopicloud_distribution_group" "test_distribution_global" {
#   name        = "kopicloud-europe-distribution-group-blue-3"
#   scope       = "Global"
#   description = "This is a very cool Global distribution group"
#   email       = "europe.distribution@kopicloud.com"
#   ou_path     = "CN=Users,DC=kopicloud,DC=local"
# }

# # Returns Created Global Distribution Group
# output "OUTPUT_global_distribution_group" {
#   description = "Created Global Distribution Group"
#   value       = resource.kopicloud_distribution_group.test_distribution_global
# }

# # Create a Global Security Group
# resource "kopicloud_security_group" "test_security_global" {
#   name        = "kopicloud-europe-security-group-red-1"
#   scope       = "Global"
#   description = "This is a very cool Global security group"
#   email       = "europe.security@kopicloud.com"
#   ou_path     = "CN=Users,DC=kopicloud,DC=local"
# }

# # Returns Created Global Security Group
# output "OUTPUT_global_security_group" {
#   description = "Created Global Security Group"
#   value       = resource.kopicloud_security_group.test_security_global
# }
