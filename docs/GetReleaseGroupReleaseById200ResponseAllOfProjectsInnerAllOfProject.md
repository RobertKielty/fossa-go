# GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject

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
**LatestRevision** | Pointer to [**GetProjectRevisions200ResponseBranchValueInner**](GetProjectRevisions200ResponseBranchValueInner.md) |  | [optional] 
**LatestBuildStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject

`func NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject(id string, title string, branch string, type_ string, public bool, teams []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, ) *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject`

NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject instantiates a new GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProjectWithDefaults

`func NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProjectWithDefaults() *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject`

NewGetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProjectWithDefaults instantiates a new GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBranch

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetVersion

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetType(v string)`

SetType sets Type field to given value.


### GetPublic

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetOriginOrganizationName() string`

GetOriginOrganizationName returns the OriginOrganizationName field if non-nil, zero value otherwise.

### GetOriginOrganizationNameOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetOriginOrganizationNameOk() (*string, bool)`

GetOriginOrganizationNameOk returns a tuple with the OriginOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetOriginOrganizationName(v string)`

SetOriginOrganizationName sets OriginOrganizationName field to given value.

### HasOriginOrganizationName

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasOriginOrganizationName() bool`

HasOriginOrganizationName returns a boolean if a field has been set.

### SetOriginOrganizationNameNil

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetOriginOrganizationNameNil(b bool)`

 SetOriginOrganizationNameNil sets the value for OriginOrganizationName to be an explicit nil

### UnsetOriginOrganizationName
`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) UnsetOriginOrganizationName()`

UnsetOriginOrganizationName ensures that no value is present for OriginOrganizationName, not even an explicit nil
### GetUrl

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScanned

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetScanned() time.Time`

GetScanned returns the Scanned field if non-nil, zero value otherwise.

### GetScannedOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetScannedOk() (*time.Time, bool)`

GetScannedOk returns a tuple with the Scanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanned

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetScanned(v time.Time)`

SetScanned sets Scanned field to given value.

### HasScanned

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasScanned() bool`

HasScanned returns a boolean if a field has been set.

### GetLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLastAnalyzed() time.Time`

GetLastAnalyzed returns the LastAnalyzed field if non-nil, zero value otherwise.

### GetLastAnalyzedOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLastAnalyzedOk() (*time.Time, bool)`

GetLastAnalyzedOk returns a tuple with the LastAnalyzed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetLastAnalyzed(v time.Time)`

SetLastAnalyzed sets LastAnalyzed field to given value.

### HasLastAnalyzed

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasLastAnalyzed() bool`

HasLastAnalyzed returns a boolean if a field has been set.

### GetTeams

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetTeams() []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetTeamsOk() (*[]GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetTeams(v []GetProjects200ResponseProjectsInnerAllOfAllOfTeamsInner)`

SetTeams sets Teams field to given value.


### GetLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLatestRevision() GetProjectRevisions200ResponseBranchValueInner`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLatestRevisionOk() (*GetProjectRevisions200ResponseBranchValueInner, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetLatestRevision(v GetProjectRevisions200ResponseBranchValueInner)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLatestBuildStatus() string`

GetLatestBuildStatus returns the LatestBuildStatus field if non-nil, zero value otherwise.

### GetLatestBuildStatusOk

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) GetLatestBuildStatusOk() (*string, bool)`

GetLatestBuildStatusOk returns a tuple with the LatestBuildStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) SetLatestBuildStatus(v string)`

SetLatestBuildStatus sets LatestBuildStatus field to given value.

### HasLatestBuildStatus

`func (o *GetReleaseGroupReleaseById200ResponseAllOfProjectsInnerAllOfProject) HasLatestBuildStatus() bool`

HasLatestBuildStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


