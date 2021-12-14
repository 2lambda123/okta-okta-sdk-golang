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

type SmsTemplate struct {
	Created time.Time `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Template string `json:"template,omitempty"`
	Translations *SmsTemplateTranslations `json:"translations,omitempty"`
	Type_ *SmsTemplateType `json:"type,omitempty"`
}