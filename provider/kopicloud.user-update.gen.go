package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	kcapi "github.com/KopiCloud-AD-API/terraform-provider-ad/api"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var result diag.Diagnostics
	var required_inputs []string

	required_inputs = make([]string, 0)

	if checkFieldsModified(d, required_inputs) {
		result = resourceUser_Update_0(ctx, d, m)
	}

	return result
}

func resourceUser_Update_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceUser_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	guid, err := uuid.Parse(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	first_name := getFieldValue("first_name", false, d).(string)

	initials := getFieldValue("initials", false, d).(string)

	last_name := getFieldValue("last_name", false, d).(string)

	display_name := getFieldValue("display_name", false, d).(string)

	description := getFieldValue("description", false, d).(string)

	email_address := getFieldValue("email_address", false, d).(string)

	department := getFieldValue("department", false, d).(string)

	office := getFieldValue("office", false, d).(string)

	company := getFieldValue("company", false, d).(string)

	ou_path := getFieldValue("ou_path", false, d).(string)

	job_title := getFieldValue("job_title", false, d).(string)

	manager := getFieldValue("manager", false, d).(string)

	street := getFieldValue("street", false, d).(string)

	po_box := getFieldValue("po_box", false, d).(string)

	city := getFieldValue("city", false, d).(string)

	state := getFieldValue("state", false, d).(string)

	zip_code := getFieldValue("zip_code", false, d).(string)

	country := getFieldValue("country", false, d).(string)

	office_phone := getFieldValue("office_phone", false, d).(string)

	home_phone := getFieldValue("home_phone", false, d).(string)

	mobile_phone := getFieldValue("mobile_phone", false, d).(string)

	profile_path := getFieldValue("profile_path", false, d).(string)

	profile_logon_script := getFieldValue("profile_logon_script", false, d).(string)

	home_folder_path := getFieldValue("home_folder_path", false, d).(string)

	home_folder_drive := getFieldValue("home_folder_drive", false, d).(string)

	home_folder_directory := getFieldValue("home_folder_directory", false, d).(string)

	rds_profile_path := getFieldValue("rds_profile_path", false, d).(string)

	rds_home_folder_path := getFieldValue("rds_home_folder_path", false, d).(string)

	rds_home_folder_drive := getFieldValue("rds_home_folder_drive", false, d).(string)

	rds_connect_drive := getFieldValue("rds_connect_drive", false, d).(bool)

	rds_allow_logon := getFieldValue("rds_allow_logon", false, d).(bool)

	change_password_next_logon := getFieldValue("change_password_next_logon", false, d).(bool)

	password_never_expires := getFieldValue("password_never_expires", false, d).(bool)

	show_fields := getFieldValue("show_fields", false, d).(string)

	params := kcapi.PutApiADUserGuidParams{
		AuthToken: c.data.Get("token").(string),

		UserGuid: guid,

		FirstName: &first_name,

		Initials: &initials,

		LastName: &last_name,

		DisplayName: &display_name,

		Description: &description,

		EmailAddress: &email_address,

		Department: &department,

		Office: &office,

		Company: &company,

		OUPath: &ou_path,

		JobTitle: &job_title,

		Manager: &manager,

		Street: &street,

		POBox: &po_box,

		City: &city,

		State: &state,

		ZipCode: &zip_code,

		Country: &country,

		OfficePhone: &office_phone,

		HomePhone: &home_phone,

		MobilePhone: &mobile_phone,

		ProfilePath: &profile_path,

		ProfileLogonScript: &profile_logon_script,

		HomeFolderPath: &home_folder_path,

		HomeFolderDrive: &home_folder_drive,

		HomeFolderDirectory: &home_folder_directory,

		RdsProfilePath: &rds_profile_path,

		RdsHomeFolderPath: &rds_home_folder_path,

		RdsHomeFolderDrive: &rds_home_folder_drive,

		RdsConnectDrive: &rds_connect_drive,

		RdsAllowLogon: &rds_allow_logon,

		ChangePasswordNextLogon: &change_password_next_logon,

		PasswordNeverExpired: &password_never_expires,

		ShowFields: &show_fields,
	}

	res, err := c.client.PutApiADUserGuidWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}

	if res != nil {
		if res.HTTPResponse.StatusCode != 200 {
			if res.JSON200 == nil {
				if json.Valid(res.Body) {
					var result map[string]interface{}
					json.Unmarshal(res.Body, &result)
					if output, ok := result["output"]; ok {
						tflog.Error(ctx, fmt.Sprintf("%s", output))
						return diag.Errorf(fmt.Sprintf("%s", output))
					}
				}
				tflog.Error(ctx, string(res.Body))
				return diag.Errorf(string(res.Body))
			} else {

				tflog.Error(ctx, *res.JSON200.Output)
				return diag.Errorf(string(*res.JSON200.Output))

			}
		} else {

			tflog.Info(ctx, *res.JSON200.Output)

			api_result := res.JSON200.Result

			resItems := UserToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted User }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted User",
					map[string]interface{}{
						"User": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("first_name", res.JSON200.Result.FirstName); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("initials", res.JSON200.Result.Initials); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("last_name", res.JSON200.Result.LastName); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("display_name", res.JSON200.Result.DisplayName); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("description", res.JSON200.Result.Description); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("email_address", res.JSON200.Result.EmailAddress); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("department", res.JSON200.Result.Department); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("office", res.JSON200.Result.Office); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("company", res.JSON200.Result.Company); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("change_password_next_logon", res.JSON200.Result.ChangePasswordNextLogon); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("password_never_expires", res.JSON200.Result.PasswordNeverExpired); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("ou_path", res.JSON200.Result.OuPath); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("job_title", res.JSON200.Result.JobTitle); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("manager", res.JSON200.Result.Manager); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("street", res.JSON200.Result.StreetAddress); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("po_box", res.JSON200.Result.StreetPoBox); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("city", res.JSON200.Result.City); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("state", res.JSON200.Result.State); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("zip_code", res.JSON200.Result.PostalCode); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("country", res.JSON200.Result.Country); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("office_phone", res.JSON200.Result.OfficePhone); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("home_phone", res.JSON200.Result.HomePhone); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("mobile_phone", res.JSON200.Result.MobilePhone); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("profile_path", res.JSON200.Result.ProfilePath); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("profile_logon_script", res.JSON200.Result.ProfileLogonScript); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("home_folder_path", res.JSON200.Result.HomeFolderPath); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("home_folder_drive", res.JSON200.Result.HomeFolderDrive); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("home_folder_directory", res.JSON200.Result.HomeFolderDirectory); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("rds_profile_path", res.JSON200.Result.RdsProfilePath); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("rds_home_folder_path", res.JSON200.Result.RdsHomeFolderPath); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("rds_home_folder_drive", res.JSON200.Result.RdsHomeFolderDrive); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("rds_connect_drive", res.JSON200.Result.RdsConnectDrive); err != nil {
				return diag.FromErr(err)
			}

			if err := d.Set("rds_allow_logon", res.JSON200.Result.RdsAllowLogon); err != nil {
				return diag.FromErr(err)
			}

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_User(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "User")
	}

	tflog.Debug(ctx, "dataSourceUserRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
