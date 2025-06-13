# PutOrganizationNugetSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]GetOrganizationNugetSettings200ResponseSourcesInner**](GetOrganizationNugetSettings200ResponseSourcesInner.md) | List of configured Nuget Feeds | [optional] 

## Methods

### NewPutOrganizationNugetSettingsRequest

`func NewPutOrganizationNugetSettingsRequest() *PutOrganizationNugetSettingsRequest`

NewPutOrganizationNugetSettingsRequest instantiates a new PutOrganizationNugetSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrganizationNugetSettingsRequestWithDefaults

`func NewPutOrganizationNugetSettingsRequestWithDefaults() *PutOrganizationNugetSettingsRequest`

NewPutOrganizationNugetSettingsRequestWithDefaults instantiates a new PutOrganizationNugetSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *PutOrganizationNugetSettingsRequest) GetSources() []GetOrganizationNugetSettings200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *PutOrganizationNugetSettingsRequest) GetSourcesOk() (*[]GetOrganizationNugetSettings200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *PutOrganizationNugetSettingsRequest) SetSources(v []GetOrganizationNugetSettings200ResponseSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *PutOrganizationNugetSettingsRequest) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


