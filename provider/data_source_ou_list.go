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

func dataSourceOuList() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceOuListRead,
		Schema:      schemaOfOUModelList(),
	}
}

func resourceOuListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceOuListRead", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	params := kcapi.GetApiOUAllParams{
		AuthToken: c.data.Get("token").(string),
	}
	res, err := c.client.GetApiOUAllWithResponse(ctx, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	if res != nil {
		//As the return item is a []Computers, lets Unmarshal it into "computers"
		resItems := FlattenOUModelList(res.JSON200)
		for _, ou := range resItems {
			tflog.Debug(ctx, fmt.Sprintf("converted OU: %v", ou))
		}
		if err := d.Set("oumodel_list", resItems); err != nil {
			return diag.FromErr(err)
		}
		d.SetId("all_ou")
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	tflog.Debug(ctx, fmt.Sprintf("converted OU: %v", d.Get("oumodel_list")))
	tflog.Debug(ctx, "[DEBUG] %s: resourceOuListRead finished successfully"+d.Id())
	return diags
}
