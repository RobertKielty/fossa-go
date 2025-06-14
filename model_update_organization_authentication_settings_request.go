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

// checks if the UpdateOrganizationAuthenticationSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOrganizationAuthenticationSettingsRequest{}

// UpdateOrganizationAuthenticationSettingsRequest struct for UpdateOrganizationAuthenticationSettingsRequest
type UpdateOrganizationAuthenticationSettingsRequest struct {
	Subdomain *string `json:"subdomain,omitempty"`
	DisableInvite *bool `json:"disableInvite,omitempty"`
}

// NewUpdateOrganizationAuthenticationSettingsRequest instantiates a new UpdateOrganizationAuthenticationSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrganizationAuthenticationSettingsRequest() *UpdateOrganizationAuthenticationSettingsRequest {
	this := UpdateOrganizationAuthenticationSettingsRequest{}
	return &this
}

// NewUpdateOrganizationAuthenticationSettingsRequestWithDefaults instantiates a new UpdateOrganizationAuthenticationSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrganizationAuthenticationSettingsRequestWithDefaults() *UpdateOrganizationAuthenticationSettingsRequest {
	this := UpdateOrganizationAuthenticationSettingsRequest{}
	return &this
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *UpdateOrganizationAuthenticationSettingsRequest) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAuthenticationSettingsRequest) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *UpdateOrganizationAuthenticationSettingsRequest) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *UpdateOrganizationAuthenticationSettingsRequest) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetDisableInvite returns the DisableInvite field value if set, zero value otherwise.
func (o *UpdateOrganizationAuthenticationSettingsRequest) GetDisableInvite() bool {
	if o == nil || IsNil(o.DisableInvite) {
		var ret bool
		return ret
	}
	return *o.DisableInvite
}

// GetDisableInviteOk returns a tuple with the DisableInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrganizationAuthenticationSettingsRequest) GetDisableInviteOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableInvite) {
		return nil, false
	}
	return o.DisableInvite, true
}

// HasDisableInvite returns a boolean if a field has been set.
func (o *UpdateOrganizationAuthenticationSettingsRequest) HasDisableInvite() bool {
	if o != nil && !IsNil(o.DisableInvite) {
		return true
	}

	return false
}

// SetDisableInvite gets a reference to the given bool and assigns it to the DisableInvite field.
func (o *UpdateOrganizationAuthenticationSettingsRequest) SetDisableInvite(v bool) {
	o.DisableInvite = &v
}

func (o UpdateOrganizationAuthenticationSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOrganizationAuthenticationSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.DisableInvite) {
		toSerialize["disableInvite"] = o.DisableInvite
	}
	return toSerialize, nil
}

type NullableUpdateOrganizationAuthenticationSettingsRequest struct {
	value *UpdateOrganizationAuthenticationSettingsRequest
	isSet bool
}

func (v NullableUpdateOrganizationAuthenticationSettingsRequest) Get() *UpdateOrganizationAuthenticationSettingsRequest {
	return v.value
}

func (v *NullableUpdateOrganizationAuthenticationSettingsRequest) Set(val *UpdateOrganizationAuthenticationSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrganizationAuthenticationSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrganizationAuthenticationSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrganizationAuthenticationSettingsRequest(val *UpdateOrganizationAuthenticationSettingsRequest) *NullableUpdateOrganizationAuthenticationSettingsRequest {
	return &NullableUpdateOrganizationAuthenticationSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateOrganizationAuthenticationSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrganizationAuthenticationSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


