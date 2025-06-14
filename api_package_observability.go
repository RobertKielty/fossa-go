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
	"reflect"
)


// PackageObservabilityAPIService PackageObservabilityAPI service
type PackageObservabilityAPIService service

type ApiGetPackageIndexExportRequest struct {
	ctx context.Context
	ApiService *PackageObservabilityAPIService
	fetchers *[]string
	packageName *string
	depth *[]string
	labels *[]string
	projectName *string
	sources *[]string
	visibility *[]string
	blockType *string
	cve *string
	cwes *[]string
	locators *[]string
	fixTypes *[]string
	severities *[]string
	teamIds *[]float32
}

// Filter packages to those from the specified fetchers.  For example, fetchers[0]&#x3D;npm&amp;fetchers[1]&#x3D;apk will match all NPM packages and all APK packages.
func (r ApiGetPackageIndexExportRequest) Fetchers(fetchers []string) ApiGetPackageIndexExportRequest {
	r.fetchers = &fetchers
	return r
}

// Filter results to only packages with the specified name.  Supports partial matches.  For example \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;foo-bar\&quot;.
func (r ApiGetPackageIndexExportRequest) PackageName(packageName string) ApiGetPackageIndexExportRequest {
	r.packageName = &packageName
	return r
}

// Filter results to only include packages which are direct or transitive dependencies of your projects.
func (r ApiGetPackageIndexExportRequest) Depth(depth []string) ApiGetPackageIndexExportRequest {
	r.depth = &depth
	return r
}

// Filter packages to those belonging to your projects with the specified labels.
func (r ApiGetPackageIndexExportRequest) Labels(labels []string) ApiGetPackageIndexExportRequest {
	r.labels = &labels
	return r
}

// Filter packages to only one of your specific projects.  Exact match only.
func (r ApiGetPackageIndexExportRequest) ProjectName(projectName string) ApiGetPackageIndexExportRequest {
	r.projectName = &projectName
	return r
}

// Filter packages to those belonging to your projects from the specified set of sources
func (r ApiGetPackageIndexExportRequest) Sources(sources []string) ApiGetPackageIndexExportRequest {
	r.sources = &sources
	return r
}

// Filter results to your projects which are public or private
func (r ApiGetPackageIndexExportRequest) Visibility(visibility []string) ApiGetPackageIndexExportRequest {
	r.visibility = &visibility
	return r
}

// Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization
func (r ApiGetPackageIndexExportRequest) BlockType(blockType string) ApiGetPackageIndexExportRequest {
	r.blockType = &blockType
	return r
}

// Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers
func (r ApiGetPackageIndexExportRequest) Cve(cve string) ApiGetPackageIndexExportRequest {
	r.cve = &cve
	return r
}

// Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers
func (r ApiGetPackageIndexExportRequest) Cwes(cwes []string) ApiGetPackageIndexExportRequest {
	r.cwes = &cwes
	return r
}

// Filter packages to those with specific package locators (exact match only)
func (r ApiGetPackageIndexExportRequest) Locators(locators []string) ApiGetPackageIndexExportRequest {
	r.locators = &locators
	return r
}

// Filter packages to those with vulnerabilities that either have or do not have a fix available
func (r ApiGetPackageIndexExportRequest) FixTypes(fixTypes []string) ApiGetPackageIndexExportRequest {
	r.fixTypes = &fixTypes
	return r
}

// Filter packages by severity levels of issues
func (r ApiGetPackageIndexExportRequest) Severities(severities []string) ApiGetPackageIndexExportRequest {
	r.severities = &severities
	return r
}

// Filter packages to just those owned by the specified teams.  Specify the string \&quot;null\&quot; to filter packages that are not owned by any team.
func (r ApiGetPackageIndexExportRequest) TeamIds(teamIds []float32) ApiGetPackageIndexExportRequest {
	r.teamIds = &teamIds
	return r
}

func (r ApiGetPackageIndexExportRequest) Execute() (*GetPackageIndexExport201Response, *http.Response, error) {
	return r.ApiService.GetPackageIndexExportExecute(r)
}

/*
GetPackageIndexExport Method for GetPackageIndexExport

Request an export of the package index with the applied set of filters.  The export will be generated asynchronously and a link to the export will be emailed to you.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackageIndexExportRequest
*/
func (a *PackageObservabilityAPIService) GetPackageIndexExport(ctx context.Context) ApiGetPackageIndexExportRequest {
	return ApiGetPackageIndexExportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageIndexExport201Response
func (a *PackageObservabilityAPIService) GetPackageIndexExportExecute(r ApiGetPackageIndexExportRequest) (*GetPackageIndexExport201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageIndexExport201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageObservabilityAPIService.GetPackageIndexExport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/report"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fetchers != nil {
		t := *r.fetchers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fetchers", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fetchers", t, "form", "multi")
		}
	}
	if r.packageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageName", r.packageName, "form", "")
	}
	if r.depth != nil {
		t := *r.depth
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "depth", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "depth", t, "form", "multi")
		}
	}
	if r.labels != nil {
		t := *r.labels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "labels", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "labels", t, "form", "multi")
		}
	}
	if r.projectName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectName", r.projectName, "form", "")
	}
	if r.sources != nil {
		t := *r.sources
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sources", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sources", t, "form", "multi")
		}
	}
	if r.visibility != nil {
		t := *r.visibility
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", t, "form", "multi")
		}
	}
	if r.blockType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "blockType", r.blockType, "form", "")
	}
	if r.cve != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cve", r.cve, "form", "")
	}
	if r.cwes != nil {
		t := *r.cwes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "cwes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "cwes", t, "form", "multi")
		}
	}
	if r.locators != nil {
		t := *r.locators
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "locators", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "locators", t, "form", "multi")
		}
	}
	if r.fixTypes != nil {
		t := *r.fixTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fixTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fixTypes", t, "form", "multi")
		}
	}
	if r.severities != nil {
		t := *r.severities
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "severities", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "severities", t, "form", "multi")
		}
	}
	if r.teamIds != nil {
		t := *r.teamIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "teamIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "teamIds", t, "form", "multi")
		}
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

type ApiGetPackagesRequest struct {
	ctx context.Context
	ApiService *PackageObservabilityAPIService
	fetchers *[]string
	packageName *string
	depth *[]string
	labels *[]string
	projectName *string
	sources *[]string
	visibility *[]string
	blockType *string
	cve *string
	cwes *[]string
	fixTypes *[]string
	severities *[]string
	teamIds *[]float32
	locators *[]string
	page *int32
	count *int32
	sort *string
}

// Filter packages to those from the specified fetchers.  For example, fetchers[0]&#x3D;npm&amp;fetchers[1]&#x3D;apk will match all NPM packages and all APK packages.
func (r ApiGetPackagesRequest) Fetchers(fetchers []string) ApiGetPackagesRequest {
	r.fetchers = &fetchers
	return r
}

// Filter results to only packages with the specified name.  Supports partial matches.  For example \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;foo-bar\&quot;.
func (r ApiGetPackagesRequest) PackageName(packageName string) ApiGetPackagesRequest {
	r.packageName = &packageName
	return r
}

// Filter results to only include packages which are direct or transitive dependencies of your projects.
func (r ApiGetPackagesRequest) Depth(depth []string) ApiGetPackagesRequest {
	r.depth = &depth
	return r
}

// Filter packages to those belonging to your projects with the specified labels.
func (r ApiGetPackagesRequest) Labels(labels []string) ApiGetPackagesRequest {
	r.labels = &labels
	return r
}

// Filter packages to only one of your specific projects.  Exact match only.
func (r ApiGetPackagesRequest) ProjectName(projectName string) ApiGetPackagesRequest {
	r.projectName = &projectName
	return r
}

// Filter packages to those belonging to your projects from the specified set of sources
func (r ApiGetPackagesRequest) Sources(sources []string) ApiGetPackagesRequest {
	r.sources = &sources
	return r
}

// Filter results to your projects which are public or private
func (r ApiGetPackagesRequest) Visibility(visibility []string) ApiGetPackagesRequest {
	r.visibility = &visibility
	return r
}

// Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization
func (r ApiGetPackagesRequest) BlockType(blockType string) ApiGetPackagesRequest {
	r.blockType = &blockType
	return r
}

// Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers
func (r ApiGetPackagesRequest) Cve(cve string) ApiGetPackagesRequest {
	r.cve = &cve
	return r
}

// Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers
func (r ApiGetPackagesRequest) Cwes(cwes []string) ApiGetPackagesRequest {
	r.cwes = &cwes
	return r
}

// Filter packages to those with vulnerabilities that either have or do not have a fix available
func (r ApiGetPackagesRequest) FixTypes(fixTypes []string) ApiGetPackagesRequest {
	r.fixTypes = &fixTypes
	return r
}

// Filter packages by severity levels of issues
func (r ApiGetPackagesRequest) Severities(severities []string) ApiGetPackagesRequest {
	r.severities = &severities
	return r
}

// Filter packages to just those owned by the specified teams.  Specify the string \&quot;null\&quot; to filter packages that are not owned by any team.
func (r ApiGetPackagesRequest) TeamIds(teamIds []float32) ApiGetPackagesRequest {
	r.teamIds = &teamIds
	return r
}

// Filter packages by the specified dependency project locators without revision / version data.  Locators are unique identifiers for packages in the FOSSA system. Exact matches only.
func (r ApiGetPackagesRequest) Locators(locators []string) ApiGetPackagesRequest {
	r.locators = &locators
	return r
}

// The page number to retrieve
func (r ApiGetPackagesRequest) Page(page int32) ApiGetPackagesRequest {
	r.page = &page
	return r
}

// The number of results to return per page
func (r ApiGetPackagesRequest) Count(count int32) ApiGetPackagesRequest {
	r.count = &count
	return r
}

// The field to sort by
func (r ApiGetPackagesRequest) Sort(sort string) ApiGetPackagesRequest {
	r.sort = &sort
	return r
}

func (r ApiGetPackagesRequest) Execute() (*GetPackageIndexExport201Response, *http.Response, error) {
	return r.ApiService.GetPackagesExecute(r)
}

/*
GetPackages Method for GetPackages

Retrieve the number of projects that a dependency is used in and how many versions of it are blocked.  This endpoint is paginated.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesRequest
*/
func (a *PackageObservabilityAPIService) GetPackages(ctx context.Context) ApiGetPackagesRequest {
	return ApiGetPackagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackageIndexExport201Response
func (a *PackageObservabilityAPIService) GetPackagesExecute(r ApiGetPackagesRequest) (*GetPackageIndexExport201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackageIndexExport201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageObservabilityAPIService.GetPackages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fetchers != nil {
		t := *r.fetchers
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fetchers", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fetchers", t, "form", "multi")
		}
	}
	if r.packageName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageName", r.packageName, "form", "")
	}
	if r.depth != nil {
		t := *r.depth
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "depth", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "depth", t, "form", "multi")
		}
	}
	if r.labels != nil {
		t := *r.labels
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "labels", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "labels", t, "form", "multi")
		}
	}
	if r.projectName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "projectName", r.projectName, "form", "")
	}
	if r.sources != nil {
		t := *r.sources
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sources", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sources", t, "form", "multi")
		}
	}
	if r.visibility != nil {
		t := *r.visibility
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "visibility", t, "form", "multi")
		}
	}
	if r.blockType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "blockType", r.blockType, "form", "")
	}
	if r.cve != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cve", r.cve, "form", "")
	}
	if r.cwes != nil {
		t := *r.cwes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "cwes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "cwes", t, "form", "multi")
		}
	}
	if r.fixTypes != nil {
		t := *r.fixTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fixTypes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fixTypes", t, "form", "multi")
		}
	}
	if r.severities != nil {
		t := *r.severities
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "severities", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "severities", t, "form", "multi")
		}
	}
	if r.teamIds != nil {
		t := *r.teamIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "teamIds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "teamIds", t, "form", "multi")
		}
	}
	if r.locators != nil {
		t := *r.locators
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "locators", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "locators", t, "form", "multi")
		}
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		r.page = &defaultValue
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
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

type ApiGetPackagesPackageLocatorsRequest struct {
	ctx context.Context
	ApiService *PackageObservabilityAPIService
	packageLocator *string
	count *int32
}

// Filter locators to those partially matching the specified locator.  For example, \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;bazfoobar\&quot;.
func (r ApiGetPackagesPackageLocatorsRequest) PackageLocator(packageLocator string) ApiGetPackagesPackageLocatorsRequest {
	r.packageLocator = &packageLocator
	return r
}

// The number of results to return.
func (r ApiGetPackagesPackageLocatorsRequest) Count(count int32) ApiGetPackagesPackageLocatorsRequest {
	r.count = &count
	return r
}

func (r ApiGetPackagesPackageLocatorsRequest) Execute() (*GetPackagesPackageLocators200Response, *http.Response, error) {
	return r.ApiService.GetPackagesPackageLocatorsExecute(r)
}

/*
GetPackagesPackageLocators Method for GetPackagesPackageLocators

Fetch the list of packages locators used by your organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesPackageLocatorsRequest
*/
func (a *PackageObservabilityAPIService) GetPackagesPackageLocators(ctx context.Context) ApiGetPackagesPackageLocatorsRequest {
	return ApiGetPackagesPackageLocatorsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackagesPackageLocators200Response
func (a *PackageObservabilityAPIService) GetPackagesPackageLocatorsExecute(r ApiGetPackagesPackageLocatorsRequest) (*GetPackagesPackageLocators200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackagesPackageLocators200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageObservabilityAPIService.GetPackagesPackageLocators")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/package-locators"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.packageLocator != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageLocator", r.packageLocator, "form", "")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "form", "")
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

type ApiGetPackagesPackageManagersRequest struct {
	ctx context.Context
	ApiService *PackageObservabilityAPIService
}

func (r ApiGetPackagesPackageManagersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetPackagesPackageManagersExecute(r)
}

/*
GetPackagesPackageManagers Method for GetPackagesPackageManagers

Fetch the list of packages managers used by your organization. If none are used/can't be found, the full list of package managers is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesPackageManagersRequest
*/
func (a *PackageObservabilityAPIService) GetPackagesPackageManagers(ctx context.Context) ApiGetPackagesPackageManagersRequest {
	return ApiGetPackagesPackageManagersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *PackageObservabilityAPIService) GetPackagesPackageManagersExecute(r ApiGetPackagesPackageManagersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageObservabilityAPIService.GetPackagesPackageManagers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/package-managers"

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

type ApiGetPackagesPackageSummaryRequest struct {
	ctx context.Context
	ApiService *PackageObservabilityAPIService
}

func (r ApiGetPackagesPackageSummaryRequest) Execute() (*GetPackagesPackageSummary200Response, *http.Response, error) {
	return r.ApiService.GetPackagesPackageSummaryExecute(r)
}

/*
GetPackagesPackageSummary Method for GetPackagesPackageSummary

Fetch the total package count and the last cache date for your organization.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPackagesPackageSummaryRequest
*/
func (a *PackageObservabilityAPIService) GetPackagesPackageSummary(ctx context.Context) ApiGetPackagesPackageSummaryRequest {
	return ApiGetPackagesPackageSummaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPackagesPackageSummary200Response
func (a *PackageObservabilityAPIService) GetPackagesPackageSummaryExecute(r ApiGetPackagesPackageSummaryRequest) (*GetPackagesPackageSummary200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPackagesPackageSummary200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PackageObservabilityAPIService.GetPackagesPackageSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/packages/package-summary"

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
