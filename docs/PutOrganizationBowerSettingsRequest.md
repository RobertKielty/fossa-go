# PutOrganizationBowerSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registries** | Pointer to [**[]GetOrganizationBowerSettings200ResponseRegistriesInner**](GetOrganizationBowerSettings200ResponseRegistriesInner.md) | List of configured Bower registries | [optional] 
**UseArtifactory** | Pointer to **bool** | Does this registry use Artifactory? | [optional] 

## Methods

### NewPutOrganizationBowerSettingsRequest

`func NewPutOrganizationBowerSettingsRequest() *PutOrganizationBowerSettingsRequest`

NewPutOrganizationBowerSettingsRequest instantiates a new PutOrganizationBowerSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutOrganizationBowerSettingsRequestWithDefaults

`func NewPutOrganizationBowerSettingsRequestWithDefaults() *PutOrganizationBowerSettingsRequest`

NewPutOrganizationBowerSettingsRequestWithDefaults instantiates a new PutOrganizationBowerSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistries

`func (o *PutOrganizationBowerSettingsRequest) GetRegistries() []GetOrganizationBowerSettings200ResponseRegistriesInner`

GetRegistries returns the Registries field if non-nil, zero value otherwise.

### GetRegistriesOk

`func (o *PutOrganizationBowerSettingsRequest) GetRegistriesOk() (*[]GetOrganizationBowerSettings200ResponseRegistriesInner, bool)`

GetRegistriesOk returns a tuple with the Registries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistries

`func (o *PutOrganizationBowerSettingsRequest) SetRegistries(v []GetOrganizationBowerSettings200ResponseRegistriesInner)`

SetRegistries sets Registries field to given value.

### HasRegistries

`func (o *PutOrganizationBowerSettingsRequest) HasRegistries() bool`

HasRegistries returns a boolean if a field has been set.

### GetUseArtifactory

`func (o *PutOrganizationBowerSettingsRequest) GetUseArtifactory() bool`

GetUseArtifactory returns the UseArtifactory field if non-nil, zero value otherwise.

### GetUseArtifactoryOk

`func (o *PutOrganizationBowerSettingsRequest) GetUseArtifactoryOk() (*bool, bool)`

GetUseArtifactoryOk returns a tuple with the UseArtifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseArtifactory

`func (o *PutOrganizationBowerSettingsRequest) SetUseArtifactory(v bool)`

SetUseArtifactory sets UseArtifactory field to given value.

### HasUseArtifactory

`func (o *PutOrganizationBowerSettingsRequest) HasUseArtifactory() bool`

HasUseArtifactory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


