/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Entry struct {

	// The origin of the word and the way in which its meaning has changed throughout history
	Etymologies *Arrayofstrings `json:"etymologies,omitempty"`

	GrammaticalFeatures *GrammaticalFeaturesList `json:"grammaticalFeatures,omitempty"`

	// Identifies the homograph grouping. The last two digits identify different entries of the same homograph. The first one/two digits identify the homograph number.
	HomographNumber string `json:"homographNumber,omitempty"`

	Notes *CategorizedTextList `json:"notes,omitempty"`

	Pronunciations *PronunciationsList `json:"pronunciations,omitempty"`

	// Complete list of senses
	Senses []Sense `json:"senses,omitempty"`

	// Various words that are used interchangeably depending on the context, e.g 'a' and 'an'
	VariantForms *VariantFormsList `json:"variantForms,omitempty"`
}
