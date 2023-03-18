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

func resourceDnsARecord() *schema.Resource {
	terraformSchema := make(map[string]*schema.Schema)

	terraformSchema["result"] = schemaOfDnsRecord(``)

	terraformSchema["hostname"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["ip_address"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	terraformSchema["zone_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
		Required: false,

		ForceNew: true,

		Description: "",
	}

	return &schema.Resource{

		CreateContext: resourceDnsARecordCreate,

		ReadContext: resourceDnsARecordRead,

		DeleteContext: resourceDnsARecordDelete,

		Schema:      terraformSchema,
		Description: "",
	}
}

func resourceDnsARecordCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsARecordRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	hostname := d.Get("hostname").(string)

	ip_address := d.Get("ip_address").(string)

	zone_name := d.Get("zone_name").(string)

	params := kcapi.PostApiDnsARecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		IPAddress: &ip_address,

		ZoneName: &zone_name,
	}

	res, err := c.client.PostApiDnsARecordWithResponse(
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

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsARecord: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_DnsARecord(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "resourceDnsARecordcreate finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsARecordRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsARecordRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
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

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsARecord: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId(getId_for_DnsARecord(api_result))

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "resourceDnsARecordread finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}

func resourceDnsARecordDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("Beginning resourceDnsARecordRead"))
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	tflog.Debug(ctx, "Terraform data", map[string]interface{}{
		"Terraform ID": d.Id(),
		"client_data":  c.data,
		"schema_data":  d,
	})

	hostname := d.Get("hostname").(string)

	ip_address := d.Get("ip_address").(string)

	zone_name := d.Get("zone_name").(string)

	params := kcapi.DeleteApiDnsARecordParams{
		AuthToken: c.data.Get("token").(string),

		DNSHostName: &hostname,

		IPAddress: &ip_address,

		ZoneName: &zone_name,
	}

	res, err := c.client.DeleteApiDnsARecordWithResponse(
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

			for _, element := range resItems {
				tflog.Debug(ctx, "converted DnsARecord: %#v", element.(map[string]interface{}))
			}

			result := wrapInArray(resItems)

			if err := d.Set("result", result); err != nil {
				return diag.FromErr(err)
			}

			d.SetId("")

		}
	} else {
		return diag.Errorf("No data found in db, insert one %s", "DnsARecord")
	}

	tflog.Debug(ctx, "resourceDnsARecorddelete finished successfully",
		map[string]interface{}{
			"result": d,
		})
	return diags
}