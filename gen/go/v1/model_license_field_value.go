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

type LicenseFieldValue struct {
	// The name of the custom field which you want to populate a value.
	Field string `json:"field"`
	// The value that you wish to populate the custom field with.
	Value string `json:"value"`
}
