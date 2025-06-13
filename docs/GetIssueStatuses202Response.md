# GetIssueStatuses202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to [**GetIssueStatuses200ResponseIssues**](GetIssueStatuses200ResponseIssues.md) |  | [optional] 
**Revisions** | Pointer to [**GetIssueStatuses200ResponseIssues**](GetIssueStatuses200ResponseIssues.md) |  | [optional] 

## Methods

### NewGetIssueStatuses202Response

`func NewGetIssueStatuses202Response() *GetIssueStatuses202Response`

NewGetIssueStatuses202Response instantiates a new GetIssueStatuses202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueStatuses202ResponseWithDefaults

`func NewGetIssueStatuses202ResponseWithDefaults() *GetIssueStatuses202Response`

NewGetIssueStatuses202ResponseWithDefaults instantiates a new GetIssueStatuses202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetIssueStatuses202Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetIssueStatuses202Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetIssueStatuses202Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetIssueStatuses202Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetIssues

`func (o *GetIssueStatuses202Response) GetIssues() GetIssueStatuses200ResponseIssues`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetIssueStatuses202Response) GetIssuesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetIssueStatuses202Response) SetIssues(v GetIssueStatuses200ResponseIssues)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetIssueStatuses202Response) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRevisions

`func (o *GetIssueStatuses202Response) GetRevisions() GetIssueStatuses200ResponseIssues`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *GetIssueStatuses202Response) GetRevisionsOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *GetIssueStatuses202Response) SetRevisions(v GetIssueStatuses200ResponseIssues)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *GetIssueStatuses202Response) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


