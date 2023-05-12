/*
Nextlinux Engine API Server

This is the Nextlinux Engine API. Provides the primary external API for users of the service.

API version: 0.6.0
Contact: nurmi@nextlinux.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package engine

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"os"
)

// Linger please
var (
	_ _context.Context
)

type ImportApi interface {

	/*
	ImportImageArchive Import an nextlinux image tar.gz archive file. This is a deprecated API replaced by the \"/imports/images\" route

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiImportImageArchiveRequest
	*/
	ImportImageArchive(ctx _context.Context) ApiImportImageArchiveRequest

	// ImportImageArchiveExecute executes the request
	//  @return []NextlinuxImage
	ImportImageArchiveExecute(r ApiImportImageArchiveRequest) ([]NextlinuxImage, *_nethttp.Response, error)
}

// ImportApiService ImportApi service
type ImportApiService service

type ApiImportImageArchiveRequest struct {
	ctx _context.Context
	ApiService ImportApi
	archiveFile **os.File
}

// nextlinux image tar archive.
func (r ApiImportImageArchiveRequest) ArchiveFile(archiveFile *os.File) ApiImportImageArchiveRequest {
	r.archiveFile = &archiveFile
	return r
}

func (r ApiImportImageArchiveRequest) Execute() ([]NextlinuxImage, *_nethttp.Response, error) {
	return r.ApiService.ImportImageArchiveExecute(r)
}

/*
ImportImageArchive Import an nextlinux image tar.gz archive file. This is a deprecated API replaced by the \"/imports/images\" route

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiImportImageArchiveRequest
*/
func (a *ImportApiService) ImportImageArchive(ctx _context.Context) ApiImportImageArchiveRequest {
	return ApiImportImageArchiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []NextlinuxImage
func (a *ImportApiService) ImportImageArchiveExecute(r ApiImportImageArchiveRequest) ([]NextlinuxImage, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []NextlinuxImage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportApiService.ImportImageArchive")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/import/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.archiveFile == nil {
		return localVarReturnValue, nil, reportError("archiveFile is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	localVarFormFileName = "archive_file"
	localVarFile := *r.archiveFile
	if localVarFile != nil {
		fbs, _ := _ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
