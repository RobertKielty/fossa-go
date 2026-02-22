# \OIDCAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOIDCProvider**](OIDCAPI.md#CreateOIDCProvider) | **Post** /oidc/providers | 
[**CreateOIDCTrustRelationship**](OIDCAPI.md#CreateOIDCTrustRelationship) | **Post** /oidc/trust-relationships | 
[**DeleteOIDCProvider**](OIDCAPI.md#DeleteOIDCProvider) | **Delete** /oidc/providers/{id} | 
[**DeleteOIDCTrustRelationship**](OIDCAPI.md#DeleteOIDCTrustRelationship) | **Delete** /oidc/trust-relationships/{id} | 
[**ExchangeOIDCToken**](OIDCAPI.md#ExchangeOIDCToken) | **Post** /oidc/token-exchange | 
[**GetOIDCProvider**](OIDCAPI.md#GetOIDCProvider) | **Get** /oidc/providers/{id} | 
[**GetOIDCProviderAvailableServiceAccounts**](OIDCAPI.md#GetOIDCProviderAvailableServiceAccounts) | **Get** /oidc/providers/{id}/available-service-accounts | 
[**GetOIDCTrustRelationship**](OIDCAPI.md#GetOIDCTrustRelationship) | **Get** /oidc/trust-relationships/{id} | 
[**ListOIDCProviders**](OIDCAPI.md#ListOIDCProviders) | **Get** /oidc/providers | 
[**ListOIDCTrustRelationships**](OIDCAPI.md#ListOIDCTrustRelationships) | **Get** /oidc/trust-relationships | 
[**UpdateOIDCTrustRelationship**](OIDCAPI.md#UpdateOIDCTrustRelationship) | **Put** /oidc/trust-relationships/{id} | 



## CreateOIDCProvider

> CreateOIDCProvider201Response CreateOIDCProvider(ctx).CreateOIDCProviderRequest(createOIDCProviderRequest).Execute()





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
	createOIDCProviderRequest := openapiclient.createOIDCProvider_request{CreateOIDCProviderRequestOneOf: openapiclient.NewCreateOIDCProviderRequestOneOf("https://oidc.example.com", "team", int32(789))} // CreateOIDCProviderRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.CreateOIDCProvider(context.Background()).CreateOIDCProviderRequest(createOIDCProviderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.CreateOIDCProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOIDCProvider`: CreateOIDCProvider201Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.CreateOIDCProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOIDCProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOIDCProviderRequest** | [**CreateOIDCProviderRequest**](CreateOIDCProviderRequest.md) |  | 

### Return type

[**CreateOIDCProvider201Response**](CreateOIDCProvider201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOIDCTrustRelationship

> CreateOIDCTrustRelationship201Response CreateOIDCTrustRelationship(ctx).CreateOIDCTrustRelationshipRequest(createOIDCTrustRelationshipRequest).Execute()





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
	createOIDCTrustRelationshipRequest := openapiclient.createOIDCTrustRelationship_request{CreateOIDCTrustRelationshipRequestOneOf: openapiclient.NewCreateOIDCTrustRelationshipRequestOneOf(int32(789), int32(321), "org", []string{"Audiences_example"}, []openapiclient.ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner{*openapiclient.NewListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner("sub", openapiclient.listOIDCTrustRelationships_200_response_allOf_results_inner_allOf_requiredClaims_inner_value{Bool: new(bool)})})} // CreateOIDCTrustRelationshipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.CreateOIDCTrustRelationship(context.Background()).CreateOIDCTrustRelationshipRequest(createOIDCTrustRelationshipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.CreateOIDCTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOIDCTrustRelationship`: CreateOIDCTrustRelationship201Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.CreateOIDCTrustRelationship`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOIDCTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOIDCTrustRelationshipRequest** | [**CreateOIDCTrustRelationshipRequest**](CreateOIDCTrustRelationshipRequest.md) |  | 

### Return type

[**CreateOIDCTrustRelationship201Response**](CreateOIDCTrustRelationship201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOIDCProvider

> DeleteOIDCProvider(ctx, id).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Provider to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OIDCAPI.DeleteOIDCProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.DeleteOIDCProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Provider to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOIDCProviderRequest struct via the builder pattern


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


## DeleteOIDCTrustRelationship

> DeleteOIDCTrustRelationship(ctx, id).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Trust Relationship to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OIDCAPI.DeleteOIDCTrustRelationship(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.DeleteOIDCTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Trust Relationship to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOIDCTrustRelationshipRequest struct via the builder pattern


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


## ExchangeOIDCToken

> ExchangeOIDCToken200Response ExchangeOIDCToken(ctx).ExchangeOIDCTokenRequest(exchangeOIDCTokenRequest).Execute()





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
	exchangeOIDCTokenRequest := *openapiclient.NewExchangeOIDCTokenRequest(int32(123), "john.doe", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...") // ExchangeOIDCTokenRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.ExchangeOIDCToken(context.Background()).ExchangeOIDCTokenRequest(exchangeOIDCTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.ExchangeOIDCToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExchangeOIDCToken`: ExchangeOIDCToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.ExchangeOIDCToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExchangeOIDCTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exchangeOIDCTokenRequest** | [**ExchangeOIDCTokenRequest**](ExchangeOIDCTokenRequest.md) |  | 

### Return type

[**ExchangeOIDCToken200Response**](ExchangeOIDCToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCProvider

> CreateOIDCProvider201Response GetOIDCProvider(ctx, id).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.GetOIDCProvider(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.GetOIDCProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOIDCProvider`: CreateOIDCProvider201Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.GetOIDCProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateOIDCProvider201Response**](CreateOIDCProvider201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCProviderAvailableServiceAccounts

> GetOIDCProviderAvailableServiceAccounts200Response GetOIDCProviderAvailableServiceAccounts(ctx, id).TeamId(teamId).PageSize(pageSize).Prev(prev).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Provider
	teamId := int32(456) // int32 | The team ID that the service account must be a member of. Required if the user making the request only has team-level permission to manage trust relationships. If the user has organization-level permission then this parameter has no effect.  (optional)
	pageSize := int32(20) // int32 | Number of service accounts to return (optional) (default to 10)
	prev := int32(789) // int32 | ID of the last service account from the previous page for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.GetOIDCProviderAvailableServiceAccounts(context.Background(), id).TeamId(teamId).PageSize(pageSize).Prev(prev).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.GetOIDCProviderAvailableServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOIDCProviderAvailableServiceAccounts`: GetOIDCProviderAvailableServiceAccounts200Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.GetOIDCProviderAvailableServiceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCProviderAvailableServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamId** | **int32** | The team ID that the service account must be a member of. Required if the user making the request only has team-level permission to manage trust relationships. If the user has organization-level permission then this parameter has no effect.  | 
 **pageSize** | **int32** | Number of service accounts to return | [default to 10]
 **prev** | **int32** | ID of the last service account from the previous page for pagination | [default to 0]

### Return type

[**GetOIDCProviderAvailableServiceAccounts200Response**](GetOIDCProviderAvailableServiceAccounts200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOIDCTrustRelationship

> ListOIDCTrustRelationships200ResponseAllOfResultsInner GetOIDCTrustRelationship(ctx, id).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Trust Relationship

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.GetOIDCTrustRelationship(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.GetOIDCTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOIDCTrustRelationship`: ListOIDCTrustRelationships200ResponseAllOfResultsInner
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.GetOIDCTrustRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Trust Relationship | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOIDCTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOIDCTrustRelationships200ResponseAllOfResultsInner**](ListOIDCTrustRelationships200ResponseAllOfResultsInner.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOIDCProviders

> ListOIDCProviders200Response ListOIDCProviders(ctx).PageSize(pageSize).Prev(prev).FilterScope(filterScope).FilterScopeId(filterScopeId).Execute()





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
	pageSize := int32(20) // int32 | Number of providers to return (optional) (default to 10)
	prev := int32(123) // int32 | ID of the last provider from the previous page for pagination (optional) (default to 0)
	filterScope := "team" // string | Filter providers by scope (org or team). When filtering by team, the `scopeId` parameter must also be provided. The org scope will only return organization-scoped providers. The team scope will return both org-scoped providers as well as team-scoped providers where the scopeId matches the team ID.  (optional)
	filterScopeId := int32(789) // int32 | Filter providers by scope ID. Required when the `scope` parameter is `team`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.ListOIDCProviders(context.Background()).PageSize(pageSize).Prev(prev).FilterScope(filterScope).FilterScopeId(filterScopeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.ListOIDCProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOIDCProviders`: ListOIDCProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.ListOIDCProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOIDCProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of providers to return | [default to 10]
 **prev** | **int32** | ID of the last provider from the previous page for pagination | [default to 0]
 **filterScope** | **string** | Filter providers by scope (org or team). When filtering by team, the &#x60;scopeId&#x60; parameter must also be provided. The org scope will only return organization-scoped providers. The team scope will return both org-scoped providers as well as team-scoped providers where the scopeId matches the team ID.  | 
 **filterScopeId** | **int32** | Filter providers by scope ID. Required when the &#x60;scope&#x60; parameter is &#x60;team&#x60;. | 

### Return type

[**ListOIDCProviders200Response**](ListOIDCProviders200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOIDCTrustRelationships

> ListOIDCTrustRelationships200Response ListOIDCTrustRelationships(ctx).UserId(userId).ProviderId(providerId).PageSize(pageSize).Prev(prev).Execute()





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
	userId := int32(789) // int32 | Filter by user ID (optional)
	providerId := int32(321) // int32 | Filter by provider ID (optional)
	pageSize := int32(20) // int32 | Number of trust relationships to return (optional) (default to 10)
	prev := int32(456) // int32 | ID of the last trust relationship from the previous page for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.ListOIDCTrustRelationships(context.Background()).UserId(userId).ProviderId(providerId).PageSize(pageSize).Prev(prev).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.ListOIDCTrustRelationships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOIDCTrustRelationships`: ListOIDCTrustRelationships200Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.ListOIDCTrustRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOIDCTrustRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **int32** | Filter by user ID | 
 **providerId** | **int32** | Filter by provider ID | 
 **pageSize** | **int32** | Number of trust relationships to return | [default to 10]
 **prev** | **int32** | ID of the last trust relationship from the previous page for pagination | [default to 0]

### Return type

[**ListOIDCTrustRelationships200Response**](ListOIDCTrustRelationships200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOIDCTrustRelationship

> CreateOIDCTrustRelationship201Response UpdateOIDCTrustRelationship(ctx, id).UpdateOIDCTrustRelationshipRequest(updateOIDCTrustRelationshipRequest).Execute()





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
	id := int32(123) // int32 | The ID of the OIDC Trust Relationship to update
	updateOIDCTrustRelationshipRequest := *openapiclient.NewUpdateOIDCTrustRelationshipRequest() // UpdateOIDCTrustRelationshipRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OIDCAPI.UpdateOIDCTrustRelationship(context.Background(), id).UpdateOIDCTrustRelationshipRequest(updateOIDCTrustRelationshipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OIDCAPI.UpdateOIDCTrustRelationship``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOIDCTrustRelationship`: CreateOIDCTrustRelationship201Response
	fmt.Fprintf(os.Stdout, "Response from `OIDCAPI.UpdateOIDCTrustRelationship`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the OIDC Trust Relationship to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOIDCTrustRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOIDCTrustRelationshipRequest** | [**UpdateOIDCTrustRelationshipRequest**](UpdateOIDCTrustRelationshipRequest.md) |  | 

### Return type

[**CreateOIDCTrustRelationship201Response**](CreateOIDCTrustRelationship201Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

