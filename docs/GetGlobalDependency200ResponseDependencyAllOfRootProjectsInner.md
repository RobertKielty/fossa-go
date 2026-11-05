# GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Conclusions** | Pointer to [**GetProjectDependencies200ResponseDependenciesInnerConclusions**](GetProjectDependencies200ResponseDependenciesInnerConclusions.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**[]GetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerAllOfTeamsInner**](GetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerAllOfTeamsInner.md) |  | [optional] 
**Paths** | Pointer to **[][]string** | The dependency paths from the root project to this dependency. | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInner

`func NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInner() *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner`

NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInner instantiates a new GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerWithDefaults

`func NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerWithDefaults() *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner`

NewGetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerWithDefaults instantiates a new GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRevision

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetBranch

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConclusions

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetConclusions() GetProjectDependencies200ResponseDependenciesInnerConclusions`

GetConclusions returns the Conclusions field if non-nil, zero value otherwise.

### GetConclusionsOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetConclusionsOk() (*GetProjectDependencies200ResponseDependenciesInnerConclusions, bool)`

GetConclusionsOk returns a tuple with the Conclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusions

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetConclusions(v GetProjectDependencies200ResponseDependenciesInnerConclusions)`

SetConclusions sets Conclusions field to given value.

### HasConclusions

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasConclusions() bool`

HasConclusions returns a boolean if a field has been set.

### GetType

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTeams

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetTeams() []GetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerAllOfTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetTeamsOk() (*[]GetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerAllOfTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetTeams(v []GetGlobalDependency200ResponseDependencyAllOfRootProjectsInnerAllOfTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetPaths

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetPaths() [][]string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetPathsOk() (*[][]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetPaths(v [][]string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetDepth

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


