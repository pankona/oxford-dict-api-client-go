# \WordlistApi

All URIs are relative to *https://od-api-demo.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WordlistSourceLangFiltersAdvancedGet**](WordlistApi.md#WordlistSourceLangFiltersAdvancedGet) | **Get** /wordlist/{source_lang}/{filters_advanced} | Retrieve list of words for category with advanced options
[**WordlistSourceLangFiltersBasicGet**](WordlistApi.md#WordlistSourceLangFiltersBasicGet) | **Get** /wordlist/{source_lang}/{filters_basic} | Retrieve a list of words in a category


# **WordlistSourceLangFiltersAdvancedGet**
> Wordlist WordlistSourceLangFiltersAdvancedGet($sourceLang, $filtersAdvanced, $appId, $appKey, $exclude, $excludeSenses, $excludePrimeSenses, $wordLength, $prefix, $exact, $limit, $offset)

Retrieve list of words for category with advanced options

Use this to apply more complex filters to the [list of words](documentation/glossary?term=wordlist). For example, you may only want to filter out words for which all [senses](documentation/glossary?term=sense) match the filter, or only its 'prime sense'. You can also filter by word length or match by substring (prefix).     <div id=\"wordlist_advanced\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **filtersAdvanced** | **string**| Semicolon separated list of wordlist parameters, presented as value pairs: LexicalCategory, domains, regions, registers. Parameters can take comma separated list of values. E.g., lexicalCategory&#x3D;noun,adjective;domains&#x3D;sport. Number of values limited to 5. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]
 **exclude** | **string**| Semicolon separated list of parameters-value pairs (same as filters). Excludes headwords that have any senses in specified exclusion attributes (lexical categories, domains, etc.) from results. | [optional] 
 **excludeSenses** | **string**| Semicolon separated list of parameters-value pairs (same as filters). Excludes only those senses of a particular headword that match specified exclusion attributes (lexical categories, domains, etc.) from results but includes the headword if it has other permitted senses. | [optional] 
 **excludePrimeSenses** | **string**| Semicolon separated list of parameters-value pairs (same as filters). Excludes a headword only if the primary sense matches the specified exclusion attributes (registers, domains only). | [optional] 
 **wordLength** | **string**| Parameter to speficy the minimum (&gt;), exact or maximum (&lt;) length of the words required. E.g., &gt;5 - more than 5 chars; &lt;4 - less than 4 chars; &gt;5&lt;10 - from 5 to 10 chars; 3 - exactly 3 chars. | [optional] [default to &gt;5,&lt;10]
 **prefix** | **string**| Filter words that start with prefix parameter | [optional] [default to goal]
 **exact** | **bool**| If exact&#x3D;true wordlist returns a list of entries that exactly matches the search string. Otherwise wordlist lists entries that start with prefix string. | [optional] [default to false]
 **limit** | **string**| Limit the number of results per response. Default and maximum limit is 5000. | [optional] 
 **offset** | **string**| Offset the start number of the result. | [optional] 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WordlistSourceLangFiltersBasicGet**
> Wordlist WordlistSourceLangFiltersBasicGet($sourceLang, $filtersBasic, $appId, $appKey, $limit, $offset)

Retrieve a list of words in a category

 Use this to retrieve a [list of words](documentation/glossary?term=wordlist) for particular [domain](documentation/glossary?term=domain), [lexical category](documentation/glossary?term=lexicalcategory), [register](documentation/glossary?term=registers) and/or [region](documentation/glossary?term=regions). View the full list of possible filters using the filters Utility endpoint.  The response only includes [headwords](documentation/glossary?term=headword), not all their possible [inflections](documentation/glossary?term=inflection). If you require a full [wordlist](documentation/glossary?term=wordlist) including [inflected forms](documentation/glossary?term=inflection), contact us and we can help.    <div id=\"wordlist\"></div> 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceLang** | **string**| IANA language code | 
 **filtersBasic** | **string**| Semicolon separated list of wordlist parameters, presented as value pairs: LexicalCategory, domains, regions, registers. Parameters can take comma separated list of values. E.g., lexicalCategory&#x3D;noun,adjective;domains&#x3D;sport. Number of values limited to 5. | 
 **appId** | **string**| App ID Authentication Parameter | [default to 5037d509]
 **appKey** | **string**| App Key Authentication Parameter | [default to 4dc1aebaa63721f0f8e79a55e2514bc7]
 **limit** | **string**| Limit the number of results per response. Default and maximum limit is 5000. | [optional] 
 **offset** | **string**| Offset the start number of the result | [optional] 

### Return type

[**Wordlist**](Wordlist.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

