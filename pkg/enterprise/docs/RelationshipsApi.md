# \RelationshipsApi

All URIs are relative to *http://localhost/enterprise*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArtifactRelationship**](RelationshipsApi.md#AddArtifactRelationship) | **Post** /artifact_relationships | 
[**DeleteArtifactRelationships**](RelationshipsApi.md#DeleteArtifactRelationships) | **Delete** /artifact_relationships | 
[**GetArtifactRelationship**](RelationshipsApi.md#GetArtifactRelationship) | **Get** /artifact_relationships/{relationship_id} | 
[**GetRelationshipSbomDiff**](RelationshipsApi.md#GetRelationshipSbomDiff) | **Get** /artifact_relationships/{relationship_id}/diffs/sbom | 
[**ListArtifactRelationships**](RelationshipsApi.md#ListArtifactRelationships) | **Get** /artifact_relationships | 



## AddArtifactRelationship

> interface{} AddArtifactRelationship(ctx).Relationship(relationship).Execute()





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
    relationship := *openapiclient.NewArtifactRelationship() // ArtifactRelationship | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationshipsApi.AddArtifactRelationship(context.Background()).Relationship(relationship).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.AddArtifactRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArtifactRelationship`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.AddArtifactRelationship`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddArtifactRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relationship** | [**ArtifactRelationship**](ArtifactRelationship.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactRelationships

> interface{} DeleteArtifactRelationships(ctx).RelationshipIds(relationshipIds).Execute()





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
    relationshipIds := []string{"Inner_example"} // []string | List of relationship Ids to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationshipsApi.DeleteArtifactRelationships(context.Background()).RelationshipIds(relationshipIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.DeleteArtifactRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteArtifactRelationships`: interface{}
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.DeleteArtifactRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relationshipIds** | **[]string** | List of relationship Ids to delete | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactRelationship

> ArtifactRelationship GetArtifactRelationship(ctx, relationshipId).Execute()





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
    relationshipId := "relationshipId_example" // string | Id of record to retrieve

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationshipsApi.GetArtifactRelationship(context.Background(), relationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.GetArtifactRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactRelationship`: ArtifactRelationship
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.GetArtifactRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipId** | **string** | Id of record to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArtifactRelationship**](ArtifactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelationshipSbomDiff

> RelationshipSbomDiff GetRelationshipSbomDiff(ctx, relationshipId).Execute()





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
    relationshipId := "relationshipId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationshipsApi.GetRelationshipSbomDiff(context.Background(), relationshipId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.GetRelationshipSbomDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationshipSbomDiff`: RelationshipSbomDiff
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.GetRelationshipSbomDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**relationshipId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationshipSbomDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RelationshipSbomDiff**](RelationshipSbomDiff.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArtifactRelationships

> []ArtifactRelationship ListArtifactRelationships(ctx).ArtifactType(artifactType).ArtifactId(artifactId).Execute()





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
    artifactType := "artifactType_example" // string | Filter for artifact type as either source or target (optional)
    artifactId := "artifactId_example" // string | Filter for artifact id as either source or target (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RelationshipsApi.ListArtifactRelationships(context.Background()).ArtifactType(artifactType).ArtifactId(artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.ListArtifactRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArtifactRelationships`: []ArtifactRelationship
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.ListArtifactRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListArtifactRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifactType** | **string** | Filter for artifact type as either source or target | 
 **artifactId** | **string** | Filter for artifact id as either source or target | 

### Return type

[**[]ArtifactRelationship**](ArtifactRelationship.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

