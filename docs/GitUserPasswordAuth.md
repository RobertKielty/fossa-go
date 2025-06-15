# GitUserPasswordAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Git Server (For FOSSA internal usage) | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Git token in FOSSA | [optional] 
**Type** | Pointer to **string** | FOSSA internal type | [optional] 
**Username** | Pointer to **string** | Username to authenticate to the remote Git server | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGitUserPasswordAuth

`func NewGitUserPasswordAuth() *GitUserPasswordAuth`

NewGitUserPasswordAuth instantiates a new GitUserPasswordAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitUserPasswordAuthWithDefaults

`func NewGitUserPasswordAuthWithDefaults() *GitUserPasswordAuth`

NewGitUserPasswordAuthWithDefaults instantiates a new GitUserPasswordAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitUserPasswordAuth) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitUserPasswordAuth) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitUserPasswordAuth) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GitUserPasswordAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GitUserPasswordAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GitUserPasswordAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GitUserPasswordAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GitUserPasswordAuth) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GitUserPasswordAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitUserPasswordAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitUserPasswordAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitUserPasswordAuth) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *GitUserPasswordAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GitUserPasswordAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GitUserPasswordAuth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GitUserPasswordAuth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GitUserPasswordAuth) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GitUserPasswordAuth) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GitUserPasswordAuth) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GitUserPasswordAuth) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GitUserPasswordAuth) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GitUserPasswordAuth) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GitUserPasswordAuth) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GitUserPasswordAuth) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


