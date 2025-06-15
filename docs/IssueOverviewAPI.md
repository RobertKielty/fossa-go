# \IssueOverviewAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportIssueOverviewCsv**](IssueOverviewAPI.md#ExportIssueOverviewCsv) | **Post** /issue_counts/export | 
[**GetIssueCounts**](IssueOverviewAPI.md#GetIssueCounts) | **Get** /issue_counts | 



## ExportIssueOverviewCsv

> ExportIssueOverviewCsv200Response ExportIssueOverviewCsv(ctx).Start(start).End(end).Labels(labels).Category(category).ProjectId(projectId).TeamId(teamId).Execute()





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
	start := time.Now() // time.Time | The start date to pull Issue Overview data for. Defaults to 30 days before the end date. (optional)
	end := time.Now() // time.Time | The start date to pull Issue Overview data for. Defaults to the current date and time. (optional)
	labels := []int32{int32(123)} // []int32 | The project labels to filter Issue Overview data to. (optional)
	category := "category_example" // string | The issue category to filter Issue Overview data to. If no category is chosen, it will fetch data for all issue categories. (optional)
	projectId := "projectId_example" // string | The specific project ID to filter Issue Overview data to. (optional)
	teamId := []openapiclient.GetIssueCWEsTeamIdParameterInner{openapiclient.getIssueCWEs_teamId___parameter_inner{Float32: new(float32)}} // []GetIssueCWEsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueOverviewAPI.ExportIssueOverviewCsv(context.Background()).Start(start).End(end).Labels(labels).Category(category).ProjectId(projectId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueOverviewAPI.ExportIssueOverviewCsv``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportIssueOverviewCsv`: ExportIssueOverviewCsv200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueOverviewAPI.ExportIssueOverviewCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIssueOverviewCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** | The start date to pull Issue Overview data for. Defaults to 30 days before the end date. | 
 **end** | **time.Time** | The start date to pull Issue Overview data for. Defaults to the current date and time. | 
 **labels** | **[]int32** | The project labels to filter Issue Overview data to. | 
 **category** | **string** | The issue category to filter Issue Overview data to. If no category is chosen, it will fetch data for all issue categories. | 
 **projectId** | **string** | The specific project ID to filter Issue Overview data to. | 
 **teamId** | [**[]GetIssueCWEsTeamIdParameterInner**](GetIssueCWEsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 

### Return type

[**ExportIssueOverviewCsv200Response**](ExportIssueOverviewCsv200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueCounts

> GetIssueCounts200Response GetIssueCounts(ctx).Start(start).End(end).Labels(labels).Category(category).ProjectId(projectId).TeamId(teamId).Execute()





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
	start := time.Now() // time.Time | The start date to pull Issue Overview data for. Defaults to 30 days before the end date. (optional)
	end := time.Now() // time.Time | The start date to pull Issue Overview data for. Defaults to the current date and time. (optional)
	labels := []int32{int32(123)} // []int32 | The project labels to filter Issue Overview data to. (optional)
	category := "category_example" // string | The issue category to filter Issue Overview data to. If no category is chosen, it will fetch data for all issue categories. (optional)
	projectId := "projectId_example" // string | The specific project ID to filter Issue Overview data to. (optional)
	teamId := []openapiclient.GetIssueCWEsTeamIdParameterInner{openapiclient.getIssueCWEs_teamId___parameter_inner{Float32: new(float32)}} // []GetIssueCWEsTeamIdParameterInner | Filter by one or more team IDs. Providing \"null\" will return all unassigned projects. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueOverviewAPI.GetIssueCounts(context.Background()).Start(start).End(end).Labels(labels).Category(category).ProjectId(projectId).TeamId(teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueOverviewAPI.GetIssueCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueCounts`: GetIssueCounts200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueOverviewAPI.GetIssueCounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **time.Time** | The start date to pull Issue Overview data for. Defaults to 30 days before the end date. | 
 **end** | **time.Time** | The start date to pull Issue Overview data for. Defaults to the current date and time. | 
 **labels** | **[]int32** | The project labels to filter Issue Overview data to. | 
 **category** | **string** | The issue category to filter Issue Overview data to. If no category is chosen, it will fetch data for all issue categories. | 
 **projectId** | **string** | The specific project ID to filter Issue Overview data to. | 
 **teamId** | [**[]GetIssueCWEsTeamIdParameterInner**](GetIssueCWEsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | 

### Return type

[**GetIssueCounts200Response**](GetIssueCounts200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

