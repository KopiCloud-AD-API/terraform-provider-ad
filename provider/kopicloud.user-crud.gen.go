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

func resourceUser() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["username"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["password"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
		Sensitive: true,

		Description: "",
	}

	terraformSchema["first_name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["initials"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["last_name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["display_name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["description"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["email_address"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["department"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["office"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["company"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["change_password_next_logon"] = &schema.Schema{
		Type:      schema.TypeBool,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["password_never_expires"] = &schema.Schema{
		Type:      schema.TypeBool,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["ou_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["job_title"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["manager"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["street"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["po_box"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["city"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["state"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["zip_code"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["country"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["office_phone"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["home_phone"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["mobile_phone"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["profile_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["profile_logon_script"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["home_folder_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["home_folder_drive"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["home_folder_directory"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["rds_profile_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["rds_home_folder_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["rds_home_folder_drive"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["rds_connect_drive"] = &schema.Schema{
		Type:      schema.TypeBool,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["rds_allow_logon"] = &schema.Schema{
		Type:      schema.TypeBool,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["show_fields"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceUserCreate,

		ReadContext: resourceUserRead,

		UpdateContext: resourceUserUpdate,

		DeleteContext: resourceUserDelete,

		Schema:      terraformSchema,
		Description: "Element to create, update, delete Users from Active Directory",
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute password"))

	password := d.Get("password").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute first_name"))

	first_name := d.Get("first_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute initials"))

	initials := d.Get("initials").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute last_name"))

	last_name := d.Get("last_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute display_name"))

	display_name := d.Get("display_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute description"))

	description := d.Get("description").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute email_address"))

	email_address := d.Get("email_address").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute department"))

	department := d.Get("department").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute office"))

	office := d.Get("office").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute company"))

	company := d.Get("company").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute job_title"))

	job_title := d.Get("job_title").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute manager"))

	manager := d.Get("manager").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute street"))

	street := d.Get("street").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute po_box"))

	po_box := d.Get("po_box").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute city"))

	city := d.Get("city").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute state"))

	state := d.Get("state").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute zip_code"))

	zip_code := d.Get("zip_code").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute country"))

	country := d.Get("country").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute office_phone"))

	office_phone := d.Get("office_phone").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute home_phone"))

	home_phone := d.Get("home_phone").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute mobile_phone"))

	mobile_phone := d.Get("mobile_phone").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute profile_path"))

	profile_path := d.Get("profile_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute profile_logon_script"))

	profile_logon_script := d.Get("profile_logon_script").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute home_folder_path"))

	home_folder_path := d.Get("home_folder_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute home_folder_drive"))

	home_folder_drive := d.Get("home_folder_drive").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute home_folder_directory"))

	home_folder_directory := d.Get("home_folder_directory").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute rds_profile_path"))

	rds_profile_path := d.Get("rds_profile_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute rds_home_folder_path"))

	rds_home_folder_path := d.Get("rds_home_folder_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute rds_home_folder_drive"))

	rds_home_folder_drive := d.Get("rds_home_folder_drive").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute rds_connect_drive"))

	rds_connect_drive := d.Get("rds_connect_drive").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute rds_allow_logon"))

	rds_allow_logon := d.Get("rds_allow_logon").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute change_password_next_logon"))

	change_password_next_logon := d.Get("change_password_next_logon").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute password_never_expires"))

	password_never_expires := d.Get("password_never_expires").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute show_fields"))

	show_fields := d.Get("show_fields").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.APIUserParams{
		AuthToken: c.data.Get("token").(string),

		Password: &password,

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

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument username"))
	username := d.Get("username").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API APIUserWithResponse"))
	res, err := c.client.APIUserWithResponse(
		ctx,

		username,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API APIUserWithResponse returned successfully"))

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

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_User(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "User")
	}

	tflog.Debug(ctx, "resourceUsercreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for GUID"))
	guid, err := uuid.Parse(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute show_fields"))

	show_fields := d.Get("show_fields").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiADUserGuidParams{
		AuthToken: c.data.Get("token").(string),

		UserGuid: guid,

		ShowFields: &show_fields,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiADUserGuidWithResponse"))
	res, err := c.client.GetApiADUserGuidWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiADUserGuidWithResponse returned successfully"))

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

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_User(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "User")
	}

	tflog.Debug(ctx, "resourceUserread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute show_fields"))

	show_fields := d.Get("show_fields").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiADUserUsernameParams{
		AuthToken: c.data.Get("token").(string),

		ShowFields: &show_fields,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument username"))
	username := d.Get("username").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiADUserUsernameWithResponse"))
	res, err := c.client.DeleteApiADUserUsernameWithResponse(
		ctx,

		username,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiADUserUsernameWithResponse returned successfully"))

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

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "User")
	}

	tflog.Debug(ctx, "resourceUserdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
