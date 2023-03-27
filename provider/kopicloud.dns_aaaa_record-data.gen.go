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

func dataSourceDnsAAAARecordReader() *schema.Resource {
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

	terraformSchema["ipv6_address"] = &schema.Schema{
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
		ReadContext: dataSourceDnsAAAARecordReaderRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceDnsAAAARecordReaderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 3)

	required_inputs[0] = "hostname"

	required_inputs[1] = "ipv6_address"

	required_inputs[2] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsAAAARecordReader_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "hostname"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsAAAARecordReader_Read_1(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "ipv6_address"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsAAAARecordReader_Read_2(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsAAAARecordReader_Read_3(ctx, d, m)
	}

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsAAAARecordReader_Read_4(ctx, d, m)
	}

	return nil
}

func dataSourceDnsAAAARecordReader_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsAAAARecordReader_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	hostname := d.Get("hostname").(string)

	ipv6_address := d.Get("ipv6_address").(string)

	zone_name := d.Get("zone_name").(string)

	params := kcapi.GetApiDnsAAAARecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		IPv6Address: &ipv6_address,

		ZoneName: &zone_name,
	}

	res, err := c.client.GetApiDnsAAAARecordWithResponse(
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
					tflog.Debug(ctx, "converted DnsAAAARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsAAAARecord",
					map[string]interface{}{
						"DnsAAAARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsAAAARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsAAAARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsAAAARecordReader_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsAAAARecordReader_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsAAAARecordHostNameDNSHostNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	hostname := d.Get("hostname").(string)

	res, err := c.client.GetApiDnsAAAARecordHostNameDNSHostNameWithResponse(
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
					tflog.Debug(ctx, "converted DnsAAAARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsAAAARecord",
					map[string]interface{}{
						"DnsAAAARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsAAAARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsAAAARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsAAAARecordReader_Read_2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsAAAARecordReader_Read_2"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsAAAARecordIPv6AddressIPv6AddressParams{
		AuthToken: c.data.Get("token").(string),
	}

	ipv6_address := d.Get("ipv6_address").(string)

	res, err := c.client.GetApiDnsAAAARecordIPv6AddressIPv6AddressWithResponse(
		ctx,

		ipv6_address,

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
					tflog.Debug(ctx, "converted DnsAAAARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsAAAARecord",
					map[string]interface{}{
						"DnsAAAARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsAAAARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsAAAARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsAAAARecordReader_Read_3(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsAAAARecordReader_Read_3"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsAAAARecordZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.GetApiDnsAAAARecordZoneNameWithResponse(
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
					tflog.Debug(ctx, "converted DnsAAAARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsAAAARecord",
					map[string]interface{}{
						"DnsAAAARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsAAAARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsAAAARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsAAAARecordReader_Read_4(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsAAAARecordReader_Read_4"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsAAAARecordAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	res, err := c.client.GetApiDnsAAAARecordAllWithResponse(
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
					tflog.Debug(ctx, "converted DnsAAAARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsAAAARecord",
					map[string]interface{}{
						"DnsAAAARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsAAAARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsAAAARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
