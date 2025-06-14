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

// checks if the GetIssueCounts200ResponseQuality type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssueCounts200ResponseQuality{}

// GetIssueCounts200ResponseQuality struct for GetIssueCounts200ResponseQuality
type GetIssueCounts200ResponseQuality struct {
	Type *GetIssueCounts200ResponseQualityType `json:"type,omitempty"`
	Status *GetIssueCounts200ResponseQualityStatus `json:"status,omitempty"`
}

// NewGetIssueCounts200ResponseQuality instantiates a new GetIssueCounts200ResponseQuality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssueCounts200ResponseQuality() *GetIssueCounts200ResponseQuality {
	this := GetIssueCounts200ResponseQuality{}
	return &this
}

// NewGetIssueCounts200ResponseQualityWithDefaults instantiates a new GetIssueCounts200ResponseQuality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssueCounts200ResponseQualityWithDefaults() *GetIssueCounts200ResponseQuality {
	this := GetIssueCounts200ResponseQuality{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseQuality) GetType() GetIssueCounts200ResponseQualityType {
	if o == nil || IsNil(o.Type) {
		var ret GetIssueCounts200ResponseQualityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseQuality) GetTypeOk() (*GetIssueCounts200ResponseQualityType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseQuality) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given GetIssueCounts200ResponseQualityType and assigns it to the Type field.
func (o *GetIssueCounts200ResponseQuality) SetType(v GetIssueCounts200ResponseQualityType) {
	o.Type = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseQuality) GetStatus() GetIssueCounts200ResponseQualityStatus {
	if o == nil || IsNil(o.Status) {
		var ret GetIssueCounts200ResponseQualityStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseQuality) GetStatusOk() (*GetIssueCounts200ResponseQualityStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseQuality) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetIssueCounts200ResponseQualityStatus and assigns it to the Status field.
func (o *GetIssueCounts200ResponseQuality) SetStatus(v GetIssueCounts200ResponseQualityStatus) {
	o.Status = &v
}

func (o GetIssueCounts200ResponseQuality) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssueCounts200ResponseQuality) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetIssueCounts200ResponseQuality struct {
	value *GetIssueCounts200ResponseQuality
	isSet bool
}

func (v NullableGetIssueCounts200ResponseQuality) Get() *GetIssueCounts200ResponseQuality {
	return v.value
}

func (v *NullableGetIssueCounts200ResponseQuality) Set(val *GetIssueCounts200ResponseQuality) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssueCounts200ResponseQuality) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssueCounts200ResponseQuality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssueCounts200ResponseQuality(val *GetIssueCounts200ResponseQuality) *NullableGetIssueCounts200ResponseQuality {
	return &NullableGetIssueCounts200ResponseQuality{value: val, isSet: true}
}

func (v NullableGetIssueCounts200ResponseQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssueCounts200ResponseQuality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


