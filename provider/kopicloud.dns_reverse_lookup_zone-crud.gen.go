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

func resourceDnsReverseLookupZone() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["network_id"] = &schema.Schema{
		Type:      schema.TypeString,
		Computed:  false,
		Optional:  true,
		Required:  false,
		Sensitive: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["result"] = schemaOfDnsZone(``)

	return &schema.Resource{

		CreateContext: resourceDnsReverseLookupZoneCreate,

		ReadContext: resourceDnsReverseLookupZoneRead,

		DeleteContext: resourceDnsReverseLookupZoneDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceDnsReverseLookupZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsReverseLookupZoneCreate"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.PostApiDnsReverseLookupZoneNetworkIDParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument network_id"))
	network_id := d.Get("network_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API PostApiDnsReverseLookupZoneNetworkIDWithResponse"))
	res, err := c.client.PostApiDnsReverseLookupZoneNetworkIDWithResponse(
		ctx,

		network_id,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API PostApiDnsReverseLookupZoneNetworkIDWithResponse returned successfully"))

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
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsReverseLookupZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsReverseLookupZone",
					map[string]interface{}{
						"DnsReverseLookupZone": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_DnsLookupZone(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsReverseLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsReverseLookupZonecreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsReverseLookupZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsReverseLookupZoneRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.GetApiDnsReverseLookupZoneNetworkIDParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument network_id"))
	network_id := d.Get("network_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API GetApiDnsReverseLookupZoneNetworkIDWithResponse"))
	res, err := c.client.GetApiDnsReverseLookupZoneNetworkIDWithResponse(
		ctx,

		network_id,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API GetApiDnsReverseLookupZoneNetworkIDWithResponse returned successfully"))

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
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsReverseLookupZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsReverseLookupZone",
					map[string]interface{}{
						"DnsReverseLookupZone": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId(getId_for_DnsLookupZone(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsReverseLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsReverseLookupZoneread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsReverseLookupZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsReverseLookupZoneDelete"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	tflog.Debug(ctx, fmt.Sprintf("Creating params structure"))
	params := kcapi.DeleteApiDnsReverseLookupZoneNetworkIDParams{
		AuthToken: c.data.Get("token").(string),
	}

	tflog.Debug(ctx, fmt.Sprintf("Creating variable for function argument network_id"))
	network_id := d.Get("network_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("Calling API DeleteApiDnsReverseLookupZoneNetworkIDWithResponse"))
	res, err := c.client.DeleteApiDnsReverseLookupZoneNetworkIDWithResponse(
		ctx,

		network_id,

		&params)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Debug(ctx, fmt.Sprintf("API DeleteApiDnsReverseLookupZoneNetworkIDWithResponse returned successfully"))

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
			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsReverseLookupZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsReverseLookupZone",
					map[string]interface{}{
						"DnsReverseLookupZone": resItems,
					})
			}

			result := wrapInArray(resItems)

			tflog.Debug(ctx, fmt.Sprintf("Ignoring result: %#v", result))

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsReverseLookupZone")
	}

	tflog.Debug(ctx, "resourceDnsReverseLookupZonedelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
