# GetIssues202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to [**GetIssues200ResponseIssues**](GetIssues200ResponseIssues.md) |  | [optional] 

## Methods

### NewGetIssues202Response

`func NewGetIssues202Response() *GetIssues202Response`

NewGetIssues202Response instantiates a new GetIssues202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssues202ResponseWithDefaults

`func NewGetIssues202ResponseWithDefaults() *GetIssues202Response`

NewGetIssues202ResponseWithDefaults instantiates a new GetIssues202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetIssues202Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetIssues202Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetIssues202Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetIssues202Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIssues

`func (o *GetIssues202Response) GetIssues() GetIssues200ResponseIssues`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetIssues202Response) GetIssuesOk() (*GetIssues200ResponseIssues, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetIssues202Response) SetIssues(v GetIssues200ResponseIssues)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetIssues202Response) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


