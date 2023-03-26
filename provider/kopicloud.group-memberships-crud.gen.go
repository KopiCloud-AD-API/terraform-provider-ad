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

func resourceGroupMembership() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["user_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: false,
		Required: true,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["group_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["result"] = schemaOfGroup(``)

	return &schema.Resource{

		CreateContext: resourceGroupMembershipCreate,

		ReadContext: resourceGroupMembershipRead,

		DeleteContext: resourceGroupMembershipDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceGroupMembershipCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupMembershipCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiADUserUsernameGroupGroupNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument user_name"))
	user_name := d.Get("user_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument group_name"))
	group_name := d.Get("group_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiADUserUsernameGroupGroupNameWithResponse"))
	res, err := c.client.PostApiADUserUsernameGroupGroupNameWithResponse(
		ctx,

		user_name,

		group_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiADUserUsernameGroupGroupNameWithResponse returned successfully"))

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

			d.SetId(getId_for_GroupMembership(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Group")
	}

	tflog.Debug(ctx, "resourceGroupMembershipcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceGroupMembershipRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupMembershipRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiADUserUsernameGroupGroupNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument user_name"))
	user_name := d.Get("user_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument group_name"))
	group_name := d.Get("group_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiADUserUsernameGroupGroupNameWithResponse"))
	res, err := c.client.GetApiADUserUsernameGroupGroupNameWithResponse(
		ctx,

		user_name,

		group_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiADUserUsernameGroupGroupNameWithResponse returned successfully"))

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

			d.SetId(getId_for_GroupMembership(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "Group")
	}

	tflog.Debug(ctx, "resourceGroupMembershipread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceGroupMembershipDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceGroupMembershipDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiADUserUsernameGroupGroupNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument user_name"))
	user_name := d.Get("user_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument group_name"))
	group_name := d.Get("group_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiADUserUsernameGroupGroupNameWithResponse"))
	res, err := c.client.DeleteApiADUserUsernameGroupGroupNameWithResponse(
		ctx,

		user_name,

		group_name,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiADUserUsernameGroupGroupNameWithResponse returned successfully"))

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

	tflog.Debug(ctx, "resourceGroupMembershipdelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
