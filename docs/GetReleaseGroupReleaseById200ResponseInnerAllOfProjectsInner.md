# GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectGroupReleaseId** | Pointer to **int32** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Project** | Pointer to [**[]GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner**](GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner

`func NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner() *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner`

NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner instantiates a new GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerWithDefaults

`func NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerWithDefaults() *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner`

NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerWithDefaults instantiates a new GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProjectGroupReleaseId() int32`

GetProjectGroupReleaseId returns the ProjectGroupReleaseId field if non-nil, zero value otherwise.

### GetProjectGroupReleaseIdOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProjectGroupReleaseIdOk() (*int32, bool)`

GetProjectGroupReleaseIdOk returns a tuple with the ProjectGroupReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetProjectGroupReleaseId(v int32)`

SetProjectGroupReleaseId sets ProjectGroupReleaseId field to given value.

### HasProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasProjectGroupReleaseId() bool`

HasProjectGroupReleaseId returns a boolean if a field has been set.

### GetBranch

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRevisionId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProject

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProject() []GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) GetProjectOk() (*[]GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) SetProject(v []GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInner) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


