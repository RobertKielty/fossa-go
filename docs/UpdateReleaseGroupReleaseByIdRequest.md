# UpdateReleaseGroupReleaseByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Projects** | Pointer to [**[]UpdateReleaseGroupReleaseByIdRequestProjectsInner**](UpdateReleaseGroupReleaseByIdRequestProjectsInner.md) |  | [optional] 
**ProjectsToDelete** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateReleaseGroupReleaseByIdRequest

`func NewUpdateReleaseGroupReleaseByIdRequest(title string, ) *UpdateReleaseGroupReleaseByIdRequest`

NewUpdateReleaseGroupReleaseByIdRequest instantiates a new UpdateReleaseGroupReleaseByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReleaseGroupReleaseByIdRequestWithDefaults

`func NewUpdateReleaseGroupReleaseByIdRequestWithDefaults() *UpdateReleaseGroupReleaseByIdRequest`

NewUpdateReleaseGroupReleaseByIdRequestWithDefaults instantiates a new UpdateReleaseGroupReleaseByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateReleaseGroupReleaseByIdRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetProjects

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetProjects() []UpdateReleaseGroupReleaseByIdRequestProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetProjectsOk() (*[]UpdateReleaseGroupReleaseByIdRequestProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *UpdateReleaseGroupReleaseByIdRequest) SetProjects(v []UpdateReleaseGroupReleaseByIdRequestProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *UpdateReleaseGroupReleaseByIdRequest) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetProjectsToDelete

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetProjectsToDelete() []string`

GetProjectsToDelete returns the ProjectsToDelete field if non-nil, zero value otherwise.

### GetProjectsToDeleteOk

`func (o *UpdateReleaseGroupReleaseByIdRequest) GetProjectsToDeleteOk() (*[]string, bool)`

GetProjectsToDeleteOk returns a tuple with the ProjectsToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsToDelete

`func (o *UpdateReleaseGroupReleaseByIdRequest) SetProjectsToDelete(v []string)`

SetProjectsToDelete sets ProjectsToDelete field to given value.

### HasProjectsToDelete

`func (o *UpdateReleaseGroupReleaseByIdRequest) HasProjectsToDelete() bool`

HasProjectsToDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


