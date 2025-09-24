# UpdateTeamUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on team users | 
**Users** | Pointer to [**[]UpdateTeamUsersRequestUsersInner**](UpdateTeamUsersRequestUsersInner.md) | List of users to add, remove, or replace | [optional] 

## Methods

### NewUpdateTeamUsersRequest

`func NewUpdateTeamUsersRequest(action string, ) *UpdateTeamUsersRequest`

NewUpdateTeamUsersRequest instantiates a new UpdateTeamUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamUsersRequestWithDefaults

`func NewUpdateTeamUsersRequestWithDefaults() *UpdateTeamUsersRequest`

NewUpdateTeamUsersRequestWithDefaults instantiates a new UpdateTeamUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateTeamUsersRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateTeamUsersRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateTeamUsersRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetUsers

`func (o *UpdateTeamUsersRequest) GetUsers() []UpdateTeamUsersRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateTeamUsersRequest) GetUsersOk() (*[]UpdateTeamUsersRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateTeamUsersRequest) SetUsers(v []UpdateTeamUsersRequestUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateTeamUsersRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


