terraform {
  required_providers {
    kopicloud = {
      source  = "gitlab.com/kopicloud/kopicloud-ad-tf-provider"
    }
  }
}

provider "kopicloud" {
  host = "http://localhost"
  port    = "3001"
  token   = "superSecretToken"
}