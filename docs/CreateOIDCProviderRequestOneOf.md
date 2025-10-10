# CreateOIDCProviderRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | The issuer URL of the OIDC Provider | 
**Scope** | **string** | The scope level of the OIDC Provider | 
**ScopeId** | **int32** | The ID of the team | 

## Methods

### NewCreateOIDCProviderRequestOneOf

`func NewCreateOIDCProviderRequestOneOf(issuer string, scope string, scopeId int32, ) *CreateOIDCProviderRequestOneOf`

NewCreateOIDCProviderRequestOneOf instantiates a new CreateOIDCProviderRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCProviderRequestOneOfWithDefaults

`func NewCreateOIDCProviderRequestOneOfWithDefaults() *CreateOIDCProviderRequestOneOf`

NewCreateOIDCProviderRequestOneOfWithDefaults instantiates a new CreateOIDCProviderRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *CreateOIDCProviderRequestOneOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateOIDCProviderRequestOneOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateOIDCProviderRequestOneOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetScope

`func (o *CreateOIDCProviderRequestOneOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCProviderRequestOneOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCProviderRequestOneOf) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *CreateOIDCProviderRequestOneOf) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateOIDCProviderRequestOneOf) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateOIDCProviderRequestOneOf) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


