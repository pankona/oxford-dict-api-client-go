/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schema for the Filters endpoint.
type Filters struct {

	// Additional Information provided by OUP
	Metadata *interface{} `json:"metadata,omitempty"`

	Results *FiltersResults `json:"results,omitempty"`
}
