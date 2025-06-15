# PutOrganizationMavenSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GetOrganizationMavenSettings200ResponseRepositoriesInner**](GetOrganizationMavenSettings200ResponseRepositoriesInner.md) | List of configured Maven Repositories | [optional] 
**Servers** | Pointer to [**[]GetOrganizationMavenSettings200ResponseServersInner**](GetOrganizationMavenSettings200ResponseServersInner.md) | List of configured Credentials for Maven Repositories | [optional] 

## Methods

### NewPutOrganizationMavenSettingsRequest

`func NewPutOrganizationMavenSettingsRequest() *PutOrganizationMavenSettingsRequest`

NewPutOrganizationMavenSettingsRequest instantiates a new PutOrganizationMavenSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrganizationMavenSettingsRequestWithDefaults

`func NewPutOrganizationMavenSettingsRequestWithDefaults() *PutOrganizationMavenSettingsRequest`

NewPutOrganizationMavenSettingsRequestWithDefaults instantiates a new PutOrganizationMavenSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *PutOrganizationMavenSettingsRequest) GetRepositories() []GetOrganizationMavenSettings200ResponseRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PutOrganizationMavenSettingsRequest) GetRepositoriesOk() (*[]GetOrganizationMavenSettings200ResponseRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PutOrganizationMavenSettingsRequest) SetRepositories(v []GetOrganizationMavenSettings200ResponseRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PutOrganizationMavenSettingsRequest) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.

### GetServers

`func (o *PutOrganizationMavenSettingsRequest) GetServers() []GetOrganizationMavenSettings200ResponseServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *PutOrganizationMavenSettingsRequest) GetServersOk() (*[]GetOrganizationMavenSettings200ResponseServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *PutOrganizationMavenSettingsRequest) SetServers(v []GetOrganizationMavenSettings200ResponseServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *PutOrganizationMavenSettingsRequest) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


