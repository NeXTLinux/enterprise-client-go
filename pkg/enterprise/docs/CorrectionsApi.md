# \CorrectionsApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCorrection**](CorrectionsApi.md#AddCorrection) | **Post** /corrections | Create a correction record
[**DeleteCorrectionByUuid**](CorrectionsApi.md#DeleteCorrectionByUuid) | **Delete** /corrections/{uuid} | Delete a correction by UUID
[**GetCorrectionByUuid**](CorrectionsApi.md#GetCorrectionByUuid) | **Get** /corrections/{uuid} | Retrieve a correction by UUID
[**GetCorrections**](CorrectionsApi.md#GetCorrections) | **Get** /corrections | Retrieve a list of corrections
[**UpdateCorrectionByUuid**](CorrectionsApi.md#UpdateCorrectionByUuid) | **Put** /corrections/{uuid} | Update a correction by UUID



## AddCorrection

> Correction AddCorrection(ctx).Correction(correction).XNextlinuxAccount(xNextlinuxAccount).Execute()

Create a correction record



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
    correction := *openapiclient.NewCorrection("Type_example", *openapiclient.NewCorrectionMatch("npm"), []openapiclient.CorrectionFieldMatch{*openapiclient.NewCorrectionFieldMatch("name", "FieldValue_example")}) // Correction | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CorrectionsApi.AddCorrection(context.Background()).Correction(correction).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrectionsApi.AddCorrection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCorrection`: Correction
    fmt.Fprintf(os.Stdout, "Response from `CorrectionsApi.AddCorrection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCorrectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correction** | [**Correction**](Correction.md) |  | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCorrectionByUuid

> DeleteCorrectionByUuid(ctx, uuid).XNextlinuxAccount(xNextlinuxAccount).Execute()

Delete a correction by UUID



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
    uuid := "uuid_example" // string | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CorrectionsApi.DeleteCorrectionByUuid(context.Background(), uuid).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrectionsApi.DeleteCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrectionByUuid

> Correction GetCorrectionByUuid(ctx, uuid).XNextlinuxAccount(xNextlinuxAccount).Execute()

Retrieve a correction by UUID



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
    uuid := "uuid_example" // string | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CorrectionsApi.GetCorrectionByUuid(context.Background(), uuid).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrectionsApi.GetCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrectionByUuid`: Correction
    fmt.Fprintf(os.Stdout, "Response from `CorrectionsApi.GetCorrectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCorrections

> []Correction GetCorrections(ctx).CorrectionType(correctionType).XNextlinuxAccount(xNextlinuxAccount).Execute()

Retrieve a list of corrections



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
    correctionType := "correctionType_example" // string |  (optional)
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CorrectionsApi.GetCorrections(context.Background()).CorrectionType(correctionType).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrectionsApi.GetCorrections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCorrections`: []Correction
    fmt.Fprintf(os.Stdout, "Response from `CorrectionsApi.GetCorrections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCorrectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correctionType** | **string** |  | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**[]Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCorrectionByUuid

> Correction UpdateCorrectionByUuid(ctx, uuid).Correction(correction).XNextlinuxAccount(xNextlinuxAccount).Execute()

Update a correction by UUID



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
    uuid := "uuid_example" // string | 
    correction := *openapiclient.NewCorrection("Type_example", *openapiclient.NewCorrectionMatch("npm"), []openapiclient.CorrectionFieldMatch{*openapiclient.NewCorrectionFieldMatch("name", "FieldValue_example")}) // Correction | 
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CorrectionsApi.UpdateCorrectionByUuid(context.Background(), uuid).Correction(correction).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CorrectionsApi.UpdateCorrectionByUuid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCorrectionByUuid`: Correction
    fmt.Fprintf(os.Stdout, "Response from `CorrectionsApi.UpdateCorrectionByUuid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCorrectionByUuidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correction** | [**Correction**](Correction.md) |  | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**Correction**](Correction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

