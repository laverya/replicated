/*
 * Vendor API V1
 *
 * Apps documentation
 *
 * API version: 1.0.0
 * Contact: info@replicated.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type InlineResponse200 struct {
	Branding map[string]Branding `json:"branding,omitempty"`
	Channels []AppChannel `json:"channels,omitempty"`
}
