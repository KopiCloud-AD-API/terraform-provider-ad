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

func resourceDnsLookupZone() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsZone(``)

	terraformSchema["zone_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceDnsLookupZoneCreate,

		ReadContext: resourceDnsLookupZoneRead,

		DeleteContext: resourceDnsLookupZoneDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceDnsLookupZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsLookupZoneRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	params := kcapi.PostApiDnsLookupZoneZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.PostApiDnsLookupZoneZoneNameWithResponse(
		ctx,

		zone_name,

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

			resItems := DnsZoneToTerraform(api_result)

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsLookupZone: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_DnsLookupZone(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsLookupZonecreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsLookupZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsLookupZoneRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	params := kcapi.GetApiDnsZonesZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.GetApiDnsZonesZoneNameWithResponse(
		ctx,

		zone_name,

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

			resItems := DnsZoneToTerraform(api_result)

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsLookupZone: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_DnsLookupZone(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsLookupZoneread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsLookupZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsLookupZoneRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	params := kcapi.DeleteApiDnsLookupZoneZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.DeleteApiDnsLookupZoneZoneNameWithResponse(
		ctx,

		zone_name,

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

			resItems := DnsZoneToTerraform(api_result)

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsLookupZone: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsLookupZonedelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
