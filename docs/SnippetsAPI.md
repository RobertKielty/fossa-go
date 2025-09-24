# \SnippetsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComparedSnippetPaths**](SnippetsAPI.md#GetComparedSnippetPaths) | **Get** /revisions/{locator}/snippets/compare/{olderRevisionLocator}/{status}/paths | 
[**GetComparedSnippets**](SnippetsAPI.md#GetComparedSnippets) | **Get** /revisions/{locator}/snippets/compare/{olderRevisionLocator}/{status} | 
[**GetSnippetById**](SnippetsAPI.md#GetSnippetById) | **Get** /revisions/{locator}/snippets/{snippetId} | 
[**GetSnippetMatchDetails**](SnippetsAPI.md#GetSnippetMatchDetails) | **Get** /revisions/{locator}/snippets/{snippetId}/matches/{path} | 
[**GetSnippetPaths**](SnippetsAPI.md#GetSnippetPaths) | **Get** /revisions/{locator}/snippets/paths | 
[**GetSnippets**](SnippetsAPI.md#GetSnippets) | **Get** /revisions/{locator}/snippets | 
[**RejectSnippets**](SnippetsAPI.md#RejectSnippets) | **Post** /revisions/{locator}/snippets/reject | 
[**UnrejectSnippets**](SnippetsAPI.md#UnrejectSnippets) | **Post** /revisions/{locator}/snippets/unreject | 



## GetComparedSnippetPaths

> GetSnippetPaths200Response GetComparedSnippetPaths(ctx, locator, olderRevisionLocator, status).Path(path).Execute()





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
	path := "path_example" // string | The path from which a single depth of files and/or directories will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetComparedSnippetPaths(context.Background(), locator, olderRevisionLocator, status).Path(path).Execute()
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



 **path** | **string** | The path from which a single depth of files and/or directories will be returned | 

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

> GetSnippets200Response GetComparedSnippets(ctx, locator, olderRevisionLocator, status).Path(path).Ids(ids).Search(search).Confidence(confidence).Sort(sort).Page(page).Count(count).Execute()





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
	search := "search_example" // string | Search term for filtering snippets (optional)
	confidence := []string{"Confidence_example"} // []string | Filter by confidence levels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "confidence_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetComparedSnippets(context.Background(), locator, olderRevisionLocator, status).Path(path).Ids(ids).Search(search).Confidence(confidence).Sort(sort).Page(page).Count(count).Execute()
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
 **search** | **string** | Search term for filtering snippets | 
 **confidence** | **[]string** | Filter by confidence levels | 
 **sort** | **string** | Sort order for results | [default to &quot;confidence_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]

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


## GetSnippetById

> GetSnippetById200Response GetSnippetById(ctx, locator, snippetId).Path(path).Execute()





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
	path := "path_example" // string | The path to filter rejection status by. Defaults to root path ('/') if not provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetById(context.Background(), locator, snippetId).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnippetsAPI.GetSnippetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnippetById`: GetSnippetById200Response
	fmt.Fprintf(os.Stdout, "Response from `SnippetsAPI.GetSnippetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** | The revision locator | 
**snippetId** | **string** | The unique identifier of the snippet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnippetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **path** | **string** | The path to filter rejection status by. Defaults to root path (&#39;/&#39;) if not provided. | 

### Return type

[**GetSnippetById200Response**](GetSnippetById200Response.md)

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


## GetSnippetPaths

> GetSnippetPaths200Response GetSnippetPaths(ctx, locator).Path(path).Execute()





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
	path := "path_example" // string | The path from which a single depth of files and/or directories will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippetPaths(context.Background(), locator).Path(path).Execute()
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

 **path** | **string** | The path from which a single depth of files and/or directories will be returned | 

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

> GetSnippets200Response GetSnippets(ctx, locator).Path(path).Ids(ids).Search(search).Confidence(confidence).Sort(sort).Page(page).Count(count).Execute()





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
	search := "search_example" // string | Search term for filtering snippets (optional)
	confidence := []string{"Confidence_example"} // []string | Filter by confidence levels (optional)
	sort := "sort_example" // string | Sort order for results (optional) (default to "confidence_desc")
	page := int32(56) // int32 | The specific page of data to return (optional) (default to 1)
	count := int32(56) // int32 | The number of items to return in each page of results (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnippetsAPI.GetSnippets(context.Background(), locator).Path(path).Ids(ids).Search(search).Confidence(confidence).Sort(sort).Page(page).Count(count).Execute()
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
 **search** | **string** | Search term for filtering snippets | 
 **confidence** | **[]string** | Filter by confidence levels | 
 **sort** | **string** | Sort order for results | [default to &quot;confidence_desc&quot;]
 **page** | **int32** | The specific page of data to return | [default to 1]
 **count** | **int32** | The number of items to return in each page of results | [default to 20]

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

