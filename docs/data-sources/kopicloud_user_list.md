---
page_title: "kopicloud_user_list Data Source - KopiCloud AD Provider"
subcategory: ""
description: "KopiCloud AD Provider -kopicloud_user_list (Data Source)"
---

# kopicloud_user_list (Data Source)
[![Terraform](https://img.shields.io/badge/terraform-v1.3+-blue.svg)](https://www.terraform.io/downloads.html) 
[![KopiCloud_AD_API](https://img.shields.io/badge/kopiCloud_ad-v1.0+-blueviolet.svg)](https://www.kopicloud-ad-api.com)

List Users in the Active Directory

## Example Usage

Get All AD Users:
```
data "kopicloud_user_list" "test_all" { }
```

Returns All AD Users:
```
output "OUTPUT_kopicloud_all_users" {
  description = "All Existing AD Users"
  value       = data.kopicloud_user_list.test_all
}
```

----

Get All AD Users Inside an AD OU:
```
data "kopicloud_user_list" "test_ou" {
  ou_path = "LDAP://CN=API Service,CN=Managed Service Accounts,DC=kopicloud,DC=local"
}
```

Returns All AD Users Inside an AD OU:
```
output "OUTPUT_kopicloud_all_users_ou" {
  description = "All Existing AD Users inside an OU"
  value       = data.kopicloud_user_list.test_ou
}
```

----

Get All AD Users Showing Specific Fields:
```
data "kopicloud_user_list" "test_filter" {
  show_fields = "firstName,lastName,displayName"
}
```

Returns all AD Users Showing Specific Fields:
```
output "kopicloud_all_users_filter" {
  description = "All Existing AD Users Showing Specific Fields"
  value       = data.kopicloud_user_list.test_filter
}
```

## Notes

Use the parameter **ShowFields** to select user fields to show. This optional argument is a comma-separated string with the name of the fields you want to be returned
- Example: ShowFields=All or ShowFields=* : Return all User Fields
- Example: ShowFields=Username,Firstname : Return Only Username and Firstname Fields
- Example: ShowFields=Null or Empty : Return Default Fields (Username, Firstname, Lastname, Display_Name, Description, Company, Office, Department, Email_Address)

## Schema

### Optional

- `ou_path` (String) OU Path (Distinguished Name)
- `recursive` (Boolean) Recursive Search Inside the OU
- `show_fields` (String) Filter Specific Fields in the Output

### Read-Only

- `id` (String) The ID of this Resource
- `result` (List of Objects) List of AD Users (see [below for nested schema](#nestedatt--result))

<a id="nestedatt--result"></a>
### Nested Schema for `result`

Read-Only:

- `change_password_next_logon` (Boolean) Status of Change Password Next Logon Setting
- `city` (String) City
- `company` (String) Company Name
- `country` (String) Country (use the English Name of the Country)
- `department` (String) Company Department
- `description` (String) User Description
- `display_name` (String) User Display Name
- `email_address` (String) User Email Address
- `first_name` (String) User First Name
- `guid` (String) User GUID
- `home_folder_directory` (String) Home Folder Directory
- `home_folder_drive` (String) Home Folder Drive
- `home_folder_path` (String) Home Folder Path
- `home_phone` (String) Home Phone
- `initials` (String) User Middle Name Initial
- `job_title` (String) Job Title
- `last_name` (String) User Last Name
- `manager` (String) - Manager Name
- `mobile_phone` (String) Mobile Phone
- `office` (String) Office Information
- `office_phone` (String) Office Phone
- `ou_path` (String) OU Path (Distinguished Name) of the User
- `password_never_expired` (Boolean) Status of the Password Never Expired Setting
- `postal_code` (String) Postal/ZIP Code
- `profile_logon_script` (String) Profile Logon Script
- `profile_path` (String) Profile Path
- `rds_allow_logon` (Boolean) Remote Desktop Allow Logon
- `rds_connect_drive` (Boolean) Remote Desktop Connect Drive
- `rds_home_folder_drive` (String) Remote Desktop Home Folder
- `rds_home_folder_path` (String) Remote Desktop Folder Path
- `rds_profile_path` (String) Remote Desktop Profile Path
- `sam_username` (String) SAM Account Name (Used by Previous Versions of Windows)
- `state` (String) State or Province
- `street_address` (String) Street Address
- `street_po_box` (String) PO Box Address
- `username` (String) Username
