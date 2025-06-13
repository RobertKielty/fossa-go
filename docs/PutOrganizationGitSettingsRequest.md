# PutOrganizationGitSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthOptions** | Pointer to [**[]GetOrganizationGitSettings200ResponseAuthOptionsInner**](GetOrganizationGitSettings200ResponseAuthOptionsInner.md) | List of Git configurations | [optional] 

## Methods

### NewPutOrganizationGitSettingsRequest

`func NewPutOrganizationGitSettingsRequest() *PutOrganizationGitSettingsRequest`

NewPutOrganizationGitSettingsRequest instantiates a new PutOrganizationGitSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrganizationGitSettingsRequestWithDefaults

`func NewPutOrganizationGitSettingsRequestWithDefaults() *PutOrganizationGitSettingsRequest`

NewPutOrganizationGitSettingsRequestWithDefaults instantiates a new PutOrganizationGitSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthOptions

`func (o *PutOrganizationGitSettingsRequest) GetAuthOptions() []GetOrganizationGitSettings200ResponseAuthOptionsInner`

GetAuthOptions returns the AuthOptions field if non-nil, zero value otherwise.

### GetAuthOptionsOk

`func (o *PutOrganizationGitSettingsRequest) GetAuthOptionsOk() (*[]GetOrganizationGitSettings200ResponseAuthOptionsInner, bool)`

GetAuthOptionsOk returns a tuple with the AuthOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthOptions

`func (o *PutOrganizationGitSettingsRequest) SetAuthOptions(v []GetOrganizationGitSettings200ResponseAuthOptionsInner)`

SetAuthOptions sets AuthOptions field to given value.

### HasAuthOptions

`func (o *PutOrganizationGitSettingsRequest) HasAuthOptions() bool`

HasAuthOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


