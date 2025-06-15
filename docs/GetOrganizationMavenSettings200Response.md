# GetOrganizationMavenSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GetOrganizationMavenSettings200ResponseRepositoriesInner**](GetOrganizationMavenSettings200ResponseRepositoriesInner.md) | List of configured Maven Repositories | [optional] 
**Servers** | Pointer to [**[]GetOrganizationMavenSettings200ResponseServersInner**](GetOrganizationMavenSettings200ResponseServersInner.md) | List of configured Credentials for Maven Repositories | [optional] 

## Methods

### NewGetOrganizationMavenSettings200Response

`func NewGetOrganizationMavenSettings200Response() *GetOrganizationMavenSettings200Response`

NewGetOrganizationMavenSettings200Response instantiates a new GetOrganizationMavenSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationMavenSettings200ResponseWithDefaults

`func NewGetOrganizationMavenSettings200ResponseWithDefaults() *GetOrganizationMavenSettings200Response`

NewGetOrganizationMavenSettings200ResponseWithDefaults instantiates a new GetOrganizationMavenSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *GetOrganizationMavenSettings200Response) GetRepositories() []GetOrganizationMavenSettings200ResponseRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *GetOrganizationMavenSettings200Response) GetRepositoriesOk() (*[]GetOrganizationMavenSettings200ResponseRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *GetOrganizationMavenSettings200Response) SetRepositories(v []GetOrganizationMavenSettings200ResponseRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *GetOrganizationMavenSettings200Response) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetServers

`func (o *GetOrganizationMavenSettings200Response) GetServers() []GetOrganizationMavenSettings200ResponseServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetOrganizationMavenSettings200Response) GetServersOk() (*[]GetOrganizationMavenSettings200ResponseServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetOrganizationMavenSettings200Response) SetServers(v []GetOrganizationMavenSettings200ResponseServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetOrganizationMavenSettings200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


