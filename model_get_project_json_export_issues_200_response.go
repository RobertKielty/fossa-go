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

// checks if the GetProjectJSONExportIssues200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectJSONExportIssues200Response{}

// GetProjectJSONExportIssues200Response struct for GetProjectJSONExportIssues200Response
type GetProjectJSONExportIssues200Response struct {
	PackageLicenseIssues []GetProjectExportIssues200ResponsePackageLicenseIssuesInner `json:"Package License Issues,omitempty"`
	VulnerabilityIssues []GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner `json:"Vulnerability Issues,omitempty"`
	QualityIssues []GetProjectExportIssues200ResponseQualityIssuesInner `json:"Quality Issues,omitempty"`
}

// NewGetProjectJSONExportIssues200Response instantiates a new GetProjectJSONExportIssues200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectJSONExportIssues200Response() *GetProjectJSONExportIssues200Response {
	this := GetProjectJSONExportIssues200Response{}
	return &this
}

// NewGetProjectJSONExportIssues200ResponseWithDefaults instantiates a new GetProjectJSONExportIssues200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectJSONExportIssues200ResponseWithDefaults() *GetProjectJSONExportIssues200Response {
	this := GetProjectJSONExportIssues200Response{}
	return &this
}

// GetPackageLicenseIssues returns the PackageLicenseIssues field value if set, zero value otherwise.
func (o *GetProjectJSONExportIssues200Response) GetPackageLicenseIssues() []GetProjectExportIssues200ResponsePackageLicenseIssuesInner {
	if o == nil || IsNil(o.PackageLicenseIssues) {
		var ret []GetProjectExportIssues200ResponsePackageLicenseIssuesInner
		return ret
	}
	return o.PackageLicenseIssues
}

// GetPackageLicenseIssuesOk returns a tuple with the PackageLicenseIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectJSONExportIssues200Response) GetPackageLicenseIssuesOk() ([]GetProjectExportIssues200ResponsePackageLicenseIssuesInner, bool) {
	if o == nil || IsNil(o.PackageLicenseIssues) {
		return nil, false
	}
	return o.PackageLicenseIssues, true
}

// HasPackageLicenseIssues returns a boolean if a field has been set.
func (o *GetProjectJSONExportIssues200Response) HasPackageLicenseIssues() bool {
	if o != nil && !IsNil(o.PackageLicenseIssues) {
		return true
	}

	return false
}

// SetPackageLicenseIssues gets a reference to the given []GetProjectExportIssues200ResponsePackageLicenseIssuesInner and assigns it to the PackageLicenseIssues field.
func (o *GetProjectJSONExportIssues200Response) SetPackageLicenseIssues(v []GetProjectExportIssues200ResponsePackageLicenseIssuesInner) {
	o.PackageLicenseIssues = v
}

// GetVulnerabilityIssues returns the VulnerabilityIssues field value if set, zero value otherwise.
func (o *GetProjectJSONExportIssues200Response) GetVulnerabilityIssues() []GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner {
	if o == nil || IsNil(o.VulnerabilityIssues) {
		var ret []GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner
		return ret
	}
	return o.VulnerabilityIssues
}

// GetVulnerabilityIssuesOk returns a tuple with the VulnerabilityIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectJSONExportIssues200Response) GetVulnerabilityIssuesOk() ([]GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner, bool) {
	if o == nil || IsNil(o.VulnerabilityIssues) {
		return nil, false
	}
	return o.VulnerabilityIssues, true
}

// HasVulnerabilityIssues returns a boolean if a field has been set.
func (o *GetProjectJSONExportIssues200Response) HasVulnerabilityIssues() bool {
	if o != nil && !IsNil(o.VulnerabilityIssues) {
		return true
	}

	return false
}

// SetVulnerabilityIssues gets a reference to the given []GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner and assigns it to the VulnerabilityIssues field.
func (o *GetProjectJSONExportIssues200Response) SetVulnerabilityIssues(v []GetProjectJSONExportIssues200ResponseVulnerabilityIssuesInner) {
	o.VulnerabilityIssues = v
}

// GetQualityIssues returns the QualityIssues field value if set, zero value otherwise.
func (o *GetProjectJSONExportIssues200Response) GetQualityIssues() []GetProjectExportIssues200ResponseQualityIssuesInner {
	if o == nil || IsNil(o.QualityIssues) {
		var ret []GetProjectExportIssues200ResponseQualityIssuesInner
		return ret
	}
	return o.QualityIssues
}

// GetQualityIssuesOk returns a tuple with the QualityIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectJSONExportIssues200Response) GetQualityIssuesOk() ([]GetProjectExportIssues200ResponseQualityIssuesInner, bool) {
	if o == nil || IsNil(o.QualityIssues) {
		return nil, false
	}
	return o.QualityIssues, true
}

// HasQualityIssues returns a boolean if a field has been set.
func (o *GetProjectJSONExportIssues200Response) HasQualityIssues() bool {
	if o != nil && !IsNil(o.QualityIssues) {
		return true
	}

	return false
}

// SetQualityIssues gets a reference to the given []GetProjectExportIssues200ResponseQualityIssuesInner and assigns it to the QualityIssues field.
func (o *GetProjectJSONExportIssues200Response) SetQualityIssues(v []GetProjectExportIssues200ResponseQualityIssuesInner) {
	o.QualityIssues = v
}

func (o GetProjectJSONExportIssues200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectJSONExportIssues200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageLicenseIssues) {
		toSerialize["Package License Issues"] = o.PackageLicenseIssues
	}
	if !IsNil(o.VulnerabilityIssues) {
		toSerialize["Vulnerability Issues"] = o.VulnerabilityIssues
	}
	if !IsNil(o.QualityIssues) {
		toSerialize["Quality Issues"] = o.QualityIssues
	}
	return toSerialize, nil
}

type NullableGetProjectJSONExportIssues200Response struct {
	value *GetProjectJSONExportIssues200Response
	isSet bool
}

func (v NullableGetProjectJSONExportIssues200Response) Get() *GetProjectJSONExportIssues200Response {
	return v.value
}

func (v *NullableGetProjectJSONExportIssues200Response) Set(val *GetProjectJSONExportIssues200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectJSONExportIssues200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectJSONExportIssues200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectJSONExportIssues200Response(val *GetProjectJSONExportIssues200Response) *NullableGetProjectJSONExportIssues200Response {
	return &NullableGetProjectJSONExportIssues200Response{value: val, isSet: true}
}

func (v NullableGetProjectJSONExportIssues200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectJSONExportIssues200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


