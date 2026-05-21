# \UsersAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](UsersAPI.md#CreateServiceAccount) | **Post** /users/service-accounts | 
[**DeleteUserInvitation**](UsersAPI.md#DeleteUserInvitation) | **Delete** /user-invitations/{email} | Delete a pending user invitation by email
[**GetAllUsers**](UsersAPI.md#GetAllUsers) | **Get** /users | 
[**GetAllUsersV2**](UsersAPI.md#GetAllUsersV2) | **Get** /v2/users | 
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{id} | 
[**ListUserInvitations**](UsersAPI.md#ListUserInvitations) | **Get** /user-invitations | List pending user invitations
[**SendUserInvitation**](UsersAPI.md#SendUserInvitation) | **Post** /organizations/{id}/invite | Send an invitation to join an organization



## CreateServiceAccount

> CreateServiceAccount201Response CreateServiceAccount(ctx).CreateServiceAccountRequest(createServiceAccountRequest).Execute()





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
	createServiceAccountRequest := *openapiclient.NewCreateServiceAccountRequest("api-service-prod") // CreateServiceAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.CreateServiceAccount(context.Background()).CreateServiceAccountRequest(createServiceAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateServiceAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceAccount`: CreateServiceAccount201Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceAccountRequest** | [**CreateServiceAccountRequest**](CreateServiceAccountRequest.md) |  | 

### Return type

[**CreateServiceAccount201Response**](CreateServiceAccount201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserInvitation

> string DeleteUserInvitation(ctx, email).Execute()

Delete a pending user invitation by email



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
	email := "email_example" // string | Email address of the invited user to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.DeleteUserInvitation(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUserInvitation`: string
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.DeleteUserInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address of the invited user to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsers

> []GetAllUsers200ResponseInner GetAllUsers(ctx).Count(count).Page(page).Execute()





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
	count := int32(56) // int32 | The numbers of users being returned (optional)
	page := int32(56) // int32 | The page number of users being returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetAllUsers(context.Background()).Count(count).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetAllUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUsers`: []GetAllUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetAllUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **int32** | The numbers of users being returned | 
 **page** | **int32** | The page number of users being returned | 

### Return type

[**[]GetAllUsers200ResponseInner**](GetAllUsers200ResponseInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsersV2

> GetAllUsersV2200Response GetAllUsersV2(ctx).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()





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
	search := "search_example" // string | Search term to filter users by username, email, or full name (max 255) (optional)
	sort := "sort_example" // string | Sort order for results. Use format `field_asc` or `field_desc`. Supported fields: username, full_name, email, created_at, last_visit. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetAllUsersV2(context.Background()).Page(page).PageSize(pageSize).Search(search).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetAllUsersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUsersV2`: GetAllUsersV2200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetAllUsersV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUsersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (1-indexed, defaults to 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (defaults to 10, max 50) | [default to 10]
 **search** | **string** | Search term to filter users by username, email, or full name (max 255) | 
 **sort** | **string** | Sort order for results. Use format &#x60;field_asc&#x60; or &#x60;field_desc&#x60;. Supported fields: username, full_name, email, created_at, last_visit. | 

### Return type

[**GetAllUsersV2200Response**](GetAllUsersV2200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> map[string]interface{} GetUser(ctx, id).Execute()





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
	id := int32(56) // int32 | The user's unique identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The user&#39;s unique identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserInvitations

> []UserInvitation ListUserInvitations(ctx).Execute()

List pending user invitations



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
	resp, r, err := apiClient.UsersAPI.ListUserInvitations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserInvitations`: []UserInvitation
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserInvitations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUserInvitationsRequest struct via the builder pattern


### Return type

[**[]UserInvitation**](UserInvitation.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendUserInvitation

> []UserInvitationCreateResult SendUserInvitation(ctx, id).SendUserInvitationRequest(sendUserInvitationRequest).Execute()

Send an invitation to join an organization



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
	id := int32(56) // int32 | Organization ID.
	sendUserInvitationRequest := *openapiclient.NewSendUserInvitationRequest([]string{"Emails_example"}) // SendUserInvitationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SendUserInvitation(context.Background(), id).SendUserInvitationRequest(sendUserInvitationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SendUserInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendUserInvitation`: []UserInvitationCreateResult
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SendUserInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Organization ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendUserInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sendUserInvitationRequest** | [**SendUserInvitationRequest**](SendUserInvitationRequest.md) |  | 

### Return type

[**[]UserInvitationCreateResult**](UserInvitationCreateResult.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

