# GetSnippetPackages200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the snippet package | 
**Name** | **string** | Name of the package | 
**VersionCount** | **int32** | Number of versions detected for this package | 
**MatchCount** | **int32** | Total number of matches for this package across all versions | 
**IssueCounts** | [**GetSnippetPackages200ResponseResultsInnerIssueCounts**](GetSnippetPackages200ResponseResultsInnerIssueCounts.md) |  | 
**Labels** | [**[]GetSnippets200ResponseResultsInnerLabelsInner**](GetSnippets200ResponseResultsInnerLabelsInner.md) | Package labels assigned to this snippet package | 
**IsFullyRejected** | **bool** | Whether all snippets in this package have been rejected | 

## Methods

### NewGetSnippetPackages200ResponseResultsInner

`func NewGetSnippetPackages200ResponseResultsInner(id string, name string, versionCount int32, matchCount int32, issueCounts GetSnippetPackages200ResponseResultsInnerIssueCounts, labels []GetSnippets200ResponseResultsInnerLabelsInner, isFullyRejected bool, ) *GetSnippetPackages200ResponseResultsInner`

NewGetSnippetPackages200ResponseResultsInner instantiates a new GetSnippetPackages200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetPackages200ResponseResultsInnerWithDefaults

`func NewGetSnippetPackages200ResponseResultsInnerWithDefaults() *GetSnippetPackages200ResponseResultsInner`

NewGetSnippetPackages200ResponseResultsInnerWithDefaults instantiates a new GetSnippetPackages200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSnippetPackages200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSnippetPackages200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetSnippetPackages200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSnippetPackages200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetVersionCount

`func (o *GetSnippetPackages200ResponseResultsInner) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *GetSnippetPackages200ResponseResultsInner) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.


### GetMatchCount

`func (o *GetSnippetPackages200ResponseResultsInner) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *GetSnippetPackages200ResponseResultsInner) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.


### GetIssueCounts

`func (o *GetSnippetPackages200ResponseResultsInner) GetIssueCounts() GetSnippetPackages200ResponseResultsInnerIssueCounts`

GetIssueCounts returns the IssueCounts field if non-nil, zero value otherwise.

### GetIssueCountsOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetIssueCountsOk() (*GetSnippetPackages200ResponseResultsInnerIssueCounts, bool)`

GetIssueCountsOk returns a tuple with the IssueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCounts

`func (o *GetSnippetPackages200ResponseResultsInner) SetIssueCounts(v GetSnippetPackages200ResponseResultsInnerIssueCounts)`

SetIssueCounts sets IssueCounts field to given value.


### GetLabels

`func (o *GetSnippetPackages200ResponseResultsInner) GetLabels() []GetSnippets200ResponseResultsInnerLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetLabelsOk() (*[]GetSnippets200ResponseResultsInnerLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetSnippetPackages200ResponseResultsInner) SetLabels(v []GetSnippets200ResponseResultsInnerLabelsInner)`

SetLabels sets Labels field to given value.


### GetIsFullyRejected

`func (o *GetSnippetPackages200ResponseResultsInner) GetIsFullyRejected() bool`

GetIsFullyRejected returns the IsFullyRejected field if non-nil, zero value otherwise.

### GetIsFullyRejectedOk

`func (o *GetSnippetPackages200ResponseResultsInner) GetIsFullyRejectedOk() (*bool, bool)`

GetIsFullyRejectedOk returns a tuple with the IsFullyRejected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullyRejected

`func (o *GetSnippetPackages200ResponseResultsInner) SetIsFullyRejected(v bool)`

SetIsFullyRejected sets IsFullyRejected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


