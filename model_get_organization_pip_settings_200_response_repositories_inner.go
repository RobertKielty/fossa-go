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

// checks if the GetOrganizationPipSettings200ResponseRepositoriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationPipSettings200ResponseRepositoriesInner{}

// GetOrganizationPipSettings200ResponseRepositoriesInner struct for GetOrganizationPipSettings200ResponseRepositoriesInner
type GetOrganizationPipSettings200ResponseRepositoriesInner struct {
	// UUID of the Pip repository (For FOSSA internal usage)
	Id *float32 `json:"_id,omitempty"`
	// Remote URL of the Pip repository
	Url *string `json:"url,omitempty"`
	// Username for authenticating to the Pip repository
	Username *string `json:"username,omitempty"`
	Password *GetOrganizationBowerSettings200ResponseRegistriesInnerUrl `json:"password,omitempty"`
	// Used when an existing password is obfuscated in the response
	HasPassword *bool `json:"hasPassword,omitempty"`
}

// NewGetOrganizationPipSettings200ResponseRepositoriesInner instantiates a new GetOrganizationPipSettings200ResponseRepositoriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationPipSettings200ResponseRepositoriesInner() *GetOrganizationPipSettings200ResponseRepositoriesInner {
	this := GetOrganizationPipSettings200ResponseRepositoriesInner{}
	return &this
}

// NewGetOrganizationPipSettings200ResponseRepositoriesInnerWithDefaults instantiates a new GetOrganizationPipSettings200ResponseRepositoriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationPipSettings200ResponseRepositoriesInnerWithDefaults() *GetOrganizationPipSettings200ResponseRepositoriesInner {
	this := GetOrganizationPipSettings200ResponseRepositoriesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetId(v float32) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetUrl(v string) {
	o.Url = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl {
	if o == nil || IsNil(o.Password) {
		var ret GetOrganizationBowerSettings200ResponseRegistriesInnerUrl
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given GetOrganizationBowerSettings200ResponseRegistriesInnerUrl and assigns it to the Password field.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl) {
	o.Password = &v
}

// GetHasPassword returns the HasPassword field value if set, zero value otherwise.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetHasPassword() bool {
	if o == nil || IsNil(o.HasPassword) {
		var ret bool
		return ret
	}
	return *o.HasPassword
}

// GetHasPasswordOk returns a tuple with the HasPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetHasPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.HasPassword) {
		return nil, false
	}
	return o.HasPassword, true
}

// HasHasPassword returns a boolean if a field has been set.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasHasPassword() bool {
	if o != nil && !IsNil(o.HasPassword) {
		return true
	}

	return false
}

// SetHasPassword gets a reference to the given bool and assigns it to the HasPassword field.
func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetHasPassword(v bool) {
	o.HasPassword = &v
}

func (o GetOrganizationPipSettings200ResponseRepositoriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationPipSettings200ResponseRepositoriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
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

type NullableGetOrganizationPipSettings200ResponseRepositoriesInner struct {
	value *GetOrganizationPipSettings200ResponseRepositoriesInner
	isSet bool
}

func (v NullableGetOrganizationPipSettings200ResponseRepositoriesInner) Get() *GetOrganizationPipSettings200ResponseRepositoriesInner {
	return v.value
}

func (v *NullableGetOrganizationPipSettings200ResponseRepositoriesInner) Set(val *GetOrganizationPipSettings200ResponseRepositoriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationPipSettings200ResponseRepositoriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationPipSettings200ResponseRepositoriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationPipSettings200ResponseRepositoriesInner(val *GetOrganizationPipSettings200ResponseRepositoriesInner) *NullableGetOrganizationPipSettings200ResponseRepositoriesInner {
	return &NullableGetOrganizationPipSettings200ResponseRepositoriesInner{value: val, isSet: true}
}

func (v NullableGetOrganizationPipSettings200ResponseRepositoriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationPipSettings200ResponseRepositoriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


