# \BuildsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuilds**](BuildsAPI.md#GetBuilds) | **Get** /builds | 
[**GetBuildsCount**](BuildsAPI.md#GetBuildsCount) | **Get** /counts/builds | 



## GetBuilds

> []GetBuilds200ResponseInner GetBuilds(ctx).Locator(locator).ProjectId(projectId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Page(page).Sort(sort).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | The revision locator to filter builds to (optional)
	projectId := "projectId_example" // string | The project locator to filter builds to (optional)
	startDate := time.Now() // time.Time | The start date to filter builds to (optional)
	endDate := time.Now() // time.Time | The end date to filter builds to (optional)
	pageSize := int32(56) // int32 | The number of builds to return (optional)
	page := int32(56) // int32 | The page number to return (optional)
	sort := "sort_example" // string | The sort order(s) to apply to the builds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.GetBuilds(context.Background()).Locator(locator).ProjectId(projectId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.GetBuilds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuilds`: []GetBuilds200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.GetBuilds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **string** | The revision locator to filter builds to | 
 **projectId** | **string** | The project locator to filter builds to | 
 **startDate** | **time.Time** | The start date to filter builds to | 
 **endDate** | **time.Time** | The end date to filter builds to | 
 **pageSize** | **int32** | The number of builds to return | 
 **page** | **int32** | The page number to return | 
 **sort** | **string** | The sort order(s) to apply to the builds | 

### Return type

[**[]GetBuilds200ResponseInner**](GetBuilds200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuildsCount

> float32 GetBuildsCount(ctx).Locator(locator).ProjectId(projectId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Page(page).Sort(sort).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	locator := "locator_example" // string | The revision locator to filter builds to (optional)
	projectId := "projectId_example" // string | The project locator to filter builds to (optional)
	startDate := time.Now() // time.Time | The start date to filter builds to (optional)
	endDate := time.Now() // time.Time | The end date to filter builds to (optional)
	pageSize := int32(56) // int32 | The number of builds to return (optional)
	page := int32(56) // int32 | The page number to return (optional)
	sort := "sort_example" // string | The sort order(s) to apply to the builds (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.GetBuildsCount(context.Background()).Locator(locator).ProjectId(projectId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Page(page).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.GetBuildsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBuildsCount`: float32
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.GetBuildsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locator** | **string** | The revision locator to filter builds to | 
 **projectId** | **string** | The project locator to filter builds to | 
 **startDate** | **time.Time** | The start date to filter builds to | 
 **endDate** | **time.Time** | The end date to filter builds to | 
 **pageSize** | **int32** | The number of builds to return | 
 **page** | **int32** | The page number to return | 
 **sort** | **string** | The sort order(s) to apply to the builds | 

### Return type

**float32**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

