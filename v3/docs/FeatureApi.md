# {{classname}}

All URIs are relative to *https://{subdomain}.{domain}*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeature**](FeatureApi.md#GetFeature) | **Get** /api/v1/features/{featureId} | 
[**ListFeatureDependencies**](FeatureApi.md#ListFeatureDependencies) | **Get** /api/v1/features/{featureId}/dependencies | 
[**ListFeatureDependents**](FeatureApi.md#ListFeatureDependents) | **Get** /api/v1/features/{featureId}/dependents | 
[**ListFeatures**](FeatureApi.md#ListFeatures) | **Get** /api/v1/features | 
[**UpdateFeatureLifecycle**](FeatureApi.md#UpdateFeatureLifecycle) | **Post** /api/v1/features/{featureId}/{lifecycle} | 

# **GetFeature**
> Feature GetFeature(ctx, featureId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureId** | **string**|  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFeatureDependencies**
> []Feature ListFeatureDependencies(ctx, featureId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureId** | **string**|  | 

### Return type

[**[]Feature**](Feature.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFeatureDependents**
> []Feature ListFeatureDependents(ctx, featureId)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureId** | **string**|  | 

### Return type

[**[]Feature**](Feature.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFeatures**
> []Feature ListFeatures(ctx, )


Success

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Feature**](Feature.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFeatureLifecycle**
> Feature UpdateFeatureLifecycle(ctx, featureId, lifecycle, optional)


Success

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **featureId** | **string**|  | 
  **lifecycle** | **string**|  | 
 **optional** | ***FeatureApiUpdateFeatureLifecycleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeatureApiUpdateFeatureLifecycleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  | 

### Return type

[**Feature**](Feature.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
