# GetOrganizationNotificationSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationDefaultEnabled** | **bool** | the on/off status of project notifications | 
**NotificationDefaultSlackScan** | **bool** | the on/off status of slack notifications for scans | 
**NotificationDefaultApiWebhookScan** | **bool** | the on/off status of API webhook notifications for scans | 
**NotificationDefaultEmailScanUsers** | **[]float32** |  | 
**NotificationDefaultEmailScanUserType** | **string** | The type of users to email for scan notifications. Each option represents a different set of users: - current: Only the current user will receive scan notifications. - all: All users will receive scan notifications. - custom: Custom set of users will receive scan notifications.  | 

## Methods

### NewGetOrganizationNotificationSettings200Response

`func NewGetOrganizationNotificationSettings200Response(notificationDefaultEnabled bool, notificationDefaultSlackScan bool, notificationDefaultApiWebhookScan bool, notificationDefaultEmailScanUsers []float32, notificationDefaultEmailScanUserType string, ) *GetOrganizationNotificationSettings200Response`

NewGetOrganizationNotificationSettings200Response instantiates a new GetOrganizationNotificationSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationNotificationSettings200ResponseWithDefaults

`func NewGetOrganizationNotificationSettings200ResponseWithDefaults() *GetOrganizationNotificationSettings200Response`

NewGetOrganizationNotificationSettings200ResponseWithDefaults instantiates a new GetOrganizationNotificationSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationDefaultEnabled

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEnabled() bool`

GetNotificationDefaultEnabled returns the NotificationDefaultEnabled field if non-nil, zero value otherwise.

### GetNotificationDefaultEnabledOk

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEnabledOk() (*bool, bool)`

GetNotificationDefaultEnabledOk returns a tuple with the NotificationDefaultEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEnabled

`func (o *GetOrganizationNotificationSettings200Response) SetNotificationDefaultEnabled(v bool)`

SetNotificationDefaultEnabled sets NotificationDefaultEnabled field to given value.


### GetNotificationDefaultSlackScan

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultSlackScan() bool`

GetNotificationDefaultSlackScan returns the NotificationDefaultSlackScan field if non-nil, zero value otherwise.

### GetNotificationDefaultSlackScanOk

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultSlackScanOk() (*bool, bool)`

GetNotificationDefaultSlackScanOk returns a tuple with the NotificationDefaultSlackScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultSlackScan

`func (o *GetOrganizationNotificationSettings200Response) SetNotificationDefaultSlackScan(v bool)`

SetNotificationDefaultSlackScan sets NotificationDefaultSlackScan field to given value.


### GetNotificationDefaultApiWebhookScan

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultApiWebhookScan() bool`

GetNotificationDefaultApiWebhookScan returns the NotificationDefaultApiWebhookScan field if non-nil, zero value otherwise.

### GetNotificationDefaultApiWebhookScanOk

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultApiWebhookScanOk() (*bool, bool)`

GetNotificationDefaultApiWebhookScanOk returns a tuple with the NotificationDefaultApiWebhookScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultApiWebhookScan

`func (o *GetOrganizationNotificationSettings200Response) SetNotificationDefaultApiWebhookScan(v bool)`

SetNotificationDefaultApiWebhookScan sets NotificationDefaultApiWebhookScan field to given value.


### GetNotificationDefaultEmailScanUsers

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEmailScanUsers() []float32`

GetNotificationDefaultEmailScanUsers returns the NotificationDefaultEmailScanUsers field if non-nil, zero value otherwise.

### GetNotificationDefaultEmailScanUsersOk

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEmailScanUsersOk() (*[]float32, bool)`

GetNotificationDefaultEmailScanUsersOk returns a tuple with the NotificationDefaultEmailScanUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEmailScanUsers

`func (o *GetOrganizationNotificationSettings200Response) SetNotificationDefaultEmailScanUsers(v []float32)`

SetNotificationDefaultEmailScanUsers sets NotificationDefaultEmailScanUsers field to given value.


### GetNotificationDefaultEmailScanUserType

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEmailScanUserType() string`

GetNotificationDefaultEmailScanUserType returns the NotificationDefaultEmailScanUserType field if non-nil, zero value otherwise.

### GetNotificationDefaultEmailScanUserTypeOk

`func (o *GetOrganizationNotificationSettings200Response) GetNotificationDefaultEmailScanUserTypeOk() (*string, bool)`

GetNotificationDefaultEmailScanUserTypeOk returns a tuple with the NotificationDefaultEmailScanUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationDefaultEmailScanUserType

`func (o *GetOrganizationNotificationSettings200Response) SetNotificationDefaultEmailScanUserType(v string)`

SetNotificationDefaultEmailScanUserType sets NotificationDefaultEmailScanUserType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


