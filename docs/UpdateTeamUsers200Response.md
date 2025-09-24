# UpdateTeamUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team ID | [optional] 
**Users** | Pointer to [**[]UpdateTeamUsers200ResponseUsersInner**](UpdateTeamUsers200ResponseUsersInner.md) | Updated list of team users | [optional] 

## Methods

### NewUpdateTeamUsers200Response

`func NewUpdateTeamUsers200Response() *UpdateTeamUsers200Response`

NewUpdateTeamUsers200Response instantiates a new UpdateTeamUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTeamUsers200ResponseWithDefaults

`func NewUpdateTeamUsers200ResponseWithDefaults() *UpdateTeamUsers200Response`

NewUpdateTeamUsers200ResponseWithDefaults instantiates a new UpdateTeamUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateTeamUsers200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateTeamUsers200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateTeamUsers200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateTeamUsers200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsers

`func (o *UpdateTeamUsers200Response) GetUsers() []UpdateTeamUsers200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UpdateTeamUsers200Response) GetUsersOk() (*[]UpdateTeamUsers200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UpdateTeamUsers200Response) SetUsers(v []UpdateTeamUsers200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UpdateTeamUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


