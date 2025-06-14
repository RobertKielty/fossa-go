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

// checks if the GetIssue200ResponseOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssue200ResponseOneOf2{}

// GetIssue200ResponseOneOf2 struct for GetIssue200ResponseOneOf2
type GetIssue200ResponseOneOf2 struct {
	Id *int32 `json:"id,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Source *GetIssue200ResponseOneOfAllOfSource1 `json:"source,omitempty"`
	Depths *GetIssue200ResponseOneOfAllOfDepths `json:"depths,omitempty"`
	ContainerLayers *GetIssue200ResponseOneOfAllOfContainerLayers `json:"containerLayers,omitempty"`
	Statuses *GetIssueStatuses200ResponseIssues `json:"statuses,omitempty"`
	Projects []GetIssue200ResponseOneOfAllOfProjectsInner `json:"projects,omitempty"`
	Type *string `json:"type,omitempty"`
	QualityRule *GetIssue200ResponseOneOf2AllOfQualityRule `json:"qualityRule,omitempty"`
}

// NewGetIssue200ResponseOneOf2 instantiates a new GetIssue200ResponseOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssue200ResponseOneOf2() *GetIssue200ResponseOneOf2 {
	this := GetIssue200ResponseOneOf2{}
	return &this
}

// NewGetIssue200ResponseOneOf2WithDefaults instantiates a new GetIssue200ResponseOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssue200ResponseOneOf2WithDefaults() *GetIssue200ResponseOneOf2 {
	this := GetIssue200ResponseOneOf2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetIssue200ResponseOneOf2) SetId(v int32) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetIssue200ResponseOneOf2) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetSource() GetIssue200ResponseOneOfAllOfSource1 {
	if o == nil || IsNil(o.Source) {
		var ret GetIssue200ResponseOneOfAllOfSource1
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource1, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GetIssue200ResponseOneOfAllOfSource1 and assigns it to the Source field.
func (o *GetIssue200ResponseOneOf2) SetSource(v GetIssue200ResponseOneOfAllOfSource1) {
	o.Source = &v
}

// GetDepths returns the Depths field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetDepths() GetIssue200ResponseOneOfAllOfDepths {
	if o == nil || IsNil(o.Depths) {
		var ret GetIssue200ResponseOneOfAllOfDepths
		return ret
	}
	return *o.Depths
}

// GetDepthsOk returns a tuple with the Depths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool) {
	if o == nil || IsNil(o.Depths) {
		return nil, false
	}
	return o.Depths, true
}

// HasDepths returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasDepths() bool {
	if o != nil && !IsNil(o.Depths) {
		return true
	}

	return false
}

// SetDepths gets a reference to the given GetIssue200ResponseOneOfAllOfDepths and assigns it to the Depths field.
func (o *GetIssue200ResponseOneOf2) SetDepths(v GetIssue200ResponseOneOfAllOfDepths) {
	o.Depths = &v
}

// GetContainerLayers returns the ContainerLayers field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers {
	if o == nil || IsNil(o.ContainerLayers) {
		var ret GetIssue200ResponseOneOfAllOfContainerLayers
		return ret
	}
	return *o.ContainerLayers
}

// GetContainerLayersOk returns a tuple with the ContainerLayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool) {
	if o == nil || IsNil(o.ContainerLayers) {
		return nil, false
	}
	return o.ContainerLayers, true
}

// HasContainerLayers returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasContainerLayers() bool {
	if o != nil && !IsNil(o.ContainerLayers) {
		return true
	}

	return false
}

// SetContainerLayers gets a reference to the given GetIssue200ResponseOneOfAllOfContainerLayers and assigns it to the ContainerLayers field.
func (o *GetIssue200ResponseOneOf2) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers) {
	o.ContainerLayers = &v
}

// GetStatuses returns the Statuses field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetStatuses() GetIssueStatuses200ResponseIssues {
	if o == nil || IsNil(o.Statuses) {
		var ret GetIssueStatuses200ResponseIssues
		return ret
	}
	return *o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// HasStatuses returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasStatuses() bool {
	if o != nil && !IsNil(o.Statuses) {
		return true
	}

	return false
}

// SetStatuses gets a reference to the given GetIssueStatuses200ResponseIssues and assigns it to the Statuses field.
func (o *GetIssue200ResponseOneOf2) SetStatuses(v GetIssueStatuses200ResponseIssues) {
	o.Statuses = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner {
	if o == nil || IsNil(o.Projects) {
		var ret []GetIssue200ResponseOneOfAllOfProjectsInner
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetProjectsOk() ([]GetIssue200ResponseOneOfAllOfProjectsInner, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []GetIssue200ResponseOneOfAllOfProjectsInner and assigns it to the Projects field.
func (o *GetIssue200ResponseOneOf2) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner) {
	o.Projects = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetIssue200ResponseOneOf2) SetType(v string) {
	o.Type = &v
}

// GetQualityRule returns the QualityRule field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOf2) GetQualityRule() GetIssue200ResponseOneOf2AllOfQualityRule {
	if o == nil || IsNil(o.QualityRule) {
		var ret GetIssue200ResponseOneOf2AllOfQualityRule
		return ret
	}
	return *o.QualityRule
}

// GetQualityRuleOk returns a tuple with the QualityRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOf2) GetQualityRuleOk() (*GetIssue200ResponseOneOf2AllOfQualityRule, bool) {
	if o == nil || IsNil(o.QualityRule) {
		return nil, false
	}
	return o.QualityRule, true
}

// HasQualityRule returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOf2) HasQualityRule() bool {
	if o != nil && !IsNil(o.QualityRule) {
		return true
	}

	return false
}

// SetQualityRule gets a reference to the given GetIssue200ResponseOneOf2AllOfQualityRule and assigns it to the QualityRule field.
func (o *GetIssue200ResponseOneOf2) SetQualityRule(v GetIssue200ResponseOneOf2AllOfQualityRule) {
	o.QualityRule = &v
}

func (o GetIssue200ResponseOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssue200ResponseOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Depths) {
		toSerialize["depths"] = o.Depths
	}
	if !IsNil(o.ContainerLayers) {
		toSerialize["containerLayers"] = o.ContainerLayers
	}
	if !IsNil(o.Statuses) {
		toSerialize["statuses"] = o.Statuses
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.QualityRule) {
		toSerialize["qualityRule"] = o.QualityRule
	}
	return toSerialize, nil
}

type NullableGetIssue200ResponseOneOf2 struct {
	value *GetIssue200ResponseOneOf2
	isSet bool
}

func (v NullableGetIssue200ResponseOneOf2) Get() *GetIssue200ResponseOneOf2 {
	return v.value
}

func (v *NullableGetIssue200ResponseOneOf2) Set(val *GetIssue200ResponseOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssue200ResponseOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssue200ResponseOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssue200ResponseOneOf2(val *GetIssue200ResponseOneOf2) *NullableGetIssue200ResponseOneOf2 {
	return &NullableGetIssue200ResponseOneOf2{value: val, isSet: true}
}

func (v NullableGetIssue200ResponseOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssue200ResponseOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


