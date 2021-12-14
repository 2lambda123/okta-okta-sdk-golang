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

type IonForm struct {
	Accepts string `json:"accepts,omitempty"`
	Href string `json:"href,omitempty"`
	Method string `json:"method,omitempty"`
	Name string `json:"name,omitempty"`
	Produces string `json:"produces,omitempty"`
	Refresh int32 `json:"refresh,omitempty"`
	Rel []string `json:"rel,omitempty"`
	RelatesTo []string `json:"relatesTo,omitempty"`
	Value []IonField `json:"value,omitempty"`
}