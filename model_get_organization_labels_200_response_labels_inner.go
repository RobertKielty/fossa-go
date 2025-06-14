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

// checks if the GetOrganizationLabels200ResponseLabelsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLabels200ResponseLabelsInner{}

// GetOrganizationLabels200ResponseLabelsInner struct for GetOrganizationLabels200ResponseLabelsInner
type GetOrganizationLabels200ResponseLabelsInner struct {
	// ID of this label
	Id *int32 `json:"id,omitempty"`
	// ID of the organization that this Label is associated with
	OrganizationId *int32 `json:"organizationId,omitempty"`
	// Text that this Label represents
	Label *string `json:"label,omitempty"`
	Projects []string `json:"projects,omitempty"`
}

// NewGetOrganizationLabels200ResponseLabelsInner instantiates a new GetOrganizationLabels200ResponseLabelsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLabels200ResponseLabelsInner() *GetOrganizationLabels200ResponseLabelsInner {
	this := GetOrganizationLabels200ResponseLabelsInner{}
	return &this
}

// NewGetOrganizationLabels200ResponseLabelsInnerWithDefaults instantiates a new GetOrganizationLabels200ResponseLabelsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLabels200ResponseLabelsInnerWithDefaults() *GetOrganizationLabels200ResponseLabelsInner {
	this := GetOrganizationLabels200ResponseLabelsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetOrganizationLabels200ResponseLabelsInner) SetId(v int32) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *GetOrganizationLabels200ResponseLabelsInner) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetOrganizationLabels200ResponseLabelsInner) SetLabel(v string) {
	o.Label = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetProjects() []string {
	if o == nil || IsNil(o.Projects) {
		var ret []string
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) GetProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GetOrganizationLabels200ResponseLabelsInner) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []string and assigns it to the Projects field.
func (o *GetOrganizationLabels200ResponseLabelsInner) SetProjects(v []string) {
	o.Projects = v
}

func (o GetOrganizationLabels200ResponseLabelsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLabels200ResponseLabelsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	return toSerialize, nil
}

type NullableGetOrganizationLabels200ResponseLabelsInner struct {
	value *GetOrganizationLabels200ResponseLabelsInner
	isSet bool
}

func (v NullableGetOrganizationLabels200ResponseLabelsInner) Get() *GetOrganizationLabels200ResponseLabelsInner {
	return v.value
}

func (v *NullableGetOrganizationLabels200ResponseLabelsInner) Set(val *GetOrganizationLabels200ResponseLabelsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLabels200ResponseLabelsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLabels200ResponseLabelsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLabels200ResponseLabelsInner(val *GetOrganizationLabels200ResponseLabelsInner) *NullableGetOrganizationLabels200ResponseLabelsInner {
	return &NullableGetOrganizationLabels200ResponseLabelsInner{value: val, isSet: true}
}

func (v NullableGetOrganizationLabels200ResponseLabelsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLabels200ResponseLabelsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


