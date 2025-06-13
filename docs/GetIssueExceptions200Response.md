# GetIssueExceptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exceptions** | Pointer to [**[]GetIssueExceptions200ResponseExceptionsInner**](GetIssueExceptions200ResponseExceptionsInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetIssueExceptions200Response

`func NewGetIssueExceptions200Response() *GetIssueExceptions200Response`

NewGetIssueExceptions200Response instantiates a new GetIssueExceptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueExceptions200ResponseWithDefaults

`func NewGetIssueExceptions200ResponseWithDefaults() *GetIssueExceptions200Response`

NewGetIssueExceptions200ResponseWithDefaults instantiates a new GetIssueExceptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptions

`func (o *GetIssueExceptions200Response) GetExceptions() []GetIssueExceptions200ResponseExceptionsInner`

GetExceptions returns the Exceptions field if non-nil, zero value otherwise.

### GetExceptionsOk

`func (o *GetIssueExceptions200Response) GetExceptionsOk() (*[]GetIssueExceptions200ResponseExceptionsInner, bool)`

GetExceptionsOk returns a tuple with the Exceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptions

`func (o *GetIssueExceptions200Response) SetExceptions(v []GetIssueExceptions200ResponseExceptionsInner)`

SetExceptions sets Exceptions field to given value.

### HasExceptions

`func (o *GetIssueExceptions200Response) HasExceptions() bool`

HasExceptions returns a boolean if a field has been set.

### GetTotal

`func (o *GetIssueExceptions200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetIssueExceptions200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetIssueExceptions200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetIssueExceptions200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


