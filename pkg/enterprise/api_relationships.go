/*
Nextlinux Enterprise API Server

This is the Nextlinux Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.7.0
Contact: dev@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package enterprise

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type RelationshipsApi interface {

	/*
	AddArtifactRelationship Method for AddArtifactRelationship

	Add a new relationship for this image to another artifact (source or image)

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiAddArtifactRelationshipRequest
	*/
	AddArtifactRelationship(ctx _context.Context) ApiAddArtifactRelationshipRequest

	// AddArtifactRelationshipExecute executes the request
	//  @return interface{}
	AddArtifactRelationshipExecute(r ApiAddArtifactRelationshipRequest) (interface{}, *_nethttp.Response, error)

	/*
	DeleteArtifactRelationships Method for DeleteArtifactRelationships

	Delete one or more relationships

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiDeleteArtifactRelationshipsRequest
	*/
	DeleteArtifactRelationships(ctx _context.Context) ApiDeleteArtifactRelationshipsRequest

	// DeleteArtifactRelationshipsExecute executes the request
	//  @return interface{}
	DeleteArtifactRelationshipsExecute(r ApiDeleteArtifactRelationshipsRequest) (interface{}, *_nethttp.Response, error)

	/*
	GetArtifactRelationship Method for GetArtifactRelationship

	Get the relationship between software supply chain artifacts (images, source revisions, etc)

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param relationshipId Id of record to retrieve
	 @return ApiGetArtifactRelationshipRequest
	*/
	GetArtifactRelationship(ctx _context.Context, relationshipId string) ApiGetArtifactRelationshipRequest

	// GetArtifactRelationshipExecute executes the request
	//  @return ArtifactRelationship
	GetArtifactRelationshipExecute(r ApiGetArtifactRelationshipRequest) (ArtifactRelationship, *_nethttp.Response, error)

	/*
	GetRelationshipSbomDiff Method for GetRelationshipSbomDiff

	Return the context-aware diff of the sboms for the relationship

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param relationshipId
	 @return ApiGetRelationshipSbomDiffRequest
	*/
	GetRelationshipSbomDiff(ctx _context.Context, relationshipId string) ApiGetRelationshipSbomDiffRequest

	// GetRelationshipSbomDiffExecute executes the request
	//  @return RelationshipSbomDiff
	GetRelationshipSbomDiffExecute(r ApiGetRelationshipSbomDiffRequest) (RelationshipSbomDiff, *_nethttp.Response, error)

	/*
	ListArtifactRelationships Method for ListArtifactRelationships

	List the relationships between software supply chain artifacts (images, source revisions, etc)

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListArtifactRelationshipsRequest
	*/
	ListArtifactRelationships(ctx _context.Context) ApiListArtifactRelationshipsRequest

	// ListArtifactRelationshipsExecute executes the request
	//  @return []ArtifactRelationship
	ListArtifactRelationshipsExecute(r ApiListArtifactRelationshipsRequest) ([]ArtifactRelationship, *_nethttp.Response, error)
}

// RelationshipsApiService RelationshipsApi service
type RelationshipsApiService service

type ApiAddArtifactRelationshipRequest struct {
	ctx _context.Context
	ApiService RelationshipsApi
	relationship *ArtifactRelationship
}

func (r ApiAddArtifactRelationshipRequest) Relationship(relationship ArtifactRelationship) ApiAddArtifactRelationshipRequest {
	r.relationship = &relationship
	return r
}

func (r ApiAddArtifactRelationshipRequest) Execute() (interface{}, *_nethttp.Response, error) {
	return r.ApiService.AddArtifactRelationshipExecute(r)
}

/*
AddArtifactRelationship Method for AddArtifactRelationship

Add a new relationship for this image to another artifact (source or image)

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddArtifactRelationshipRequest
*/
func (a *RelationshipsApiService) AddArtifactRelationship(ctx _context.Context) ApiAddArtifactRelationshipRequest {
	return ApiAddArtifactRelationshipRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *RelationshipsApiService) AddArtifactRelationshipExecute(r ApiAddArtifactRelationshipRequest) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.AddArtifactRelationship")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact_relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.relationship == nil {
		return localVarReturnValue, nil, reportError("relationship is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.relationship
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteArtifactRelationshipsRequest struct {
	ctx _context.Context
	ApiService RelationshipsApi
	relationshipIds *[]string
}

// List of relationship Ids to delete
func (r ApiDeleteArtifactRelationshipsRequest) RelationshipIds(relationshipIds []string) ApiDeleteArtifactRelationshipsRequest {
	r.relationshipIds = &relationshipIds
	return r
}

func (r ApiDeleteArtifactRelationshipsRequest) Execute() (interface{}, *_nethttp.Response, error) {
	return r.ApiService.DeleteArtifactRelationshipsExecute(r)
}

/*
DeleteArtifactRelationships Method for DeleteArtifactRelationships

Delete one or more relationships

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteArtifactRelationshipsRequest
*/
func (a *RelationshipsApiService) DeleteArtifactRelationships(ctx _context.Context) ApiDeleteArtifactRelationshipsRequest {
	return ApiDeleteArtifactRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return interface{}
func (a *RelationshipsApiService) DeleteArtifactRelationshipsExecute(r ApiDeleteArtifactRelationshipsRequest) (interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.DeleteArtifactRelationships")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact_relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.relationshipIds == nil {
		return localVarReturnValue, nil, reportError("relationshipIds is required and must be specified")
	}

	localVarQueryParams.Add("relationship_ids", parameterToString(*r.relationshipIds, "csv"))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetArtifactRelationshipRequest struct {
	ctx _context.Context
	ApiService RelationshipsApi
	relationshipId string
}


func (r ApiGetArtifactRelationshipRequest) Execute() (ArtifactRelationship, *_nethttp.Response, error) {
	return r.ApiService.GetArtifactRelationshipExecute(r)
}

/*
GetArtifactRelationship Method for GetArtifactRelationship

Get the relationship between software supply chain artifacts (images, source revisions, etc)

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param relationshipId Id of record to retrieve
 @return ApiGetArtifactRelationshipRequest
*/
func (a *RelationshipsApiService) GetArtifactRelationship(ctx _context.Context, relationshipId string) ApiGetArtifactRelationshipRequest {
	return ApiGetArtifactRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		relationshipId: relationshipId,
	}
}

// Execute executes the request
//  @return ArtifactRelationship
func (a *RelationshipsApiService) GetArtifactRelationshipExecute(r ApiGetArtifactRelationshipRequest) (ArtifactRelationship, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ArtifactRelationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.GetArtifactRelationship")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact_relationships/{relationship_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"relationship_id"+"}", _neturl.PathEscape(parameterToString(r.relationshipId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetRelationshipSbomDiffRequest struct {
	ctx _context.Context
	ApiService RelationshipsApi
	relationshipId string
}


func (r ApiGetRelationshipSbomDiffRequest) Execute() (RelationshipSbomDiff, *_nethttp.Response, error) {
	return r.ApiService.GetRelationshipSbomDiffExecute(r)
}

/*
GetRelationshipSbomDiff Method for GetRelationshipSbomDiff

Return the context-aware diff of the sboms for the relationship

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param relationshipId
 @return ApiGetRelationshipSbomDiffRequest
*/
func (a *RelationshipsApiService) GetRelationshipSbomDiff(ctx _context.Context, relationshipId string) ApiGetRelationshipSbomDiffRequest {
	return ApiGetRelationshipSbomDiffRequest{
		ApiService: a,
		ctx: ctx,
		relationshipId: relationshipId,
	}
}

// Execute executes the request
//  @return RelationshipSbomDiff
func (a *RelationshipsApiService) GetRelationshipSbomDiffExecute(r ApiGetRelationshipSbomDiffRequest) (RelationshipSbomDiff, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  RelationshipSbomDiff
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.GetRelationshipSbomDiff")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact_relationships/{relationship_id}/diffs/sbom"
	localVarPath = strings.Replace(localVarPath, "{"+"relationship_id"+"}", _neturl.PathEscape(parameterToString(r.relationshipId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListArtifactRelationshipsRequest struct {
	ctx _context.Context
	ApiService RelationshipsApi
	artifactType *string
	artifactId *string
}

// Filter for artifact type as either source or target
func (r ApiListArtifactRelationshipsRequest) ArtifactType(artifactType string) ApiListArtifactRelationshipsRequest {
	r.artifactType = &artifactType
	return r
}
// Filter for artifact id as either source or target
func (r ApiListArtifactRelationshipsRequest) ArtifactId(artifactId string) ApiListArtifactRelationshipsRequest {
	r.artifactId = &artifactId
	return r
}

func (r ApiListArtifactRelationshipsRequest) Execute() ([]ArtifactRelationship, *_nethttp.Response, error) {
	return r.ApiService.ListArtifactRelationshipsExecute(r)
}

/*
ListArtifactRelationships Method for ListArtifactRelationships

List the relationships between software supply chain artifacts (images, source revisions, etc)

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListArtifactRelationshipsRequest
*/
func (a *RelationshipsApiService) ListArtifactRelationships(ctx _context.Context) ApiListArtifactRelationshipsRequest {
	return ApiListArtifactRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ArtifactRelationship
func (a *RelationshipsApiService) ListArtifactRelationshipsExecute(r ApiListArtifactRelationshipsRequest) ([]ArtifactRelationship, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []ArtifactRelationship
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RelationshipsApiService.ListArtifactRelationships")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artifact_relationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.artifactType != nil {
		localVarQueryParams.Add("artifact_type", parameterToString(*r.artifactType, ""))
	}
	if r.artifactId != nil {
		localVarQueryParams.Add("artifact_id", parameterToString(*r.artifactId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
