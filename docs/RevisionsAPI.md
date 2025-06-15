# \RevisionsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FixPlans**](RevisionsAPI.md#FixPlans) | **Get** /revisions/{locator}/report/remediation-guidance | 
[**GetRevisionAttributionEmail**](RevisionsAPI.md#GetRevisionAttributionEmail) | **Get** /revisions/{locator}/attribution/email | 
[**GetRevisionAttributionJSON**](RevisionsAPI.md#GetRevisionAttributionJSON) | **Get** /revisions/{locator}/attribution/json | 
[**NoticeFiles**](RevisionsAPI.md#NoticeFiles) | **Get** /revisions/{locator}/notice-files | 
[**OriginalSbom**](RevisionsAPI.md#OriginalSbom) | **Get** /revisions/{locator}/original-sbom | 
[**UpdateRevision**](RevisionsAPI.md#UpdateRevision) | **Patch** /revisions/{locator} | 



## FixPlans

> *os.File FixPlans(ctx, locator).Preview(preview).Format(format).Bundle(bundle).ExcludeQuickWins(excludeQuickWins).ExcludeHighPriority(excludeHighPriority).ExcludedLowPriority(excludedLowPriority).ExcludeOutdatedDependencies(excludeOutdatedDependencies).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | 
	preview := true // bool | Whether to preview the report (default is false) (optional)
	format := "format_example" // string | What format to return the report in (default is PDF) (optional)
	bundle := true // bool | Whether to bundle the report with json file and return as a zip (default is false) (optional)
	excludeQuickWins := true // bool | Whether to exclude Quick Wins section (default is false) (optional)
	excludeHighPriority := true // bool | Whether to exclude High Priority section (default is false) (optional)
	excludedLowPriority := true // bool | Whether to exclude Low Priority section (default is false) (optional)
	excludeOutdatedDependencies := true // bool | Whether to exclude Outdated Dependencies section (default is false) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.FixPlans(context.Background(), locator).Preview(preview).Format(format).Bundle(bundle).ExcludeQuickWins(excludeQuickWins).ExcludeHighPriority(excludeHighPriority).ExcludedLowPriority(excludedLowPriority).ExcludeOutdatedDependencies(excludeOutdatedDependencies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.FixPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FixPlans`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.FixPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFixPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preview** | **bool** | Whether to preview the report (default is false) | 
 **format** | **string** | What format to return the report in (default is PDF) | 
 **bundle** | **bool** | Whether to bundle the report with json file and return as a zip (default is false) | 
 **excludeQuickWins** | **bool** | Whether to exclude Quick Wins section (default is false) | 
 **excludeHighPriority** | **bool** | Whether to exclude High Priority section (default is false) | 
 **excludedLowPriority** | **bool** | Whether to exclude Low Priority section (default is false) | 
 **excludeOutdatedDependencies** | **bool** | Whether to exclude Outdated Dependencies section (default is false) | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/zip, text/html, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionEmail

> GetRevisionAttributionEmail200Response GetRevisionAttributionEmail(ctx, locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludePackageLabels(excludePackageLabels).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | the url-encoded locator of the revision
	access := "access_example" // string | The public ID (optional)
	preview := true // bool | Whether to preview the report (default is false) (optional)
	format := "format_example" // string | The format of the report (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is true) (optional)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is true) (optional)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional)
	includeLicenseScan := true // bool | Whether to include the license scan (default is false) (optional)
	includeProjectLicense := true // bool | Whether to include the project license (default is false) (optional)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional)
	includeFileMatches := true // bool | Whether to include the file matches (default is false) (optional)
	includeOpenVulnerabilities := true // bool | Whether to include the open vulnerabilities (default is false) (optional)
	includeClosedVulnerabilities := true // bool | Whether to include the closed vulnerabilities (default is false) (optional)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional)
	includeLicenseHeaders := true // bool | Whether to include the license headers (default is false) (optional)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false). (optional)
	excludePackageLabels := []string{"Inner_example"} // []string | Exclude dependencies with particular package labels from the report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionEmail(context.Background(), locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludePackageLabels(excludePackageLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionEmail`: GetRevisionAttributionEmail200Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **access** | **string** | The public ID | 
 **preview** | **bool** | Whether to preview the report (default is false) | 
 **format** | **string** | The format of the report | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is true) | 
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is true) | 
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | 
 **includeLicenseScan** | **bool** | Whether to include the license scan (default is false) | 
 **includeProjectLicense** | **bool** | Whether to include the project license (default is false) | 
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | 
 **includeFileMatches** | **bool** | Whether to include the file matches (default is false) | 
 **includeOpenVulnerabilities** | **bool** | Whether to include the open vulnerabilities (default is false) | 
 **includeClosedVulnerabilities** | **bool** | Whether to include the closed vulnerabilities (default is false) | 
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | 
 **includeLicenseHeaders** | **bool** | Whether to include the license headers (default is false) | 
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false). | 
 **excludePackageLabels** | **[]string** | Exclude dependencies with particular package labels from the report | 

### Return type

[**GetRevisionAttributionEmail200Response**](GetRevisionAttributionEmail200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionJSON

> GetRevisionAttributionJSON200Response GetRevisionAttributionJSON(ctx, locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludePackageLabels(excludePackageLabels).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | the url-encoded locator of the revision
	preview := true // bool | Whether to preview the report (default is false) (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is true) (optional)
	includeHashAndVersionData := true // bool | Whether to include hash and version data (default is false) (optional)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional)
	includeFileMatches := true // bool | Whether to include the file matches (default is false) (optional)
	includeOpenVulnerabilities := true // bool | Whether to include the open vulnerabilities (default is false) (optional)
	includeClosedVulnerabilities := true // bool | Whether to include the closed vulnerabilities (default is false) (optional)
	includeNoticeFiles := true // bool | Whether to include the notice files match data (default is false) (optional)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional)
	excludePackageLabels := []string{"Inner_example"} // []string | Exclude dependencies with particular package labels from the report (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionJSON(context.Background(), locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludePackageLabels(excludePackageLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionJSON``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionJSON`: GetRevisionAttributionJSON200Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionJSON`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionJSONRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preview** | **bool** | Whether to preview the report (default is false) | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is true) | 
 **includeHashAndVersionData** | **bool** | Whether to include hash and version data (default is false) | 
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | 
 **includeFileMatches** | **bool** | Whether to include the file matches (default is false) | 
 **includeOpenVulnerabilities** | **bool** | Whether to include the open vulnerabilities (default is false) | 
 **includeClosedVulnerabilities** | **bool** | Whether to include the closed vulnerabilities (default is false) | 
 **includeNoticeFiles** | **bool** | Whether to include the notice files match data (default is false) | 
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | 
 **excludePackageLabels** | **[]string** | Exclude dependencies with particular package labels from the report | 

### Return type

[**GetRevisionAttributionJSON200Response**](GetRevisionAttributionJSON200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NoticeFiles

> []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner NoticeFiles(ctx, locator).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.NoticeFiles(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.NoticeFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NoticeFiles`: []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.NoticeFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNoticeFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OriginalSbom

> OriginalSbom(ctx, locator).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RevisionsAPI.OriginalSbom(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.OriginalSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOriginalSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRevision

> GetProjectRevisions200ResponseBranchValueInner UpdateRevision(ctx, locator).UpdateRevisionRequest(updateRevisionRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | 
	updateRevisionRequest := *openapiclient.NewUpdateRevisionRequest() // UpdateRevisionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.UpdateRevision(context.Background(), locator).UpdateRevisionRequest(updateRevisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.UpdateRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRevision`: GetProjectRevisions200ResponseBranchValueInner
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.UpdateRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRevisionRequest** | [**UpdateRevisionRequest**](UpdateRevisionRequest.md) |  | 

### Return type

[**GetProjectRevisions200ResponseBranchValueInner**](GetProjectRevisions200ResponseBranchValueInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

