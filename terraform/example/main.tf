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

data "kopicloud_computers" "test" {
  oupath= "OU=Domain Controllers,DC=kopicloud,DC=local"
  recursive=true
}

# Returns all computers
output "computers_list" {
  description = "Existing computers"
  value = data.kopicloud_computers.test
}


data "kopicloud_all_computers" "test" {}

# Returns all computers
output "all_computers_list" {
  description = "Existing computers"
  value = data.kopicloud_all_computers.test
}

data "kopicloud_all_ad_groups" "test" {}

# Returns all ADGroups
output "kopicloud_all_ad_groups" {
  description = "Existing AD Groups"
  value = data.kopicloud_all_ad_groups.test
}

data "kopicloud_ad_group_membership" "test" {
  username = "guillermo"
}

# Returns all ADGroups for a User
output "kopicloud_ad_group_membership" {
  description = "AD Groups for user"
  value = data.kopicloud_ad_group_membership.test
}


data "kopicloud_all_ad_users" "test" {}

# Returns all ADUsers
output "kopicloud_all_ad_users" {
  description = "Existing AD Users"
  value = data.kopicloud_all_ad_users.test
}

data "kopicloud_all_ous" "test" {}

# Returns all ous
output "all_ous_list" {
  description = "Existing OUs"
  value = data.kopicloud_all_ous.test
}

data "kopicloud_all_dns_a_records" "test" {}

# Returns all DNS A Records
output "all_dns_a_records" {
  description = "Existing DNS A Records"
  value = data.kopicloud_all_dns_a_records.test
}

data "kopicloud_all_dns_aaaa_records" "test" {}

# Returns all DNS AAAA Records
output "all_dns_aaaa_records" {
  description = "Existing DNS AAAA Records"
  value = data.kopicloud_all_dns_aaaa_records.test
}

data "kopicloud_all_dns_cname_records" "test" {}

# Returns all DNS CNAME Records
output "all_dns_cname_records" {
  description = "Existing DNS AAAA Records"
  value = data.kopicloud_all_dns_cname_records.test
}

data "kopicloud_all_dns_zones" "test" {}

# Returns all DNS CNAME Records
output "all_dns_zones" {
  description = "Existing DNS Zones"
  value = data.kopicloud_all_dns_zones.test
}
