# \ThesaurusApi

All URIs are relative to *https://od-api.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceLangWordIdAntonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdAntonymsGet) | **Get** /entries/{source_lang}/{word_id}/antonyms | Retrieve words that mean the opposite
[**EntriesSourceLangWordIdSynonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdSynonymsGet) | **Get** /entries/{source_lang}/{word_id}/synonyms | Retrieve words that are similar
[**EntriesSourceLangWordIdSynonymsantonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdSynonymsantonymsGet) | **Get** /entries/{source_lang}/{word_id}/synonyms;antonyms | Retrieve synonyms and antonyms for a given word


# **EntriesSourceLangWordIdAntonymsGet**
> Thesaurus EntriesSourceLangWordIdAntonymsGet(ctx, sourceLang, wordId, appId, appKey)
Retrieve words that mean the opposite

 Retrieve words that are opposite in meaning to the input word ([antonym](documentation/glossary?term=thesaurus)).    <div id=\"antonyms\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdSynonymsGet**
> Thesaurus EntriesSourceLangWordIdSynonymsGet(ctx, sourceLang, wordId, appId, appKey)
Retrieve words that are similar

 Use this to retrieve words that are similar in meaning to the input word ([synonym](documentation/glossary?term=synonym)).    <div id=\"synonyms\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdSynonymsantonymsGet**
> Thesaurus EntriesSourceLangWordIdSynonymsantonymsGet(ctx, sourceLang, wordId, appId, appKey)
Retrieve synonyms and antonyms for a given word

 Retrieve available [synonyms](documentation/glossary?term=thesaurus) and [antonyms](documentation/glossary?term=thesaurus) for a given word and language.     <div id=\"synonyms_and_antonyms\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

