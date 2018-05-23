/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Link to a sense of a specific entry in the thesaurus Dictionary
type ThesaurusLink struct {

	// identifier of a word
	EntryId string `json:"entry_id"`

	// identifier of a sense
	SenseId string `json:"sense_id"`
}