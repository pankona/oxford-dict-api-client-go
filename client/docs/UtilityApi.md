# \UtilityApi

All URIs are relative to *https://od-api.oxforddictionaries.com:443/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsSourceDomainsLanguageTargetDomainsLanguageGet**](UtilityApi.md#DomainsSourceDomainsLanguageTargetDomainsLanguageGet) | **Get** /domains/{source_domains_language}/{target_domains_language} | Lists available domains in a bilingual dataset
[**DomainsSourceLanguageGet**](UtilityApi.md#DomainsSourceLanguageGet) | **Get** /domains/{source_language} | Lists available domains in a monolingual dataset
[**FiltersEndpointGet**](UtilityApi.md#FiltersEndpointGet) | **Get** /filters/{endpoint} | Lists available filters for specific endpoint
[**FiltersGet**](UtilityApi.md#FiltersGet) | **Get** /filters | Lists available filters
[**GrammaticalFeaturesSourceLanguageGet**](UtilityApi.md#GrammaticalFeaturesSourceLanguageGet) | **Get** /grammaticalFeatures/{source_language} | Lists available grammatical features in a dataset
[**LanguagesGet**](UtilityApi.md#LanguagesGet) | **Get** /languages | Lists available dictionaries
[**LexicalcategoriesLanguageGet**](UtilityApi.md#LexicalcategoriesLanguageGet) | **Get** /lexicalcategories/{language} | Lists available lexical categories in a dataset
[**RegionsSourceLanguageGet**](UtilityApi.md#RegionsSourceLanguageGet) | **Get** /regions/{source_language} | Lists available regions in a monolingual dataset
[**RegistersSourceLanguageGet**](UtilityApi.md#RegistersSourceLanguageGet) | **Get** /registers/{source_language} | Lists available registers in a  monolingual dataset
[**RegistersSourceRegisterLanguageTargetRegisterLanguageGet**](UtilityApi.md#RegistersSourceRegisterLanguageTargetRegisterLanguageGet) | **Get** /registers/{source_register_language}/{target_register_language} | Lists available registers in a bilingual dataset


# **DomainsSourceDomainsLanguageTargetDomainsLanguageGet**
> UtilityLabels DomainsSourceDomainsLanguageTargetDomainsLanguageGet(ctx, sourceDomainsLanguage, targetDomainsLanguage, appId, appKey)
Lists available domains in a bilingual dataset

Returns a list of the available [domains](documentation/glossary?term=domain) for a given bilingual language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceDomainsLanguage** | **string**| IANA language code | 
  **targetDomainsLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DomainsSourceLanguageGet**
> UtilityLabels DomainsSourceLanguageGet(ctx, sourceLanguage, appId, appKey)
Lists available domains in a monolingual dataset

Returns a list of the available [domains](documentation/glossary?term=domain) for a given monolingual language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FiltersEndpointGet**
> Filters FiltersEndpointGet(ctx, endpoint, appId, appKey)
Lists available filters for specific endpoint

Returns a list of all the valid filters for a given endpoint to construct API calls. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **endpoint** | **string**| Name of the endpoint. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FiltersGet**
> Filters FiltersGet(ctx, appId, appKey)
Lists available filters

Returns a list of all the valid filters to construct API calls. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Filters**](Filters.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrammaticalFeaturesSourceLanguageGet**
> UtilityLabels GrammaticalFeaturesSourceLanguageGet(ctx, sourceLanguage, appId, appKey)
Lists available grammatical features in a dataset

Returns a list of the available [grammatical features](documentation/glossary?term=grammaticalfeatures) for a given language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLanguage** | **string**| IANA language code. If provided output will be filtered by sourceLanguage. | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LanguagesGet**
> Languages LanguagesGet(ctx, appId, appKey, optional)
Lists available dictionaries

Returns a list of monolingual and bilingual language datasets available in the API 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appId** | **string**| App ID Authentication Parameter | [default to ]
 **appKey** | **string**| App Key Authentication Parameter | [default to ]
 **sourceLanguage** | **string**| IANA language code. If provided output will be filtered by sourceLanguage. | 
 **targetLanguage** | **string**| IANA language code. If provided output will be filtered by sourceLanguage. | 

### Return type

[**Languages**](Languages.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LexicalcategoriesLanguageGet**
> UtilityLabels LexicalcategoriesLanguageGet(ctx, language, appId, appKey)
Lists available lexical categories in a dataset

Returns a list of available [lexical categories](documentation/glossary?term=lexicalcategory) for a given language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **language** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegionsSourceLanguageGet**
> Regions RegionsSourceLanguageGet(ctx, sourceLanguage, appId, appKey)
Lists available regions in a monolingual dataset

Returns a list of the available [regions](documentation/glossary?term=regions) for a given monolingual language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**Regions**](Regions.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistersSourceLanguageGet**
> UtilityLabels RegistersSourceLanguageGet(ctx, sourceLanguage, appId, appKey)
Lists available registers in a  monolingual dataset

Returns a list of the available [registers](documentation/glossary?term=registers) for a given monolingual language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistersSourceRegisterLanguageTargetRegisterLanguageGet**
> UtilityLabels RegistersSourceRegisterLanguageTargetRegisterLanguageGet(ctx, sourceRegisterLanguage, targetRegisterLanguage, appId, appKey)
Lists available registers in a bilingual dataset

Returns a list of the available [registers](documentation/glossary?term=registers) for a given bilingual language dataset. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **sourceRegisterLanguage** | **string**| IANA language code | 
  **targetRegisterLanguage** | **string**| IANA language code | 
  **appId** | **string**| App ID Authentication Parameter | [default to ]
  **appKey** | **string**| App Key Authentication Parameter | [default to ]

### Return type

[**UtilityLabels**](UtilityLabels.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

