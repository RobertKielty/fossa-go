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

// checks if the GetProjectDependencies200ResponseDependenciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectDependencies200ResponseDependenciesInner{}

// GetProjectDependencies200ResponseDependenciesInner struct for GetProjectDependencies200ResponseDependenciesInner
type GetProjectDependencies200ResponseDependenciesInner struct {
	Locator *string `json:"locator,omitempty"`
	Title *string `json:"title,omitempty"`
	IsManual *bool `json:"isManual,omitempty"`
	IsIgnored *bool `json:"isIgnored,omitempty"`
	IsUnknown *bool `json:"isUnknown,omitempty"`
	Licenses []string `json:"licenses,omitempty"`
	DeclaredLicenses []string `json:"declaredLicenses,omitempty"`
	Depth *int32 `json:"depth,omitempty"`
	OriginPaths []string `json:"originPaths,omitempty"`
	Status *GetProjectDependencies200ResponseDependenciesInnerStatus `json:"status,omitempty"`
	Issues []GetProjectDependencies200ResponseDependenciesInnerIssuesInner `json:"issues,omitempty"`
	RootProjects []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner `json:"rootProjects,omitempty"`
	LayerDepth *float32 `json:"layerDepth,omitempty"`
	Cpes []string `json:"cpes,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewGetProjectDependencies200ResponseDependenciesInner instantiates a new GetProjectDependencies200ResponseDependenciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectDependencies200ResponseDependenciesInner() *GetProjectDependencies200ResponseDependenciesInner {
	this := GetProjectDependencies200ResponseDependenciesInner{}
	return &this
}

// NewGetProjectDependencies200ResponseDependenciesInnerWithDefaults instantiates a new GetProjectDependencies200ResponseDependenciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectDependencies200ResponseDependenciesInnerWithDefaults() *GetProjectDependencies200ResponseDependenciesInner {
	this := GetProjectDependencies200ResponseDependenciesInner{}
	return &this
}

// GetLocator returns the Locator field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLocator() string {
	if o == nil || IsNil(o.Locator) {
		var ret string
		return ret
	}
	return *o.Locator
}

// GetLocatorOk returns a tuple with the Locator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLocatorOk() (*string, bool) {
	if o == nil || IsNil(o.Locator) {
		return nil, false
	}
	return o.Locator, true
}

// HasLocator returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasLocator() bool {
	if o != nil && !IsNil(o.Locator) {
		return true
	}

	return false
}

// SetLocator gets a reference to the given string and assigns it to the Locator field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetLocator(v string) {
	o.Locator = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetTitle(v string) {
	o.Title = &v
}

// GetIsManual returns the IsManual field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsManual() bool {
	if o == nil || IsNil(o.IsManual) {
		var ret bool
		return ret
	}
	return *o.IsManual
}

// GetIsManualOk returns a tuple with the IsManual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsManualOk() (*bool, bool) {
	if o == nil || IsNil(o.IsManual) {
		return nil, false
	}
	return o.IsManual, true
}

// HasIsManual returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsManual() bool {
	if o != nil && !IsNil(o.IsManual) {
		return true
	}

	return false
}

// SetIsManual gets a reference to the given bool and assigns it to the IsManual field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsManual(v bool) {
	o.IsManual = &v
}

// GetIsIgnored returns the IsIgnored field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsIgnored() bool {
	if o == nil || IsNil(o.IsIgnored) {
		var ret bool
		return ret
	}
	return *o.IsIgnored
}

// GetIsIgnoredOk returns a tuple with the IsIgnored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsIgnoredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIgnored) {
		return nil, false
	}
	return o.IsIgnored, true
}

// HasIsIgnored returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsIgnored() bool {
	if o != nil && !IsNil(o.IsIgnored) {
		return true
	}

	return false
}

// SetIsIgnored gets a reference to the given bool and assigns it to the IsIgnored field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsIgnored(v bool) {
	o.IsIgnored = &v
}

// GetIsUnknown returns the IsUnknown field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsUnknown() bool {
	if o == nil || IsNil(o.IsUnknown) {
		var ret bool
		return ret
	}
	return *o.IsUnknown
}

// GetIsUnknownOk returns a tuple with the IsUnknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsUnknownOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnknown) {
		return nil, false
	}
	return o.IsUnknown, true
}

// HasIsUnknown returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsUnknown() bool {
	if o != nil && !IsNil(o.IsUnknown) {
		return true
	}

	return false
}

// SetIsUnknown gets a reference to the given bool and assigns it to the IsUnknown field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsUnknown(v bool) {
	o.IsUnknown = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLicenses() []string {
	if o == nil || IsNil(o.Licenses) {
		var ret []string
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.Licenses) {
		return nil, false
	}
	return o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasLicenses() bool {
	if o != nil && !IsNil(o.Licenses) {
		return true
	}

	return false
}

// SetLicenses gets a reference to the given []string and assigns it to the Licenses field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetLicenses(v []string) {
	o.Licenses = v
}

// GetDeclaredLicenses returns the DeclaredLicenses field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetDeclaredLicenses() []string {
	if o == nil || IsNil(o.DeclaredLicenses) {
		var ret []string
		return ret
	}
	return o.DeclaredLicenses
}

// GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetDeclaredLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeclaredLicenses) {
		return nil, false
	}
	return o.DeclaredLicenses, true
}

// HasDeclaredLicenses returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasDeclaredLicenses() bool {
	if o != nil && !IsNil(o.DeclaredLicenses) {
		return true
	}

	return false
}

// SetDeclaredLicenses gets a reference to the given []string and assigns it to the DeclaredLicenses field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetDeclaredLicenses(v []string) {
	o.DeclaredLicenses = v
}

// GetDepth returns the Depth field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetDepth() int32 {
	if o == nil || IsNil(o.Depth) {
		var ret int32
		return ret
	}
	return *o.Depth
}

// GetDepthOk returns a tuple with the Depth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetDepthOk() (*int32, bool) {
	if o == nil || IsNil(o.Depth) {
		return nil, false
	}
	return o.Depth, true
}

// HasDepth returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasDepth() bool {
	if o != nil && !IsNil(o.Depth) {
		return true
	}

	return false
}

// SetDepth gets a reference to the given int32 and assigns it to the Depth field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetDepth(v int32) {
	o.Depth = &v
}

// GetOriginPaths returns the OriginPaths field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetOriginPaths() []string {
	if o == nil || IsNil(o.OriginPaths) {
		var ret []string
		return ret
	}
	return o.OriginPaths
}

// GetOriginPathsOk returns a tuple with the OriginPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetOriginPathsOk() ([]string, bool) {
	if o == nil || IsNil(o.OriginPaths) {
		return nil, false
	}
	return o.OriginPaths, true
}

// HasOriginPaths returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasOriginPaths() bool {
	if o != nil && !IsNil(o.OriginPaths) {
		return true
	}

	return false
}

// SetOriginPaths gets a reference to the given []string and assigns it to the OriginPaths field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetOriginPaths(v []string) {
	o.OriginPaths = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetStatus() GetProjectDependencies200ResponseDependenciesInnerStatus {
	if o == nil || IsNil(o.Status) {
		var ret GetProjectDependencies200ResponseDependenciesInnerStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetStatusOk() (*GetProjectDependencies200ResponseDependenciesInnerStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GetProjectDependencies200ResponseDependenciesInnerStatus and assigns it to the Status field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetStatus(v GetProjectDependencies200ResponseDependenciesInnerStatus) {
	o.Status = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIssues() []GetProjectDependencies200ResponseDependenciesInnerIssuesInner {
	if o == nil || IsNil(o.Issues) {
		var ret []GetProjectDependencies200ResponseDependenciesInnerIssuesInner
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetIssuesOk() ([]GetProjectDependencies200ResponseDependenciesInnerIssuesInner, bool) {
	if o == nil || IsNil(o.Issues) {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasIssues() bool {
	if o != nil && !IsNil(o.Issues) {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []GetProjectDependencies200ResponseDependenciesInnerIssuesInner and assigns it to the Issues field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetIssues(v []GetProjectDependencies200ResponseDependenciesInnerIssuesInner) {
	o.Issues = v
}

// GetRootProjects returns the RootProjects field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetRootProjects() []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner {
	if o == nil || IsNil(o.RootProjects) {
		var ret []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner
		return ret
	}
	return o.RootProjects
}

// GetRootProjectsOk returns a tuple with the RootProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetRootProjectsOk() ([]GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner, bool) {
	if o == nil || IsNil(o.RootProjects) {
		return nil, false
	}
	return o.RootProjects, true
}

// HasRootProjects returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasRootProjects() bool {
	if o != nil && !IsNil(o.RootProjects) {
		return true
	}

	return false
}

// SetRootProjects gets a reference to the given []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner and assigns it to the RootProjects field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetRootProjects(v []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner) {
	o.RootProjects = v
}

// GetLayerDepth returns the LayerDepth field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLayerDepth() float32 {
	if o == nil || IsNil(o.LayerDepth) {
		var ret float32
		return ret
	}
	return *o.LayerDepth
}

// GetLayerDepthOk returns a tuple with the LayerDepth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetLayerDepthOk() (*float32, bool) {
	if o == nil || IsNil(o.LayerDepth) {
		return nil, false
	}
	return o.LayerDepth, true
}

// HasLayerDepth returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasLayerDepth() bool {
	if o != nil && !IsNil(o.LayerDepth) {
		return true
	}

	return false
}

// SetLayerDepth gets a reference to the given float32 and assigns it to the LayerDepth field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetLayerDepth(v float32) {
	o.LayerDepth = &v
}

// GetCpes returns the Cpes field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetCpes() []string {
	if o == nil || IsNil(o.Cpes) {
		var ret []string
		return ret
	}
	return o.Cpes
}

// GetCpesOk returns a tuple with the Cpes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetCpesOk() ([]string, bool) {
	if o == nil || IsNil(o.Cpes) {
		return nil, false
	}
	return o.Cpes, true
}

// HasCpes returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasCpes() bool {
	if o != nil && !IsNil(o.Cpes) {
		return true
	}

	return false
}

// SetCpes gets a reference to the given []string and assigns it to the Cpes field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetCpes(v []string) {
	o.Cpes = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetProjectDependencies200ResponseDependenciesInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetProjectDependencies200ResponseDependenciesInner) SetVersion(v string) {
	o.Version = &v
}

func (o GetProjectDependencies200ResponseDependenciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectDependencies200ResponseDependenciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locator) {
		toSerialize["locator"] = o.Locator
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.IsManual) {
		toSerialize["isManual"] = o.IsManual
	}
	if !IsNil(o.IsIgnored) {
		toSerialize["isIgnored"] = o.IsIgnored
	}
	if !IsNil(o.IsUnknown) {
		toSerialize["isUnknown"] = o.IsUnknown
	}
	if !IsNil(o.Licenses) {
		toSerialize["licenses"] = o.Licenses
	}
	if !IsNil(o.DeclaredLicenses) {
		toSerialize["declaredLicenses"] = o.DeclaredLicenses
	}
	if !IsNil(o.Depth) {
		toSerialize["depth"] = o.Depth
	}
	if !IsNil(o.OriginPaths) {
		toSerialize["originPaths"] = o.OriginPaths
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Issues) {
		toSerialize["issues"] = o.Issues
	}
	if !IsNil(o.RootProjects) {
		toSerialize["rootProjects"] = o.RootProjects
	}
	if !IsNil(o.LayerDepth) {
		toSerialize["layerDepth"] = o.LayerDepth
	}
	if !IsNil(o.Cpes) {
		toSerialize["cpes"] = o.Cpes
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableGetProjectDependencies200ResponseDependenciesInner struct {
	value *GetProjectDependencies200ResponseDependenciesInner
	isSet bool
}

func (v NullableGetProjectDependencies200ResponseDependenciesInner) Get() *GetProjectDependencies200ResponseDependenciesInner {
	return v.value
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInner) Set(val *GetProjectDependencies200ResponseDependenciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectDependencies200ResponseDependenciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectDependencies200ResponseDependenciesInner(val *GetProjectDependencies200ResponseDependenciesInner) *NullableGetProjectDependencies200ResponseDependenciesInner {
	return &NullableGetProjectDependencies200ResponseDependenciesInner{value: val, isSet: true}
}

func (v NullableGetProjectDependencies200ResponseDependenciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectDependencies200ResponseDependenciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


