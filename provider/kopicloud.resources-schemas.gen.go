package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func checkFieldsModified(d *schema.ResourceData, args []string) bool {
	return len(args) == 0 || d.HasChanges(args...)
}

func getFieldValue(field string, useOld bool, d *schema.ResourceData) interface{} {
	old, new := d.GetChange(field)
	if useOld {
		return old
	} else {
		return new
	}
}

func resources() map[string]*schema.Resource {
	return map[string]*schema.Resource{

		"kopicloud_computer": resourceComputer(),

		"kopicloud_ou": resourceOU(),

		"kopicloud_dns_a_record": resourceDnsARecord(),

		"kopicloud_dns_aaaa_record": resourceDnsAAAARecord(),

		"kopicloud_dns_cname_record": resourceDnsCNameRecord(),

		"kopicloud_dns_lookup_zone": resourceDnsLookupZone(),

		"kopicloud_dns_reverse_lookup_zone": resourceDnsReverseLookupZone(),

		"kopicloud_security_group": resourceGroup(),

		"kopicloud_distribution_group": resourceDistributionGroup(),

		"kopicloud_group_membership": resourceGroupMembership(),

		"kopicloud_user": resourceUser(),

		"kopicloud_user_password_reset": resourceUserPasswordReset(),

		"kopicloud_user_enable_account": resourceUserEnableAccount(),

		"kopicloud_user_disable_account": resourceUserDisableAccount(),

		"kopicloud_user_unlock_account": resourceUserUnlockAccount(),

		"kopicloud_user_rename_account": resourceUserRename(),
	}
}
