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
import (
	"time"
)

type OrgSetting struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Address1 string `json:"address1,omitempty"`
	Address2 string `json:"address2,omitempty"`
	City string `json:"city,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	Country string `json:"country,omitempty"`
	Created time.Time `json:"created,omitempty"`
	EndUserSupportHelpURL string `json:"endUserSupportHelpURL,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	Subdomain string `json:"subdomain,omitempty"`
	SupportPhoneNumber string `json:"supportPhoneNumber,omitempty"`
	Website string `json:"website,omitempty"`
}