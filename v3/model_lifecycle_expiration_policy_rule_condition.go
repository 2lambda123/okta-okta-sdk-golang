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

type LifecycleExpirationPolicyRuleCondition struct {
	LifecycleStatus string `json:"lifecycleStatus,omitempty"`
	Number int32 `json:"number,omitempty"`
	Unit string `json:"unit,omitempty"`
}