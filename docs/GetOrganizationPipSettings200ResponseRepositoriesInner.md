# GetOrganizationPipSettings200ResponseRepositoriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Pip repository (For FOSSA internal usage) | [optional] 
**Url** | Pointer to **string** | Remote URL of the Pip repository | [optional] 
**Username** | Pointer to **string** | Username for authenticating to the Pip repository | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGetOrganizationPipSettings200ResponseRepositoriesInner

`func NewGetOrganizationPipSettings200ResponseRepositoriesInner() *GetOrganizationPipSettings200ResponseRepositoriesInner`

NewGetOrganizationPipSettings200ResponseRepositoriesInner instantiates a new GetOrganizationPipSettings200ResponseRepositoriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationPipSettings200ResponseRepositoriesInnerWithDefaults

`func NewGetOrganizationPipSettings200ResponseRepositoriesInnerWithDefaults() *GetOrganizationPipSettings200ResponseRepositoriesInner`

NewGetOrganizationPipSettings200ResponseRepositoriesInnerWithDefaults instantiates a new GetOrganizationPipSettings200ResponseRepositoriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetOrganizationPipSettings200ResponseRepositoriesInner) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


