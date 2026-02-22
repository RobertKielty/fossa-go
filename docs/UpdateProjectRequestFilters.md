# UpdateProjectRequestFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Licensing** | Pointer to **int32** | Saved filter ID for licensing issues | [optional] 
**Vulnerability** | Pointer to **int32** | Saved filter ID for security/vulnerability issues | [optional] 
**Quality** | Pointer to **int32** | Saved filter ID for quality issues | [optional] 

## Methods

### NewUpdateProjectRequestFilters

`func NewUpdateProjectRequestFilters() *UpdateProjectRequestFilters`

NewUpdateProjectRequestFilters instantiates a new UpdateProjectRequestFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestFiltersWithDefaults

`func NewUpdateProjectRequestFiltersWithDefaults() *UpdateProjectRequestFilters`

NewUpdateProjectRequestFiltersWithDefaults instantiates a new UpdateProjectRequestFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicensing

`func (o *UpdateProjectRequestFilters) GetLicensing() int32`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *UpdateProjectRequestFilters) GetLicensingOk() (*int32, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *UpdateProjectRequestFilters) SetLicensing(v int32)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *UpdateProjectRequestFilters) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetVulnerability

`func (o *UpdateProjectRequestFilters) GetVulnerability() int32`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *UpdateProjectRequestFilters) GetVulnerabilityOk() (*int32, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *UpdateProjectRequestFilters) SetVulnerability(v int32)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *UpdateProjectRequestFilters) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.

### GetQuality

`func (o *UpdateProjectRequestFilters) GetQuality() int32`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *UpdateProjectRequestFilters) GetQualityOk() (*int32, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *UpdateProjectRequestFilters) SetQuality(v int32)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *UpdateProjectRequestFilters) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


