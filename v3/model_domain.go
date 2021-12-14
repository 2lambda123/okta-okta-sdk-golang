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

type Domain struct {
	CertificateSourceType *DomainCertificateSourceType `json:"certificateSourceType,omitempty"`
	DnsRecords []DnsRecord `json:"dnsRecords,omitempty"`
	Domain string `json:"domain,omitempty"`
	Id string `json:"id,omitempty"`
	PublicCertificate *DomainCertificateMetadata `json:"publicCertificate,omitempty"`
	ValidationStatus *DomainValidationStatus `json:"validationStatus,omitempty"`
}