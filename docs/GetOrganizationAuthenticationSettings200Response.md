# GetOrganizationAuthenticationSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subdomain** | Pointer to **string** |  | [optional] 
**LoginSubdomain** | Pointer to **string** |  | [optional] 
**DisableInvite** | Pointer to **bool** |  | [optional] 
**Saml** | Pointer to [**GetOrganizationAuthenticationSettings200ResponseSaml**](GetOrganizationAuthenticationSettings200ResponseSaml.md) |  | [optional] 
**Ldap** | Pointer to [**GetOrganizationAuthenticationSettings200ResponseLdap**](GetOrganizationAuthenticationSettings200ResponseLdap.md) |  | [optional] 
**Sso** | Pointer to [**GetOrganizationAuthenticationSettings200ResponseSso**](GetOrganizationAuthenticationSettings200ResponseSso.md) |  | [optional] 

## Methods

### NewGetOrganizationAuthenticationSettings200Response

`func NewGetOrganizationAuthenticationSettings200Response() *GetOrganizationAuthenticationSettings200Response`

NewGetOrganizationAuthenticationSettings200Response instantiates a new GetOrganizationAuthenticationSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAuthenticationSettings200ResponseWithDefaults

`func NewGetOrganizationAuthenticationSettings200ResponseWithDefaults() *GetOrganizationAuthenticationSettings200Response`

NewGetOrganizationAuthenticationSettings200ResponseWithDefaults instantiates a new GetOrganizationAuthenticationSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetLoginSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) GetLoginSubdomain() string`

GetLoginSubdomain returns the LoginSubdomain field if non-nil, zero value otherwise.

### GetLoginSubdomainOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetLoginSubdomainOk() (*string, bool)`

GetLoginSubdomainOk returns a tuple with the LoginSubdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) SetLoginSubdomain(v string)`

SetLoginSubdomain sets LoginSubdomain field to given value.

### HasLoginSubdomain

`func (o *GetOrganizationAuthenticationSettings200Response) HasLoginSubdomain() bool`

HasLoginSubdomain returns a boolean if a field has been set.

### GetDisableInvite

`func (o *GetOrganizationAuthenticationSettings200Response) GetDisableInvite() bool`

GetDisableInvite returns the DisableInvite field if non-nil, zero value otherwise.

### GetDisableInviteOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetDisableInviteOk() (*bool, bool)`

GetDisableInviteOk returns a tuple with the DisableInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableInvite

`func (o *GetOrganizationAuthenticationSettings200Response) SetDisableInvite(v bool)`

SetDisableInvite sets DisableInvite field to given value.

### HasDisableInvite

`func (o *GetOrganizationAuthenticationSettings200Response) HasDisableInvite() bool`

HasDisableInvite returns a boolean if a field has been set.

### GetSaml

`func (o *GetOrganizationAuthenticationSettings200Response) GetSaml() GetOrganizationAuthenticationSettings200ResponseSaml`

GetSaml returns the Saml field if non-nil, zero value otherwise.

### GetSamlOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetSamlOk() (*GetOrganizationAuthenticationSettings200ResponseSaml, bool)`

GetSamlOk returns a tuple with the Saml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml

`func (o *GetOrganizationAuthenticationSettings200Response) SetSaml(v GetOrganizationAuthenticationSettings200ResponseSaml)`

SetSaml sets Saml field to given value.

### HasSaml

`func (o *GetOrganizationAuthenticationSettings200Response) HasSaml() bool`

HasSaml returns a boolean if a field has been set.

### GetLdap

`func (o *GetOrganizationAuthenticationSettings200Response) GetLdap() GetOrganizationAuthenticationSettings200ResponseLdap`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetLdapOk() (*GetOrganizationAuthenticationSettings200ResponseLdap, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *GetOrganizationAuthenticationSettings200Response) SetLdap(v GetOrganizationAuthenticationSettings200ResponseLdap)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *GetOrganizationAuthenticationSettings200Response) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetSso

`func (o *GetOrganizationAuthenticationSettings200Response) GetSso() GetOrganizationAuthenticationSettings200ResponseSso`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *GetOrganizationAuthenticationSettings200Response) GetSsoOk() (*GetOrganizationAuthenticationSettings200ResponseSso, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *GetOrganizationAuthenticationSettings200Response) SetSso(v GetOrganizationAuthenticationSettings200ResponseSso)`

SetSso sets Sso field to given value.

### HasSso

`func (o *GetOrganizationAuthenticationSettings200Response) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


