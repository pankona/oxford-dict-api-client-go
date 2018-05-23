/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schema for wordlist endpoint.
type Wordlist struct {

	// Additional Information provided by OUP
	Metadata *interface{} `json:"metadata,omitempty"`

	// A list of found words
	Results []WordlistResults `json:"results,omitempty"`
}
