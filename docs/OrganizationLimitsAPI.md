# \OrganizationLimitsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationContributorLimits**](OrganizationLimitsAPI.md#GetOrganizationContributorLimits) | **Get** /organizations/{id}/limits/contributors | 
[**GetOrganizationReleaseGroupLimits**](OrganizationLimitsAPI.md#GetOrganizationReleaseGroupLimits) | **Get** /organizations/{id}/limits/release-groups | 



## GetOrganizationContributorLimits

> GetOrganizationContributorLimits200Response GetOrganizationContributorLimits(ctx, id).Execute()





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
	id := int32(56) // int32 | The organization ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLimitsAPI.GetOrganizationContributorLimits(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLimitsAPI.GetOrganizationContributorLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationContributorLimits`: GetOrganizationContributorLimits200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLimitsAPI.GetOrganizationContributorLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The organization ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationContributorLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationContributorLimits200Response**](GetOrganizationContributorLimits200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationReleaseGroupLimits

> GetOrganizationContributorLimits200Response GetOrganizationReleaseGroupLimits(ctx, id).Execute()





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
	id := int32(56) // int32 | The organization ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationLimitsAPI.GetOrganizationReleaseGroupLimits(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationLimitsAPI.GetOrganizationReleaseGroupLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationReleaseGroupLimits`: GetOrganizationContributorLimits200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationLimitsAPI.GetOrganizationReleaseGroupLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The organization ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationReleaseGroupLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationContributorLimits200Response**](GetOrganizationContributorLimits200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

