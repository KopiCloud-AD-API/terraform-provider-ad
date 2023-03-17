# KopiCloud AD API Terraform Provider

Terraform Provider for KopiCloud AD API

**Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live -Martin Golding-**

-----

List of Methods

## AD Computer

[ ] GET /api/Computers/All - List ALL Computers in AD

[ ] GET /api/Computers- List Computers inside an AD OU

[ ] GET /api/Computers/{ADComputerName} - Show AD Computer Details

[ ] GET /api/Computers/{ADComputerName}/Exists - Check If AD Computer Exists

[ ] POST /api/Computers/{ADComputerName}/Register - Register AD Computer

[ ] PUT /api/Computers/{ADComputerName}/Rename/{ADNewName} - Rename AD Computer

[ ] DELETE /api/Computers/{ADComputerName}/Remove - Remove AD Computer

[ ] DELETE /api/Computers/Remove - Remove Multiple AD Computers

[ ] DELETE /api/Computers/CleanUp - CleanUp Inactive AD Computers

###########################################

## AD Group Membership

[ ] GET /api/ADUser/{Username}/Group/All - Get AD User Group Membership

[ ] GET /api/ADUser/{Username}/Group/{GroupName} - Check If User Is a Member of an AD Group

[ ] POST /api/ADUser/{Username}/Group/{GroupName} - Add AD User To an AD Group

[ ] DELETE / api/ADUser/{Username}/Group/{GroupName} - Remove AD User From an AD Group

###########################################

## AD Groups

[ ] GET /api/ADGroups/All - List All AD Groups

[ ] GET /api/ADGroups - List AD Groups Inside an OU (Review)

[ ] GET /api/ADGroups/{GroupName}/Exists - Check If AD Group Exists

[ ] POST /api/ADGroups/{GroupName}/Security - Create an AD Security Group

[ ] POST /api/ADGroups/{GroupName}/Distribution - Create an AD Distribution Group

[ ] PUT /api/ADGroups/{GroupName}/Rename/{NewGroupName} - Rename an AD Group

[ ] DELETE /api/ADGroups/{GroupName}/Remove - Delete an AD Group

###########################################

## AD Organization Units (OU)

[ ] GET /api/OU/All - List All AD OUs

[ ] GET /api/OU/{OUName}/Details - Get AD OU Details

[ ] GET /api/OU/{OUName}/Exists - Check If AD OU Exists

[ ] POST /api/OU/{OUName} - Create AD OU

[ ] PUT /api/OU/{OUName}/Rename/{OUNewName} - Rename AD OU

[ ] PUT /api/OU - Update AD OU

[ ] DELETE /api/OU/{OUName}/Remove - Delete AD OU

###########################################

## AD Users

[ ] GET /api/ADUser/ListUsers/All - Get List of ALL AD Users

[ ] GET /api/ADUser/ListUsers - Get List of AD Users Inside an OU

[ ] GET /api/ADUser/{Username}/Exists - Check If AD User Exist

[ ] GET /api/ADUser/{Username}/LastLogon - Check If AD User Exist

[ ] POST /api/ADUser - Create an AD User

[ ] POST /api/ADUser/{Username}/Organization - Set AD User Organization Settings

[ ] POST /api/ADUser/{Username}/Organization/Clear - Clear AD User Organization Settings

[ ] POST /api/ADUser/{Username}/Address - Set AD User Address Settings

[ ] POST /api/ADUser/{Username}/Address/Clear - Clear AD User Address Settings

[ ] POST /api/ADUser/{Username}/Phone - Set AD User Phone Settings

[ ] POST /api/ADUser/{Username}/Phone/Clear - Clear AD User Phone Settings

[ ] POST /api/ADUser/{Username}/Profile - Set AD User Profile Settings

[ ] POST /api/ADUser/{Username}/Profile/Clear - Clear AD User Profile Settings

[ ] POST /api/ADUser/{Username}/RDSProfile - Set AD User RDS Profile Settings

[ ] POST /api/ADUser/{Username}/RDSProfile/Clear - Clear AD User RDS Profile Settings

[ ] POST /api/ADUser/{Username}/ResetPassword - Reset AD User Password

[ ] PUT /api/ADUser/{Username}/Enable - Enable AD User

[ ] PUT /api/ADUser/{Username}/Disable - Disable AD User

[ ] PUT /api/ADUser/{Username}/Unlock - Unlock AD User

[ ] PUT /api/ADUser/{Username}/Rename/{NewUsername} - Rename AD User

[ ] DELETE /api/ADUser/{Username} - Delete AD User

###########################################

## DNS Zones

[ ] GET /api/DnsZones/{ZoneName}/Exists - Check If DNS Zone Exists

[x] GET /api/DnsZones/All - List All DNS Zones

[x] POST /api/DnsLookupZone/{ZoneName} - Create a DNS Lookup Zone

[ ] POST /api/DnsReverseLookupZone/{ZoneName} - Create a DNS Reverse Lookup Zone

[x] DELETE /api/DnsZone/{ZoneName} - Remove a DNS Zone

###########################################

## DNS A Records

[x] GET /api/DnsARecord/All - List All DNS A Records in All Zones

[x] GET /api/DnsARecord - List All DNS A Records in a DNS Zone

[x] POST /api/DnsARecord - Create a DNS A Records

[x] DELETE /api/DnsARecord - Delete a DNS A Record

###########################################

## DNS AAAA Records

[x] GET /api/DnsAAAARecord/All - List All DNS AAAA Records in All Zones

[x] GET /api/DnsAAAARecord - List All DNS AAAA Records in a DNS Zone

[x] POST /api/DnsAAAARecord - Create a DNS AAAA Records

[x] DELETE /api/DnsAAAARecord - Delete a DNS A Record

###########################################

## DNS CNAME Records

[x] GET /api/DnsCNameRecord/All - List All DNS CNAME Records in All Zones

[x] GET /api/DnsCNameRecord - List All DNS CNAME Records in a DNS Zone

[x] POST /api/DnsCNameRecord - Create a DNS CNAME Records

[x] DELETE /api/DnsCNameRecord - Delete a DNS CNAME Record
