# PutOrganizationPipSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GetOrganizationRubyGemsSettings200ResponseSourcesInner**](GetOrganizationRubyGemsSettings200ResponseSourcesInner.md) | List of configured Pip Repositories | [optional] 

## Methods

### NewPutOrganizationPipSettingsRequest

`func NewPutOrganizationPipSettingsRequest() *PutOrganizationPipSettingsRequest`

NewPutOrganizationPipSettingsRequest instantiates a new PutOrganizationPipSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrganizationPipSettingsRequestWithDefaults

`func NewPutOrganizationPipSettingsRequestWithDefaults() *PutOrganizationPipSettingsRequest`

NewPutOrganizationPipSettingsRequestWithDefaults instantiates a new PutOrganizationPipSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *PutOrganizationPipSettingsRequest) GetRepositories() []GetOrganizationRubyGemsSettings200ResponseSourcesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *PutOrganizationPipSettingsRequest) GetRepositoriesOk() (*[]GetOrganizationRubyGemsSettings200ResponseSourcesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *PutOrganizationPipSettingsRequest) SetRepositories(v []GetOrganizationRubyGemsSettings200ResponseSourcesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *PutOrganizationPipSettingsRequest) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


