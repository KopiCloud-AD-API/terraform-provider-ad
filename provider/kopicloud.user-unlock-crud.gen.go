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

func resourceUserUnlockAccount() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfUser(``)

	terraformSchema["username"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: false,
		Required: true,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["show_fields"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceUserUnlockAccountCreate,

		ReadContext: resourceUserUnlockAccountRead,

		DeleteContext: resourceUserUnlockAccountDelete,

		Schema:      terraformSchema,
		Description: "Element to create, update, delete Users from Active Directory",
	}
}

func resourceUserUnlockAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserUnlockAccountCreate"))
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
	params := kcapi.PutApiADUserUsernameUnlockParams{
		AuthToken: c.data.Get("token").(string),

		ShowFields: &show_fields,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument username"))
	username := d.Get("username").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API PutApiADUserUsernameUnlockWithResponse"))
	res, err := c.client.PutApiADUserUsernameUnlockWithResponse(
		ctx,

		username,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PutApiADUserUsernameUnlockWithResponse returned successfully"))

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
					tflog.Debug(ctx, "converted UserUnlockAccount }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted UserUnlockAccount",
					map[string]interface{}{
						"UserUnlockAccount": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_User(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "UserUnlockAccount")
	}

	tflog.Debug(ctx, "resourceUserUnlockAccountcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceUserUnlockAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserUnlockAccountRead"))
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
					tflog.Debug(ctx, "converted UserUnlockAccount }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted UserUnlockAccount",
					map[string]interface{}{
						"UserUnlockAccount": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_User(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "UserUnlockAccount")
	}

	tflog.Debug(ctx, "resourceUserUnlockAccountread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceUserUnlockAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceUserUnlockAccountDelete"))
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
					tflog.Debug(ctx, "converted UserUnlockAccount }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted UserUnlockAccount",
					map[string]interface{}{
						"UserUnlockAccount": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "UserUnlockAccount")
	}

	tflog.Debug(ctx, "resourceUserUnlockAccountdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
