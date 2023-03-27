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

func dataSourceDnsCNameRecordReader() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsRecordList(``)

	terraformSchema["hostname"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Sensitive:   false,
		Description: "",
	}

	terraformSchema["hostname_alias"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Sensitive:   false,
		Description: "",
	}

	terraformSchema["zone_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Sensitive:   false,
		Description: "",
	}

	return &schema.Resource{
		ReadContext: dataSourceDnsCNameRecordReaderRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceDnsCNameRecordReaderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 3)

	required_inputs[0] = "hostname"

	required_inputs[1] = "hostname_alias"

	required_inputs[2] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsCNameRecordReader_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "hostname"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsCNameRecordReader_Read_1(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "hostname_alias"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsCNameRecordReader_Read_2(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsCNameRecordReader_Read_3(ctx, d, m)
	}

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsCNameRecordReader_Read_4(ctx, d, m)
	}

	return nil
}

func dataSourceDnsCNameRecordReader_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsCNameRecordReader_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	hostname := d.Get("hostname").(string)

	hostname_alias := d.Get("hostname_alias").(string)

	zone_name := d.Get("zone_name").(string)

	params := kcapi.GetApiDnsCNameRecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		DNSHostNameAlias: &hostname_alias,

		ZoneName: &zone_name,
	}

	res, err := c.client.GetApiDnsCNameRecordWithResponse(
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
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "dataSourceDnsCNameRecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsCNameRecordReader_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsCNameRecordReader_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsCNameRecordHostNameDNSHostNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	hostname := d.Get("hostname").(string)

	res, err := c.client.GetApiDnsCNameRecordHostNameDNSHostNameWithResponse(
		ctx,

		hostname,

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

			resItems := DnsRecordListToTerraform(api_result)

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
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "dataSourceDnsCNameRecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsCNameRecordReader_Read_2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsCNameRecordReader_Read_2"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsCNameRecordHostNameAliasDNSHostNameAliasParams{
		AuthToken: c.data.Get("token").(string),
	}

	hostname_alias := d.Get("hostname_alias").(string)

	res, err := c.client.GetApiDnsCNameRecordHostNameAliasDNSHostNameAliasWithResponse(
		ctx,

		hostname_alias,

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

			resItems := DnsRecordListToTerraform(api_result)

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
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "dataSourceDnsCNameRecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsCNameRecordReader_Read_3(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsCNameRecordReader_Read_3"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsCNameRecordZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.GetApiDnsCNameRecordZoneNameWithResponse(
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

			resItems := DnsRecordListToTerraform(api_result)

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
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "dataSourceDnsCNameRecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsCNameRecordReader_Read_4(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsCNameRecordReader_Read_4"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsCNameRecordAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	res, err := c.client.GetApiDnsCNameRecordAllWithResponse(
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

			resItems := DnsRecordListToTerraform(api_result)

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
		return diag.Errorf("No data found in db, insert one %s", "DnsCNameRecord")
	}

	tflog.Debug(ctx, "dataSourceDnsCNameRecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
