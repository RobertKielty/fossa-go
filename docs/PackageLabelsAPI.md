# \PackageLabelsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePackageLabel**](PackageLabelsAPI.md#CreatePackageLabel) | **Post** /package-labels | 
[**CreatePackageLabelAssignments**](PackageLabelsAPI.md#CreatePackageLabelAssignments) | **Post** /package-label-assignments | 
[**DeletePackageLabelAssignments**](PackageLabelsAPI.md#DeletePackageLabelAssignments) | **Delete** /package-label-assignments | 
[**DeletePackageLabels**](PackageLabelsAPI.md#DeletePackageLabels) | **Delete** /package-labels | 
[**GetPackageLabelAssignments**](PackageLabelsAPI.md#GetPackageLabelAssignments) | **Get** /package-label-assignments | 
[**GetPackageLabels**](PackageLabelsAPI.md#GetPackageLabels) | **Get** /package-labels | 



## CreatePackageLabel

> GetPackageLabels200Response CreatePackageLabel(ctx).CreatePackageLabelRequest(createPackageLabelRequest).Execute()





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
	createPackageLabelRequest := *openapiclient.NewCreatePackageLabelRequest([]string{"Labels_example"}) // CreatePackageLabelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageLabelsAPI.CreatePackageLabel(context.Background()).CreatePackageLabelRequest(createPackageLabelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.CreatePackageLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePackageLabel`: GetPackageLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageLabelsAPI.CreatePackageLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackageLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPackageLabelRequest** | [**CreatePackageLabelRequest**](CreatePackageLabelRequest.md) |  | 

### Return type

[**GetPackageLabels200Response**](GetPackageLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePackageLabelAssignments

> GetPackageLabelAssignments200Response CreatePackageLabelAssignments(ctx).CreatePackageLabelAssignmentsRequest(createPackageLabelAssignmentsRequest).Execute()





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
	createPackageLabelAssignmentsRequest := *openapiclient.NewCreatePackageLabelAssignmentsRequest("PackageId_example") // CreatePackageLabelAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageLabelsAPI.CreatePackageLabelAssignments(context.Background()).CreatePackageLabelAssignmentsRequest(createPackageLabelAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.CreatePackageLabelAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePackageLabelAssignments`: GetPackageLabelAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageLabelsAPI.CreatePackageLabelAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePackageLabelAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPackageLabelAssignmentsRequest** | [**CreatePackageLabelAssignmentsRequest**](CreatePackageLabelAssignmentsRequest.md) |  | 

### Return type

[**GetPackageLabelAssignments200Response**](GetPackageLabelAssignments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePackageLabelAssignments

> DeletePackageLabelAssignments(ctx).DeletePackageLabelAssignmentsRequest(deletePackageLabelAssignmentsRequest).Execute()





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
	deletePackageLabelAssignmentsRequest := *openapiclient.NewDeletePackageLabelAssignmentsRequest() // DeletePackageLabelAssignmentsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageLabelsAPI.DeletePackageLabelAssignments(context.Background()).DeletePackageLabelAssignmentsRequest(deletePackageLabelAssignmentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.DeletePackageLabelAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackageLabelAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePackageLabelAssignmentsRequest** | [**DeletePackageLabelAssignmentsRequest**](DeletePackageLabelAssignmentsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePackageLabels

> DeletePackageLabels(ctx).DeletePackageLabelsRequest(deletePackageLabelsRequest).Execute()





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
	deletePackageLabelsRequest := *openapiclient.NewDeletePackageLabelsRequest([]int32{int32(123)}) // DeletePackageLabelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PackageLabelsAPI.DeletePackageLabels(context.Background()).DeletePackageLabelsRequest(deletePackageLabelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.DeletePackageLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePackageLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deletePackageLabelsRequest** | [**DeletePackageLabelsRequest**](DeletePackageLabelsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageLabelAssignments

> GetPackageLabelAssignments200Response GetPackageLabelAssignments(ctx).PackageId(packageId).PackageVersion(packageVersion).Scope(scope).ScopeId(scopeId).Execute()





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
	packageId := "npm+lodash" // string |  (optional)
	packageVersion := "4.15.0" // string |  (optional)
	scope := "project" // string |  (optional)
	scopeId := "custom+1/my-cli-project or custom+1/my-cli-project/$revision1" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageLabelsAPI.GetPackageLabelAssignments(context.Background()).PackageId(packageId).PackageVersion(packageVersion).Scope(scope).ScopeId(scopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.GetPackageLabelAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageLabelAssignments`: GetPackageLabelAssignments200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageLabelsAPI.GetPackageLabelAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageLabelAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageId** | **string** |  | 
 **packageVersion** | **string** |  | 
 **scope** | **string** |  | 
 **scopeId** | **string** |  | 

### Return type

[**GetPackageLabelAssignments200Response**](GetPackageLabelAssignments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackageLabels

> GetPackageLabels200Response GetPackageLabels(ctx).Execute()





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
	resp, r, err := apiClient.PackageLabelsAPI.GetPackageLabels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageLabelsAPI.GetPackageLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageLabels`: GetPackageLabels200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageLabelsAPI.GetPackageLabels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageLabelsRequest struct via the builder pattern


### Return type

[**GetPackageLabels200Response**](GetPackageLabels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

