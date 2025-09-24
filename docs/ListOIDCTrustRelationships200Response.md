# ListOIDCTrustRelationships200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ListOIDCTrustRelationships200ResponseAllOfResultsInner**](ListOIDCTrustRelationships200ResponseAllOfResultsInner.md) | Array of OIDC Trust Relationships | 
**PageSize** | **int32** | Number of items requested per page | [default to 10]
**Last** | **int32** | ID of the last OIDC trust relationship in the results | 

## Methods

### NewListOIDCTrustRelationships200Response

`func NewListOIDCTrustRelationships200Response(results []ListOIDCTrustRelationships200ResponseAllOfResultsInner, pageSize int32, last int32, ) *ListOIDCTrustRelationships200Response`

NewListOIDCTrustRelationships200Response instantiates a new ListOIDCTrustRelationships200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCTrustRelationships200ResponseWithDefaults

`func NewListOIDCTrustRelationships200ResponseWithDefaults() *ListOIDCTrustRelationships200Response`

NewListOIDCTrustRelationships200ResponseWithDefaults instantiates a new ListOIDCTrustRelationships200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListOIDCTrustRelationships200Response) GetResults() []ListOIDCTrustRelationships200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListOIDCTrustRelationships200Response) GetResultsOk() (*[]ListOIDCTrustRelationships200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListOIDCTrustRelationships200Response) SetResults(v []ListOIDCTrustRelationships200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *ListOIDCTrustRelationships200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListOIDCTrustRelationships200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListOIDCTrustRelationships200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetLast

`func (o *ListOIDCTrustRelationships200Response) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ListOIDCTrustRelationships200Response) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ListOIDCTrustRelationships200Response) SetLast(v int32)`

SetLast sets Last field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


