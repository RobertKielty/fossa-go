# \TeamsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAddableTeamProjectsAndReleaseGroups**](TeamsAPI.md#GetAddableTeamProjectsAndReleaseGroups) | **Get** /teams/{id}/addable-projects-and-release-groups | 
[**GetAddableTeamUsers**](TeamsAPI.md#GetAddableTeamUsers) | **Get** /teams/{id}/members/addable | 
[**GetAllTeams**](TeamsAPI.md#GetAllTeams) | **Get** /teams | 
[**GetTeamByIdV2**](TeamsAPI.md#GetTeamByIdV2) | **Get** /v2/teams/{id} | 
[**GetTeamMembers**](TeamsAPI.md#GetTeamMembers) | **Get** /teams/{id}/members | 
[**GetTeamProjects**](TeamsAPI.md#GetTeamProjects) | **Get** /teams/{id}/projects | 
[**GetTeamReleaseGroups**](TeamsAPI.md#GetTeamReleaseGroups) | **Get** /teams/{id}/release-groups | 



## GetAddableTeamProjectsAndReleaseGroups

> GetAddableTeamProjectsAndReleaseGroups200Response GetAddableTeamProjectsAndReleaseGroups(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()





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
	id := int32(56) // int32 | ID of the team
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter projects and Release Groups by title (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetAddableTeamProjectsAndReleaseGroups(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetAddableTeamProjectsAndReleaseGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddableTeamProjectsAndReleaseGroups`: GetAddableTeamProjectsAndReleaseGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetAddableTeamProjectsAndReleaseGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddableTeamProjectsAndReleaseGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter projects and Release Groups by title | 

### Return type

[**GetAddableTeamProjectsAndReleaseGroups200Response**](GetAddableTeamProjectsAndReleaseGroups200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddableTeamUsers

> GetAddableTeamUsers200Response GetAddableTeamUsers(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()





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
	id := int32(56) // int32 | ID of the team
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter users by username (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetAddableTeamUsers(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetAddableTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddableTeamUsers`: GetAddableTeamUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetAddableTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddableTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter users by username | 

### Return type

[**GetAddableTeamUsers200Response**](GetAddableTeamUsers200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTeams

> []GetAllTeams200ResponseInner GetAllTeams(ctx).Execute()





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
	resp, r, err := apiClient.TeamsAPI.GetAllTeams(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetAllTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTeams`: []GetAllTeams200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetAllTeams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTeamsRequest struct via the builder pattern


### Return type

[**[]GetAllTeams200ResponseInner**](GetAllTeams200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamByIdV2

> GetTeamByIdV2200Response GetTeamByIdV2(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the team

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamByIdV2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamByIdV2`: GetTeamByIdV2200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTeamByIdV2200Response**](GetTeamByIdV2200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamMembers

> GetTeamMembers200Response GetTeamMembers(ctx, id).Page(page).PageSize(pageSize).Search(search).Execute()





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
	id := int32(56) // int32 | ID of the team
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter team members by username or email (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamMembers(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamMembers`: GetTeamMembers200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter team members by username or email | 

### Return type

[**GetTeamMembers200Response**](GetTeamMembers200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamProjects

> GetTeamProjects200Response GetTeamProjects(ctx, id).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()





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
	id := int32(56) // int32 | ID of the team
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter projects by title (optional)
	sort := "sort_example" // string | Sort order for the projects by title (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamProjects(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamProjects`: GetTeamProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter projects by title | 
 **sort** | **string** | Sort order for the projects by title | 

### Return type

[**GetTeamProjects200Response**](GetTeamProjects200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamReleaseGroups

> GetTeamReleaseGroups200Response GetTeamReleaseGroups(ctx, id).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()





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
	id := int32(56) // int32 | ID of the team
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter Release Groups by title (optional)
	sort := "sort_example" // string | Sort order for the Release Groups by title (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetTeamReleaseGroups(context.Background(), id).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetTeamReleaseGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamReleaseGroups`: GetTeamReleaseGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetTeamReleaseGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamReleaseGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter Release Groups by title | 
 **sort** | **string** | Sort order for the Release Groups by title | 

### Return type

[**GetTeamReleaseGroups200Response**](GetTeamReleaseGroups200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

