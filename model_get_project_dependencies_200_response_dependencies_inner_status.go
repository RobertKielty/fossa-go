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

// checks if the GetProjectDependencies200ResponseDependenciesInnerStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectDependencies200ResponseDependenciesInnerStatus{}

// GetProjectDependencies200ResponseDependenciesInnerStatus struct for GetProjectDependencies200ResponseDependenciesInnerStatus
type GetProjectDependencies200ResponseDependenciesInnerStatus struct {
	Error *string `json:"error,omitempty"`
	Resolved *bool `json:"resolved,omitempty"`
	Unsupported *bool `json:"unsupported,omitempty"`
	Analyzing *bool `json:"analyzing,omitempty"`
}

// NewGetProjectDependencies200ResponseDependenciesInnerStatus instantiates a new GetProjectDependencies200ResponseDependenciesInnerStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectDependencies200ResponseDependenciesInnerStatus() *GetProjectDependencies200ResponseDependenciesInnerStatus {
	this := GetProjectDependencies200ResponseDependenciesInnerStatus{}
	return &this
}

// NewGetProjectDependencies200ResponseDependenciesInnerStatusWithDefaults instantiates a new GetProjectDependencies200ResponseDependenciesInnerStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectDependencies200ResponseDependenciesInnerStatusWithDefaults() *GetProjectDependencies200ResponseDependenciesInnerStatus {
	this := GetProjectDependencies200ResponseDependenciesInnerStatus{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) SetError(v string) {
	o.Error = &v
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetResolved() bool {
	if o == nil || IsNil(o.Resolved) {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) SetResolved(v bool) {
	o.Resolved = &v
}

// GetUnsupported returns the Unsupported field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetUnsupported() bool {
	if o == nil || IsNil(o.Unsupported) {
		var ret bool
		return ret
	}
	return *o.Unsupported
}

// GetUnsupportedOk returns a tuple with the Unsupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetUnsupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.Unsupported) {
		return nil, false
	}
	return o.Unsupported, true
}

// HasUnsupported returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) HasUnsupported() bool {
	if o != nil && !IsNil(o.Unsupported) {
		return true
	}

	return false
}

// SetUnsupported gets a reference to the given bool and assigns it to the Unsupported field.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) SetUnsupported(v bool) {
	o.Unsupported = &v
}

// GetAnalyzing returns the Analyzing field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetAnalyzing() bool {
	if o == nil || IsNil(o.Analyzing) {
		var ret bool
		return ret
	}
	return *o.Analyzing
}

// GetAnalyzingOk returns a tuple with the Analyzing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) GetAnalyzingOk() (*bool, bool) {
	if o == nil || IsNil(o.Analyzing) {
		return nil, false
	}
	return o.Analyzing, true
}

// HasAnalyzing returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) HasAnalyzing() bool {
	if o != nil && !IsNil(o.Analyzing) {
		return true
	}

	return false
}

// SetAnalyzing gets a reference to the given bool and assigns it to the Analyzing field.
func (o *GetProjectDependencies200ResponseDependenciesInnerStatus) SetAnalyzing(v bool) {
	o.Analyzing = &v
}

func (o GetProjectDependencies200ResponseDependenciesInnerStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectDependencies200ResponseDependenciesInnerStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	if !IsNil(o.Unsupported) {
		toSerialize["unsupported"] = o.Unsupported
	}
	if !IsNil(o.Analyzing) {
		toSerialize["analyzing"] = o.Analyzing
	}
	return toSerialize, nil
}

type NullableGetProjectDependencies200ResponseDependenciesInnerStatus struct {
	value *GetProjectDependencies200ResponseDependenciesInnerStatus
	isSet bool
}

func (v NullableGetProjectDependencies200ResponseDependenciesInnerStatus) Get() *GetProjectDependencies200ResponseDependenciesInnerStatus {
	return v.value
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInnerStatus) Set(val *GetProjectDependencies200ResponseDependenciesInnerStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectDependencies200ResponseDependenciesInnerStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInnerStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectDependencies200ResponseDependenciesInnerStatus(val *GetProjectDependencies200ResponseDependenciesInnerStatus) *NullableGetProjectDependencies200ResponseDependenciesInnerStatus {
	return &NullableGetProjectDependencies200ResponseDependenciesInnerStatus{value: val, isSet: true}
}

func (v NullableGetProjectDependencies200ResponseDependenciesInnerStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInnerStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


