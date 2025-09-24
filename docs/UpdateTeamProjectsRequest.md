# UpdateTeamProjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on team projects | 
**Projects** | Pointer to [**UpdateTeamProjectsRequestProjects**](UpdateTeamProjectsRequestProjects.md) |  | [optional] 
**Filters** | Pointer to [**UpdateTeamProjectsRequestFilters**](UpdateTeamProjectsRequestFilters.md) |  | [optional] 

## Methods

### NewUpdateTeamProjectsRequest

`func NewUpdateTeamProjectsRequest(action string, ) *UpdateTeamProjectsRequest`

NewUpdateTeamProjectsRequest instantiates a new UpdateTeamProjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamProjectsRequestWithDefaults

`func NewUpdateTeamProjectsRequestWithDefaults() *UpdateTeamProjectsRequest`

NewUpdateTeamProjectsRequestWithDefaults instantiates a new UpdateTeamProjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateTeamProjectsRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateTeamProjectsRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateTeamProjectsRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetProjects

`func (o *UpdateTeamProjectsRequest) GetProjects() UpdateTeamProjectsRequestProjects`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UpdateTeamProjectsRequest) GetProjectsOk() (*UpdateTeamProjectsRequestProjects, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UpdateTeamProjectsRequest) SetProjects(v UpdateTeamProjectsRequestProjects)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UpdateTeamProjectsRequest) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetFilters

`func (o *UpdateTeamProjectsRequest) GetFilters() UpdateTeamProjectsRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateTeamProjectsRequest) GetFiltersOk() (*UpdateTeamProjectsRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateTeamProjectsRequest) SetFilters(v UpdateTeamProjectsRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateTeamProjectsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


