# GetIssuesByRevision200ResponseRevisionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionId** | Pointer to **string** |  | [optional] 
**DependencyName** | Pointer to **string** |  | [optional] 
**IssueCount** | Pointer to **int32** |  | [optional] 
**ProjectCount** | Pointer to **int32** |  | [optional] 
**JiraTicketCount** | Pointer to **int32** |  | [optional] 
**Depth** | Pointer to [**GetIssuesByRevision200ResponseRevisionsInnerDepth**](GetIssuesByRevision200ResponseRevisionsInnerDepth.md) |  | [optional] 
**Type** | Pointer to [**GetIssuesByRevision200ResponseRevisionsInnerType**](GetIssuesByRevision200ResponseRevisionsInnerType.md) |  | [optional] 

## Methods

### NewGetIssuesByRevision200ResponseRevisionsInner

`func NewGetIssuesByRevision200ResponseRevisionsInner() *GetIssuesByRevision200ResponseRevisionsInner`

NewGetIssuesByRevision200ResponseRevisionsInner instantiates a new GetIssuesByRevision200ResponseRevisionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssuesByRevision200ResponseRevisionsInnerWithDefaults

`func NewGetIssuesByRevision200ResponseRevisionsInnerWithDefaults() *GetIssuesByRevision200ResponseRevisionsInner`

NewGetIssuesByRevision200ResponseRevisionsInnerWithDefaults instantiates a new GetIssuesByRevision200ResponseRevisionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionId

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetDependencyName

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetDependencyName() string`

GetDependencyName returns the DependencyName field if non-nil, zero value otherwise.

### GetDependencyNameOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetDependencyNameOk() (*string, bool)`

GetDependencyNameOk returns a tuple with the DependencyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyName

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetDependencyName(v string)`

SetDependencyName sets DependencyName field to given value.

### HasDependencyName

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasDependencyName() bool`

HasDependencyName returns a boolean if a field has been set.

### GetIssueCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetIssueCount() int32`

GetIssueCount returns the IssueCount field if non-nil, zero value otherwise.

### GetIssueCountOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetIssueCountOk() (*int32, bool)`

GetIssueCountOk returns a tuple with the IssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetIssueCount(v int32)`

SetIssueCount sets IssueCount field to given value.

### HasIssueCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasIssueCount() bool`

HasIssueCount returns a boolean if a field has been set.

### GetProjectCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetProjectCount() int32`

GetProjectCount returns the ProjectCount field if non-nil, zero value otherwise.

### GetProjectCountOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetProjectCountOk() (*int32, bool)`

GetProjectCountOk returns a tuple with the ProjectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetProjectCount(v int32)`

SetProjectCount sets ProjectCount field to given value.

### HasProjectCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasProjectCount() bool`

HasProjectCount returns a boolean if a field has been set.

### GetJiraTicketCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetJiraTicketCount() int32`

GetJiraTicketCount returns the JiraTicketCount field if non-nil, zero value otherwise.

### GetJiraTicketCountOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetJiraTicketCountOk() (*int32, bool)`

GetJiraTicketCountOk returns a tuple with the JiraTicketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraTicketCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetJiraTicketCount(v int32)`

SetJiraTicketCount sets JiraTicketCount field to given value.

### HasJiraTicketCount

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasJiraTicketCount() bool`

HasJiraTicketCount returns a boolean if a field has been set.

### GetDepth

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetDepth() GetIssuesByRevision200ResponseRevisionsInnerDepth`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetDepthOk() (*GetIssuesByRevision200ResponseRevisionsInnerDepth, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetDepth(v GetIssuesByRevision200ResponseRevisionsInnerDepth)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetType

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetType() GetIssuesByRevision200ResponseRevisionsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssuesByRevision200ResponseRevisionsInner) GetTypeOk() (*GetIssuesByRevision200ResponseRevisionsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssuesByRevision200ResponseRevisionsInner) SetType(v GetIssuesByRevision200ResponseRevisionsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssuesByRevision200ResponseRevisionsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


