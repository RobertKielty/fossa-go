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
	"time"
)

// checks if the GetIssueCounts200ResponseCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssueCounts200ResponseCounts{}

// GetIssueCounts200ResponseCounts An array of daily issue count summaries.
type GetIssueCounts200ResponseCounts struct {
	Date *time.Time `json:"date,omitempty"`
	Active *int32 `json:"active,omitempty"`
	Ignored *int32 `json:"ignored,omitempty"`
	Remediated *int32 `json:"remediated,omitempty"`
}

// NewGetIssueCounts200ResponseCounts instantiates a new GetIssueCounts200ResponseCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssueCounts200ResponseCounts() *GetIssueCounts200ResponseCounts {
	this := GetIssueCounts200ResponseCounts{}
	return &this
}

// NewGetIssueCounts200ResponseCountsWithDefaults instantiates a new GetIssueCounts200ResponseCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssueCounts200ResponseCountsWithDefaults() *GetIssueCounts200ResponseCounts {
	this := GetIssueCounts200ResponseCounts{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseCounts) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseCounts) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseCounts) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *GetIssueCounts200ResponseCounts) SetDate(v time.Time) {
	o.Date = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseCounts) GetActive() int32 {
	if o == nil || IsNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseCounts) GetActiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseCounts) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *GetIssueCounts200ResponseCounts) SetActive(v int32) {
	o.Active = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseCounts) GetIgnored() int32 {
	if o == nil || IsNil(o.Ignored) {
		var ret int32
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseCounts) GetIgnoredOk() (*int32, bool) {
	if o == nil || IsNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseCounts) HasIgnored() bool {
	if o != nil && !IsNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given int32 and assigns it to the Ignored field.
func (o *GetIssueCounts200ResponseCounts) SetIgnored(v int32) {
	o.Ignored = &v
}

// GetRemediated returns the Remediated field value if set, zero value otherwise.
func (o *GetIssueCounts200ResponseCounts) GetRemediated() int32 {
	if o == nil || IsNil(o.Remediated) {
		var ret int32
		return ret
	}
	return *o.Remediated
}

// GetRemediatedOk returns a tuple with the Remediated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueCounts200ResponseCounts) GetRemediatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Remediated) {
		return nil, false
	}
	return o.Remediated, true
}

// HasRemediated returns a boolean if a field has been set.
func (o *GetIssueCounts200ResponseCounts) HasRemediated() bool {
	if o != nil && !IsNil(o.Remediated) {
		return true
	}

	return false
}

// SetRemediated gets a reference to the given int32 and assigns it to the Remediated field.
func (o *GetIssueCounts200ResponseCounts) SetRemediated(v int32) {
	o.Remediated = &v
}

func (o GetIssueCounts200ResponseCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssueCounts200ResponseCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Ignored) {
		toSerialize["ignored"] = o.Ignored
	}
	if !IsNil(o.Remediated) {
		toSerialize["remediated"] = o.Remediated
	}
	return toSerialize, nil
}

type NullableGetIssueCounts200ResponseCounts struct {
	value *GetIssueCounts200ResponseCounts
	isSet bool
}

func (v NullableGetIssueCounts200ResponseCounts) Get() *GetIssueCounts200ResponseCounts {
	return v.value
}

func (v *NullableGetIssueCounts200ResponseCounts) Set(val *GetIssueCounts200ResponseCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssueCounts200ResponseCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssueCounts200ResponseCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssueCounts200ResponseCounts(val *GetIssueCounts200ResponseCounts) *NullableGetIssueCounts200ResponseCounts {
	return &NullableGetIssueCounts200ResponseCounts{value: val, isSet: true}
}

func (v NullableGetIssueCounts200ResponseCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssueCounts200ResponseCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


