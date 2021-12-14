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

type SwaThreeFieldApplicationSettingsApplication struct {
	ButtonSelector string `json:"buttonSelector,omitempty"`
	ExtraFieldSelector string `json:"extraFieldSelector,omitempty"`
	ExtraFieldValue string `json:"extraFieldValue,omitempty"`
	LoginUrlRegex string `json:"loginUrlRegex,omitempty"`
	PasswordSelector string `json:"passwordSelector,omitempty"`
	TargetURL string `json:"targetURL,omitempty"`
	UserNameSelector string `json:"userNameSelector,omitempty"`
}