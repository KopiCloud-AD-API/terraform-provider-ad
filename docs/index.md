---
page_title: "KopiCloud AD API Terraform Provider"
subcategory: ""
description: "KopiCloud AD API Terraform Provider"
---

# KopiCloud AD API Terraform Provider

**KopiCloud AD API** is a production-ready REST API designed to securely manage **Microsoft Active Directory (AD)** and **Microsoft DNS** from custom applications, automation tools, DevOps pipelines, and Terraform.

Instead of using usernames and passwords with WinRM to access Active Directory or DNS, we use an API with authentication tokens. WinRM is not required. These tokens can be used for a limited time or forever.

Every task or action executed is written in a log, so you know who and when they call any Terraform action.

## Terraform Provider Authentication
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

To set up the **KopiCloud AD Terraform Provider**, log in to the KopiCloud AD API Management Portal and generate a token.

In your Terraform provider file (for example, provider.tf), configure the KopiCloud AD Terraform Provider with the hostname of your API server and the authentication token.

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

## Parameters

- `host` (String) KopiCloud AD API Server URL
- `token` (String) Bearer (JWT) or Basic Authentication Token


## Resources

- KopiCloud AD API Official Web Site: https://adapi.kopicloud.com

- KopiCloud AD API Documentation: https://adapihelp.kopicloud.com

- KopiClod AD API Terraform Sample Code Repo: https://github.com/KopiCloud-AD-API
