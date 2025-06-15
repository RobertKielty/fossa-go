# \OrganizationLabelsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationLabel**](OrganizationLabelsAPI.md#CreateOrganizationLabel) | **Post** /organizations/labels | 
[**DeleteOrganizationLabel**](OrganizationLabelsAPI.md#DeleteOrganizationLabel) | **Delete** /organizations/labels/{id} | 
[**GetOrganizationLabel**](OrganizationLabelsAPI.md#GetOrganizationLabel) | **Get** /organizations/labels/{id} | 
[**GetOrganizationLabels**](OrganizationLabelsAPI.md#GetOrganizationLabels) | **Get** /organizations/labels | 



## CreateOrganizationLabel

> CreateOrganizationLabel201Response CreateOrganizationLabel(ctx).CreateOrganizationLabelRequest(createOrganizationLabelRequest).Execute()





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
	createOrganizationLabelRequest := *openapiclient.NewCreateOrganizationLabelRequest() // CreateOrganizationLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsAPI.CreateOrganizationLabel(context.Background()).CreateOrganizationLabelRequest(createOrganizationLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsAPI.CreateOrganizationLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationLabel`: CreateOrganizationLabel201Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsAPI.CreateOrganizationLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationLabelRequest** | [**CreateOrganizationLabelRequest**](CreateOrganizationLabelRequest.md) |  | 

### Return type

[**CreateOrganizationLabel201Response**](CreateOrganizationLabel201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationLabel

> DeleteOrganizationLabel(ctx, id).Execute()





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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationLabelsAPI.DeleteOrganizationLabel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsAPI.DeleteOrganizationLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationLabelRequest struct via the builder pattern


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


## GetOrganizationLabel

> GetOrganizationLabels200ResponseLabelsInner GetOrganizationLabel(ctx, id).Execute()





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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLabelsAPI.GetOrganizationLabel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsAPI.GetOrganizationLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLabel`: GetOrganizationLabels200ResponseLabelsInner
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsAPI.GetOrganizationLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationLabels200ResponseLabelsInner**](GetOrganizationLabels200ResponseLabelsInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationLabels

> GetOrganizationLabels200Response GetOrganizationLabels(ctx).Execute()





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
	resp, r, err := apiClient.OrganizationLabelsAPI.GetOrganizationLabels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLabelsAPI.GetOrganizationLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLabels`: GetOrganizationLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLabelsAPI.GetOrganizationLabels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLabelsRequest struct via the builder pattern


### Return type

[**GetOrganizationLabels200Response**](GetOrganizationLabels200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

