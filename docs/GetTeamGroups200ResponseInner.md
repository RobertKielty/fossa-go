# GetTeamGroups200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team group ID | [optional] 
**Name** | Pointer to **string** | Team group name | [optional] 
**DefaultRoleId** | Pointer to **int32** | Default role ID for members | [optional] 
**Teams** | Pointer to [**[]GetTeamGroups200ResponseInnerTeamsInner**](GetTeamGroups200ResponseInnerTeamsInner.md) | Teams within this team group | [optional] 
**Members** | Pointer to [**[]GetAllTeams200ResponseInnerTeamUsersInner**](GetAllTeams200ResponseInnerTeamUsersInner.md) | Members of this team group | [optional] 

## Methods

### NewGetTeamGroups200ResponseInner

`func NewGetTeamGroups200ResponseInner() *GetTeamGroups200ResponseInner`

NewGetTeamGroups200ResponseInner instantiates a new GetTeamGroups200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamGroups200ResponseInnerWithDefaults

`func NewGetTeamGroups200ResponseInnerWithDefaults() *GetTeamGroups200ResponseInner`

NewGetTeamGroups200ResponseInnerWithDefaults instantiates a new GetTeamGroups200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTeamGroups200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamGroups200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamGroups200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTeamGroups200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetTeamGroups200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTeamGroups200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTeamGroups200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTeamGroups200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *GetTeamGroups200ResponseInner) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *GetTeamGroups200ResponseInner) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *GetTeamGroups200ResponseInner) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *GetTeamGroups200ResponseInner) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetTeams

`func (o *GetTeamGroups200ResponseInner) GetTeams() []GetTeamGroups200ResponseInnerTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetTeamGroups200ResponseInner) GetTeamsOk() (*[]GetTeamGroups200ResponseInnerTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetTeamGroups200ResponseInner) SetTeams(v []GetTeamGroups200ResponseInnerTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GetTeamGroups200ResponseInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetMembers

`func (o *GetTeamGroups200ResponseInner) GetMembers() []GetAllTeams200ResponseInnerTeamUsersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetTeamGroups200ResponseInner) GetMembersOk() (*[]GetAllTeams200ResponseInnerTeamUsersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetTeamGroups200ResponseInner) SetMembers(v []GetAllTeams200ResponseInnerTeamUsersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetTeamGroups200ResponseInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


