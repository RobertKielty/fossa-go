# UpdateTeamUsersRequestUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the user | 
**RoleId** | Pointer to **int32** | Role ID for the user (required for add and replace actions) | [optional] 

## Methods

### NewUpdateTeamUsersRequestUsersInner

`func NewUpdateTeamUsersRequestUsersInner(id int32, ) *UpdateTeamUsersRequestUsersInner`

NewUpdateTeamUsersRequestUsersInner instantiates a new UpdateTeamUsersRequestUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamUsersRequestUsersInnerWithDefaults

`func NewUpdateTeamUsersRequestUsersInnerWithDefaults() *UpdateTeamUsersRequestUsersInner`

NewUpdateTeamUsersRequestUsersInnerWithDefaults instantiates a new UpdateTeamUsersRequestUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTeamUsersRequestUsersInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTeamUsersRequestUsersInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTeamUsersRequestUsersInner) SetId(v int32)`

SetId sets Id field to given value.


### GetRoleId

`func (o *UpdateTeamUsersRequestUsersInner) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *UpdateTeamUsersRequestUsersInner) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *UpdateTeamUsersRequestUsersInner) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *UpdateTeamUsersRequestUsersInner) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


