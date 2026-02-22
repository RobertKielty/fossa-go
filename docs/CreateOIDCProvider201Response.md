# CreateOIDCProvider201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the OIDC Provider | 
**OrganizationId** | **int32** | The ID of the organization this provider belongs to | 
**Issuer** | **string** | The issuer URL of the OIDC Provider | 
**Scope** | **string** | The scope level of the OIDC Provider | 
**ScopeId** | **int32** | The ID associated with the scope: either the organization ID or the team ID | 
**CreatedAt** | **time.Time** | When the OIDC Provider was created | 
**UpdatedAt** | **time.Time** | When the OIDC Provider was last updated | 

## Methods

### NewCreateOIDCProvider201Response

`func NewCreateOIDCProvider201Response(id int32, organizationId int32, issuer string, scope string, scopeId int32, createdAt time.Time, updatedAt time.Time, ) *CreateOIDCProvider201Response`

NewCreateOIDCProvider201Response instantiates a new CreateOIDCProvider201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCProvider201ResponseWithDefaults

`func NewCreateOIDCProvider201ResponseWithDefaults() *CreateOIDCProvider201Response`

NewCreateOIDCProvider201ResponseWithDefaults instantiates a new CreateOIDCProvider201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOIDCProvider201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOIDCProvider201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOIDCProvider201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *CreateOIDCProvider201Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOIDCProvider201Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOIDCProvider201Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetIssuer

`func (o *CreateOIDCProvider201Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateOIDCProvider201Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateOIDCProvider201Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetScope

`func (o *CreateOIDCProvider201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCProvider201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCProvider201Response) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *CreateOIDCProvider201Response) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateOIDCProvider201Response) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateOIDCProvider201Response) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.


### GetCreatedAt

`func (o *CreateOIDCProvider201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateOIDCProvider201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateOIDCProvider201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateOIDCProvider201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateOIDCProvider201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateOIDCProvider201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


