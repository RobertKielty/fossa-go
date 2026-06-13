# \TeamsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddReleaseGroupsToTeam**](TeamsAPI.md#AddReleaseGroupsToTeam) | **Post** /teams/{id}/release-groups | 
[**CreateTeam**](TeamsAPI.md#CreateTeam) | **Post** /teams | 
[**DeleteTeam**](TeamsAPI.md#DeleteTeam) | **Delete** /teams/{id} | 
[**GetAddableProjectsFromReleaseGroup**](TeamsAPI.md#GetAddableProjectsFromReleaseGroup) | **Get** /teams/{id}/release-groups/{releaseGroupId}/addable-projects | 
[**GetAddableTeamProjectsAndReleaseGroups**](TeamsAPI.md#GetAddableTeamProjectsAndReleaseGroups) | **Get** /teams/{id}/addable-projects-and-release-groups | 
[**GetAddableTeamUsers**](TeamsAPI.md#GetAddableTeamUsers) | **Get** /teams/{id}/members/addable | 
[**GetAllTeams**](TeamsAPI.md#GetAllTeams) | **Get** /teams | 
[**GetAllTeamsV2**](TeamsAPI.md#GetAllTeamsV2) | **Get** /v2/teams | 
[**GetTeamByIdV2**](TeamsAPI.md#GetTeamByIdV2) | **Get** /v2/teams/{id} | 
[**GetTeamMembers**](TeamsAPI.md#GetTeamMembers) | **Get** /teams/{id}/members | 
[**GetTeamProjects**](TeamsAPI.md#GetTeamProjects) | **Get** /teams/{id}/projects | 
[**GetTeamReleaseGroups**](TeamsAPI.md#GetTeamReleaseGroups) | **Get** /teams/{id}/release-groups | 
[**RemoveReleaseGroupsFromTeam**](TeamsAPI.md#RemoveReleaseGroupsFromTeam) | **Delete** /teams/{id}/release-groups | 
[**UpdateTeam**](TeamsAPI.md#UpdateTeam) | **Put** /teams/{id} | 
[**UpdateTeamProjects**](TeamsAPI.md#UpdateTeamProjects) | **Put** /teams/{id}/projects | 
[**UpdateTeamUsers**](TeamsAPI.md#UpdateTeamUsers) | **Put** /teams/{id}/users | 



## AddReleaseGroupsToTeam

> AddReleaseGroupsToTeam200Response AddReleaseGroupsToTeam(ctx, id).AddReleaseGroupsToTeamRequest(addReleaseGroupsToTeamRequest).Execute()





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
	addReleaseGroupsToTeamRequest := *openapiclient.NewAddReleaseGroupsToTeamRequest([]int32{int32(123)}) // AddReleaseGroupsToTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.AddReleaseGroupsToTeam(context.Background(), id).AddReleaseGroupsToTeamRequest(addReleaseGroupsToTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.AddReleaseGroupsToTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddReleaseGroupsToTeam`: AddReleaseGroupsToTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.AddReleaseGroupsToTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddReleaseGroupsToTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseGroupsToTeamRequest** | [**AddReleaseGroupsToTeamRequest**](AddReleaseGroupsToTeamRequest.md) |  | 

### Return type

[**AddReleaseGroupsToTeam200Response**](AddReleaseGroupsToTeam200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeam

> GetAllTeams200ResponseInner CreateTeam(ctx).CreateTeamRequest(createTeamRequest).Execute()





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
	createTeamRequest := *openapiclient.NewCreateTeamRequest("Engineering", int32(2)) // CreateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.CreateTeam(context.Background()).CreateTeamRequest(createTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.CreateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeam`: GetAllTeams200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.CreateTeam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeamRequest** | [**CreateTeamRequest**](CreateTeamRequest.md) |  | 

### Return type

[**GetAllTeams200ResponseInner**](GetAllTeams200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeam

> DeleteTeam(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the team to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamsAPI.DeleteTeam(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.DeleteTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamRequest struct via the builder pattern


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


## GetAddableProjectsFromReleaseGroup

> GetAddableProjectsFromReleaseGroup200Response GetAddableProjectsFromReleaseGroup(ctx, id, releaseGroupId).Execute()





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
	releaseGroupId := int32(56) // int32 | ID of the Release Group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetAddableProjectsFromReleaseGroup(context.Background(), id, releaseGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetAddableProjectsFromReleaseGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddableProjectsFromReleaseGroup`: GetAddableProjectsFromReleaseGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetAddableProjectsFromReleaseGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 
**releaseGroupId** | **int32** | ID of the Release Group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddableProjectsFromReleaseGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAddableProjectsFromReleaseGroup200Response**](GetAddableProjectsFromReleaseGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetAllTeamsV2

> GetAllTeamsV2200Response GetAllTeamsV2(ctx).Page(page).PageSize(pageSize).Search(search).Execute()





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
	page := int32(56) // int32 | Page number (1-indexed, defaults to 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (defaults to 10, max 50) (optional) (default to 10)
	search := "search_example" // string | Search term to filter teams by name (max 255) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.GetAllTeamsV2(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.GetAllTeamsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTeamsV2`: GetAllTeamsV2200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.GetAllTeamsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTeamsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter teams by name (max 255) | 

### Return type

[**GetAllTeamsV2200Response**](GetAllTeamsV2200Response.md)

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


## RemoveReleaseGroupsFromTeam

> AddReleaseGroupsToTeam200Response RemoveReleaseGroupsFromTeam(ctx, id).AddReleaseGroupsToTeamRequest(addReleaseGroupsToTeamRequest).Execute()





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
	addReleaseGroupsToTeamRequest := *openapiclient.NewAddReleaseGroupsToTeamRequest([]int32{int32(123)}) // AddReleaseGroupsToTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.RemoveReleaseGroupsFromTeam(context.Background(), id).AddReleaseGroupsToTeamRequest(addReleaseGroupsToTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.RemoveReleaseGroupsFromTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveReleaseGroupsFromTeam`: AddReleaseGroupsToTeam200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.RemoveReleaseGroupsFromTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReleaseGroupsFromTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addReleaseGroupsToTeamRequest** | [**AddReleaseGroupsToTeamRequest**](AddReleaseGroupsToTeamRequest.md) |  | 

### Return type

[**AddReleaseGroupsToTeam200Response**](AddReleaseGroupsToTeam200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> GetAllTeams200ResponseInner UpdateTeam(ctx, id).UpdateTeamRequest(updateTeamRequest).Execute()





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
	id := int32(56) // int32 | ID of the team to update
	updateTeamRequest := *openapiclient.NewUpdateTeamRequest() // UpdateTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.UpdateTeam(context.Background(), id).UpdateTeamRequest(updateTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.UpdateTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeam`: GetAllTeams200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamRequest** | [**UpdateTeamRequest**](UpdateTeamRequest.md) |  | 

### Return type

[**GetAllTeams200ResponseInner**](GetAllTeams200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamProjects

> UpdateTeamProjects200Response UpdateTeamProjects(ctx, id).UpdateTeamProjectsRequest(updateTeamProjectsRequest).Execute()





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
	updateTeamProjectsRequest := *openapiclient.NewUpdateTeamProjectsRequest("add") // UpdateTeamProjectsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.UpdateTeamProjects(context.Background(), id).UpdateTeamProjectsRequest(updateTeamProjectsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.UpdateTeamProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamProjects`: UpdateTeamProjects200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.UpdateTeamProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamProjectsRequest** | [**UpdateTeamProjectsRequest**](UpdateTeamProjectsRequest.md) |  | 

### Return type

[**UpdateTeamProjects200Response**](UpdateTeamProjects200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamUsers

> UpdateTeamUsers200Response UpdateTeamUsers(ctx, id).UpdateTeamUsersRequest(updateTeamUsersRequest).Execute()





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
	updateTeamUsersRequest := *openapiclient.NewUpdateTeamUsersRequest("add") // UpdateTeamUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsAPI.UpdateTeamUsers(context.Background(), id).UpdateTeamUsersRequest(updateTeamUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsAPI.UpdateTeamUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamUsers`: UpdateTeamUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamsAPI.UpdateTeamUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamUsersRequest** | [**UpdateTeamUsersRequest**](UpdateTeamUsersRequest.md) |  | 

### Return type

[**UpdateTeamUsers200Response**](UpdateTeamUsers200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

