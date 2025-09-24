# UpdateTeamGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team group ID | [optional] 
**Name** | Pointer to **string** | Team group name | [optional] 
**DefaultRoleId** | Pointer to **int32** | Default role ID for members | [optional] 
**AutoAddUsers** | Pointer to **bool** | Whether to automatically add users | [optional] 
**UniqueIdentifier** | Pointer to **string** | Unique identifier | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update timestamp | [optional] 

## Methods

### NewUpdateTeamGroup200Response

`func NewUpdateTeamGroup200Response() *UpdateTeamGroup200Response`

NewUpdateTeamGroup200Response instantiates a new UpdateTeamGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamGroup200ResponseWithDefaults

`func NewUpdateTeamGroup200ResponseWithDefaults() *UpdateTeamGroup200Response`

NewUpdateTeamGroup200ResponseWithDefaults instantiates a new UpdateTeamGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTeamGroup200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTeamGroup200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTeamGroup200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTeamGroup200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateTeamGroup200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTeamGroup200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTeamGroup200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTeamGroup200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *UpdateTeamGroup200Response) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *UpdateTeamGroup200Response) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *UpdateTeamGroup200Response) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *UpdateTeamGroup200Response) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetAutoAddUsers

`func (o *UpdateTeamGroup200Response) GetAutoAddUsers() bool`

GetAutoAddUsers returns the AutoAddUsers field if non-nil, zero value otherwise.

### GetAutoAddUsersOk

`func (o *UpdateTeamGroup200Response) GetAutoAddUsersOk() (*bool, bool)`

GetAutoAddUsersOk returns a tuple with the AutoAddUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddUsers

`func (o *UpdateTeamGroup200Response) SetAutoAddUsers(v bool)`

SetAutoAddUsers sets AutoAddUsers field to given value.

### HasAutoAddUsers

`func (o *UpdateTeamGroup200Response) HasAutoAddUsers() bool`

HasAutoAddUsers returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *UpdateTeamGroup200Response) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *UpdateTeamGroup200Response) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *UpdateTeamGroup200Response) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *UpdateTeamGroup200Response) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateTeamGroup200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateTeamGroup200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateTeamGroup200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateTeamGroup200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UpdateTeamGroup200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateTeamGroup200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateTeamGroup200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateTeamGroup200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


