# GemServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the gem server (For FOSSA internal usage) | [optional] 
**Url** | Pointer to **string** | URL of the external gem server | [optional] 
**Username** | Pointer to **string** | Username for authenticating to the gem server | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGemServer

`func NewGemServer() *GemServer`

NewGemServer instantiates a new GemServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGemServerWithDefaults

`func NewGemServerWithDefaults() *GemServer`

NewGemServerWithDefaults instantiates a new GemServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GemServer) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GemServer) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GemServer) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GemServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *GemServer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GemServer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GemServer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GemServer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *GemServer) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GemServer) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GemServer) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GemServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GemServer) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GemServer) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GemServer) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GemServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GemServer) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GemServer) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GemServer) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GemServer) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


