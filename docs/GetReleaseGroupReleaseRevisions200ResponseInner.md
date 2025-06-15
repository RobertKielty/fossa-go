# GetReleaseGroupReleaseRevisions200ResponseInner

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
**Licenses** | Pointer to [**[]GetReleaseGroupReleaseRevisions200ResponseInnerAllOfLicensesInner**](GetReleaseGroupReleaseRevisions200ResponseInnerAllOfLicensesInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseRevisions200ResponseInner

`func NewGetReleaseGroupReleaseRevisions200ResponseInner() *GetReleaseGroupReleaseRevisions200ResponseInner`

NewGetReleaseGroupReleaseRevisions200ResponseInner instantiates a new GetReleaseGroupReleaseRevisions200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseRevisions200ResponseInnerWithDefaults

`func NewGetReleaseGroupReleaseRevisions200ResponseInnerWithDefaults() *GetReleaseGroupReleaseRevisions200ResponseInner`

NewGetReleaseGroupReleaseRevisions200ResponseInnerWithDefaults instantiates a new GetReleaseGroupReleaseRevisions200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLoc() GetProjectRevisions200ResponseBranchValueInnerLoc`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLocOk() (*GetProjectRevisions200ResponseBranchValueInnerLoc, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLoc(v GetProjectRevisions200ResponseBranchValueInnerLoc)`

SetLoc sets Loc field to given value.

### HasLoc

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLoc() bool`

HasLoc returns a boolean if a field has been set.

### GetLocator

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetResolved

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetProjectId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetSourceType

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetError

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetMessage

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRevisionTimestamp

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetRevisionTimestamp() string`

GetRevisionTimestamp returns the RevisionTimestamp field if non-nil, zero value otherwise.

### GetRevisionTimestampOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetRevisionTimestampOk() (*string, bool)`

GetRevisionTimestampOk returns a tuple with the RevisionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionTimestamp

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetRevisionTimestamp(v string)`

SetRevisionTimestamp sets RevisionTimestamp field to given value.

### HasRevisionTimestamp

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasRevisionTimestamp() bool`

HasRevisionTimestamp returns a boolean if a field has been set.

### SetRevisionTimestampNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetRevisionTimestampNil(b bool)`

 SetRevisionTimestampNil sets the value for RevisionTimestamp to be an explicit nil

### UnsetRevisionTimestamp
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetRevisionTimestamp()`

UnsetRevisionTimestamp ensures that no value is present for RevisionTimestamp, not even an explicit nil
### GetLatestRevisionScanId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLatestRevisionScanId() float32`

GetLatestRevisionScanId returns the LatestRevisionScanId field if non-nil, zero value otherwise.

### GetLatestRevisionScanIdOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLatestRevisionScanIdOk() (*float32, bool)`

GetLatestRevisionScanIdOk returns a tuple with the LatestRevisionScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionScanId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLatestRevisionScanId(v float32)`

SetLatestRevisionScanId sets LatestRevisionScanId field to given value.

### HasLatestRevisionScanId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLatestRevisionScanId() bool`

HasLatestRevisionScanId returns a boolean if a field has been set.

### SetLatestRevisionScanIdNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLatestRevisionScanIdNil(b bool)`

 SetLatestRevisionScanIdNil sets the value for LatestRevisionScanId to be an explicit nil

### UnsetLatestRevisionScanId
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetLatestRevisionScanId()`

UnsetLatestRevisionScanId ensures that no value is present for LatestRevisionScanId, not even an explicit nil
### GetLatestHubbleAnalysisId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLatestHubbleAnalysisId() float32`

GetLatestHubbleAnalysisId returns the LatestHubbleAnalysisId field if non-nil, zero value otherwise.

### GetLatestHubbleAnalysisIdOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLatestHubbleAnalysisIdOk() (*float32, bool)`

GetLatestHubbleAnalysisIdOk returns a tuple with the LatestHubbleAnalysisId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestHubbleAnalysisId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLatestHubbleAnalysisId(v float32)`

SetLatestHubbleAnalysisId sets LatestHubbleAnalysisId field to given value.

### HasLatestHubbleAnalysisId

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLatestHubbleAnalysisId() bool`

HasLatestHubbleAnalysisId returns a boolean if a field has been set.

### SetLatestHubbleAnalysisIdNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLatestHubbleAnalysisIdNil(b bool)`

 SetLatestHubbleAnalysisIdNil sets the value for LatestHubbleAnalysisId to be an explicit nil

### UnsetLatestHubbleAnalysisId
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetLatestHubbleAnalysisId()`

UnsetLatestHubbleAnalysisId ensures that no value is present for LatestHubbleAnalysisId, not even an explicit nil
### GetCreatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAuthor

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetLink

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetUrl

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetLicenses

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLicenses() []GetReleaseGroupReleaseRevisions200ResponseInnerAllOfLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) GetLicensesOk() (*[]GetReleaseGroupReleaseRevisions200ResponseInnerAllOfLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) SetLicenses(v []GetReleaseGroupReleaseRevisions200ResponseInnerAllOfLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetReleaseGroupReleaseRevisions200ResponseInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


