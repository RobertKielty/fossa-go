# GetProjectRevisions200ResponseBranchValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | Pointer to [**GetProjectRevisions200ResponseBranchValueInnerLoc**](GetProjectRevisions200ResponseBranchValueInnerLoc.md) |  | [optional] 
**Locator** | Pointer to **string** | Text ID that uniquely identifies a project | [optional] 
**Resolved** | Pointer to **bool** | Has the Revision been fully analyzed by FOSSA | [optional] 
**ProjectId** | Pointer to **string** | The Project locator that the Revision belongs to | [optional] 
**SourceType** | Pointer to **NullableString** | FOSSA internal representation of the source language for the given repository/project | [optional] 
**Source** | Pointer to **NullableString** | The source the Revision originated from (for example &#x60;github&#x60;, &#x60;cli&#x60;, &#x60;archive&#x60;, &#x60;container&#x60;, &#x60;sbom&#x60;, or &#x60;binary&#x60;) | [optional] 
**UnresolvedIssueCount** | Pointer to **NullableFloat32** | The number of unresolved issues found for this Revision | [optional] 
**Error** | Pointer to **NullableString** | Error message during analysis (if any) | [optional] 
**Message** | Pointer to **NullableString** | Message of the revision or commit | [optional] 
**RevisionTimestamp** | Pointer to **NullableString** | timestamp of when the Revision was published | [optional] 
**LatestRevisionScanId** | Pointer to **NullableFloat32** | The Revision Scan ID of the latest policy scan | [optional] 
**LatestHubbleAnalysisId** | Pointer to **NullableFloat32** | The Hubble Analysis ID of the latest analysis | [optional] 
**CreatedAt** | Pointer to **string** | when the Revision was added to the FOSSA Database | [optional] 
**UpdatedAt** | Pointer to **string** | when the Revision was last updated in the FOSSA Database | [optional] 

## Methods

### NewGetProjectRevisions200ResponseBranchValueInner

`func NewGetProjectRevisions200ResponseBranchValueInner() *GetProjectRevisions200ResponseBranchValueInner`

NewGetProjectRevisions200ResponseBranchValueInner instantiates a new GetProjectRevisions200ResponseBranchValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectRevisions200ResponseBranchValueInnerWithDefaults

`func NewGetProjectRevisions200ResponseBranchValueInnerWithDefaults() *GetProjectRevisions200ResponseBranchValueInner`

NewGetProjectRevisions200ResponseBranchValueInnerWithDefaults instantiates a new GetProjectRevisions200ResponseBranchValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLoc() GetProjectRevisions200ResponseBranchValueInnerLoc`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLocOk() (*GetProjectRevisions200ResponseBranchValueInnerLoc, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLoc(v GetProjectRevisions200ResponseBranchValueInnerLoc)`

SetLoc sets Loc field to given value.

### HasLoc

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasLoc() bool`

HasLoc returns a boolean if a field has been set.

### GetLocator

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetResolved

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetProjectId

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSourceType

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetSource

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetUnresolvedIssueCount

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUnresolvedIssueCount() float32`

GetUnresolvedIssueCount returns the UnresolvedIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedIssueCountOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUnresolvedIssueCountOk() (*float32, bool)`

GetUnresolvedIssueCountOk returns a tuple with the UnresolvedIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedIssueCount

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetUnresolvedIssueCount(v float32)`

SetUnresolvedIssueCount sets UnresolvedIssueCount field to given value.

### HasUnresolvedIssueCount

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasUnresolvedIssueCount() bool`

HasUnresolvedIssueCount returns a boolean if a field has been set.

### SetUnresolvedIssueCountNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetUnresolvedIssueCountNil(b bool)`

 SetUnresolvedIssueCountNil sets the value for UnresolvedIssueCount to be an explicit nil

### UnsetUnresolvedIssueCount
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetUnresolvedIssueCount()`

UnsetUnresolvedIssueCount ensures that no value is present for UnresolvedIssueCount, not even an explicit nil
### GetError

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMessage

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRevisionTimestamp

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetRevisionTimestamp() string`

GetRevisionTimestamp returns the RevisionTimestamp field if non-nil, zero value otherwise.

### GetRevisionTimestampOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetRevisionTimestampOk() (*string, bool)`

GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionTimestamp

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetRevisionTimestamp(v string)`

SetRevisionTimestamp sets RevisionTimestamp field to given value.

### HasRevisionTimestamp

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasRevisionTimestamp() bool`

HasRevisionTimestamp returns a boolean if a field has been set.

### SetRevisionTimestampNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetRevisionTimestampNil(b bool)`

 SetRevisionTimestampNil sets the value for RevisionTimestamp to be an explicit nil

### UnsetRevisionTimestamp
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetRevisionTimestamp()`

UnsetRevisionTimestamp ensures that no value is present for RevisionTimestamp, not even an explicit nil
### GetLatestRevisionScanId

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLatestRevisionScanId() float32`

GetLatestRevisionScanId returns the LatestRevisionScanId field if non-nil, zero value otherwise.

### GetLatestRevisionScanIdOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLatestRevisionScanIdOk() (*float32, bool)`

GetLatestRevisionScanIdOk returns a tuple with the LatestRevisionScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionScanId

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLatestRevisionScanId(v float32)`

SetLatestRevisionScanId sets LatestRevisionScanId field to given value.

### HasLatestRevisionScanId

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasLatestRevisionScanId() bool`

HasLatestRevisionScanId returns a boolean if a field has been set.

### SetLatestRevisionScanIdNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLatestRevisionScanIdNil(b bool)`

 SetLatestRevisionScanIdNil sets the value for LatestRevisionScanId to be an explicit nil

### UnsetLatestRevisionScanId
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetLatestRevisionScanId()`

UnsetLatestRevisionScanId ensures that no value is present for LatestRevisionScanId, not even an explicit nil
### GetLatestHubbleAnalysisId

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLatestHubbleAnalysisId() float32`

GetLatestHubbleAnalysisId returns the LatestHubbleAnalysisId field if non-nil, zero value otherwise.

### GetLatestHubbleAnalysisIdOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLatestHubbleAnalysisIdOk() (*float32, bool)`

GetLatestHubbleAnalysisIdOk returns a tuple with the LatestHubbleAnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHubbleAnalysisId

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLatestHubbleAnalysisId(v float32)`

SetLatestHubbleAnalysisId sets LatestHubbleAnalysisId field to given value.

### HasLatestHubbleAnalysisId

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasLatestHubbleAnalysisId() bool`

HasLatestHubbleAnalysisId returns a boolean if a field has been set.

### SetLatestHubbleAnalysisIdNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLatestHubbleAnalysisIdNil(b bool)`

 SetLatestHubbleAnalysisIdNil sets the value for LatestHubbleAnalysisId to be an explicit nil

### UnsetLatestHubbleAnalysisId
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetLatestHubbleAnalysisId()`

UnsetLatestHubbleAnalysisId ensures that no value is present for LatestHubbleAnalysisId, not even an explicit nil
### GetCreatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


