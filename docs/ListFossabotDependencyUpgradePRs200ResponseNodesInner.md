# ListFossabotDependencyUpgradePRs200ResponseNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**PrNumber** | **NullableInt32** |  | 
**VcsUrl** | **NullableString** |  | 
**State** | **string** |  | 
**AnalysisBadge** | **NullableString** |  | 
**InitiatedByUserId** | **NullableInt32** | FOSSA user id of the initiator; null for automated PRs. | 
**OpenedAt** | **time.Time** |  | 
**ClosedAt** | **NullableTime** |  | 
**MergedAt** | **NullableTime** |  | 

## Methods

### NewListFossabotDependencyUpgradePRs200ResponseNodesInner

`func NewListFossabotDependencyUpgradePRs200ResponseNodesInner(title string, prNumber NullableInt32, vcsUrl NullableString, state string, analysisBadge NullableString, initiatedByUserId NullableInt32, openedAt time.Time, closedAt NullableTime, mergedAt NullableTime, ) *ListFossabotDependencyUpgradePRs200ResponseNodesInner`

NewListFossabotDependencyUpgradePRs200ResponseNodesInner instantiates a new ListFossabotDependencyUpgradePRs200ResponseNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFossabotDependencyUpgradePRs200ResponseNodesInnerWithDefaults

`func NewListFossabotDependencyUpgradePRs200ResponseNodesInnerWithDefaults() *ListFossabotDependencyUpgradePRs200ResponseNodesInner`

NewListFossabotDependencyUpgradePRs200ResponseNodesInnerWithDefaults instantiates a new ListFossabotDependencyUpgradePRs200ResponseNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrNumber

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetPrNumber() int32`

GetPrNumber returns the PrNumber field if non-nil, zero value otherwise.

### GetPrNumberOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetPrNumberOk() (*int32, bool)`

GetPrNumberOk returns a tuple with the PrNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrNumber

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetPrNumber(v int32)`

SetPrNumber sets PrNumber field to given value.


### SetPrNumberNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetPrNumberNil(b bool)`

 SetPrNumberNil sets the value for PrNumber to be an explicit nil

### UnsetPrNumber
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetPrNumber()`

UnsetPrNumber ensures that no value is present for PrNumber, not even an explicit nil
### GetVcsUrl

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetVcsUrl() string`

GetVcsUrl returns the VcsUrl field if non-nil, zero value otherwise.

### GetVcsUrlOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetVcsUrlOk() (*string, bool)`

GetVcsUrlOk returns a tuple with the VcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUrl

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetVcsUrl(v string)`

SetVcsUrl sets VcsUrl field to given value.


### SetVcsUrlNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetVcsUrlNil(b bool)`

 SetVcsUrlNil sets the value for VcsUrl to be an explicit nil

### UnsetVcsUrl
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetVcsUrl()`

UnsetVcsUrl ensures that no value is present for VcsUrl, not even an explicit nil
### GetState

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetState(v string)`

SetState sets State field to given value.


### GetAnalysisBadge

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetAnalysisBadge() string`

GetAnalysisBadge returns the AnalysisBadge field if non-nil, zero value otherwise.

### GetAnalysisBadgeOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetAnalysisBadgeOk() (*string, bool)`

GetAnalysisBadgeOk returns a tuple with the AnalysisBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisBadge

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetAnalysisBadge(v string)`

SetAnalysisBadge sets AnalysisBadge field to given value.


### SetAnalysisBadgeNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetAnalysisBadgeNil(b bool)`

 SetAnalysisBadgeNil sets the value for AnalysisBadge to be an explicit nil

### UnsetAnalysisBadge
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetAnalysisBadge()`

UnsetAnalysisBadge ensures that no value is present for AnalysisBadge, not even an explicit nil
### GetInitiatedByUserId

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetInitiatedByUserId() int32`

GetInitiatedByUserId returns the InitiatedByUserId field if non-nil, zero value otherwise.

### GetInitiatedByUserIdOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetInitiatedByUserIdOk() (*int32, bool)`

GetInitiatedByUserIdOk returns a tuple with the InitiatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedByUserId

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetInitiatedByUserId(v int32)`

SetInitiatedByUserId sets InitiatedByUserId field to given value.


### SetInitiatedByUserIdNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetInitiatedByUserIdNil(b bool)`

 SetInitiatedByUserIdNil sets the value for InitiatedByUserId to be an explicit nil

### UnsetInitiatedByUserId
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetInitiatedByUserId()`

UnsetInitiatedByUserId ensures that no value is present for InitiatedByUserId, not even an explicit nil
### GetOpenedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetOpenedAt() time.Time`

GetOpenedAt returns the OpenedAt field if non-nil, zero value otherwise.

### GetOpenedAtOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetOpenedAtOk() (*time.Time, bool)`

GetOpenedAtOk returns a tuple with the OpenedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetOpenedAt(v time.Time)`

SetOpenedAt sets OpenedAt field to given value.


### GetClosedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetClosedAt() time.Time`

GetClosedAt returns the ClosedAt field if non-nil, zero value otherwise.

### GetClosedAtOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetClosedAtOk() (*time.Time, bool)`

GetClosedAtOk returns a tuple with the ClosedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetClosedAt(v time.Time)`

SetClosedAt sets ClosedAt field to given value.


### SetClosedAtNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetClosedAtNil(b bool)`

 SetClosedAtNil sets the value for ClosedAt to be an explicit nil

### UnsetClosedAt
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetClosedAt()`

UnsetClosedAt ensures that no value is present for ClosedAt, not even an explicit nil
### GetMergedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.


### SetMergedAtNil

`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *ListFossabotDependencyUpgradePRs200ResponseNodesInner) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


