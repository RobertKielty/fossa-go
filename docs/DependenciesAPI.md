# \DependenciesAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomLicenses**](DependenciesAPI.md#GetCustomLicenses) | **Get** /v2/dependencies/custom-licenses | 
[**GetGlobalDependency**](DependenciesAPI.md#GetGlobalDependency) | **Get** /v2/dependencies/{locator} | 
[**GetProjectDependencies**](DependenciesAPI.md#GetProjectDependencies) | **Get** /v2/revisions/{locator}/dependencies | 
[**GetProjectDependency**](DependenciesAPI.md#GetProjectDependency) | **Get** /v2/revisions/{locator}/dependencies/{dependencyRevisionLocator} | 
[**GetProjectDependencyCount**](DependenciesAPI.md#GetProjectDependencyCount) | **Get** /v2/revisions/{locator}/dependencies/count | 
[**GetProjectDependencyPackageManagers**](DependenciesAPI.md#GetProjectDependencyPackageManagers) | **Get** /v2/revisions/{locator}/dependencies/package-managers | 
[**GetReleaseGroupDependencies**](DependenciesAPI.md#GetReleaseGroupDependencies) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies | 
[**GetReleaseGroupDependency**](DependenciesAPI.md#GetReleaseGroupDependency) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies/{dependencyRevisionLocator} | 
[**GetReleaseGroupDependencyCount**](DependenciesAPI.md#GetReleaseGroupDependencyCount) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies/count | 
[**GetReleaseGroupDependencyPackageManagers**](DependenciesAPI.md#GetReleaseGroupDependencyPackageManagers) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/dependencies/package-managers | 
[**GetReleaseGroupDependencyRootProjects**](DependenciesAPI.md#GetReleaseGroupDependencyRootProjects) | **Get** /v2/release-groups/{projectGroupId}/releases/{projectGroupReleaseId}/root-projects | 
[**GetRevisionDependencies**](DependenciesAPI.md#GetRevisionDependencies) | **Get** /revisions/{locator}/dependencies | 
[**GetRevisionDependenciesPost**](DependenciesAPI.md#GetRevisionDependenciesPost) | **Post** /revisions/{locator}/list-dependencies | 



## GetCustomLicenses

> GetCustomLicenses200Response GetCustomLicenses(ctx).ProjectLocator(projectLocator).Page(page).PageSize(pageSize).Execute()





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
	projectLocator := "custom+1234/my-project" // string | Optional project locator to filter results to custom licenses used within a specific project. If not provided, returns custom licenses across all projects the user has access to.  (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetCustomLicenses(context.Background()).ProjectLocator(projectLocator).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetCustomLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomLicenses`: GetCustomLicenses200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetCustomLicenses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectLocator** | **string** | Optional project locator to filter results to custom licenses used within a specific project. If not provided, returns custom licenses across all projects the user has access to.  | 
 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetCustomLicenses200Response**](GetCustomLicenses200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalDependency

> GetGlobalDependency200Response GetGlobalDependency(ctx, locator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()





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
	locator := "locator_example" // string | The locator of the dependency to retrieve
	includeResolutionNotes := true // bool | Include resolution notes in issue data (optional)
	includeLicenseText := true // bool | Include full license text in license data (optional)
	includeCopyright := true // bool | Include copyright information in license data (optional)
	includeMatches := true // bool | Include license match details in license data (optional)
	includeDownloadUrl := true // bool | Include download URL in package data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetGlobalDependency(context.Background(), locator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetGlobalDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGlobalDependency`: GetGlobalDependency200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetGlobalDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The locator of the dependency to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeResolutionNotes** | **bool** | Include resolution notes in issue data | 
 **includeLicenseText** | **bool** | Include full license text in license data | 
 **includeCopyright** | **bool** | Include copyright information in license data | 
 **includeMatches** | **bool** | Include license match details in license data | 
 **includeDownloadUrl** | **bool** | Include download URL in package data | 

### Return type

[**GetGlobalDependency200Response**](GetGlobalDependency200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDependencies

> GetProjectDependencies200Response GetProjectDependencies(ctx, locator).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Sources(sources).RootProjects(rootProjects).PackageLabels(packageLabels).VendoredPath(vendoredPath).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Page(page).Count(count).Execute()





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
	sources := []string{"Sources_example"} // []string | Filter dependencies by source type (managed or vendored). Only supported on project scope. (optional)
	rootProjects := []string{"Inner_example"} // []string | Filter release group dependencies by root projects (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter dependencies by package label IDs (optional)
	vendoredPath := "vendoredPath_example" // string | Filter to vendored dependencies found under this path prefix. Only supported on project scope. (optional)
	includeResolutionNotes := true // bool | Include resolution notes in issue data (optional)
	includeLicenseText := true // bool | Include full license text in license data (optional)
	includeCopyright := true // bool | Include copyright information in license data (optional)
	includeMatches := true // bool | Include license match details in license data (optional)
	includeDownloadUrl := true // bool | Include download URL in package data (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetProjectDependencies(context.Background(), locator).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Sources(sources).RootProjects(rootProjects).PackageLabels(packageLabels).VendoredPath(vendoredPath).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Page(page).Count(count).Execute()
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
 **sources** | **[]string** | Filter dependencies by source type (managed or vendored). Only supported on project scope. | 
 **rootProjects** | **[]string** | Filter release group dependencies by root projects | 
 **packageLabels** | **[]string** | Filter dependencies by package label IDs | 
 **vendoredPath** | **string** | Filter to vendored dependencies found under this path prefix. Only supported on project scope. | 
 **includeResolutionNotes** | **bool** | Include resolution notes in issue data | 
 **includeLicenseText** | **bool** | Include full license text in license data | 
 **includeCopyright** | **bool** | Include copyright information in license data | 
 **includeMatches** | **bool** | Include license match details in license data | 
 **includeDownloadUrl** | **bool** | Include download URL in package data | 
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


## GetProjectDependency

> GetProjectDependency200Response GetProjectDependency(ctx, locator, dependencyRevisionLocator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()





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
	locator := "locator_example" // string | The locator of the project revision
	dependencyRevisionLocator := "dependencyRevisionLocator_example" // string | The locator of the dependency to retrieve
	includeResolutionNotes := true // bool | Include resolution notes in issue data (optional)
	includeLicenseText := true // bool | Include full license text in license data (optional)
	includeCopyright := true // bool | Include copyright information in license data (optional)
	includeMatches := true // bool | Include license match details in license data (optional)
	includeDownloadUrl := true // bool | Include download URL in package data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetProjectDependency(context.Background(), locator, dependencyRevisionLocator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetProjectDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectDependency`: GetProjectDependency200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetProjectDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The locator of the project revision | 
**dependencyRevisionLocator** | **string** | The locator of the dependency to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeResolutionNotes** | **bool** | Include resolution notes in issue data | 
 **includeLicenseText** | **bool** | Include full license text in license data | 
 **includeCopyright** | **bool** | Include copyright information in license data | 
 **includeMatches** | **bool** | Include license match details in license data | 
 **includeDownloadUrl** | **bool** | Include download URL in package data | 

### Return type

[**GetProjectDependency200Response**](GetProjectDependency200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectDependencyCount

> GetProjectDependencyCount200Response GetProjectDependencyCount(ctx, locator).Sources(sources).Execute()





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
	locator := "locator_example" // string | The locator of the project revision
	sources := []string{"Sources_example"} // []string | Filter dependencies by source type (managed or vendored). Only supported on project scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetProjectDependencyCount(context.Background(), locator).Sources(sources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetProjectDependencyCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectDependencyCount`: GetProjectDependencyCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetProjectDependencyCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The locator of the project revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectDependencyCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sources** | **[]string** | Filter dependencies by source type (managed or vendored). Only supported on project scope. | 

### Return type

[**GetProjectDependencyCount200Response**](GetProjectDependencyCount200Response.md)

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

> GetProjectDependencies200Response GetReleaseGroupDependencies(ctx, projectGroupId, projectGroupReleaseId).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Sources(sources).RootProjects(rootProjects).PackageLabels(packageLabels).VendoredPath(vendoredPath).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Page(page).Count(count).Execute()





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
	confidence := []string{"Confidence_example"} // []string | Filter dependencies by confidence (optional)
	sources := []string{"Sources_example"} // []string | Filter dependencies by source type (managed or vendored). Only supported on project scope. (optional)
	rootProjects := []string{"Inner_example"} // []string | Filter release group dependencies by root projects (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter dependencies by package label IDs (optional)
	vendoredPath := "vendoredPath_example" // string | Filter to vendored dependencies found under this path prefix. Only supported on project scope. (optional)
	includeResolutionNotes := true // bool | Include resolution notes in issue data (optional)
	includeLicenseText := true // bool | Include full license text in license data (optional)
	includeCopyright := true // bool | Include copyright information in license data (optional)
	includeMatches := true // bool | Include license match details in license data (optional)
	includeDownloadUrl := true // bool | Include download URL in package data (optional)
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependencies(context.Background(), projectGroupId, projectGroupReleaseId).Locators(locators).Title(title).Status(status).Depth(depth).LayerDepth(layerDepth).HasIssues(hasIssues).Licenses(licenses).Fetchers(fetchers).ShowIgnored(showIgnored).Confidence(confidence).Sources(sources).RootProjects(rootProjects).PackageLabels(packageLabels).VendoredPath(vendoredPath).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Page(page).Count(count).Execute()
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
 **confidence** | **[]string** | Filter dependencies by confidence | 
 **sources** | **[]string** | Filter dependencies by source type (managed or vendored). Only supported on project scope. | 
 **rootProjects** | **[]string** | Filter release group dependencies by root projects | 
 **packageLabels** | **[]string** | Filter dependencies by package label IDs | 
 **vendoredPath** | **string** | Filter to vendored dependencies found under this path prefix. Only supported on project scope. | 
 **includeResolutionNotes** | **bool** | Include resolution notes in issue data | 
 **includeLicenseText** | **bool** | Include full license text in license data | 
 **includeCopyright** | **bool** | Include copyright information in license data | 
 **includeMatches** | **bool** | Include license match details in license data | 
 **includeDownloadUrl** | **bool** | Include download URL in package data | 
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


## GetReleaseGroupDependency

> GetProjectDependency200Response GetReleaseGroupDependency(ctx, projectGroupId, projectGroupReleaseId, dependencyRevisionLocator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()





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
	projectGroupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release
	dependencyRevisionLocator := "dependencyRevisionLocator_example" // string | The locator of the dependency to retrieve
	includeResolutionNotes := true // bool | Include resolution notes in issue data (optional)
	includeLicenseText := true // bool | Include full license text in license data (optional)
	includeCopyright := true // bool | Include copyright information in license data (optional)
	includeMatches := true // bool | Include license match details in license data (optional)
	includeDownloadUrl := true // bool | Include download URL in package data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependency(context.Background(), projectGroupId, projectGroupReleaseId, dependencyRevisionLocator).IncludeResolutionNotes(includeResolutionNotes).IncludeLicenseText(includeLicenseText).IncludeCopyright(includeCopyright).IncludeMatches(includeMatches).IncludeDownloadUrl(includeDownloadUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetReleaseGroupDependency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupDependency`: GetProjectDependency200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetReleaseGroupDependency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 
**dependencyRevisionLocator** | **string** | The locator of the dependency to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupDependencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeResolutionNotes** | **bool** | Include resolution notes in issue data | 
 **includeLicenseText** | **bool** | Include full license text in license data | 
 **includeCopyright** | **bool** | Include copyright information in license data | 
 **includeMatches** | **bool** | Include license match details in license data | 
 **includeDownloadUrl** | **bool** | Include download URL in package data | 

### Return type

[**GetProjectDependency200Response**](GetProjectDependency200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupDependencyCount

> GetReleaseGroupDependencyCount200Response GetReleaseGroupDependencyCount(ctx, projectGroupId, projectGroupReleaseId).Sources(sources).Execute()





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
	projectGroupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release
	sources := []string{"Sources_example"} // []string | Filter dependencies by source type (managed or vendored). Only supported on project scope. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DependenciesAPI.GetReleaseGroupDependencyCount(context.Background(), projectGroupId, projectGroupReleaseId).Sources(sources).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetReleaseGroupDependencyCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupDependencyCount`: GetReleaseGroupDependencyCount200Response
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetReleaseGroupDependencyCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectGroupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupDependencyCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sources** | **[]string** | Filter dependencies by source type (managed or vendored). Only supported on project scope. | 

### Return type

[**GetReleaseGroupDependencyCount200Response**](GetReleaseGroupDependencyCount200Response.md)

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
	resp, r, err := apiClient.DependenciesAPI.GetRevisionDependencies(context.Background(), locator).Limit(limit).Offset(offset).IncludeIgnored(includeIgnored).IncludeHashData(includeHashData).IncludeLicenseText(includeLicenseText).IncludeLocators(includeLocators).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetRevisionDependencies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependencies`: []GetRevisionDependenciesPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetRevisionDependencies`: %v\n", resp)
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
	resp, r, err := apiClient.DependenciesAPI.GetRevisionDependenciesPost(context.Background(), locator).GetRevisionDependenciesPostRequest(getRevisionDependenciesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DependenciesAPI.GetRevisionDependenciesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependenciesPost`: []GetRevisionDependenciesPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `DependenciesAPI.GetRevisionDependenciesPost`: %v\n", resp)
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

