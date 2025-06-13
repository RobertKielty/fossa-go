# \CLIAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCLI**](CLIAPI.md#GetOrganizationCLI) | **Get** /cli/organization | 



## GetOrganizationCLI

> GetOrganizationCLI200Response GetOrganizationCLI(ctx).Execute()





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
	resp, r, err := apiClient.CLIAPI.GetOrganizationCLI(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CLIAPI.GetOrganizationCLI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationCLI`: GetOrganizationCLI200Response
	fmt.Fprintf(os.Stdout, "Response from `CLIAPI.GetOrganizationCLI`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCLIRequest struct via the builder pattern


### Return type

[**GetOrganizationCLI200Response**](GetOrganizationCLI200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

