/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type TranslationsListInner struct {

	// A subject, discipline, or branch of knowledge particular to the translation
	Domains *Arrayofstrings `json:"domains,omitempty"`

	GrammaticalFeatures *GrammaticalFeaturesList `json:"grammaticalFeatures,omitempty"`

	// IANA language code specifying the language of the translation
	Language string `json:"language"`

	Notes *CategorizedTextList `json:"notes,omitempty"`

	// A particular area in which the translation occurs, e.g. 'Great Britain'
	Regions *Arrayofstrings `json:"regions,omitempty"`

	// A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal'
	Registers *Arrayofstrings `json:"registers,omitempty"`

	Text string `json:"text"`
}