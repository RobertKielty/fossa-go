# UpdateAuthenticationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomain** | Pointer to **string** |  | [optional] 
**DisableInvite** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateAuthenticationSettings

`func NewUpdateAuthenticationSettings() *UpdateAuthenticationSettings`

NewUpdateAuthenticationSettings instantiates a new UpdateAuthenticationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAuthenticationSettingsWithDefaults

`func NewUpdateAuthenticationSettingsWithDefaults() *UpdateAuthenticationSettings`

NewUpdateAuthenticationSettingsWithDefaults instantiates a new UpdateAuthenticationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomain

`func (o *UpdateAuthenticationSettings) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UpdateAuthenticationSettings) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UpdateAuthenticationSettings) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UpdateAuthenticationSettings) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDisableInvite

`func (o *UpdateAuthenticationSettings) GetDisableInvite() bool`

GetDisableInvite returns the DisableInvite field if non-nil, zero value otherwise.

### GetDisableInviteOk

`func (o *UpdateAuthenticationSettings) GetDisableInviteOk() (*bool, bool)`

GetDisableInviteOk returns a tuple with the DisableInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInvite

`func (o *UpdateAuthenticationSettings) SetDisableInvite(v bool)`

SetDisableInvite sets DisableInvite field to given value.

### HasDisableInvite

`func (o *UpdateAuthenticationSettings) HasDisableInvite() bool`

HasDisableInvite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


