# \SBOMAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShareRequest**](SBOMAPI.md#CreateShareRequest) | **Post** /api/v1/share-requests | 
[**GetLinkedOrganizations**](SBOMAPI.md#GetLinkedOrganizations) | **Get** /api/v1/shared-organizations | 
[**GetShareRequests**](SBOMAPI.md#GetShareRequests) | **Get** /api/v1/share-requests | 



## CreateShareRequest

> CreateShareRequest201Response CreateShareRequest(ctx).CreateShareRequestRequest(createShareRequestRequest).Execute()





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
	createShareRequestRequest := *openapiclient.NewCreateShareRequestRequest("RevisionId_example", int64(123)) // CreateShareRequestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMAPI.CreateShareRequest(context.Background()).CreateShareRequestRequest(createShareRequestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.CreateShareRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShareRequest`: CreateShareRequest201Response
	fmt.Fprintf(os.Stdout, "Response from `SBOMAPI.CreateShareRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShareRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShareRequestRequest** | [**CreateShareRequestRequest**](CreateShareRequestRequest.md) |  | 

### Return type

[**CreateShareRequest201Response**](CreateShareRequest201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedOrganizations

> GetLinkedOrganizations200Response GetLinkedOrganizations(ctx).Execute()





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
	resp, r, err := apiClient.SBOMAPI.GetLinkedOrganizations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetLinkedOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedOrganizations`: GetLinkedOrganizations200Response
	fmt.Fprintf(os.Stdout, "Response from `SBOMAPI.GetLinkedOrganizations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedOrganizationsRequest struct via the builder pattern


### Return type

[**GetLinkedOrganizations200Response**](GetLinkedOrganizations200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShareRequests

> GetShareRequests200Response GetShareRequests(ctx).ProjectLocator(projectLocator).Execute()





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
	projectLocator := "projectLocator_example" // string | Filter share requests by project locator (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SBOMAPI.GetShareRequests(context.Background()).ProjectLocator(projectLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetShareRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShareRequests`: GetShareRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `SBOMAPI.GetShareRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShareRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectLocator** | **string** | Filter share requests by project locator | 

### Return type

[**GetShareRequests200Response**](GetShareRequests200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

