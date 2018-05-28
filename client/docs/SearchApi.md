# \SearchApi

All URIs are relative to *https://od-api.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchSourceLangGet**](SearchApi.md#SearchSourceLangGet) | **Get** /search/{source_lang} | Retrieve possible matches to input
[**SearchSourceSearchLanguageTranslationstargetSearchLanguageGet**](SearchApi.md#SearchSourceSearchLanguageTranslationstargetSearchLanguageGet) | **Get** /search/{source_search_language}/translations&#x3D;{target_search_language} | Retrieve possible translation matches to input


# **SearchSourceLangGet**
> Wordlist SearchSourceLangGet(ctx, sourceLang, appId, appKey, optional)
Retrieve possible matches to input

 Use this to retrieve possible [headword](documentation/glossary?term=headword) matches for a given string of text. The results are culculated using headword matching, fuzzy matching, and [lemmatization](documentation/glossary?term=lemma)     <div id=\"search\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLang** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **appId** | **string**| App ID Authentication Parameter | [default to ]
 **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **q** | **string**| Search string | [default to eye]
 **prefix** | **bool**| Set prefix to true if you&#39;d like to get results only starting with search string. | [default to false]
 **regions** | **string**| If searching in English, filter words with specific region(s) either &#39;us&#39; or &#39;gb&#39;. | 
 **limit** | **string**| Limit the number of results per response. Default and maximum limit is 5000. | 
 **offset** | **string**| Offset the start number of the result. | 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchSourceSearchLanguageTranslationstargetSearchLanguageGet**
> Wordlist SearchSourceSearchLanguageTranslationstargetSearchLanguageGet(ctx, sourceSearchLanguage, targetSearchLanguage, appId, appKey, optional)
Retrieve possible translation matches to input

 Use this to find matches in our translation dictionaries.    <div id=\"search_translation\"></div> 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceSearchLanguage** | **string**| IANA language code | 
  **targetSearchLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceSearchLanguage** | **string**| IANA language code | 
 **targetSearchLanguage** | **string**| IANA language code | 
 **appId** | **string**| App ID Authentication Parameter | [default to ]
 **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **q** | **string**| Search string. | [default to eye]
 **prefix** | **bool**| Set prefix to true if you&#39;d like to get results only starting with search string. | [default to false]
 **regions** | **string**| Filter words with specific region(s) E.g., regions&#x3D;us. For now gb, us are available for en language. | 
 **limit** | **string**| Limit the number of results per response. Default and maximum limit is 5000. | 
 **offset** | **string**| Offset the start number of the result. | 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

