# BuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archives** | Pointer to [**BuildRequestArchives**](BuildRequestArchives.md) |  | [optional] 
**SelectedTeams** | Pointer to [**[]BuildRequestSelectedTeamsInner**](BuildRequestSelectedTeamsInner.md) |  | [optional] 
**ForceRebuild** | Pointer to **bool** | Force a rebuild no matter what | [optional] 

## Methods

### NewBuildRequest

`func NewBuildRequest() *BuildRequest`

NewBuildRequest instantiates a new BuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRequestWithDefaults

`func NewBuildRequestWithDefaults() *BuildRequest`

NewBuildRequestWithDefaults instantiates a new BuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchives

`func (o *BuildRequest) GetArchives() BuildRequestArchives`

GetArchives returns the Archives field if non-nil, zero value otherwise.

### GetArchivesOk

`func (o *BuildRequest) GetArchivesOk() (*BuildRequestArchives, bool)`

GetArchivesOk returns a tuple with the Archives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchives

`func (o *BuildRequest) SetArchives(v BuildRequestArchives)`

SetArchives sets Archives field to given value.

### HasArchives

`func (o *BuildRequest) HasArchives() bool`

HasArchives returns a boolean if a field has been set.

### GetSelectedTeams

`func (o *BuildRequest) GetSelectedTeams() []BuildRequestSelectedTeamsInner`

GetSelectedTeams returns the SelectedTeams field if non-nil, zero value otherwise.

### GetSelectedTeamsOk

`func (o *BuildRequest) GetSelectedTeamsOk() (*[]BuildRequestSelectedTeamsInner, bool)`

GetSelectedTeamsOk returns a tuple with the SelectedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedTeams

`func (o *BuildRequest) SetSelectedTeams(v []BuildRequestSelectedTeamsInner)`

SetSelectedTeams sets SelectedTeams field to given value.

### HasSelectedTeams

`func (o *BuildRequest) HasSelectedTeams() bool`

HasSelectedTeams returns a boolean if a field has been set.

### GetForceRebuild

`func (o *BuildRequest) GetForceRebuild() bool`

GetForceRebuild returns the ForceRebuild field if non-nil, zero value otherwise.

### GetForceRebuildOk

`func (o *BuildRequest) GetForceRebuildOk() (*bool, bool)`

GetForceRebuildOk returns a tuple with the ForceRebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRebuild

`func (o *BuildRequest) SetForceRebuild(v bool)`

SetForceRebuild sets ForceRebuild field to given value.

### HasForceRebuild

`func (o *BuildRequest) HasForceRebuild() bool`

HasForceRebuild returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


