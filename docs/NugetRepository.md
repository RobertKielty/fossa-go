# NugetRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Nuget Feed (For FOSSA internal usage) | [optional] 
**Url** | Pointer to **string** | Remote URL of the Nuget Feed | [optional] 
**Username** | Pointer to **string** | Username for authenticating to the Nuget Feed | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewNugetRepository

`func NewNugetRepository() *NugetRepository`

NewNugetRepository instantiates a new NugetRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNugetRepositoryWithDefaults

`func NewNugetRepositoryWithDefaults() *NugetRepository`

NewNugetRepositoryWithDefaults instantiates a new NugetRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NugetRepository) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NugetRepository) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NugetRepository) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *NugetRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *NugetRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NugetRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NugetRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NugetRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *NugetRepository) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NugetRepository) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NugetRepository) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NugetRepository) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NugetRepository) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NugetRepository) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NugetRepository) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NugetRepository) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *NugetRepository) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *NugetRepository) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *NugetRepository) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *NugetRepository) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


