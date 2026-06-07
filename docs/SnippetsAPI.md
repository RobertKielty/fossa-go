# \SnippetsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComparedSnippetPackages**](SnippetsAPI.md#GetComparedSnippetPackages) | **Get** /revisions/{locator}/snippets/compare/{olderRevisionLocator}/{status}/packages | Get compared snippet packages between two revisions
[**GetComparedSnippetPaths**](SnippetsAPI.md#GetComparedSnippetPaths) | **Get** /revisions/{locator}/snippets/compare/{olderRevisionLocator}/{status}/paths | Get compared snippet paths between two revisions
[**GetComparedSnippets**](SnippetsAPI.md#GetComparedSnippets) | **Get** /revisions/{locator}/snippets/compare/{olderRevisionLocator}/{status} | Get compared snippets between two revisions
[**GetSnippetCount**](SnippetsAPI.md#GetSnippetCount) | **Get** /revisions/{locator}/snippets/count | Get snippet count
[**GetSnippetDetails**](SnippetsAPI.md#GetSnippetDetails) | **Get** /revisions/{locator}/snippets/{snippetId} | Get the details of a specific snippet
[**GetSnippetMatchDetails**](SnippetsAPI.md#GetSnippetMatchDetails) | **Get** /revisions/{locator}/snippets/{snippetId}/matches/{path} | Get the details of a specific snippet match
[**GetSnippetPackages**](SnippetsAPI.md#GetSnippetPackages) | **Get** /revisions/{locator}/snippets/packages | Get snippet packages
[**GetSnippetPaths**](SnippetsAPI.md#GetSnippetPaths) | **Get** /revisions/{locator}/snippets/paths | Get snippet paths
[**GetSnippets**](SnippetsAPI.md#GetSnippets) | **Get** /revisions/{locator}/snippets | Get snippets
[**RejectSnippets**](SnippetsAPI.md#RejectSnippets) | **Post** /revisions/{locator}/snippets/reject | Reject snippet matches
[**UnrejectSnippets**](SnippetsAPI.md#UnrejectSnippets) | **Post** /revisions/{locator}/snippets/unreject | Unreject snippet matches



## GetComparedSnippetPackages

> GetSnippetPackages200Response GetComparedSnippetPackages(ctx, locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()

Get compared snippet packages between two revisions



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
	locator := "locator_example" // string | The base revision locator
	olderRevisionLocator := "olderRevisionLocator_example" // string | An older revision locator from the same project
	status := "status_example" // string | The status of snippet packages to retrieve
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "matchCount_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetComparedSnippetPackages(context.Background(), locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetComparedSnippetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComparedSnippetPackages`: GetSnippetPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetComparedSnippetPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The base revision locator | 
**olderRevisionLocator** | **string** | An older revision locator from the same project | 
**status** | **string** | The status of snippet packages to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComparedSnippetPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 
 **sort** | **string** | Sort order for results | [default to &quot;matchCount_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetSnippetPackages200Response**](GetSnippetPackages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComparedSnippetPaths

> GetSnippetPaths200Response GetComparedSnippetPaths(ctx, locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()

Get compared snippet paths between two revisions



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
	locator := "locator_example" // string | The base revision locator
	olderRevisionLocator := "olderRevisionLocator_example" // string | An older revision locator from the same project
	status := "status_example" // string | The status of snippets to retrieve
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetComparedSnippetPaths(context.Background(), locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetComparedSnippetPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComparedSnippetPaths`: GetSnippetPaths200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetComparedSnippetPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The base revision locator | 
**olderRevisionLocator** | **string** | An older revision locator from the same project | 
**status** | **string** | The status of snippets to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComparedSnippetPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 

### Return type

[**GetSnippetPaths200Response**](GetSnippetPaths200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComparedSnippets

> GetSnippets200Response GetComparedSnippets(ctx, locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()

Get compared snippets between two revisions



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
	locator := "locator_example" // string | The base revision locator
	olderRevisionLocator := "olderRevisionLocator_example" // string | An older revision locator from the same project
	status := "status_example" // string | The status of snippets to retrieve
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "matchCount_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetComparedSnippets(context.Background(), locator, olderRevisionLocator, status).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetComparedSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComparedSnippets`: GetSnippets200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetComparedSnippets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The base revision locator | 
**olderRevisionLocator** | **string** | An older revision locator from the same project | 
**status** | **string** | The status of snippets to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComparedSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 
 **sort** | **string** | Sort order for results | [default to &quot;matchCount_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetSnippets200Response**](GetSnippets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippetCount

> UpdateIssues200Response GetSnippetCount(ctx, locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()

Get snippet count



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
	locator := "locator_example" // string | The revision locator
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetCount(context.Background(), locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetCount`: UpdateIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 

### Return type

[**UpdateIssues200Response**](UpdateIssues200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippetDetails

> GetSnippetDetails200Response GetSnippetDetails(ctx, locator, snippetId).Execute()

Get the details of a specific snippet



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
	locator := "locator_example" // string | The revision locator
	snippetId := "snippetId_example" // string | The unique identifier of the snippet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetDetails(context.Background(), locator, snippetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetDetails`: GetSnippetDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 
**snippetId** | **string** | The unique identifier of the snippet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetSnippetDetails200Response**](GetSnippetDetails200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippetMatchDetails

> GetSnippetMatchDetails200Response GetSnippetMatchDetails(ctx, locator, snippetId, path).Execute()

Get the details of a specific snippet match



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
	locator := "locator_example" // string | The revision locator
	snippetId := "snippetId_example" // string | The unique identifier of the snippet
	path := "path_example" // string | The file path where the snippet match occurred

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetMatchDetails(context.Background(), locator, snippetId, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetMatchDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetMatchDetails`: GetSnippetMatchDetails200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetMatchDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 
**snippetId** | **string** | The unique identifier of the snippet | 
**path** | **string** | The file path where the snippet match occurred | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetMatchDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetSnippetMatchDetails200Response**](GetSnippetMatchDetails200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippetPackages

> GetSnippetPackages200Response GetSnippetPackages(ctx, locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()

Get snippet packages



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
	locator := "locator_example" // string | The revision locator
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "matchCount_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetPackages(context.Background(), locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetPackages`: GetSnippetPackages200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetPackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 
 **sort** | **string** | Sort order for results | [default to &quot;matchCount_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetSnippetPackages200Response**](GetSnippetPackages200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippetPaths

> GetSnippetPaths200Response GetSnippetPaths(ctx, locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()

Get snippet paths



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
	locator := "locator_example" // string | The revision locator
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetPaths(context.Background(), locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetPaths`: GetSnippetPaths200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 

### Return type

[**GetSnippetPaths200Response**](GetSnippetPaths200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnippets

> GetSnippets200Response GetSnippets(ctx, locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()

Get snippets



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
	locator := "locator_example" // string | The revision locator
	path := "path_example" // string | The path to filter snippets by
	ids := []string{"Inner_example"} // []string | Filter by specific snippet IDs (optional)
	packageIds := []string{"Inner_example"} // []string | Filter by specific snippet package IDs (optional)
	search := "search_example" // string | Search term for filtering snippets by package name (optional)
	rejectionStatus := []string{"RejectionStatus_example"} // []string | Filter by rejection status (optional)
	packageLabels := []string{"Inner_example"} // []string | Filter by package labels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "matchCount_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippets(context.Background(), locator).Path(path).Ids(ids).PackageIds(packageIds).Search(search).RejectionStatus(rejectionStatus).PackageLabels(packageLabels).Sort(sort).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippets`: GetSnippets200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | The path to filter snippets by | 
 **ids** | **[]string** | Filter by specific snippet IDs | 
 **packageIds** | **[]string** | Filter by specific snippet package IDs | 
 **search** | **string** | Search term for filtering snippets by package name | 
 **rejectionStatus** | **[]string** | Filter by rejection status | 
 **packageLabels** | **[]string** | Filter by package labels | 
 **sort** | **string** | Sort order for results | [default to &quot;matchCount_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetSnippets200Response**](GetSnippets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectSnippets

> RejectSnippets(ctx, locator).RejectSnippetsRequest(rejectSnippetsRequest).Execute()

Reject snippet matches



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
	locator := "locator_example" // string | The revision locator
	rejectSnippetsRequest := *openapiclient.NewRejectSnippetsRequest("Path_example") // RejectSnippetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SnippetsAPI.RejectSnippets(context.Background(), locator).RejectSnippetsRequest(rejectSnippetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.RejectSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectSnippetsRequest** | [**RejectSnippetsRequest**](RejectSnippetsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnrejectSnippets

> UnrejectSnippets(ctx, locator).RejectSnippetsRequest(rejectSnippetsRequest).Execute()

Unreject snippet matches



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
	locator := "locator_example" // string | The revision locator
	rejectSnippetsRequest := *openapiclient.NewRejectSnippetsRequest("Path_example") // RejectSnippetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SnippetsAPI.UnrejectSnippets(context.Background(), locator).RejectSnippetsRequest(rejectSnippetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.UnrejectSnippets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrejectSnippetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectSnippetsRequest** | [**RejectSnippetsRequest**](RejectSnippetsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

