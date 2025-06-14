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
	"time"
	"bytes"
	"fmt"
)

// checks if the GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse{}

// GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse struct for GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse
type GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse struct {
	// The ID of the package label assignment.
	Id int32 `json:"id"`
	// The date and time the package label assignment was created.
	CreatedAt time.Time `json:"createdAt"`
	// The date and time the package label assignment was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// The ID of the organization that owns the package label assignment.
	OrganizationId int32 `json:"organizationId"`
	// The ID of the label that was assigned to the package.
	LabelId int32 `json:"labelId"`
	// The ID of the package that the label was assigned to.
	PackageId string `json:"packageId"`
	// The version of the package that the label was assigned to or null if the label was assigned to all versions.
	PackageVersion *string `json:"packageVersion,omitempty"`
	// The scope of the package label assignment.
	Scope string `json:"scope"`
	// The ID of the scope that the label was assigned to or null if the label was assigned to all scopes.
	ScopeId string `json:"scopeId"`
}

type _GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse

// NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse instantiates a new GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse(id int32, createdAt time.Time, updatedAt time.Time, organizationId int32, labelId int32, packageId string, scope string, scopeId string) *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse {
	this := GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.OrganizationId = organizationId
	this.LabelId = labelId
	this.PackageId = packageId
	this.Scope = scope
	this.ScopeId = scopeId
	return &this
}

// NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponseWithDefaults instantiates a new GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponseWithDefaults() *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse {
	this := GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse{}
	return &this
}

// GetId returns the Id field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetId(v int32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetOrganizationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetOrganizationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetOrganizationId(v int32) {
	o.OrganizationId = v
}

// GetLabelId returns the LabelId field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetLabelId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetLabelIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelId, true
}

// SetLabelId sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetLabelId(v int32) {
	o.LabelId = v
}

// GetPackageId returns the PackageId field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetPackageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageId, true
}

// SetPackageId sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetPackageId(v string) {
	o.PackageId = v
}

// GetPackageVersion returns the PackageVersion field value if set, zero value otherwise.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetPackageVersion() string {
	if o == nil || IsNil(o.PackageVersion) {
		var ret string
		return ret
	}
	return *o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetPackageVersionOk() (*string, bool) {
	if o == nil || IsNil(o.PackageVersion) {
		return nil, false
	}
	return o.PackageVersion, true
}

// HasPackageVersion returns a boolean if a field has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) HasPackageVersion() bool {
	if o != nil && !IsNil(o.PackageVersion) {
		return true
	}

	return false
}

// SetPackageVersion gets a reference to the given string and assigns it to the PackageVersion field.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetPackageVersion(v string) {
	o.PackageVersion = &v
}

// GetScope returns the Scope field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetScope(v string) {
	o.Scope = v
}

// GetScopeId returns the ScopeId field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetScopeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value
// and a boolean to check if the value has been set.
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) GetScopeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScopeId, true
}

// SetScopeId sets field value
func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) SetScopeId(v string) {
	o.ScopeId = v
}

func (o GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["labelId"] = o.LabelId
	toSerialize["packageId"] = o.PackageId
	if !IsNil(o.PackageVersion) {
		toSerialize["packageVersion"] = o.PackageVersion
	}
	toSerialize["scope"] = o.Scope
	toSerialize["scopeId"] = o.ScopeId
	return toSerialize, nil
}

func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdAt",
		"updatedAt",
		"organizationId",
		"labelId",
		"packageId",
		"scope",
		"scopeId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse := _GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse)

	if err != nil {
		return err
	}

	*o = GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse(varGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse)

	return err
}

type NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse struct {
	value *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse
	isSet bool
}

func (v NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) Get() *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse {
	return v.value
}

func (v *NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) Set(val *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse(val *GetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) *NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse {
	return &NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse{value: val, isSet: true}
}

func (v NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPackageLabelAssignments200ResponsePackageLabelAssignmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


