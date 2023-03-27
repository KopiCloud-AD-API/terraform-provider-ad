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

func dataSourceDnsZoneReader() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsZoneList(``)

	return &schema.Resource{
		ReadContext: dataSourceDnsZoneReaderRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceDnsZoneReaderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsZoneReader_Read_0(ctx, d, m)
	}

	return nil
}

func dataSourceDnsZoneReader_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsZoneReader_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsZonesAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	res, err := c.client.GetApiDnsZonesAllWithResponse(
		ctx,

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

			resItems := DnsZoneListToTerraform(api_result)

			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsZone",
					map[string]interface{}{
						"DnsZone": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsZone")
	}

	tflog.Debug(ctx, "dataSourceDnsZoneReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsLookupZoneReader() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsZoneList(``)

	terraformSchema["zone_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Sensitive:   false,
		Description: "",
	}

	return &schema.Resource{
		ReadContext: dataSourceDnsLookupZoneReaderRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceDnsLookupZoneReaderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 1)

	required_inputs[0] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsLookupZoneReader_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsLookupZoneReader_Read_1(ctx, d, m)
	}

	return nil
}

func dataSourceDnsLookupZoneReader_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsLookupZoneReader_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
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

			resItems := DnsZoneToTerraform(api_result)

			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsLookupZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsLookupZone",
					map[string]interface{}{
						"DnsLookupZone": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsLookupZone")
	}

	tflog.Debug(ctx, "dataSourceDnsLookupZoneReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsLookupZoneReader_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsLookupZoneReader_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsLookupZoneAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	res, err := c.client.GetApiDnsLookupZoneAllWithResponse(
		ctx,

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

			resItems := DnsZoneListToTerraform(api_result)

			rt := reflect.TypeOf(resItems)
			switch rt.Kind() {
			case reflect.Slice | reflect.Array:
				for _, element := range resItems {
					tflog.Debug(ctx, "converted DnsLookupZone }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsLookupZone",
					map[string]interface{}{
						"DnsLookupZone": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsLookupZone")
	}

	tflog.Debug(ctx, "dataSourceDnsLookupZoneReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
