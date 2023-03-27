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

func resourceGroup() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["scope"] = &schema.Schema{
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
		Optional:  false,
		Required:  true,
		Sensitive: false,

		Description: "",
	}

	terraformSchema["email"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  false,
		Required:  true,
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

	terraformSchema["result"] = schemaOfGroup(``)

	return &schema.Resource{

		CreateContext: resourceGroupCreate,

		ReadContext: resourceGroupRead,

		UpdateContext: resourceGroupUpdate,

		DeleteContext: resourceGroupDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute scope"))

	scope := d.Get("scope").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute description"))

	description := d.Get("description").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute email"))

	email := d.Get("email").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiADGroupsGroupNameSecurityParams{
		AuthToken: c.data.Get("token").(string),

		GroupScope: stringToGroupScope(scope),

		GroupDescription: &description,

		GroupEmail: &email,

		OUPath: &ou_path,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument name"))
	name := d.Get("name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiADGroupsGroupNameSecurityWithResponse"))
	res, err := c.client.PostApiADGroupsGroupNameSecurityWithResponse(
		ctx,

		name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiADGroupsGroupNameSecurityWithResponse returned successfully"))

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

			resItems := GroupToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Group }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Group",
					map[string]interface{}{
						"Group": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_Group(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Group")
	}

	tflog.Debug(ctx, "resourceGroupcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiADGroupsGroupNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument name"))
	name := d.Get("name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiADGroupsGroupNameWithResponse"))
	res, err := c.client.GetApiADGroupsGroupNameWithResponse(
		ctx,

		name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiADGroupsGroupNameWithResponse returned successfully"))

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

			resItems := GroupToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Group }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Group",
					map[string]interface{}{
						"Group": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_Group(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Group")
	}

	tflog.Debug(ctx, "resourceGroupread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiADGroupsGroupNameRemoveParams{
		AuthToken: c.data.Get("token").(string),

		OUPath: &ou_path,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument name"))
	name := d.Get("name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiADGroupsGroupNameRemoveWithResponse"))
	res, err := c.client.DeleteApiADGroupsGroupNameRemoveWithResponse(
		ctx,

		name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiADGroupsGroupNameRemoveWithResponse returned successfully"))

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

			resItems := GroupToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Group }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Group",
					map[string]interface{}{
						"Group": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Group")
	}

	tflog.Debug(ctx, "resourceGroupdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
