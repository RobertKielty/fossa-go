# GitAuth

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

### NewGitAuth

`func NewGitAuth() *GitAuth`

NewGitAuth instantiates a new GitAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitAuthWithDefaults

`func NewGitAuthWithDefaults() *GitAuth`

NewGitAuthWithDefaults instantiates a new GitAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitAuth) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitAuth) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitAuth) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GitAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GitAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GitAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GitAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GitAuth) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GitAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitAuth) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GitAuth) GetValue() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GitAuth) GetValueOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GitAuth) SetValue(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetValue sets Value field to given value.

### HasValue

`func (o *GitAuth) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHasToken

`func (o *GitAuth) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *GitAuth) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *GitAuth) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.

### HasHasToken

`func (o *GitAuth) HasHasToken() bool`

HasHasToken returns a boolean if a field has been set.

### GetUsername

`func (o *GitAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitAuth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GitAuth) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GitAuth) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GitAuth) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GitAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GitAuth) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GitAuth) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GitAuth) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GitAuth) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


