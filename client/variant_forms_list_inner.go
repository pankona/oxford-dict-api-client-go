/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type VariantFormsListInner struct {

	// A particular area in which the variant form occurs, e.g. 'Great Britain'
	Regions *Arrayofstrings `json:"regions,omitempty"`

	Text string `json:"text"`
}
