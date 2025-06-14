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

// checks if the GetIssue200ResponseOneOfAllOfSource1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssue200ResponseOneOfAllOfSource1{}

// GetIssue200ResponseOneOfAllOfSource1 struct for GetIssue200ResponseOneOfAllOfSource1
type GetIssue200ResponseOneOfAllOfSource1 struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Version *string `json:"version,omitempty"`
	PackageManager *string `json:"packageManager,omitempty"`
	IsCorrected *bool `json:"isCorrected,omitempty"`
}

// NewGetIssue200ResponseOneOfAllOfSource1 instantiates a new GetIssue200ResponseOneOfAllOfSource1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssue200ResponseOneOfAllOfSource1() *GetIssue200ResponseOneOfAllOfSource1 {
	this := GetIssue200ResponseOneOfAllOfSource1{}
	return &this
}

// NewGetIssue200ResponseOneOfAllOfSource1WithDefaults instantiates a new GetIssue200ResponseOneOfAllOfSource1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssue200ResponseOneOfAllOfSource1WithDefaults() *GetIssue200ResponseOneOfAllOfSource1 {
	this := GetIssue200ResponseOneOfAllOfSource1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetUrl(v string) {
	o.Url = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetVersion(v string) {
	o.Version = &v
}

// GetPackageManager returns the PackageManager field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetPackageManager() string {
	if o == nil || IsNil(o.PackageManager) {
		var ret string
		return ret
	}
	return *o.PackageManager
}

// GetPackageManagerOk returns a tuple with the PackageManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetPackageManagerOk() (*string, bool) {
	if o == nil || IsNil(o.PackageManager) {
		return nil, false
	}
	return o.PackageManager, true
}

// HasPackageManager returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasPackageManager() bool {
	if o != nil && !IsNil(o.PackageManager) {
		return true
	}

	return false
}

// SetPackageManager gets a reference to the given string and assigns it to the PackageManager field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetPackageManager(v string) {
	o.PackageManager = &v
}

// GetIsCorrected returns the IsCorrected field value if set, zero value otherwise.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetIsCorrected() bool {
	if o == nil || IsNil(o.IsCorrected) {
		var ret bool
		return ret
	}
	return *o.IsCorrected
}

// GetIsCorrectedOk returns a tuple with the IsCorrected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) GetIsCorrectedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCorrected) {
		return nil, false
	}
	return o.IsCorrected, true
}

// HasIsCorrected returns a boolean if a field has been set.
func (o *GetIssue200ResponseOneOfAllOfSource1) HasIsCorrected() bool {
	if o != nil && !IsNil(o.IsCorrected) {
		return true
	}

	return false
}

// SetIsCorrected gets a reference to the given bool and assigns it to the IsCorrected field.
func (o *GetIssue200ResponseOneOfAllOfSource1) SetIsCorrected(v bool) {
	o.IsCorrected = &v
}

func (o GetIssue200ResponseOneOfAllOfSource1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssue200ResponseOneOfAllOfSource1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.PackageManager) {
		toSerialize["packageManager"] = o.PackageManager
	}
	if !IsNil(o.IsCorrected) {
		toSerialize["isCorrected"] = o.IsCorrected
	}
	return toSerialize, nil
}

type NullableGetIssue200ResponseOneOfAllOfSource1 struct {
	value *GetIssue200ResponseOneOfAllOfSource1
	isSet bool
}

func (v NullableGetIssue200ResponseOneOfAllOfSource1) Get() *GetIssue200ResponseOneOfAllOfSource1 {
	return v.value
}

func (v *NullableGetIssue200ResponseOneOfAllOfSource1) Set(val *GetIssue200ResponseOneOfAllOfSource1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssue200ResponseOneOfAllOfSource1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssue200ResponseOneOfAllOfSource1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssue200ResponseOneOfAllOfSource1(val *GetIssue200ResponseOneOfAllOfSource1) *NullableGetIssue200ResponseOneOfAllOfSource1 {
	return &NullableGetIssue200ResponseOneOfAllOfSource1{value: val, isSet: true}
}

func (v NullableGetIssue200ResponseOneOfAllOfSource1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssue200ResponseOneOfAllOfSource1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


