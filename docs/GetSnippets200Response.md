# GetSnippets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]GetSnippets200ResponseResultsInner**](GetSnippets200ResponseResultsInner.md) | Array of snippets | 
**TotalCount** | **int32** | Total number of snippets matching the filter criteria | 
**Page** | **int32** | Current page number | 
**Count** | **int32** | Number of items per page | 

## Methods

### NewGetSnippets200Response

`func NewGetSnippets200Response(results []GetSnippets200ResponseResultsInner, totalCount int32, page int32, count int32, ) *GetSnippets200Response`

NewGetSnippets200Response instantiates a new GetSnippets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippets200ResponseWithDefaults

`func NewGetSnippets200ResponseWithDefaults() *GetSnippets200Response`

NewGetSnippets200ResponseWithDefaults instantiates a new GetSnippets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetSnippets200Response) GetResults() []GetSnippets200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSnippets200Response) GetResultsOk() (*[]GetSnippets200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSnippets200Response) SetResults(v []GetSnippets200ResponseResultsInner)`

SetResults sets Results field to given value.


### GetTotalCount

`func (o *GetSnippets200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetSnippets200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetSnippets200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetPage

`func (o *GetSnippets200Response) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetSnippets200Response) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetSnippets200Response) SetPage(v int32)`

SetPage sets Page field to given value.


### GetCount

`func (o *GetSnippets200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetSnippets200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetSnippets200Response) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


