# GetProjectRevisions200ResponseBranchValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | Pointer to [**GetProjectRevisions200ResponseBranchValueInnerLoc**](GetProjectRevisions200ResponseBranchValueInnerLoc.md) |  | [optional] 
**Locator** | Pointer to **string** | Text ID that uniquely identifies a project | [optional] 
**Resolved** | Pointer to **bool** | Has the Revision been fully analyzed by FOSSA | [optional] 
**ProjectId** | Pointer to **string** | The Project locator that the Revision belongs to | [optional] 
**SourceType** | Pointer to **NullableString** | FOSSA internal representation of the source language for the given repository/project | [optional] 
**Error** | Pointer to **NullableString** | Error message during analysis (if any) | [optional] 
**Message** | Pointer to **NullableString** | Message of the revision or commit | [optional] 
**RevisionTimestamp** | Pointer to **NullableString** | timestamp of when the Revision was published | [optional] 
**LatestRevisionScanId** | Pointer to **NullableFloat32** | The Revision Scan ID of the latest policy scan | [optional] 
**LatestHubbleAnalysisId** | Pointer to **NullableFloat32** | The Hubble Analysis ID of the latest analysis | [optional] 
**CreatedAt** | Pointer to **string** | when the Revision was added to the FOSSA Database | [optional] 
**UpdatedAt** | Pointer to **string** | when the Revision was last updated in the FOSSA Database | [optional] 
**Author** | Pointer to **NullableString** | The author of the Revision | [optional] 
**Link** | Pointer to **NullableString** | The link associated with the Revision | [optional] 
**Url** | Pointer to **NullableString** | The url associated with the Revision | [optional] 

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

### GetAuthor

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetLink

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetUrl

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetProjectRevisions200ResponseBranchValueInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetProjectRevisions200ResponseBranchValueInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GetProjectRevisions200ResponseBranchValueInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GetProjectRevisions200ResponseBranchValueInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


