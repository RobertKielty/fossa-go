# \JiraIntegrationSettingsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateJiraConfiguration**](JiraIntegrationSettingsAPI.md#CreateJiraConfiguration) | **Post** /jira | 
[**DeleteJiraConfiguration**](JiraIntegrationSettingsAPI.md#DeleteJiraConfiguration) | **Delete** /jira/{id} | 
[**GetJiraConfigurations**](JiraIntegrationSettingsAPI.md#GetJiraConfigurations) | **Get** /jira | 
[**PatchJiraConfiguration**](JiraIntegrationSettingsAPI.md#PatchJiraConfiguration) | **Patch** /jira/{id} | 



## CreateJiraConfiguration

> PatchJiraConfigurationRequest CreateJiraConfiguration(ctx).PatchJiraConfigurationRequest(patchJiraConfigurationRequest).Execute()





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
	patchJiraConfigurationRequest := *openapiclient.NewPatchJiraConfigurationRequest() // PatchJiraConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraIntegrationSettingsAPI.CreateJiraConfiguration(context.Background()).PatchJiraConfigurationRequest(patchJiraConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationSettingsAPI.CreateJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateJiraConfiguration`: PatchJiraConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationSettingsAPI.CreateJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchJiraConfigurationRequest** | [**PatchJiraConfigurationRequest**](PatchJiraConfigurationRequest.md) |  | 

### Return type

[**PatchJiraConfigurationRequest**](PatchJiraConfigurationRequest.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteJiraConfiguration

> DeleteJiraConfiguration200Response DeleteJiraConfiguration(ctx, id).Execute()





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
	resp, r, err := apiClient.JiraIntegrationSettingsAPI.DeleteJiraConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationSettingsAPI.DeleteJiraConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteJiraConfiguration`: DeleteJiraConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationSettingsAPI.DeleteJiraConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteJiraConfiguration200Response**](DeleteJiraConfiguration200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJiraConfigurations

> []PatchJiraConfigurationRequest GetJiraConfigurations(ctx).Execute()





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
	resp, r, err := apiClient.JiraIntegrationSettingsAPI.GetJiraConfigurations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationSettingsAPI.GetJiraConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJiraConfigurations`: []PatchJiraConfigurationRequest
	fmt.Fprintf(os.Stdout, "Response from `JiraIntegrationSettingsAPI.GetJiraConfigurations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJiraConfigurationsRequest struct via the builder pattern


### Return type

[**[]PatchJiraConfigurationRequest**](PatchJiraConfigurationRequest.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchJiraConfiguration

> PatchJiraConfiguration(ctx, id).PatchJiraConfigurationRequest(patchJiraConfigurationRequest).Execute()





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
	patchJiraConfigurationRequest := *openapiclient.NewPatchJiraConfigurationRequest() // PatchJiraConfigurationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.JiraIntegrationSettingsAPI.PatchJiraConfiguration(context.Background(), id).PatchJiraConfigurationRequest(patchJiraConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraIntegrationSettingsAPI.PatchJiraConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchJiraConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchJiraConfigurationRequest** | [**PatchJiraConfigurationRequest**](PatchJiraConfigurationRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

