/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Description of an entry for a particular part of speech
type ThesaurusLexicalEntry struct {
	Entries []ThesaurusEntry `json:"entries,omitempty"`

	// IANA language code
	Language string `json:"language"`

	// A linguistic category of words (or more precisely lexical items), generally defined by the syntactic or morphological behaviour of the lexical item in question, such as noun or verb
	LexicalCategory string `json:"lexicalCategory"`

	// A given written or spoken realisation of a an entry.
	Text string `json:"text"`

	// Various words that are used interchangeably depending on the context, e.g 'a' and 'an'
	VariantForms *VariantFormsList `json:"variantForms,omitempty"`
}