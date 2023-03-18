---
page_title: "KopiCloud AD Provider"
subcategory: ""
description: "Setup the KopiCloud AD Terraform Provider"
  
---

# Set up the KopiCloud AD Terraform Provider
To configure the KopiCloud AD API Terraform provider, log in to your **KopiCloud AD API Management Website** and generate an Authentication Token.

KopiCloud AD uses Authentication Tokens to authenticate to AD instead of usernames and passwords.

Then configure the KopiCloud AD Terraform Provider with the hostname of your API server and the authentication token.

```
terraform {
  required_providers {
    kopicloud = {
      source = "kopicloud-ad-api/ad"
    }
  }
}

provider "kopicloud" {
  host  = "https://api.kopicloud.local"
  token = "Basic b3NjYWI8UzFsdkyQMVsuD70"
}
```

## Schema

- `host` (String) KopiCloud AD API Server URL
- `token` (String) Bearer (JWT) or Basic Authentication Token
