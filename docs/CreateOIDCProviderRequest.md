# CreateOIDCProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | The issuer URL of the OIDC Provider | 
**Thumbprints** | Pointer to **[]string** | Array of thumbprints for certificate validation | [optional] 
**Scope** | **string** | The scope level of the OIDC Provider | 
**ScopeId** | **string** | The ID associated with the scope | 

## Methods

### NewCreateOIDCProviderRequest

`func NewCreateOIDCProviderRequest(issuer string, scope string, scopeId string, ) *CreateOIDCProviderRequest`

NewCreateOIDCProviderRequest instantiates a new CreateOIDCProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCProviderRequestWithDefaults

`func NewCreateOIDCProviderRequestWithDefaults() *CreateOIDCProviderRequest`

NewCreateOIDCProviderRequestWithDefaults instantiates a new CreateOIDCProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *CreateOIDCProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateOIDCProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateOIDCProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetThumbprints

`func (o *CreateOIDCProviderRequest) GetThumbprints() []string`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *CreateOIDCProviderRequest) GetThumbprintsOk() (*[]string, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *CreateOIDCProviderRequest) SetThumbprints(v []string)`

SetThumbprints sets Thumbprints field to given value.

### HasThumbprints

`func (o *CreateOIDCProviderRequest) HasThumbprints() bool`

HasThumbprints returns a boolean if a field has been set.

### GetScope

`func (o *CreateOIDCProviderRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCProviderRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCProviderRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *CreateOIDCProviderRequest) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateOIDCProviderRequest) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateOIDCProviderRequest) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


