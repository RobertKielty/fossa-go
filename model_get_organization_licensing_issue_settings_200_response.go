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

// checks if the GetOrganizationLicensingIssueSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationLicensingIssueSettings200Response{}

// GetOrganizationLicensingIssueSettings200Response struct for GetOrganizationLicensingIssueSettings200Response
type GetOrganizationLicensingIssueSettings200Response struct {
	DefaultPolicyId *int32 `json:"defaultPolicyId,omitempty"`
	ProjectDefaultLicensingIssueScanningEnabled *bool `json:"projectDefaultLicensingIssueScanningEnabled,omitempty"`
	ProjectDefaultLicensingStatusCheckEnabled *bool `json:"projectDefaultLicensingStatusCheckEnabled,omitempty"`
	ProjectDefaultStatusCheckFilterLicensing *int32 `json:"projectDefaultStatusCheckFilterLicensing,omitempty"`
}

// NewGetOrganizationLicensingIssueSettings200Response instantiates a new GetOrganizationLicensingIssueSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationLicensingIssueSettings200Response() *GetOrganizationLicensingIssueSettings200Response {
	this := GetOrganizationLicensingIssueSettings200Response{}
	return &this
}

// NewGetOrganizationLicensingIssueSettings200ResponseWithDefaults instantiates a new GetOrganizationLicensingIssueSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationLicensingIssueSettings200ResponseWithDefaults() *GetOrganizationLicensingIssueSettings200Response {
	this := GetOrganizationLicensingIssueSettings200Response{}
	return &this
}

// GetDefaultPolicyId returns the DefaultPolicyId field value if set, zero value otherwise.
func (o *GetOrganizationLicensingIssueSettings200Response) GetDefaultPolicyId() int32 {
	if o == nil || IsNil(o.DefaultPolicyId) {
		var ret int32
		return ret
	}
	return *o.DefaultPolicyId
}

// GetDefaultPolicyIdOk returns a tuple with the DefaultPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) GetDefaultPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultPolicyId) {
		return nil, false
	}
	return o.DefaultPolicyId, true
}

// HasDefaultPolicyId returns a boolean if a field has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) HasDefaultPolicyId() bool {
	if o != nil && !IsNil(o.DefaultPolicyId) {
		return true
	}

	return false
}

// SetDefaultPolicyId gets a reference to the given int32 and assigns it to the DefaultPolicyId field.
func (o *GetOrganizationLicensingIssueSettings200Response) SetDefaultPolicyId(v int32) {
	o.DefaultPolicyId = &v
}

// GetProjectDefaultLicensingIssueScanningEnabled returns the ProjectDefaultLicensingIssueScanningEnabled field value if set, zero value otherwise.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultLicensingIssueScanningEnabled() bool {
	if o == nil || IsNil(o.ProjectDefaultLicensingIssueScanningEnabled) {
		var ret bool
		return ret
	}
	return *o.ProjectDefaultLicensingIssueScanningEnabled
}

// GetProjectDefaultLicensingIssueScanningEnabledOk returns a tuple with the ProjectDefaultLicensingIssueScanningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultLicensingIssueScanningEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProjectDefaultLicensingIssueScanningEnabled) {
		return nil, false
	}
	return o.ProjectDefaultLicensingIssueScanningEnabled, true
}

// HasProjectDefaultLicensingIssueScanningEnabled returns a boolean if a field has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) HasProjectDefaultLicensingIssueScanningEnabled() bool {
	if o != nil && !IsNil(o.ProjectDefaultLicensingIssueScanningEnabled) {
		return true
	}

	return false
}

// SetProjectDefaultLicensingIssueScanningEnabled gets a reference to the given bool and assigns it to the ProjectDefaultLicensingIssueScanningEnabled field.
func (o *GetOrganizationLicensingIssueSettings200Response) SetProjectDefaultLicensingIssueScanningEnabled(v bool) {
	o.ProjectDefaultLicensingIssueScanningEnabled = &v
}

// GetProjectDefaultLicensingStatusCheckEnabled returns the ProjectDefaultLicensingStatusCheckEnabled field value if set, zero value otherwise.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultLicensingStatusCheckEnabled() bool {
	if o == nil || IsNil(o.ProjectDefaultLicensingStatusCheckEnabled) {
		var ret bool
		return ret
	}
	return *o.ProjectDefaultLicensingStatusCheckEnabled
}

// GetProjectDefaultLicensingStatusCheckEnabledOk returns a tuple with the ProjectDefaultLicensingStatusCheckEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultLicensingStatusCheckEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProjectDefaultLicensingStatusCheckEnabled) {
		return nil, false
	}
	return o.ProjectDefaultLicensingStatusCheckEnabled, true
}

// HasProjectDefaultLicensingStatusCheckEnabled returns a boolean if a field has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) HasProjectDefaultLicensingStatusCheckEnabled() bool {
	if o != nil && !IsNil(o.ProjectDefaultLicensingStatusCheckEnabled) {
		return true
	}

	return false
}

// SetProjectDefaultLicensingStatusCheckEnabled gets a reference to the given bool and assigns it to the ProjectDefaultLicensingStatusCheckEnabled field.
func (o *GetOrganizationLicensingIssueSettings200Response) SetProjectDefaultLicensingStatusCheckEnabled(v bool) {
	o.ProjectDefaultLicensingStatusCheckEnabled = &v
}

// GetProjectDefaultStatusCheckFilterLicensing returns the ProjectDefaultStatusCheckFilterLicensing field value if set, zero value otherwise.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultStatusCheckFilterLicensing() int32 {
	if o == nil || IsNil(o.ProjectDefaultStatusCheckFilterLicensing) {
		var ret int32
		return ret
	}
	return *o.ProjectDefaultStatusCheckFilterLicensing
}

// GetProjectDefaultStatusCheckFilterLicensingOk returns a tuple with the ProjectDefaultStatusCheckFilterLicensing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) GetProjectDefaultStatusCheckFilterLicensingOk() (*int32, bool) {
	if o == nil || IsNil(o.ProjectDefaultStatusCheckFilterLicensing) {
		return nil, false
	}
	return o.ProjectDefaultStatusCheckFilterLicensing, true
}

// HasProjectDefaultStatusCheckFilterLicensing returns a boolean if a field has been set.
func (o *GetOrganizationLicensingIssueSettings200Response) HasProjectDefaultStatusCheckFilterLicensing() bool {
	if o != nil && !IsNil(o.ProjectDefaultStatusCheckFilterLicensing) {
		return true
	}

	return false
}

// SetProjectDefaultStatusCheckFilterLicensing gets a reference to the given int32 and assigns it to the ProjectDefaultStatusCheckFilterLicensing field.
func (o *GetOrganizationLicensingIssueSettings200Response) SetProjectDefaultStatusCheckFilterLicensing(v int32) {
	o.ProjectDefaultStatusCheckFilterLicensing = &v
}

func (o GetOrganizationLicensingIssueSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationLicensingIssueSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultPolicyId) {
		toSerialize["defaultPolicyId"] = o.DefaultPolicyId
	}
	if !IsNil(o.ProjectDefaultLicensingIssueScanningEnabled) {
		toSerialize["projectDefaultLicensingIssueScanningEnabled"] = o.ProjectDefaultLicensingIssueScanningEnabled
	}
	if !IsNil(o.ProjectDefaultLicensingStatusCheckEnabled) {
		toSerialize["projectDefaultLicensingStatusCheckEnabled"] = o.ProjectDefaultLicensingStatusCheckEnabled
	}
	if !IsNil(o.ProjectDefaultStatusCheckFilterLicensing) {
		toSerialize["projectDefaultStatusCheckFilterLicensing"] = o.ProjectDefaultStatusCheckFilterLicensing
	}
	return toSerialize, nil
}

type NullableGetOrganizationLicensingIssueSettings200Response struct {
	value *GetOrganizationLicensingIssueSettings200Response
	isSet bool
}

func (v NullableGetOrganizationLicensingIssueSettings200Response) Get() *GetOrganizationLicensingIssueSettings200Response {
	return v.value
}

func (v *NullableGetOrganizationLicensingIssueSettings200Response) Set(val *GetOrganizationLicensingIssueSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationLicensingIssueSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationLicensingIssueSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationLicensingIssueSettings200Response(val *GetOrganizationLicensingIssueSettings200Response) *NullableGetOrganizationLicensingIssueSettings200Response {
	return &NullableGetOrganizationLicensingIssueSettings200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationLicensingIssueSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationLicensingIssueSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


