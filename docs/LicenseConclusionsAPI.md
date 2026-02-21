# \LicenseConclusionsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLicenseConclusion**](LicenseConclusionsAPI.md#AddLicenseConclusion) | **Put** /license-conclusions/conclude | 
[**RemoveLicenseConclusion**](LicenseConclusionsAPI.md#RemoveLicenseConclusion) | **Put** /license-conclusions/unconclude | 



## AddLicenseConclusion

> AddLicenseConclusion201Response AddLicenseConclusion(ctx).AddLicenseConclusionRequest(addLicenseConclusionRequest).Execute()





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
	addLicenseConclusionRequest := *openapiclient.NewAddLicenseConclusionRequest("DependencyRevisionLocator_example", openapiclient.addLicenseConclusion_request_scope{AddLicenseConclusionRequestScopeOneOf: openapiclient.NewAddLicenseConclusionRequestScopeOneOf("Scope_example", "ProjectLocator_example")}, "LicenseId_example") // AddLicenseConclusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseConclusionsAPI.AddLicenseConclusion(context.Background()).AddLicenseConclusionRequest(addLicenseConclusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseConclusionsAPI.AddLicenseConclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLicenseConclusion`: AddLicenseConclusion201Response
	fmt.Fprintf(os.Stdout, "Response from `LicenseConclusionsAPI.AddLicenseConclusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLicenseConclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLicenseConclusionRequest** | [**AddLicenseConclusionRequest**](AddLicenseConclusionRequest.md) |  | 

### Return type

[**AddLicenseConclusion201Response**](AddLicenseConclusion201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLicenseConclusion

> AddLicenseConclusion201Response RemoveLicenseConclusion(ctx).AddLicenseConclusionRequest(addLicenseConclusionRequest).Execute()





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
	addLicenseConclusionRequest := *openapiclient.NewAddLicenseConclusionRequest("DependencyRevisionLocator_example", openapiclient.addLicenseConclusion_request_scope{AddLicenseConclusionRequestScopeOneOf: openapiclient.NewAddLicenseConclusionRequestScopeOneOf("Scope_example", "ProjectLocator_example")}, "LicenseId_example") // AddLicenseConclusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseConclusionsAPI.RemoveLicenseConclusion(context.Background()).AddLicenseConclusionRequest(addLicenseConclusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseConclusionsAPI.RemoveLicenseConclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveLicenseConclusion`: AddLicenseConclusion201Response
	fmt.Fprintf(os.Stdout, "Response from `LicenseConclusionsAPI.RemoveLicenseConclusion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLicenseConclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addLicenseConclusionRequest** | [**AddLicenseConclusionRequest**](AddLicenseConclusionRequest.md) |  | 

### Return type

[**AddLicenseConclusion201Response**](AddLicenseConclusion201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

