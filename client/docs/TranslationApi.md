# \TranslationApi

All URIs are relative to *https://od-api-demo.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EntriesSourceTranslationLanguageWordIdTranslationstargetTranslationLanguageGet**](TranslationApi.md#EntriesSourceTranslationLanguageWordIdTranslationstargetTranslationLanguageGet) | **Get** /entries/{source_translation_language}/{word_id}/translations&#x3D;{target_translation_language} | Retrieve translation for a given word


# **EntriesSourceTranslationLanguageWordIdTranslationstargetTranslationLanguageGet**
> RetrieveEntry EntriesSourceTranslationLanguageWordIdTranslationstargetTranslationLanguageGet($sourceTranslationLanguage, $wordId, $targetTranslationLanguage, $appId, $appKey)

Retrieve translation for a given word

 Use this to return translations for a given word. In the event that a word in the dataset does not have a direct translation, the response will be a [definition](documentation/glossary?term=entry) in the target language.    <div id=\"translation\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceTranslationLanguage** | **string**| IANA language code | 
 **wordId** | **string**| The source word | 
 **targetTranslationLanguage** | **string**| IANA language code | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]

### Return type

[**RetrieveEntry**](RetrieveEntry.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

