# CreateServiceAccountRequestTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the team to assign the service account to | 
**RoleId** | **int32** | ID of the team role to assign | 

## Methods

### NewCreateServiceAccountRequestTeam

`func NewCreateServiceAccountRequestTeam(id int32, roleId int32, ) *CreateServiceAccountRequestTeam`

NewCreateServiceAccountRequestTeam instantiates a new CreateServiceAccountRequestTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountRequestTeamWithDefaults

`func NewCreateServiceAccountRequestTeamWithDefaults() *CreateServiceAccountRequestTeam`

NewCreateServiceAccountRequestTeamWithDefaults instantiates a new CreateServiceAccountRequestTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateServiceAccountRequestTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateServiceAccountRequestTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateServiceAccountRequestTeam) SetId(v int32)`

SetId sets Id field to given value.


### GetRoleId

`func (o *CreateServiceAccountRequestTeam) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *CreateServiceAccountRequestTeam) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *CreateServiceAccountRequestTeam) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


