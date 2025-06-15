# \DependenciesAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectDependencies**](DependenciesAPI.md#GetProjectDependencies) | **Get** /v2/revisions/{locator}/dependencies | 
[**GetProjectDependencyPackageManagers**](DependenciesAPI.md#GetProjectDependencyPackageManagers) | **Get** /v2/revisions/{locator}/dependencies/package-managers | 
[**GetReleaseGroupDependencies**](DependenciesAPI.md#GetReleaseGroupDependencies) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies | 
[**GetReleaseGroupDependencyPackageManagers**](DependenciesAPI.md#GetReleaseGroupDependencyPackageManagers) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies/package-managers | 
[**GetReleaseGroupDependencyRootProjects**](DependenciesAPI.md#GetReleaseGroupDependencyRootProjects) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/root-projects | 



## GetProjectDependencies

> GetProjectDependencies200Response GetProjectDependencies(ctx, locator).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Page(page).Count(count).Execute()





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
	locators := []string{"Inner_example"} // []string | Filter dependencies by locators (exact match) (optional)
	title := "title_example" // string | Filter dependencies by title (optional)
	status := []string{"Status_example"} // []string | Filter dependencies by status (optional)
	depth := []string{"Depth_example"} // []string | Filter dependencies by depth (optional)
	layerDepth := []string{"LayerDepth_example"} // []string | Filter dependencies by depth (for container projects) (optional)
	hasIssues := []string{"HasIssues_example"} // []string | Filter dependencies by the presence of issues (optional)
	licenses := []string{"Inner_example"} // []string | Filter dependencies by licenses (optional)
	fetchers := []string{"Inner_example"} // []string | Filter dependencies by package manager (optional)
	showIgnored := true // bool | Includes ignored dependencies (optional)
	confidence := []string{"Confidence_example"} // []string | Filter dependencies by confidence (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetProjectDependencies(context.Background(), locator).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetProjectDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectDependencies`: GetProjectDependencies200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetProjectDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locators** | **[]string** | Filter dependencies by locators (exact match) | 
 **title** | **string** | Filter dependencies by title | 
 **status** | **[]string** | Filter dependencies by status | 
 **depth** | **[]string** | Filter dependencies by depth | 
 **layerDepth** | **[]string** | Filter dependencies by depth (for container projects) | 
 **hasIssues** | **[]string** | Filter dependencies by the presence of issues | 
 **licenses** | **[]string** | Filter dependencies by licenses | 
 **fetchers** | **[]string** | Filter dependencies by package manager | 
 **showIgnored** | **bool** | Includes ignored dependencies | 
 **confidence** | **[]string** | Filter dependencies by confidence | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 50]

### Return type

[**GetProjectDependencies200Response**](GetProjectDependencies200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDependencyPackageManagers

> GetIssuePackageManagers200Response GetProjectDependencyPackageManagers(ctx, locator).Execute()





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
	resp, r, err := apiClient.DependenciesAPI.GetProjectDependencyPackageManagers(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetProjectDependencyPackageManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectDependencyPackageManagers`: GetIssuePackageManagers200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetProjectDependencyPackageManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDependencyPackageManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetReleaseGroupDependencies

> GetProjectDependencies200Response GetReleaseGroupDependencies(ctx, projectGroupId, projectGroupReleaseId).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).RootProjects(rootProjects).Page(page).Count(count).Execute()





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
	projectGroupId := int32(56) // int32 | 
	projectGroupReleaseId := int32(56) // int32 | The ID of the release
	locators := []string{"Inner_example"} // []string | Filter dependencies by locators (exact match) (optional)
	title := "title_example" // string | Filter dependencies by title (optional)
	status := []string{"Status_example"} // []string | Filter dependencies by status (optional)
	depth := []string{"Depth_example"} // []string | Filter dependencies by depth (optional)
	layerDepth := []string{"LayerDepth_example"} // []string | Filter dependencies by depth (for container projects) (optional)
	hasIssues := []string{"HasIssues_example"} // []string | Filter dependencies by the presence of issues (optional)
	licenses := []string{"Inner_example"} // []string | Filter dependencies by licenses (optional)
	fetchers := []string{"Inner_example"} // []string | Filter dependencies by package manager (optional)
	showIgnored := true // bool | Includes ignored dependencies (optional)
	rootProjects := []string{"Inner_example"} // []string | Filter release group dependencies by root projects (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependencies(context.Background(), projectGroupId, projectGroupReleaseId).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).RootProjects(rootProjects).Page(page).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetReleaseGroupDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupDependencies`: GetProjectDependencies200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetReleaseGroupDependencies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupDependenciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locators** | **[]string** | Filter dependencies by locators (exact match) | 
 **title** | **string** | Filter dependencies by title | 
 **status** | **[]string** | Filter dependencies by status | 
 **depth** | **[]string** | Filter dependencies by depth | 
 **layerDepth** | **[]string** | Filter dependencies by depth (for container projects) | 
 **hasIssues** | **[]string** | Filter dependencies by the presence of issues | 
 **licenses** | **[]string** | Filter dependencies by licenses | 
 **fetchers** | **[]string** | Filter dependencies by package manager | 
 **showIgnored** | **bool** | Includes ignored dependencies | 
 **rootProjects** | **[]string** | Filter release group dependencies by root projects | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 50]

### Return type

[**GetProjectDependencies200Response**](GetProjectDependencies200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupDependencyPackageManagers

> GetIssuePackageManagers200Response GetReleaseGroupDependencyPackageManagers(ctx, projectGroupId, projectGroupReleaseId).Execute()





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
	projectGroupId := int32(56) // int32 | 
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependencyPackageManagers(context.Background(), projectGroupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetReleaseGroupDependencyPackageManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupDependencyPackageManagers`: GetIssuePackageManagers200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetReleaseGroupDependencyPackageManagers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupDependencyPackageManagersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetReleaseGroupDependencyRootProjects

> GetReleaseGroupDependencyRootProjects200Response GetReleaseGroupDependencyRootProjects(ctx, projectGroupId, projectGroupReleaseId).Execute()





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
	projectGroupId := int32(56) // int32 | 
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependencyRootProjects(context.Background(), projectGroupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetReleaseGroupDependencyRootProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupDependencyRootProjects`: GetReleaseGroupDependencyRootProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetReleaseGroupDependencyRootProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** |  | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupDependencyRootProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetReleaseGroupDependencyRootProjects200Response**](GetReleaseGroupDependencyRootProjects200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

