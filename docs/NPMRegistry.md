# NPMRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | UUID of the NPM registry (For FOSSA internal usage) | [optional] 
**Url** | Pointer to **string** | Remote URL of the NPM registry | [optional] 
**Auth** | Pointer to **map[string]interface{}** |  | [optional] 
**Ca** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasCa** | Pointer to **bool** | Used when an existing CA is obfuscated in the response | [optional] [readonly] 

## Methods

### NewNPMRegistry

`func NewNPMRegistry() *NPMRegistry`

NewNPMRegistry instantiates a new NPMRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNPMRegistryWithDefaults

`func NewNPMRegistryWithDefaults() *NPMRegistry`

NewNPMRegistryWithDefaults instantiates a new NPMRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NPMRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NPMRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NPMRegistry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NPMRegistry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NPMRegistry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NPMRegistry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NPMRegistry) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NPMRegistry) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAuth

`func (o *NPMRegistry) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *NPMRegistry) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *NPMRegistry) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *NPMRegistry) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCa

`func (o *NPMRegistry) GetCa() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetCa returns the Ca field if non-nil, zero value otherwise.

### GetCaOk

`func (o *NPMRegistry) GetCaOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetCaOk returns a tuple with the Ca field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCa

`func (o *NPMRegistry) SetCa(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetCa sets Ca field to given value.

### HasCa

`func (o *NPMRegistry) HasCa() bool`

HasCa returns a boolean if a field has been set.

### GetHasCa

`func (o *NPMRegistry) GetHasCa() bool`

GetHasCa returns the HasCa field if non-nil, zero value otherwise.

### GetHasCaOk

`func (o *NPMRegistry) GetHasCaOk() (*bool, bool)`

GetHasCaOk returns a tuple with the HasCa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCa

`func (o *NPMRegistry) SetHasCa(v bool)`

SetHasCa sets HasCa field to given value.

### HasHasCa

`func (o *NPMRegistry) HasHasCa() bool`

HasHasCa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


