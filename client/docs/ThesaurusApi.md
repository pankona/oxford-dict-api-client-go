# \ThesaurusApi

All URIs are relative to *https://od-api-demo.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceLangWordIdAntonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdAntonymsGet) | **Get** /entries/{source_lang}/{word_id}/antonyms | Retrieve words that mean the opposite
[**EntriesSourceLangWordIdSynonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdSynonymsGet) | **Get** /entries/{source_lang}/{word_id}/synonyms | Retrieve words that are similar
[**EntriesSourceLangWordIdSynonymsantonymsGet**](ThesaurusApi.md#EntriesSourceLangWordIdSynonymsantonymsGet) | **Get** /entries/{source_lang}/{word_id}/synonyms;antonyms | Retrieve synonyms and antonyms for a given word


# **EntriesSourceLangWordIdAntonymsGet**
> Thesaurus EntriesSourceLangWordIdAntonymsGet($sourceLang, $wordId, $appId, $appKey)

Retrieve words that mean the opposite

 Retrieve words that are opposite in meaning to the input word ([antonym](documentation/glossary?term=thesaurus)).    <div id=\"antonyms\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **wordId** | **string**| An Entry identifier. Case-sensitive. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdSynonymsGet**
> Thesaurus EntriesSourceLangWordIdSynonymsGet($sourceLang, $wordId, $appId, $appKey)

Retrieve words that are similar

 Use this to retrieve words that are similar in meaning to the input word ([synonym](documentation/glossary?term=synonym)).    <div id=\"synonyms\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **wordId** | **string**| An Entry identifier. Case-sensitive. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EntriesSourceLangWordIdSynonymsantonymsGet**
> Thesaurus EntriesSourceLangWordIdSynonymsantonymsGet($sourceLang, $wordId, $appId, $appKey)

Retrieve synonyms and antonyms for a given word

 Retrieve available [synonyms](documentation/glossary?term=thesaurus) and [antonyms](documentation/glossary?term=thesaurus) for a given word and language.     <div id=\"synonyms_and_antonyms\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **wordId** | **string**| An Entry identifier. Case-sensitive. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**Thesaurus**](Thesaurus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

