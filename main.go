package main

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	cli "github.com/jawher/mow.cli"

	kcp "github.com/KopiCloud-AD-API/terraform-provider-ad/provider"
)

// Generate the client for the API
//go:generate oapi-codegen -package=api   -generate=types   -o api/kopicloud.types.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json
//go:generate oapi-codegen -package=api   -generate=client  -o api/kopicloud.client.gen.go https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json

// Generate the Terraform Schemas && converters from API Types to Terraform
//go:generate openapi-terraform-provider-generator schemas     -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.schemas.gen.go     		   -i Computer -i ou -i DnsRecord -i DnsZone -i Group -i User -e _IncludeAllProperties -e _ShowFields
//go:generate openapi-terraform-provider-generator converters  -u https://labapi.kopicloud-ad-api.com/swagger/v1/swagger.json --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.converters.gen.go  		   -i Computer -i ou -i DnsRecord -i DnsZone -i Group -i User -e _IncludeAllProperties -e _ShowFields

// Generate the Code for Computers
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.computer-crud.gen.go 	     -i generator-inputs/computer.json
//go:generate openapi-terraform-provider-generator resource-update 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.computer-update.gen.go 		 -i generator-inputs/computer.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.computer-data.gen.go 	     -i generator-inputs/computer.json

// Generate the Code for Computers Cleanup
//:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.computer-cleanup-crud.gen.go -i generator-inputs/computer_cleanup.json

// Generate the Code for Groups Memberships
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.group-memberships-crud.gen.go 	 -i generator-inputs/group_membership.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.group-memberships-data.gen.go 	 -i generator-inputs/group_membership.json

// Generate the Code for Security Groups
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.security-group-crud.gen.go 	 -i generator-inputs/security_group.json
//go:generate openapi-terraform-provider-generator resource-update 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.security-group-update.gen.go -i generator-inputs/security_group.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.security-group-data.gen.go 	 -i generator-inputs/security_group.json

// Generate the Code for Distribution Groups
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.distribution-group-crud.gen.go 	 -i generator-inputs/distribution_group.json
//go:generate openapi-terraform-provider-generator resource-update 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.distribution-group-update.gen.go -i generator-inputs/distribution_group.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.distribution-group-data.gen.go 	 -i generator-inputs/distribution_group.json

// Generate the Code for Organizational Units
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.ou-crud.gen.go 	     		 -i generator-inputs/organizational_unit.json
//go:generate openapi-terraform-provider-generator resource-update 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.ou-update.gen.go 	     	 -i generator-inputs/organizational_unit.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.ou-data.gen.go 	 			 -i generator-inputs/organizational_unit.json

// Generate the Code for Users
//:generate openapi-terraform-provider-generator resource-create 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-create.gen.go 	     	 -i generator-inputs/user.json
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-crud.gen.go 	     	 -i generator-inputs/user.json
//go:generate openapi-terraform-provider-generator resource-update 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-update.gen.go 		 	 -i generator-inputs/user.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-data.gen.go 	     	 -i generator-inputs/user.json

// Generate the Code for User Password Reset
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-password-reset-crud.gen.go -i generator-inputs/user_password_reset.json

// Generate the Code for User Enabling
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-enable-crud.gen.go -i generator-inputs/user_enable.json

// Generate the Code for User Disabling
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-disable-crud.gen.go -i generator-inputs/user_disable.json

// Generate the Code for User Unlocking
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-unlock-crud.gen.go -i generator-inputs/user_unlock.json

// Generate the Code for User Renaming
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.user-rename.gen.go -i generator-inputs/user_rename.json

// Generate the Code for the DNS A Records
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_a_record-crud.gen.go 	 -i generator-inputs/dns_a_record.json
//go:generate openapi-terraform-provider-generator datasource-read													          --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_a_record-data.gen.go 	 -i generator-inputs/dns_a_record.json

// Generate the Code for the DNS AAAA Records
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_aaaa_record-crud.gen.go   -i generator-inputs/dns_aaaa_record.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_aaaa_record-data.gen.go   -i generator-inputs/dns_aaaa_record.json

// Generate the Code for the DNS CName Records
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_cname_record-crud.gen.go  -i generator-inputs/dns_cname_record.json
//go:generate openapi-terraform-provider-generator datasource-read 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_cname_record-data.gen.go  -i generator-inputs/dns_cname_record.json

// Generate the Code for the DNS Lookup Zones
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_lookup_zone-crud.gen.go   -i generator-inputs/dns_lookup_zone.json
//go:generate openapi-terraform-provider-generator datasource-read  														  --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_lookup_zone-data.gen.go   -i generator-inputs/dns_lookup_zone.json

// Generate the Code for the DNS Reverse Lookup Zones (Not yet ready)
//go:generate openapi-terraform-provider-generator resource-crud 														      --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_reverse_lookup_zone-crud.gen.go   -i generator-inputs/dns_reverse_lookup_zone.json
//go:generate openapi-terraform-provider-generator datasource-read  														  --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.dns_reverse_lookup_zone-data.gen.go  -i generator-inputs/dns_reverse_lookup_zone.json

// Generate the schemas for the resources (should list all the JSON files for which resources were generated)
//go:generate openapi-terraform-provider-generator resource-schemas 														  --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.resources-schemas.gen.go 	-i generator-inputs/computer.json -i generator-inputs/organizational_unit.json -i generator-inputs/dns_a_record.json -i generator-inputs/dns_aaaa_record.json -i generator-inputs/dns_cname_record.json -i generator-inputs/dns_lookup_zone.json -i generator-inputs/dns_reverse_lookup_zone.json -i generator-inputs/security_group.json -i generator-inputs/distribution_group.json -i generator-inputs/group_membership.json -i generator-inputs/user.json -i generator-inputs/user_password_reset.json -i generator-inputs/user_enable.json -i generator-inputs/user_disable.json -i generator-inputs/user_unlock.json -i generator-inputs/user_rename.json
// -i generator-inputs/computer_cleanup.json

// Generate the code for the datasources (Schemas)
//go:generate openapi-terraform-provider-generator datasource-schemas 														  --api github.com/KopiCloud-AD-API/terraform-provider-ad/api -p provider -o provider/kopicloud.datasources-schemas.gen.go 	-i generator-inputs/computer.json -i generator-inputs/organizational_unit.json -i generator-inputs/dns_a_record.json -i generator-inputs/dns_aaaa_record.json -i generator-inputs/dns_cname_record.json -i generator-inputs/dns_lookup_zone.json -i generator-inputs/dns_reverse_lookup_zone.json -i generator-inputs/security_group.json -i generator-inputs/distribution_group.json -i generator-inputs/group_membership.json -i generator-inputs/user.json

//go:generate goimports -w provider

// Generate the Terraform provider documentation using `tfplugindocs`:
//:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

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
