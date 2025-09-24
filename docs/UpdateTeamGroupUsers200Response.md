# UpdateTeamGroupUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team group ID | [optional] 
**Users** | Pointer to [**[]UpdateTeamGroupUsers200ResponseUsersInner**](UpdateTeamGroupUsers200ResponseUsersInner.md) | Updated list of team group users | [optional] 

## Methods

### NewUpdateTeamGroupUsers200Response

`func NewUpdateTeamGroupUsers200Response() *UpdateTeamGroupUsers200Response`

NewUpdateTeamGroupUsers200Response instantiates a new UpdateTeamGroupUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamGroupUsers200ResponseWithDefaults

`func NewUpdateTeamGroupUsers200ResponseWithDefaults() *UpdateTeamGroupUsers200Response`

NewUpdateTeamGroupUsers200ResponseWithDefaults instantiates a new UpdateTeamGroupUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTeamGroupUsers200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTeamGroupUsers200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTeamGroupUsers200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTeamGroupUsers200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateTeamGroupUsers200Response) GetUsers() []UpdateTeamGroupUsers200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateTeamGroupUsers200Response) GetUsersOk() (*[]UpdateTeamGroupUsers200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateTeamGroupUsers200Response) SetUsers(v []UpdateTeamGroupUsers200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateTeamGroupUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


