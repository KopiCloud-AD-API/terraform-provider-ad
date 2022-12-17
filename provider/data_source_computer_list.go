package provider

import (
	"context"
	"fmt"
	kcapi "kopicloud/api"
	"log"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComputerList() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceComputersRead,
		Schema:      schemaOfComputerList(),
	}
}

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
		resItems := FlattenComputerList(res.JSON200)
		for _, computer := range resItems {
			tflog.Debug(ctx, fmt.Sprintf("converted computer: %v", computer))
		}
		if err := d.Set("computer_list", resItems); err != nil {
			return diag.FromErr(err)
		}
		d.SetId("all_computers")
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	tflog.Debug(ctx, fmt.Sprintf("converted computers: %v", d.Get("computer_list")))
	tflog.Debug(ctx, "[DEBUG] %s: resourceComputereRead finished successfully"+d.Id())
	return diags
}
