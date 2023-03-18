package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func checkRequiredArgumentsSet(ctx context.Context, d *schema.ResourceData, args []string) bool {
	for _, arg := range args {
		_, ok := d.GetOk(arg)
		if !ok {
			return false
		}
	}
	return true
}

func dataSources() map[string]*schema.Resource {
	return map[string]*schema.Resource{

		"kopicloud_computer_list": dataSourceComputer(),

		"kopicloud_ou_list": dataSourceOU(),

		"kopicloud_dns_a_records_list": dataSourceDnsARecordReader(),

		"kopicloud_dns_aaaa_records_list": dataSourceDnsAAAARecordReader(),

		"kopicloud_dns_cname_records_list": dataSourceDnsCNameRecordReader(),

		"kopicloud_dns_zone_list": dataSourceDnsZoneReader(),

		"kopicloud_dns_lookup_zone_list": dataSourceDnsLookupZoneReader(),

		"kopicloud_dns_reverse_lookup_zone_list": dataSourceDnsReverseLookupZoneReader(),

		"kopicloud_group_list": dataSourceGroup(),

		"kopicloud_security_group_list": dataSourceSecurityGroup(),

		"kopicloud_distribution_group_list": dataSourceDistributioGroup(),

		"kopicloud_group_membership_list": dataSourceGroupMembership(),

		"kopicloud_user_list": dataSourceUser(),
	}
}
