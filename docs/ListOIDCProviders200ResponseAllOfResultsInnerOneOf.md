# ListOIDCProviders200ResponseAllOfResultsInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the OIDC Provider | 
**OrganizationId** | **int32** | The ID of the organization this provider belongs to | 
**Issuer** | **string** | The issuer URL of the OIDC Provider | 
**Scope** | **string** | The scope level of the OIDC Provider | 
**ScopeId** | **int32** | The organization ID | 
**TeamName** | Pointer to [**nil**](nil.md) | Team name (null for org-scoped providers) | [optional] 
**CreatedAt** | **time.Time** | When the OIDC Provider was created | 
**UpdatedAt** | **time.Time** | When the OIDC Provider was last updated | 

## Methods

### NewListOIDCProviders200ResponseAllOfResultsInnerOneOf

`func NewListOIDCProviders200ResponseAllOfResultsInnerOneOf(id int32, organizationId int32, issuer string, scope string, scopeId int32, createdAt time.Time, updatedAt time.Time, ) *ListOIDCProviders200ResponseAllOfResultsInnerOneOf`

NewListOIDCProviders200ResponseAllOfResultsInnerOneOf instantiates a new ListOIDCProviders200ResponseAllOfResultsInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCProviders200ResponseAllOfResultsInnerOneOfWithDefaults

`func NewListOIDCProviders200ResponseAllOfResultsInnerOneOfWithDefaults() *ListOIDCProviders200ResponseAllOfResultsInnerOneOf`

NewListOIDCProviders200ResponseAllOfResultsInnerOneOfWithDefaults instantiates a new ListOIDCProviders200ResponseAllOfResultsInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetIssuer

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetScope

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.


### GetTeamName

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetTeamName() nil`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetTeamNameOk() (*nil, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetTeamName(v nil)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


