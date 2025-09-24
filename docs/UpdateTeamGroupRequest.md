# UpdateTeamGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | New name for the team group | 
**DefaultRoleId** | **int32** | Default role ID for team group members | 

## Methods

### NewUpdateTeamGroupRequest

`func NewUpdateTeamGroupRequest(name string, defaultRoleId int32, ) *UpdateTeamGroupRequest`

NewUpdateTeamGroupRequest instantiates a new UpdateTeamGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamGroupRequestWithDefaults

`func NewUpdateTeamGroupRequestWithDefaults() *UpdateTeamGroupRequest`

NewUpdateTeamGroupRequestWithDefaults instantiates a new UpdateTeamGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTeamGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTeamGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTeamGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultRoleId

`func (o *UpdateTeamGroupRequest) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *UpdateTeamGroupRequest) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *UpdateTeamGroupRequest) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


