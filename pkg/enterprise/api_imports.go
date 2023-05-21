/*
Anchore Enterprise API Server

This is the Anchore Enterprise API. It provides additional external API routes and functionality for enterprise users.

API version: 0.2.1
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

type ImportsApi interface {

	/*
	CreateOperation Begin the import of a source code repository analyzed by Syft into the system

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiCreateOperationRequest
	*/
	CreateOperation(ctx _context.Context) ApiCreateOperationRequest

	// CreateOperationExecute executes the request
	//  @return SourceImportOperation
	CreateOperationExecute(r ApiCreateOperationRequest) (SourceImportOperation, *_nethttp.Response, error)

	/*
	FinalizeOperation Add source records to catalog db

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param operationId
	 @return ApiFinalizeOperationRequest
	*/
	FinalizeOperation(ctx _context.Context, operationId string) ApiFinalizeOperationRequest

	// FinalizeOperationExecute executes the request
	//  @return SourceManifest
	FinalizeOperationExecute(r ApiFinalizeOperationRequest) (SourceManifest, *_nethttp.Response, error)

	/*
	GetImportSourcesSbom list the packages of an imported source code repository

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param operationId
	 @return ApiGetImportSourcesSbomRequest
	*/
	GetImportSourcesSbom(ctx _context.Context, operationId string) ApiGetImportSourcesSbomRequest

	// GetImportSourcesSbomExecute executes the request
	//  @return SourceImportContentResponse
	GetImportSourcesSbomExecute(r ApiGetImportSourcesSbomRequest) (SourceImportContentResponse, *_nethttp.Response, error)

	/*
	GetOperation Get detail on a single import

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param operationId
	 @return ApiGetOperationRequest
	*/
	GetOperation(ctx _context.Context, operationId string) ApiGetOperationRequest

	// GetOperationExecute executes the request
	//  @return SourceImportOperation
	GetOperationExecute(r ApiGetOperationRequest) (SourceImportOperation, *_nethttp.Response, error)

	/*
	InvalidateOperation Invalidate operation ID so it can be garbage collected

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param operationId
	 @return ApiInvalidateOperationRequest
	*/
	InvalidateOperation(ctx _context.Context, operationId string) ApiInvalidateOperationRequest

	// InvalidateOperationExecute executes the request
	//  @return SourceImportOperation
	InvalidateOperationExecute(r ApiInvalidateOperationRequest) (SourceImportOperation, *_nethttp.Response, error)

	/*
	ListOperations Lists in-progress imports

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListOperationsRequest
	*/
	ListOperations(ctx _context.Context) ApiListOperationsRequest

	// ListOperationsExecute executes the request
	//  @return []SourceImportOperation
	ListOperationsExecute(r ApiListOperationsRequest) ([]SourceImportOperation, *_nethttp.Response, error)

	/*
	UploadImportSourcesSbom Begin the import of a source code repository analyzed by Syft into the system

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param operationId
	 @return ApiUploadImportSourcesSbomRequest
	*/
	UploadImportSourcesSbom(ctx _context.Context, operationId string) ApiUploadImportSourcesSbomRequest

	// UploadImportSourcesSbomExecute executes the request
	//  @return SourceImportContentResponse
	UploadImportSourcesSbomExecute(r ApiUploadImportSourcesSbomRequest) (SourceImportContentResponse, *_nethttp.Response, error)
}

// ImportsApiService ImportsApi service
type ImportsApiService service

type ApiCreateOperationRequest struct {
	ctx _context.Context
	ApiService ImportsApi
}


func (r ApiCreateOperationRequest) Execute() (SourceImportOperation, *_nethttp.Response, error) {
	return r.ApiService.CreateOperationExecute(r)
}

/*
CreateOperation Begin the import of a source code repository analyzed by Syft into the system

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOperationRequest
*/
func (a *ImportsApiService) CreateOperation(ctx _context.Context) ApiCreateOperationRequest {
	return ApiCreateOperationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SourceImportOperation
func (a *ImportsApiService) CreateOperationExecute(r ApiCreateOperationRequest) (SourceImportOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceImportOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.CreateOperation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources"

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

type ApiFinalizeOperationRequest struct {
	ctx _context.Context
	ApiService ImportsApi
	operationId string
	metadata *SourceImportMetadata
}

func (r ApiFinalizeOperationRequest) Metadata(metadata SourceImportMetadata) ApiFinalizeOperationRequest {
	r.metadata = &metadata
	return r
}

func (r ApiFinalizeOperationRequest) Execute() (SourceManifest, *_nethttp.Response, error) {
	return r.ApiService.FinalizeOperationExecute(r)
}

/*
FinalizeOperation Add source records to catalog db

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId
 @return ApiFinalizeOperationRequest
*/
func (a *ImportsApiService) FinalizeOperation(ctx _context.Context, operationId string) ApiFinalizeOperationRequest {
	return ApiFinalizeOperationRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
//  @return SourceManifest
func (a *ImportsApiService) FinalizeOperationExecute(r ApiFinalizeOperationRequest) (SourceManifest, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceManifest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.FinalizeOperation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources/{operation_id}/finalize"
	localVarPath = strings.Replace(localVarPath, "{"+"operation_id"+"}", _neturl.PathEscape(parameterToString(r.operationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.metadata == nil {
		return localVarReturnValue, nil, reportError("metadata is required and must be specified")
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
	localVarPostBody = r.metadata
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

type ApiGetImportSourcesSbomRequest struct {
	ctx _context.Context
	ApiService ImportsApi
	operationId string
}


func (r ApiGetImportSourcesSbomRequest) Execute() (SourceImportContentResponse, *_nethttp.Response, error) {
	return r.ApiService.GetImportSourcesSbomExecute(r)
}

/*
GetImportSourcesSbom list the packages of an imported source code repository

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId
 @return ApiGetImportSourcesSbomRequest
*/
func (a *ImportsApiService) GetImportSourcesSbom(ctx _context.Context, operationId string) ApiGetImportSourcesSbomRequest {
	return ApiGetImportSourcesSbomRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
//  @return SourceImportContentResponse
func (a *ImportsApiService) GetImportSourcesSbomExecute(r ApiGetImportSourcesSbomRequest) (SourceImportContentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceImportContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.GetImportSourcesSbom")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources/{operation_id}/sbom"
	localVarPath = strings.Replace(localVarPath, "{"+"operation_id"+"}", _neturl.PathEscape(parameterToString(r.operationId, "")), -1)

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

type ApiGetOperationRequest struct {
	ctx _context.Context
	ApiService ImportsApi
	operationId string
}


func (r ApiGetOperationRequest) Execute() (SourceImportOperation, *_nethttp.Response, error) {
	return r.ApiService.GetOperationExecute(r)
}

/*
GetOperation Get detail on a single import

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId
 @return ApiGetOperationRequest
*/
func (a *ImportsApiService) GetOperation(ctx _context.Context, operationId string) ApiGetOperationRequest {
	return ApiGetOperationRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
//  @return SourceImportOperation
func (a *ImportsApiService) GetOperationExecute(r ApiGetOperationRequest) (SourceImportOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceImportOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.GetOperation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources/{operation_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"operation_id"+"}", _neturl.PathEscape(parameterToString(r.operationId, "")), -1)

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

type ApiInvalidateOperationRequest struct {
	ctx _context.Context
	ApiService ImportsApi
	operationId string
}


func (r ApiInvalidateOperationRequest) Execute() (SourceImportOperation, *_nethttp.Response, error) {
	return r.ApiService.InvalidateOperationExecute(r)
}

/*
InvalidateOperation Invalidate operation ID so it can be garbage collected

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId
 @return ApiInvalidateOperationRequest
*/
func (a *ImportsApiService) InvalidateOperation(ctx _context.Context, operationId string) ApiInvalidateOperationRequest {
	return ApiInvalidateOperationRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
//  @return SourceImportOperation
func (a *ImportsApiService) InvalidateOperationExecute(r ApiInvalidateOperationRequest) (SourceImportOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceImportOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.InvalidateOperation")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources/{operation_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"operation_id"+"}", _neturl.PathEscape(parameterToString(r.operationId, "")), -1)

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

type ApiListOperationsRequest struct {
	ctx _context.Context
	ApiService ImportsApi
}


func (r ApiListOperationsRequest) Execute() ([]SourceImportOperation, *_nethttp.Response, error) {
	return r.ApiService.ListOperationsExecute(r)
}

/*
ListOperations Lists in-progress imports

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOperationsRequest
*/
func (a *ImportsApiService) ListOperations(ctx _context.Context) ApiListOperationsRequest {
	return ApiListOperationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []SourceImportOperation
func (a *ImportsApiService) ListOperationsExecute(r ApiListOperationsRequest) ([]SourceImportOperation, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []SourceImportOperation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.ListOperations")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources"

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

type ApiUploadImportSourcesSbomRequest struct {
	ctx _context.Context
	ApiService ImportsApi
	operationId string
	sbom *NativeSBOM
}

func (r ApiUploadImportSourcesSbomRequest) Sbom(sbom NativeSBOM) ApiUploadImportSourcesSbomRequest {
	r.sbom = &sbom
	return r
}

func (r ApiUploadImportSourcesSbomRequest) Execute() (SourceImportContentResponse, *_nethttp.Response, error) {
	return r.ApiService.UploadImportSourcesSbomExecute(r)
}

/*
UploadImportSourcesSbom Begin the import of a source code repository analyzed by Syft into the system

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param operationId
 @return ApiUploadImportSourcesSbomRequest
*/
func (a *ImportsApiService) UploadImportSourcesSbom(ctx _context.Context, operationId string) ApiUploadImportSourcesSbomRequest {
	return ApiUploadImportSourcesSbomRequest{
		ApiService: a,
		ctx: ctx,
		operationId: operationId,
	}
}

// Execute executes the request
//  @return SourceImportContentResponse
func (a *ImportsApiService) UploadImportSourcesSbomExecute(r ApiUploadImportSourcesSbomRequest) (SourceImportContentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SourceImportContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.UploadImportSourcesSbom")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/imports/sources/{operation_id}/sbom"
	localVarPath = strings.Replace(localVarPath, "{"+"operation_id"+"}", _neturl.PathEscape(parameterToString(r.operationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.sbom == nil {
		return localVarReturnValue, nil, reportError("sbom is required and must be specified")
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
	localVarPostBody = r.sbom
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
