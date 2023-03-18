package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	kcapi "github.com/KopiCloud-AD-API/terraform-provider-ad/api"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceOUUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var result diag.Diagnostics
	var required_inputs []string

	required_inputs = make([]string, 1)

	required_inputs[0] = "ou_name"

	if checkFieldsModified(d, required_inputs) {
		result = resourceOU_Update_0(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "ou_name"

	if checkFieldsModified(d, required_inputs) {
		result = resourceOU_Update_1(ctx, d, m)
	}

	return result
}

func resourceOU_Update_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceOU_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	ou_path := getFieldValue("result.0.path", false, d).(string)

	ou_name := getFieldValue("ou_name", false, d).(string)

	params := kcapi.PutApiOURenameParams{
		AuthToken: c.data.Get("token").(string),

		OUPath: ou_path,

		OUNewName: ou_name,
	}

	res, err := c.client.PutApiOURenameWithResponse(
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

			resItems := OUToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted OU }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted OU",
					map[string]interface{}{
						"OU": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_OU(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "OU")
	}

	tflog.Debug(ctx, "dataSourceOURead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceOU_Update_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceOU_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	ou_path := getFieldValue("ou_path", false, d).(string)

	description := getFieldValue("description", false, d).(string)

	protected := getFieldValue("protected", false, d).(bool)

	params := kcapi.PutApiOUParams{
		AuthToken: c.data.Get("token").(string),

		OUPath: ou_path,

		OUDescription: &description,

		IsProtected: &protected,
	}

	res, err := c.client.PutApiOUWithResponse(
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

			resItems := OUToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted OU }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted OU",
					map[string]interface{}{
						"OU": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_OU(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "OU")
	}

	tflog.Debug(ctx, "dataSourceOURead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
