/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A mapping of filters available per endpoints.
type FiltersResults struct {

	// A list of filters available for Retrieve Entry endpoint
	Entries *Arrayofstrings `json:"entries,omitempty"`

	// A list of filters available for LEMMATRON endpoint
	Inflections *Arrayofstrings `json:"inflections,omitempty"`

	// A list of filters available for Translations endpoint
	Translations *Arrayofstrings `json:"translations,omitempty"`

	// A list of filters available for Translations endpoint
	Wordlist *Arrayofstrings `json:"wordlist,omitempty"`
}
