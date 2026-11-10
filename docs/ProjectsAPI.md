# \ProjectsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProject**](ProjectsAPI.md#DeleteProject) | **Delete** /projects/{locator} | Delete a project.
[**DeleteProjectGenerateAttributionSlug**](ProjectsAPI.md#DeleteProjectGenerateAttributionSlug) | **Delete** /projects/{locator}/generate_attribution_slug | 
[**DeleteProjects**](ProjectsAPI.md#DeleteProjects) | **Delete** /v2/projects | 
[**DeleteReleaseGroups**](ProjectsAPI.md#DeleteReleaseGroups) | **Delete** /v2/release-groups | 
[**GenerateProjectGenerateAttributionSlug**](ProjectsAPI.md#GenerateProjectGenerateAttributionSlug) | **Put** /projects/{locator}/generate_attribution_slug | 
[**GetProject**](ProjectsAPI.md#GetProject) | **Get** /projects/{locator} | Get a single project.
[**GetProjectCSVExportIssues**](ProjectsAPI.md#GetProjectCSVExportIssues) | **Get** /projects/{locator}/export-issues/csv | 
[**GetProjectExportIssues**](ProjectsAPI.md#GetProjectExportIssues) | **Get** /projects/{locator}/export-issues | 
[**GetProjectJSONExportIssues**](ProjectsAPI.md#GetProjectJSONExportIssues) | **Get** /projects/{locator}/export-issues/json | 
[**GetProjectLastPublished**](ProjectsAPI.md#GetProjectLastPublished) | **Get** /projects/{locator}/last-published | 
[**GetProjectRevisions**](ProjectsAPI.md#GetProjectRevisions) | **Get** /projects/{locator}/revisions | 
[**GetProjects**](ProjectsAPI.md#GetProjects) | **Get** /v2/projects | 
[**GetProjectsByPost**](ProjectsAPI.md#GetProjectsByPost) | **Post** /v2/projects | 
[**GetProjectsSummary**](ProjectsAPI.md#GetProjectsSummary) | **Get** /v2/projects/summary | 
[**GetReleaseGroups**](ProjectsAPI.md#GetReleaseGroups) | **Get** /v2/release-groups | 
[**ListReleaseGroupsForProject**](ProjectsAPI.md#ListReleaseGroupsForProject) | **Get** /v2/projects/{locator}/release-groups | 
[**UpdateProject**](ProjectsAPI.md#UpdateProject) | **Put** /projects/{locator} | Update a project&#39;s settings and configuration.
[**UpdateProjectsLabels**](ProjectsAPI.md#UpdateProjectsLabels) | **Put** /v2/projects/labels | 
[**UpdateProjectsPolicies**](ProjectsAPI.md#UpdateProjectsPolicies) | **Put** /v2/projects/policy | 
[**UpdateReleaseGroupsPolicies**](ProjectsAPI.md#UpdateReleaseGroupsPolicies) | **Put** /v2/release-groups/policy | 



## DeleteProject

> DeleteProject(ctx, locator).Execute()

Delete a project.



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
	locator := "git%2Bgithub.com%2Ffossas%2Ffossa-cli" // string | The URL-encoded locator of the project (e.g., \"git+github.com/owner/repo\")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the project (e.g., \&quot;git+github.com/owner/repo\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


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
	r, err := apiClient.ProjectsAPI.DeleteProjectGenerateAttributionSlug(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjectGenerateAttributionSlug``: %v\n", err)
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


## DeleteProjects

> DeleteProjects(ctx).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Locators(locators).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()





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
	title := "title_example" // string | Filter by project name. (optional)
	type_ := []string{"Type_example"} // []string | Filter by project type. (optional)
	isPublic := true // bool | Filter by project being public or private. (optional)
	labels := []string{"Inner_example"} // []string | Filter by project labels. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	lastRevisionWithin := int32(56) // int32 | Filter by last revision analyzed within N days. (optional)
	locators := openapiclient.deleteProjects_locators_parameter{ArrayOfString: new([]string)} // DeleteProjectsLocatorsParameter | The list of locators for the projects to delete. If \"all\" is provided, then all projects that meet the provided filters will be deleted.  (optional)
	url := "url_example" // string | Filter by a project's URL. (optional)
	includeSharedProjects := true // bool | Include shared projects. (optional)
	onlyIncludeSharedProjects := true // bool | Only show projects that have been shared with other organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteProjects(context.Background()).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Locators(locators).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Filter by project name. | 
 **type_** | **[]string** | Filter by project type. | 
 **isPublic** | **bool** | Filter by project being public or private. | 
 **labels** | **[]string** | Filter by project labels. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **lastRevisionWithin** | **int32** | Filter by last revision analyzed within N days. | 
 **locators** | [**DeleteProjectsLocatorsParameter**](DeleteProjectsLocatorsParameter.md) | The list of locators for the projects to delete. If \&quot;all\&quot; is provided, then all projects that meet the provided filters will be deleted.  | 
 **url** | **string** | Filter by a project&#39;s URL. | 
 **includeSharedProjects** | **bool** | Include shared projects. | 
 **onlyIncludeSharedProjects** | **bool** | Only show projects that have been shared with other organizations. | 

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


## DeleteReleaseGroups

> DeleteReleaseGroups(ctx).Title(title).TeamId(teamId).LatestScan(latestScan).Ids(ids).Execute()





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
	title := "title_example" // string | Filter by release group name. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned release groups. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	ids := openapiclient.deleteReleaseGroups_ids_parameter{ArrayOfInt32: new([]int32)} // DeleteReleaseGroupsIdsParameter | The list of ids for the release groups to delete. If \"all\" is provided, then all release groups that meet the provided filters will be deleted.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.DeleteReleaseGroups(context.Background()).Title(title).TeamId(teamId).LatestScan(latestScan).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.DeleteReleaseGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Filter by release group name. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned release groups. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **ids** | [**DeleteReleaseGroupsIdsParameter**](DeleteReleaseGroupsIdsParameter.md) | The list of ids for the release groups to delete. If \&quot;all\&quot; is provided, then all release groups that meet the provided filters will be deleted.  | 

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


## GenerateProjectGenerateAttributionSlug

> string GenerateProjectGenerateAttributionSlug(ctx, locator).Execute()





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
	resp, r, err := apiClient.ProjectsAPI.GenerateProjectGenerateAttributionSlug(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GenerateProjectGenerateAttributionSlug``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateProjectGenerateAttributionSlug`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GenerateProjectGenerateAttributionSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateProjectGenerateAttributionSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> map[string]interface{} GetProject(ctx, locator).Ref(ref).RefType(refType).Execute()

Get a single project.



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
	locator := "git%2Bgithub.com%2Ffossas%2Ffossa-cli" // string | The URL-encoded locator of the project (e.g., \"git+github.com/owner/repo\")
	ref := "ref_example" // string | The branch or tag to use when computing the head/latest/last-analyzed revision metadata. Defaults to the project's default branch. (optional)
	refType := "refType_example" // string | Whether `ref` refers to a branch or a tag. Defaults to `branch`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProject(context.Background(), locator).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the project (e.g., \&quot;git+github.com/owner/repo\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ref** | **string** | The branch or tag to use when computing the head/latest/last-analyzed revision metadata. Defaults to the project&#39;s default branch. | 
 **refType** | **string** | Whether &#x60;ref&#x60; refers to a branch or a tag. Defaults to &#x60;branch&#x60;. | 

### Return type

**map[string]interface{}**

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
	resp, r, err := apiClient.ProjectsAPI.GetProjectCSVExportIssues(context.Background(), locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectCSVExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectCSVExportIssues`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectCSVExportIssues`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.GetProjectExportIssues(context.Background(), locator).Format(format).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectExportIssues`: GetProjectExportIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectExportIssues`: %v\n", resp)
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
	resp, r, err := apiClient.ProjectsAPI.GetProjectJSONExportIssues(context.Background(), locator).RevisionId(revisionId).Status(status).Ref(ref).RefType(refType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectJSONExportIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectJSONExportIssues`: GetProjectJSONExportIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectJSONExportIssues`: %v\n", resp)
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


## GetProjectLastPublished

> string GetProjectLastPublished(ctx, locator).Execute()





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
	resp, r, err := apiClient.ProjectsAPI.GetProjectLastPublished(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectLastPublished``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectLastPublished`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectLastPublished`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectLastPublishedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRevisions

> GetProjectRevisions200Response GetProjectRevisions(ctx, locator).Offset(offset).Count(count).Resolved(resolved).Refs(refs).RefsType(refsType).Source(source).IsMinimal(isMinimal).Locator2(locator2).Execute()





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
	locator := "locator_example" // string | the url-encoded locator of the project - matched by substring
	offset := float32(8.14) // float32 | the number of revisions to skip per ref (branch or tag) for pagination. (optional)
	count := float32(8.14) // float32 | the number of revisions to return per ref (branch or tag), max 1000. (optional)
	resolved := true // bool | If true, we will only return Revisions that have been successfully analyzed by FOSSA (resolved) (optional)
	refs := []string{"Inner_example"} // []string | the list of branches or tags being requested (optional)
	refsType := "refsType_example" // string | Specify whether the list should be tags OR branches (optional)
	source := "source_example" // string | Filter the revisions by source (optional)
	isMinimal := true // bool | If `true`, each Revision in the response is reduced to only its `locator` and `message` fields. Useful for limiting response size.  (optional)
	locator2 := "locator_example" // string | Filter the revisions to those whose locator contains this value (case-insensitive substring match). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectRevisions(context.Background(), locator).Offset(offset).Count(count).Resolved(resolved).Refs(refs).RefsType(refsType).Source(source).IsMinimal(isMinimal).Locator2(locator2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRevisions`: GetProjectRevisions200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the project - matched by substring | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **float32** | the number of revisions to skip per ref (branch or tag) for pagination. | 
 **count** | **float32** | the number of revisions to return per ref (branch or tag), max 1000. | 
 **resolved** | **bool** | If true, we will only return Revisions that have been successfully analyzed by FOSSA (resolved) | 
 **refs** | **[]string** | the list of branches or tags being requested | 
 **refsType** | **string** | Specify whether the list should be tags OR branches | 
 **source** | **string** | Filter the revisions by source | 
 **isMinimal** | **bool** | If &#x60;true&#x60;, each Revision in the response is reduced to only its &#x60;locator&#x60; and &#x60;message&#x60; fields. Useful for limiting response size.  | 
 **locator2** | **string** | Filter the revisions to those whose locator contains this value (case-insensitive substring match). | 

### Return type

[**GetProjectRevisions200Response**](GetProjectRevisions200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> GetProjects200Response GetProjects(ctx).Sort(sort).Page(page).Count(count).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Locators(locators).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Url(url).Inventory(inventory).Execute()





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
	sort := "sort_example" // string | The category to order the results by and sort direction. (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)
	title := "title_example" // string | Filter by project name. (optional)
	type_ := []string{"Type_example"} // []string | Filter by project type. (optional)
	isPublic := true // bool | Filter by project being public or private. (optional)
	labels := []string{"Inner_example"} // []string | Filter by project labels. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	lastRevisionWithin := int32(56) // int32 | Filter by last revision analyzed within N days. (optional)
	locators := []string{"Inner_example"} // []string | Filter by project locators (exact match). (optional)
	includeSharedProjects := true // bool | Include shared projects. (optional)
	onlyIncludeSharedProjects := true // bool | Only show projects that have been shared with other organizations. (optional)
	url := "url_example" // string | Filter by a project's URL. (optional)
	inventory := []string{"Inventory_example"} // []string | Filter by additional inventory types. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjects(context.Background()).Sort(sort).Page(page).Count(count).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Locators(locators).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Url(url).Inventory(inventory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjects`: GetProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The category to order the results by and sort direction. | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]
 **title** | **string** | Filter by project name. | 
 **type_** | **[]string** | Filter by project type. | 
 **isPublic** | **bool** | Filter by project being public or private. | 
 **labels** | **[]string** | Filter by project labels. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **lastRevisionWithin** | **int32** | Filter by last revision analyzed within N days. | 
 **locators** | **[]string** | Filter by project locators (exact match). | 
 **includeSharedProjects** | **bool** | Include shared projects. | 
 **onlyIncludeSharedProjects** | **bool** | Only show projects that have been shared with other organizations. | 
 **url** | **string** | Filter by a project&#39;s URL. | 
 **inventory** | **[]string** | Filter by additional inventory types. | 

### Return type

[**GetProjects200Response**](GetProjects200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsByPost

> GetProjects200Response GetProjectsByPost(ctx).GetProjectsByPostRequest(getProjectsByPostRequest).Execute()





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
	getProjectsByPostRequest := *openapiclient.NewGetProjectsByPostRequest() // GetProjectsByPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsByPost(context.Background()).GetProjectsByPostRequest(getProjectsByPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsByPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsByPost`: GetProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getProjectsByPostRequest** | [**GetProjectsByPostRequest**](GetProjectsByPostRequest.md) |  | 

### Return type

[**GetProjects200Response**](GetProjects200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsSummary

> GetProjectsSummary200Response GetProjectsSummary(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectsSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsSummary`: GetProjectsSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectsSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsSummaryRequest struct via the builder pattern


### Return type

[**GetProjectsSummary200Response**](GetProjectsSummary200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroups

> GetReleaseGroups200Response GetReleaseGroups(ctx).Sort(sort).Page(page).Count(count).Title(title).TeamId(teamId).LatestScan(latestScan).IsPublic(isPublic).Execute()





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
	sort := "sort_example" // string | The category to order the results by and sort direction. (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)
	title := "title_example" // string | Filter by release group name. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned release groups. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	isPublic := true // bool | Filter by whether the release group is public on the portal. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetReleaseGroups(context.Background()).Sort(sort).Page(page).Count(count).Title(title).TeamId(teamId).LatestScan(latestScan).IsPublic(isPublic).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetReleaseGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroups`: GetReleaseGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetReleaseGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | The category to order the results by and sort direction. | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 10]
 **title** | **string** | Filter by release group name. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned release groups. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **isPublic** | **bool** | Filter by whether the release group is public on the portal. | 

### Return type

[**GetReleaseGroups200Response**](GetReleaseGroups200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReleaseGroupsForProject

> ListReleaseGroupsForProject200Response ListReleaseGroupsForProject(ctx, locator).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ListReleaseGroupsForProject(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ListReleaseGroupsForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReleaseGroupsForProject`: ListReleaseGroupsForProject200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ListReleaseGroupsForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | the url-encoded locator of the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiListReleaseGroupsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListReleaseGroupsForProject200Response**](ListReleaseGroupsForProject200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> map[string]interface{} UpdateProject(ctx, locator).UpdateProjectRequest(updateProjectRequest).Execute()

Update a project's settings and configuration.



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
	locator := "git%2Bgithub.com%2Ffossas%2Ffossa-cli" // string | The URL-encoded locator of the project (e.g., \"git+github.com/owner/repo\")
	updateProjectRequest := *openapiclient.NewUpdateProjectRequest() // UpdateProjectRequest | Project fields to update. All fields are optional. Only the fields provided will be updated; omitted fields remain unchanged. Note: The endpoint filters out any fields not in the allowed list automatically.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProject(context.Background(), locator).UpdateProjectRequest(updateProjectRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProject`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The URL-encoded locator of the project (e.g., \&quot;git+github.com/owner/repo\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProjectRequest** | [**UpdateProjectRequest**](UpdateProjectRequest.md) | Project fields to update. All fields are optional. Only the fields provided will be updated; omitted fields remain unchanged. Note: The endpoint filters out any fields not in the allowed list automatically.  | 

### Return type

**map[string]interface{}**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectsLabels

> map[string]*interface{} UpdateProjectsLabels(ctx).LabelId(labelId).Locators(locators).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()





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
	labelId := float32(8.14) // float32 | The ID of the label you want to apply to projects.
	locators := openapiclient.deleteProjects_locators_parameter{ArrayOfString: new([]string)} // DeleteProjectsLocatorsParameter | The list of locators for the projects to update. If \"all\" is provided, then all projects that meet the provided filters will have the label applied. 
	title := "title_example" // string | Filter by project name. (optional)
	type_ := []string{"Type_example"} // []string | Filter by project type. (optional)
	isPublic := true // bool | Filter by project being public or private. (optional)
	labels := []string{"Inner_example"} // []string | Filter by project labels. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	lastRevisionWithin := int32(56) // int32 | Filter by last revision analyzed within N days. (optional)
	url := "url_example" // string | Filter by a project's URL. (optional)
	includeSharedProjects := true // bool | Include shared projects. (optional)
	onlyIncludeSharedProjects := true // bool | Only show projects that have been shared with other organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProjectsLabels(context.Background()).LabelId(labelId).Locators(locators).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectsLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectsLabels`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProjectsLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectsLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelId** | **float32** | The ID of the label you want to apply to projects. | 
 **locators** | [**DeleteProjectsLocatorsParameter**](DeleteProjectsLocatorsParameter.md) | The list of locators for the projects to update. If \&quot;all\&quot; is provided, then all projects that meet the provided filters will have the label applied.  | 
 **title** | **string** | Filter by project name. | 
 **type_** | **[]string** | Filter by project type. | 
 **isPublic** | **bool** | Filter by project being public or private. | 
 **labels** | **[]string** | Filter by project labels. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **lastRevisionWithin** | **int32** | Filter by last revision analyzed within N days. | 
 **url** | **string** | Filter by a project&#39;s URL. | 
 **includeSharedProjects** | **bool** | Include shared projects. | 
 **onlyIncludeSharedProjects** | **bool** | Only show projects that have been shared with other organizations. | 

### Return type

**map[string]*interface{}**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectsPolicies

> map[string]*interface{} UpdateProjectsPolicies(ctx).PolicyId(policyId).Locators(locators).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()





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
	policyId := float32(8.14) // float32 | The ID of the policy you want to apply to projects.
	locators := openapiclient.deleteProjects_locators_parameter{ArrayOfString: new([]string)} // DeleteProjectsLocatorsParameter | The list of locators for the projects to update. If \"all\" is provided, then all projects that meet the provided filters will have the policy applied. 
	title := "title_example" // string | Filter by project name. (optional)
	type_ := []string{"Type_example"} // []string | Filter by project type. (optional)
	isPublic := true // bool | Filter by project being public or private. (optional)
	labels := []string{"Inner_example"} // []string | Filter by project labels. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)
	lastRevisionWithin := int32(56) // int32 | Filter by last revision analyzed within N days. (optional)
	url := "url_example" // string | Filter by a project's URL. (optional)
	includeSharedProjects := true // bool | Include shared projects. (optional)
	onlyIncludeSharedProjects := true // bool | Only show projects that have been shared with other organizations. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.UpdateProjectsPolicies(context.Background()).PolicyId(policyId).Locators(locators).Title(title).Type_(type_).IsPublic(isPublic).Labels(labels).TeamId(teamId).LatestScan(latestScan).LastRevisionWithin(lastRevisionWithin).Url(url).IncludeSharedProjects(includeSharedProjects).OnlyIncludeSharedProjects(onlyIncludeSharedProjects).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateProjectsPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectsPolicies`: map[string]*interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.UpdateProjectsPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectsPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **float32** | The ID of the policy you want to apply to projects. | 
 **locators** | [**DeleteProjectsLocatorsParameter**](DeleteProjectsLocatorsParameter.md) | The list of locators for the projects to update. If \&quot;all\&quot; is provided, then all projects that meet the provided filters will have the policy applied.  | 
 **title** | **string** | Filter by project name. | 
 **type_** | **[]string** | Filter by project type. | 
 **isPublic** | **bool** | Filter by project being public or private. | 
 **labels** | **[]string** | Filter by project labels. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 
 **lastRevisionWithin** | **int32** | Filter by last revision analyzed within N days. | 
 **url** | **string** | Filter by a project&#39;s URL. | 
 **includeSharedProjects** | **bool** | Include shared projects. | 
 **onlyIncludeSharedProjects** | **bool** | Only show projects that have been shared with other organizations. | 

### Return type

**map[string]*interface{}**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseGroupsPolicies

> UpdateReleaseGroupsPolicies(ctx).PolicyId(policyId).Ids(ids).Title(title).TeamId(teamId).LatestScan(latestScan).Execute()





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
	policyId := float32(8.14) // float32 | The ID of the policy you want to apply to projects.
	ids := openapiclient.updateReleaseGroupsPolicies_ids_parameter{ArrayOfFloat32: new([]float32)} // UpdateReleaseGroupsPoliciesIdsParameter | The list of IDs for the release groups to update. If \"all\" is provided, then all release groups that meet the provided filters will have the policy applied. If omitted, the release groups are selected by the supplied filter params (`title`, `latestScan`, `teamId`); at least one filter must then be provided.  (optional)
	title := "title_example" // string | Filter by release group name. (optional)
	teamId := []openapiclient.GetIssueCountsTeamIdParameterInner{openapiclient.getIssueCounts_teamId_parameter_inner{Float32: new(float32)}} // []GetIssueCountsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned release groups. (optional)
	latestScan := int32(56) // int32 | Filter by last policy scan within N days. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.UpdateReleaseGroupsPolicies(context.Background()).PolicyId(policyId).Ids(ids).Title(title).TeamId(teamId).LatestScan(latestScan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.UpdateReleaseGroupsPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseGroupsPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **float32** | The ID of the policy you want to apply to projects. | 
 **ids** | [**UpdateReleaseGroupsPoliciesIdsParameter**](UpdateReleaseGroupsPoliciesIdsParameter.md) | The list of IDs for the release groups to update. If \&quot;all\&quot; is provided, then all release groups that meet the provided filters will have the policy applied. If omitted, the release groups are selected by the supplied filter params (&#x60;title&#x60;, &#x60;latestScan&#x60;, &#x60;teamId&#x60;); at least one filter must then be provided.  | 
 **title** | **string** | Filter by release group name. | 
 **teamId** | [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned release groups. | 
 **latestScan** | **int32** | Filter by last policy scan within N days. | 

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

