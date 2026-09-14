# \BinaryAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReleaseComponentsCount**](BinaryAPI.md#GetReleaseComponentsCount) | **Get** /binary/release-group/{releaseGroupId}/release/{releaseId}/components/count | 
[**GetReleaseComponentsPaths**](BinaryAPI.md#GetReleaseComponentsPaths) | **Get** /binary/release-group/{releaseGroupId}/release/{releaseId}/components/paths | 
[**GetReleaseDependencyConfidence**](BinaryAPI.md#GetReleaseDependencyConfidence) | **Get** /binary/release/{releaseId}/dependency-confidence | 
[**GetRevisionComponentMatches**](BinaryAPI.md#GetRevisionComponentMatches) | **Get** /binary/{revisionLocator}/{componentId}/matches | 
[**GetRevisionComponentsCount**](BinaryAPI.md#GetRevisionComponentsCount) | **Get** /binary/revision/{revisionLocator}/components/count | 
[**GetRevisionComponentsPaths**](BinaryAPI.md#GetRevisionComponentsPaths) | **Get** /binary/revision/{revisionLocator}/components/paths | 
[**GetRevisionDependencyComponents**](BinaryAPI.md#GetRevisionDependencyComponents) | **Get** /binary/{revisionLocator}/{dependencyLocator}/components | 
[**GetRevisionDependencyConfidence**](BinaryAPI.md#GetRevisionDependencyConfidence) | **Get** /binary/{revisionLocator}/dependency-confidence | 
[**GetSingleReleaseDependencyConfidence**](BinaryAPI.md#GetSingleReleaseDependencyConfidence) | **Get** /binary/release/{releaseId}/dependency-confidence/{dependencyLocator} | 
[**GetSingleRevisionDependencyConfidence**](BinaryAPI.md#GetSingleRevisionDependencyConfidence) | **Get** /binary/{revisionLocator}/dependency-confidence/{dependencyLocator} | 



## GetReleaseComponentsCount

> GetReleaseComponentsCount200Response GetReleaseComponentsCount(ctx, releaseGroupId, releaseId).Execute()





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
	releaseGroupId := float32(8.14) // float32 | Release group id
	releaseId := float32(8.14) // float32 | Release id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetReleaseComponentsCount(context.Background(), releaseGroupId, releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetReleaseComponentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseComponentsCount`: GetReleaseComponentsCount200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetReleaseComponentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseGroupId** | **float32** | Release group id | 
**releaseId** | **float32** | Release id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseComponentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetReleaseComponentsCount200Response**](GetReleaseComponentsCount200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseComponentsPaths

> GetRevisionComponentsPaths200Response GetReleaseComponentsPaths(ctx, releaseGroupId, releaseId).Path(path).Search(search).Execute()





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
	releaseGroupId := float32(8.14) // float32 | Release group id
	releaseId := float32(8.14) // float32 | Release id
	path := "path_example" // string | Path to find components at (optional)
	search := "search_example" // string | Filter the results to dependencies whose title matches this search string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetReleaseComponentsPaths(context.Background(), releaseGroupId, releaseId).Path(path).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetReleaseComponentsPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseComponentsPaths`: GetRevisionComponentsPaths200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetReleaseComponentsPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseGroupId** | **float32** | Release group id | 
**releaseId** | **float32** | Release id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseComponentsPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **string** | Path to find components at | 
 **search** | **string** | Filter the results to dependencies whose title matches this search string. | 

### Return type

[**GetRevisionComponentsPaths200Response**](GetRevisionComponentsPaths200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseDependencyConfidence

> GetRevisionDependencyConfidence200Response GetReleaseDependencyConfidence(ctx, releaseId).Execute()





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
	releaseId := float32(8.14) // float32 | Release id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetReleaseDependencyConfidence(context.Background(), releaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetReleaseDependencyConfidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseDependencyConfidence`: GetRevisionDependencyConfidence200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetReleaseDependencyConfidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **float32** | Release id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseDependencyConfidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRevisionDependencyConfidence200Response**](GetRevisionDependencyConfidence200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionComponentMatches

> GetRevisionComponentMatches200Response GetRevisionComponentMatches(ctx, revisionLocator, componentId).Page(page).PageSize(pageSize).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision
	componentId := "componentId_example" // string | The ID of the component
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetRevisionComponentMatches(context.Background(), revisionLocator, componentId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetRevisionComponentMatches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionComponentMatches`: GetRevisionComponentMatches200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetRevisionComponentMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 
**componentId** | **string** | The ID of the component | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionComponentMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetRevisionComponentMatches200Response**](GetRevisionComponentMatches200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionComponentsCount

> GetRevisionComponentsCount200Response GetRevisionComponentsCount(ctx, revisionLocator).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetRevisionComponentsCount(context.Background(), revisionLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetRevisionComponentsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionComponentsCount`: GetRevisionComponentsCount200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetRevisionComponentsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionComponentsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRevisionComponentsCount200Response**](GetRevisionComponentsCount200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionComponentsPaths

> GetRevisionComponentsPaths200Response GetRevisionComponentsPaths(ctx, revisionLocator).Path(path).Search(search).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision
	path := "path_example" // string | Path to find components at (optional)
	search := "search_example" // string | Filter the results to dependencies whose title matches this search string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetRevisionComponentsPaths(context.Background(), revisionLocator).Path(path).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetRevisionComponentsPaths``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionComponentsPaths`: GetRevisionComponentsPaths200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetRevisionComponentsPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionComponentsPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path to find components at | 
 **search** | **string** | Filter the results to dependencies whose title matches this search string. | 

### Return type

[**GetRevisionComponentsPaths200Response**](GetRevisionComponentsPaths200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionDependencyComponents

> GetRevisionDependencyComponents200Response GetRevisionDependencyComponents(ctx, revisionLocator, dependencyLocator).Page(page).PageSize(pageSize).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision
	dependencyLocator := "dependencyLocator_example" // string | Dependency revision
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetRevisionDependencyComponents(context.Background(), revisionLocator, dependencyLocator).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetRevisionDependencyComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependencyComponents`: GetRevisionDependencyComponents200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetRevisionDependencyComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 
**dependencyLocator** | **string** | Dependency revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionDependencyComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The specific page of data to return | [default to 1]
 **pageSize** | **int32** | The number of items to return in each page of results | [default to 10]

### Return type

[**GetRevisionDependencyComponents200Response**](GetRevisionDependencyComponents200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRevisionDependencyConfidence

> GetRevisionDependencyConfidence200Response GetRevisionDependencyConfidence(ctx, revisionLocator).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetRevisionDependencyConfidence(context.Background(), revisionLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetRevisionDependencyConfidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRevisionDependencyConfidence`: GetRevisionDependencyConfidence200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetRevisionDependencyConfidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRevisionDependencyConfidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRevisionDependencyConfidence200Response**](GetRevisionDependencyConfidence200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleReleaseDependencyConfidence

> GetRevisionDependencyConfidence200Response GetSingleReleaseDependencyConfidence(ctx, releaseId, dependencyLocator).Execute()





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
	releaseId := float32(8.14) // float32 | Release id
	dependencyLocator := "dependencyLocator_example" // string | Dependency revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetSingleReleaseDependencyConfidence(context.Background(), releaseId, dependencyLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetSingleReleaseDependencyConfidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleReleaseDependencyConfidence`: GetRevisionDependencyConfidence200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetSingleReleaseDependencyConfidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**releaseId** | **float32** | Release id | 
**dependencyLocator** | **string** | Dependency revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleReleaseDependencyConfidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRevisionDependencyConfidence200Response**](GetRevisionDependencyConfidence200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleRevisionDependencyConfidence

> GetRevisionDependencyConfidence200Response GetSingleRevisionDependencyConfidence(ctx, revisionLocator, dependencyLocator).Execute()





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
	revisionLocator := "revisionLocator_example" // string | Binary Decomposition project revision
	dependencyLocator := "dependencyLocator_example" // string | Dependency revision

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BinaryAPI.GetSingleRevisionDependencyConfidence(context.Background(), revisionLocator, dependencyLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BinaryAPI.GetSingleRevisionDependencyConfidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSingleRevisionDependencyConfidence`: GetRevisionDependencyConfidence200Response
	fmt.Fprintf(os.Stdout, "Response from `BinaryAPI.GetSingleRevisionDependencyConfidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**revisionLocator** | **string** | Binary Decomposition project revision | 
**dependencyLocator** | **string** | Dependency revision | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleRevisionDependencyConfidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetRevisionDependencyConfidence200Response**](GetRevisionDependencyConfidence200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

