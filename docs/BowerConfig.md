# BowerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Registries** | Pointer to [**[]GetOrganizationBowerSettings200ResponseRegistriesInner**](GetOrganizationBowerSettings200ResponseRegistriesInner.md) | List of configured Bower registries | [optional] 
**UseArtifactory** | Pointer to **bool** | Does this registry use Artifactory? | [optional] 

## Methods

### NewBowerConfig

`func NewBowerConfig() *BowerConfig`

NewBowerConfig instantiates a new BowerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBowerConfigWithDefaults

`func NewBowerConfigWithDefaults() *BowerConfig`

NewBowerConfigWithDefaults instantiates a new BowerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistries

`func (o *BowerConfig) GetRegistries() []GetOrganizationBowerSettings200ResponseRegistriesInner`

GetRegistries returns the Registries field if non-nil, zero value otherwise.

### GetRegistriesOk

`func (o *BowerConfig) GetRegistriesOk() (*[]GetOrganizationBowerSettings200ResponseRegistriesInner, bool)`

GetRegistriesOk returns a tuple with the Registries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistries

`func (o *BowerConfig) SetRegistries(v []GetOrganizationBowerSettings200ResponseRegistriesInner)`

SetRegistries sets Registries field to given value.

### HasRegistries

`func (o *BowerConfig) HasRegistries() bool`

HasRegistries returns a boolean if a field has been set.

### GetUseArtifactory

`func (o *BowerConfig) GetUseArtifactory() bool`

GetUseArtifactory returns the UseArtifactory field if non-nil, zero value otherwise.

### GetUseArtifactoryOk

`func (o *BowerConfig) GetUseArtifactoryOk() (*bool, bool)`

GetUseArtifactoryOk returns a tuple with the UseArtifactory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseArtifactory

`func (o *BowerConfig) SetUseArtifactory(v bool)`

SetUseArtifactory sets UseArtifactory field to given value.

### HasUseArtifactory

`func (o *BowerConfig) HasUseArtifactory() bool`

HasUseArtifactory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


