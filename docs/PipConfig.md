# PipConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GetOrganizationPipSettings200ResponseRepositoriesInner**](GetOrganizationPipSettings200ResponseRepositoriesInner.md) | List of configured Pip Repositories | [optional] 

## Methods

### NewPipConfig

`func NewPipConfig() *PipConfig`

NewPipConfig instantiates a new PipConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipConfigWithDefaults

`func NewPipConfigWithDefaults() *PipConfig`

NewPipConfigWithDefaults instantiates a new PipConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *PipConfig) GetRepositories() []GetOrganizationPipSettings200ResponseRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PipConfig) GetRepositoriesOk() (*[]GetOrganizationPipSettings200ResponseRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PipConfig) SetRepositories(v []GetOrganizationPipSettings200ResponseRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PipConfig) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


