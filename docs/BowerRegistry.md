# BowerRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the config (For FOSSA internal usage) | [optional] 
**Url** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasUrl** | Pointer to **bool** | Used when an existing URL is obfuscated in the response | [optional] [readonly] 

## Methods

### NewBowerRegistry

`func NewBowerRegistry() *BowerRegistry`

NewBowerRegistry instantiates a new BowerRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBowerRegistryWithDefaults

`func NewBowerRegistryWithDefaults() *BowerRegistry`

NewBowerRegistryWithDefaults instantiates a new BowerRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BowerRegistry) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BowerRegistry) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BowerRegistry) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *BowerRegistry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *BowerRegistry) GetUrl() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BowerRegistry) GetUrlOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BowerRegistry) SetUrl(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BowerRegistry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHasUrl

`func (o *BowerRegistry) GetHasUrl() bool`

GetHasUrl returns the HasUrl field if non-nil, zero value otherwise.

### GetHasUrlOk

`func (o *BowerRegistry) GetHasUrlOk() (*bool, bool)`

GetHasUrlOk returns a tuple with the HasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUrl

`func (o *BowerRegistry) SetHasUrl(v bool)`

SetHasUrl sets HasUrl field to given value.

### HasHasUrl

`func (o *BowerRegistry) HasHasUrl() bool`

HasHasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


