# GetAllTeamsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetAllTeamsV2200ResponseResultsInner**](GetAllTeamsV2200ResponseResultsInner.md) | Array of teams | 
**PageSize** | **int32** | Number of items per page | 
**Page** | **int32** | Current page number (1-indexed) | 
**TotalCount** | **int32** | Total number of teams | 

## Methods

### NewGetAllTeamsV2200Response

`func NewGetAllTeamsV2200Response(results []GetAllTeamsV2200ResponseResultsInner, pageSize int32, page int32, totalCount int32, ) *GetAllTeamsV2200Response`

NewGetAllTeamsV2200Response instantiates a new GetAllTeamsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllTeamsV2200ResponseWithDefaults

`func NewGetAllTeamsV2200ResponseWithDefaults() *GetAllTeamsV2200Response`

NewGetAllTeamsV2200ResponseWithDefaults instantiates a new GetAllTeamsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetAllTeamsV2200Response) GetResults() []GetAllTeamsV2200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetAllTeamsV2200Response) GetResultsOk() (*[]GetAllTeamsV2200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetAllTeamsV2200Response) SetResults(v []GetAllTeamsV2200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *GetAllTeamsV2200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetAllTeamsV2200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetAllTeamsV2200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetAllTeamsV2200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetAllTeamsV2200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetAllTeamsV2200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetTotalCount

`func (o *GetAllTeamsV2200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetAllTeamsV2200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetAllTeamsV2200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


