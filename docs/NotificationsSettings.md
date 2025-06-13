# NotificationsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationDefaultEnabled** | Pointer to **bool** | the on/off status of project notifications | [optional] 
**NotificationDefaultSlackScan** | Pointer to **bool** | the on/off status of slack notifications for scans | [optional] 
**NotificationDefaultEmailScanUsers** | Pointer to **[]float32** |  | [optional] 
**NotificationDefaultEmailScanUserType** | Pointer to **string** | The type of users to email for scan notifications. Each option represents a different set of users: - current: Only the current user will receive scan notifications. - all: All users will receive scan notifications. - custom: Custom set of users will receive scan notifications.  | [optional] 

## Methods

### NewNotificationsSettings

`func NewNotificationsSettings() *NotificationsSettings`

NewNotificationsSettings instantiates a new NotificationsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsWithDefaults

`func NewNotificationsSettingsWithDefaults() *NotificationsSettings`

NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationDefaultEnabled

`func (o *NotificationsSettings) GetNotificationDefaultEnabled() bool`

GetNotificationDefaultEnabled returns the NotificationDefaultEnabled field if non-nil, zero value otherwise.

### GetNotificationDefaultEnabledOk

`func (o *NotificationsSettings) GetNotificationDefaultEnabledOk() (*bool, bool)`

GetNotificationDefaultEnabledOk returns a tuple with the NotificationDefaultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEnabled

`func (o *NotificationsSettings) SetNotificationDefaultEnabled(v bool)`

SetNotificationDefaultEnabled sets NotificationDefaultEnabled field to given value.

### HasNotificationDefaultEnabled

`func (o *NotificationsSettings) HasNotificationDefaultEnabled() bool`

HasNotificationDefaultEnabled returns a boolean if a field has been set.

### GetNotificationDefaultSlackScan

`func (o *NotificationsSettings) GetNotificationDefaultSlackScan() bool`

GetNotificationDefaultSlackScan returns the NotificationDefaultSlackScan field if non-nil, zero value otherwise.

### GetNotificationDefaultSlackScanOk

`func (o *NotificationsSettings) GetNotificationDefaultSlackScanOk() (*bool, bool)`

GetNotificationDefaultSlackScanOk returns a tuple with the NotificationDefaultSlackScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultSlackScan

`func (o *NotificationsSettings) SetNotificationDefaultSlackScan(v bool)`

SetNotificationDefaultSlackScan sets NotificationDefaultSlackScan field to given value.

### HasNotificationDefaultSlackScan

`func (o *NotificationsSettings) HasNotificationDefaultSlackScan() bool`

HasNotificationDefaultSlackScan returns a boolean if a field has been set.

### GetNotificationDefaultEmailScanUsers

`func (o *NotificationsSettings) GetNotificationDefaultEmailScanUsers() []float32`

GetNotificationDefaultEmailScanUsers returns the NotificationDefaultEmailScanUsers field if non-nil, zero value otherwise.

### GetNotificationDefaultEmailScanUsersOk

`func (o *NotificationsSettings) GetNotificationDefaultEmailScanUsersOk() (*[]float32, bool)`

GetNotificationDefaultEmailScanUsersOk returns a tuple with the NotificationDefaultEmailScanUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEmailScanUsers

`func (o *NotificationsSettings) SetNotificationDefaultEmailScanUsers(v []float32)`

SetNotificationDefaultEmailScanUsers sets NotificationDefaultEmailScanUsers field to given value.

### HasNotificationDefaultEmailScanUsers

`func (o *NotificationsSettings) HasNotificationDefaultEmailScanUsers() bool`

HasNotificationDefaultEmailScanUsers returns a boolean if a field has been set.

### GetNotificationDefaultEmailScanUserType

`func (o *NotificationsSettings) GetNotificationDefaultEmailScanUserType() string`

GetNotificationDefaultEmailScanUserType returns the NotificationDefaultEmailScanUserType field if non-nil, zero value otherwise.

### GetNotificationDefaultEmailScanUserTypeOk

`func (o *NotificationsSettings) GetNotificationDefaultEmailScanUserTypeOk() (*string, bool)`

GetNotificationDefaultEmailScanUserTypeOk returns a tuple with the NotificationDefaultEmailScanUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEmailScanUserType

`func (o *NotificationsSettings) SetNotificationDefaultEmailScanUserType(v string)`

SetNotificationDefaultEmailScanUserType sets NotificationDefaultEmailScanUserType field to given value.

### HasNotificationDefaultEmailScanUserType

`func (o *NotificationsSettings) HasNotificationDefaultEmailScanUserType() bool`

HasNotificationDefaultEmailScanUserType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


