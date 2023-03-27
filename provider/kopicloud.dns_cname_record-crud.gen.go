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

func resourceDnsCNameRecord() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["hostname"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["hostname_alias"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["zone_name"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		ForceNew: true,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceDnsCNameRecordCreate,

		ReadContext: resourceDnsCNameRecordRead,

		DeleteContext: resourceDnsCNameRecordDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceDnsCNameRecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsCNameRecordCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname"))

	hostname := d.Get("hostname").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname_alias"))

	hostname_alias := d.Get("hostname_alias").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute zone_name"))

	zone_name := d.Get("zone_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiDnsCNameRecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		DNSHostNameAlias: &hostname_alias,

		ZoneName: &zone_name,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiDnsCNameRecordWithResponse"))
	res, err := c.client.PostApiDnsCNameRecordWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiDnsCNameRecordWithResponse returned successfully"))

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

			resItems := DnsRecordToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsCNameRecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsCNameRecord",
					map[string]interface{}{
						"DnsCNameRecord": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_DnsCNameRecord(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "resourceDnsCNameRecordcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsCNameRecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsCNameRecordRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname"))

	hostname := d.Get("hostname").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname_alias"))

	hostname_alias := d.Get("hostname_alias").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute zone_name"))

	zone_name := d.Get("zone_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiDnsCNameRecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		DNSHostNameAlias: &hostname_alias,

		ZoneName: &zone_name,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiDnsCNameRecordWithResponse"))
	res, err := c.client.GetApiDnsCNameRecordWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiDnsCNameRecordWithResponse returned successfully"))

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

			resItems := DnsRecordToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsCNameRecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsCNameRecord",
					map[string]interface{}{
						"DnsCNameRecord": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_DnsCNameRecord(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "resourceDnsCNameRecordread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsCNameRecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsCNameRecordDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname"))

	hostname := d.Get("hostname").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute hostname_alias"))

	hostname_alias := d.Get("hostname_alias").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for attribute zone_name"))

	zone_name := d.Get("zone_name").(string)

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiDnsCNameRecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		DNSHostNameAlias: &hostname_alias,

		ZoneName: &zone_name,
	}

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiDnsCNameRecordWithResponse"))
	res, err := c.client.DeleteApiDnsCNameRecordWithResponse(
		ctx,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiDnsCNameRecordWithResponse returned successfully"))

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

			resItems := DnsRecordToTerraform(api_result)
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsCNameRecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsCNameRecord",
					map[string]interface{}{
						"DnsCNameRecord": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "resourceDnsCNameRecorddelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
