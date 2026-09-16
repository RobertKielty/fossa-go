# ListOIDCTrustRelationships200ResponseAllOfResultsInner

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
**RequiredClaims** | [**[]ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner**](ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner.md) | Array of claim objects. Must contain at least one object with claim: \&quot;sub\&quot;. Additional objects with other claims are optional.  | 
**CreatedAt** | **time.Time** | When the trust relationship was created | 
**UpdatedAt** | **time.Time** | When the trust relationship was last updated | 
**Username** | **string** | The username of the service account associated with this trust relationship | 
**Email** | Pointer to **NullableString** | The email of the service account associated with this trust relationship (&#x60;null&#x60; if the account has no email) | [optional] 
**TeamName** | Pointer to **NullableString** | The name of the team. &#x60;null&#x60; for organization-scoped trust relationships (only team-scoped relationships have a team name). | [optional] 
**Issuer** | **string** | The issuer URL of the OIDC Provider | 

## Methods

### NewListOIDCTrustRelationships200ResponseAllOfResultsInner

`func NewListOIDCTrustRelationships200ResponseAllOfResultsInner(id int32, organizationId int32, userId int32, providerId int32, audiences []string, requiredClaims []ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner, createdAt time.Time, updatedAt time.Time, username string, issuer string, ) *ListOIDCTrustRelationships200ResponseAllOfResultsInner`

NewListOIDCTrustRelationships200ResponseAllOfResultsInner instantiates a new ListOIDCTrustRelationships200ResponseAllOfResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCTrustRelationships200ResponseAllOfResultsInnerWithDefaults

`func NewListOIDCTrustRelationships200ResponseAllOfResultsInnerWithDefaults() *ListOIDCTrustRelationships200ResponseAllOfResultsInner`

NewListOIDCTrustRelationships200ResponseAllOfResultsInnerWithDefaults instantiates a new ListOIDCTrustRelationships200ResponseAllOfResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetScope

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetAudiences

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetRequiredClaims

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetRequiredClaims() []ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetRequiredClaimsOk() (*[]ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetRequiredClaims(v []ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.


### GetCreatedAt

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUsername

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetTeamName

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.

### HasTeamName

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) HasTeamName() bool`

HasTeamName returns a boolean if a field has been set.

### SetTeamNameNil

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetTeamNameNil(b bool)`

 SetTeamNameNil sets the value for TeamName to be an explicit nil

### UnsetTeamName
`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) UnsetTeamName()`

UnsetTeamName ensures that no value is present for TeamName, not even an explicit nil
### GetIssuer

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInner) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


