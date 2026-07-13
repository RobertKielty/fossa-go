# GetSnippetPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetSnippetPackages200ResponseResultsInner**](GetSnippetPackages200ResponseResultsInner.md) | Array of snippet packages | 
**TotalCount** | **int32** | Total number of snippet packages matching the filter criteria | 
**Page** | **int32** | Current page number | 
**PageSize** | **int32** | Number of items requested per page | 

## Methods

### NewGetSnippetPackages200Response

`func NewGetSnippetPackages200Response(results []GetSnippetPackages200ResponseResultsInner, totalCount int32, page int32, pageSize int32, ) *GetSnippetPackages200Response`

NewGetSnippetPackages200Response instantiates a new GetSnippetPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetPackages200ResponseWithDefaults

`func NewGetSnippetPackages200ResponseWithDefaults() *GetSnippetPackages200Response`

NewGetSnippetPackages200ResponseWithDefaults instantiates a new GetSnippetPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetSnippetPackages200Response) GetResults() []GetSnippetPackages200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSnippetPackages200Response) GetResultsOk() (*[]GetSnippetPackages200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSnippetPackages200Response) SetResults(v []GetSnippetPackages200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *GetSnippetPackages200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetSnippetPackages200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetSnippetPackages200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *GetSnippetPackages200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSnippetPackages200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSnippetPackages200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *GetSnippetPackages200Response) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetSnippetPackages200Response) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetSnippetPackages200Response) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


