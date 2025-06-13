# GetTeamMembers200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | ID of the user | 
**RoleId** | **int32** | ID of the role assigned to the user in the team | 
**Username** | **string** | Username of the user | 
**Email** | **string** | Email of the user | 

## Methods

### NewGetTeamMembers200ResponseResultsInner

`func NewGetTeamMembers200ResponseResultsInner(userId int32, roleId int32, username string, email string, ) *GetTeamMembers200ResponseResultsInner`

NewGetTeamMembers200ResponseResultsInner instantiates a new GetTeamMembers200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamMembers200ResponseResultsInnerWithDefaults

`func NewGetTeamMembers200ResponseResultsInnerWithDefaults() *GetTeamMembers200ResponseResultsInner`

NewGetTeamMembers200ResponseResultsInnerWithDefaults instantiates a new GetTeamMembers200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetTeamMembers200ResponseResultsInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetTeamMembers200ResponseResultsInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetTeamMembers200ResponseResultsInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetRoleId

`func (o *GetTeamMembers200ResponseResultsInner) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *GetTeamMembers200ResponseResultsInner) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *GetTeamMembers200ResponseResultsInner) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.


### GetUsername

`func (o *GetTeamMembers200ResponseResultsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetTeamMembers200ResponseResultsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetTeamMembers200ResponseResultsInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *GetTeamMembers200ResponseResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetTeamMembers200ResponseResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetTeamMembers200ResponseResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


