package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	cli "github.com/jawher/mow.cli"

	kcp "gitlab.com/KopiCloud/kopicloud-ad-tf-provider/provider"
)

// :generate openapi-terraform-provider-generator resource-schemas 														  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.resources-schemas.gen.go -i generator-inputs/dns-records.json -i generator-inputs/dns-zones.json
// :generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dnsrecords.gen.go 		 -i generator-inputs/dns-records.json
// :generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dnszones.gen.go 		 -i generator-inputs/dns-zones.json
//
//go:generate oapi-codegen -package=api   -generate=types   -o api/kopicloud.types.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate oapi-codegen -package=api   -generate=client  -o api/kopicloud.client.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate openapi-terraform-provider-generator schemas     -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.schemas.gen.go     -i computer
//go:generate openapi-terraform-provider-generator converters  -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.converters.gen.go  -i computer
//go:generate openapi-terraform-provider-generator datasources 																  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.datasources.gen.go -i generator-inputs/datasources.json
var (
	build_version string
	build_date    string
	build_time    string
)

func version() string {
	version := fmt.Sprintf("Build Version: %s\n", build_version)
	version += fmt.Sprintf("Build Date: %s - %s\n", build_date, build_time)
	version += "\n"

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			version += fmt.Sprintf("%s = %s\n", setting.Key, setting.Value)
		}
	}
	return version
}

func main() {
	app := cli.App("openapi-terrafor-generator", "Generate code for Terraform providers based on an OpenAPI schema and input files")

	app.Version("v version", version())

	app.Action = func() {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() *schema.Provider {
				return kcp.Provider()
			},
		})
	}

	app.Run(os.Args)
}
