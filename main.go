package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	kcp "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/provider"
)

//go:generate oapi-codegen -package=api   -generate=types   -o api/kopicloud.types.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate oapi-codegen -package=api   -generate=client  -o api/kopicloud.client.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate openapi-terraform-provider-generator schemas     -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.schemas.gen.go     -i computer
//go:generate openapi-terraform-provider-generator converters  -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.converters.gen.go  -i computer
//go:generate openapi-terraform-provider-generator datasources 																  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.datasources.gen.go -i generator-inputs/datasources.json
//go:generate openapi-terraform-provider-generator resource-schemas 														  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.resources-schemas.gen.go -i generator-inputs/dns-records.json -i generator-inputs/dns-zones.json
//go:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dnsrecords.gen.go 		 -i generator-inputs/dns-records.json
//go:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dnszones.gen.go 		 -i generator-inputs/dns-zones.json

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return kcp.Provider()
		},
	})
}
