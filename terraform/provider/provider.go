package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	kcapi "kopicloud/api"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		// First some scafolding
		DataSourcesMap: dataSources(),
		// ResourcesMap: map[string]*schema.Resource{
		// 	"kopicloud_computer": resourceComputer(),
		// },
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Kopiclcoud Server URL",
				DefaultFunc: schema.EnvDefaultFunc("API_HOST", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bearer (JWT) or Basic Authentication Token",
				DefaultFunc: schema.EnvDefaultFunc("API_AUTHENTICATION_TOKEN", ""),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

type ApiClient struct {
	data   *schema.ResourceData
	client *kcapi.ClientWithResponses
}

func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	client, err := c.NewKopiCloudClient()
	if err != nil {
		return c, diag.FromErr(err)
	}
	c.client = client
	return c, nil

}

func (a *ApiClient) NewKopiCloudClient() (*kcapi.ClientWithResponses, error) {
	host := a.data.Get("host").(string)
	c, err := kcapi.NewClientWithResponses(host)
	if err != nil {
		return c, err
	}
	return c, nil
}
