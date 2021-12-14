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

type GroupSchema struct {
	Schema string `json:"$schema,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Created string `json:"created,omitempty"`
	Definitions *GroupSchemaDefinitions `json:"definitions,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Properties *UserSchemaProperties `json:"properties,omitempty"`
	Title string `json:"title,omitempty"`
	Type_ string `json:"type,omitempty"`
}