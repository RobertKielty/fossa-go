# CreateOIDCTrustRelationshipRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | The ID of the user to associate with this trust relationship | 
**ProviderId** | **int32** | The ID of the OIDC Provider to use for this trust relationship | 
**Scope** | **string** | The scope level of the trust relationship | 
**Audiences** | **[]string** | Array of valid audiences for this trust relationship | 
**RequiredClaims** | [**[]CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner**](CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner.md) | Array of claim objects. Must contain at least one object with claim: \&quot;sub\&quot;. Additional objects with other claims are optional.  | 

## Methods

### NewCreateOIDCTrustRelationshipRequestOneOf

`func NewCreateOIDCTrustRelationshipRequestOneOf(userId int32, providerId int32, scope string, audiences []string, requiredClaims []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, ) *CreateOIDCTrustRelationshipRequestOneOf`

NewCreateOIDCTrustRelationshipRequestOneOf instantiates a new CreateOIDCTrustRelationshipRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCTrustRelationshipRequestOneOfWithDefaults

`func NewCreateOIDCTrustRelationshipRequestOneOfWithDefaults() *CreateOIDCTrustRelationshipRequestOneOf`

NewCreateOIDCTrustRelationshipRequestOneOfWithDefaults instantiates a new CreateOIDCTrustRelationshipRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateOIDCTrustRelationshipRequestOneOf) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *CreateOIDCTrustRelationshipRequestOneOf) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetScope

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateOIDCTrustRelationshipRequestOneOf) SetScope(v string)`

SetScope sets Scope field to given value.


### GetAudiences

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *CreateOIDCTrustRelationshipRequestOneOf) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.


### GetRequiredClaims

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetRequiredClaims() []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *CreateOIDCTrustRelationshipRequestOneOf) GetRequiredClaimsOk() (*[]CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *CreateOIDCTrustRelationshipRequestOneOf) SetRequiredClaims(v []CreateOIDCTrustRelationshipRequestOneOfRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


