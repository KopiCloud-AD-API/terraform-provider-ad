package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	kcapi "github.com/KopiCloud-AD-API/terraform-provider-ad/api"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfUserList(``)

	terraformSchema["show_fields"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	terraformSchema["ou_path"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	terraformSchema["recursive"] = &schema.Schema{
		Type:        schema.TypeBool,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema:      terraformSchema,
		Description: "Read Users from Active Directory",
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 1)

	required_inputs[0] = "ou_path"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceUser_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceUser_Read_1(ctx, d, m)
	}

	return nil
}

func dataSourceUser_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceUser_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	show_fields := d.Get("show_fields").(string)

	ou_path := d.Get("ou_path").(string)

	params := kcapi.GetApiADUserListUsersParams{
		AuthToken: c.data.Get("token").(string),

		ShowFields: &show_fields,

		OUPath: &ou_path,
	}

	res, err := c.client.GetApiADUserListUsersWithResponse(
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
					if output, ok := result["output"]; !ok {
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

			resItems := UserListToTerraform(api_result)

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

			result := resItems

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			if len(d.Id()) == 0 {
				if uuid, err := uuid.NewV4(); err == nil {
					d.SetId(uuid.String())
				}
			} else {
				panic(err)
			}

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

func dataSourceUser_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceUser_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	show_fields := d.Get("show_fields").(string)

	params := kcapi.GetApiADUserListUsersAllParams{
		AuthToken: c.data.Get("token").(string),

		ShowFields: &show_fields,
	}

	res, err := c.client.GetApiADUserListUsersAllWithResponse(
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
					if output, ok := result["output"]; !ok {
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

			resItems := UserListToTerraform(api_result)

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

			result := resItems

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			if len(d.Id()) == 0 {
				if uuid, err := uuid.NewV4(); err == nil {
					d.SetId(uuid.String())
				}
			} else {
				panic(err)
			}

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
