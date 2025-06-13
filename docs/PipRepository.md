# PipRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Pip repository (For FOSSA internal usage) | [optional] 
**Url** | Pointer to **string** | Remote URL of the Pip repository | [optional] 
**Username** | Pointer to **string** | Username for authenticating to the Pip repository | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewPipRepository

`func NewPipRepository() *PipRepository`

NewPipRepository instantiates a new PipRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipRepositoryWithDefaults

`func NewPipRepositoryWithDefaults() *PipRepository`

NewPipRepositoryWithDefaults instantiates a new PipRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipRepository) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipRepository) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipRepository) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PipRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *PipRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PipRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PipRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PipRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *PipRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PipRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PipRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PipRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *PipRepository) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PipRepository) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PipRepository) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PipRepository) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *PipRepository) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *PipRepository) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *PipRepository) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *PipRepository) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


