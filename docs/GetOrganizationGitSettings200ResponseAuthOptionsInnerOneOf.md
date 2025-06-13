# GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Git Server (For FOSSA internal usage) | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Git token in FOSSA | [optional] 
**Type** | Pointer to **string** | FOSSA internal type | [optional] 
**Value** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasToken** | Pointer to **bool** | Used when an existing token is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf`

NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOfWithDefaults

`func NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOfWithDefaults() *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf`

NewGetOrganizationGitSettings200ResponseAuthOptionsInnerOneOfWithDefaults instantiates a new GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetValue() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetValueOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) SetValue(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetHasToken() bool`

GetHasToken returns the HasToken field if non-nil, zero value otherwise.

### GetHasTokenOk

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) GetHasTokenOk() (*bool, bool)`

GetHasTokenOk returns a tuple with the HasToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) SetHasToken(v bool)`

SetHasToken sets HasToken field to given value.

### HasHasToken

`func (o *GetOrganizationGitSettings200ResponseAuthOptionsInnerOneOf) HasHasToken() bool`

HasHasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


