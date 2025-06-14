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

// checks if the GetIssueStatuses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssueStatuses200Response{}

// GetIssueStatuses200Response struct for GetIssueStatuses200Response
type GetIssueStatuses200Response struct {
	Issues *GetIssueStatuses200ResponseIssues `json:"issues,omitempty"`
	Revisions *GetIssueStatuses200ResponseIssues `json:"revisions,omitempty"`
}

// NewGetIssueStatuses200Response instantiates a new GetIssueStatuses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssueStatuses200Response() *GetIssueStatuses200Response {
	this := GetIssueStatuses200Response{}
	return &this
}

// NewGetIssueStatuses200ResponseWithDefaults instantiates a new GetIssueStatuses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssueStatuses200ResponseWithDefaults() *GetIssueStatuses200Response {
	this := GetIssueStatuses200Response{}
	return &this
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *GetIssueStatuses200Response) GetIssues() GetIssueStatuses200ResponseIssues {
	if o == nil || IsNil(o.Issues) {
		var ret GetIssueStatuses200ResponseIssues
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueStatuses200Response) GetIssuesOk() (*GetIssueStatuses200ResponseIssues, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *GetIssueStatuses200Response) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given GetIssueStatuses200ResponseIssues and assigns it to the Issues field.
func (o *GetIssueStatuses200Response) SetIssues(v GetIssueStatuses200ResponseIssues) {
	o.Issues = &v
}

// GetRevisions returns the Revisions field value if set, zero value otherwise.
func (o *GetIssueStatuses200Response) GetRevisions() GetIssueStatuses200ResponseIssues {
	if o == nil || IsNil(o.Revisions) {
		var ret GetIssueStatuses200ResponseIssues
		return ret
	}
	return *o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssueStatuses200Response) GetRevisionsOk() (*GetIssueStatuses200ResponseIssues, bool) {
	if o == nil || IsNil(o.Revisions) {
		return nil, false
	}
	return o.Revisions, true
}

// HasRevisions returns a boolean if a field has been set.
func (o *GetIssueStatuses200Response) HasRevisions() bool {
	if o != nil && !IsNil(o.Revisions) {
		return true
	}

	return false
}

// SetRevisions gets a reference to the given GetIssueStatuses200ResponseIssues and assigns it to the Revisions field.
func (o *GetIssueStatuses200Response) SetRevisions(v GetIssueStatuses200ResponseIssues) {
	o.Revisions = &v
}

func (o GetIssueStatuses200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssueStatuses200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.Revisions) {
		toSerialize["revisions"] = o.Revisions
	}
	return toSerialize, nil
}

type NullableGetIssueStatuses200Response struct {
	value *GetIssueStatuses200Response
	isSet bool
}

func (v NullableGetIssueStatuses200Response) Get() *GetIssueStatuses200Response {
	return v.value
}

func (v *NullableGetIssueStatuses200Response) Set(val *GetIssueStatuses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssueStatuses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssueStatuses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssueStatuses200Response(val *GetIssueStatuses200Response) *NullableGetIssueStatuses200Response {
	return &NullableGetIssueStatuses200Response{value: val, isSet: true}
}

func (v NullableGetIssueStatuses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssueStatuses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


