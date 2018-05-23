# \LemmatronApi

All URIs are relative to *https://od-api-demo.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InflectionsSourceLangWordIdFiltersGet**](LemmatronApi.md#InflectionsSourceLangWordIdFiltersGet) | **Get** /inflections/{source_lang}/{word_id}/{filters} | Apply optional filters to Lemmatron response
[**InflectionsSourceLangWordIdGet**](LemmatronApi.md#InflectionsSourceLangWordIdGet) | **Get** /inflections/{source_lang}/{word_id} | Check a word exists in the dictionary and retrieve its root form


# **InflectionsSourceLangWordIdFiltersGet**
> Lemmatron InflectionsSourceLangWordIdFiltersGet($sourceLang, $wordId, $filters, $appId, $appKey)

Apply optional filters to Lemmatron response

 Retrieve available [lemmas](documentation/glossary?term=lemma) for a given [inflected](documentation/glossary?term=inflection) wordform. Filter results by categories.      <div id=\"lemmatron_filters\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **wordId** | **string**| The input word | 
 **filters** | [**[]string**](string.md)| Separate filtering conditions using a semicolon. Conditions take values grammaticalFeatures and/or lexicalCategory and are case-sensitive. To list multiple values in single condition divide them with comma. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**Lemmatron**](Lemmatron.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InflectionsSourceLangWordIdGet**
> Lemmatron InflectionsSourceLangWordIdGet($sourceLang, $wordId, $appId, $appKey)

Check a word exists in the dictionary and retrieve its root form

 Use this to check if a word exists in the dictionary, or what 'root' form it links to (e.g., swimming > swim). The response tells you the possible [lemmas](documentation/glossary?term=lemma) for a given [inflected](documentation/glossary?term=inflection) word. This can then be combined with other endpoints to retrieve more information.    <div id=\"lemmatron\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **wordId** | **string**| The input word | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**Lemmatron**](Lemmatron.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

