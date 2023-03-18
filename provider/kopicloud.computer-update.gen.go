package provider

import (
	"context"
	"encoding/json"
	"fmt"

	kcapi "github.com/KopiCloud-AD-API/terraform-provider-ad/api"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var result diag.Diagnostics
	var required_inputs []string

	required_inputs = make([]string, 1)

	required_inputs[0] = "ad_computer_name"

	if checkFieldsModified(d, required_inputs) {
		result = resourceComputer_Update_0(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "description"

	if checkFieldsModified(d, required_inputs) {
		result = resourceComputer_Update_1(ctx, d, m)
	}

	return result
}

func resourceComputer_Update_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceComputer_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.PutApiComputersADComputerNameRenameNewADComputerNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	ad_computer_name_old := getFieldValue("ad_computer_name", true, d).(string)

	ad_computer_name := getFieldValue("ad_computer_name", false, d).(string)

	res, err := c.client.PutApiComputersADComputerNameRenameNewADComputerNameWithResponse(
		ctx,

		ad_computer_name_old,

		ad_computer_name,

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

			resItems := ComputerToTerraform(api_result)

			for _, element := range resItems {
				tflog.Debug(ctx, "converted Computer",
					map[string]interface{}{
						"Computer": element,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_Computer(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "dataSourceComputerRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceComputer_Update_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceComputer_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	description := getFieldValue("description", false, d).(string)

	params := kcapi.PutApiComputersADComputerNameUpdateParams{
		AuthToken: c.data.Get("token").(string),

		ComputerDescription: &description,
	}

	ad_computer_name := getFieldValue("ad_computer_name", false, d).(string)

	res, err := c.client.PutApiComputersADComputerNameUpdateWithResponse(
		ctx,

		ad_computer_name,

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

			resItems := ComputerToTerraform(api_result)

			for _, element := range resItems {
				tflog.Debug(ctx, "converted Computer",
					map[string]interface{}{
						"Computer": element,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_Computer(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Computer")
	}

	tflog.Debug(ctx, "dataSourceComputerRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
