# \ProjectLabelsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectLabels**](ProjectLabelsAPI.md#GetProjectLabels) | **Get** /projects/{locator}/labels | 
[**UpdateProjectLabels**](ProjectLabelsAPI.md#UpdateProjectLabels) | **Put** /projects/{locator} | 



## GetProjectLabels

> []GetProjectLabels200ResponseInner GetProjectLabels(ctx, locator).Execute()





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
	resp, r, err := apiClient.ProjectLabelsAPI.GetProjectLabels(context.Background(), locator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectLabelsAPI.GetProjectLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectLabels`: []GetProjectLabels200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ProjectLabelsAPI.GetProjectLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetProjectLabels200ResponseInner**](GetProjectLabels200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectLabels

> UpdateProjectLabels200Response UpdateProjectLabels(ctx, locator).UpdateProjectLabelsRequest(updateProjectLabelsRequest).Execute()





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
	updateProjectLabelsRequest := *openapiclient.NewUpdateProjectLabelsRequest() // UpdateProjectLabelsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectLabelsAPI.UpdateProjectLabels(context.Background(), locator).UpdateProjectLabelsRequest(updateProjectLabelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectLabelsAPI.UpdateProjectLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectLabels`: UpdateProjectLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `ProjectLabelsAPI.UpdateProjectLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locator** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProjectLabelsRequest** | [**UpdateProjectLabelsRequest**](UpdateProjectLabelsRequest.md) |  | 

### Return type

[**UpdateProjectLabels200Response**](UpdateProjectLabels200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

