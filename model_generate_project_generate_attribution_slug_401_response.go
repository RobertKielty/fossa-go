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

// checks if the GenerateProjectGenerateAttributionSlug401Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateProjectGenerateAttributionSlug401Response{}

// GenerateProjectGenerateAttributionSlug401Response struct for GenerateProjectGenerateAttributionSlug401Response
type GenerateProjectGenerateAttributionSlug401Response struct {
	// Message describing the error
	Message *string `json:"message,omitempty"`
}

// NewGenerateProjectGenerateAttributionSlug401Response instantiates a new GenerateProjectGenerateAttributionSlug401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateProjectGenerateAttributionSlug401Response() *GenerateProjectGenerateAttributionSlug401Response {
	this := GenerateProjectGenerateAttributionSlug401Response{}
	return &this
}

// NewGenerateProjectGenerateAttributionSlug401ResponseWithDefaults instantiates a new GenerateProjectGenerateAttributionSlug401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateProjectGenerateAttributionSlug401ResponseWithDefaults() *GenerateProjectGenerateAttributionSlug401Response {
	this := GenerateProjectGenerateAttributionSlug401Response{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GenerateProjectGenerateAttributionSlug401Response) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateProjectGenerateAttributionSlug401Response) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GenerateProjectGenerateAttributionSlug401Response) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GenerateProjectGenerateAttributionSlug401Response) SetMessage(v string) {
	o.Message = &v
}

func (o GenerateProjectGenerateAttributionSlug401Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenerateProjectGenerateAttributionSlug401Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableGenerateProjectGenerateAttributionSlug401Response struct {
	value *GenerateProjectGenerateAttributionSlug401Response
	isSet bool
}

func (v NullableGenerateProjectGenerateAttributionSlug401Response) Get() *GenerateProjectGenerateAttributionSlug401Response {
	return v.value
}

func (v *NullableGenerateProjectGenerateAttributionSlug401Response) Set(val *GenerateProjectGenerateAttributionSlug401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateProjectGenerateAttributionSlug401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateProjectGenerateAttributionSlug401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateProjectGenerateAttributionSlug401Response(val *GenerateProjectGenerateAttributionSlug401Response) *NullableGenerateProjectGenerateAttributionSlug401Response {
	return &NullableGenerateProjectGenerateAttributionSlug401Response{value: val, isSet: true}
}

func (v NullableGenerateProjectGenerateAttributionSlug401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateProjectGenerateAttributionSlug401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


