# GetOrganizationMavenSettings200ResponseServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Maven Server (For FOSSA internal usage) | [optional] 
**Id** | Pointer to **string** | User configured ID (corresponds with a Maven Repository) | [optional] 
**Username** | Pointer to **string** | Username for authenticating to the Maven repository | [optional] 
**Password** | Pointer to [**GetOrganizationBowerSettings200ResponseRegistriesInnerUrl**](GetOrganizationBowerSettings200ResponseRegistriesInnerUrl.md) |  | [optional] 
**HasPassword** | Pointer to **bool** | Used when an existing password is obfuscated in the response | [optional] [readonly] 

## Methods

### NewGetOrganizationMavenSettings200ResponseServersInner

`func NewGetOrganizationMavenSettings200ResponseServersInner() *GetOrganizationMavenSettings200ResponseServersInner`

NewGetOrganizationMavenSettings200ResponseServersInner instantiates a new GetOrganizationMavenSettings200ResponseServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationMavenSettings200ResponseServersInnerWithDefaults

`func NewGetOrganizationMavenSettings200ResponseServersInnerWithDefaults() *GetOrganizationMavenSettings200ResponseServersInner`

NewGetOrganizationMavenSettings200ResponseServersInnerWithDefaults instantiates a new GetOrganizationMavenSettings200ResponseServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationMavenSettings200ResponseServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationMavenSettings200ResponseServersInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationMavenSettings200ResponseServersInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetPassword() GetOrganizationBowerSettings200ResponseRegistriesInnerUrl`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetPasswordOk() (*GetOrganizationBowerSettings200ResponseRegistriesInnerUrl, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) SetPassword(v GetOrganizationBowerSettings200ResponseRegistriesInnerUrl)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetHasPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *GetOrganizationMavenSettings200ResponseServersInner) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *GetOrganizationMavenSettings200ResponseServersInner) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


