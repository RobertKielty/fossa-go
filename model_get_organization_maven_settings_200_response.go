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

// checks if the GetOrganizationMavenSettings200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationMavenSettings200Response{}

// GetOrganizationMavenSettings200Response struct for GetOrganizationMavenSettings200Response
type GetOrganizationMavenSettings200Response struct {
	// List of configured Maven Repositories
	Repositories []GetOrganizationMavenSettings200ResponseRepositoriesInner `json:"repositories,omitempty"`
	// List of configured Credentials for Maven Repositories
	Servers []GetOrganizationMavenSettings200ResponseServersInner `json:"servers,omitempty"`
}

// NewGetOrganizationMavenSettings200Response instantiates a new GetOrganizationMavenSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationMavenSettings200Response() *GetOrganizationMavenSettings200Response {
	this := GetOrganizationMavenSettings200Response{}
	return &this
}

// NewGetOrganizationMavenSettings200ResponseWithDefaults instantiates a new GetOrganizationMavenSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationMavenSettings200ResponseWithDefaults() *GetOrganizationMavenSettings200Response {
	this := GetOrganizationMavenSettings200Response{}
	return &this
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *GetOrganizationMavenSettings200Response) GetRepositories() []GetOrganizationMavenSettings200ResponseRepositoriesInner {
	if o == nil || IsNil(o.Repositories) {
		var ret []GetOrganizationMavenSettings200ResponseRepositoriesInner
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationMavenSettings200Response) GetRepositoriesOk() ([]GetOrganizationMavenSettings200ResponseRepositoriesInner, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *GetOrganizationMavenSettings200Response) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []GetOrganizationMavenSettings200ResponseRepositoriesInner and assigns it to the Repositories field.
func (o *GetOrganizationMavenSettings200Response) SetRepositories(v []GetOrganizationMavenSettings200ResponseRepositoriesInner) {
	o.Repositories = v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *GetOrganizationMavenSettings200Response) GetServers() []GetOrganizationMavenSettings200ResponseServersInner {
	if o == nil || IsNil(o.Servers) {
		var ret []GetOrganizationMavenSettings200ResponseServersInner
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationMavenSettings200Response) GetServersOk() ([]GetOrganizationMavenSettings200ResponseServersInner, bool) {
	if o == nil || IsNil(o.Servers) {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *GetOrganizationMavenSettings200Response) HasServers() bool {
	if o != nil && !IsNil(o.Servers) {
		return true
	}

	return false
}

// SetServers gets a reference to the given []GetOrganizationMavenSettings200ResponseServersInner and assigns it to the Servers field.
func (o *GetOrganizationMavenSettings200Response) SetServers(v []GetOrganizationMavenSettings200ResponseServersInner) {
	o.Servers = v
}

func (o GetOrganizationMavenSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationMavenSettings200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	if !IsNil(o.Servers) {
		toSerialize["servers"] = o.Servers
	}
	return toSerialize, nil
}

type NullableGetOrganizationMavenSettings200Response struct {
	value *GetOrganizationMavenSettings200Response
	isSet bool
}

func (v NullableGetOrganizationMavenSettings200Response) Get() *GetOrganizationMavenSettings200Response {
	return v.value
}

func (v *NullableGetOrganizationMavenSettings200Response) Set(val *GetOrganizationMavenSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationMavenSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationMavenSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationMavenSettings200Response(val *GetOrganizationMavenSettings200Response) *NullableGetOrganizationMavenSettings200Response {
	return &NullableGetOrganizationMavenSettings200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationMavenSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationMavenSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


