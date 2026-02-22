# UpdateOIDCTrustRelationshipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** | Array of valid audiences for this trust relationship | [optional] 
**RequiredClaims** | Pointer to [**[]ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner**](ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner.md) | Array of claim objects. Must contain at least one object with claim: \&quot;sub\&quot;. Additional objects with other claims are optional.  | [optional] 

## Methods

### NewUpdateOIDCTrustRelationshipRequest

`func NewUpdateOIDCTrustRelationshipRequest() *UpdateOIDCTrustRelationshipRequest`

NewUpdateOIDCTrustRelationshipRequest instantiates a new UpdateOIDCTrustRelationshipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOIDCTrustRelationshipRequestWithDefaults

`func NewUpdateOIDCTrustRelationshipRequestWithDefaults() *UpdateOIDCTrustRelationshipRequest`

NewUpdateOIDCTrustRelationshipRequestWithDefaults instantiates a new UpdateOIDCTrustRelationshipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *UpdateOIDCTrustRelationshipRequest) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *UpdateOIDCTrustRelationshipRequest) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *UpdateOIDCTrustRelationshipRequest) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *UpdateOIDCTrustRelationshipRequest) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetRequiredClaims

`func (o *UpdateOIDCTrustRelationshipRequest) GetRequiredClaims() []ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner`

GetRequiredClaims returns the RequiredClaims field if non-nil, zero value otherwise.

### GetRequiredClaimsOk

`func (o *UpdateOIDCTrustRelationshipRequest) GetRequiredClaimsOk() (*[]ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner, bool)`

GetRequiredClaimsOk returns a tuple with the RequiredClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredClaims

`func (o *UpdateOIDCTrustRelationshipRequest) SetRequiredClaims(v []ListOIDCTrustRelationships200ResponseAllOfResultsInnerAllOfRequiredClaimsInner)`

SetRequiredClaims sets RequiredClaims field to given value.

### HasRequiredClaims

`func (o *UpdateOIDCTrustRelationshipRequest) HasRequiredClaims() bool`

HasRequiredClaims returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


