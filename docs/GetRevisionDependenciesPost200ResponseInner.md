# GetRevisionDependenciesPost200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Loc** | Pointer to [**GetRevisionDependenciesPost200ResponseInnerLoc**](GetRevisionDependenciesPost200ResponseInnerLoc.md) |  | [optional] 
**Locator** | Pointer to **string** | The full locator string for this dependency | [optional] 
**ProjectId** | Pointer to **string** | The package identifier (fetcher+package) | [optional] 
**Resolved** | Pointer to **bool** | Whether the dependency has been successfully resolved | [optional] 
**Unsupported** | Pointer to **bool** | Whether the dependency fetcher/type is unsupported | [optional] 
**LatestRevisionScanId** | Pointer to **int32** | ID of the latest revision scan (always null for dependencies) | [optional] 
**DependencyLock** | Pointer to [**GetRevisionDependenciesPost200ResponseInnerDependencyLock**](GetRevisionDependenciesPost200ResponseInnerDependencyLock.md) |  | [optional] 
**Project** | Pointer to [**GetRevisionDependenciesPost200ResponseInnerProject**](GetRevisionDependenciesPost200ResponseInnerProject.md) |  | [optional] 
**OriginPaths** | Pointer to **[]string** | Paths from the root to this dependency | [optional] 
**Manual** | Pointer to **bool** | Whether this dependency was manually added | [optional] 
**UnresolvedLocators** | Pointer to **[]string** | List of unresolved locators | [optional] 
**IsSubmodule** | Pointer to **bool** | Whether this is a submodule (deprecated, always false) | [optional] 
**Type** | Pointer to **string** | Type of dependency (deprecated, always \&quot;MEDIATED\&quot;) | [optional] 
**Depth** | Pointer to **int32** | Depth in the dependency tree | [optional] 
**IssueTargets** | Pointer to [**[]GetRevisionDependenciesPost200ResponseInnerIssueTargetsInner**](GetRevisionDependenciesPost200ResponseInnerIssueTargetsInner.md) |  | [optional] 
**Licenses** | Pointer to [**[]GetRevisionDependenciesPost200ResponseInnerLicensesInner**](GetRevisionDependenciesPost200ResponseInnerLicensesInner.md) |  | [optional] 
**Ignored** | Pointer to **bool** | Whether this dependency is ignored | [optional] 
**DownloadUrl** | Pointer to **string** | Download URL for the dependency package | [optional] 
**Hash** | Pointer to **string** | Hash of the dependency (included if includeHashData is true) | [optional] 
**Version** | Pointer to **string** | Version of the dependency (included if includeHashData is true) | [optional] 
**IsGolang** | Pointer to **bool** | Whether this is a Go dependency (included if includeHashData is true) | [optional] 
**Layers** | Pointer to [**[]GetRevisionDependenciesPost200ResponseInnerLayersInner**](GetRevisionDependenciesPost200ResponseInnerLayersInner.md) | Container layer information (only for container dependencies) | [optional] 

## Methods

### NewGetRevisionDependenciesPost200ResponseInner

`func NewGetRevisionDependenciesPost200ResponseInner() *GetRevisionDependenciesPost200ResponseInner`

NewGetRevisionDependenciesPost200ResponseInner instantiates a new GetRevisionDependenciesPost200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionDependenciesPost200ResponseInnerWithDefaults

`func NewGetRevisionDependenciesPost200ResponseInnerWithDefaults() *GetRevisionDependenciesPost200ResponseInner`

NewGetRevisionDependenciesPost200ResponseInnerWithDefaults instantiates a new GetRevisionDependenciesPost200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoc

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLoc() GetRevisionDependenciesPost200ResponseInnerLoc`

GetLoc returns the Loc field if non-nil, zero value otherwise.

### GetLocOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLocOk() (*GetRevisionDependenciesPost200ResponseInnerLoc, bool)`

GetLocOk returns a tuple with the Loc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoc

`func (o *GetRevisionDependenciesPost200ResponseInner) SetLoc(v GetRevisionDependenciesPost200ResponseInnerLoc)`

SetLoc sets Loc field to given value.

### HasLoc

`func (o *GetRevisionDependenciesPost200ResponseInner) HasLoc() bool`

HasLoc returns a boolean if a field has been set.

### GetLocator

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetRevisionDependenciesPost200ResponseInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetRevisionDependenciesPost200ResponseInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetProjectId

`func (o *GetRevisionDependenciesPost200ResponseInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetRevisionDependenciesPost200ResponseInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetRevisionDependenciesPost200ResponseInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetResolved

`func (o *GetRevisionDependenciesPost200ResponseInner) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetRevisionDependenciesPost200ResponseInner) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetRevisionDependenciesPost200ResponseInner) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetUnsupported

`func (o *GetRevisionDependenciesPost200ResponseInner) GetUnsupported() bool`

GetUnsupported returns the Unsupported field if non-nil, zero value otherwise.

### GetUnsupportedOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetUnsupportedOk() (*bool, bool)`

GetUnsupportedOk returns a tuple with the Unsupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupported

`func (o *GetRevisionDependenciesPost200ResponseInner) SetUnsupported(v bool)`

SetUnsupported sets Unsupported field to given value.

### HasUnsupported

`func (o *GetRevisionDependenciesPost200ResponseInner) HasUnsupported() bool`

HasUnsupported returns a boolean if a field has been set.

### GetLatestRevisionScanId

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLatestRevisionScanId() int32`

GetLatestRevisionScanId returns the LatestRevisionScanId field if non-nil, zero value otherwise.

### GetLatestRevisionScanIdOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLatestRevisionScanIdOk() (*int32, bool)`

GetLatestRevisionScanIdOk returns a tuple with the LatestRevisionScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisionScanId

`func (o *GetRevisionDependenciesPost200ResponseInner) SetLatestRevisionScanId(v int32)`

SetLatestRevisionScanId sets LatestRevisionScanId field to given value.

### HasLatestRevisionScanId

`func (o *GetRevisionDependenciesPost200ResponseInner) HasLatestRevisionScanId() bool`

HasLatestRevisionScanId returns a boolean if a field has been set.

### GetDependencyLock

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDependencyLock() GetRevisionDependenciesPost200ResponseInnerDependencyLock`

GetDependencyLock returns the DependencyLock field if non-nil, zero value otherwise.

### GetDependencyLockOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDependencyLockOk() (*GetRevisionDependenciesPost200ResponseInnerDependencyLock, bool)`

GetDependencyLockOk returns a tuple with the DependencyLock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyLock

`func (o *GetRevisionDependenciesPost200ResponseInner) SetDependencyLock(v GetRevisionDependenciesPost200ResponseInnerDependencyLock)`

SetDependencyLock sets DependencyLock field to given value.

### HasDependencyLock

`func (o *GetRevisionDependenciesPost200ResponseInner) HasDependencyLock() bool`

HasDependencyLock returns a boolean if a field has been set.

### GetProject

`func (o *GetRevisionDependenciesPost200ResponseInner) GetProject() GetRevisionDependenciesPost200ResponseInnerProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetProjectOk() (*GetRevisionDependenciesPost200ResponseInnerProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetRevisionDependenciesPost200ResponseInner) SetProject(v GetRevisionDependenciesPost200ResponseInnerProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetRevisionDependenciesPost200ResponseInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInner) GetOriginPaths() []string`

GetOriginPaths returns the OriginPaths field if non-nil, zero value otherwise.

### GetOriginPathsOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetOriginPathsOk() (*[]string, bool)`

GetOriginPathsOk returns a tuple with the OriginPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInner) SetOriginPaths(v []string)`

SetOriginPaths sets OriginPaths field to given value.

### HasOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInner) HasOriginPaths() bool`

HasOriginPaths returns a boolean if a field has been set.

### GetManual

`func (o *GetRevisionDependenciesPost200ResponseInner) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *GetRevisionDependenciesPost200ResponseInner) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *GetRevisionDependenciesPost200ResponseInner) HasManual() bool`

HasManual returns a boolean if a field has been set.

### GetUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInner) GetUnresolvedLocators() []string`

GetUnresolvedLocators returns the UnresolvedLocators field if non-nil, zero value otherwise.

### GetUnresolvedLocatorsOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetUnresolvedLocatorsOk() (*[]string, bool)`

GetUnresolvedLocatorsOk returns a tuple with the UnresolvedLocators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInner) SetUnresolvedLocators(v []string)`

SetUnresolvedLocators sets UnresolvedLocators field to given value.

### HasUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInner) HasUnresolvedLocators() bool`

HasUnresolvedLocators returns a boolean if a field has been set.

### GetIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIsSubmodule() bool`

GetIsSubmodule returns the IsSubmodule field if non-nil, zero value otherwise.

### GetIsSubmoduleOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIsSubmoduleOk() (*bool, bool)`

GetIsSubmoduleOk returns a tuple with the IsSubmodule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInner) SetIsSubmodule(v bool)`

SetIsSubmodule sets IsSubmodule field to given value.

### HasIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInner) HasIsSubmodule() bool`

HasIsSubmodule returns a boolean if a field has been set.

### GetType

`func (o *GetRevisionDependenciesPost200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRevisionDependenciesPost200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRevisionDependenciesPost200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDepth

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetRevisionDependenciesPost200ResponseInner) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetRevisionDependenciesPost200ResponseInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetIssueTargets

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIssueTargets() []GetRevisionDependenciesPost200ResponseInnerIssueTargetsInner`

GetIssueTargets returns the IssueTargets field if non-nil, zero value otherwise.

### GetIssueTargetsOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIssueTargetsOk() (*[]GetRevisionDependenciesPost200ResponseInnerIssueTargetsInner, bool)`

GetIssueTargetsOk returns a tuple with the IssueTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTargets

`func (o *GetRevisionDependenciesPost200ResponseInner) SetIssueTargets(v []GetRevisionDependenciesPost200ResponseInnerIssueTargetsInner)`

SetIssueTargets sets IssueTargets field to given value.

### HasIssueTargets

`func (o *GetRevisionDependenciesPost200ResponseInner) HasIssueTargets() bool`

HasIssueTargets returns a boolean if a field has been set.

### GetLicenses

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLicenses() []GetRevisionDependenciesPost200ResponseInnerLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLicensesOk() (*[]GetRevisionDependenciesPost200ResponseInnerLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetRevisionDependenciesPost200ResponseInner) SetLicenses(v []GetRevisionDependenciesPost200ResponseInnerLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetRevisionDependenciesPost200ResponseInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetIgnored

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *GetRevisionDependenciesPost200ResponseInner) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *GetRevisionDependenciesPost200ResponseInner) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *GetRevisionDependenciesPost200ResponseInner) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *GetRevisionDependenciesPost200ResponseInner) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetHash

`func (o *GetRevisionDependenciesPost200ResponseInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetRevisionDependenciesPost200ResponseInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetRevisionDependenciesPost200ResponseInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetVersion

`func (o *GetRevisionDependenciesPost200ResponseInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRevisionDependenciesPost200ResponseInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRevisionDependenciesPost200ResponseInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsGolang

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIsGolang() bool`

GetIsGolang returns the IsGolang field if non-nil, zero value otherwise.

### GetIsGolangOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetIsGolangOk() (*bool, bool)`

GetIsGolangOk returns a tuple with the IsGolang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGolang

`func (o *GetRevisionDependenciesPost200ResponseInner) SetIsGolang(v bool)`

SetIsGolang sets IsGolang field to given value.

### HasIsGolang

`func (o *GetRevisionDependenciesPost200ResponseInner) HasIsGolang() bool`

HasIsGolang returns a boolean if a field has been set.

### GetLayers

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLayers() []GetRevisionDependenciesPost200ResponseInnerLayersInner`

GetLayers returns the Layers field if non-nil, zero value otherwise.

### GetLayersOk

`func (o *GetRevisionDependenciesPost200ResponseInner) GetLayersOk() (*[]GetRevisionDependenciesPost200ResponseInnerLayersInner, bool)`

GetLayersOk returns a tuple with the Layers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayers

`func (o *GetRevisionDependenciesPost200ResponseInner) SetLayers(v []GetRevisionDependenciesPost200ResponseInnerLayersInner)`

SetLayers sets Layers field to given value.

### HasLayers

`func (o *GetRevisionDependenciesPost200ResponseInner) HasLayers() bool`

HasLayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


