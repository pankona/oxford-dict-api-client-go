/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A lexical sense represents the lexical meaning of a lexical entry when interpreted as referring to the corresponding ontology element
type ThesaurusSense struct {

	// antonym of word
	Antonyms *SynonymsAntonyms `json:"antonyms,omitempty"`

	// A subject, discipline, or branch of knowledge particular to the Sense
	Domains *Arrayofstrings `json:"domains,omitempty"`

	Examples *ExamplesList `json:"examples,omitempty"`

	// The id of the sense that is required for the delete procedure
	Id string `json:"id,omitempty"`

	// A particular area in which the Sense occurs, e.g. 'Great Britain'
	Regions *Arrayofstrings `json:"regions,omitempty"`

	// A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal'
	Registers *Arrayofstrings `json:"registers,omitempty"`

	// subsenses of word
	Subsenses []ThesaurusSense `json:"subsenses,omitempty"`

	// synonym of word
	Synonyms *SynonymsAntonyms `json:"synonyms,omitempty"`
}
