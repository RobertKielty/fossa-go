# \CustomRiskScoresAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomRiskScore**](CustomRiskScoresAPI.md#CreateCustomRiskScore) | **Post** /custom-risk-scores/{issueId} | Create a custom risk score
[**DeleteCustomRiskScore**](CustomRiskScoresAPI.md#DeleteCustomRiskScore) | **Delete** /custom-risk-scores/{issueId} | Delete a custom risk score
[**UpdateCustomRiskScore**](CustomRiskScoresAPI.md#UpdateCustomRiskScore) | **Patch** /custom-risk-scores/{issueId} | Update a custom risk score



## CreateCustomRiskScore

> CreateCustomRiskScore201Response CreateCustomRiskScore(ctx, issueId).ScopeType(scopeType).ScopeId(scopeId).CreateCustomRiskScoreRequest(createCustomRiskScoreRequest).Execute()

Create a custom risk score



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
	issueId := int32(56) // int32 | ID of the vulnerability issue
	scopeType := "scopeType_example" // string | Scope type for the custom risk score
	scopeId := "scopeId_example" // string | Project locator or release group ID
	createCustomRiskScoreRequest := *openapiclient.NewCreateCustomRiskScoreRequest(int32(75)) // CreateCustomRiskScoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRiskScoresAPI.CreateCustomRiskScore(context.Background(), issueId).ScopeType(scopeType).ScopeId(scopeId).CreateCustomRiskScoreRequest(createCustomRiskScoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRiskScoresAPI.CreateCustomRiskScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomRiskScore`: CreateCustomRiskScore201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomRiskScoresAPI.CreateCustomRiskScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | ID of the vulnerability issue | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeType** | **string** | Scope type for the custom risk score | 
 **scopeId** | **string** | Project locator or release group ID | 
 **createCustomRiskScoreRequest** | [**CreateCustomRiskScoreRequest**](CreateCustomRiskScoreRequest.md) |  | 

### Return type

[**CreateCustomRiskScore201Response**](CreateCustomRiskScore201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomRiskScore

> DeleteCustomRiskScore(ctx, issueId).ScopeType(scopeType).ScopeId(scopeId).Execute()

Delete a custom risk score



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
	issueId := int32(56) // int32 | ID of the vulnerability issue
	scopeType := "scopeType_example" // string | Scope type for the custom risk score
	scopeId := "scopeId_example" // string | Project locator or release group ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CustomRiskScoresAPI.DeleteCustomRiskScore(context.Background(), issueId).ScopeType(scopeType).ScopeId(scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRiskScoresAPI.DeleteCustomRiskScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | ID of the vulnerability issue | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeType** | **string** | Scope type for the custom risk score | 
 **scopeId** | **string** | Project locator or release group ID | 

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


## UpdateCustomRiskScore

> CreateCustomRiskScore201Response UpdateCustomRiskScore(ctx, issueId).ScopeType(scopeType).ScopeId(scopeId).CreateCustomRiskScoreRequest(createCustomRiskScoreRequest).Execute()

Update a custom risk score



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
	issueId := int32(56) // int32 | ID of the vulnerability issue
	scopeType := "scopeType_example" // string | Scope type for the custom risk score
	scopeId := "scopeId_example" // string | Project locator or release group ID
	createCustomRiskScoreRequest := *openapiclient.NewCreateCustomRiskScoreRequest(int32(75)) // CreateCustomRiskScoreRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomRiskScoresAPI.UpdateCustomRiskScore(context.Background(), issueId).ScopeType(scopeType).ScopeId(scopeId).CreateCustomRiskScoreRequest(createCustomRiskScoreRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomRiskScoresAPI.UpdateCustomRiskScore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomRiskScore`: CreateCustomRiskScore201Response
	fmt.Fprintf(os.Stdout, "Response from `CustomRiskScoresAPI.UpdateCustomRiskScore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueId** | **int32** | ID of the vulnerability issue | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomRiskScoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scopeType** | **string** | Scope type for the custom risk score | 
 **scopeId** | **string** | Project locator or release group ID | 
 **createCustomRiskScoreRequest** | [**CreateCustomRiskScoreRequest**](CreateCustomRiskScoreRequest.md) |  | 

### Return type

[**CreateCustomRiskScore201Response**](CreateCustomRiskScore201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

