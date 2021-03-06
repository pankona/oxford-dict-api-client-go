/*
 *
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.11.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/context"
)

// Linger please
var (
	_ context.Context
)

type ThesaurusApiService service

/* ThesaurusApiService Retrieve words that mean the opposite
 Retrieve words that are opposite in meaning to the input word ([antonym](documentation/glossary?term&#x3D;thesaurus)).    &lt;div id&#x3D;\&quot;antonyms\&quot;&gt;&lt;/div&gt;
* @param ctx context.Context for authentication, logging, tracing, etc.
@param sourceLang IANA language code
@param wordId An Entry identifier. Case-sensitive.
@param appId App ID Authentication Parameter
@param appKey App Key Authentication Parameter
@return Thesaurus*/
func (a *ThesaurusApiService) EntriesSourceLangWordIdAntonymsGet(ctx context.Context, sourceLang string, wordId string, appId string, appKey string) (Thesaurus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Thesaurus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entries/{source_lang}/{word_id}/antonyms"
	localVarPath = strings.Replace(localVarPath, "{"+"source_lang"+"}", fmt.Sprintf("%v", sourceLang), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word_id"+"}", fmt.Sprintf("%v", wordId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["app_id"] = parameterToString(appId, "")
	localVarHeaderParams["app_key"] = parameterToString(appKey, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ThesaurusApiService Retrieve words that are similar
 Use this to retrieve words that are similar in meaning to the input word ([synonym](documentation/glossary?term&#x3D;synonym)).    &lt;div id&#x3D;\&quot;synonyms\&quot;&gt;&lt;/div&gt;
* @param ctx context.Context for authentication, logging, tracing, etc.
@param sourceLang IANA language code
@param wordId An Entry identifier. Case-sensitive.
@param appId App ID Authentication Parameter
@param appKey App Key Authentication Parameter
@return Thesaurus*/
func (a *ThesaurusApiService) EntriesSourceLangWordIdSynonymsGet(ctx context.Context, sourceLang string, wordId string, appId string, appKey string) (Thesaurus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Thesaurus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entries/{source_lang}/{word_id}/synonyms"
	localVarPath = strings.Replace(localVarPath, "{"+"source_lang"+"}", fmt.Sprintf("%v", sourceLang), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word_id"+"}", fmt.Sprintf("%v", wordId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["app_id"] = parameterToString(appId, "")
	localVarHeaderParams["app_key"] = parameterToString(appKey, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ThesaurusApiService Retrieve synonyms and antonyms for a given word
 Retrieve available [synonyms](documentation/glossary?term&#x3D;thesaurus) and [antonyms](documentation/glossary?term&#x3D;thesaurus) for a given word and language.     &lt;div id&#x3D;\&quot;synonyms_and_antonyms\&quot;&gt;&lt;/div&gt;
* @param ctx context.Context for authentication, logging, tracing, etc.
@param sourceLang IANA language code
@param wordId An Entry identifier. Case-sensitive.
@param appId App ID Authentication Parameter
@param appKey App Key Authentication Parameter
@return Thesaurus*/
func (a *ThesaurusApiService) EntriesSourceLangWordIdSynonymsantonymsGet(ctx context.Context, sourceLang string, wordId string, appId string, appKey string) (Thesaurus, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     Thesaurus
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/entries/{source_lang}/{word_id}/synonyms;antonyms"
	localVarPath = strings.Replace(localVarPath, "{"+"source_lang"+"}", fmt.Sprintf("%v", sourceLang), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"word_id"+"}", fmt.Sprintf("%v", wordId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["app_id"] = parameterToString(appId, "")
	localVarHeaderParams["app_key"] = parameterToString(appKey, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
