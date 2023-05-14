package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaOfSingleValueList(t schema.ValueType, field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprint("List of Single Field Object"),
		Elem: &schema.Resource{
			Schema: schemaMapOfScalarElement(t, field_name),
		},
	}
}

func schemaOfSingleValueObject(t schema.ValueType, field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprint("List of Single Field Object"),
		Elem: &schema.Resource{
			Schema: schemaMapOfScalarElement(t, field_name),
		},
	}
}

func schemaMapOfScalarElement(t schema.ValueType, field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		field_name: &schema.Schema{
			Type:     t,
			Computed: true,
		},
	}
}

func schemaOfOUList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "OU"),
		Elem: &schema.Resource{
			Schema: schemaMapOfOU(scalar_field_name),
		},
	}
}

func schemaOfOU(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "OU"),
		Elem: &schema.Resource{
			Schema: schemaMapOfOU(scalar_field_name),
		},
	}
}

func schemaMapOfOU(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"protected": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},

		"guid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}

func schemaOfDnsZoneList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "DnsZone"),
		Elem: &schema.Resource{
			Schema: schemaMapOfDnsZone(scalar_field_name),
		},
	}
}

func schemaOfDnsZone(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "DnsZone"),
		Elem: &schema.Resource{
			Schema: schemaMapOfDnsZone(scalar_field_name),
		},
	}
}

func schemaMapOfDnsZone(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"distinguished_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"zone_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"zone_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}

func schemaOfDnsRecordList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "DnsRecord"),
		Elem: &schema.Resource{
			Schema: schemaMapOfDnsRecord(scalar_field_name),
		},
	}
}

func schemaOfDnsRecord(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "DnsRecord"),
		Elem: &schema.Resource{
			Schema: schemaMapOfDnsRecord(scalar_field_name),
		},
	}
}

func schemaMapOfDnsRecord(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"zone": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"timestamp": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"data": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}

func schemaOfUserList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "User"),
		Elem: &schema.Resource{
			Schema: schemaMapOfUser(scalar_field_name),
		},
	}
}

func schemaOfUser(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "User"),
		Elem: &schema.Resource{
			Schema: schemaMapOfUser(scalar_field_name),
		},
	}
}

func schemaMapOfUser(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"first_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"password_never_expired": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},

		"job_title": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"home_folder_directory": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"rds_profile_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"last_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"company": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"change_password_next_logon": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},

		"street_po_box": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"city": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"office_phone": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"rds_allow_logon": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},

		"guid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"state": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"department": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"mobile_phone": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"profile_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"email_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"sam_username": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"home_phone": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"manager": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"country": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"home_folder_drive": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"rds_home_folder_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"rds_home_folder_drive": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"rds_connect_drive": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "",
		},

		"initials": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"ou_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"street_address": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"username": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"display_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"office": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"postal_code": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"profile_logon_script": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"home_folder_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}

func schemaOfGroupList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "Group"),
		Elem: &schema.Resource{
			Schema: schemaMapOfGroup(scalar_field_name),
		},
	}
}

func schemaOfGroup(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "Group"),
		Elem: &schema.Resource{
			Schema: schemaMapOfGroup(scalar_field_name),
		},
	}
}

func schemaMapOfGroup(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"email": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"ou_path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"guid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"scope": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}

func schemaOfComputerList(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("List of %s", "Computer"),
		Elem: &schema.Resource{
			Schema: schemaMapOfComputer(scalar_field_name),
		},
	}
}

func schemaOfComputer(scalar_field_name string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: fmt.Sprintf("Single Element List of %s", "Computer"),
		Elem: &schema.Resource{
			Schema: schemaMapOfComputer(scalar_field_name),
		},
	}
}

func schemaMapOfComputer(scalar_field_name string) map[string]*schema.Schema {
	return map[string]*schema.Schema{

		"operating_system": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"dns_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"path": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"created": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"sid": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},

		"computer_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "",
		},
	}
}
