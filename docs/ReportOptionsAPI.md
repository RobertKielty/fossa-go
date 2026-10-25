# \ReportOptionsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReportOption**](ReportOptionsAPI.md#CreateReportOption) | **Post** /report-options | 
[**DeleteReportOptionById**](ReportOptionsAPI.md#DeleteReportOptionById) | **Delete** /report-options/{id} | 
[**GetAllReportOptions**](ReportOptionsAPI.md#GetAllReportOptions) | **Get** /report-options | 
[**UpdateReportOptionById**](ReportOptionsAPI.md#UpdateReportOptionById) | **Put** /report-options/{id} | 



## CreateReportOption

> CreateReportOption201Response CreateReportOption(ctx).CreateReportOptionRequest(createReportOptionRequest).Execute()





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
	createReportOptionRequest := *openapiclient.NewCreateReportOptionRequest("My Custom Report", *openapiclient.NewCreateReportOptionRequestOptions(*openapiclient.NewGetAllReportOptions200ResponseResultsInnerOptionsSections(false, false, false, false, false, false), *openapiclient.NewGetAllReportOptions200ResponseResultsInnerOptionsToggles(false), *openapiclient.NewGetAllReportOptions200ResponseResultsInnerOptionsExcludeFields([]int32{int32(123)}), *openapiclient.NewCreateReportOptionRequestOptionsDependencyData(false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false))) // CreateReportOptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportOptionsAPI.CreateReportOption(context.Background()).CreateReportOptionRequest(createReportOptionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportOptionsAPI.CreateReportOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReportOption`: CreateReportOption201Response
	fmt.Fprintf(os.Stdout, "Response from `ReportOptionsAPI.CreateReportOption`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReportOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReportOptionRequest** | [**CreateReportOptionRequest**](CreateReportOptionRequest.md) |  | 

### Return type

[**CreateReportOption201Response**](CreateReportOption201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReportOptionById

> DeleteReportOptionById(ctx, id).Execute()





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
	id := int32(1) // int32 | The unique identifier of the report option to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportOptionsAPI.DeleteReportOptionById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportOptionsAPI.DeleteReportOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique identifier of the report option to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReportOptionByIdRequest struct via the builder pattern


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


## GetAllReportOptions

> GetAllReportOptions200Response GetAllReportOptions(ctx).Execute()





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
	resp, r, err := apiClient.ReportOptionsAPI.GetAllReportOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportOptionsAPI.GetAllReportOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReportOptions`: GetAllReportOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `ReportOptionsAPI.GetAllReportOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReportOptionsRequest struct via the builder pattern


### Return type

[**GetAllReportOptions200Response**](GetAllReportOptions200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReportOptionById

> CreateReportOption201Response UpdateReportOptionById(ctx, id).UpdateReportOptionByIdRequest(updateReportOptionByIdRequest).Execute()





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
	id := int32(1) // int32 | The unique identifier of the report option to update
	updateReportOptionByIdRequest := *openapiclient.NewUpdateReportOptionByIdRequest() // UpdateReportOptionByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReportOptionsAPI.UpdateReportOptionById(context.Background(), id).UpdateReportOptionByIdRequest(updateReportOptionByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportOptionsAPI.UpdateReportOptionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReportOptionById`: CreateReportOption201Response
	fmt.Fprintf(os.Stdout, "Response from `ReportOptionsAPI.UpdateReportOptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique identifier of the report option to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReportOptionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReportOptionByIdRequest** | [**UpdateReportOptionByIdRequest**](UpdateReportOptionByIdRequest.md) |  | 

### Return type

[**CreateReportOption201Response**](CreateReportOption201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

