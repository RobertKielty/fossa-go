# GetIssue200ResponseOneOfAllOfProjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ScannedAt** | Pointer to **time.Time** |  | [optional] 
**AnalyzedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** | A link to the project the issue is from. | [optional] 
**OriginPaths** | Pointer to **[]string** | Optional array of direct dependency origin paths, only present when includeDirectDependencyOriginPaths&#x3D;true is supplied as a URL param.  | [optional] 
**FirstFoundAt** | Pointer to **time.Time** | The date when the issue was first found for this project or release group. | [optional] 
**DefaultBranch** | Pointer to **string** | The default branch of the project. | [optional] 
**Latest** | Pointer to **bool** | Whether this issue is present in the most recent scan of the project. If false, the issue may be considered remediated. | [optional] 
**RevisionId** | Pointer to **string** | The ID of the revision that is affected by this issue. | [optional] 
**RevisionScanId** | Pointer to **float32** | The ID of the revision scan associated with this issue. | [optional] 

## Methods

### NewGetIssue200ResponseOneOfAllOfProjectsInner

`func NewGetIssue200ResponseOneOfAllOfProjectsInner() *GetIssue200ResponseOneOfAllOfProjectsInner`

NewGetIssue200ResponseOneOfAllOfProjectsInner instantiates a new GetIssue200ResponseOneOfAllOfProjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseOneOfAllOfProjectsInnerWithDefaults

`func NewGetIssue200ResponseOneOfAllOfProjectsInnerWithDefaults() *GetIssue200ResponseOneOfAllOfProjectsInner`

NewGetIssue200ResponseOneOfAllOfProjectsInnerWithDefaults instantiates a new GetIssue200ResponseOneOfAllOfProjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDepth

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetTitle

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetAnalyzedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetAnalyzedAt() time.Time`

GetAnalyzedAt returns the AnalyzedAt field if non-nil, zero value otherwise.

### GetAnalyzedAtOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetAnalyzedAtOk() (*time.Time, bool)`

GetAnalyzedAtOk returns a tuple with the AnalyzedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetAnalyzedAt(v time.Time)`

SetAnalyzedAt sets AnalyzedAt field to given value.

### HasAnalyzedAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasAnalyzedAt() bool`

HasAnalyzedAt returns a boolean if a field has been set.

### GetUrl

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetOriginPaths

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetOriginPaths() []string`

GetOriginPaths returns the OriginPaths field if non-nil, zero value otherwise.

### GetOriginPathsOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetOriginPathsOk() (*[]string, bool)`

GetOriginPathsOk returns a tuple with the OriginPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPaths

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetOriginPaths(v []string)`

SetOriginPaths sets OriginPaths field to given value.

### HasOriginPaths

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasOriginPaths() bool`

HasOriginPaths returns a boolean if a field has been set.

### GetFirstFoundAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetFirstFoundAt() time.Time`

GetFirstFoundAt returns the FirstFoundAt field if non-nil, zero value otherwise.

### GetFirstFoundAtOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetFirstFoundAtOk() (*time.Time, bool)`

GetFirstFoundAtOk returns a tuple with the FirstFoundAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFoundAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetFirstFoundAt(v time.Time)`

SetFirstFoundAt sets FirstFoundAt field to given value.

### HasFirstFoundAt

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasFirstFoundAt() bool`

HasFirstFoundAt returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetLatest

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetRevisionId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetRevisionScanId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetRevisionScanId() float32`

GetRevisionScanId returns the RevisionScanId field if non-nil, zero value otherwise.

### GetRevisionScanIdOk

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) GetRevisionScanIdOk() (*float32, bool)`

GetRevisionScanIdOk returns a tuple with the RevisionScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionScanId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) SetRevisionScanId(v float32)`

SetRevisionScanId sets RevisionScanId field to given value.

### HasRevisionScanId

`func (o *GetIssue200ResponseOneOfAllOfProjectsInner) HasRevisionScanId() bool`

HasRevisionScanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


