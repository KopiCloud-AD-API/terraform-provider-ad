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

func resourceComputerCleanUp() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfComputerList(``)

	terraformSchema["days"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: false,
		Optional: false,
		Required: true,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["ou_path"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: false,
		Required: true,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["recursive"] = &schema.Schema{
		Type:     schema.TypeBool,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceComputerCleanUpCreate,

		ReadContext: resourceComputerCleanUpRead,

		DeleteContext: resourceComputerCleanUpDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceComputerCleanUpCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerCleanUpCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute days"))

	days := d.Get("days").(int)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute recursive"))

	recursive := d.Get("recursive").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiComputersCleanUpParams{
		AuthToken: c.data.Get("token").(string),

		Days: intToPInt32(days),

		OUPath: &ou_path,

		Recursive: &recursive,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiComputersCleanUpWithResponse"))
	res, err := c.client.DeleteApiComputersCleanUpWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiComputersCleanUpWithResponse returned successfully"))

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

			resItems := ComputerListToTerraform(api_result)
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

			result := resItems

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			if len(d.Id()) == 0 {
				if uuid, err := uuid.NewRandom(); err == nil {
					d.SetId(uuid.String())
				} else {
					return diag.FromErr(err)
				}
			}

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputerCleanUpcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceComputerCleanUpRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerCleanUpRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute days"))

	days := d.Get("days").(int)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute recursive"))

	recursive := d.Get("recursive").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiComputersCleanUpParams{
		AuthToken: c.data.Get("token").(string),

		Days: intToPInt32(days),

		OUPath: &ou_path,

		Recursive: &recursive,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiComputersCleanUpWithResponse"))
	res, err := c.client.DeleteApiComputersCleanUpWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiComputersCleanUpWithResponse returned successfully"))

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

			resItems := ComputerListToTerraform(api_result)
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

			result := resItems

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			if len(d.Id()) == 0 {
				if uuid, err := uuid.NewRandom(); err == nil {
					d.SetId(uuid.String())
				} else {
					return diag.FromErr(err)
				}
			}

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputerCleanUpread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceComputerCleanUpDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceComputerCleanUpDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute days"))

	days := d.Get("days").(int)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute ou_path"))

	ou_path := d.Get("ou_path").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute recursive"))

	recursive := d.Get("recursive").(bool)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiComputersCleanUpParams{
		AuthToken: c.data.Get("token").(string),

		Days: intToPInt32(days),

		OUPath: &ou_path,

		Recursive: &recursive,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiComputersCleanUpWithResponse"))
	res, err := c.client.DeleteApiComputersCleanUpWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiComputersCleanUpWithResponse returned successfully"))

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

			resItems := ComputerListToTerraform(api_result)
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

			result := resItems

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "resourceComputerCleanUpdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
