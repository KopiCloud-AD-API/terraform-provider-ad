package provider

import (
	"context"
	"fmt"
	kcapi "kopicloud/api"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceComputer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceComputerCreate,
		ReadContext:   resourceComputerRead,
		Schema:        schemaOfComputer(),
	}
}

func resourceComputerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s: Beginning resourceComputersRead", d.Id()))

	var diags diag.Diagnostics
	c := m.(*ApiClient)
	params := kcapi.PostApiComputersADComputerNameRegisterParams{
		OUPath:    new(string),
		AuthToken: c.data.Get("token").(string),
	}

	_, err := c.client.PostApiComputersADComputerNameRegisterWithResponse(ctx, "", &params)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceComputerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	tflog.Debug(ctx, fmt.Sprintf("[DEBUG] %s: Beginning resourceComputersRead", d.Id()))

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
		resItems := flattenComputerList(res.JSON200)
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
