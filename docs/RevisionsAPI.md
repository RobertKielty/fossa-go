# \RevisionsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRevisionAttributionPublicReportV2**](RevisionsAPI.md#CreateRevisionAttributionPublicReportV2) | **Post** /v2/revisions/{locator}/attribution/public | 
[**FixPlans**](RevisionsAPI.md#FixPlans) | **Get** /revisions/{locator}/report/remediation-guidance | 
[**GetRevisionAttributionDownloadV2**](RevisionsAPI.md#GetRevisionAttributionDownloadV2) | **Get** /v2/revisions/{locator}/attribution/download | 
[**GetRevisionAttributionEmail**](RevisionsAPI.md#GetRevisionAttributionEmail) | **Get** /revisions/{locator}/attribution/email | 
[**GetRevisionAttributionEmailV2**](RevisionsAPI.md#GetRevisionAttributionEmailV2) | **Get** /v2/revisions/{locator}/attribution/email | 
[**GetRevisionAttributionFullV2**](RevisionsAPI.md#GetRevisionAttributionFullV2) | **Get** /v2/revisions/{locator}/attribution/full/{format} | 
[**GetRevisionAttributionJSON**](RevisionsAPI.md#GetRevisionAttributionJSON) | **Get** /revisions/{locator}/attribution/json | 
[**GetRevisionAttributionJSONV2**](RevisionsAPI.md#GetRevisionAttributionJSONV2) | **Get** /v2/revisions/{locator}/attribution/json | 
[**GetRevisionAttributionPreviewV2**](RevisionsAPI.md#GetRevisionAttributionPreviewV2) | **Get** /v2/revisions/{locator}/attribution/preview | 
[**GetRevisionAttributionV2**](RevisionsAPI.md#GetRevisionAttributionV2) | **Get** /v2/revisions/{locator}/attribution | 
[**GetRevisionDependencies**](RevisionsAPI.md#GetRevisionDependencies) | **Get** /revisions/{locator}/dependencies | 
[**GetRevisionDependenciesPost**](RevisionsAPI.md#GetRevisionDependenciesPost) | **Post** /revisions/{locator}/list-dependencies | 
[**GetRevisionScans**](RevisionsAPI.md#GetRevisionScans) | **Get** /revisions/{locator}/scans | 
[**NoticeFiles**](RevisionsAPI.md#NoticeFiles) | **Get** /revisions/{locator}/notice-files | 
[**OriginalSbom**](RevisionsAPI.md#OriginalSbom) | **Get** /revisions/{locator}/original-sbom | 
[**UpdateRevision**](RevisionsAPI.md#UpdateRevision) | **Patch** /revisions/{locator} | 



## CreateRevisionAttributionPublicReportV2

> CreateRevisionAttributionPublicReportV2202Response CreateRevisionAttributionPublicReportV2(ctx, locator).Format(format).Emails(emails).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	format := "format_example" // string | The format of the report (optional)
	emails := "emails_example" // string | A single email address used as the report recipient and, for SPDX-style reports, the author identifier. Despite the plural name, this accepts exactly one address — not a list.  (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is false) (optional) (default to false)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional) (default to false)
	includeLicenseScan := true // bool | Whether to include the first-party license scan (default is false) (optional) (default to false)
	includeProjectLicense := true // bool | Whether to include the project's declared license (default is false) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional) (default to false)
	includeLicenseHeaders := true // bool | Whether to include license headers (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.CreateRevisionAttributionPublicReportV2(context.Background(), locator).Format(format).Emails(emails).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.CreateRevisionAttributionPublicReportV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRevisionAttributionPublicReportV2`: CreateRevisionAttributionPublicReportV2202Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.CreateRevisionAttributionPublicReportV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRevisionAttributionPublicReportV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The format of the report | 
 **emails** | **string** | A single email address used as the report recipient and, for SPDX-style reports, the author identifier. Despite the plural name, this accepts exactly one address — not a list.  | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is false) | [default to false]
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | [default to false]
 **includeLicenseScan** | **bool** | Whether to include the first-party license scan (default is false) | [default to false]
 **includeProjectLicense** | **bool** | Whether to include the project&#39;s declared license (default is false) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | [default to false]
 **includeLicenseHeaders** | **bool** | Whether to include license headers (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

### Return type

[**CreateRevisionAttributionPublicReportV2202Response**](CreateRevisionAttributionPublicReportV2202Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FixPlans

> *os.File FixPlans(ctx, locator).Preview(preview).Format(format).Bundle(bundle).ExcludeQuickWins(excludeQuickWins).ExcludeHighPriority(excludeHighPriority).ExcludeLowPriority(excludeLowPriority).ExcludeOutdatedDependencies(excludeOutdatedDependencies).Demo(demo).IncludeTransitiveVulns(includeTransitiveVulns).DeduplicateOutdatedDeps(deduplicateOutdatedDeps).IncludeMalware(includeMalware).Execute()





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
	excludeLowPriority := true // bool | Whether to exclude Low Priority section (default is false) (optional)
	excludeOutdatedDependencies := true // bool | Whether to exclude Outdated Dependencies section (default is false) (optional)
	demo := true // bool | Whether to generate the report in demo mode (default is false) (optional)
	includeTransitiveVulns := true // bool | Whether to include transitive vulnerabilities (default is false) (optional)
	deduplicateOutdatedDeps := true // bool | Whether to deduplicate outdated dependencies (default is false) (optional)
	includeMalware := true // bool | Whether to include malware findings (default is false). Only takes effect when the organization has the malware-issues feature enabled.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.FixPlans(context.Background(), locator).Preview(preview).Format(format).Bundle(bundle).ExcludeQuickWins(excludeQuickWins).ExcludeHighPriority(excludeHighPriority).ExcludeLowPriority(excludeLowPriority).ExcludeOutdatedDependencies(excludeOutdatedDependencies).Demo(demo).IncludeTransitiveVulns(includeTransitiveVulns).DeduplicateOutdatedDeps(deduplicateOutdatedDeps).IncludeMalware(includeMalware).Execute()
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
 **excludeLowPriority** | **bool** | Whether to exclude Low Priority section (default is false) | 
 **excludeOutdatedDependencies** | **bool** | Whether to exclude Outdated Dependencies section (default is false) | 
 **demo** | **bool** | Whether to generate the report in demo mode (default is false) | 
 **includeTransitiveVulns** | **bool** | Whether to include transitive vulnerabilities (default is false) | 
 **deduplicateOutdatedDeps** | **bool** | Whether to deduplicate outdated dependencies (default is false) | 
 **includeMalware** | **bool** | Whether to include malware findings (default is false). Only takes effect when the organization has the malware-issues feature enabled.  | 

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


## GetRevisionAttributionDownloadV2

> *os.File GetRevisionAttributionDownloadV2(ctx, locator).Format(format).Access(access).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	format := "format_example" // string | The format of the report (optional)
	access := "access_example" // string | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is false) (optional) (default to false)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional) (default to false)
	includeLicenseScan := true // bool | Whether to include the first-party license scan (default is false) (optional) (default to false)
	includeProjectLicense := true // bool | Whether to include the project's declared license (default is false) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional) (default to false)
	includeLicenseHeaders := true // bool | Whether to include license headers (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	includeHashAndVersionData := true // bool | Whether to include hash and version data (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionDownloadV2(context.Background(), locator).Format(format).Access(access).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionDownloadV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionDownloadV2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionDownloadV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionDownloadV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The format of the report | 
 **access** | **string** | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is false) | [default to false]
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | [default to false]
 **includeLicenseScan** | **bool** | Whether to include the first-party license scan (default is false) | [default to false]
 **includeProjectLicense** | **bool** | Whether to include the project&#39;s declared license (default is false) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | [default to false]
 **includeLicenseHeaders** | **bool** | Whether to include license headers (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **includeHashAndVersionData** | **bool** | Whether to include hash and version data (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionEmail

> GetRevisionAttributionEmail200Response GetRevisionAttributionEmail(ctx, locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()





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
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionEmail(context.Background(), locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()
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
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  | 

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


## GetRevisionAttributionEmailV2

> GetRevisionAttributionEmailV2200Response GetRevisionAttributionEmailV2(ctx, locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	access := "access_example" // string | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. (optional)
	preview := true // bool | Whether to preview the report inline rather than generate a downloadable file (default is false) (optional) (default to false)
	format := "format_example" // string | The format of the report (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is false) (optional) (default to false)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional) (default to false)
	includeLicenseScan := true // bool | Whether to include the first-party license scan (default is false) (optional) (default to false)
	includeProjectLicense := true // bool | Whether to include the project's declared license (default is false) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional) (default to false)
	includeLicenseHeaders := true // bool | Whether to include license headers (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionEmailV2(context.Background(), locator).Access(access).Preview(preview).Format(format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionEmailV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionEmailV2`: GetRevisionAttributionEmailV2200Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionEmailV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionEmailV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **access** | **string** | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. | 
 **preview** | **bool** | Whether to preview the report inline rather than generate a downloadable file (default is false) | [default to false]
 **format** | **string** | The format of the report | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is false) | [default to false]
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | [default to false]
 **includeLicenseScan** | **bool** | Whether to include the first-party license scan (default is false) | [default to false]
 **includeProjectLicense** | **bool** | Whether to include the project&#39;s declared license (default is false) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | [default to false]
 **includeLicenseHeaders** | **bool** | Whether to include license headers (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

### Return type

[**GetRevisionAttributionEmailV2200Response**](GetRevisionAttributionEmailV2200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionFullV2

> *os.File GetRevisionAttributionFullV2(ctx, locator, format).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	format := "format_example" // string | The format of the report

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionFullV2(context.Background(), locator, format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionFullV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionFullV2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionFullV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 
**format** | **string** | The format of the report | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionFullV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionJSON

> GetRevisionAttributionJSON200Response GetRevisionAttributionJSON(ctx, locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()





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
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionJSON(context.Background(), locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()
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
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

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


## GetRevisionAttributionJSONV2

> GetRevisionAttributionJSON200Response GetRevisionAttributionJSONV2(ctx, locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	preview := true // bool | Whether to preview the report inline rather than generate a downloadable file (default is false) (optional) (default to false)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeHashAndVersionData := true // bool | Whether to include hash and version data (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeNoticeFiles := true // bool | Whether to include the notice file match data (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionJSONV2(context.Background(), locator).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeNoticeFiles(includeNoticeFiles).IncludePackageLabels(includePackageLabels).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionJSONV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionJSONV2`: GetRevisionAttributionJSON200Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionJSONV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionJSONV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **preview** | **bool** | Whether to preview the report inline rather than generate a downloadable file (default is false) | [default to false]
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeHashAndVersionData** | **bool** | Whether to include hash and version data (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeNoticeFiles** | **bool** | Whether to include the notice file match data (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

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


## GetRevisionAttributionPreviewV2

> string GetRevisionAttributionPreviewV2(ctx, locator).Format(format).Access(access).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	format := "format_example" // string | The format of the report (optional)
	access := "access_example" // string | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. (optional)
	preview := true // bool | Whether to preview the report inline rather than generate a downloadable file (default is false) (optional) (default to false)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is false) (optional) (default to false)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional) (default to false)
	includeLicenseScan := true // bool | Whether to include the first-party license scan (default is false) (optional) (default to false)
	includeProjectLicense := true // bool | Whether to include the project's declared license (default is false) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional) (default to false)
	includeLicenseHeaders := true // bool | Whether to include license headers (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	includeHashAndVersionData := true // bool | Whether to include hash and version data (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionPreviewV2(context.Background(), locator).Format(format).Access(access).Preview(preview).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionPreviewV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionPreviewV2`: string
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionPreviewV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionPreviewV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The format of the report | 
 **access** | **string** | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. | 
 **preview** | **bool** | Whether to preview the report inline rather than generate a downloadable file (default is false) | [default to false]
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is false) | [default to false]
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | [default to false]
 **includeLicenseScan** | **bool** | Whether to include the first-party license scan (default is false) | [default to false]
 **includeProjectLicense** | **bool** | Whether to include the project&#39;s declared license (default is false) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | [default to false]
 **includeLicenseHeaders** | **bool** | Whether to include license headers (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **includeHashAndVersionData** | **bool** | Whether to include hash and version data (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionAttributionV2

> *os.File GetRevisionAttributionV2(ctx, locator).Format(format).Access(access).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()





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
	locator := "locator_example" // string | The URL-encoded locator of the revision
	format := "format_example" // string | The format of the report (optional)
	access := "access_example" // string | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. (optional)
	includeDeepDependencies := true // bool | Whether to include deep dependencies (default is false) (optional) (default to false)
	includeDirectDependencies := true // bool | Whether to include direct dependencies (default is false) (optional) (default to false)
	includeLicenseList := true // bool | Whether to include the license list (default is false) (optional) (default to false)
	includeLicenseScan := true // bool | Whether to include the first-party license scan (default is false) (optional) (default to false)
	includeProjectLicense := true // bool | Whether to include the project's declared license (default is false) (optional) (default to false)
	includeCopyrightList := true // bool | Whether to include the copyright list (default is false) (optional) (default to false)
	includeFileMatches := true // bool | Whether to include license file matches (default is false) (optional) (default to false)
	includeOpenVulnerabilities := true // bool | Whether to include open vulnerabilities (default is false) (optional) (default to false)
	includeClosedVulnerabilities := true // bool | Whether to include closed vulnerabilities (default is false) (optional) (default to false)
	includeDependencySummary := true // bool | Whether to include the dependency summary (default is false) (optional) (default to false)
	includeLicenseHeaders := true // bool | Whether to include license headers (default is false) (optional) (default to false)
	includePackageLabels := true // bool | Whether to include the package labels assigned to each dependency (default is false) (optional) (default to false)
	includeHashAndVersionData := true // bool | Whether to include hash and version data (default is false) (optional) (default to false)
	excludeUnknownDependencies := true // bool | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) (optional) (default to false)
	excludeFields := *openapiclient.NewQueueReleaseGroupAttributionReportV2ExcludeFieldsParameter() // QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter | Object controlling which dependencies are excluded from the report. The only supported nested field is `packageLabels`: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the `qs` library, so the array is sent using bracket-and-index notation rather than standard OpenAPI `deepObject` serialization. For example, to exclude two labels send (before URL-encoding): `excludeFields[packageLabels][0]=internal&excludeFields[packageLabels][1]=vendored`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionAttributionV2(context.Background(), locator).Format(format).Access(access).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).IncludePackageLabels(includePackageLabels).IncludeHashAndVersionData(includeHashAndVersionData).ExcludeUnknownDependencies(excludeUnknownDependencies).ExcludeFields(excludeFields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionAttributionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionAttributionV2`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionAttributionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionAttributionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The format of the report | 
 **access** | **string** | The public BOM access ID. When provided, the report is generated for the public BOM associated with this ID. | 
 **includeDeepDependencies** | **bool** | Whether to include deep dependencies (default is false) | [default to false]
 **includeDirectDependencies** | **bool** | Whether to include direct dependencies (default is false) | [default to false]
 **includeLicenseList** | **bool** | Whether to include the license list (default is false) | [default to false]
 **includeLicenseScan** | **bool** | Whether to include the first-party license scan (default is false) | [default to false]
 **includeProjectLicense** | **bool** | Whether to include the project&#39;s declared license (default is false) | [default to false]
 **includeCopyrightList** | **bool** | Whether to include the copyright list (default is false) | [default to false]
 **includeFileMatches** | **bool** | Whether to include license file matches (default is false) | [default to false]
 **includeOpenVulnerabilities** | **bool** | Whether to include open vulnerabilities (default is false) | [default to false]
 **includeClosedVulnerabilities** | **bool** | Whether to include closed vulnerabilities (default is false) | [default to false]
 **includeDependencySummary** | **bool** | Whether to include the dependency summary (default is false) | [default to false]
 **includeLicenseHeaders** | **bool** | Whether to include license headers (default is false) | [default to false]
 **includePackageLabels** | **bool** | Whether to include the package labels assigned to each dependency (default is false) | [default to false]
 **includeHashAndVersionData** | **bool** | Whether to include hash and version data (default is false) | [default to false]
 **excludeUnknownDependencies** | **bool** | Whether to exclude unknown (unresolved) dependencies from the report (default is false, meaning unknown dependencies are included) | [default to false]
 **excludeFields** | [**QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter**](QueueReleaseGroupAttributionReportV2ExcludeFieldsParameter.md) | Object controlling which dependencies are excluded from the report. The only supported nested field is &#x60;packageLabels&#x60;: a non-empty array of package-label values; dependencies carrying any of these labels are excluded from the report.  The server parses the query string with the &#x60;qs&#x60; library, so the array is sent using bracket-and-index notation rather than standard OpenAPI &#x60;deepObject&#x60; serialization. For example, to exclude two labels send (before URL-encoding): &#x60;excludeFields[packageLabels][0]&#x3D;internal&amp;excludeFields[packageLabels][1]&#x3D;vendored&#x60;.  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionDependencies

> []GetRevisionDependenciesPost200ResponseInner GetRevisionDependencies(ctx, locator).Limit(limit).Offset(offset).IncludeIgnored(includeIgnored).IncludeHashData(includeHashData).IncludeLicenseText(includeLicenseText).IncludeLocators(includeLocators).Execute()





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
	locator := "custom+1234/my-project$abcd1234" // string | The URL-encoded locator of the revision
	limit := int32(100) // int32 | Maximum number of dependencies to return. The value is clamped server-side to the range 25–100: any value below 25 is treated as 25, and any value above 100 is treated as 100.  (optional)
	offset := int32(0) // int32 | Number of dependencies to skip for pagination (optional)
	includeIgnored := true // bool | Whether to include ignored dependencies in the response (optional) (default to false)
	includeHashData := true // bool | Whether to include hash and version data for dependencies (optional) (default to false)
	includeLicenseText := true // bool | Whether to include full license text in the license information (optional) (default to false)
	includeLocators := []string{"Inner_example"} // []string | Array of locators to filter dependencies. Only dependencies matching these locators will be returned. Note: For large lists of locators that may exceed URL length limits, use POST /api/revisions/:locator/deps instead.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionDependencies(context.Background(), locator).Limit(limit).Offset(offset).IncludeIgnored(includeIgnored).IncludeHashData(includeHashData).IncludeLicenseText(includeLicenseText).IncludeLocators(includeLocators).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependencies`: []GetRevisionDependenciesPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of dependencies to return. The value is clamped server-side to the range 25–100: any value below 25 is treated as 25, and any value above 100 is treated as 100.  | 
 **offset** | **int32** | Number of dependencies to skip for pagination | 
 **includeIgnored** | **bool** | Whether to include ignored dependencies in the response | [default to false]
 **includeHashData** | **bool** | Whether to include hash and version data for dependencies | [default to false]
 **includeLicenseText** | **bool** | Whether to include full license text in the license information | [default to false]
 **includeLocators** | **[]string** | Array of locators to filter dependencies. Only dependencies matching these locators will be returned. Note: For large lists of locators that may exceed URL length limits, use POST /api/revisions/:locator/deps instead.  | 

### Return type

[**[]GetRevisionDependenciesPost200ResponseInner**](GetRevisionDependenciesPost200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionDependenciesPost

> []GetRevisionDependenciesPost200ResponseInner GetRevisionDependenciesPost(ctx, locator).GetRevisionDependenciesPostRequest(getRevisionDependenciesPostRequest).Execute()





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
	locator := "custom+1234/my-project$abcd1234" // string | The URL-encoded locator of the revision
	getRevisionDependenciesPostRequest := *openapiclient.NewGetRevisionDependenciesPostRequest() // GetRevisionDependenciesPostRequest | Query parameters for filtering and configuring the dependency response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionDependenciesPost(context.Background(), locator).GetRevisionDependenciesPostRequest(getRevisionDependenciesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionDependenciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependenciesPost`: []GetRevisionDependenciesPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionDependenciesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionDependenciesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getRevisionDependenciesPostRequest** | [**GetRevisionDependenciesPostRequest**](GetRevisionDependenciesPostRequest.md) | Query parameters for filtering and configuring the dependency response | 

### Return type

[**[]GetRevisionDependenciesPost200ResponseInner**](GetRevisionDependenciesPost200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionScans

> GetRevisionScans200Response GetRevisionScans(ctx, locator).Page(page).PageSize(pageSize).Execute()





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
	locator := "custom%2B1234%2Fmy-project%24abcd1234" // string | The URL-encoded locator of the revision
	page := int32(1) // int32 | The 1-indexed page of results to return (optional) (default to 1)
	pageSize := int32(10) // int32 | The number of scans to return per page (maximum 50) (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RevisionsAPI.GetRevisionScans(context.Background(), locator).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RevisionsAPI.GetRevisionScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionScans`: GetRevisionScans200Response
	fmt.Fprintf(os.Stdout, "Response from `RevisionsAPI.GetRevisionScans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The 1-indexed page of results to return | [default to 1]
 **pageSize** | **int32** | The number of scans to return per page (maximum 50) | [default to 10]

### Return type

[**GetRevisionScans200Response**](GetRevisionScans200Response.md)

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
	updateRevisionRequest := *openapiclient.NewUpdateRevisionRequest() // UpdateRevisionRequest | 

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

