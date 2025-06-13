# GitTokenAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Git Server (For FOSSA internal usage) | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Git token in FOSSA | [optional] 
**Type** | Pointer to **string** | FOSSA internal type | [optional] 
**Value** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasToken** | Pointer to **bool** | Used when an existing token is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGitTokenAuth

`func NewGitTokenAuth() *GitTokenAuth`

NewGitTokenAuth instantiates a new GitTokenAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTokenAuthWithDefaults

`func NewGitTokenAuthWithDefaults() *GitTokenAuth`

NewGitTokenAuthWithDefaults instantiates a new GitTokenAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GitTokenAuth) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GitTokenAuth) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GitTokenAuth) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GitTokenAuth) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GitTokenAuth) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GitTokenAuth) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GitTokenAuth) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GitTokenAuth) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GitTokenAuth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTokenAuth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTokenAuth) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitTokenAuth) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GitTokenAuth) GetValue() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GitTokenAuth) GetValueOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GitTokenAuth) SetValue(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetValue sets Value field to given value.

### HasValue

`func (o *GitTokenAuth) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHasToken

`func (o *GitTokenAuth) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *GitTokenAuth) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *GitTokenAuth) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.

### HasHasToken

`func (o *GitTokenAuth) HasHasToken() bool`

HasHasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


