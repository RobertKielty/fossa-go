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

// checks if the GetProjectsSummary200ResponseSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectsSummary200ResponseSummary{}

// GetProjectsSummary200ResponseSummary struct for GetProjectsSummary200ResponseSummary
type GetProjectsSummary200ResponseSummary struct {
	Projects *int32 `json:"projects,omitempty"`
	ReleaseGroups *int32 `json:"releaseGroups,omitempty"`
}

// NewGetProjectsSummary200ResponseSummary instantiates a new GetProjectsSummary200ResponseSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectsSummary200ResponseSummary() *GetProjectsSummary200ResponseSummary {
	this := GetProjectsSummary200ResponseSummary{}
	return &this
}

// NewGetProjectsSummary200ResponseSummaryWithDefaults instantiates a new GetProjectsSummary200ResponseSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectsSummary200ResponseSummaryWithDefaults() *GetProjectsSummary200ResponseSummary {
	this := GetProjectsSummary200ResponseSummary{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GetProjectsSummary200ResponseSummary) GetProjects() int32 {
	if o == nil || IsNil(o.Projects) {
		var ret int32
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectsSummary200ResponseSummary) GetProjectsOk() (*int32, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GetProjectsSummary200ResponseSummary) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given int32 and assigns it to the Projects field.
func (o *GetProjectsSummary200ResponseSummary) SetProjects(v int32) {
	o.Projects = &v
}

// GetReleaseGroups returns the ReleaseGroups field value if set, zero value otherwise.
func (o *GetProjectsSummary200ResponseSummary) GetReleaseGroups() int32 {
	if o == nil || IsNil(o.ReleaseGroups) {
		var ret int32
		return ret
	}
	return *o.ReleaseGroups
}

// GetReleaseGroupsOk returns a tuple with the ReleaseGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectsSummary200ResponseSummary) GetReleaseGroupsOk() (*int32, bool) {
	if o == nil || IsNil(o.ReleaseGroups) {
		return nil, false
	}
	return o.ReleaseGroups, true
}

// HasReleaseGroups returns a boolean if a field has been set.
func (o *GetProjectsSummary200ResponseSummary) HasReleaseGroups() bool {
	if o != nil && !IsNil(o.ReleaseGroups) {
		return true
	}

	return false
}

// SetReleaseGroups gets a reference to the given int32 and assigns it to the ReleaseGroups field.
func (o *GetProjectsSummary200ResponseSummary) SetReleaseGroups(v int32) {
	o.ReleaseGroups = &v
}

func (o GetProjectsSummary200ResponseSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectsSummary200ResponseSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.ReleaseGroups) {
		toSerialize["releaseGroups"] = o.ReleaseGroups
	}
	return toSerialize, nil
}

type NullableGetProjectsSummary200ResponseSummary struct {
	value *GetProjectsSummary200ResponseSummary
	isSet bool
}

func (v NullableGetProjectsSummary200ResponseSummary) Get() *GetProjectsSummary200ResponseSummary {
	return v.value
}

func (v *NullableGetProjectsSummary200ResponseSummary) Set(val *GetProjectsSummary200ResponseSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectsSummary200ResponseSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectsSummary200ResponseSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectsSummary200ResponseSummary(val *GetProjectsSummary200ResponseSummary) *NullableGetProjectsSummary200ResponseSummary {
	return &NullableGetProjectsSummary200ResponseSummary{value: val, isSet: true}
}

func (v NullableGetProjectsSummary200ResponseSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectsSummary200ResponseSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


