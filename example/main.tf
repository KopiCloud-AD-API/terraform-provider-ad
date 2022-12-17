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

data "kopicloud_all_computers" "test" {}

# Returns all computers
output "all_computers_list" {
  description = "Existing computers"
  value = data.kopicloud_all_computers.test
}

data "kopicloud_all_ous" "test" {}

# Returns all ous
output "all_ous_list" {
  description = "Existing OUs"
  value = data.kopicloud_all_ous.test
}
