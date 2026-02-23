# GetAllReportOptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetAllReportOptions200ResponseResultsInner**](GetAllReportOptions200ResponseResultsInner.md) |  | 
**Page** | **int32** | Current page number | 
**PageSize** | **int32** | Number of items per page | 
**Total** | **int32** | Total number of report options | 

## Methods

### NewGetAllReportOptions200Response

`func NewGetAllReportOptions200Response(results []GetAllReportOptions200ResponseResultsInner, page int32, pageSize int32, total int32, ) *GetAllReportOptions200Response`

NewGetAllReportOptions200Response instantiates a new GetAllReportOptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllReportOptions200ResponseWithDefaults

`func NewGetAllReportOptions200ResponseWithDefaults() *GetAllReportOptions200Response`

NewGetAllReportOptions200ResponseWithDefaults instantiates a new GetAllReportOptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetAllReportOptions200Response) GetResults() []GetAllReportOptions200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetAllReportOptions200Response) GetResultsOk() (*[]GetAllReportOptions200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetAllReportOptions200Response) SetResults(v []GetAllReportOptions200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPage

`func (o *GetAllReportOptions200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetAllReportOptions200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetAllReportOptions200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *GetAllReportOptions200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetAllReportOptions200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetAllReportOptions200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotal

`func (o *GetAllReportOptions200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllReportOptions200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllReportOptions200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


