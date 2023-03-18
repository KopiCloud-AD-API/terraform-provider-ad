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

func dataSourceDnsARecordReader() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsRecordList(``)

	terraformSchema["hostname"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	terraformSchema["ip_address"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	terraformSchema["zone_name"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    false,
		Optional:    true,
		Required:    false,
		Description: "",
	}

	return &schema.Resource{
		ReadContext: dataSourceDnsARecordReaderRead,
		Schema:      terraformSchema,
		Description: "",
	}
}

func dataSourceDnsARecordReaderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var required_inputs []string

	required_inputs = make([]string, 3)

	required_inputs[0] = "hostname"

	required_inputs[1] = "ip_address"

	required_inputs[2] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsARecordReader_Read_0(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "hostname"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsARecordReader_Read_1(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "ip_address"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsARecordReader_Read_2(ctx, d, m)
	}

	required_inputs = make([]string, 1)

	required_inputs[0] = "zone_name"

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsARecordReader_Read_3(ctx, d, m)
	}

	required_inputs = make([]string, 0)

	if checkRequiredArgumentsSet(ctx, d, required_inputs) {
		return dataSourceDnsARecordReader_Read_4(ctx, d, m)
	}

	return nil
}

func dataSourceDnsARecordReader_Read_0(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsARecordReader_Read_0"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	hostname := d.Get("hostname").(string)

	ip_address := d.Get("ip_address").(string)

	zone_name := d.Get("zone_name").(string)

	params := kcapi.GetApiDnsARecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		IPAddress: &ip_address,

		ZoneName: &zone_name,
	}

	res, err := c.client.GetApiDnsARecordWithResponse(
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
					tflog.Debug(ctx, "converted DnsARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsARecord",
					map[string]interface{}{
						"DnsARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsARecordReader_Read_1(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsARecordReader_Read_1"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsARecordHostNameDNSHostNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	hostname := d.Get("hostname").(string)

	res, err := c.client.GetApiDnsARecordHostNameDNSHostNameWithResponse(
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
					tflog.Debug(ctx, "converted DnsARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsARecord",
					map[string]interface{}{
						"DnsARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsARecordReader_Read_2(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsARecordReader_Read_2"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsARecordIPAddressIPAddressParams{
		AuthToken: c.data.Get("token").(string),
	}

	ip_address := d.Get("ip_address").(string)

	res, err := c.client.GetApiDnsARecordIPAddressIPAddressWithResponse(
		ctx,

		ip_address,

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
					tflog.Debug(ctx, "converted DnsARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsARecord",
					map[string]interface{}{
						"DnsARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsARecordReader_Read_3(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsARecordReader_Read_3"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsARecordZoneNameParams{
		AuthToken: c.data.Get("token").(string),
	}

	zone_name := d.Get("zone_name").(string)

	res, err := c.client.GetApiDnsARecordZoneNameWithResponse(
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
					tflog.Debug(ctx, "converted DnsARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsARecord",
					map[string]interface{}{
						"DnsARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func dataSourceDnsARecordReader_Read_4(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning dataSourceDnsARecordReader_Read_4"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"client_data": c.data,
		"schema_data": d,
	})

	params := kcapi.GetApiDnsARecordAllParams{
		AuthToken: c.data.Get("token").(string),
	}

	res, err := c.client.GetApiDnsARecordAllWithResponse(
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
					tflog.Debug(ctx, "converted DnsARecord }: %#v", element.(map[string]interface{}))
				}
			default:
				tflog.Debug(ctx, "converted DnsARecord",
					map[string]interface{}{
						"DnsARecord": resItems,
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
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "dataSourceDnsARecordReaderRead finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}
