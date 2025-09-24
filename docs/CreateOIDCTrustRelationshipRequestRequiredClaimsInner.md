# CreateOIDCTrustRelationshipRequestRequiredClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | **string** | The claim type identifier | 
**Value** | [**CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue**](CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue.md) |  | 
**HasWildcards** | Pointer to **bool** | Whether this claim supports wildcard matching. Defaults to false. If true, then the following characters in &#x60;value&#x60; are treated as special characters: - &#x60;?&#x60; matches exactly one character - &#x60;*&#x60; matches zero or more characters - &#x60;\\&#x60; escapes the following character  | [optional] [default to false]

## Methods

### NewCreateOIDCTrustRelationshipRequestRequiredClaimsInner

`func NewCreateOIDCTrustRelationshipRequestRequiredClaimsInner(claim string, value CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue, ) *CreateOIDCTrustRelationshipRequestRequiredClaimsInner`

NewCreateOIDCTrustRelationshipRequestRequiredClaimsInner instantiates a new CreateOIDCTrustRelationshipRequestRequiredClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOIDCTrustRelationshipRequestRequiredClaimsInnerWithDefaults

`func NewCreateOIDCTrustRelationshipRequestRequiredClaimsInnerWithDefaults() *CreateOIDCTrustRelationshipRequestRequiredClaimsInner`

NewCreateOIDCTrustRelationshipRequestRequiredClaimsInnerWithDefaults instantiates a new CreateOIDCTrustRelationshipRequestRequiredClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetClaimOk() (*string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) SetClaim(v string)`

SetClaim sets Claim field to given value.


### GetValue

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetValue() CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetValueOk() (*CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) SetValue(v CreateOIDCTrustRelationshipRequestRequiredClaimsInnerValue)`

SetValue sets Value field to given value.


### GetHasWildcards

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetHasWildcards() bool`

GetHasWildcards returns the HasWildcards field if non-nil, zero value otherwise.

### GetHasWildcardsOk

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) GetHasWildcardsOk() (*bool, bool)`

GetHasWildcardsOk returns a tuple with the HasWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWildcards

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) SetHasWildcards(v bool)`

SetHasWildcards sets HasWildcards field to given value.

### HasHasWildcards

`func (o *CreateOIDCTrustRelationshipRequestRequiredClaimsInner) HasHasWildcards() bool`

HasHasWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


