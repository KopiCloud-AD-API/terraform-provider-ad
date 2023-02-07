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

// Generate the client for the API
//:generate oapi-codegen -package=api   -generate=types   -o api/kopicloud.types.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//:generate oapi-codegen -package=api   -generate=client  -o api/kopicloud.client.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json

// Generate the Terraform Schemas && converters from API Types to Terraform
//:generate openapi-terraform-provider-generator schemas     -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.schemas.gen.go     -i computer -i dnsrecord
//:generate openapi-terraform-provider-generator converters  -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.converters.gen.go  -i computer -i dnsrecord

// Generate the code for the datasources (Schemas && Code)
//:generate openapi-terraform-provider-generator datasources 																  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.datasources.gen.go 		 -i generator-inputs/computer.json -i generator-inputs/dns_a_record.json -i generator-inputs/dns_aaaa_record.json -i generator-inputs/dns_cname_record.json

// Generate the Code for Computers
//:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.computer.gen.go 	     -i generator-inputs/computer.json

// Generate the Code for the DNS A Records
//:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dns_a_record.gen.go 	 -i generator-inputs/dns_a_record.json

// Generate the Code for the DNS AAAA Records
//:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dns_aaaa_record.gen.go   -i generator-inputs/dns_aaaa_record.json

// Generate the Code for the DNS CName Records
//:generate openapi-terraform-provider-generator resource-crud 														      --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dns_cname_record.gen.go  -i generator-inputs/dns_cname_record.json

// Generate the Code for the DNS Lookup Zones
//:generate openapi-terraform-provider-generator resource-crud 														          --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dns_lookup_zone.gen.go   -i generator-inputs/dns_lookup_zone.json

// Generate the Code for the DNS Reverse Lookup Zones (Not yet ready)
//:generate openapi-terraform-provider-generator resource-crud 														          --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.dns_reverse_lookup_zone.gen.go   -i generator-inputs/dns_reverse_lookup_zone.json

// Generate the schemas for the resources (should list all the JSON files for which resources were generated)
//:generate openapi-terraform-provider-generator resource-schemas 														  --api gitlab.com/KopiCloud/kopicloud-ad-tf-provider/api -p provider -o provider/kopicloud.resources-schemas.gen.go -i generator-inputs/computer.json -i generator-inputs/dns_a_record.json -i generator-inputs/dns_aaaa_record.json -i generator-inputs/dns_cname_record.json -i generator-inputs/dns_lookup_zone.json -i generator-inputs/dns_reverse_lookup_zone.json

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
