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

func resourceComputer() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["ad_computer_name"] = &schema.Schema{
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

	terraformSchema["description"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceComputerCreate,

		ReadContext: resourceComputerRead,

		UpdateContext: resourceComputerUpdate,

		DeleteContext: resourceComputerDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceComputerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute description"))

	description := d.Get("description").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiComputersADComputerNameRegisterParams{
		AuthToken: c.data.Get("token").(string),

		OUPath: &ou_path,

		ADComputerDescription: &description,
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument ad_computer_name"))
	ad_computer_name := d.Get("ad_computer_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiComputersADComputerNameRegisterWithResponse"))
	res, err := c.client.PostApiComputersADComputerNameRegisterWithResponse(
		ctx,

		ad_computer_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiComputersADComputerNameRegisterWithResponse returned successfully"))

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

			resItems := ComputerToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Computer }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Computer",
					map[string]interface{}{
						"Computer": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_Computer(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputercreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceComputerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiComputersADComputerNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument ad_computer_name"))
	ad_computer_name := d.Get("ad_computer_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiComputersADComputerNameWithResponse"))
	res, err := c.client.GetApiComputersADComputerNameWithResponse(
		ctx,

		ad_computer_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiComputersADComputerNameWithResponse returned successfully"))

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

			resItems := ComputerToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Computer }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Computer",
					map[string]interface{}{
						"Computer": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_Computer(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputerread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceComputerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiComputersADComputerNameRemoveParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument ad_computer_name"))
	ad_computer_name := d.Get("ad_computer_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiComputersADComputerNameRemoveWithResponse"))
	res, err := c.client.DeleteApiComputersADComputerNameRemoveWithResponse(
		ctx,

		ad_computer_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiComputersADComputerNameRemoveWithResponse returned successfully"))

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

			resItems := ComputerToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted Computer }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted Computer",
					map[string]interface{}{
						"Computer": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputerdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
