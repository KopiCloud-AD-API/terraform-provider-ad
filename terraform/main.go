package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	kcp "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/provider"
)

//go:generate oapi-codegen -package=api   -generate=types   -o api/kopicloud.types.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate oapi-codegen -package=api   -generate=client  -o api/kopicloud.client.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate openapi-terraform-provider-generator -api=gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -package=provider -generate=schemas -o provider/kopicloud.schemas.gen.go -url https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate openapi-terraform-provider-generator -api=gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -package=provider -generate=converters -o provider/kopicloud.converters.gen.go -url https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate ../bin/kopicloud-ad-tf-provider-generator -i generator-inputs/datasources.json -o provider/datasources.gen.go -generate datasources
//go:generate ../bin/kopicloud-ad-tf-provider-generator -i generator-inputs/resources.json -o provider/resources.gen.go -generate resources

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return kcp.Provider()
		},
	})
}
