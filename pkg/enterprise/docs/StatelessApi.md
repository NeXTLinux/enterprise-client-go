# \StatelessApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatelessSbomVulnerabilities**](StatelessApi.md#GetStatelessSbomVulnerabilities) | **Post** /stateless/sbom/vuln/{vtype} | Get vulnerabilities for input sbom by type



## GetStatelessSbomVulnerabilities

> SBOMVulnerabilitiesResponse GetStatelessSbomVulnerabilities(ctx, vtype).Sbom(sbom).WillNotFix(willNotFix).XNextlinuxAccount(xNextlinuxAccount).Execute()

Get vulnerabilities for input sbom by type

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
    vtype := "vtype_example" // string | 
    sbom := interface{}(Object) // interface{} | 
    willNotFix := true // bool | Vulnerability data publishers explicitly won't fix some vulnerabilities. This is captured by will_not_fix attribute of each result. If the query parameter is set, results matching it's value will be filtered. Results are not filtered if the query parameter is unset (optional)
    xNextlinuxAccount := "xNextlinuxAccount_example" // string | An account name to change the resource scope of the request to that account, if permissions allow (admin only) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StatelessApi.GetStatelessSbomVulnerabilities(context.Background(), vtype).Sbom(sbom).WillNotFix(willNotFix).XNextlinuxAccount(xNextlinuxAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatelessApi.GetStatelessSbomVulnerabilities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatelessSbomVulnerabilities`: SBOMVulnerabilitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `StatelessApi.GetStatelessSbomVulnerabilities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vtype** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatelessSbomVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sbom** | **interface{}** |  | 
 **willNotFix** | **bool** | Vulnerability data publishers explicitly won&#39;t fix some vulnerabilities. This is captured by will_not_fix attribute of each result. If the query parameter is set, results matching it&#39;s value will be filtered. Results are not filtered if the query parameter is unset | 
 **xNextlinuxAccount** | **string** | An account name to change the resource scope of the request to that account, if permissions allow (admin only) | 

### Return type

[**SBOMVulnerabilitiesResponse**](SBOMVulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

