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

// checks if the GitTokenAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitTokenAuth{}

// GitTokenAuth struct for GitTokenAuth
type GitTokenAuth struct {
	// UUID of the Git Server (For FOSSA internal usage)
	Id *float32 `json:"_id,omitempty"`
	// Display name of the Git token in FOSSA
	DisplayName *string `json:"displayName,omitempty"`
	// FOSSA internal type
	Type *string `json:"type,omitempty"`
	Value *GetOrganizationBowerSettings200ResponseRegistriesInnerUrl `json:"value,omitempty"`
	// Used when an existing token is obfuscated in the response
	HasToken *bool `json:"hasToken,omitempty"`
}

// NewGitTokenAuth instantiates a new GitTokenAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitTokenAuth() *GitTokenAuth {
	this := GitTokenAuth{}
	return &this
}

// NewGitTokenAuthWithDefaults instantiates a new GitTokenAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitTokenAuthWithDefaults() *GitTokenAuth {
	this := GitTokenAuth{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitTokenAuth) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAuth) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitTokenAuth) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GitTokenAuth) SetId(v float32) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GitTokenAuth) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAuth) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GitTokenAuth) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GitTokenAuth) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GitTokenAuth) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAuth) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GitTokenAuth) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GitTokenAuth) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GitTokenAuth) GetValue() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl {
	if o == nil || IsNil(o.Value) {
		var ret GetOrganizationBowerSettings200ResponseRegistriesInnerUrl
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAuth) GetValueOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GitTokenAuth) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given GetOrganizationBowerSettings200ResponseRegistriesInnerUrl and assigns it to the Value field.
func (o *GitTokenAuth) SetValue(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl) {
	o.Value = &v
}

// GetHasToken returns the HasToken field value if set, zero value otherwise.
func (o *GitTokenAuth) GetHasToken() bool {
	if o == nil || IsNil(o.HasToken) {
		var ret bool
		return ret
	}
	return *o.HasToken
}

// GetHasTokenOk returns a tuple with the HasToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitTokenAuth) GetHasTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.HasToken) {
		return nil, false
	}
	return o.HasToken, true
}

// HasHasToken returns a boolean if a field has been set.
func (o *GitTokenAuth) HasHasToken() bool {
	if o != nil && !IsNil(o.HasToken) {
		return true
	}

	return false
}

// SetHasToken gets a reference to the given bool and assigns it to the HasToken field.
func (o *GitTokenAuth) SetHasToken(v bool) {
	o.HasToken = &v
}

func (o GitTokenAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitTokenAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.HasToken) {
		toSerialize["hasToken"] = o.HasToken
	}
	return toSerialize, nil
}

type NullableGitTokenAuth struct {
	value *GitTokenAuth
	isSet bool
}

func (v NullableGitTokenAuth) Get() *GitTokenAuth {
	return v.value
}

func (v *NullableGitTokenAuth) Set(val *GitTokenAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableGitTokenAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableGitTokenAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitTokenAuth(val *GitTokenAuth) *NullableGitTokenAuth {
	return &NullableGitTokenAuth{value: val, isSet: true}
}

func (v NullableGitTokenAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitTokenAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


