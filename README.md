# KopiCloud AD API Terraform Provider (Beta)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html)
[![KopiCloud-AD](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

The **KopiCloud AD API Terraform Provider** allows you to manage AD Users, AD Groups, AD OUs, AD Computers, DNS Records, and DNS Zones securely.

## What is KopiCloud AD API?

KopiCloud AD API is a production-ready REST API designed to securely manage Microsoft Active Directory and DNS from custom applications, automation tools, DevOps pipelines, and Terraform.

**KopiCloud AD API** is used in all kinds of organizations to:

- Integrate Microsoft AD and DNS with your applications or scripts.

- Use this Terraform Provider in your CI/CD pipelines without credentials.

- Manage AD Users, AD Groups, AD OUs, and AD Computers.

- Manage DNS Records and DNS Zones.

- Compatible with AWS, Azure, GCP, and OCI VMs running AD DCs.

- Compatible with AWS, Azure, and GCP Active Directory Managed Services.

## Why KopiCloud AD API?

Here are a few reasons to choose KopiCloud AD API to automate your Active Directory deployments:

- **No official Microsoft API**: There is no official Microsoft API, so if you want to automate access to the Active Directory or DNS, you must write your own API or execute PowerShell commands.

- **It is secure**: We use authentication tokens instead of using usernames and passwords to access Active Directory or DNS. These tokens can be used for a limited time or forever.

- **We keep a log of everything**: Every task or action executed is written in a log, so you know who and when they call any API method. Coming soon, you will be able to forward events to several SIEMs.

- **Automate AD with our Terraform Provider**: Create service accounts in AD, create DNS records, clean up orphan machine accounts in AD after a Terraform destroy, reset user passwords, and more.

- **Designed for all kinds of companies**: We have plenty of pre-configured security access groups. The API provides many options if you are a small company or a large enterprise with a dedicated security team.

- **Production or Test Environment?**: Both. If you are in production, every call is secured using a token, and everything is logged. Or you can disable the token authentication if running in a test environment.



## Resources

- KopiCloud AD API Official Web Site: [https://www.kopicloud-ad-api.com](https://www.kopicloud-ad-api.com)

- KopiCloud AD API Documentation: [https://help.kopicloud-ad-api.com](https://help.kopicloud-ad-api.com)

- KopiClod AD API Terraform Sample Code Repo: [https://github.com/KopiCloud-AD-API](https://github.com/KopiCloud-AD-API)
