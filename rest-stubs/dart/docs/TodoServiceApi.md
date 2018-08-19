# swagger.api.TodoServiceApi

## Load the API package
```dart
import 'package:swagger/api.dart';
```

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTodo**](TodoServiceApi.md#createTodo) | **POST** /v1/todo/records | 
[**deleteTodo**](TodoServiceApi.md#deleteTodo) | **DELETE** /v1/todo/records/{id} | 
[**getTodo**](TodoServiceApi.md#getTodo) | **GET** /v1/todo/records/{id} | 
[**queryTodo**](TodoServiceApi.md#queryTodo) | **GET** /v1/todo/records | 
[**updateTodo**](TodoServiceApi.md#updateTodo) | **PUT** /v1/todo/records/{id} | 


# **createTodo**
> TodoMutateResponse createTodo(body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new TodoServiceApi();
var body = new TodoCreateRequest(); // TodoCreateRequest | 

try { 
    var result = api_instance.createTodo(body);
    print(result);
} catch (e) {
    print("Exception when calling TodoServiceApi->createTodo: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TodoCreateRequest**](TodoCreateRequest.md)|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deleteTodo**
> TodoMutateResponse deleteTodo(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new TodoServiceApi();
var id = 56; // int | 

try { 
    var result = api_instance.deleteTodo(id);
    print(result);
} catch (e) {
    print("Exception when calling TodoServiceApi->deleteTodo: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getTodo**
> TodoGetResponse getTodo(id)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new TodoServiceApi();
var id = 56; // int | 

try { 
    var result = api_instance.getTodo(id);
    print(result);
} catch (e) {
    print("Exception when calling TodoServiceApi->getTodo: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 

### Return type

[**TodoGetResponse**](TodoGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **queryTodo**
> TodoQueryRequest queryTodo()



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new TodoServiceApi();

try { 
    var result = api_instance.queryTodo();
    print(result);
} catch (e) {
    print("Exception when calling TodoServiceApi->queryTodo: $e\n");
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**TodoQueryRequest**](TodoQueryRequest.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **updateTodo**
> TodoMutateResponse updateTodo(id, body)



### Example 
```dart
import 'package:swagger/api.dart';

var api_instance = new TodoServiceApi();
var id = 56; // int | 
var body = new TodoUpdateRequest(); // TodoUpdateRequest | 

try { 
    var result = api_instance.updateTodo(id, body);
    print(result);
} catch (e) {
    print("Exception when calling TodoServiceApi->updateTodo: $e\n");
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  | 
 **body** | [**TodoUpdateRequest**](TodoUpdateRequest.md)|  | 

### Return type

[**TodoMutateResponse**](TodoMutateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

