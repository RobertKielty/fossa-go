# GetAddableTeamUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetAddableTeamUsers200ResponseResultsInner**](GetAddableTeamUsers200ResponseResultsInner.md) | Array of users that can be added to the team | 
**PageSize** | **int32** | Number of items per page | 
**Page** | **int32** | Current page number (1-indexed) | 
**TotalCount** | **int32** | Total number of users that can be added to the team | 

## Methods

### NewGetAddableTeamUsers200Response

`func NewGetAddableTeamUsers200Response(results []GetAddableTeamUsers200ResponseResultsInner, pageSize int32, page int32, totalCount int32, ) *GetAddableTeamUsers200Response`

NewGetAddableTeamUsers200Response instantiates a new GetAddableTeamUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAddableTeamUsers200ResponseWithDefaults

`func NewGetAddableTeamUsers200ResponseWithDefaults() *GetAddableTeamUsers200Response`

NewGetAddableTeamUsers200ResponseWithDefaults instantiates a new GetAddableTeamUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetAddableTeamUsers200Response) GetResults() []GetAddableTeamUsers200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetAddableTeamUsers200Response) GetResultsOk() (*[]GetAddableTeamUsers200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetAddableTeamUsers200Response) SetResults(v []GetAddableTeamUsers200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *GetAddableTeamUsers200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetAddableTeamUsers200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetAddableTeamUsers200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetAddableTeamUsers200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetAddableTeamUsers200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetAddableTeamUsers200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetTotalCount

`func (o *GetAddableTeamUsers200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetAddableTeamUsers200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetAddableTeamUsers200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


