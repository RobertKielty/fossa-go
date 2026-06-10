# GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner

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

## Methods

### NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner

`func NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner(id string, title string, branch string, type_ string, public bool, teams []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, ) *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner`

NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner instantiates a new GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInnerWithDefaults

`func NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInnerWithDefaults() *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner`

NewGetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInnerWithDefaults instantiates a new GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBranch

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetVersion

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetType(v string)`

SetType sets Type field to given value.


### GetPublic

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetOriginOrganizationName() string`

GetOriginOrganizationName returns the OriginOrganizationName field if non-nil, zero value otherwise.

### GetOriginOrganizationNameOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetOriginOrganizationNameOk() (*string, bool)`

GetOriginOrganizationNameOk returns a tuple with the OriginOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetOriginOrganizationName(v string)`

SetOriginOrganizationName sets OriginOrganizationName field to given value.

### HasOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasOriginOrganizationName() bool`

HasOriginOrganizationName returns a boolean if a field has been set.

### SetOriginOrganizationNameNil

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetOriginOrganizationNameNil(b bool)`

 SetOriginOrganizationNameNil sets the value for OriginOrganizationName to be an explicit nil

### UnsetOriginOrganizationName
`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) UnsetOriginOrganizationName()`

UnsetOriginOrganizationName ensures that no value is present for OriginOrganizationName, not even an explicit nil
### GetUrl

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScanned

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetScanned() time.Time`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetScannedOk() (*time.Time, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetScanned(v time.Time)`

SetScanned sets Scanned field to given value.

### HasScanned

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasScanned() bool`

HasScanned returns a boolean if a field has been set.

### GetLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLastAnalyzed() time.Time`

GetLastAnalyzed returns the LastAnalyzed field if non-nil, zero value otherwise.

### GetLastAnalyzedOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLastAnalyzedOk() (*time.Time, bool)`

GetLastAnalyzedOk returns a tuple with the LastAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetLastAnalyzed(v time.Time)`

SetLastAnalyzed sets LastAnalyzed field to given value.

### HasLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasLastAnalyzed() bool`

HasLastAnalyzed returns a boolean if a field has been set.

### GetTeams

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetTeams() []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetTeamsOk() (*[]GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetTeams(v []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner)`

SetTeams sets Teams field to given value.


### GetLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLatestRevision() GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLatestRevisionOk() (*GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetLatestRevision(v GetProjects200ResponseProjectsInnerAllOfAllOfLatestRevision)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLatestBuildStatus() string`

GetLatestBuildStatus returns the LatestBuildStatus field if non-nil, zero value otherwise.

### GetLatestBuildStatusOk

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) GetLatestBuildStatusOk() (*string, bool)`

GetLatestBuildStatusOk returns a tuple with the LatestBuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) SetLatestBuildStatus(v string)`

SetLatestBuildStatus sets LatestBuildStatus field to given value.

### HasLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseInnerAllOfProjectsInnerAllOfProjectInner) HasLatestBuildStatus() bool`

HasLatestBuildStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


