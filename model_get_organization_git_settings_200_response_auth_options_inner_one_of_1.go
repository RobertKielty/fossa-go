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

// checks if the GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1{}

// GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 struct for GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1
type GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 struct {
	// UUID of the Git Server (For FOSSA internal usage)
	Id *float32 `json:"_id,omitempty"`
	// Display name of the Git token in FOSSA
	DisplayName *string `json:"displayName,omitempty"`
	// FOSSA internal type
	Type *string `json:"type,omitempty"`
	// Username to authenticate to the remote Git server
	Username *string `json:"username,omitempty"`
	Password *GetOrganizationBowerSettings200ResponseRegistriesInnerUrl `json:"password,omitempty"`
	// Used when an existing password is obfuscated in the response
	HasPassword *bool `json:"hasPassword,omitempty"`
}

// NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 {
	this := GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1{}
	return &this
}

// NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1WithDefaults instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1WithDefaults() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 {
	this := GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetId(v float32) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl {
	if o == nil || IsNil(o.Password) {
		var ret GetOrganizationBowerSettings200ResponseRegistriesInnerUrl
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given GetOrganizationBowerSettings200ResponseRegistriesInnerUrl and assigns it to the Password field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl) {
	o.Password = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetHasPassword(v bool) {
	o.HasPassword = &v
}

func (o GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.HasPassword) {
		toSerialize["hasPassword"] = o.HasPassword
	}
	return toSerialize, nil
}

type NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 struct {
	value *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1
	isSet bool
}

func (v NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) Get() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 {
	return v.value
}

func (v *NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) Set(val *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1(val *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) *NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 {
	return &NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1{value: val, isSet: true}
}

func (v NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


