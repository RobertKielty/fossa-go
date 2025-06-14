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

// checks if the UpdateReleaseGroupByIdRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateReleaseGroupByIdRequest{}

// UpdateReleaseGroupByIdRequest struct for UpdateReleaseGroupByIdRequest
type UpdateReleaseGroupByIdRequest struct {
	Title *string `json:"title,omitempty"`
	LicensingPolicyId *int32 `json:"licensingPolicyId,omitempty"`
	SecurityPolicyId *int32 `json:"securityPolicyId,omitempty"`
	QualityPolicyId *int32 `json:"qualityPolicyId,omitempty"`
	PublicOnPortal *bool `json:"publicOnPortal,omitempty"`
	ReportCustomText *string `json:"reportCustomText,omitempty"`
}

// NewUpdateReleaseGroupByIdRequest instantiates a new UpdateReleaseGroupByIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReleaseGroupByIdRequest() *UpdateReleaseGroupByIdRequest {
	this := UpdateReleaseGroupByIdRequest{}
	return &this
}

// NewUpdateReleaseGroupByIdRequestWithDefaults instantiates a new UpdateReleaseGroupByIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReleaseGroupByIdRequestWithDefaults() *UpdateReleaseGroupByIdRequest {
	this := UpdateReleaseGroupByIdRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateReleaseGroupByIdRequest) SetTitle(v string) {
	o.Title = &v
}

// GetLicensingPolicyId returns the LicensingPolicyId field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetLicensingPolicyId() int32 {
	if o == nil || IsNil(o.LicensingPolicyId) {
		var ret int32
		return ret
	}
	return *o.LicensingPolicyId
}

// GetLicensingPolicyIdOk returns a tuple with the LicensingPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetLicensingPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LicensingPolicyId) {
		return nil, false
	}
	return o.LicensingPolicyId, true
}

// HasLicensingPolicyId returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasLicensingPolicyId() bool {
	if o != nil && !IsNil(o.LicensingPolicyId) {
		return true
	}

	return false
}

// SetLicensingPolicyId gets a reference to the given int32 and assigns it to the LicensingPolicyId field.
func (o *UpdateReleaseGroupByIdRequest) SetLicensingPolicyId(v int32) {
	o.LicensingPolicyId = &v
}

// GetSecurityPolicyId returns the SecurityPolicyId field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetSecurityPolicyId() int32 {
	if o == nil || IsNil(o.SecurityPolicyId) {
		var ret int32
		return ret
	}
	return *o.SecurityPolicyId
}

// GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetSecurityPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SecurityPolicyId) {
		return nil, false
	}
	return o.SecurityPolicyId, true
}

// HasSecurityPolicyId returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasSecurityPolicyId() bool {
	if o != nil && !IsNil(o.SecurityPolicyId) {
		return true
	}

	return false
}

// SetSecurityPolicyId gets a reference to the given int32 and assigns it to the SecurityPolicyId field.
func (o *UpdateReleaseGroupByIdRequest) SetSecurityPolicyId(v int32) {
	o.SecurityPolicyId = &v
}

// GetQualityPolicyId returns the QualityPolicyId field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetQualityPolicyId() int32 {
	if o == nil || IsNil(o.QualityPolicyId) {
		var ret int32
		return ret
	}
	return *o.QualityPolicyId
}

// GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetQualityPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QualityPolicyId) {
		return nil, false
	}
	return o.QualityPolicyId, true
}

// HasQualityPolicyId returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasQualityPolicyId() bool {
	if o != nil && !IsNil(o.QualityPolicyId) {
		return true
	}

	return false
}

// SetQualityPolicyId gets a reference to the given int32 and assigns it to the QualityPolicyId field.
func (o *UpdateReleaseGroupByIdRequest) SetQualityPolicyId(v int32) {
	o.QualityPolicyId = &v
}

// GetPublicOnPortal returns the PublicOnPortal field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetPublicOnPortal() bool {
	if o == nil || IsNil(o.PublicOnPortal) {
		var ret bool
		return ret
	}
	return *o.PublicOnPortal
}

// GetPublicOnPortalOk returns a tuple with the PublicOnPortal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetPublicOnPortalOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicOnPortal) {
		return nil, false
	}
	return o.PublicOnPortal, true
}

// HasPublicOnPortal returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasPublicOnPortal() bool {
	if o != nil && !IsNil(o.PublicOnPortal) {
		return true
	}

	return false
}

// SetPublicOnPortal gets a reference to the given bool and assigns it to the PublicOnPortal field.
func (o *UpdateReleaseGroupByIdRequest) SetPublicOnPortal(v bool) {
	o.PublicOnPortal = &v
}

// GetReportCustomText returns the ReportCustomText field value if set, zero value otherwise.
func (o *UpdateReleaseGroupByIdRequest) GetReportCustomText() string {
	if o == nil || IsNil(o.ReportCustomText) {
		var ret string
		return ret
	}
	return *o.ReportCustomText
}

// GetReportCustomTextOk returns a tuple with the ReportCustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReleaseGroupByIdRequest) GetReportCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.ReportCustomText) {
		return nil, false
	}
	return o.ReportCustomText, true
}

// HasReportCustomText returns a boolean if a field has been set.
func (o *UpdateReleaseGroupByIdRequest) HasReportCustomText() bool {
	if o != nil && !IsNil(o.ReportCustomText) {
		return true
	}

	return false
}

// SetReportCustomText gets a reference to the given string and assigns it to the ReportCustomText field.
func (o *UpdateReleaseGroupByIdRequest) SetReportCustomText(v string) {
	o.ReportCustomText = &v
}

func (o UpdateReleaseGroupByIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateReleaseGroupByIdRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.LicensingPolicyId) {
		toSerialize["licensingPolicyId"] = o.LicensingPolicyId
	}
	if !IsNil(o.SecurityPolicyId) {
		toSerialize["securityPolicyId"] = o.SecurityPolicyId
	}
	if !IsNil(o.QualityPolicyId) {
		toSerialize["qualityPolicyId"] = o.QualityPolicyId
	}
	if !IsNil(o.PublicOnPortal) {
		toSerialize["publicOnPortal"] = o.PublicOnPortal
	}
	if !IsNil(o.ReportCustomText) {
		toSerialize["reportCustomText"] = o.ReportCustomText
	}
	return toSerialize, nil
}

type NullableUpdateReleaseGroupByIdRequest struct {
	value *UpdateReleaseGroupByIdRequest
	isSet bool
}

func (v NullableUpdateReleaseGroupByIdRequest) Get() *UpdateReleaseGroupByIdRequest {
	return v.value
}

func (v *NullableUpdateReleaseGroupByIdRequest) Set(val *UpdateReleaseGroupByIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReleaseGroupByIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReleaseGroupByIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReleaseGroupByIdRequest(val *UpdateReleaseGroupByIdRequest) *NullableUpdateReleaseGroupByIdRequest {
	return &NullableUpdateReleaseGroupByIdRequest{value: val, isSet: true}
}

func (v NullableUpdateReleaseGroupByIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReleaseGroupByIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


