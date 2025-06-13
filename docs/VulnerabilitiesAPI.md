# \VulnerabilitiesAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRemediationGuidance**](VulnerabilitiesAPI.md#GetRemediationGuidance) | **Get** /vulns/{vulnId}/revisions/{revisionId}/remediation-guidance | 



## GetRemediationGuidance

> GetIssue200ResponseOneOfAllOfRemediation GetRemediationGuidance(ctx, vulnId, revisionId).Execute()





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
	vulnId := "vulnId_example" // string | 
	revisionId := "revisionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VulnerabilitiesAPI.GetRemediationGuidance(context.Background(), vulnId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.GetRemediationGuidance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemediationGuidance`: GetIssue200ResponseOneOfAllOfRemediation
	fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.GetRemediationGuidance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vulnId** | **string** |  | 
**revisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemediationGuidanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIssue200ResponseOneOfAllOfRemediation**](GetIssue200ResponseOneOfAllOfRemediation.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

