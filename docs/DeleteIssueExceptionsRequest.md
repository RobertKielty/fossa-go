# DeleteIssueExceptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**DeleteIssueExceptionsRequestOneOfFilters**](DeleteIssueExceptionsRequestOneOfFilters.md) |  | [optional] 
**Value** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewDeleteIssueExceptionsRequest

`func NewDeleteIssueExceptionsRequest() *DeleteIssueExceptionsRequest`

NewDeleteIssueExceptionsRequest instantiates a new DeleteIssueExceptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIssueExceptionsRequestWithDefaults

`func NewDeleteIssueExceptionsRequestWithDefaults() *DeleteIssueExceptionsRequest`

NewDeleteIssueExceptionsRequestWithDefaults instantiates a new DeleteIssueExceptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DeleteIssueExceptionsRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DeleteIssueExceptionsRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DeleteIssueExceptionsRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DeleteIssueExceptionsRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetFilters

`func (o *DeleteIssueExceptionsRequest) GetFilters() DeleteIssueExceptionsRequestOneOfFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DeleteIssueExceptionsRequest) GetFiltersOk() (*DeleteIssueExceptionsRequestOneOfFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DeleteIssueExceptionsRequest) SetFilters(v DeleteIssueExceptionsRequestOneOfFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DeleteIssueExceptionsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetValue

`func (o *DeleteIssueExceptionsRequest) GetValue() []int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeleteIssueExceptionsRequest) GetValueOk() (*[]int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeleteIssueExceptionsRequest) SetValue(v []int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *DeleteIssueExceptionsRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


