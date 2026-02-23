# UpdateProjectRequestNotificationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Notification ID (for updates to existing notifications) | [optional] 
**Channel** | **string** | The event channel that triggers this notification | 
**Service** | **string** | The service to use for sending notifications | 
**SubscribedUsers** | Pointer to **[]int32** | Array of user IDs to receive EMAIL notifications | [optional] 

## Methods

### NewUpdateProjectRequestNotificationsInner

`func NewUpdateProjectRequestNotificationsInner(channel string, service string, ) *UpdateProjectRequestNotificationsInner`

NewUpdateProjectRequestNotificationsInner instantiates a new UpdateProjectRequestNotificationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestNotificationsInnerWithDefaults

`func NewUpdateProjectRequestNotificationsInnerWithDefaults() *UpdateProjectRequestNotificationsInner`

NewUpdateProjectRequestNotificationsInnerWithDefaults instantiates a new UpdateProjectRequestNotificationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateProjectRequestNotificationsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateProjectRequestNotificationsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateProjectRequestNotificationsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateProjectRequestNotificationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChannel

`func (o *UpdateProjectRequestNotificationsInner) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateProjectRequestNotificationsInner) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateProjectRequestNotificationsInner) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetService

`func (o *UpdateProjectRequestNotificationsInner) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UpdateProjectRequestNotificationsInner) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UpdateProjectRequestNotificationsInner) SetService(v string)`

SetService sets Service field to given value.


### GetSubscribedUsers

`func (o *UpdateProjectRequestNotificationsInner) GetSubscribedUsers() []int32`

GetSubscribedUsers returns the SubscribedUsers field if non-nil, zero value otherwise.

### GetSubscribedUsersOk

`func (o *UpdateProjectRequestNotificationsInner) GetSubscribedUsersOk() (*[]int32, bool)`

GetSubscribedUsersOk returns a tuple with the SubscribedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedUsers

`func (o *UpdateProjectRequestNotificationsInner) SetSubscribedUsers(v []int32)`

SetSubscribedUsers sets SubscribedUsers field to given value.

### HasSubscribedUsers

`func (o *UpdateProjectRequestNotificationsInner) HasSubscribedUsers() bool`

HasSubscribedUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


