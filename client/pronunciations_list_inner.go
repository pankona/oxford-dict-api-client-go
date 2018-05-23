/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// A grouping of pronunciation information
type PronunciationsListInner struct {

	// The URL of the sound file
	AudioFile string `json:"audioFile,omitempty"`

	// A local or regional variation where the pronunciation occurs, e.g. 'British English'
	Dialects *Arrayofstrings `json:"dialects,omitempty"`

	// The alphabetic system used to display the phonetic spelling
	PhoneticNotation string `json:"phoneticNotation,omitempty"`

	// Phonetic spelling is the representation of vocal sounds which express pronunciations of words. It is a system of spelling in which each letter represents invariably the same spoken sound
	PhoneticSpelling string `json:"phoneticSpelling,omitempty"`

	// A particular area in which the pronunciation occurs, e.g. 'Great Britain'
	Regions *Arrayofstrings `json:"regions,omitempty"`
}