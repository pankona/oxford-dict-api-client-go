/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Schema for lexi-stats results for a word/trueCase/lemma/lexicalCategory returned as a frequency
type StatsWordResult struct {

	// Additional Information provided by OUP
	Metadata *interface{} `json:"metadata,omitempty"`

	Result *StatsWordResultResult `json:"result,omitempty"`
}
