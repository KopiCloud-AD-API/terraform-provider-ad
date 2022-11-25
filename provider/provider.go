package kopicloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		// First some scafolding
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_HOST", ""),
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_TOKEN", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

type ApiClient struct {
	data *schema.ResourceData
}

func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	//	client, err := c.NewKopicloudClient()
	//	if err != nil {
	//		return c, diag.FromErr(err)
	//	}
	//	c.kopicloudclient = client
	return c, nil

}
