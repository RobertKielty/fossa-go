# UpdateTeamProjectsRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Filter by project title | [optional] 
**Labels** | Pointer to **[]int32** | Filter by label IDs | [optional] 
**Type** | Pointer to **string** | Filter by project type | [optional] 
**LastRevisionWithin** | Pointer to **string** | Filter by last revision date | [optional] 
**IsPublic** | Pointer to **bool** | Filter by public/private status | [optional] 

## Methods

### NewUpdateTeamProjectsRequestFilters

`func NewUpdateTeamProjectsRequestFilters() *UpdateTeamProjectsRequestFilters`

NewUpdateTeamProjectsRequestFilters instantiates a new UpdateTeamProjectsRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamProjectsRequestFiltersWithDefaults

`func NewUpdateTeamProjectsRequestFiltersWithDefaults() *UpdateTeamProjectsRequestFilters`

NewUpdateTeamProjectsRequestFiltersWithDefaults instantiates a new UpdateTeamProjectsRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateTeamProjectsRequestFilters) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateTeamProjectsRequestFilters) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateTeamProjectsRequestFilters) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateTeamProjectsRequestFilters) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateTeamProjectsRequestFilters) GetLabels() []int32`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateTeamProjectsRequestFilters) GetLabelsOk() (*[]int32, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateTeamProjectsRequestFilters) SetLabels(v []int32)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateTeamProjectsRequestFilters) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *UpdateTeamProjectsRequestFilters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateTeamProjectsRequestFilters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateTeamProjectsRequestFilters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateTeamProjectsRequestFilters) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLastRevisionWithin

`func (o *UpdateTeamProjectsRequestFilters) GetLastRevisionWithin() string`

GetLastRevisionWithin returns the LastRevisionWithin field if non-nil, zero value otherwise.

### GetLastRevisionWithinOk

`func (o *UpdateTeamProjectsRequestFilters) GetLastRevisionWithinOk() (*string, bool)`

GetLastRevisionWithinOk returns a tuple with the LastRevisionWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionWithin

`func (o *UpdateTeamProjectsRequestFilters) SetLastRevisionWithin(v string)`

SetLastRevisionWithin sets LastRevisionWithin field to given value.

### HasLastRevisionWithin

`func (o *UpdateTeamProjectsRequestFilters) HasLastRevisionWithin() bool`

HasLastRevisionWithin returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateTeamProjectsRequestFilters) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateTeamProjectsRequestFilters) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateTeamProjectsRequestFilters) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateTeamProjectsRequestFilters) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


