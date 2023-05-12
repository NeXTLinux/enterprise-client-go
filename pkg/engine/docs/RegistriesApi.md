# \RegistriesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistry**](RegistriesApi.md#CreateRegistry) | **Post** /registries | Add a new registry
[**DeleteRegistry**](RegistriesApi.md#DeleteRegistry) | **Delete** /registries/{registry} | Delete a registry configuration
[**GetRegistry**](RegistriesApi.md#GetRegistry) | **Get** /registries/{registry} | Get a specific registry configuration
[**ListRegistries**](RegistriesApi.md#ListRegistries) | **Get** /registries | List configured registries
[**UpdateRegistry**](RegistriesApi.md#UpdateRegistry) | **Put** /registries/{registry} | Update/replace a registry configuration



## CreateRegistry

> []RegistryConfiguration CreateRegistry(ctx).Registrydata(registrydata).Validate(validate).XNextlinuxAccount(xNextlinuxAccount).Execute()

Add a new registry



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registrydata := *openapiclient.NewRegistryConfigurationRequest() // RegistryConfigurationRequest | 
    validate := true // bool | flag to determine whether or not to validate registry/credential at registry add time (optional)
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.CreateRegistry(context.Background()).Registrydata(registrydata).Validate(validate).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.CreateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.CreateRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registrydata** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md) |  | 
 **validate** | **bool** | flag to determine whether or not to validate registry/credential at registry add time | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRegistry

> DeleteRegistry(ctx, registry).XNextlinuxAccount(xNextlinuxAccount).Execute()

Delete a registry configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registry := "registry_example" // string | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.DeleteRegistry(context.Background(), registry).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.DeleteRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegistry

> []RegistryConfiguration GetRegistry(ctx, registry).XNextlinuxAccount(xNextlinuxAccount).Execute()

Get a specific registry configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registry := "registry_example" // string | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.GetRegistry(context.Background(), registry).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.GetRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.GetRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRegistries

> []RegistryConfiguration ListRegistries(ctx).XNextlinuxAccount(xNextlinuxAccount).Execute()

List configured registries



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.ListRegistries(context.Background()).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.ListRegistries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRegistries`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.ListRegistries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRegistriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegistry

> []RegistryConfiguration UpdateRegistry(ctx, registry).Registrydata(registrydata).Validate(validate).XNextlinuxAccount(xNextlinuxAccount).Execute()

Update/replace a registry configuration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    registry := "registry_example" // string | 
    registrydata := *openapiclient.NewRegistryConfigurationRequest() // RegistryConfigurationRequest | 
    validate := true // bool | flag to determine whether or not to validate registry/credential at registry update time (optional)
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RegistriesApi.UpdateRegistry(context.Background(), registry).Registrydata(registrydata).Validate(validate).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegistriesApi.UpdateRegistry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRegistry`: []RegistryConfiguration
    fmt.Fprintf(os.Stdout, "Response from `RegistriesApi.UpdateRegistry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registry** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registrydata** | [**RegistryConfigurationRequest**](RegistryConfigurationRequest.md) |  | 
 **validate** | **bool** | flag to determine whether or not to validate registry/credential at registry update time | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]RegistryConfiguration**](RegistryConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

