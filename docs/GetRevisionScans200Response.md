# GetRevisionScans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetRevisionScans200ResponseResultsInner**](GetRevisionScans200ResponseResultsInner.md) |  | 
**Page** | **int32** | The current page number | 
**PageSize** | **int32** | The number of items requested per page | 
**TotalCount** | **int32** | The total number of scans matching the query across all pages | 

## Methods

### NewGetRevisionScans200Response

`func NewGetRevisionScans200Response(results []GetRevisionScans200ResponseResultsInner, page int32, pageSize int32, totalCount int32, ) *GetRevisionScans200Response`

NewGetRevisionScans200Response instantiates a new GetRevisionScans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionScans200ResponseWithDefaults

`func NewGetRevisionScans200ResponseWithDefaults() *GetRevisionScans200Response`

NewGetRevisionScans200ResponseWithDefaults instantiates a new GetRevisionScans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetRevisionScans200Response) GetResults() []GetRevisionScans200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetRevisionScans200Response) GetResultsOk() (*[]GetRevisionScans200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetRevisionScans200Response) SetResults(v []GetRevisionScans200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPage

`func (o *GetRevisionScans200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetRevisionScans200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetRevisionScans200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *GetRevisionScans200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetRevisionScans200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetRevisionScans200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalCount

`func (o *GetRevisionScans200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetRevisionScans200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetRevisionScans200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


