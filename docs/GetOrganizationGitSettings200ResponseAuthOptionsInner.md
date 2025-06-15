# GetOrganizationGitSettings200ResponseAuthOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Git Server (For FOSSA internal usage) | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Git token in FOSSA | [optional] 
**Type** | Pointer to **string** | FOSSA internal type | [optional] 
**Value** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasToken** | Pointer to **bool** | Used when an existing token is obfuscated in the response | [optional] [readonly] 
**Username** | Pointer to **string** | Username to authenticate to the remote Git server | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGetOrganizationGitSettings200ResponseAuthOptionsInner

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInner() *GetOrganizationGitSettings200ResponseAuthOptionsInner`

NewGetOrganizationGitSettings200ResponseAuthOptionsInner instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationGitSettings200ResponseAuthOptionsInnerWithDefaults

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerWithDefaults() *GetOrganizationGitSettings200ResponseAuthOptionsInner`

NewGetOrganizationGitSettings200ResponseAuthOptionsInnerWithDefaults instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetValue() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetValueOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetValue(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.

### HasHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasHasToken() bool`

HasHasToken returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInner) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


