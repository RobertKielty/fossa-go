# UpdateTeamGroupUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Action to perform on team group users | 
**Users** | [**[]UpdateTeamGroupUsersRequestUsersInner**](UpdateTeamGroupUsersRequestUsersInner.md) | List of users to add, remove, or replace | 

## Methods

### NewUpdateTeamGroupUsersRequest

`func NewUpdateTeamGroupUsersRequest(action string, users []UpdateTeamGroupUsersRequestUsersInner, ) *UpdateTeamGroupUsersRequest`

NewUpdateTeamGroupUsersRequest instantiates a new UpdateTeamGroupUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamGroupUsersRequestWithDefaults

`func NewUpdateTeamGroupUsersRequestWithDefaults() *UpdateTeamGroupUsersRequest`

NewUpdateTeamGroupUsersRequestWithDefaults instantiates a new UpdateTeamGroupUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateTeamGroupUsersRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateTeamGroupUsersRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateTeamGroupUsersRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetUsers

`func (o *UpdateTeamGroupUsersRequest) GetUsers() []UpdateTeamGroupUsersRequestUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateTeamGroupUsersRequest) GetUsersOk() (*[]UpdateTeamGroupUsersRequestUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateTeamGroupUsersRequest) SetUsers(v []UpdateTeamGroupUsersRequestUsersInner)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


