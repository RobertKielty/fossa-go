# GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1

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

### NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1`

NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1WithDefaults

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1WithDefaults() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1`

NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1WithDefaults instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf1) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


