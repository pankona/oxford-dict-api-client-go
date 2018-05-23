/* 
 * 
 *
 * No descripton provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.11.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

type ArrayOfRelatedEntriesInner struct {

	// A subject, discipline, or branch of knowledge particular to the Sense
	Domains Arrayofstrings `json:"domains,omitempty"`

	// The identifier of the word
	Id string `json:"id,omitempty"`

	// IANA language code specifying the language of the word
	Language string `json:"language,omitempty"`

	// A particular area in which the pronunciation occurs, e.g. 'Great Britain'
	Regions Arrayofstrings `json:"regions,omitempty"`

	// A level of language usage, typically with respect to formality. e.g. 'offensive', 'informal'
	Registers Arrayofstrings `json:"registers,omitempty"`

	Text string `json:"text,omitempty"`
}
