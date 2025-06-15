# \IssueFiltersAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSavedFilter**](IssueFiltersAPI.md#CreateSavedFilter) | **Post** /issue-filters | Create a saved filter for your organization
[**DeleteSavedFilter**](IssueFiltersAPI.md#DeleteSavedFilter) | **Delete** /issue-filters/{filterId} | Delete a saved filter by ID
[**GetSavedFilterById**](IssueFiltersAPI.md#GetSavedFilterById) | **Get** /issue-filters/{filterId} | Get a saved filter by ID
[**ListSavedFilters**](IssueFiltersAPI.md#ListSavedFilters) | **Get** /issue-filters | List out saved filters
[**UpdateSavedFilter**](IssueFiltersAPI.md#UpdateSavedFilter) | **Put** /issue-filters/{filterId} | Update a saved filter for your organization



## CreateSavedFilter

> ListSavedFilters200ResponseInner CreateSavedFilter(ctx).CreateSavedFilterRequest(createSavedFilterRequest).Execute()

Create a saved filter for your organization



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
	createSavedFilterRequest := *openapiclient.NewCreateSavedFilterRequest() // CreateSavedFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFiltersAPI.CreateSavedFilter(context.Background()).CreateSavedFilterRequest(createSavedFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFiltersAPI.CreateSavedFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSavedFilter`: ListSavedFilters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IssueFiltersAPI.CreateSavedFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSavedFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSavedFilterRequest** | [**CreateSavedFilterRequest**](CreateSavedFilterRequest.md) |  | 

### Return type

[**ListSavedFilters200ResponseInner**](ListSavedFilters200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSavedFilter

> DeleteSavedFilter(ctx, filterId).Execute()

Delete a saved filter by ID



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
	filterId := int32(56) // int32 | The ID of the saved filter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueFiltersAPI.DeleteSavedFilter(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFiltersAPI.DeleteSavedFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32** | The ID of the saved filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSavedFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSavedFilterById

> ListSavedFilters200ResponseInner GetSavedFilterById(ctx, filterId).Execute()

Get a saved filter by ID



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
	filterId := int32(56) // int32 | The ID of the saved filter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFiltersAPI.GetSavedFilterById(context.Background(), filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFiltersAPI.GetSavedFilterById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSavedFilterById`: ListSavedFilters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IssueFiltersAPI.GetSavedFilterById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32** | The ID of the saved filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSavedFilterByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSavedFilters200ResponseInner**](ListSavedFilters200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSavedFilters

> []ListSavedFilters200ResponseInner ListSavedFilters(ctx).Category(category).Execute()

List out saved filters



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
	category := "category_example" // string | The category of the issue filters.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFiltersAPI.ListSavedFilters(context.Background()).Category(category).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFiltersAPI.ListSavedFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSavedFilters`: []ListSavedFilters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IssueFiltersAPI.ListSavedFilters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSavedFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | The category of the issue filters. | 

### Return type

[**[]ListSavedFilters200ResponseInner**](ListSavedFilters200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSavedFilter

> ListSavedFilters200ResponseInner UpdateSavedFilter(ctx, filterId).UpdateSavedFilterRequest(updateSavedFilterRequest).Execute()

Update a saved filter for your organization



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
	filterId := int32(56) // int32 | The ID of the saved filter.
	updateSavedFilterRequest := *openapiclient.NewUpdateSavedFilterRequest() // UpdateSavedFilterRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFiltersAPI.UpdateSavedFilter(context.Background(), filterId).UpdateSavedFilterRequest(updateSavedFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFiltersAPI.UpdateSavedFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSavedFilter`: ListSavedFilters200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `IssueFiltersAPI.UpdateSavedFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int32** | The ID of the saved filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSavedFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSavedFilterRequest** | [**UpdateSavedFilterRequest**](UpdateSavedFilterRequest.md) |  | 

### Return type

[**ListSavedFilters200ResponseInner**](ListSavedFilters200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

