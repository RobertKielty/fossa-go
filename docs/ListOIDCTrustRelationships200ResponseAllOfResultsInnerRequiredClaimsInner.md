# ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Claim** | **string** | The claim type identifier | 
**Value** | [**ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue**](ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue.md) |  | 
**HasWildcards** | Pointer to **bool** | Whether this claim supports wildcard matching. Defaults to false. If true, then the following characters in &#x60;value&#x60; are treated as special characters: - &#x60;?&#x60; matches exactly one character - &#x60;*&#x60; matches zero or more characters - &#x60;\\&#x60; escapes the following character  | [optional] [default to false]

## Methods

### NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner

`func NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner(claim string, value ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue, ) *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner`

NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner instantiates a new ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerWithDefaults

`func NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerWithDefaults() *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner`

NewListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerWithDefaults instantiates a new ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaim

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetClaimOk() (*string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) SetClaim(v string)`

SetClaim sets Claim field to given value.


### GetValue

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetValue() ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetValueOk() (*ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) SetValue(v ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInnerValue)`

SetValue sets Value field to given value.


### GetHasWildcards

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetHasWildcards() bool`

GetHasWildcards returns the HasWildcards field if non-nil, zero value otherwise.

### GetHasWildcardsOk

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) GetHasWildcardsOk() (*bool, bool)`

GetHasWildcardsOk returns a tuple with the HasWildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWildcards

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) SetHasWildcards(v bool)`

SetHasWildcards sets HasWildcards field to given value.

### HasHasWildcards

`func (o *ListOIDCTrustRelationships200ResponseAllOfResultsInnerRequiredClaimsInner) HasHasWildcards() bool`

HasHasWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


