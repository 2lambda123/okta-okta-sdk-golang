/*
 * Okta API
 *
 * Allows customers to easily access the Okta API
 *
 * API version: 2.8.1
 * Contact: devex-public@okta.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DnsRecordType string

// List of DNSRecordType
const (
	TXT_DnsRecordType DnsRecordType = "TXT"
	CNAME_DnsRecordType DnsRecordType = "CNAME"
)