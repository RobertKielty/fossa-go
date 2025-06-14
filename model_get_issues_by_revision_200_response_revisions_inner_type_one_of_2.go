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

// checks if the GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2{}

// GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 struct for GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
type GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 struct {
	Critical *int32 `json:"critical,omitempty"`
	High *int32 `json:"high,omitempty"`
	Medium *int32 `json:"medium,omitempty"`
	Low *int32 `json:"low,omitempty"`
	Unknown *int32 `json:"unknown,omitempty"`
}

// NewGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 instantiates a new GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2() *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 {
	this := GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2{}
	return &this
}

// NewGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2WithDefaults instantiates a new GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2WithDefaults() *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 {
	this := GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetCritical() int32 {
	if o == nil || IsNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetCriticalOk() (*int32, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) SetCritical(v int32) {
	o.Critical = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetHigh() int32 {
	if o == nil || IsNil(o.High) {
		var ret int32
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetHighOk() (*int32, bool) {
	if o == nil || IsNil(o.High) {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) HasHigh() bool {
	if o != nil && !IsNil(o.High) {
		return true
	}

	return false
}

// SetHigh gets a reference to the given int32 and assigns it to the High field.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) SetHigh(v int32) {
	o.High = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetMedium() int32 {
	if o == nil || IsNil(o.Medium) {
		var ret int32
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetMediumOk() (*int32, bool) {
	if o == nil || IsNil(o.Medium) {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) HasMedium() bool {
	if o != nil && !IsNil(o.Medium) {
		return true
	}

	return false
}

// SetMedium gets a reference to the given int32 and assigns it to the Medium field.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) SetMedium(v int32) {
	o.Medium = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetLow() int32 {
	if o == nil || IsNil(o.Low) {
		var ret int32
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetLowOk() (*int32, bool) {
	if o == nil || IsNil(o.Low) {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) HasLow() bool {
	if o != nil && !IsNil(o.Low) {
		return true
	}

	return false
}

// SetLow gets a reference to the given int32 and assigns it to the Low field.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) SetLow(v int32) {
	o.Low = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetUnknown() int32 {
	if o == nil || IsNil(o.Unknown) {
		var ret int32
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) GetUnknownOk() (*int32, bool) {
	if o == nil || IsNil(o.Unknown) {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) HasUnknown() bool {
	if o != nil && !IsNil(o.Unknown) {
		return true
	}

	return false
}

// SetUnknown gets a reference to the given int32 and assigns it to the Unknown field.
func (o *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) SetUnknown(v int32) {
	o.Unknown = &v
}

func (o GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !IsNil(o.High) {
		toSerialize["high"] = o.High
	}
	if !IsNil(o.Medium) {
		toSerialize["medium"] = o.Medium
	}
	if !IsNil(o.Low) {
		toSerialize["low"] = o.Low
	}
	if !IsNil(o.Unknown) {
		toSerialize["unknown"] = o.Unknown
	}
	return toSerialize, nil
}

type NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 struct {
	value *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2
	isSet bool
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) Get() *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 {
	return v.value
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) Set(val *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2(val *GetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) *NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2 {
	return &NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2{value: val, isSet: true}
}

func (v NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssuesByRevision200ResponseRevisionsInnerTypeOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


