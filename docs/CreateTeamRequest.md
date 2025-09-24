# CreateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the team | 
**DefaultRoleId** | **int32** | Default role ID for team members | 
**AutoAddUsers** | Pointer to **bool** | Whether to automatically add users to this team | [optional] [default to false]
**UniqueIdentifier** | Pointer to **string** | Unique identifier for the team | [optional] 
**TeamGroupIds** | Pointer to **[]int32** | IDs of team groups to add this team to (requires team groups feature) | [optional] 

## Methods

### NewCreateTeamRequest

`func NewCreateTeamRequest(name string, defaultRoleId int32, ) *CreateTeamRequest`

NewCreateTeamRequest instantiates a new CreateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamRequestWithDefaults

`func NewCreateTeamRequestWithDefaults() *CreateTeamRequest`

NewCreateTeamRequestWithDefaults instantiates a new CreateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTeamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultRoleId

`func (o *CreateTeamRequest) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *CreateTeamRequest) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *CreateTeamRequest) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.


### GetAutoAddUsers

`func (o *CreateTeamRequest) GetAutoAddUsers() bool`

GetAutoAddUsers returns the AutoAddUsers field if non-nil, zero value otherwise.

### GetAutoAddUsersOk

`func (o *CreateTeamRequest) GetAutoAddUsersOk() (*bool, bool)`

GetAutoAddUsersOk returns a tuple with the AutoAddUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddUsers

`func (o *CreateTeamRequest) SetAutoAddUsers(v bool)`

SetAutoAddUsers sets AutoAddUsers field to given value.

### HasAutoAddUsers

`func (o *CreateTeamRequest) HasAutoAddUsers() bool`

HasAutoAddUsers returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *CreateTeamRequest) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *CreateTeamRequest) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *CreateTeamRequest) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *CreateTeamRequest) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### GetTeamGroupIds

`func (o *CreateTeamRequest) GetTeamGroupIds() []int32`

GetTeamGroupIds returns the TeamGroupIds field if non-nil, zero value otherwise.

### GetTeamGroupIdsOk

`func (o *CreateTeamRequest) GetTeamGroupIdsOk() (*[]int32, bool)`

GetTeamGroupIdsOk returns a tuple with the TeamGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamGroupIds

`func (o *CreateTeamRequest) SetTeamGroupIds(v []int32)`

SetTeamGroupIds sets TeamGroupIds field to given value.

### HasTeamGroupIds

`func (o *CreateTeamRequest) HasTeamGroupIds() bool`

HasTeamGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


