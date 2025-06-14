/*
FOSSA API

OpenAPI Specification for public FOSSA APIs

API version: 4.28.61
Contact: support@fossa.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fossa

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// PackageLabelsAPIService PackageLabelsAPI service
type PackageLabelsAPIService service

type ApiCreatePackageLabelRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
	createPackageLabelRequest *CreatePackageLabelRequest
}

func (r ApiCreatePackageLabelRequest) CreatePackageLabelRequest(createPackageLabelRequest CreatePackageLabelRequest) ApiCreatePackageLabelRequest {
	r.createPackageLabelRequest = &createPackageLabelRequest
	return r
}

func (r ApiCreatePackageLabelRequest) Execute() (*GetPackageLabels200Response, *http.Response, error) {
	return r.ApiService.CreatePackageLabelExecute(r)
}

/*
CreatePackageLabel Method for CreatePackageLabel

Create one or more new Package Labels. Note - This endpoint is not available for organizations on the Free plan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePackageLabelRequest
*/
func (a *PackageLabelsAPIService) CreatePackageLabel(ctx context.Context) ApiCreatePackageLabelRequest {
	return ApiCreatePackageLabelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageLabels200Response
func (a *PackageLabelsAPIService) CreatePackageLabelExecute(r ApiCreatePackageLabelRequest) (*GetPackageLabels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageLabels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.CreatePackageLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-labels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createPackageLabelRequest == nil {
		return localVarReturnValue, nil, reportError("createPackageLabelRequest is required and must be specified")
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
	localVarPostBody = r.createPackageLabelRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreatePackageLabelAssignmentsRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
	createPackageLabelAssignmentsRequest *CreatePackageLabelAssignmentsRequest
}

func (r ApiCreatePackageLabelAssignmentsRequest) CreatePackageLabelAssignmentsRequest(createPackageLabelAssignmentsRequest CreatePackageLabelAssignmentsRequest) ApiCreatePackageLabelAssignmentsRequest {
	r.createPackageLabelAssignmentsRequest = &createPackageLabelAssignmentsRequest
	return r
}

func (r ApiCreatePackageLabelAssignmentsRequest) Execute() (*GetPackageLabelAssignments200Response, *http.Response, error) {
	return r.ApiService.CreatePackageLabelAssignmentsExecute(r)
}

/*
CreatePackageLabelAssignments Method for CreatePackageLabelAssignments

Assign one or more Package Labels to a single package. Note - This endpoint is not available for organizations on the Free plan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePackageLabelAssignmentsRequest
*/
func (a *PackageLabelsAPIService) CreatePackageLabelAssignments(ctx context.Context) ApiCreatePackageLabelAssignmentsRequest {
	return ApiCreatePackageLabelAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageLabelAssignments200Response
func (a *PackageLabelsAPIService) CreatePackageLabelAssignmentsExecute(r ApiCreatePackageLabelAssignmentsRequest) (*GetPackageLabelAssignments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageLabelAssignments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.CreatePackageLabelAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-label-assignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createPackageLabelAssignmentsRequest == nil {
		return localVarReturnValue, nil, reportError("createPackageLabelAssignmentsRequest is required and must be specified")
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
	localVarPostBody = r.createPackageLabelAssignmentsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeletePackageLabelAssignmentsRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
	deletePackageLabelAssignmentsRequest *DeletePackageLabelAssignmentsRequest
}

func (r ApiDeletePackageLabelAssignmentsRequest) DeletePackageLabelAssignmentsRequest(deletePackageLabelAssignmentsRequest DeletePackageLabelAssignmentsRequest) ApiDeletePackageLabelAssignmentsRequest {
	r.deletePackageLabelAssignmentsRequest = &deletePackageLabelAssignmentsRequest
	return r
}

func (r ApiDeletePackageLabelAssignmentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePackageLabelAssignmentsExecute(r)
}

/*
DeletePackageLabelAssignments Method for DeletePackageLabelAssignments

Remove one or more Package Labels from a single package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeletePackageLabelAssignmentsRequest
*/
func (a *PackageLabelsAPIService) DeletePackageLabelAssignments(ctx context.Context) ApiDeletePackageLabelAssignmentsRequest {
	return ApiDeletePackageLabelAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PackageLabelsAPIService) DeletePackageLabelAssignmentsExecute(r ApiDeletePackageLabelAssignmentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.DeletePackageLabelAssignments")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-label-assignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deletePackageLabelAssignmentsRequest == nil {
		return nil, reportError("deletePackageLabelAssignmentsRequest is required and must be specified")
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
	localVarPostBody = r.deletePackageLabelAssignmentsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeletePackageLabelsRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
	deletePackageLabelsRequest *DeletePackageLabelsRequest
}

func (r ApiDeletePackageLabelsRequest) DeletePackageLabelsRequest(deletePackageLabelsRequest DeletePackageLabelsRequest) ApiDeletePackageLabelsRequest {
	r.deletePackageLabelsRequest = &deletePackageLabelsRequest
	return r
}

func (r ApiDeletePackageLabelsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePackageLabelsExecute(r)
}

/*
DeletePackageLabels Method for DeletePackageLabels

Delete one or more Package Labels. Note - This endpoint is not available for organizations on the Free plan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeletePackageLabelsRequest
*/
func (a *PackageLabelsAPIService) DeletePackageLabels(ctx context.Context) ApiDeletePackageLabelsRequest {
	return ApiDeletePackageLabelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PackageLabelsAPIService) DeletePackageLabelsExecute(r ApiDeletePackageLabelsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.DeletePackageLabels")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-labels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deletePackageLabelsRequest == nil {
		return nil, reportError("deletePackageLabelsRequest is required and must be specified")
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
	localVarPostBody = r.deletePackageLabelsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPackageLabelAssignmentsRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
	packageId *string
	packageVersion *string
	scope *string
	scopeId *string
}

func (r ApiGetPackageLabelAssignmentsRequest) PackageId(packageId string) ApiGetPackageLabelAssignmentsRequest {
	r.packageId = &packageId
	return r
}

func (r ApiGetPackageLabelAssignmentsRequest) PackageVersion(packageVersion string) ApiGetPackageLabelAssignmentsRequest {
	r.packageVersion = &packageVersion
	return r
}

func (r ApiGetPackageLabelAssignmentsRequest) Scope(scope string) ApiGetPackageLabelAssignmentsRequest {
	r.scope = &scope
	return r
}

func (r ApiGetPackageLabelAssignmentsRequest) ScopeId(scopeId string) ApiGetPackageLabelAssignmentsRequest {
	r.scopeId = &scopeId
	return r
}

func (r ApiGetPackageLabelAssignmentsRequest) Execute() (*GetPackageLabelAssignments200Response, *http.Response, error) {
	return r.ApiService.GetPackageLabelAssignmentsExecute(r)
}

/*
GetPackageLabelAssignments Method for GetPackageLabelAssignments

Get all Package Labels assigned to a single package.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackageLabelAssignmentsRequest
*/
func (a *PackageLabelsAPIService) GetPackageLabelAssignments(ctx context.Context) ApiGetPackageLabelAssignmentsRequest {
	return ApiGetPackageLabelAssignmentsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageLabelAssignments200Response
func (a *PackageLabelsAPIService) GetPackageLabelAssignmentsExecute(r ApiGetPackageLabelAssignmentsRequest) (*GetPackageLabelAssignments200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageLabelAssignments200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.GetPackageLabelAssignments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-label-assignments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.packageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageId", r.packageId, "form", "")
	}
	if r.packageVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageVersion", r.packageVersion, "form", "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
	}
	if r.scopeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scopeId", r.scopeId, "form", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPackageLabelsRequest struct {
	ctx context.Context
	ApiService *PackageLabelsAPIService
}

func (r ApiGetPackageLabelsRequest) Execute() (*GetPackageLabels200Response, *http.Response, error) {
	return r.ApiService.GetPackageLabelsExecute(r)
}

/*
GetPackageLabels Method for GetPackageLabels

Retrieve all of the Package Labels defined for the current organization. Note - This endpoint is not available for organizations on the Free plan.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackageLabelsRequest
*/
func (a *PackageLabelsAPIService) GetPackageLabels(ctx context.Context) ApiGetPackageLabelsRequest {
	return ApiGetPackageLabelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageLabels200Response
func (a *PackageLabelsAPIService) GetPackageLabelsExecute(r ApiGetPackageLabelsRequest) (*GetPackageLabels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageLabels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageLabelsAPIService.GetPackageLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/package-labels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v GetGitHubAppInstallationUrl403Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
