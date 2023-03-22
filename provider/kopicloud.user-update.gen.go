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

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

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
