# GetOIDCTrustRelationship200Response

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
**Username** | **string** | The username of the service account associated with this trust relationship | 
**Email** | Pointer to **NullableString** | The email of the service account associated with this trust relationship (&#x60;null&#x60; if the account has no email) | [optional] 
**TeamName** | Pointer to **NullableString** | The name of the team. &#x60;null&#x60; for organization-scoped trust relationships (only team-scoped relationships have a team name). | [optional] 
**Issuer** | **string** | The issuer URL of the OIDC Provider | 

## Methods

### NewGetOIDCTrustRelationship200Response

`func NewGetOIDCTrustRelationship200Response(id int32, organizationId int32, userId int32, providerId int32, audiences []string, requiredClaims []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, createdAt time.Time, updatedAt time.Time, username string, issuer string, ) *GetOIDCTrustRelationship200Response`

NewGetOIDCTrustRelationship200Response instantiates a new GetOIDCTrustRelationship200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOIDCTrustRelationship200ResponseWithDefaults

`func NewGetOIDCTrustRelationship200ResponseWithDefaults() *GetOIDCTrustRelationship200Response`

NewGetOIDCTrustRelationship200ResponseWithDefaults instantiates a new GetOIDCTrustRelationship200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOIDCTrustRelationship200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOIDCTrustRelationship200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOIDCTrustRelationship200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *GetOIDCTrustRelationship200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOIDCTrustRelationship200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOIDCTrustRelationship200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *GetOIDCTrustRelationship200Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetOIDCTrustRelationship200Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetOIDCTrustRelationship200Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *GetOIDCTrustRelationship200Response) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *GetOIDCTrustRelationship200Response) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *GetOIDCTrustRelationship200Response) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetScope

`func (o *GetOIDCTrustRelationship200Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetOIDCTrustRelationship200Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetOIDCTrustRelationship200Response) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetOIDCTrustRelationship200Response) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeId

`func (o *GetOIDCTrustRelationship200Response) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *GetOIDCTrustRelationship200Response) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *GetOIDCTrustRelationship200Response) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *GetOIDCTrustRelationship200Response) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetAudiences

`func (o *GetOIDCTrustRelationship200Response) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *GetOIDCTrustRelationship200Response) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *GetOIDCTrustRelationship200Response) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetRequiredClaims

`func (o *GetOIDCTrustRelationship200Response) GetRequiredClaims() []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *GetOIDCTrustRelationship200Response) GetRequiredClaimsOk() (*[]CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *GetOIDCTrustRelationship200Response) SetRequiredClaims(v []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.


### GetCreatedAt

`func (o *GetOIDCTrustRelationship200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOIDCTrustRelationship200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOIDCTrustRelationship200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetOIDCTrustRelationship200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOIDCTrustRelationship200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOIDCTrustRelationship200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUsername

`func (o *GetOIDCTrustRelationship200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetOIDCTrustRelationship200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetOIDCTrustRelationship200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *GetOIDCTrustRelationship200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOIDCTrustRelationship200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOIDCTrustRelationship200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOIDCTrustRelationship200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetOIDCTrustRelationship200Response) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetOIDCTrustRelationship200Response) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTeamName

`func (o *GetOIDCTrustRelationship200Response) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *GetOIDCTrustRelationship200Response) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *GetOIDCTrustRelationship200Response) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *GetOIDCTrustRelationship200Response) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamNameNil

`func (o *GetOIDCTrustRelationship200Response) SetTeamNameNil(b bool)`

 SetTeamNameNil sets the value for TeamName to be an explicit nil

### UnsetTeamName
`func (o *GetOIDCTrustRelationship200Response) UnsetTeamName()`

UnsetTeamName ensures that no value is present for TeamName, not even an explicit nil
### GetIssuer

`func (o *GetOIDCTrustRelationship200Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetOIDCTrustRelationship200Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetOIDCTrustRelationship200Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


