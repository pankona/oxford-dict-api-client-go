# \TheSentenceDictionaryApi

All URIs are relative to *https://od-api.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceLanguageWordIdSentencesGet**](TheSentenceDictionaryApi.md#EntriesSourceLanguageWordIdSentencesGet) | **Get** /entries/{source_language}/{word_id}/sentences | Retrieve corpus sentences for a given word


# **EntriesSourceLanguageWordIdSentencesGet**
> SentencesResults EntriesSourceLanguageWordIdSentencesGet(ctx, sourceLanguage, wordId, appId, appKey)
Retrieve corpus sentences for a given word

 Use this to retrieve sentences extracted from  corpora which show how a word is used in the language. This is available for English and Spanish. For English, the sentences are linked to the correct [sense](documentation/glossary?term=sense) of the word in the dictionary. In Spanish, they are linked at the [headword](documentation/glossary?term=headword) level.   <div id=\"sentences\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLanguage** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**SentencesResults**](SentencesResults.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

