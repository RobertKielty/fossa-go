# GetTeamMembers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetTeamMembers200ResponseResultsInner**](GetTeamMembers200ResponseResultsInner.md) | Array of team members | 
**PageSize** | **int32** | Number of items per page | 
**Page** | **int32** | Current page number (1-indexed) | 
**TotalCount** | **int32** | Total number of team members | 

## Methods

### NewGetTeamMembers200Response

`func NewGetTeamMembers200Response(results []GetTeamMembers200ResponseResultsInner, pageSize int32, page int32, totalCount int32, ) *GetTeamMembers200Response`

NewGetTeamMembers200Response instantiates a new GetTeamMembers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamMembers200ResponseWithDefaults

`func NewGetTeamMembers200ResponseWithDefaults() *GetTeamMembers200Response`

NewGetTeamMembers200ResponseWithDefaults instantiates a new GetTeamMembers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetTeamMembers200Response) GetResults() []GetTeamMembers200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetTeamMembers200Response) GetResultsOk() (*[]GetTeamMembers200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetTeamMembers200Response) SetResults(v []GetTeamMembers200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetPageSize

`func (o *GetTeamMembers200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetTeamMembers200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetTeamMembers200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetTeamMembers200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetTeamMembers200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetTeamMembers200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetTotalCount

`func (o *GetTeamMembers200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetTeamMembers200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetTeamMembers200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


