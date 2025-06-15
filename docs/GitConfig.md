# GitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthOptions** | Pointer to [**[]GetOrganizationGitSettings200ResponseAuthOptionsInner**](GetOrganizationGitSettings200ResponseAuthOptionsInner.md) | List of Git configurations | [optional] 

## Methods

### NewGitConfig

`func NewGitConfig() *GitConfig`

NewGitConfig instantiates a new GitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitConfigWithDefaults

`func NewGitConfigWithDefaults() *GitConfig`

NewGitConfigWithDefaults instantiates a new GitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthOptions

`func (o *GitConfig) GetAuthOptions() []GetOrganizationGitSettings200ResponseAuthOptionsInner`

GetAuthOptions returns the AuthOptions field if non-nil, zero value otherwise.

### GetAuthOptionsOk

`func (o *GitConfig) GetAuthOptionsOk() (*[]GetOrganizationGitSettings200ResponseAuthOptionsInner, bool)`

GetAuthOptionsOk returns a tuple with the AuthOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOptions

`func (o *GitConfig) SetAuthOptions(v []GetOrganizationGitSettings200ResponseAuthOptionsInner)`

SetAuthOptions sets AuthOptions field to given value.

### HasAuthOptions

`func (o *GitConfig) HasAuthOptions() bool`

HasAuthOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


