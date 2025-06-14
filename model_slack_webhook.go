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

// checks if the SlackWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlackWebhook{}

// SlackWebhook struct for SlackWebhook
type SlackWebhook struct {
	// The Slack channel to send updates to
	Channel *string `json:"channel,omitempty"`
	// The configured webhook url in your Slack Application
	WebhookUrl *string `json:"webhook_url,omitempty"`
}

// NewSlackWebhook instantiates a new SlackWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlackWebhook() *SlackWebhook {
	this := SlackWebhook{}
	return &this
}

// NewSlackWebhookWithDefaults instantiates a new SlackWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlackWebhookWithDefaults() *SlackWebhook {
	this := SlackWebhook{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *SlackWebhook) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackWebhook) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *SlackWebhook) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *SlackWebhook) SetChannel(v string) {
	o.Channel = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *SlackWebhook) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackWebhook) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *SlackWebhook) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *SlackWebhook) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o SlackWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlackWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	return toSerialize, nil
}

type NullableSlackWebhook struct {
	value *SlackWebhook
	isSet bool
}

func (v NullableSlackWebhook) Get() *SlackWebhook {
	return v.value
}

func (v *NullableSlackWebhook) Set(val *SlackWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableSlackWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableSlackWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlackWebhook(val *SlackWebhook) *NullableSlackWebhook {
	return &NullableSlackWebhook{value: val, isSet: true}
}

func (v NullableSlackWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlackWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


