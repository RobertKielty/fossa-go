# \TeamGroupsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTeamsToTeamGroup**](TeamGroupsAPI.md#AddTeamsToTeamGroup) | **Post** /teams/groups/{id}/teams | 
[**CreateTeamGroup**](TeamGroupsAPI.md#CreateTeamGroup) | **Post** /teams/groups | 
[**DeleteTeamGroup**](TeamGroupsAPI.md#DeleteTeamGroup) | **Delete** /teams/groups/{id} | 
[**GetTeamGroupById**](TeamGroupsAPI.md#GetTeamGroupById) | **Get** /teams/groups/{id} | 
[**GetTeamGroups**](TeamGroupsAPI.md#GetTeamGroups) | **Get** /teams/groups | 
[**RemoveTeamFromTeamGroup**](TeamGroupsAPI.md#RemoveTeamFromTeamGroup) | **Delete** /teams/groups/{id}/teams/{teamId} | 
[**UpdateTeamGroup**](TeamGroupsAPI.md#UpdateTeamGroup) | **Put** /teams/groups/{id} | 
[**UpdateTeamGroupUsers**](TeamGroupsAPI.md#UpdateTeamGroupUsers) | **Put** /teams/groups/{id}/users | 



## AddTeamsToTeamGroup

> AddTeamsToTeamGroup200Response AddTeamsToTeamGroup(ctx, id).AddTeamsToTeamGroupRequest(addTeamsToTeamGroupRequest).Execute()





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
	id := int32(56) // int32 | ID of the team group
	addTeamsToTeamGroupRequest := *openapiclient.NewAddTeamsToTeamGroupRequest([]int32{int32(123)}) // AddTeamsToTeamGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamGroupsAPI.AddTeamsToTeamGroup(context.Background(), id).AddTeamsToTeamGroupRequest(addTeamsToTeamGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.AddTeamsToTeamGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTeamsToTeamGroup`: AddTeamsToTeamGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.AddTeamsToTeamGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamsToTeamGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTeamsToTeamGroupRequest** | [**AddTeamsToTeamGroupRequest**](AddTeamsToTeamGroupRequest.md) |  | 

### Return type

[**AddTeamsToTeamGroup200Response**](AddTeamsToTeamGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamGroup

> CreateTeamGroup200Response CreateTeamGroup(ctx).CreateTeamGroupRequest(createTeamGroupRequest).Execute()





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
	createTeamGroupRequest := *openapiclient.NewCreateTeamGroupRequest("Engineering Team Group", int32(2)) // CreateTeamGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamGroupsAPI.CreateTeamGroup(context.Background()).CreateTeamGroupRequest(createTeamGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.CreateTeamGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTeamGroup`: CreateTeamGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.CreateTeamGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTeamGroupRequest** | [**CreateTeamGroupRequest**](CreateTeamGroupRequest.md) |  | 

### Return type

[**CreateTeamGroup200Response**](CreateTeamGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTeamGroup

> DeleteTeamGroup(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the team group to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamGroupsAPI.DeleteTeamGroup(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.DeleteTeamGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTeamGroupRequest struct via the builder pattern


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


## GetTeamGroupById

> GetTeamGroupById200Response GetTeamGroupById(ctx, id).Execute()





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
	id := int32(56) // int32 | ID of the team group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamGroupsAPI.GetTeamGroupById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.GetTeamGroupById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamGroupById`: GetTeamGroupById200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.GetTeamGroupById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamGroupByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetTeamGroupById200Response**](GetTeamGroupById200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamGroups

> []GetTeamGroups200ResponseInner GetTeamGroups(ctx).Execute()





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
	resp, r, err := apiClient.TeamGroupsAPI.GetTeamGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.GetTeamGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamGroups`: []GetTeamGroups200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.GetTeamGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamGroupsRequest struct via the builder pattern


### Return type

[**[]GetTeamGroups200ResponseInner**](GetTeamGroups200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTeamFromTeamGroup

> RemoveTeamFromTeamGroup(ctx, id, teamId).Execute()





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
	id := int32(56) // int32 | ID of the team group
	teamId := int32(56) // int32 | ID of the team to remove from the team group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TeamGroupsAPI.RemoveTeamFromTeamGroup(context.Background(), id, teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.RemoveTeamFromTeamGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group | 
**teamId** | **int32** | ID of the team to remove from the team group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTeamFromTeamGroupRequest struct via the builder pattern


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


## UpdateTeamGroup

> UpdateTeamGroup200Response UpdateTeamGroup(ctx, id).UpdateTeamGroupRequest(updateTeamGroupRequest).Execute()





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
	id := int32(56) // int32 | ID of the team group to update
	updateTeamGroupRequest := *openapiclient.NewUpdateTeamGroupRequest("Updated Engineering Team Group", int32(3)) // UpdateTeamGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamGroupsAPI.UpdateTeamGroup(context.Background(), id).UpdateTeamGroupRequest(updateTeamGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.UpdateTeamGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamGroup`: UpdateTeamGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.UpdateTeamGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamGroupRequest** | [**UpdateTeamGroupRequest**](UpdateTeamGroupRequest.md) |  | 

### Return type

[**UpdateTeamGroup200Response**](UpdateTeamGroup200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamGroupUsers

> UpdateTeamGroupUsers200Response UpdateTeamGroupUsers(ctx, id).UpdateTeamGroupUsersRequest(updateTeamGroupUsersRequest).Execute()





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
	id := int32(56) // int32 | ID of the team group
	updateTeamGroupUsersRequest := *openapiclient.NewUpdateTeamGroupUsersRequest("add", []openapiclient.UpdateTeamUsersRequestUsersInner{*openapiclient.NewUpdateTeamUsersRequestUsersInner(int32(123))}) // UpdateTeamGroupUsersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamGroupsAPI.UpdateTeamGroupUsers(context.Background(), id).UpdateTeamGroupUsersRequest(updateTeamGroupUsersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamGroupsAPI.UpdateTeamGroupUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTeamGroupUsers`: UpdateTeamGroupUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `TeamGroupsAPI.UpdateTeamGroupUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the team group | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamGroupUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTeamGroupUsersRequest** | [**UpdateTeamGroupUsersRequest**](UpdateTeamGroupUsersRequest.md) |  | 

### Return type

[**UpdateTeamGroupUsers200Response**](UpdateTeamGroupUsers200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

