# GetOIDCProviderAvailableServiceAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner**](GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner.md) | Array of available service accounts | 
**PageSize** | **int32** | Number of items requested per page | [default to 10]
**Last** | **int32** | ID of the last service account in the results array for pagination | 

## Methods

### NewGetOIDCProviderAvailableServiceAccounts200Response

`func NewGetOIDCProviderAvailableServiceAccounts200Response(results []GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner, pageSize int32, last int32, ) *GetOIDCProviderAvailableServiceAccounts200Response`

NewGetOIDCProviderAvailableServiceAccounts200Response instantiates a new GetOIDCProviderAvailableServiceAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOIDCProviderAvailableServiceAccounts200ResponseWithDefaults

`func NewGetOIDCProviderAvailableServiceAccounts200ResponseWithDefaults() *GetOIDCProviderAvailableServiceAccounts200Response`

NewGetOIDCProviderAvailableServiceAccounts200ResponseWithDefaults instantiates a new GetOIDCProviderAvailableServiceAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetResults() []GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetResultsOk() (*[]GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) SetResults(v []GetOIDCProviderAvailableServiceAccounts200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetLast

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetOIDCProviderAvailableServiceAccounts200Response) SetLast(v int32)`

SetLast sets Last field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


