package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	api "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api"
	kcapi "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		// First some scafolding
		DataSourcesMap: dataSources(),
		ResourcesMap:   resources(),
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

func schemaOfStringList(list_name string, field_name string, description string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		list_name: {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					field_name: {
						Type:        schema.TypeString,
						Computed:    true,
						Description: description,
					},
				},
			},
		},
	}
}

func wrapInArray(e interface{}) []interface{} {
	return []interface{}{e}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

func getId_for_SingleValue(name string) func(*string) string {
	return func(value *string) string {
		return fmt.Sprintf("%s_%s", name, value)
	}
}

func getId_for_Computer(obj *api.Computer) string {
	return fmt.Sprintf("Computer_%s_%s", *obj.ComputerName)
}

func getId_for_DnsARecords(obj *api.DnsRecords) string {
	return fmt.Sprintf("DnsARecord_%s_%s", *obj.Name, *obj.Data)
}

func getId_for_DnsAAAARecords(obj *api.DnsRecords) string {
	return fmt.Sprintf("DnsAAAARecord_%s_%s_%s", *obj.Zone, *obj.Name, *obj.Data)
}

func getId_for_DnsCNameRecords(obj *api.DnsRecords) string {
	return fmt.Sprintf("DnsCNameRecord_%s_%s", *obj.Name, *obj.Data)
}

func getId_for_DnsLookupZone(name *string) string {
	return fmt.Sprintf("DnsLookupZone_%s", *name)
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
