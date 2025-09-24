# GetRevisionComponentMatches200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]GetRevisionComponentMatches200ResponseResultsInner**](GetRevisionComponentMatches200ResponseResultsInner.md) |  | [optional] 
**Page** | Pointer to **float32** | pagination page | [optional] 
**PageSize** | Pointer to **float32** | pagination page page size | [optional] 
**TotalCount** | Pointer to **float32** | total count of component matches | [optional] 

## Methods

### NewGetRevisionComponentMatches200Response

`func NewGetRevisionComponentMatches200Response() *GetRevisionComponentMatches200Response`

NewGetRevisionComponentMatches200Response instantiates a new GetRevisionComponentMatches200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionComponentMatches200ResponseWithDefaults

`func NewGetRevisionComponentMatches200ResponseWithDefaults() *GetRevisionComponentMatches200Response`

NewGetRevisionComponentMatches200ResponseWithDefaults instantiates a new GetRevisionComponentMatches200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetRevisionComponentMatches200Response) GetResults() []GetRevisionComponentMatches200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetRevisionComponentMatches200Response) GetResultsOk() (*[]GetRevisionComponentMatches200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetRevisionComponentMatches200Response) SetResults(v []GetRevisionComponentMatches200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetRevisionComponentMatches200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetPage

`func (o *GetRevisionComponentMatches200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetRevisionComponentMatches200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetRevisionComponentMatches200Response) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetRevisionComponentMatches200Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GetRevisionComponentMatches200Response) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetRevisionComponentMatches200Response) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetRevisionComponentMatches200Response) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GetRevisionComponentMatches200Response) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetRevisionComponentMatches200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetRevisionComponentMatches200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetRevisionComponentMatches200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetRevisionComponentMatches200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


