# UpdateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the team | [optional] 
**AutoAddUsers** | Pointer to **bool** | Whether to automatically add users to this team | [optional] 
**DefaultRoleId** | Pointer to **int32** | Default role ID for team members | [optional] 
**UniqueIdentifier** | Pointer to **string** | Unique identifier for the team | [optional] 

## Methods

### NewUpdateTeamRequest

`func NewUpdateTeamRequest() *UpdateTeamRequest`

NewUpdateTeamRequest instantiates a new UpdateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamRequestWithDefaults

`func NewUpdateTeamRequestWithDefaults() *UpdateTeamRequest`

NewUpdateTeamRequestWithDefaults instantiates a new UpdateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTeamRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTeamRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoAddUsers

`func (o *UpdateTeamRequest) GetAutoAddUsers() bool`

GetAutoAddUsers returns the AutoAddUsers field if non-nil, zero value otherwise.

### GetAutoAddUsersOk

`func (o *UpdateTeamRequest) GetAutoAddUsersOk() (*bool, bool)`

GetAutoAddUsersOk returns a tuple with the AutoAddUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddUsers

`func (o *UpdateTeamRequest) SetAutoAddUsers(v bool)`

SetAutoAddUsers sets AutoAddUsers field to given value.

### HasAutoAddUsers

`func (o *UpdateTeamRequest) HasAutoAddUsers() bool`

HasAutoAddUsers returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *UpdateTeamRequest) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *UpdateTeamRequest) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *UpdateTeamRequest) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *UpdateTeamRequest) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateTeamRequest) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateTeamRequest) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateTeamRequest) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *UpdateTeamRequest) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


