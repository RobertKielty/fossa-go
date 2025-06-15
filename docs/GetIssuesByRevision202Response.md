# GetIssuesByRevision202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Revisions** | Pointer to [**[]GetIssuesByRevision200ResponseRevisionsInner**](GetIssuesByRevision200ResponseRevisionsInner.md) |  | [optional] 

## Methods

### NewGetIssuesByRevision202Response

`func NewGetIssuesByRevision202Response() *GetIssuesByRevision202Response`

NewGetIssuesByRevision202Response instantiates a new GetIssuesByRevision202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssuesByRevision202ResponseWithDefaults

`func NewGetIssuesByRevision202ResponseWithDefaults() *GetIssuesByRevision202Response`

NewGetIssuesByRevision202ResponseWithDefaults instantiates a new GetIssuesByRevision202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *GetIssuesByRevision202Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetIssuesByRevision202Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetIssuesByRevision202Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetIssuesByRevision202Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRevisions

`func (o *GetIssuesByRevision202Response) GetRevisions() []GetIssuesByRevision200ResponseRevisionsInner`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *GetIssuesByRevision202Response) GetRevisionsOk() (*[]GetIssuesByRevision200ResponseRevisionsInner, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *GetIssuesByRevision202Response) SetRevisions(v []GetIssuesByRevision200ResponseRevisionsInner)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *GetIssuesByRevision202Response) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


