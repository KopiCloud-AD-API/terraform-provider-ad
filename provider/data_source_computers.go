package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComputers() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceComputersRead,
		Schema: map[string]*schema.Schema{
			"computers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"computer_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Computer name",
						},
						"created": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation Date",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description",
						},
						"dns_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Computer name",
						},
						"operating_system": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Computer name",
						},
						"path": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Computer name",
						},
					},
				},
			},
		},
	}
}
