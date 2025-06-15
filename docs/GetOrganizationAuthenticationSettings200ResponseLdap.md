# GetOrganizationAuthenticationSettings200ResponseLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PasswordIsSet** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**OrganizationId** | Pointer to **float32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **float32** |  | [optional] 
**Ssl** | Pointer to **bool** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**SearchBase** | Pointer to **string** |  | [optional] 
**SearchFilter** | Pointer to **string** |  | [optional] 
**UseGroups** | Pointer to **bool** |  | [optional] 
**GroupBase** | Pointer to **string** |  | [optional] 
**GroupFilter** | Pointer to **string** |  | [optional] 
**GroupNameAttr** | Pointer to **string** |  | [optional] 
**GroupDNAttr** | Pointer to **string** |  | [optional] 
**UserEmailAttr** | Pointer to **string** |  | [optional] 
**UserNameAttr** | Pointer to **string** |  | [optional] 
**UidAttr** | Pointer to **string** |  | [optional] 
**AllowEmptyEmail** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetOrganizationAuthenticationSettings200ResponseLdap

`func NewGetOrganizationAuthenticationSettings200ResponseLdap() *GetOrganizationAuthenticationSettings200ResponseLdap`

NewGetOrganizationAuthenticationSettings200ResponseLdap instantiates a new GetOrganizationAuthenticationSettings200ResponseLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAuthenticationSettings200ResponseLdapWithDefaults

`func NewGetOrganizationAuthenticationSettings200ResponseLdapWithDefaults() *GetOrganizationAuthenticationSettings200ResponseLdap`

NewGetOrganizationAuthenticationSettings200ResponseLdapWithDefaults instantiates a new GetOrganizationAuthenticationSettings200ResponseLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasswordIsSet

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetPasswordIsSet() bool`

GetPasswordIsSet returns the PasswordIsSet field if non-nil, zero value otherwise.

### GetPasswordIsSetOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetPasswordIsSetOk() (*bool, bool)`

GetPasswordIsSetOk returns a tuple with the PasswordIsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordIsSet

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetPasswordIsSet(v bool)`

SetPasswordIsSet sets PasswordIsSet field to given value.

### HasPasswordIsSet

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasPasswordIsSet() bool`

HasPasswordIsSet returns a boolean if a field has been set.

### GetId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetOrganizationId() float32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetOrganizationIdOk() (*float32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetOrganizationId(v float32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEnabled

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServer

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetPort

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSsl

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetSsl(v bool)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetUsername

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSearchBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSearchBase() string`

GetSearchBase returns the SearchBase field if non-nil, zero value otherwise.

### GetSearchBaseOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSearchBaseOk() (*string, bool)`

GetSearchBaseOk returns a tuple with the SearchBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetSearchBase(v string)`

SetSearchBase sets SearchBase field to given value.

### HasSearchBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasSearchBase() bool`

HasSearchBase returns a boolean if a field has been set.

### GetSearchFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSearchFilter() string`

GetSearchFilter returns the SearchFilter field if non-nil, zero value otherwise.

### GetSearchFilterOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetSearchFilterOk() (*string, bool)`

GetSearchFilterOk returns a tuple with the SearchFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetSearchFilter(v string)`

SetSearchFilter sets SearchFilter field to given value.

### HasSearchFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasSearchFilter() bool`

HasSearchFilter returns a boolean if a field has been set.

### GetUseGroups

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUseGroups() bool`

GetUseGroups returns the UseGroups field if non-nil, zero value otherwise.

### GetUseGroupsOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUseGroupsOk() (*bool, bool)`

GetUseGroupsOk returns a tuple with the UseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGroups

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUseGroups(v bool)`

SetUseGroups sets UseGroups field to given value.

### HasUseGroups

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUseGroups() bool`

HasUseGroups returns a boolean if a field has been set.

### GetGroupBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupBase() string`

GetGroupBase returns the GroupBase field if non-nil, zero value otherwise.

### GetGroupBaseOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupBaseOk() (*string, bool)`

GetGroupBaseOk returns a tuple with the GroupBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetGroupBase(v string)`

SetGroupBase sets GroupBase field to given value.

### HasGroupBase

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasGroupBase() bool`

HasGroupBase returns a boolean if a field has been set.

### GetGroupFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupFilter() string`

GetGroupFilter returns the GroupFilter field if non-nil, zero value otherwise.

### GetGroupFilterOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupFilterOk() (*string, bool)`

GetGroupFilterOk returns a tuple with the GroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetGroupFilter(v string)`

SetGroupFilter sets GroupFilter field to given value.

### HasGroupFilter

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasGroupFilter() bool`

HasGroupFilter returns a boolean if a field has been set.

### GetGroupNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupNameAttr() string`

GetGroupNameAttr returns the GroupNameAttr field if non-nil, zero value otherwise.

### GetGroupNameAttrOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupNameAttrOk() (*string, bool)`

GetGroupNameAttrOk returns a tuple with the GroupNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetGroupNameAttr(v string)`

SetGroupNameAttr sets GroupNameAttr field to given value.

### HasGroupNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasGroupNameAttr() bool`

HasGroupNameAttr returns a boolean if a field has been set.

### GetGroupDNAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupDNAttr() string`

GetGroupDNAttr returns the GroupDNAttr field if non-nil, zero value otherwise.

### GetGroupDNAttrOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetGroupDNAttrOk() (*string, bool)`

GetGroupDNAttrOk returns a tuple with the GroupDNAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDNAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetGroupDNAttr(v string)`

SetGroupDNAttr sets GroupDNAttr field to given value.

### HasGroupDNAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasGroupDNAttr() bool`

HasGroupDNAttr returns a boolean if a field has been set.

### GetUserEmailAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUserEmailAttr() string`

GetUserEmailAttr returns the UserEmailAttr field if non-nil, zero value otherwise.

### GetUserEmailAttrOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUserEmailAttrOk() (*string, bool)`

GetUserEmailAttrOk returns a tuple with the UserEmailAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmailAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUserEmailAttr(v string)`

SetUserEmailAttr sets UserEmailAttr field to given value.

### HasUserEmailAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUserEmailAttr() bool`

HasUserEmailAttr returns a boolean if a field has been set.

### GetUserNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUserNameAttr() string`

GetUserNameAttr returns the UserNameAttr field if non-nil, zero value otherwise.

### GetUserNameAttrOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUserNameAttrOk() (*string, bool)`

GetUserNameAttrOk returns a tuple with the UserNameAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUserNameAttr(v string)`

SetUserNameAttr sets UserNameAttr field to given value.

### HasUserNameAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUserNameAttr() bool`

HasUserNameAttr returns a boolean if a field has been set.

### GetUidAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUidAttr() string`

GetUidAttr returns the UidAttr field if non-nil, zero value otherwise.

### GetUidAttrOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUidAttrOk() (*string, bool)`

GetUidAttrOk returns a tuple with the UidAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUidAttr(v string)`

SetUidAttr sets UidAttr field to given value.

### HasUidAttr

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUidAttr() bool`

HasUidAttr returns a boolean if a field has been set.

### GetAllowEmptyEmail

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetAllowEmptyEmail() bool`

GetAllowEmptyEmail returns the AllowEmptyEmail field if non-nil, zero value otherwise.

### GetAllowEmptyEmailOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetAllowEmptyEmailOk() (*bool, bool)`

GetAllowEmptyEmailOk returns a tuple with the AllowEmptyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmptyEmail

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetAllowEmptyEmail(v bool)`

SetAllowEmptyEmail sets AllowEmptyEmail field to given value.

### HasAllowEmptyEmail

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasAllowEmptyEmail() bool`

HasAllowEmptyEmail returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetOrganizationAuthenticationSettings200ResponseLdap) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


