package provider

import (
	api "github.com/KopiCloud-AD-API/terraform-provider-ad/api"
)

func singleStringToTerraform(field_name string) func(*string) map[string]interface{} {
	return func(value *string) map[string]interface{} {
		result := make(map[string]interface{})
		result[field_name] = value
		return result
	}
}

func stringListToTerraform(field_name string) func(*[]string) []interface{} {
	return func(list *[]string) []interface{} {
		if list != nil {
			results := make([]interface{}, len(*list))
			for i, value := range *list {
				results[i] = singleStringToTerraform(field_name)(&value)
			}
			return results
		}
		return make([]interface{}, 0)
	}
}

func isNotEmpty(f interface{}) bool {
	switch v := f.(type) {
	case string:
		return len(v) > 0
	default:
		return true
	}
}

func UserToTerraform(obj *api.User) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.City != nil && isNotEmpty(obj.City) {

		result["city"] = obj.City

	}

	if obj.ProfileLogonScript != nil && isNotEmpty(obj.ProfileLogonScript) {

		result["profile_logon_script"] = obj.ProfileLogonScript

	}

	if obj.HomeFolderDrive != nil && isNotEmpty(obj.HomeFolderDrive) {

		result["home_folder_drive"] = obj.HomeFolderDrive

	}

	if obj.RdsHomeFolderDrive != nil && isNotEmpty(obj.RdsHomeFolderDrive) {

		result["rds_home_folder_drive"] = obj.RdsHomeFolderDrive

	}

	if obj.Guid != nil && isNotEmpty(obj.Guid) {

		result["guid"] = UuidToTerraform(obj.Guid)

	}

	if obj.JobTitle != nil && isNotEmpty(obj.JobTitle) {

		result["job_title"] = obj.JobTitle

	}

	if obj.HomeFolderPath != nil && isNotEmpty(obj.HomeFolderPath) {

		result["home_folder_path"] = obj.HomeFolderPath

	}

	if obj.RdsHomeFolderPath != nil && isNotEmpty(obj.RdsHomeFolderPath) {

		result["rds_home_folder_path"] = obj.RdsHomeFolderPath

	}

	if obj.RdsAllowLogon != nil && isNotEmpty(obj.RdsAllowLogon) {

		result["rds_allow_logon"] = obj.RdsAllowLogon

	}

	if obj.Company != nil && isNotEmpty(obj.Company) {

		result["company"] = obj.Company

	}

	if obj.StreetAddress != nil && isNotEmpty(obj.StreetAddress) {

		result["street_address"] = obj.StreetAddress

	}

	if obj.State != nil && isNotEmpty(obj.State) {

		result["state"] = obj.State

	}

	if obj.RdsProfilePath != nil && isNotEmpty(obj.RdsProfilePath) {

		result["rds_profile_path"] = obj.RdsProfilePath

	}

	if obj.LastName != nil && isNotEmpty(obj.LastName) {

		result["last_name"] = obj.LastName

	}

	if obj.Office != nil && isNotEmpty(obj.Office) {

		result["office"] = obj.Office

	}

	if obj.MobilePhone != nil && isNotEmpty(obj.MobilePhone) {

		result["mobile_phone"] = obj.MobilePhone

	}

	if obj.DisplayName != nil && isNotEmpty(obj.DisplayName) {

		result["display_name"] = obj.DisplayName

	}

	if obj.Initials != nil && isNotEmpty(obj.Initials) {

		result["initials"] = obj.Initials

	}

	if obj.PostalCode != nil && isNotEmpty(obj.PostalCode) {

		result["postal_code"] = obj.PostalCode

	}

	if obj.OfficePhone != nil && isNotEmpty(obj.OfficePhone) {

		result["office_phone"] = obj.OfficePhone

	}

	if obj.HomeFolderDirectory != nil && isNotEmpty(obj.HomeFolderDirectory) {

		result["home_folder_directory"] = obj.HomeFolderDirectory

	}

	if obj.RdsConnectDrive != nil && isNotEmpty(obj.RdsConnectDrive) {

		result["rds_connect_drive"] = obj.RdsConnectDrive

	}

	if obj.Description != nil && isNotEmpty(obj.Description) {

		result["description"] = obj.Description

	}

	if obj.SamUsername != nil && isNotEmpty(obj.SamUsername) {

		result["sam_username"] = obj.SamUsername

	}

	if obj.HomePhone != nil && isNotEmpty(obj.HomePhone) {

		result["home_phone"] = obj.HomePhone

	}

	if obj.StreetPoBox != nil && isNotEmpty(obj.StreetPoBox) {

		result["street_po_box"] = obj.StreetPoBox

	}

	if obj.Country != nil && isNotEmpty(obj.Country) {

		result["country"] = obj.Country

	}

	if obj.ChangePasswordNextLogon != nil && isNotEmpty(obj.ChangePasswordNextLogon) {

		result["change_password_next_logon"] = obj.ChangePasswordNextLogon

	}

	if obj.EmailAddress != nil && isNotEmpty(obj.EmailAddress) {

		result["email_address"] = obj.EmailAddress

	}

	if obj.OuPath != nil && isNotEmpty(obj.OuPath) {

		result["ou_path"] = obj.OuPath

	}

	if obj.Department != nil && isNotEmpty(obj.Department) {

		result["department"] = obj.Department

	}

	if obj.PasswordNeverExpired != nil && isNotEmpty(obj.PasswordNeverExpired) {

		result["password_never_expired"] = obj.PasswordNeverExpired

	}

	if obj.Manager != nil && isNotEmpty(obj.Manager) {

		result["manager"] = obj.Manager

	}

	if obj.ProfilePath != nil && isNotEmpty(obj.ProfilePath) {

		result["profile_path"] = obj.ProfilePath

	}

	if obj.Username != nil && isNotEmpty(obj.Username) {

		result["username"] = obj.Username

	}

	if obj.FirstName != nil && isNotEmpty(obj.FirstName) {

		result["first_name"] = obj.FirstName

	}

	return result
}

func UserListToTerraform(list *[]api.User) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, User := range *list {
			results[i] = UserToTerraform(&User)
		}
		return results
	}
	return make([]interface{}, 0)
}

func GroupToTerraform(obj *api.Group) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.Description != nil && isNotEmpty(obj.Description) {

		result["description"] = obj.Description

	}

	if obj.Email != nil && isNotEmpty(obj.Email) {

		result["email"] = obj.Email

	}

	if obj.OuPath != nil && isNotEmpty(obj.OuPath) {

		result["ou_path"] = obj.OuPath

	}

	if obj.Type != nil && isNotEmpty(obj.Type) {

		result["type"] = obj.Type

	}

	if obj.Guid != nil && isNotEmpty(obj.Guid) {

		result["guid"] = UuidToTerraform(obj.Guid)

	}

	if obj.Name != nil && isNotEmpty(obj.Name) {

		result["name"] = obj.Name

	}

	if obj.Scope != nil && isNotEmpty(obj.Scope) {

		result["scope"] = obj.Scope

	}

	return result
}

func GroupListToTerraform(list *[]api.Group) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, Group := range *list {
			results[i] = GroupToTerraform(&Group)
		}
		return results
	}
	return make([]interface{}, 0)
}

func OUToTerraform(obj *api.OU) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.Description != nil && isNotEmpty(obj.Description) {

		result["description"] = obj.Description

	}

	if obj.Path != nil && isNotEmpty(obj.Path) {

		result["path"] = obj.Path

	}

	if obj.Protected != nil && isNotEmpty(obj.Protected) {

		result["protected"] = obj.Protected

	}

	if obj.Guid != nil && isNotEmpty(obj.Guid) {

		result["guid"] = UuidToTerraform(obj.Guid)

	}

	if obj.Name != nil && isNotEmpty(obj.Name) {

		result["name"] = obj.Name

	}

	return result
}

func OUListToTerraform(list *[]api.OU) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, OU := range *list {
			results[i] = OUToTerraform(&OU)
		}
		return results
	}
	return make([]interface{}, 0)
}

func ComputerToTerraform(obj *api.Computer) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.OperatingSystem != nil && isNotEmpty(obj.OperatingSystem) {

		result["operating_system"] = obj.OperatingSystem

	}

	if obj.Description != nil && isNotEmpty(obj.Description) {

		result["description"] = obj.Description

	}

	if obj.DnsName != nil && isNotEmpty(obj.DnsName) {

		result["dns_name"] = obj.DnsName

	}

	if obj.Path != nil && isNotEmpty(obj.Path) {

		result["path"] = obj.Path

	}

	if obj.Created != nil && isNotEmpty(obj.Created) {

		result["created"] = obj.Created

	}

	if obj.Sid != nil && isNotEmpty(obj.Sid) {

		result["sid"] = obj.Sid

	}

	if obj.ComputerName != nil && isNotEmpty(obj.ComputerName) {

		result["computer_name"] = obj.ComputerName

	}

	return result
}

func ComputerListToTerraform(list *[]api.Computer) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, Computer := range *list {
			results[i] = ComputerToTerraform(&Computer)
		}
		return results
	}
	return make([]interface{}, 0)
}

func DnsRecordToTerraform(obj *api.DnsRecord) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.Name != nil && isNotEmpty(obj.Name) {

		result["name"] = obj.Name

	}

	if obj.Type != nil && isNotEmpty(obj.Type) {

		result["type"] = obj.Type

	}

	if obj.Data != nil && isNotEmpty(obj.Data) {

		result["data"] = obj.Data

	}

	if obj.Zone != nil && isNotEmpty(obj.Zone) {

		result["zone"] = obj.Zone

	}

	if obj.Timestamp != nil && isNotEmpty(obj.Timestamp) {

		result["timestamp"] = obj.Timestamp

	}

	return result
}

func DnsRecordListToTerraform(list *[]api.DnsRecord) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, DnsRecord := range *list {
			results[i] = DnsRecordToTerraform(&DnsRecord)
		}
		return results
	}
	return make([]interface{}, 0)
}

func DnsZoneToTerraform(obj *api.DnsZone) map[string]interface{} {
	result := make(map[string]interface{})

	if obj.DistinguishedName != nil && isNotEmpty(obj.DistinguishedName) {

		result["distinguished_name"] = obj.DistinguishedName

	}

	if obj.ZoneName != nil && isNotEmpty(obj.ZoneName) {

		result["zone_name"] = obj.ZoneName

	}

	if obj.ZoneType != nil && isNotEmpty(obj.ZoneType) {

		result["zone_type"] = obj.ZoneType

	}

	if obj.Type != nil && isNotEmpty(obj.Type) {

		result["type"] = obj.Type

	}

	return result
}

func DnsZoneListToTerraform(list *[]api.DnsZone) []interface{} {
	if list != nil {
		results := make([]interface{}, len(*list))
		for i, DnsZone := range *list {
			results[i] = DnsZoneToTerraform(&DnsZone)
		}
		return results
	}
	return make([]interface{}, 0)
}
