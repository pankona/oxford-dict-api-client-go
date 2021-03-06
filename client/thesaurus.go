/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schema for thesaurus endpoint
type Thesaurus struct {

	// Additional Information provided by OUP
	Metadata *interface{} `json:"metadata,omitempty"`

	// A list of found synonyms or antonyms
	Results []HeadwordThesaurus `json:"results,omitempty"`
}
