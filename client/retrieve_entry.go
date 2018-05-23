/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schema for the 'entries' endpoints
type RetrieveEntry struct {

	// Additional Information provided by OUP
	Metadata *interface{} `json:"metadata,omitempty"`

	// A list of entries and all the data related to them
	Results []HeadwordEntry `json:"results,omitempty"`
}