/*
FOSSA API

OpenAPI Specification for public FOSSA APIs

API version: 4.28.61
Contact: support@fossa.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fossa

import (
	"encoding/json"
)

// checks if the NotificationsSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettings{}

// NotificationsSettings struct for NotificationsSettings
type NotificationsSettings struct {
	// the on/off status of project notifications
	NotificationDefaultEnabled *bool `json:"notificationDefaultEnabled,omitempty"`
	// the on/off status of slack notifications for scans
	NotificationDefaultSlackScan *bool `json:"notificationDefaultSlackScan,omitempty"`
	NotificationDefaultEmailScanUsers []float32 `json:"notificationDefaultEmailScanUsers,omitempty"`
	// The type of users to email for scan notifications. Each option represents a different set of users: - current: Only the current user will receive scan notifications. - all: All users will receive scan notifications. - custom: Custom set of users will receive scan notifications. 
	NotificationDefaultEmailScanUserType *string `json:"notificationDefaultEmailScanUserType,omitempty"`
}

// NewNotificationsSettings instantiates a new NotificationsSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettings() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsWithDefaults() *NotificationsSettings {
	this := NotificationsSettings{}
	return &this
}

// GetNotificationDefaultEnabled returns the NotificationDefaultEnabled field value if set, zero value otherwise.
func (o *NotificationsSettings) GetNotificationDefaultEnabled() bool {
	if o == nil || IsNil(o.NotificationDefaultEnabled) {
		var ret bool
		return ret
	}
	return *o.NotificationDefaultEnabled
}

// GetNotificationDefaultEnabledOk returns a tuple with the NotificationDefaultEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetNotificationDefaultEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.NotificationDefaultEnabled) {
		return nil, false
	}
	return o.NotificationDefaultEnabled, true
}

// HasNotificationDefaultEnabled returns a boolean if a field has been set.
func (o *NotificationsSettings) HasNotificationDefaultEnabled() bool {
	if o != nil && !IsNil(o.NotificationDefaultEnabled) {
		return true
	}

	return false
}

// SetNotificationDefaultEnabled gets a reference to the given bool and assigns it to the NotificationDefaultEnabled field.
func (o *NotificationsSettings) SetNotificationDefaultEnabled(v bool) {
	o.NotificationDefaultEnabled = &v
}

// GetNotificationDefaultSlackScan returns the NotificationDefaultSlackScan field value if set, zero value otherwise.
func (o *NotificationsSettings) GetNotificationDefaultSlackScan() bool {
	if o == nil || IsNil(o.NotificationDefaultSlackScan) {
		var ret bool
		return ret
	}
	return *o.NotificationDefaultSlackScan
}

// GetNotificationDefaultSlackScanOk returns a tuple with the NotificationDefaultSlackScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetNotificationDefaultSlackScanOk() (*bool, bool) {
	if o == nil || IsNil(o.NotificationDefaultSlackScan) {
		return nil, false
	}
	return o.NotificationDefaultSlackScan, true
}

// HasNotificationDefaultSlackScan returns a boolean if a field has been set.
func (o *NotificationsSettings) HasNotificationDefaultSlackScan() bool {
	if o != nil && !IsNil(o.NotificationDefaultSlackScan) {
		return true
	}

	return false
}

// SetNotificationDefaultSlackScan gets a reference to the given bool and assigns it to the NotificationDefaultSlackScan field.
func (o *NotificationsSettings) SetNotificationDefaultSlackScan(v bool) {
	o.NotificationDefaultSlackScan = &v
}

// GetNotificationDefaultEmailScanUsers returns the NotificationDefaultEmailScanUsers field value if set, zero value otherwise.
func (o *NotificationsSettings) GetNotificationDefaultEmailScanUsers() []float32 {
	if o == nil || IsNil(o.NotificationDefaultEmailScanUsers) {
		var ret []float32
		return ret
	}
	return o.NotificationDefaultEmailScanUsers
}

// GetNotificationDefaultEmailScanUsersOk returns a tuple with the NotificationDefaultEmailScanUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetNotificationDefaultEmailScanUsersOk() ([]float32, bool) {
	if o == nil || IsNil(o.NotificationDefaultEmailScanUsers) {
		return nil, false
	}
	return o.NotificationDefaultEmailScanUsers, true
}

// HasNotificationDefaultEmailScanUsers returns a boolean if a field has been set.
func (o *NotificationsSettings) HasNotificationDefaultEmailScanUsers() bool {
	if o != nil && !IsNil(o.NotificationDefaultEmailScanUsers) {
		return true
	}

	return false
}

// SetNotificationDefaultEmailScanUsers gets a reference to the given []float32 and assigns it to the NotificationDefaultEmailScanUsers field.
func (o *NotificationsSettings) SetNotificationDefaultEmailScanUsers(v []float32) {
	o.NotificationDefaultEmailScanUsers = v
}

// GetNotificationDefaultEmailScanUserType returns the NotificationDefaultEmailScanUserType field value if set, zero value otherwise.
func (o *NotificationsSettings) GetNotificationDefaultEmailScanUserType() string {
	if o == nil || IsNil(o.NotificationDefaultEmailScanUserType) {
		var ret string
		return ret
	}
	return *o.NotificationDefaultEmailScanUserType
}

// GetNotificationDefaultEmailScanUserTypeOk returns a tuple with the NotificationDefaultEmailScanUserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettings) GetNotificationDefaultEmailScanUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDefaultEmailScanUserType) {
		return nil, false
	}
	return o.NotificationDefaultEmailScanUserType, true
}

// HasNotificationDefaultEmailScanUserType returns a boolean if a field has been set.
func (o *NotificationsSettings) HasNotificationDefaultEmailScanUserType() bool {
	if o != nil && !IsNil(o.NotificationDefaultEmailScanUserType) {
		return true
	}

	return false
}

// SetNotificationDefaultEmailScanUserType gets a reference to the given string and assigns it to the NotificationDefaultEmailScanUserType field.
func (o *NotificationsSettings) SetNotificationDefaultEmailScanUserType(v string) {
	o.NotificationDefaultEmailScanUserType = &v
}

func (o NotificationsSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationDefaultEnabled) {
		toSerialize["notificationDefaultEnabled"] = o.NotificationDefaultEnabled
	}
	if !IsNil(o.NotificationDefaultSlackScan) {
		toSerialize["notificationDefaultSlackScan"] = o.NotificationDefaultSlackScan
	}
	if !IsNil(o.NotificationDefaultEmailScanUsers) {
		toSerialize["notificationDefaultEmailScanUsers"] = o.NotificationDefaultEmailScanUsers
	}
	if !IsNil(o.NotificationDefaultEmailScanUserType) {
		toSerialize["notificationDefaultEmailScanUserType"] = o.NotificationDefaultEmailScanUserType
	}
	return toSerialize, nil
}

type NullableNotificationsSettings struct {
	value *NotificationsSettings
	isSet bool
}

func (v NullableNotificationsSettings) Get() *NotificationsSettings {
	return v.value
}

func (v *NullableNotificationsSettings) Set(val *NotificationsSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettings(val *NotificationsSettings) *NullableNotificationsSettings {
	return &NullableNotificationsSettings{value: val, isSet: true}
}

func (v NullableNotificationsSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


