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

func resourceOU() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["ou_name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["ou_destination_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["ou_path"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  true,
		Optional:  false,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["description"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["protected"] = &schema.Schema{
		Type:      schema.TypeBool,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceOUCreate,

		ReadContext: resourceOURead,

		UpdateContext: resourceOUUpdate,

		DeleteContext: resourceOUDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceOUCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceOUCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_name"))

	ou_name := d.Get("ou_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute description"))

	description := d.Get("description").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_destination_path"))

	ou_destination_path := d.Get("ou_destination_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute protected"))

	protected := d.Get("protected").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiOUParams{
		AuthToken: c.data.Get("token").(string),

		OUName: ou_name,

		OUDescription: &description,

		OUDestinationPath: ou_destination_path,

		IsProtected: &protected,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiOUWithResponse"))
	res, err := c.client.PostApiOUWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiOUWithResponse returned successfully"))

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

			if err := d.Set("ou_path", res.JSON200.Result.Path); err != nil {
				return diag.FromErr(err)
			}

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_OU(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "OU")
	}

	tflog.Debug(ctx, "resourceOUcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceOURead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceOURead"))
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

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetOUByGuidParams{
		AuthToken: c.data.Get("token").(string),

		OUGuid: guid,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetOUByGuidWithResponse"))
	res, err := c.client.GetOUByGuidWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetOUByGuidWithResponse returned successfully"))

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

			if err := d.Set("ou_path", res.JSON200.Result.Path); err != nil {
				return diag.FromErr(err)
			}

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_OU(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "OU")
	}

	tflog.Debug(ctx, "resourceOUread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceOUDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceOUDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute protected"))

	protected := d.Get("protected").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiOUPathRemoveParams{
		AuthToken: c.data.Get("token").(string),

		OUPath: ou_path,

		Force: &protected,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiOUPathRemoveWithResponse"))
	res, err := c.client.DeleteApiOUPathRemoveWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiOUPathRemoveWithResponse returned successfully"))

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

			if err := d.Set("ou_path", res.JSON200.Result.Path); err != nil {
				return diag.FromErr(err)
			}

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "OU")
	}

	tflog.Debug(ctx, "resourceOUdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
