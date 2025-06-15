# \ReleaseGroupsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReleaseGroup**](ReleaseGroupsAPI.md#CreateReleaseGroup) | **Post** /project_group | 
[**CreateReleaseGroupReleases**](ReleaseGroupsAPI.md#CreateReleaseGroupReleases) | **Post** /project_group/{groupId}/release | 
[**DeleteReleaseGroupById**](ReleaseGroupsAPI.md#DeleteReleaseGroupById) | **Delete** /project_group/{groupId} | 
[**DeleteReleaseGroupReleaseById**](ReleaseGroupsAPI.md#DeleteReleaseGroupReleaseById) | **Delete** /project_group/{groupId}/release/{projectGroupReleaseId} | 
[**DeprecatedGetReleaseGroupReleases**](ReleaseGroupsAPI.md#DeprecatedGetReleaseGroupReleases) | **Get** /project_group/{groupId}/release | 
[**GetAllProjectsInReleaseGroup**](ReleaseGroupsAPI.md#GetAllProjectsInReleaseGroup) | **Get** /project_group/{groupId}/all_projects | 
[**GetAllReleaseGroupTeams**](ReleaseGroupsAPI.md#GetAllReleaseGroupTeams) | **Get** /project_group/{groupId}/teams | 
[**GetReleaseGroupAttributionReportStatus**](ReleaseGroupsAPI.md#GetReleaseGroupAttributionReportStatus) | **Get** /project_group/attribution/{taskId} | 
[**GetReleaseGroupById**](ReleaseGroupsAPI.md#GetReleaseGroupById) | **Get** /project_group/{groupId} | 
[**GetReleaseGroupReleaseById**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseById) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId} | 
[**GetReleaseGroupReleaseLicenses**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseLicenses) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId}/licenses | 
[**GetReleaseGroupReleaseObligations**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseObligations) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId}/obligations | 
[**GetReleaseGroupReleaseRevisions**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseRevisions) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId}/revisions | 
[**GetReleaseGroupReleaseScans**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseScans) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId}/scans | 
[**GetReleaseGroupReleaseSummary**](ReleaseGroupsAPI.md#GetReleaseGroupReleaseSummary) | **Get** /project_group/{groupId}/release/{projectGroupReleaseId}/summary | 
[**GetReleaseGroupReleases**](ReleaseGroupsAPI.md#GetReleaseGroupReleases) | **Get** /project_group/{groupId}/releases | 
[**QueueReleaseGroupAttributionReport**](ReleaseGroupsAPI.md#QueueReleaseGroupAttributionReport) | **Post** /project_group/{groupId}/release/{releaseId}/attribution/{format} | 
[**UpdateReleaseGroupById**](ReleaseGroupsAPI.md#UpdateReleaseGroupById) | **Put** /project_group/{groupId} | 
[**UpdateReleaseGroupReleaseById**](ReleaseGroupsAPI.md#UpdateReleaseGroupReleaseById) | **Put** /project_group/{groupId}/release/{projectGroupReleaseId} | 



## CreateReleaseGroup

> CreateReleaseGroup200Response CreateReleaseGroup(ctx).CreateReleaseGroupRequest(createReleaseGroupRequest).Execute()





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
	createReleaseGroupRequest := *openapiclient.NewCreateReleaseGroupRequest("Title_example", *openapiclient.NewCreateReleaseGroupRequestRelease()) // CreateReleaseGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.CreateReleaseGroup(context.Background()).CreateReleaseGroupRequest(createReleaseGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.CreateReleaseGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReleaseGroup`: CreateReleaseGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.CreateReleaseGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReleaseGroupRequest** | [**CreateReleaseGroupRequest**](CreateReleaseGroupRequest.md) |  | 

### Return type

[**CreateReleaseGroup200Response**](CreateReleaseGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReleaseGroupReleases

> CreateReleaseGroup200Response CreateReleaseGroupReleases(ctx, groupId).CreateReleaseGroupReleasesRequest(createReleaseGroupReleasesRequest).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	createReleaseGroupReleasesRequest := *openapiclient.NewCreateReleaseGroupReleasesRequest() // CreateReleaseGroupReleasesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.CreateReleaseGroupReleases(context.Background(), groupId).CreateReleaseGroupReleasesRequest(createReleaseGroupReleasesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.CreateReleaseGroupReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReleaseGroupReleases`: CreateReleaseGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.CreateReleaseGroupReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReleaseGroupReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReleaseGroupReleasesRequest** | [**CreateReleaseGroupReleasesRequest**](CreateReleaseGroupReleasesRequest.md) |  | 

### Return type

[**CreateReleaseGroup200Response**](CreateReleaseGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReleaseGroupById

> DeleteReleaseGroupById(ctx, groupId).Execute()





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
	groupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReleaseGroupsAPI.DeleteReleaseGroupById(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.DeleteReleaseGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteReleaseGroupReleaseById

> DeleteReleaseGroupReleaseById(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReleaseGroupsAPI.DeleteReleaseGroupReleaseById(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.DeleteReleaseGroupReleaseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReleaseGroupReleaseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DeprecatedGetReleaseGroupReleases

> []DeprecatedGetReleaseGroupReleases200ResponseInner DeprecatedGetReleaseGroupReleases(ctx, groupId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.DeprecatedGetReleaseGroupReleases(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.DeprecatedGetReleaseGroupReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeprecatedGetReleaseGroupReleases`: []DeprecatedGetReleaseGroupReleases200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.DeprecatedGetReleaseGroupReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedGetReleaseGroupReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DeprecatedGetReleaseGroupReleases200ResponseInner**](DeprecatedGetReleaseGroupReleases200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectsInReleaseGroup

> []string GetAllProjectsInReleaseGroup(ctx, groupId).Execute()





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
	groupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetAllProjectsInReleaseGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetAllProjectsInReleaseGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectsInReleaseGroup`: []string
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetAllProjectsInReleaseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectsInReleaseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReleaseGroupTeams

> GetAllReleaseGroupTeams200Response GetAllReleaseGroupTeams(ctx, groupId).Execute()





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
	groupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetAllReleaseGroupTeams(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetAllReleaseGroupTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllReleaseGroupTeams`: GetAllReleaseGroupTeams200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetAllReleaseGroupTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReleaseGroupTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllReleaseGroupTeams200Response**](GetAllReleaseGroupTeams200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupAttributionReportStatus

> GetReleaseGroupAttributionReportStatus200Response GetReleaseGroupAttributionReportStatus(ctx, taskId).Execute()





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
	taskId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupAttributionReportStatus(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupAttributionReportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupAttributionReportStatus`: GetReleaseGroupAttributionReportStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupAttributionReportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupAttributionReportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReleaseGroupAttributionReportStatus200Response**](GetReleaseGroupAttributionReportStatus200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupById

> GetReleaseGroupById200Response GetReleaseGroupById(ctx, groupId).Execute()





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
	groupId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupById(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupById`: GetReleaseGroupById200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReleaseGroupById200Response**](GetReleaseGroupById200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseById

> []GetReleaseGroupReleaseById200ResponseInner GetReleaseGroupReleaseById(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseById(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseById`: []GetReleaseGroupReleaseById200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetReleaseGroupReleaseById200ResponseInner**](GetReleaseGroupReleaseById200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseLicenses

> map[string]GetReleaseGroupReleaseLicenses200ResponseValue GetReleaseGroupReleaseLicenses(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseLicenses(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseLicenses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseLicenses`: map[string]GetReleaseGroupReleaseLicenses200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]GetReleaseGroupReleaseLicenses200ResponseValue**](GetReleaseGroupReleaseLicenses200ResponseValue.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseObligations

> map[string][]GetReleaseGroupReleaseObligations200ResponseValueInner GetReleaseGroupReleaseObligations(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseObligations(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseObligations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseObligations`: map[string][]GetReleaseGroupReleaseObligations200ResponseValueInner
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseObligations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseObligationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string][]GetReleaseGroupReleaseObligations200ResponseValueInner**](array.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseRevisions

> []GetReleaseGroupReleaseRevisions200ResponseInner GetReleaseGroupReleaseRevisions(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseRevisions(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseRevisions`: []GetReleaseGroupReleaseRevisions200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetReleaseGroupReleaseRevisions200ResponseInner**](GetReleaseGroupReleaseRevisions200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseScans

> GetReleaseGroupReleaseScans200Response GetReleaseGroupReleaseScans(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseScans(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseScans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseScans`: GetReleaseGroupReleaseScans200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseScans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseScansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetReleaseGroupReleaseScans200Response**](GetReleaseGroupReleaseScans200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleaseSummary

> GetReleaseGroupReleaseSummary200Response GetReleaseGroupReleaseSummary(ctx, groupId, projectGroupReleaseId).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleaseSummary(context.Background(), groupId, projectGroupReleaseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleaseSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleaseSummary`: GetReleaseGroupReleaseSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleaseSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleaseSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetReleaseGroupReleaseSummary200Response**](GetReleaseGroupReleaseSummary200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReleaseGroupReleases

> GetReleaseGroupReleases200Response GetReleaseGroupReleases(ctx, groupId).Count(count).Page(page).Search(search).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	count := int32(56) // int32 |  (optional) (default to 10)
	page := int32(56) // int32 |  (optional) (default to 1)
	search := "search_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.GetReleaseGroupReleases(context.Background(), groupId).Count(count).Page(page).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.GetReleaseGroupReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReleaseGroupReleases`: GetReleaseGroupReleases200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.GetReleaseGroupReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseGroupReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** |  | [default to 10]
 **page** | **int32** |  | [default to 1]
 **search** | **string** |  | 

### Return type

[**GetReleaseGroupReleases200Response**](GetReleaseGroupReleases200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueueReleaseGroupAttributionReport

> QueueReleaseGroupAttributionReport200Response QueueReleaseGroupAttributionReport(ctx, groupId, releaseId, format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).DependencyInfoOptions(dependencyInfoOptions).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	releaseId := int32(56) // int32 | 
	format := "format_example" // string | 
	includeDeepDependencies := true // bool |  (optional)
	includeDirectDependencies := true // bool |  (optional)
	includeLicenseList := true // bool |  (optional)
	includeLicenseScan := true // bool |  (optional)
	includeProjectLicense := true // bool |  (optional)
	includeCopyrightList := true // bool |  (optional)
	includeFileMatches := true // bool |  (optional)
	includeOpenVulnerabilities := true // bool |  (optional)
	includeClosedVulnerabilities := true // bool |  (optional)
	includeDependencySummary := true // bool |  (optional)
	includeLicenseHeaders := true // bool |  (optional)
	dependencyInfoOptions := []string{"DependencyInfoOptions_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.QueueReleaseGroupAttributionReport(context.Background(), groupId, releaseId, format).IncludeDeepDependencies(includeDeepDependencies).IncludeDirectDependencies(includeDirectDependencies).IncludeLicenseList(includeLicenseList).IncludeLicenseScan(includeLicenseScan).IncludeProjectLicense(includeProjectLicense).IncludeCopyrightList(includeCopyrightList).IncludeFileMatches(includeFileMatches).IncludeOpenVulnerabilities(includeOpenVulnerabilities).IncludeClosedVulnerabilities(includeClosedVulnerabilities).IncludeDependencySummary(includeDependencySummary).IncludeLicenseHeaders(includeLicenseHeaders).DependencyInfoOptions(dependencyInfoOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.QueueReleaseGroupAttributionReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueueReleaseGroupAttributionReport`: QueueReleaseGroupAttributionReport200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.QueueReleaseGroupAttributionReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**releaseId** | **int32** |  | 
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueueReleaseGroupAttributionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeDeepDependencies** | **bool** |  | 
 **includeDirectDependencies** | **bool** |  | 
 **includeLicenseList** | **bool** |  | 
 **includeLicenseScan** | **bool** |  | 
 **includeProjectLicense** | **bool** |  | 
 **includeCopyrightList** | **bool** |  | 
 **includeFileMatches** | **bool** |  | 
 **includeOpenVulnerabilities** | **bool** |  | 
 **includeClosedVulnerabilities** | **bool** |  | 
 **includeDependencySummary** | **bool** |  | 
 **includeLicenseHeaders** | **bool** |  | 
 **dependencyInfoOptions** | **[]string** |  | 

### Return type

[**QueueReleaseGroupAttributionReport200Response**](QueueReleaseGroupAttributionReport200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseGroupById

> CreateReleaseGroup200Response UpdateReleaseGroupById(ctx, groupId).UpdateReleaseGroupByIdRequest(updateReleaseGroupByIdRequest).Execute()





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
	groupId := int32(56) // int32 | 
	updateReleaseGroupByIdRequest := *openapiclient.NewUpdateReleaseGroupByIdRequest() // UpdateReleaseGroupByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.UpdateReleaseGroupById(context.Background(), groupId).UpdateReleaseGroupByIdRequest(updateReleaseGroupByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.UpdateReleaseGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReleaseGroupById`: CreateReleaseGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.UpdateReleaseGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReleaseGroupByIdRequest** | [**UpdateReleaseGroupByIdRequest**](UpdateReleaseGroupByIdRequest.md) |  | 

### Return type

[**CreateReleaseGroup200Response**](CreateReleaseGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseGroupReleaseById

> CreateReleaseGroup200Response UpdateReleaseGroupReleaseById(ctx, groupId, projectGroupReleaseId).UpdateReleaseGroupReleaseByIdRequest(updateReleaseGroupReleaseByIdRequest).Execute()





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
	groupId := int32(56) // int32 | The ID of the release group
	projectGroupReleaseId := int32(56) // int32 | The ID of the release
	updateReleaseGroupReleaseByIdRequest := *openapiclient.NewUpdateReleaseGroupReleaseByIdRequest("Title_example") // UpdateReleaseGroupReleaseByIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseGroupsAPI.UpdateReleaseGroupReleaseById(context.Background(), groupId, projectGroupReleaseId).UpdateReleaseGroupReleaseByIdRequest(updateReleaseGroupReleaseByIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseGroupsAPI.UpdateReleaseGroupReleaseById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReleaseGroupReleaseById`: CreateReleaseGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseGroupsAPI.UpdateReleaseGroupReleaseById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | The ID of the release group | 
**projectGroupReleaseId** | **int32** | The ID of the release | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseGroupReleaseByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateReleaseGroupReleaseByIdRequest** | [**UpdateReleaseGroupReleaseByIdRequest**](UpdateReleaseGroupReleaseByIdRequest.md) |  | 

### Return type

[**CreateReleaseGroup200Response**](CreateReleaseGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

