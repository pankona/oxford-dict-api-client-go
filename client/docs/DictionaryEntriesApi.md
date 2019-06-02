# \DictionaryEntriesApi

All URIs are relative to *https://od-api.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceLangWordIdFiltersGet**](DictionaryEntriesApi.md#EntriesSourceLangWordIdFiltersGet) | **Get** /entries/{source_lang}/{word_id}/{filters} | Apply filters to response
[**EntriesSourceLangWordIdGet**](DictionaryEntriesApi.md#EntriesSourceLangWordIdGet) | **Get** /entries/{source_lang}/{word_id} | Retrieve dictionary information for a given word
[**EntriesSourceLangWordIdRegionsregionGet**](DictionaryEntriesApi.md#EntriesSourceLangWordIdRegionsregionGet) | **Get** /entries/{source_lang}/{word_id}/regions&#x3D;{region} | Specify GB or US dictionary for English entry search


# **EntriesSourceLangWordIdFiltersGet**
> RetrieveEntry EntriesSourceLangWordIdFiltersGet(ctx, sourceLang, wordId, filters, appId, appKey)
Apply filters to response

 Use filters to limit the [entry](documentation/glossary?term=entry) information that is returned. For example, you may only require definitions and not everything else, or just [pronunciations](documentation/glossary?term=pronunciation). The full list of filters can be retrieved from the filters Utility endpoint. You can also specify values within the filter using '='. For example 'grammaticalFeatures=singular'. Filters can also be combined using a semicolon.    <div id=\"dictionary_entries_filters\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **filters** | **string**| Separate filtering conditions using a semicolon. Conditions take values grammaticalFeatures and/or lexicalCategory and are case-sensitive. To list multiple values in single condition divide them with comma. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**RetrieveEntry**](RetrieveEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdGet**
> RetrieveEntry EntriesSourceLangWordIdGet(ctx, sourceLang, wordId, appId, appKey)
Retrieve dictionary information for a given word

 Use this to retrieve definitions, [pronunciations](documentation/glossary?term=pronunciation), example sentences, [grammatical information](documentation/glossary?term=grammaticalfeatures) and [word origins](documentation/glossary?term=etymology). It only works for dictionary [headwords](documentation/glossary?term=headword), so you may need to use the [Lemmatron](documentation/glossary?term=lemma) first if your input is likely to be an [inflected](documentation/glossary?term=inflection) form (e.g., 'swimming'). This would return the linked [headword](documentation/glossary?term=headword) (e.g., 'swim') which you can then use in the Entries endpoint. Unless specified using a region filter, the default lookup will be the Oxford Dictionary of English (GB).    <div id=\"dictionary_entries\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**RetrieveEntry**](RetrieveEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdRegionsregionGet**
> RetrieveEntry EntriesSourceLangWordIdRegionsregionGet(ctx, sourceLang, wordId, region, appId, appKey)
Specify GB or US dictionary for English entry search

 USe this filter to restrict the lookup to either our Oxford Dictionary of English (GB) or New Oxford American Dictionary (US). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **wordId** | **string**| An Entry identifier. Case-sensitive. | 
  **region** | **string**| Region filter parameter. gb &#x3D; Oxford Dictionary of English. us &#x3D; New Oxford American Dictionary. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**RetrieveEntry**](RetrieveEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

