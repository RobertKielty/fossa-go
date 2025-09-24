# ListOIDCProviders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ListOIDCProviders200ResponseAllOfResultsInner**](ListOIDCProviders200ResponseAllOfResultsInner.md) | Array of OIDC Providers | 
**PageSize** | **int32** | Number of items requested per page | [default to 10]
**Last** | **int32** | ID of the last OIDC provider in the results | 

## Methods

### NewListOIDCProviders200Response

`func NewListOIDCProviders200Response(results []ListOIDCProviders200ResponseAllOfResultsInner, pageSize int32, last int32, ) *ListOIDCProviders200Response`

NewListOIDCProviders200Response instantiates a new ListOIDCProviders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOIDCProviders200ResponseWithDefaults

`func NewListOIDCProviders200ResponseWithDefaults() *ListOIDCProviders200Response`

NewListOIDCProviders200ResponseWithDefaults instantiates a new ListOIDCProviders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ListOIDCProviders200Response) GetResults() []ListOIDCProviders200ResponseAllOfResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListOIDCProviders200Response) GetResultsOk() (*[]ListOIDCProviders200ResponseAllOfResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListOIDCProviders200Response) SetResults(v []ListOIDCProviders200ResponseAllOfResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *ListOIDCProviders200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ListOIDCProviders200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ListOIDCProviders200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetLast

`func (o *ListOIDCProviders200Response) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *ListOIDCProviders200Response) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *ListOIDCProviders200Response) SetLast(v int32)`

SetLast sets Last field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


