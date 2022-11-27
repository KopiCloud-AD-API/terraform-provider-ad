package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	kcapi "kopicloud/api"
)

func resourceComputersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceComputersRead", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	params := kcapi.GetApiComputersAllParams{
		AuthToken: c.data.Get("token").(string),
	}
	res, err := c.client.GetApiComputersAllWithResponse(ctx, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	if res != nil {
		//As the return item is a []Computers, lets Unmarshal it into "computers"
		resItems := flattenComputers(res.JSON200)
		for _, computer := range resItems {
			tflog.Debug(ctx, fmt.Sprintf("converted computer: %v", computer))
		}
		if err := d.Set("computers", resItems); err != nil {
			return diag.FromErr(err)
		}
		d.SetId("all_computers")
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	tflog.Debug(ctx, fmt.Sprintf("converted computers: %v", d.Get("computers")))
	tflog.Debug(ctx, "[DEBUG] %s: resourceComputereRead finished successfully"+d.Id())
	return diags
}

func flattenComputers(computersList *[]kcapi.Computer) []interface{} {
	if computersList != nil {
		computers := make([]interface{}, len(*computersList))
		for i, computer := range *computersList {
			cl := make(map[string]interface{})

			cl["computer_name"] = computer.ComputerName
			cl["created"] = computer.Created
			cl["description"] = computer.Description
			cl["dns_name"] = computer.DnsName
			cl["operating_system"] = computer.OperatingSystem
			cl["path"] = computer.Path

			computers[i] = cl
		}
		return computers
	}
	return make([]interface{}, 0)
}

// func resourceComputers() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext:   resourceComputersRead,
// 		Schema: map[string]*schema.Schema{

// 		},
// 	}
// }

// func resourceComputersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	// Warning or errors can be collected in a slice type
// 	log.Printf("[DEBUG] %s: Beginning resourceAvengersRead", d.Id())
// 	var diags diag.Diagnostics
// 	c := m.(*ApiClient)
// 	res, err := c.client.GetApiComputersAll(ctx)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	if res != nil {
// 		//As the return item is a []Avengers, lets Unmarshal it into "avengers"
// 		resItems := flattenAvengers(&res)
// 		if err := d.Set("avengers", resItems); err != nil {
// 			return diag.FromErr(err)
// 		}
// 	} else {
// 		return diag.Errorf("no data found in db, insert one")
// 	}
// 	log.Printf("[DEBUG] %s: resourceAvengersRead finished successfully", d.Id())
// 	return diags
// }
