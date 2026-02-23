# CreateOIDCTrustRelationship201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The unique identifier of the OIDC Trust Relationship | 
**OrganizationId** | **int32** | The ID of the organization this trust relationship belongs to | 
**UserId** | **int32** | The ID of the user associated with this trust relationship | 
**ProviderId** | **int32** | The ID of the OIDC Provider this trust relationship uses | 
**Scope** | Pointer to **string** | The scope level of the trust relationship | [optional] 
**ScopeId** | Pointer to **int32** | The ID associated with the scope: either the organization ID or the team ID | [optional] 
**Audiences** | **[]string** | Array of valid audiences for this trust relationship | 
**RequiredClaims** | [**[]CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner**](CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner.md) | Array of claim objects. Must contain at least one object with claim: \&quot;sub\&quot;. Additional objects with other claims are optional.  | 
**CreatedAt** | **time.Time** | When the trust relationship was created | 
**UpdatedAt** | **time.Time** | When the trust relationship was last updated | 

## Methods

### NewCreateOIDCTrustRelationship201Response

`func NewCreateOIDCTrustRelationship201Response(id int32, organizationId int32, userId int32, providerId int32, audiences []string, requiredClaims []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, createdAt time.Time, updatedAt time.Time, ) *CreateOIDCTrustRelationship201Response`

NewCreateOIDCTrustRelationship201Response instantiates a new CreateOIDCTrustRelationship201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCTrustRelationship201ResponseWithDefaults

`func NewCreateOIDCTrustRelationship201ResponseWithDefaults() *CreateOIDCTrustRelationship201Response`

NewCreateOIDCTrustRelationship201ResponseWithDefaults instantiates a new CreateOIDCTrustRelationship201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateOIDCTrustRelationship201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateOIDCTrustRelationship201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateOIDCTrustRelationship201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *CreateOIDCTrustRelationship201Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateOIDCTrustRelationship201Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateOIDCTrustRelationship201Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *CreateOIDCTrustRelationship201Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateOIDCTrustRelationship201Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateOIDCTrustRelationship201Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *CreateOIDCTrustRelationship201Response) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateOIDCTrustRelationship201Response) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateOIDCTrustRelationship201Response) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetScope

`func (o *CreateOIDCTrustRelationship201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCTrustRelationship201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCTrustRelationship201Response) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateOIDCTrustRelationship201Response) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeId

`func (o *CreateOIDCTrustRelationship201Response) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateOIDCTrustRelationship201Response) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateOIDCTrustRelationship201Response) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *CreateOIDCTrustRelationship201Response) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetAudiences

`func (o *CreateOIDCTrustRelationship201Response) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CreateOIDCTrustRelationship201Response) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CreateOIDCTrustRelationship201Response) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetRequiredClaims

`func (o *CreateOIDCTrustRelationship201Response) GetRequiredClaims() []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *CreateOIDCTrustRelationship201Response) GetRequiredClaimsOk() (*[]CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *CreateOIDCTrustRelationship201Response) SetRequiredClaims(v []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.


### GetCreatedAt

`func (o *CreateOIDCTrustRelationship201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateOIDCTrustRelationship201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateOIDCTrustRelationship201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateOIDCTrustRelationship201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateOIDCTrustRelationship201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateOIDCTrustRelationship201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


