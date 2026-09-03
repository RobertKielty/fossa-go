# \IssuesAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssueDispute**](IssuesAPI.md#CreateIssueDispute) | **Post** /v2/issues/{issueId}/disputes | 
[**DeleteIssueException**](IssuesAPI.md#DeleteIssueException) | **Delete** /v2/issues/exceptions/{id} | 
[**DeleteIssueExceptions**](IssuesAPI.md#DeleteIssueExceptions) | **Delete** /v2/issues/exceptions | 
[**DeleteProjectGenerateAttributionSlug**](IssuesAPI.md#DeleteProjectGenerateAttributionSlug) | **Delete** /projects/{locator}/generate_attribution_slug | 
[**ExtendIssueException**](IssuesAPI.md#ExtendIssueException) | **Put** /v2/issues/exceptions/{id} | Extend an issue exception expiration date
[**GetGlobalIssuesCSV**](IssuesAPI.md#GetGlobalIssuesCSV) | **Get** /v2/issues/csv/global | 
[**GetIssue**](IssuesAPI.md#GetIssue) | **Get** /v2/issues/{issueId} | 
[**GetIssueAffectedProjects**](IssuesAPI.md#GetIssueAffectedProjects) | **Get** /v2/issues/{issueId}/affected-projects | 
[**GetIssueCWEs**](IssuesAPI.md#GetIssueCWEs) | **Get** /v2/issues/cwes | 
[**GetIssueDiffComparisonSummaries**](IssuesAPI.md#GetIssueDiffComparisonSummaries) | **Get** /v2/issues/compare/summaries | 
[**GetIssueException**](IssuesAPI.md#GetIssueException) | **Get** /v2/issues/exceptions/{id} | 
[**GetIssueExceptions**](IssuesAPI.md#GetIssueExceptions) | **Get** /v2/issues/exceptions | 
[**GetIssuePackageManagers**](IssuesAPI.md#GetIssuePackageManagers) | **Get** /v2/issues/package-managers | 
[**GetIssueStatuses**](IssuesAPI.md#GetIssueStatuses) | **Get** /v2/issues/statuses | 
[**GetIssues**](IssuesAPI.md#GetIssues) | **Get** /v2/issues | 
[**GetIssuesByCategory**](IssuesAPI.md#GetIssuesByCategory) | **Get** /v2/issues/categories | 
[**GetIssuesByRevision**](IssuesAPI.md#GetIssuesByRevision) | **Get** /v2/issues/revisions | 
[**GetIssuesByType**](IssuesAPI.md#GetIssuesByType) | **Get** /v2/issues/types | 
[**GetLicenseList**](IssuesAPI.md#GetLicenseList) | **Get** /v2/issues/license-list | 
[**GetProjectCSVExportIssues**](IssuesAPI.md#GetProjectCSVExportIssues) | **Get** /projects/{locator}/export-issues/csv | 
[**GetProjectExportIssues**](IssuesAPI.md#GetProjectExportIssues) | **Get** /projects/{locator}/export-issues | 
[**GetProjectJSONExportIssues**](IssuesAPI.md#GetProjectJSONExportIssues) | **Get** /projects/{locator}/export-issues/json | 
[**UpdateIssues**](IssuesAPI.md#UpdateIssues) | **Put** /v2/issues | 



## CreateIssueDispute

> CreateIssueDispute200Response CreateIssueDispute(ctx, issueId).CreateIssueDisputeRequest(createIssueDisputeRequest).Execute()





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
	issueId := int32(56) // int32 | ID of the issue that is being disputed.
	createIssueDisputeRequest := *openapiclient.NewCreateIssueDisputeRequest("Reason_example") // CreateIssueDisputeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.CreateIssueDispute(context.Background(), issueId).CreateIssueDisputeRequest(createIssueDisputeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.CreateIssueDispute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueDispute`: CreateIssueDispute200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.CreateIssueDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | ID of the issue that is being disputed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIssueDisputeRequest** | [**CreateIssueDisputeRequest**](CreateIssueDisputeRequest.md) |  | 

### Return type

[**CreateIssueDispute200Response**](CreateIssueDispute200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueException

> int32 DeleteIssueException(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the issue exception to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.DeleteIssueException(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.DeleteIssueException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueException`: int32
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.DeleteIssueException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the issue exception to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**int32**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueExceptions

> UpdateIssues200Response DeleteIssueExceptions(ctx).DeleteIssueExceptionsRequest(deleteIssueExceptionsRequest).Execute()





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
	deleteIssueExceptionsRequest := openapiclient.deleteIssueExceptions_request{DeleteIssueExceptionsRequestOneOf: openapiclient.NewDeleteIssueExceptionsRequestOneOf()} // DeleteIssueExceptionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.DeleteIssueExceptions(context.Background()).DeleteIssueExceptionsRequest(deleteIssueExceptionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.DeleteIssueExceptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueExceptions`: UpdateIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.DeleteIssueExceptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteIssueExceptionsRequest** | [**DeleteIssueExceptionsRequest**](DeleteIssueExceptionsRequest.md) |  | 

### Return type

[**UpdateIssues200Response**](UpdateIssues200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectGenerateAttributionSlug

> DeleteProjectGenerateAttributionSlug(ctx, locator).Execute()





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
	r, err := apiClient.IssuesAPI.DeleteProjectGenerateAttributionSlug(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.DeleteProjectGenerateAttributionSlug``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteProjectGenerateAttributionSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtendIssueException

> ExtendIssueException(ctx, id).ExtendIssueExceptionRequest(extendIssueExceptionRequest).Execute()

Extend an issue exception expiration date



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	id := int32(56) // int32 | ID of the issue exception to extend
	extendIssueExceptionRequest := *openapiclient.NewExtendIssueExceptionRequest(time.Now()) // ExtendIssueExceptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuesAPI.ExtendIssueException(context.Background(), id).ExtendIssueExceptionRequest(extendIssueExceptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.ExtendIssueException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the issue exception to extend | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendIssueExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extendIssueExceptionRequest** | [**ExtendIssueExceptionRequest**](ExtendIssueExceptionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalIssuesCSV

> *os.File GetGlobalIssuesCSV(ctx).Email(email).TeamIds(teamIds).ReleaseGroupId(releaseGroupId).IncludeIssuesList(includeIssuesList).IncludeDailyCounts(includeDailyCounts).Execute()





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
	email := true // bool | When `true`, we will submit the report for background processing and deliver via email when ready. Otherwise the report will be streamed via API (optional)
	teamIds := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Scope the report to one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"`. Invalid (non-integer) team IDs are rejected with a `400`. Ignored for organizations on the free tier.  (optional)
	releaseGroupId := int32(56) // int32 | Scope the issue-list export to a single release group, using its latest release (the issues shown on the release group's Issues page). Requires view access to the release group; otherwise rejected with a `403`. When set, `teamIds` is ignored — the export is scoped to the release group only, not an intersection of both filters. When combined with `includeDailyCounts`, the bundle's counts CSV is `release_counts.csv` rather than the daily `issue_counts.csv`.  (optional)
	includeIssuesList := true // bool | Whether to include the per-category issue-list CSVs in the bundle. Defaults to `true` when omitted. Setting this to `false` without `includeDailyCounts=true` selects no reports and is rejected with a `400`.  (optional) (default to true)
	includeDailyCounts := true // bool | Whether to include a counts CSV in the bundle. Defaults to `false`. For team/organization scope this is `issue_counts.csv`, a daily time series of active, ignored, and remediated counts. When `releaseGroupId` is set it is instead `release_counts.csv`, holding one row per release (New / Remediated / Unchanged versus the previous release), capped at the 10 most recent scanned releases.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetGlobalIssuesCSV(context.Background()).Email(email).TeamIds(teamIds).ReleaseGroupId(releaseGroupId).IncludeIssuesList(includeIssuesList).IncludeDailyCounts(includeDailyCounts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetGlobalIssuesCSV``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalIssuesCSV`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetGlobalIssuesCSV`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalIssuesCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **bool** | When &#x60;true&#x60;, we will submit the report for background processing and deliver via email when ready. Otherwise the report will be streamed via API | 
 **teamIds** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Scope the report to one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60;. Invalid (non-integer) team IDs are rejected with a &#x60;400&#x60;. Ignored for organizations on the free tier.  | 
 **releaseGroupId** | **int32** | Scope the issue-list export to a single release group, using its latest release (the issues shown on the release group&#39;s Issues page). Requires view access to the release group; otherwise rejected with a &#x60;403&#x60;. When set, &#x60;teamIds&#x60; is ignored — the export is scoped to the release group only, not an intersection of both filters. When combined with &#x60;includeDailyCounts&#x60;, the bundle&#39;s counts CSV is &#x60;release_counts.csv&#x60; rather than the daily &#x60;issue_counts.csv&#x60;.  | 
 **includeIssuesList** | **bool** | Whether to include the per-category issue-list CSVs in the bundle. Defaults to &#x60;true&#x60; when omitted. Setting this to &#x60;false&#x60; without &#x60;includeDailyCounts&#x3D;true&#x60; selects no reports and is rejected with a &#x60;400&#x60;.  | [default to true]
 **includeDailyCounts** | **bool** | Whether to include a counts CSV in the bundle. Defaults to &#x60;false&#x60;. For team/organization scope this is &#x60;issue_counts.csv&#x60;, a daily time series of active, ignored, and remediated counts. When &#x60;releaseGroupId&#x60; is set it is instead &#x60;release_counts.csv&#x60;, holding one row per release (New / Remediated / Unchanged versus the previous release), capped at the 10 most recent scanned releases.  | [default to false]

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


## GetIssue

> GetIssue200Response GetIssue(ctx, issueId).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).Execute()





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
	issueId := int32(56) // int32 | Issue ID
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssue(context.Background(), issueId).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssue`: GetIssue200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | Issue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 

### Return type

[**GetIssue200Response**](GetIssue200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueAffectedProjects

> []GetIssueAffectedProjects200ResponseInner GetIssueAffectedProjects(ctx, issueId).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).Execute()





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
	issueId := int32(56) // int32 | Issue ID
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueAffectedProjects(context.Background(), issueId).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueAffectedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueAffectedProjects`: []GetIssueAffectedProjects200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueAffectedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | Issue ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueAffectedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 

### Return type

[**[]GetIssueAffectedProjects200ResponseInner**](GetIssueAffectedProjects200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueCWEs

> GetIssueCWEs200Response GetIssueCWEs(ctx).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).TeamId(teamId).Execute()





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
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	status := "status_example" // string | Issue status (optional)
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueCWEs(context.Background()).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueCWEs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueCWEs`: GetIssueCWEs200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueCWEs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueCWEsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeType** | **string** | Scope of issues to view / update | 
 **status** | **string** | Issue status | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssueCWEs200Response**](GetIssueCWEs200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueDiffComparisonSummaries

> GetIssueDiffComparisonSummaries200Response GetIssueDiffComparisonSummaries(ctx).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToRelease(scopeCompareToRelease).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterCwes(filterCwes).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterEpss(filterEpss).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToRelease := "scopeCompareToRelease_example" // string | The release ID to compare issues with. Only available for Release Group Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	ids := []int32{int32(123)} // []int32 | Filter by specific issue IDs (optional)
	filterRevisionIds := []string{"Inner_example"} // []string | Filter by specific revision IDs (optional)
	filterSearch := "filterSearch_example" // string | Filter by package name or CVE (when category is \"vulnerability\") (optional)
	filterDepths := []string{"FilterDepths_example"} // []string | Filter by issue depth (optional)
	filterTicketed := []string{"FilterTicketed_example"} // []string | Filter by ticketed status.  Only available to premium users. (optional)
	filterContainerLayers := []string{"FilterContainerLayers_example"} // []string | Filter by container layer (optional)
	filterType := openapiclient.getIssueStatuses_filter_type____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterTypeParameter | Filter by licensing issue type (when category is \"licensing\") or quality issue type (when category is \"quality\")  (optional)
	filterPackageManagers := []string{"Inner_example"} // []string | Filter by specific package managers (optional)
	filterCwes := []string{"Inner_example"} // []string | Filter by specific CWE identifiers (optional)
	filterProjectLabels := []string{"Inner_example"} // []string | Filter by specific project labels (optional)
	filterIdentification := []string{"FilterIdentification_example"} // []string | Filter by license identification (when category is \"licensing\") (optional)
	filterSeverity := []string{"FilterSeverity_example"} // []string | Filter by vuln severity (when category is \"vulnerability\") (optional)
	filterFoundAfter := time.Now() // time.Time | Include only issues found on after a given ISO timestamp.  Only available to premium users (optional)
	filterHasFix := []string{"FilterHasFix_example"} // []string | Filter by vuln fixability (when category is \"vulnerability\") (optional)
	filterUpgradeDistance := []string{"FilterUpgradeDistance_example"} // []string | Filter by vuln upgrade distance (when category is \"vulnerability\") (optional)
	filterExploitMaturity := []string{"FilterExploitMaturity_example"} // []string | Filter by vuln exploit maturity (when category is \"vulnerability\") (optional)
	filterIgnoreReason := []string{"FilterIgnoreReason_example"} // []string | Filter by vuln ignore reason (when category is \"vulnerability\") This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  (optional)
	filterEpss := *openapiclient.NewGetIssueDiffComparisonSummariesFilterEpssParameter() // GetIssueDiffComparisonSummariesFilterEpssParameter | Filter by epss 'score' or 'percentile'. All fields are required.  Only available to premium users. (optional)
	filterConfidence := []string{"FilterConfidence_example"} // []string | Filter issues by their binary dependency confidence level(s) (optional)
	filterIssueSource := openapiclient.getIssueStatuses_filter_issueSource____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterIssueSourceParameter | Filter by issue source. Use 'dependency' and 'snippet' to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use 'managed-dependency' and 'vendored-dependency' to filter dependency issues by whether the dependency is managed or vendored.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueDiffComparisonSummaries(context.Background()).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToRelease(scopeCompareToRelease).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterCwes(filterCwes).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterEpss(filterEpss).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueDiffComparisonSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueDiffComparisonSummaries`: GetIssueDiffComparisonSummaries200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueDiffComparisonSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueDiffComparisonSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToRelease** | **string** | The release ID to compare issues with. Only available for Release Group Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **ids** | **[]int32** | Filter by specific issue IDs | 
 **filterRevisionIds** | **[]string** | Filter by specific revision IDs | 
 **filterSearch** | **string** | Filter by package name or CVE (when category is \&quot;vulnerability\&quot;) | 
 **filterDepths** | **[]string** | Filter by issue depth | 
 **filterTicketed** | **[]string** | Filter by ticketed status.  Only available to premium users. | 
 **filterContainerLayers** | **[]string** | Filter by container layer | 
 **filterType** | [**GetIssueStatusesFilterTypeParameter**](GetIssueStatusesFilterTypeParameter.md) | Filter by licensing issue type (when category is \&quot;licensing\&quot;) or quality issue type (when category is \&quot;quality\&quot;)  | 
 **filterPackageManagers** | **[]string** | Filter by specific package managers | 
 **filterCwes** | **[]string** | Filter by specific CWE identifiers | 
 **filterProjectLabels** | **[]string** | Filter by specific project labels | 
 **filterIdentification** | **[]string** | Filter by license identification (when category is \&quot;licensing\&quot;) | 
 **filterSeverity** | **[]string** | Filter by vuln severity (when category is \&quot;vulnerability\&quot;) | 
 **filterFoundAfter** | **time.Time** | Include only issues found on after a given ISO timestamp.  Only available to premium users | 
 **filterHasFix** | **[]string** | Filter by vuln fixability (when category is \&quot;vulnerability\&quot;) | 
 **filterUpgradeDistance** | **[]string** | Filter by vuln upgrade distance (when category is \&quot;vulnerability\&quot;) | 
 **filterExploitMaturity** | **[]string** | Filter by vuln exploit maturity (when category is \&quot;vulnerability\&quot;) | 
 **filterIgnoreReason** | **[]string** | Filter by vuln ignore reason (when category is \&quot;vulnerability\&quot;) This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | 
 **filterEpss** | [**GetIssueDiffComparisonSummariesFilterEpssParameter**](GetIssueDiffComparisonSummariesFilterEpssParameter.md) | Filter by epss &#39;score&#39; or &#39;percentile&#39;. All fields are required.  Only available to premium users. | 
 **filterConfidence** | **[]string** | Filter issues by their binary dependency confidence level(s) | 
 **filterIssueSource** | [**GetIssueStatusesFilterIssueSourceParameter**](GetIssueStatusesFilterIssueSourceParameter.md) | Filter by issue source. Use &#39;dependency&#39; and &#39;snippet&#39; to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use &#39;managed-dependency&#39; and &#39;vendored-dependency&#39; to filter dependency issues by whether the dependency is managed or vendored.  | 

### Return type

[**GetIssueDiffComparisonSummaries200Response**](GetIssueDiffComparisonSummaries200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueException

> GetIssueException200Response GetIssueException(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the issue exception

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueException(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueException``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueException`: GetIssueException200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueException`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the issue exception | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueExceptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetIssueException200Response**](GetIssueException200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueExceptions

> GetIssueExceptions200Response GetIssueExceptions(ctx).FiltersCategory(filtersCategory).FiltersProjectId(filtersProjectId).FiltersReleaseGroupId(filtersReleaseGroupId).Search(search).SortBy(sortBy).OrderBy(orderBy).Page(page).Count(count).Execute()





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
	filtersCategory := "filtersCategory_example" // string | Exception category
	filtersProjectId := "filtersProjectId_example" // string | Exception project ID (optional)
	filtersReleaseGroupId := int32(56) // int32 | Exception release group ID (optional)
	search := "search_example" // string | Search term (search by created by, revision, note, or package) (optional)
	sortBy := "sortBy_example" // string | Sort by field (id, package, created by, and ignore scope) (optional)
	orderBy := "orderBy_example" // string | Sort order (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueExceptions(context.Background()).FiltersCategory(filtersCategory).FiltersProjectId(filtersProjectId).FiltersReleaseGroupId(filtersReleaseGroupId).Search(search).SortBy(sortBy).OrderBy(orderBy).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueExceptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueExceptions`: GetIssueExceptions200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueExceptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueExceptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filtersCategory** | **string** | Exception category | 
 **filtersProjectId** | **string** | Exception project ID | 
 **filtersReleaseGroupId** | **int32** | Exception release group ID | 
 **search** | **string** | Search term (search by created by, revision, note, or package) | 
 **sortBy** | **string** | Sort by field (id, package, created by, and ignore scope) | 
 **orderBy** | **string** | Sort order | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]

### Return type

[**GetIssueExceptions200Response**](GetIssueExceptions200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuePackageManagers

> GetIssuePackageManagers200Response GetIssuePackageManagers(ctx).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()





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
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	status := "status_example" // string | Issue status (optional)
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssuePackageManagers(context.Background()).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssuePackageManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuePackageManagers`: GetIssuePackageManagers200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssuePackageManagers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuePackageManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **status** | **string** | Issue status | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssuePackageManagers200Response**](GetIssuePackageManagers200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueStatuses

> GetIssueStatuses200Response GetIssueStatuses(ctx).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterIssueSource(filterIssueSource).TeamId(teamId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	ids := []int32{int32(123)} // []int32 | Filter by specific issue IDs (optional)
	filterRevisionIds := []string{"Inner_example"} // []string | Filter by specific revision IDs (optional)
	filterSearch := "filterSearch_example" // string | Filter by package name or CVE (when category is \"vulnerability\") (optional)
	filterDepths := []string{"FilterDepths_example"} // []string | Filter by issue depth (optional)
	filterTicketed := []string{"FilterTicketed_example"} // []string | Filter by ticketed status.  Only available to premium users. (optional)
	filterContainerLayers := []string{"FilterContainerLayers_example"} // []string | Filter by container layer (optional)
	filterType := openapiclient.getIssueStatuses_filter_type____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterTypeParameter | Filter by licensing issue type (when category is \"licensing\") or quality issue type (when category is \"quality\")  (optional)
	filterPackageManagers := []string{"Inner_example"} // []string | Filter by specific package managers (optional)
	filterProjectLabels := []string{"Inner_example"} // []string | Filter by specific project labels (optional)
	filterIdentification := []string{"FilterIdentification_example"} // []string | Filter by license identification (when category is \"licensing\") (optional)
	filterSeverity := []string{"FilterSeverity_example"} // []string | Filter by vuln severity (when category is \"vulnerability\") (optional)
	filterFoundAfter := time.Now() // time.Time | Include only issues found on after a given ISO timestamp.  Only available to premium users (optional)
	filterHasFix := []string{"FilterHasFix_example"} // []string | Filter by vuln fixability (when category is \"vulnerability\") (optional)
	filterUpgradeDistance := []string{"FilterUpgradeDistance_example"} // []string | Filter by vuln upgrade distance (when category is \"vulnerability\") (optional)
	filterExploitMaturity := []string{"FilterExploitMaturity_example"} // []string | Filter by vuln exploit maturity (when category is \"vulnerability\") (optional)
	filterIgnoreReason := []string{"FilterIgnoreReason_example"} // []string | Filter by vuln ignore reason (when category is \"vulnerability\") This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  (optional)
	filterLicenses := []string{"Inner_example"} // []string | Filter by issues affected by a set of license ID's (when category is \"licensing\") (optional)
	filterIssueSource := openapiclient.getIssueStatuses_filter_issueSource____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterIssueSourceParameter | Filter by issue source. Use 'dependency' and 'snippet' to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use 'managed-dependency' and 'vendored-dependency' to filter dependency issues by whether the dependency is managed or vendored.  (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueStatuses(context.Background()).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterIssueSource(filterIssueSource).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueStatuses`: GetIssueStatuses200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **ids** | **[]int32** | Filter by specific issue IDs | 
 **filterRevisionIds** | **[]string** | Filter by specific revision IDs | 
 **filterSearch** | **string** | Filter by package name or CVE (when category is \&quot;vulnerability\&quot;) | 
 **filterDepths** | **[]string** | Filter by issue depth | 
 **filterTicketed** | **[]string** | Filter by ticketed status.  Only available to premium users. | 
 **filterContainerLayers** | **[]string** | Filter by container layer | 
 **filterType** | [**GetIssueStatusesFilterTypeParameter**](GetIssueStatusesFilterTypeParameter.md) | Filter by licensing issue type (when category is \&quot;licensing\&quot;) or quality issue type (when category is \&quot;quality\&quot;)  | 
 **filterPackageManagers** | **[]string** | Filter by specific package managers | 
 **filterProjectLabels** | **[]string** | Filter by specific project labels | 
 **filterIdentification** | **[]string** | Filter by license identification (when category is \&quot;licensing\&quot;) | 
 **filterSeverity** | **[]string** | Filter by vuln severity (when category is \&quot;vulnerability\&quot;) | 
 **filterFoundAfter** | **time.Time** | Include only issues found on after a given ISO timestamp.  Only available to premium users | 
 **filterHasFix** | **[]string** | Filter by vuln fixability (when category is \&quot;vulnerability\&quot;) | 
 **filterUpgradeDistance** | **[]string** | Filter by vuln upgrade distance (when category is \&quot;vulnerability\&quot;) | 
 **filterExploitMaturity** | **[]string** | Filter by vuln exploit maturity (when category is \&quot;vulnerability\&quot;) | 
 **filterIgnoreReason** | **[]string** | Filter by vuln ignore reason (when category is \&quot;vulnerability\&quot;) This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | 
 **filterLicenses** | **[]string** | Filter by issues affected by a set of license ID&#39;s (when category is \&quot;licensing\&quot;) | 
 **filterIssueSource** | [**GetIssueStatusesFilterIssueSourceParameter**](GetIssueStatusesFilterIssueSourceParameter.md) | Filter by issue source. Use &#39;dependency&#39; and &#39;snippet&#39; to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use &#39;managed-dependency&#39; and &#39;vendored-dependency&#39; to filter dependency issues by whether the dependency is managed or vendored.  | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssueStatuses200Response**](GetIssueStatuses200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssues

> GetIssues200Response GetIssues(ctx).Category(category).ScopeType(scopeType).Csv(csv).IncludeDirectDependencyOriginPaths(includeDirectDependencyOriginPaths).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterCwes(filterCwes).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterSeveritySource(filterSeveritySource).FilterFoundBefore(filterFoundBefore).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterEpss(filterEpss).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).FilterCvssAttackVector(filterCvssAttackVector).FilterCvssAttackComplexity(filterCvssAttackComplexity).FilterCvssPrivilegesRequired(filterCvssPrivilegesRequired).Sort(sort).Page(page).Count(count).TeamId(teamId).Email(email).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	csv := true // bool | Retrieves issues as a CSV (optional)
	includeDirectDependencyOriginPaths := true // bool | Include origin paths for the direct dependency responsible for including the revision(s) affected by this issue  (optional)
	status := "status_example" // string | Issue status (optional)
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	ids := []int32{int32(123)} // []int32 | Filter by specific issue IDs (optional)
	filterRevisionIds := []string{"Inner_example"} // []string | Filter by specific revision IDs (optional)
	filterSearch := "filterSearch_example" // string | Filter by package name or CVE (when category is \"vulnerability\") (optional)
	filterDepths := []string{"FilterDepths_example"} // []string | Filter by issue depth (optional)
	filterTicketed := []string{"FilterTicketed_example"} // []string | Filter by ticketed status.  Only available to premium users. (optional)
	filterContainerLayers := []string{"FilterContainerLayers_example"} // []string | Filter by container layer (optional)
	filterType := openapiclient.getIssueStatuses_filter_type____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterTypeParameter | Filter by licensing issue type (when category is \"licensing\") or quality issue type (when category is \"quality\")  (optional)
	filterPackageManagers := []string{"Inner_example"} // []string | Filter by specific package managers (optional)
	filterCwes := []string{"Inner_example"} // []string | Filter by specific CWE identifiers (optional)
	filterProjectLabels := []string{"Inner_example"} // []string | Filter by specific project labels (optional)
	filterIdentification := []string{"FilterIdentification_example"} // []string | Filter by license identification (when category is \"licensing\") (optional)
	filterSeverity := []string{"FilterSeverity_example"} // []string | Filter by vuln severity (when category is \"vulnerability\") (optional)
	filterSeveritySource := []string{"FilterSeveritySource_example"} // []string | Filter by severity source (when category is \"vulnerability\"). Use 'standard' to filter by CVSS score, 'custom' to filter by custom risk score. When both are provided, issues matching either source are returned. Defaults to 'standard' when not provided. Custom risk score filtering is not available for global scope.  (optional)
	filterFoundBefore := time.Now() // time.Time | Include only issues found on before a given ISO timestamp.  Only available to premium users (optional)
	filterFoundAfter := time.Now() // time.Time | Include only issues found on after a given ISO timestamp.  Only available to premium users (optional)
	filterHasFix := []string{"FilterHasFix_example"} // []string | Filter by vuln fixability (when category is \"vulnerability\") (optional)
	filterUpgradeDistance := []string{"FilterUpgradeDistance_example"} // []string | Filter by vuln upgrade distance (when category is \"vulnerability\") (optional)
	filterExploitMaturity := []string{"FilterExploitMaturity_example"} // []string | Filter by vuln exploit maturity (when category is \"vulnerability\") (optional)
	filterIgnoreReason := []string{"FilterIgnoreReason_example"} // []string | Filter by vuln ignore reason (when category is \"vulnerability\") This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  (optional)
	filterEpss := *openapiclient.NewGetIssueDiffComparisonSummariesFilterEpssParameter() // GetIssueDiffComparisonSummariesFilterEpssParameter | Filter by epss 'score' or 'percentile'. All fields are required.  Only available to premium users. (optional)
	filterConfidence := []string{"FilterConfidence_example"} // []string | Filter issues by their binary dependency confidence level(s) (optional)
	filterIssueSource := openapiclient.getIssueStatuses_filter_issueSource____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterIssueSourceParameter | Filter by issue source. Use 'dependency' and 'snippet' to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use 'managed-dependency' and 'vendored-dependency' to filter dependency issues by whether the dependency is managed or vendored.  (optional)
	filterCvssAttackVector := []string{"FilterCvssAttackVector_example"} // []string | Filter by CVSS attack vector (when category is \"vulnerability\") (optional)
	filterCvssAttackComplexity := []string{"FilterCvssAttackComplexity_example"} // []string | Filter by CVSS attack complexity (when category is \"vulnerability\"). For CVSS v4, this includes the Attack Requirements (AT) metric. (optional)
	filterCvssPrivilegesRequired := []string{"FilterCvssPrivilegesRequired_example"} // []string | Filter by CVSS privileges required (when category is \"vulnerability\") (optional)
	sort := "sort_example" // string | Sort by package name, when the issue was created, or severity (when category is \"vulnerability\")  (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)
	email := true // bool | Only applies when `csv=true`. When provided, the CSV report is generated as a background task and emailed to the requesting user once ready; the response is then a `201` with the task metadata. When omitted, the CSV is streamed directly in the response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssues(context.Background()).Category(category).ScopeType(scopeType).Csv(csv).IncludeDirectDependencyOriginPaths(includeDirectDependencyOriginPaths).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterCwes(filterCwes).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterSeveritySource(filterSeveritySource).FilterFoundBefore(filterFoundBefore).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterEpss(filterEpss).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).FilterCvssAttackVector(filterCvssAttackVector).FilterCvssAttackComplexity(filterCvssAttackComplexity).FilterCvssPrivilegesRequired(filterCvssPrivilegesRequired).Sort(sort).Page(page).Count(count).TeamId(teamId).Email(email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssues`: GetIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **csv** | **bool** | Retrieves issues as a CSV | 
 **includeDirectDependencyOriginPaths** | **bool** | Include origin paths for the direct dependency responsible for including the revision(s) affected by this issue  | 
 **status** | **string** | Issue status | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **ids** | **[]int32** | Filter by specific issue IDs | 
 **filterRevisionIds** | **[]string** | Filter by specific revision IDs | 
 **filterSearch** | **string** | Filter by package name or CVE (when category is \&quot;vulnerability\&quot;) | 
 **filterDepths** | **[]string** | Filter by issue depth | 
 **filterTicketed** | **[]string** | Filter by ticketed status.  Only available to premium users. | 
 **filterContainerLayers** | **[]string** | Filter by container layer | 
 **filterType** | [**GetIssueStatusesFilterTypeParameter**](GetIssueStatusesFilterTypeParameter.md) | Filter by licensing issue type (when category is \&quot;licensing\&quot;) or quality issue type (when category is \&quot;quality\&quot;)  | 
 **filterPackageManagers** | **[]string** | Filter by specific package managers | 
 **filterCwes** | **[]string** | Filter by specific CWE identifiers | 
 **filterProjectLabels** | **[]string** | Filter by specific project labels | 
 **filterIdentification** | **[]string** | Filter by license identification (when category is \&quot;licensing\&quot;) | 
 **filterSeverity** | **[]string** | Filter by vuln severity (when category is \&quot;vulnerability\&quot;) | 
 **filterSeveritySource** | **[]string** | Filter by severity source (when category is \&quot;vulnerability\&quot;). Use &#39;standard&#39; to filter by CVSS score, &#39;custom&#39; to filter by custom risk score. When both are provided, issues matching either source are returned. Defaults to &#39;standard&#39; when not provided. Custom risk score filtering is not available for global scope.  | 
 **filterFoundBefore** | **time.Time** | Include only issues found on before a given ISO timestamp.  Only available to premium users | 
 **filterFoundAfter** | **time.Time** | Include only issues found on after a given ISO timestamp.  Only available to premium users | 
 **filterHasFix** | **[]string** | Filter by vuln fixability (when category is \&quot;vulnerability\&quot;) | 
 **filterUpgradeDistance** | **[]string** | Filter by vuln upgrade distance (when category is \&quot;vulnerability\&quot;) | 
 **filterExploitMaturity** | **[]string** | Filter by vuln exploit maturity (when category is \&quot;vulnerability\&quot;) | 
 **filterIgnoreReason** | **[]string** | Filter by vuln ignore reason (when category is \&quot;vulnerability\&quot;) This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | 
 **filterEpss** | [**GetIssueDiffComparisonSummariesFilterEpssParameter**](GetIssueDiffComparisonSummariesFilterEpssParameter.md) | Filter by epss &#39;score&#39; or &#39;percentile&#39;. All fields are required.  Only available to premium users. | 
 **filterConfidence** | **[]string** | Filter issues by their binary dependency confidence level(s) | 
 **filterIssueSource** | [**GetIssueStatusesFilterIssueSourceParameter**](GetIssueStatusesFilterIssueSourceParameter.md) | Filter by issue source. Use &#39;dependency&#39; and &#39;snippet&#39; to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use &#39;managed-dependency&#39; and &#39;vendored-dependency&#39; to filter dependency issues by whether the dependency is managed or vendored.  | 
 **filterCvssAttackVector** | **[]string** | Filter by CVSS attack vector (when category is \&quot;vulnerability\&quot;) | 
 **filterCvssAttackComplexity** | **[]string** | Filter by CVSS attack complexity (when category is \&quot;vulnerability\&quot;). For CVSS v4, this includes the Attack Requirements (AT) metric. | 
 **filterCvssPrivilegesRequired** | **[]string** | Filter by CVSS privileges required (when category is \&quot;vulnerability\&quot;) | 
 **sort** | **string** | Sort by package name, when the issue was created, or severity (when category is \&quot;vulnerability\&quot;)  | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 
 **email** | **bool** | Only applies when &#x60;csv&#x3D;true&#x60;. When provided, the CSV report is generated as a background task and emailed to the requesting user once ready; the response is then a &#x60;201&#x60; with the task metadata. When omitted, the CSV is streamed directly in the response.  | 

### Return type

[**GetIssues200Response**](GetIssues200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesByCategory

> GetIssuesByCategory200Response GetIssuesByCategory(ctx).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ReleaseGroupId(releaseGroupId).TeamId(teamId).Execute()





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
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	releaseGroupId := int32(56) // int32 | Scope the category counts to a single release group, using its latest release. When set, the counts are release-scoped, so the `teamId` filter does not apply. Requires view access to the release group; otherwise rejected with a `403`. Rejected with a `404` when the release group has no release to scope to.  (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssuesByCategory(context.Background()).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ReleaseGroupId(releaseGroupId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssuesByCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuesByCategory`: GetIssuesByCategory200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssuesByCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesByCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **releaseGroupId** | **int32** | Scope the category counts to a single release group, using its latest release. When set, the counts are release-scoped, so the &#x60;teamId&#x60; filter does not apply. Requires view access to the release group; otherwise rejected with a &#x60;403&#x60;. Rejected with a &#x60;404&#x60; when the release group has no release to scope to.  | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssuesByCategory200Response**](GetIssuesByCategory200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesByRevision

> GetIssuesByRevision200Response GetIssuesByRevision(ctx).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterIssueSource(filterIssueSource).Sort(sort).Page(page).Count(count).TeamId(teamId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	status := "status_example" // string | Issue status (optional)
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	ids := []int32{int32(123)} // []int32 | Filter by specific issue IDs (optional)
	filterRevisionIds := []string{"Inner_example"} // []string | Filter by specific revision IDs (optional)
	filterSearch := "filterSearch_example" // string | Filter by package name or CVE (when category is \"vulnerability\") (optional)
	filterDepths := []string{"FilterDepths_example"} // []string | Filter by issue depth (optional)
	filterTicketed := []string{"FilterTicketed_example"} // []string | Filter by ticketed status.  Only available to premium users. (optional)
	filterContainerLayers := []string{"FilterContainerLayers_example"} // []string | Filter by container layer (optional)
	filterType := openapiclient.getIssueStatuses_filter_type____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterTypeParameter | Filter by licensing issue type (when category is \"licensing\") or quality issue type (when category is \"quality\")  (optional)
	filterPackageManagers := []string{"Inner_example"} // []string | Filter by specific package managers (optional)
	filterProjectLabels := []string{"Inner_example"} // []string | Filter by specific project labels (optional)
	filterIdentification := []string{"FilterIdentification_example"} // []string | Filter by license identification (when category is \"licensing\") (optional)
	filterSeverity := []string{"FilterSeverity_example"} // []string | Filter by vuln severity (when category is \"vulnerability\") (optional)
	filterFoundAfter := time.Now() // time.Time | Include only issues found on after a given ISO timestamp.  Only available to premium users (optional)
	filterHasFix := []string{"FilterHasFix_example"} // []string | Filter by vuln fixability (when category is \"vulnerability\") (optional)
	filterUpgradeDistance := []string{"FilterUpgradeDistance_example"} // []string | Filter by vuln upgrade distance (when category is \"vulnerability\") (optional)
	filterExploitMaturity := []string{"FilterExploitMaturity_example"} // []string | Filter by vuln exploit maturity (when category is \"vulnerability\") (optional)
	filterIgnoreReason := []string{"FilterIgnoreReason_example"} // []string | Filter by vuln ignore reason (when category is \"vulnerability\") This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  (optional)
	filterLicenses := []string{"Inner_example"} // []string | Filter by issues affected by a set of license ID's (when category is \"licensing\") (optional)
	filterIssueSource := openapiclient.getIssueStatuses_filter_issueSource____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterIssueSourceParameter | Filter by issue source. Use 'dependency' and 'snippet' to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use 'managed-dependency' and 'vendored-dependency' to filter dependency issues by whether the dependency is managed or vendored.  (optional)
	sort := "sort_example" // string | Sort by package name, when the issue was created, or total number of issues  (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssuesByRevision(context.Background()).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterIssueSource(filterIssueSource).Sort(sort).Page(page).Count(count).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssuesByRevision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuesByRevision`: GetIssuesByRevision200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssuesByRevision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesByRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **status** | **string** | Issue status | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **ids** | **[]int32** | Filter by specific issue IDs | 
 **filterRevisionIds** | **[]string** | Filter by specific revision IDs | 
 **filterSearch** | **string** | Filter by package name or CVE (when category is \&quot;vulnerability\&quot;) | 
 **filterDepths** | **[]string** | Filter by issue depth | 
 **filterTicketed** | **[]string** | Filter by ticketed status.  Only available to premium users. | 
 **filterContainerLayers** | **[]string** | Filter by container layer | 
 **filterType** | [**GetIssueStatusesFilterTypeParameter**](GetIssueStatusesFilterTypeParameter.md) | Filter by licensing issue type (when category is \&quot;licensing\&quot;) or quality issue type (when category is \&quot;quality\&quot;)  | 
 **filterPackageManagers** | **[]string** | Filter by specific package managers | 
 **filterProjectLabels** | **[]string** | Filter by specific project labels | 
 **filterIdentification** | **[]string** | Filter by license identification (when category is \&quot;licensing\&quot;) | 
 **filterSeverity** | **[]string** | Filter by vuln severity (when category is \&quot;vulnerability\&quot;) | 
 **filterFoundAfter** | **time.Time** | Include only issues found on after a given ISO timestamp.  Only available to premium users | 
 **filterHasFix** | **[]string** | Filter by vuln fixability (when category is \&quot;vulnerability\&quot;) | 
 **filterUpgradeDistance** | **[]string** | Filter by vuln upgrade distance (when category is \&quot;vulnerability\&quot;) | 
 **filterExploitMaturity** | **[]string** | Filter by vuln exploit maturity (when category is \&quot;vulnerability\&quot;) | 
 **filterIgnoreReason** | **[]string** | Filter by vuln ignore reason (when category is \&quot;vulnerability\&quot;) This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | 
 **filterLicenses** | **[]string** | Filter by issues affected by a set of license ID&#39;s (when category is \&quot;licensing\&quot;) | 
 **filterIssueSource** | [**GetIssueStatusesFilterIssueSourceParameter**](GetIssueStatusesFilterIssueSourceParameter.md) | Filter by issue source. Use &#39;dependency&#39; and &#39;snippet&#39; to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use &#39;managed-dependency&#39; and &#39;vendored-dependency&#39; to filter dependency issues by whether the dependency is managed or vendored.  | 
 **sort** | **string** | Sort by package name, when the issue was created, or total number of issues  | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssuesByRevision200Response**](GetIssuesByRevision200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesByType

> GetIssuesByType200Response GetIssuesByType(ctx).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()





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
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssuesByType(context.Background()).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssuesByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuesByType`: GetIssuesByType200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssuesByType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetIssuesByType200Response**](GetIssuesByType200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseList

> GetLicenseList200Response GetLicenseList(ctx).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()





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
	category := "category_example" // string | Issue category (must be licensing)
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetLicenseList(context.Background()).Category(category).ScopeType(scopeType).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetLicenseList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseList`: GetLicenseList200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetLicenseList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category (must be licensing) | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 

### Return type

[**GetLicenseList200Response**](GetLicenseList200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectCSVExportIssues

> *os.File GetProjectCSVExportIssues(ctx, locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()





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
	locator := "locator_example" // string | the url-encoded locator of the project
	revisionId := "revisionId_example" // string | The revisionId of the project to get issues for. If not provided, the latest revision will be used. (optional)
	status := TODO // interface{} | The status of the issues to return. If not provided, all active issues will be returned. (optional)
	ref := "ref_example" // string | the branch or tag to get issues for. If not provided, the default branch will be used. (optional)
	refType := "refType_example" // string | Specify whether the ref is a branch or tag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetProjectCSVExportIssues(context.Background(), locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetProjectCSVExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectCSVExportIssues`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetProjectCSVExportIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectCSVExportIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** | The revisionId of the project to get issues for. If not provided, the latest revision will be used. | 
 **status** | [**interface{}**](interface{}.md) | The status of the issues to return. If not provided, all active issues will be returned. | 
 **ref** | **string** | the branch or tag to get issues for. If not provided, the default branch will be used. | 
 **refType** | **string** | Specify whether the ref is a branch or tag. | 

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


## GetProjectExportIssues

> GetProjectExportIssues200Response GetProjectExportIssues(ctx, locator).Format(format).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()





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
	locator := "locator_example" // string | the url-encoded locator of the project
	format := TODO // interface{} | report format to return (JSON or CSV). If not provided, JSON will be returned. (optional)
	revisionId := "revisionId_example" // string | The revisionId of the project to get issues for. If not provided, the latest revision will be used. (optional)
	status := TODO // interface{} | The status of the issues to return. If not provided, all active issues will be returned. (optional)
	ref := "ref_example" // string | the branch or tag to get issues for. If not provided, the default branch will be used. (optional)
	refType := "refType_example" // string | Specify whether the ref is a branch or tag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetProjectExportIssues(context.Background(), locator).Format(format).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetProjectExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectExportIssues`: GetProjectExportIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetProjectExportIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectExportIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | [**interface{}**](interface{}.md) | report format to return (JSON or CSV). If not provided, JSON will be returned. | 
 **revisionId** | **string** | The revisionId of the project to get issues for. If not provided, the latest revision will be used. | 
 **status** | [**interface{}**](interface{}.md) | The status of the issues to return. If not provided, all active issues will be returned. | 
 **ref** | **string** | the branch or tag to get issues for. If not provided, the default branch will be used. | 
 **refType** | **string** | Specify whether the ref is a branch or tag. | 

### Return type

[**GetProjectExportIssues200Response**](GetProjectExportIssues200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectJSONExportIssues

> GetProjectJSONExportIssues200Response GetProjectJSONExportIssues(ctx, locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()





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
	locator := "locator_example" // string | the url-encoded locator of the project
	revisionId := "revisionId_example" // string | The revisionId of the project to get issues for. If not provided, the latest revision will be used. (optional)
	status := TODO // interface{} | The status of the issues to return. If not provided, all active issues will be returned. (optional)
	ref := "ref_example" // string | the branch or tag to get issues for. If not provided, the default branch will be used. (optional)
	refType := "refType_example" // string | Specify whether the ref is a branch or tag. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetProjectJSONExportIssues(context.Background(), locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetProjectJSONExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectJSONExportIssues`: GetProjectJSONExportIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetProjectJSONExportIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectJSONExportIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** | The revisionId of the project to get issues for. If not provided, the latest revision will be used. | 
 **status** | [**interface{}**](interface{}.md) | The status of the issues to return. If not provided, all active issues will be returned. | 
 **ref** | **string** | the branch or tag to get issues for. If not provided, the default branch will be used. | 
 **refType** | **string** | Specify whether the ref is a branch or tag. | 

### Return type

[**GetProjectJSONExportIssues200Response**](GetProjectJSONExportIssues200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssues

> UpdateIssues200Response UpdateIssues(ctx).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterSeveritySource(filterSeveritySource).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).FilterCvssAttackVector(filterCvssAttackVector).FilterCvssAttackComplexity(filterCvssAttackComplexity).FilterCvssPrivilegesRequired(filterCvssPrivilegesRequired).TeamId(teamId).UpdateIssuesRequest(updateIssuesRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	category := "category_example" // string | Issue category
	scopeType := "scopeType_example" // string | Scope of issues to view / update
	status := "status_example" // string | Issue status (optional)
	scopeId := "scopeId_example" // string | Project or release group ID (required when scope[type] is \"project\" or \"releaseGroup\") (optional)
	scopeRevision := "scopeRevision_example" // string | Revision ID (when scope[type] is \"project\") (optional)
	scopeRevisionScanId := int32(56) // int32 | Revision scan ID (when scope[type] is \"project\") (optional)
	scopeRelease := "scopeRelease_example" // string | Release group ID (when scope[type] is \"releaseGroup\") (optional)
	scopeReleaseScanId := "scopeReleaseScanId_example" // string | Release scan ID (when scope[type] is \"releaseGroup\") (optional)
	scopeCompareToRevision := "scopeCompareToRevision_example" // string | The revision ID to compare issues with. Only available for Project Scope. (optional)
	scopeCompareToChangeStatus := "scopeCompareToChangeStatus_example" // string | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  (optional)
	ids := []int32{int32(123)} // []int32 | Filter by specific issue IDs (optional)
	filterRevisionIds := []string{"Inner_example"} // []string | Filter by specific revision IDs (optional)
	filterSearch := "filterSearch_example" // string | Filter by package name or CVE (when category is \"vulnerability\") (optional)
	filterDepths := []string{"FilterDepths_example"} // []string | Filter by issue depth (optional)
	filterTicketed := []string{"FilterTicketed_example"} // []string | Filter by ticketed status.  Only available to premium users. (optional)
	filterContainerLayers := []string{"FilterContainerLayers_example"} // []string | Filter by container layer (optional)
	filterType := openapiclient.getIssueStatuses_filter_type____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterTypeParameter | Filter by licensing issue type (when category is \"licensing\") or quality issue type (when category is \"quality\")  (optional)
	filterPackageManagers := []string{"Inner_example"} // []string | Filter by specific package managers (optional)
	filterProjectLabels := []string{"Inner_example"} // []string | Filter by specific project labels (optional)
	filterIdentification := []string{"FilterIdentification_example"} // []string | Filter by license identification (when category is \"licensing\") (optional)
	filterSeverity := []string{"FilterSeverity_example"} // []string | Filter by vuln severity (when category is \"vulnerability\") (optional)
	filterSeveritySource := []string{"FilterSeveritySource_example"} // []string | Filter by severity source (when category is \"vulnerability\"). Use 'standard' to filter by CVSS score, 'custom' to filter by custom risk score. When both are provided, issues matching either source are returned. Defaults to 'standard' when not provided. Custom risk score filtering is not available for global scope.  (optional)
	filterFoundAfter := time.Now() // time.Time | Include only issues found on after a given ISO timestamp.  Only available to premium users (optional)
	filterHasFix := []string{"FilterHasFix_example"} // []string | Filter by vuln fixability (when category is \"vulnerability\") (optional)
	filterUpgradeDistance := []string{"FilterUpgradeDistance_example"} // []string | Filter by vuln upgrade distance (when category is \"vulnerability\") (optional)
	filterExploitMaturity := []string{"FilterExploitMaturity_example"} // []string | Filter by vuln exploit maturity (when category is \"vulnerability\") (optional)
	filterIgnoreReason := []string{"FilterIgnoreReason_example"} // []string | Filter by vuln ignore reason (when category is \"vulnerability\") This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  (optional)
	filterLicenses := []string{"Inner_example"} // []string | Filter by issues affected by a set of license ID's (when category is \"licensing\") (optional)
	filterConfidence := []string{"FilterConfidence_example"} // []string | Filter issues by their binary dependency confidence level(s) (optional)
	filterIssueSource := openapiclient.getIssueStatuses_filter_issueSource____parameter{ArrayOfString: new([]string)} // GetIssueStatusesFilterIssueSourceParameter | Filter by issue source. Use 'dependency' and 'snippet' to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use 'managed-dependency' and 'vendored-dependency' to filter dependency issues by whether the dependency is managed or vendored.  (optional)
	filterCvssAttackVector := []string{"FilterCvssAttackVector_example"} // []string | Filter by CVSS attack vector (when category is \"vulnerability\") (optional)
	filterCvssAttackComplexity := []string{"FilterCvssAttackComplexity_example"} // []string | Filter by CVSS attack complexity (when category is \"vulnerability\"). For CVSS v4, this includes the Attack Requirements (AT) metric. (optional)
	filterCvssPrivilegesRequired := []string{"FilterCvssPrivilegesRequired_example"} // []string | Filter by CVSS privileges required (when category is \"vulnerability\") (optional)
	teamId := openapiclient.getIssuesByCategory_teamId_parameter{ArrayOfString: new([]string)} // GetIssuesByCategoryTeamIdParameter | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string `\"null\"` to scope to unassigned projects. Requires the `View` permission on each requested team; otherwise the request is rejected with a `403`. Ignored for organizations on the free tier.  (optional)
	updateIssuesRequest := openapiclient.updateIssues_request{UpdateIssuesRequestOneOf: openapiclient.NewUpdateIssuesRequestOneOf()} // UpdateIssuesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.UpdateIssues(context.Background()).Category(category).ScopeType(scopeType).Status(status).ScopeId(scopeId).ScopeRevision(scopeRevision).ScopeRevisionScanId(scopeRevisionScanId).ScopeRelease(scopeRelease).ScopeReleaseScanId(scopeReleaseScanId).ScopeCompareToRevision(scopeCompareToRevision).ScopeCompareToChangeStatus(scopeCompareToChangeStatus).Ids(ids).FilterRevisionIds(filterRevisionIds).FilterSearch(filterSearch).FilterDepths(filterDepths).FilterTicketed(filterTicketed).FilterContainerLayers(filterContainerLayers).FilterType(filterType).FilterPackageManagers(filterPackageManagers).FilterProjectLabels(filterProjectLabels).FilterIdentification(filterIdentification).FilterSeverity(filterSeverity).FilterSeveritySource(filterSeveritySource).FilterFoundAfter(filterFoundAfter).FilterHasFix(filterHasFix).FilterUpgradeDistance(filterUpgradeDistance).FilterExploitMaturity(filterExploitMaturity).FilterIgnoreReason(filterIgnoreReason).FilterLicenses(filterLicenses).FilterConfidence(filterConfidence).FilterIssueSource(filterIssueSource).FilterCvssAttackVector(filterCvssAttackVector).FilterCvssAttackComplexity(filterCvssAttackComplexity).FilterCvssPrivilegesRequired(filterCvssPrivilegesRequired).TeamId(teamId).UpdateIssuesRequest(updateIssuesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.UpdateIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssues`: UpdateIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.UpdateIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Issue category | 
 **scopeType** | **string** | Scope of issues to view / update | 
 **status** | **string** | Issue status | 
 **scopeId** | **string** | Project or release group ID (required when scope[type] is \&quot;project\&quot; or \&quot;releaseGroup\&quot;) | 
 **scopeRevision** | **string** | Revision ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRevisionScanId** | **int32** | Revision scan ID (when scope[type] is \&quot;project\&quot;) | 
 **scopeRelease** | **string** | Release group ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeReleaseScanId** | **string** | Release scan ID (when scope[type] is \&quot;releaseGroup\&quot;) | 
 **scopeCompareToRevision** | **string** | The revision ID to compare issues with. Only available for Project Scope. | 
 **scopeCompareToChangeStatus** | **string** | The status of issues to fetch when comparing issues. - New issues are present in the current revision but not in the comparison revision. - Remediated issues are present in the comparison revision but not in the current revision. - Unchanged issues are present in both revisions. Only available for Project Scope.  | 
 **ids** | **[]int32** | Filter by specific issue IDs | 
 **filterRevisionIds** | **[]string** | Filter by specific revision IDs | 
 **filterSearch** | **string** | Filter by package name or CVE (when category is \&quot;vulnerability\&quot;) | 
 **filterDepths** | **[]string** | Filter by issue depth | 
 **filterTicketed** | **[]string** | Filter by ticketed status.  Only available to premium users. | 
 **filterContainerLayers** | **[]string** | Filter by container layer | 
 **filterType** | [**GetIssueStatusesFilterTypeParameter**](GetIssueStatusesFilterTypeParameter.md) | Filter by licensing issue type (when category is \&quot;licensing\&quot;) or quality issue type (when category is \&quot;quality\&quot;)  | 
 **filterPackageManagers** | **[]string** | Filter by specific package managers | 
 **filterProjectLabels** | **[]string** | Filter by specific project labels | 
 **filterIdentification** | **[]string** | Filter by license identification (when category is \&quot;licensing\&quot;) | 
 **filterSeverity** | **[]string** | Filter by vuln severity (when category is \&quot;vulnerability\&quot;) | 
 **filterSeveritySource** | **[]string** | Filter by severity source (when category is \&quot;vulnerability\&quot;). Use &#39;standard&#39; to filter by CVSS score, &#39;custom&#39; to filter by custom risk score. When both are provided, issues matching either source are returned. Defaults to &#39;standard&#39; when not provided. Custom risk score filtering is not available for global scope.  | 
 **filterFoundAfter** | **time.Time** | Include only issues found on after a given ISO timestamp.  Only available to premium users | 
 **filterHasFix** | **[]string** | Filter by vuln fixability (when category is \&quot;vulnerability\&quot;) | 
 **filterUpgradeDistance** | **[]string** | Filter by vuln upgrade distance (when category is \&quot;vulnerability\&quot;) | 
 **filterExploitMaturity** | **[]string** | Filter by vuln exploit maturity (when category is \&quot;vulnerability\&quot;) | 
 **filterIgnoreReason** | **[]string** | Filter by vuln ignore reason (when category is \&quot;vulnerability\&quot;) This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | 
 **filterLicenses** | **[]string** | Filter by issues affected by a set of license ID&#39;s (when category is \&quot;licensing\&quot;) | 
 **filterConfidence** | **[]string** | Filter issues by their binary dependency confidence level(s) | 
 **filterIssueSource** | [**GetIssueStatusesFilterIssueSourceParameter**](GetIssueStatusesFilterIssueSourceParameter.md) | Filter by issue source. Use &#39;dependency&#39; and &#39;snippet&#39; to filter by whether the issue comes from a dependency or a code snippet. When the vendored dependency detection feature is enabled, use &#39;managed-dependency&#39; and &#39;vendored-dependency&#39; to filter dependency issues by whether the dependency is managed or vendored.  | 
 **filterCvssAttackVector** | **[]string** | Filter by CVSS attack vector (when category is \&quot;vulnerability\&quot;) | 
 **filterCvssAttackComplexity** | **[]string** | Filter by CVSS attack complexity (when category is \&quot;vulnerability\&quot;). For CVSS v4, this includes the Attack Requirements (AT) metric. | 
 **filterCvssPrivilegesRequired** | **[]string** | Filter by CVSS privileges required (when category is \&quot;vulnerability\&quot;) | 
 **teamId** | [**GetIssuesByCategoryTeamIdParameter**](GetIssuesByCategoryTeamIdParameter.md) | Filter issues by one or more team IDs. Accepts a single team ID, an array of team IDs, or the literal string &#x60;\&quot;null\&quot;&#x60; to scope to unassigned projects. Requires the &#x60;View&#x60; permission on each requested team; otherwise the request is rejected with a &#x60;403&#x60;. Ignored for organizations on the free tier.  | 
 **updateIssuesRequest** | [**UpdateIssuesRequest**](UpdateIssuesRequest.md) |  | 

### Return type

[**UpdateIssues200Response**](UpdateIssues200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

