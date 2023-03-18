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

func dataSourceGroupMembership() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfGroupList(``)

	terraformSchema["user_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    false,
		Required:    true,
		Description: "",
	}

	terraformSchema["group_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	return &schema.Resource{
		ReadContext: dataSourceGroupMembershipRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceGroupMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 1)

	required_inputs[0] = "user_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceGroupMembership_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 2)

	required_inputs[0] = "user_name"

	required_inputs[1] = "group_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceGroupMembership_Read_1(ctx, d, m)
	}

	return nil
}

func dataSourceGroupMembership_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceGroupMembership_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiADUserUsernameGroupAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	user_name := d.Get("user_name").(string)

	res, err := c.client.GetApiADUserUsernameGroupAllWithResponse(
		ctx,

		user_name,

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

			resItems := GroupListToTerraform(api_result)

			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted GroupMembership }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted GroupMembership",
					map[string]interface{}{
						"GroupMembership": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "GroupMembership")
	}

	tflog.Debug(ctx, "dataSourceGroupMembershipRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceGroupMembership_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceGroupMembership_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiADUserUsernameGroupGroupNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	user_name := d.Get("user_name").(string)

	group_name := d.Get("group_name").(string)

	res, err := c.client.GetApiADUserUsernameGroupGroupNameWithResponse(
		ctx,

		user_name,

		group_name,

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

			resItems := GroupToTerraform(api_result)

			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted GroupMembership }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted GroupMembership",
					map[string]interface{}{
						"GroupMembership": resItems,
					})
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_GroupMembership(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "GroupMembership")
	}

	tflog.Debug(ctx, "dataSourceGroupMembershipRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
