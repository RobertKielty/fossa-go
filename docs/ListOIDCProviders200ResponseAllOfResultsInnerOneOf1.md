# ListOIDCProviders200ResponseAllOfResultsInnerOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the OIDC Provider | 
**OrganizationId** | **int32** | The ID of the organization this provider belongs to | 
**Issuer** | **string** | The issuer URL of the OIDC Provider | 
**Scope** | **string** | The scope level of the OIDC Provider | 
**ScopeId** | **int32** | The team ID | 
**TeamName** | **string** | The name of the team | 
**CreatedAt** | **time.Time** | When the OIDC Provider was created | 
**UpdatedAt** | **time.Time** | When the OIDC Provider was last updated | 

## Methods

### NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1

`func NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1(id int32, organizationId int32, issuer string, scope string, scopeId int32, teamName string, createdAt time.Time, updatedAt time.Time, ) *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1`

NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1 instantiates a new ListOIDCProviders200ResponseAllOfResultsInnerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1WithDefaults

`func NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1WithDefaults() *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1`

NewListOIDCProviders200ResponseAllOfResultsInnerOneOf1WithDefaults instantiates a new ListOIDCProviders200ResponseAllOfResultsInnerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetIssuer

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetScope

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.


### GetTeamName

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetCreatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListOIDCProviders200ResponseAllOfResultsInnerOneOf1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


