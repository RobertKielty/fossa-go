# \AuditLogsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLogsAPI.md#GetAuditLogs) | **Get** /audit_logs | 
[**GetAuditLogsCount**](AuditLogsAPI.md#GetAuditLogsCount) | **Get** /count/audit_logs | 
[**GetAuditLogsExport**](AuditLogsAPI.md#GetAuditLogsExport) | **Post** /audit_logs/export | 



## GetAuditLogs

> []GetAuditLogs200ResponseInner GetAuditLogs(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDir(sortDir).StartDate(startDate).EndDate(endDate).ActingUserIds(actingUserIds).Actions(actions).Topics(topics).TopicActions(topicActions).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()





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
	offset := int32(56) // int32 | The number of records to skip (optional)
	limit := int32(56) // int32 | The number of records to return (optional)
	sortBy := "sortBy_example" // string | The field to sort by (defaults to createdAt) (optional)
	sortDir := "sortDir_example" // string | The direction to sort by (defaults to DESC) (optional)
	startDate := time.Now() // time.Time | The start date to filter audit logs to (optional)
	endDate := time.Now() // time.Time | The end date to filter audit logs to (optional)
	actingUserIds := []string{"Inner_example"} // []string | The acting user IDs to filter audit logs to (optional)
	actions := []string{"Inner_example"} // []string | The actions to filter audit logs to (optional)
	topics := []string{"Inner_example"} // []string | The topics to filter audit logs to (optional)
	topicActions := []string{"Inner_example"} // []string | The topic actions to filter audit logs to (optional)
	startingAfter := time.Now() // time.Time | The id to start after to filter audit logs to (optional)
	endingBefore := "endingBefore_example" // string | The id to end before to filter audit logs to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLogs(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDir(sortDir).StartDate(startDate).EndDate(endDate).ActingUserIds(actingUserIds).Actions(actions).Topics(topics).TopicActions(topicActions).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogs`: []GetAuditLogs200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of records to skip | 
 **limit** | **int32** | The number of records to return | 
 **sortBy** | **string** | The field to sort by (defaults to createdAt) | 
 **sortDir** | **string** | The direction to sort by (defaults to DESC) | 
 **startDate** | **time.Time** | The start date to filter audit logs to | 
 **endDate** | **time.Time** | The end date to filter audit logs to | 
 **actingUserIds** | **[]string** | The acting user IDs to filter audit logs to | 
 **actions** | **[]string** | The actions to filter audit logs to | 
 **topics** | **[]string** | The topics to filter audit logs to | 
 **topicActions** | **[]string** | The topic actions to filter audit logs to | 
 **startingAfter** | **time.Time** | The id to start after to filter audit logs to | 
 **endingBefore** | **string** | The id to end before to filter audit logs to | 

### Return type

[**[]GetAuditLogs200ResponseInner**](GetAuditLogs200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsCount

> int32 GetAuditLogsCount(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDir(sortDir).StartDate(startDate).EndDate(endDate).ActingUserIds(actingUserIds).Actions(actions).Topics(topics).TopicActions(topicActions).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()





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
	offset := int32(56) // int32 | The number of records to skip (optional)
	limit := int32(56) // int32 | The number of records to return (optional)
	sortBy := "sortBy_example" // string | The field to sort by (defaults to createdAt) (optional)
	sortDir := "sortDir_example" // string | The direction to sort by (defaults to DESC) (optional)
	startDate := time.Now() // time.Time | The start date to filter audit logs to (optional)
	endDate := time.Now() // time.Time | The end date to filter audit logs to (optional)
	actingUserIds := []string{"Inner_example"} // []string | The acting user IDs to filter audit logs to (optional)
	actions := []string{"Inner_example"} // []string | The actions to filter audit logs to (optional)
	topics := []string{"Inner_example"} // []string | The topics to filter audit logs to (optional)
	topicActions := []string{"Inner_example"} // []string | The topic actions to filter audit logs to (optional)
	startingAfter := time.Now() // time.Time | The id to start after to filter audit logs to (optional)
	endingBefore := "endingBefore_example" // string | The id to end before to filter audit logs to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLogsCount(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDir(sortDir).StartDate(startDate).EndDate(endDate).ActingUserIds(actingUserIds).Actions(actions).Topics(topics).TopicActions(topicActions).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsCount`: int32
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditLogsCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of records to skip | 
 **limit** | **int32** | The number of records to return | 
 **sortBy** | **string** | The field to sort by (defaults to createdAt) | 
 **sortDir** | **string** | The direction to sort by (defaults to DESC) | 
 **startDate** | **time.Time** | The start date to filter audit logs to | 
 **endDate** | **time.Time** | The end date to filter audit logs to | 
 **actingUserIds** | **[]string** | The acting user IDs to filter audit logs to | 
 **actions** | **[]string** | The actions to filter audit logs to | 
 **topics** | **[]string** | The topics to filter audit logs to | 
 **topicActions** | **[]string** | The topic actions to filter audit logs to | 
 **startingAfter** | **time.Time** | The id to start after to filter audit logs to | 
 **endingBefore** | **string** | The id to end before to filter audit logs to | 

### Return type

**int32**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuditLogsExport

> GetAuditLogsExport201Response GetAuditLogsExport(ctx).GetAuditLogsExportRequest(getAuditLogsExportRequest).Execute()





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
	getAuditLogsExportRequest := *openapiclient.NewGetAuditLogsExportRequest(time.Now(), time.Now()) // GetAuditLogsExportRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.GetAuditLogsExport(context.Background()).GetAuditLogsExportRequest(getAuditLogsExportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogsExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditLogsExport`: GetAuditLogsExport201Response
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.GetAuditLogsExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAuditLogsExportRequest** | [**GetAuditLogsExportRequest**](GetAuditLogsExportRequest.md) |  | 

### Return type

[**GetAuditLogsExport201Response**](GetAuditLogsExport201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

