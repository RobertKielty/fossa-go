# CreateReleaseGroupRequestRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Projects** | Pointer to [**[]CreateReleaseGroupRequestReleaseProjectsInner**](CreateReleaseGroupRequestReleaseProjectsInner.md) |  | [optional] 

## Methods

### NewCreateReleaseGroupRequestRelease

`func NewCreateReleaseGroupRequestRelease() *CreateReleaseGroupRequestRelease`

NewCreateReleaseGroupRequestRelease instantiates a new CreateReleaseGroupRequestRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseGroupRequestReleaseWithDefaults

`func NewCreateReleaseGroupRequestReleaseWithDefaults() *CreateReleaseGroupRequestRelease`

NewCreateReleaseGroupRequestReleaseWithDefaults instantiates a new CreateReleaseGroupRequestRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateReleaseGroupRequestRelease) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateReleaseGroupRequestRelease) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateReleaseGroupRequestRelease) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateReleaseGroupRequestRelease) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProjects

`func (o *CreateReleaseGroupRequestRelease) GetProjects() []CreateReleaseGroupRequestReleaseProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateReleaseGroupRequestRelease) GetProjectsOk() (*[]CreateReleaseGroupRequestReleaseProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateReleaseGroupRequestRelease) SetProjects(v []CreateReleaseGroupRequestReleaseProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CreateReleaseGroupRequestRelease) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


