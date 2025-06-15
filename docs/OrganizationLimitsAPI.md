# \OrganizationLimitsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationLimits**](OrganizationLimitsAPI.md#GetOrganizationLimits) | **Get** /organizations/{id}/limits/{resource} | 



## GetOrganizationLimits

> GetOrganizationLimits200Response GetOrganizationLimits(ctx, id, resource).Execute()





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
	id := float32(8.14) // float32 | The organization ID.
	resource := "resource_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLimitsAPI.GetOrganizationLimits(context.Background(), id, resource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLimitsAPI.GetOrganizationLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationLimits`: GetOrganizationLimits200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLimitsAPI.GetOrganizationLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The organization ID. | 
**resource** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationLimits200Response**](GetOrganizationLimits200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

