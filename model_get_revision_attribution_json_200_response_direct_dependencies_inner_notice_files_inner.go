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

// checks if the GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner{}

// GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner struct for GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner
type GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner struct {
	// The ID of the notice match
	Id *int32 `json:"id,omitempty"`
	// The revision ID of the notice match
	RevisionId *string `json:"revisionId,omitempty"`
	// The contents of the notice match
	Contents *string `json:"contents,omitempty"`
	// The copyrights of the notice match
	Copyrights []string `json:"copyrights,omitempty"`
	// The creation date of the notice match
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The update date of the notice match
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Whether the notice match has been manually corrected
	Corrected *bool `json:"corrected,omitempty"`
	// Whether the notice match has been manually ignored
	Ignored *bool `json:"ignored,omitempty"`
}

// NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner instantiates a new GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner() *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner {
	this := GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner{}
	return &this
}

// NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInnerWithDefaults instantiates a new GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInnerWithDefaults() *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner {
	this := GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetId(v int32) {
	o.Id = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetRevisionId() string {
	if o == nil || IsNil(o.RevisionId) {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetRevisionIdOk() (*string, bool) {
	if o == nil || IsNil(o.RevisionId) {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasRevisionId() bool {
	if o != nil && !IsNil(o.RevisionId) {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetContents() string {
	if o == nil || IsNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetContentsOk() (*string, bool) {
	if o == nil || IsNil(o.Contents) {
		return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasContents() bool {
	if o != nil && !IsNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetContents(v string) {
	o.Contents = &v
}

// GetCopyrights returns the Copyrights field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCopyrights() []string {
	if o == nil || IsNil(o.Copyrights) {
		var ret []string
		return ret
	}
	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCopyrightsOk() ([]string, bool) {
	if o == nil || IsNil(o.Copyrights) {
		return nil, false
	}
	return o.Copyrights, true
}

// HasCopyrights returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasCopyrights() bool {
	if o != nil && !IsNil(o.Copyrights) {
		return true
	}

	return false
}

// SetCopyrights gets a reference to the given []string and assigns it to the Copyrights field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetCopyrights(v []string) {
	o.Copyrights = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCorrected returns the Corrected field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCorrected() bool {
	if o == nil || IsNil(o.Corrected) {
		var ret bool
		return ret
	}
	return *o.Corrected
}

// GetCorrectedOk returns a tuple with the Corrected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetCorrectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Corrected) {
		return nil, false
	}
	return o.Corrected, true
}

// HasCorrected returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasCorrected() bool {
	if o != nil && !IsNil(o.Corrected) {
		return true
	}

	return false
}

// SetCorrected gets a reference to the given bool and assigns it to the Corrected field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetCorrected(v bool) {
	o.Corrected = &v
}

// GetIgnored returns the Ignored field value if set, zero value otherwise.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetIgnored() bool {
	if o == nil || IsNil(o.Ignored) {
		var ret bool
		return ret
	}
	return *o.Ignored
}

// GetIgnoredOk returns a tuple with the Ignored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) GetIgnoredOk() (*bool, bool) {
	if o == nil || IsNil(o.Ignored) {
		return nil, false
	}
	return o.Ignored, true
}

// HasIgnored returns a boolean if a field has been set.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) HasIgnored() bool {
	if o != nil && !IsNil(o.Ignored) {
		return true
	}

	return false
}

// SetIgnored gets a reference to the given bool and assigns it to the Ignored field.
func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) SetIgnored(v bool) {
	o.Ignored = &v
}

func (o GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RevisionId) {
		toSerialize["revisionId"] = o.RevisionId
	}
	if !IsNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	if !IsNil(o.Copyrights) {
		toSerialize["copyrights"] = o.Copyrights
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Corrected) {
		toSerialize["corrected"] = o.Corrected
	}
	if !IsNil(o.Ignored) {
		toSerialize["ignored"] = o.Ignored
	}
	return toSerialize, nil
}

type NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner struct {
	value *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner
	isSet bool
}

func (v NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) Get() *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner {
	return v.value
}

func (v *NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) Set(val *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner(val *GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) *NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner {
	return &NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner{value: val, isSet: true}
}

func (v NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


