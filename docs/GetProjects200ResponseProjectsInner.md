# GetProjects200ResponseProjectsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Title** | **string** |  | 
**Branch** | **string** |  | 
**Version** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Public** | **bool** |  | 
**OriginOrganizationName** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Scanned** | Pointer to **time.Time** |  | [optional] 
**LastAnalyzed** | Pointer to **time.Time** |  | [optional] 
**Teams** | [**[]GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner**](GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner.md) |  | 
**LatestRevision** | Pointer to [**GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision**](GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision.md) |  | [optional] 
**LatestBuildStatus** | Pointer to **string** |  | [optional] 
**Issues** | [**GetProjects200ResponseProjectsInnerAllOfAllOfIssues**](GetProjects200ResponseProjectsInnerAllOfAllOfIssues.md) |  | 
**Labels** | **[]string** |  | 

## Methods

### NewGetProjects200ResponseProjectsInner

`func NewGetProjects200ResponseProjectsInner(id string, title string, branch string, type_ string, public bool, teams []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, issues GetProjects200ResponseProjectsInnerAllOfAllOfIssues, labels []string, ) *GetProjects200ResponseProjectsInner`

NewGetProjects200ResponseProjectsInner instantiates a new GetProjects200ResponseProjectsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjects200ResponseProjectsInnerWithDefaults

`func NewGetProjects200ResponseProjectsInnerWithDefaults() *GetProjects200ResponseProjectsInner`

NewGetProjects200ResponseProjectsInnerWithDefaults instantiates a new GetProjects200ResponseProjectsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProjects200ResponseProjectsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProjects200ResponseProjectsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProjects200ResponseProjectsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetProjects200ResponseProjectsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetProjects200ResponseProjectsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetProjects200ResponseProjectsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBranch

`func (o *GetProjects200ResponseProjectsInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetProjects200ResponseProjectsInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetProjects200ResponseProjectsInner) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetVersion

`func (o *GetProjects200ResponseProjectsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetProjects200ResponseProjectsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetProjects200ResponseProjectsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetProjects200ResponseProjectsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *GetProjects200ResponseProjectsInner) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *GetProjects200ResponseProjectsInner) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *GetProjects200ResponseProjectsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProjects200ResponseProjectsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProjects200ResponseProjectsInner) SetType(v string)`

SetType sets Type field to given value.


### GetPublic

`func (o *GetProjects200ResponseProjectsInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetProjects200ResponseProjectsInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetProjects200ResponseProjectsInner) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetOriginOrganizationName

`func (o *GetProjects200ResponseProjectsInner) GetOriginOrganizationName() string`

GetOriginOrganizationName returns the OriginOrganizationName field if non-nil, zero value otherwise.

### GetOriginOrganizationNameOk

`func (o *GetProjects200ResponseProjectsInner) GetOriginOrganizationNameOk() (*string, bool)`

GetOriginOrganizationNameOk returns a tuple with the OriginOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOrganizationName

`func (o *GetProjects200ResponseProjectsInner) SetOriginOrganizationName(v string)`

SetOriginOrganizationName sets OriginOrganizationName field to given value.

### HasOriginOrganizationName

`func (o *GetProjects200ResponseProjectsInner) HasOriginOrganizationName() bool`

HasOriginOrganizationName returns a boolean if a field has been set.

### SetOriginOrganizationNameNil

`func (o *GetProjects200ResponseProjectsInner) SetOriginOrganizationNameNil(b bool)`

 SetOriginOrganizationNameNil sets the value for OriginOrganizationName to be an explicit nil

### UnsetOriginOrganizationName
`func (o *GetProjects200ResponseProjectsInner) UnsetOriginOrganizationName()`

UnsetOriginOrganizationName ensures that no value is present for OriginOrganizationName, not even an explicit nil
### GetUrl

`func (o *GetProjects200ResponseProjectsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetProjects200ResponseProjectsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetProjects200ResponseProjectsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetProjects200ResponseProjectsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScanned

`func (o *GetProjects200ResponseProjectsInner) GetScanned() time.Time`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *GetProjects200ResponseProjectsInner) GetScannedOk() (*time.Time, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *GetProjects200ResponseProjectsInner) SetScanned(v time.Time)`

SetScanned sets Scanned field to given value.

### HasScanned

`func (o *GetProjects200ResponseProjectsInner) HasScanned() bool`

HasScanned returns a boolean if a field has been set.

### GetLastAnalyzed

`func (o *GetProjects200ResponseProjectsInner) GetLastAnalyzed() time.Time`

GetLastAnalyzed returns the LastAnalyzed field if non-nil, zero value otherwise.

### GetLastAnalyzedOk

`func (o *GetProjects200ResponseProjectsInner) GetLastAnalyzedOk() (*time.Time, bool)`

GetLastAnalyzedOk returns a tuple with the LastAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalyzed

`func (o *GetProjects200ResponseProjectsInner) SetLastAnalyzed(v time.Time)`

SetLastAnalyzed sets LastAnalyzed field to given value.

### HasLastAnalyzed

`func (o *GetProjects200ResponseProjectsInner) HasLastAnalyzed() bool`

HasLastAnalyzed returns a boolean if a field has been set.

### GetTeams

`func (o *GetProjects200ResponseProjectsInner) GetTeams() []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetProjects200ResponseProjectsInner) GetTeamsOk() (*[]GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetProjects200ResponseProjectsInner) SetTeams(v []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner)`

SetTeams sets Teams field to given value.


### GetLatestRevision

`func (o *GetProjects200ResponseProjectsInner) GetLatestRevision() GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *GetProjects200ResponseProjectsInner) GetLatestRevisionOk() (*GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *GetProjects200ResponseProjectsInner) SetLatestRevision(v GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *GetProjects200ResponseProjectsInner) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetLatestBuildStatus

`func (o *GetProjects200ResponseProjectsInner) GetLatestBuildStatus() string`

GetLatestBuildStatus returns the LatestBuildStatus field if non-nil, zero value otherwise.

### GetLatestBuildStatusOk

`func (o *GetProjects200ResponseProjectsInner) GetLatestBuildStatusOk() (*string, bool)`

GetLatestBuildStatusOk returns a tuple with the LatestBuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBuildStatus

`func (o *GetProjects200ResponseProjectsInner) SetLatestBuildStatus(v string)`

SetLatestBuildStatus sets LatestBuildStatus field to given value.

### HasLatestBuildStatus

`func (o *GetProjects200ResponseProjectsInner) HasLatestBuildStatus() bool`

HasLatestBuildStatus returns a boolean if a field has been set.

### GetIssues

`func (o *GetProjects200ResponseProjectsInner) GetIssues() GetProjects200ResponseProjectsInnerAllOfAllOfIssues`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetProjects200ResponseProjectsInner) GetIssuesOk() (*GetProjects200ResponseProjectsInnerAllOfAllOfIssues, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetProjects200ResponseProjectsInner) SetIssues(v GetProjects200ResponseProjectsInnerAllOfAllOfIssues)`

SetIssues sets Issues field to given value.


### GetLabels

`func (o *GetProjects200ResponseProjectsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetProjects200ResponseProjectsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetProjects200ResponseProjectsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


