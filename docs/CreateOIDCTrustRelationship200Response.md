# CreateOIDCTrustRelationship200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the OIDC Trust Relationship | 
**OrganizationId** | **int32** | The ID of the organization this trust relationship belongs to | 
**UserId** | **int32** | The ID of the user associated with this trust relationship | 
**ProviderId** | **int32** | The ID of the OIDC Provider this trust relationship uses | 
**Scope** | Pointer to **string** | The scope level of the trust relationship | [optional] 
**ScopeId** | Pointer to **string** | The ID associated with the scope | [optional] 
**Audiences** | **[]string** | Array of valid audiences for this trust relationship (max 5) | 
**RequiredClaims** | [**[]CreateOIDCTrustRelationshipRequestRequiredClaimsInner**](CreateOIDCTrustRelationshipRequestRequiredClaimsInner.md) | Array of claim objects. Must contain at least one object with claim: \&quot;sub\&quot;. Additional objects with other claims are optional.  | 
**CreatedAt** | **time.Time** | When the trust relationship was created | 
**UpdatedAt** | **time.Time** | When the trust relationship was last updated | 

## Methods

### NewCreateOIDCTrustRelationship200Response

`func NewCreateOIDCTrustRelationship200Response(id int32, organizationId int32, userId int32, providerId int32, audiences []string, requiredClaims []CreateOIDCTrustRelationshipRequestRequiredClaimsInner, createdAt time.Time, updatedAt time.Time, ) *CreateOIDCTrustRelationship200Response`

NewCreateOIDCTrustRelationship200Response instantiates a new CreateOIDCTrustRelationship200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCTrustRelationship200ResponseWithDefaults

`func NewCreateOIDCTrustRelationship200ResponseWithDefaults() *CreateOIDCTrustRelationship200Response`

NewCreateOIDCTrustRelationship200ResponseWithDefaults instantiates a new CreateOIDCTrustRelationship200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOIDCTrustRelationship200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOIDCTrustRelationship200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOIDCTrustRelationship200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *CreateOIDCTrustRelationship200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOIDCTrustRelationship200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOIDCTrustRelationship200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *CreateOIDCTrustRelationship200Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateOIDCTrustRelationship200Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateOIDCTrustRelationship200Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *CreateOIDCTrustRelationship200Response) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateOIDCTrustRelationship200Response) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateOIDCTrustRelationship200Response) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetScope

`func (o *CreateOIDCTrustRelationship200Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCTrustRelationship200Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCTrustRelationship200Response) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateOIDCTrustRelationship200Response) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeId

`func (o *CreateOIDCTrustRelationship200Response) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateOIDCTrustRelationship200Response) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateOIDCTrustRelationship200Response) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *CreateOIDCTrustRelationship200Response) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetAudiences

`func (o *CreateOIDCTrustRelationship200Response) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CreateOIDCTrustRelationship200Response) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CreateOIDCTrustRelationship200Response) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetRequiredClaims

`func (o *CreateOIDCTrustRelationship200Response) GetRequiredClaims() []CreateOIDCTrustRelationshipRequestRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *CreateOIDCTrustRelationship200Response) GetRequiredClaimsOk() (*[]CreateOIDCTrustRelationshipRequestRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *CreateOIDCTrustRelationship200Response) SetRequiredClaims(v []CreateOIDCTrustRelationshipRequestRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.


### GetCreatedAt

`func (o *CreateOIDCTrustRelationship200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateOIDCTrustRelationship200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateOIDCTrustRelationship200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateOIDCTrustRelationship200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateOIDCTrustRelationship200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateOIDCTrustRelationship200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


