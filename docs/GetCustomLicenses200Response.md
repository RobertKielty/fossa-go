# GetCustomLicenses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetCustomLicenses200ResponseResultsInner**](GetCustomLicenses200ResponseResultsInner.md) | Array of custom license objects | 
**TotalCount** | **int32** | Total number of custom licenses available (across all pages) | 
**Page** | **int32** | Current page number | 
**PageSize** | **int32** | Number of items per page | 

## Methods

### NewGetCustomLicenses200Response

`func NewGetCustomLicenses200Response(results []GetCustomLicenses200ResponseResultsInner, totalCount int32, page int32, pageSize int32, ) *GetCustomLicenses200Response`

NewGetCustomLicenses200Response instantiates a new GetCustomLicenses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomLicenses200ResponseWithDefaults

`func NewGetCustomLicenses200ResponseWithDefaults() *GetCustomLicenses200Response`

NewGetCustomLicenses200ResponseWithDefaults instantiates a new GetCustomLicenses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetCustomLicenses200Response) GetResults() []GetCustomLicenses200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetCustomLicenses200Response) GetResultsOk() (*[]GetCustomLicenses200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetCustomLicenses200Response) SetResults(v []GetCustomLicenses200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *GetCustomLicenses200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetCustomLicenses200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetCustomLicenses200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *GetCustomLicenses200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetCustomLicenses200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetCustomLicenses200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *GetCustomLicenses200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetCustomLicenses200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetCustomLicenses200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


