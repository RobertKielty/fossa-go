# CreateIssueDisputeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | The reason why this issue is being disputed. | [optional] 
**Comment** | Pointer to **string** | Any additional information that is important for this dispute. | [optional] 

## Methods

### NewCreateIssueDisputeRequest

`func NewCreateIssueDisputeRequest() *CreateIssueDisputeRequest`

NewCreateIssueDisputeRequest instantiates a new CreateIssueDisputeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueDisputeRequestWithDefaults

`func NewCreateIssueDisputeRequestWithDefaults() *CreateIssueDisputeRequest`

NewCreateIssueDisputeRequestWithDefaults instantiates a new CreateIssueDisputeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *CreateIssueDisputeRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateIssueDisputeRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateIssueDisputeRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateIssueDisputeRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *CreateIssueDisputeRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CreateIssueDisputeRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CreateIssueDisputeRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CreateIssueDisputeRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


