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

type UserFactor struct {
	Embedded map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Created time.Time `json:"created,omitempty"`
	FactorType *FactorType `json:"factorType,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Provider *FactorProvider `json:"provider,omitempty"`
	Status *FactorStatus `json:"status,omitempty"`
	Verify *VerifyFactorRequest `json:"verify,omitempty"`
}