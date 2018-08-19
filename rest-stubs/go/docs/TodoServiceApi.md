# \TodoServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTodo**](TodoServiceApi.md#CreateTodo) | **Post** /v1/todo/records | 
[**DeleteTodo**](TodoServiceApi.md#DeleteTodo) | **Delete** /v1/todo/records/{id} | 
[**GetTodo**](TodoServiceApi.md#GetTodo) | **Get** /v1/todo/records/{id} | 
[**QueryTodo**](TodoServiceApi.md#QueryTodo) | **Get** /v1/todo/records | 
[**UpdateTodo**](TodoServiceApi.md#UpdateTodo) | **Put** /v1/todo/records/{id} | 


# **CreateTodo**
> TodoMutateResponse CreateTodo(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**TodoCreateRequest**](TodoCreateRequest.md)|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTodo**
> TodoMutateResponse DeleteTodo(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTodo**
> TodoGetResponse GetTodo(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 

### Return type

[**TodoGetResponse**](TodoGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QueryTodo**
> TodoQueryRequest QueryTodo(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TodoQueryRequest**](TodoQueryRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTodo**
> TodoMutateResponse UpdateTodo(ctx, id, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **int32**|  | 
  **body** | [**TodoUpdateRequest**](TodoUpdateRequest.md)|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

