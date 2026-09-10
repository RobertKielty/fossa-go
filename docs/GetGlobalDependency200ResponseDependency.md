# GetGlobalDependency200ResponseDependency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locator** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**IsManual** | Pointer to **bool** |  | [optional] 
**IsIgnored** | Pointer to **bool** |  | [optional] 
**IsUnknown** | Pointer to **bool** |  | [optional] 
**IsSnippetConversion** | Pointer to **bool** |  | [optional] 
**Conclusions** | Pointer to [**GetProjectDependencies200ResponseDependenciesInnerConclusions**](GetProjectDependencies200ResponseDependenciesInnerConclusions.md) |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**DeclaredLicenses** | Pointer to **[]string** |  | [optional] 
**DiscoveredLicenses** | Pointer to **[]string** |  | [optional] 
**LicenseGroups** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInnerLicenseGroupsInner**](GetProjectDependencies200ResponseDependenciesInnerLicenseGroupsInner.md) | The grouped licenses discovered for this dependency. | [optional] 
**Labels** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner**](GetProjectDependencies200ResponseDependenciesInnerLabelsInner.md) | The package label assignments applied to this dependency. | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 
**OriginPaths** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**GetProjectDependencies200ResponseDependenciesInnerStatus**](GetProjectDependencies200ResponseDependenciesInnerStatus.md) |  | [optional] 
**Issues** | Pointer to [**[]GetGlobalDependency200ResponseDependencyAllOfIssuesInner**](GetGlobalDependency200ResponseDependencyAllOfIssuesInner.md) |  | [optional] 
**RootProjects** | Pointer to [**[]GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner**](GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner.md) |  | [optional] 
**LayerDepth** | Pointer to **float32** |  | [optional] 
**Cpes** | Pointer to **[]string** |  | [optional] 
**VendoredPaths** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Authors** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**IndividualLicenses** | Pointer to **[]string** |  | [optional] 
**GroupedLicenses** | Pointer to [**[]GetGlobalDependency200ResponseDependencyAllOfGroupedLicensesInner**](GetGlobalDependency200ResponseDependencyAllOfGroupedLicensesInner.md) | The flattened, grouped licenses discovered for this dependency. | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**LastPublishedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetGlobalDependency200ResponseDependency

`func NewGetGlobalDependency200ResponseDependency() *GetGlobalDependency200ResponseDependency`

NewGetGlobalDependency200ResponseDependency instantiates a new GetGlobalDependency200ResponseDependency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalDependency200ResponseDependencyWithDefaults

`func NewGetGlobalDependency200ResponseDependencyWithDefaults() *GetGlobalDependency200ResponseDependency`

NewGetGlobalDependency200ResponseDependencyWithDefaults instantiates a new GetGlobalDependency200ResponseDependency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocator

`func (o *GetGlobalDependency200ResponseDependency) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetGlobalDependency200ResponseDependency) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetGlobalDependency200ResponseDependency) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetGlobalDependency200ResponseDependency) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetTitle

`func (o *GetGlobalDependency200ResponseDependency) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetGlobalDependency200ResponseDependency) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetGlobalDependency200ResponseDependency) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetGlobalDependency200ResponseDependency) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsManual

`func (o *GetGlobalDependency200ResponseDependency) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *GetGlobalDependency200ResponseDependency) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *GetGlobalDependency200ResponseDependency) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *GetGlobalDependency200ResponseDependency) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsIgnored

`func (o *GetGlobalDependency200ResponseDependency) GetIsIgnored() bool`

GetIsIgnored returns the IsIgnored field if non-nil, zero value otherwise.

### GetIsIgnoredOk

`func (o *GetGlobalDependency200ResponseDependency) GetIsIgnoredOk() (*bool, bool)`

GetIsIgnoredOk returns a tuple with the IsIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIgnored

`func (o *GetGlobalDependency200ResponseDependency) SetIsIgnored(v bool)`

SetIsIgnored sets IsIgnored field to given value.

### HasIsIgnored

`func (o *GetGlobalDependency200ResponseDependency) HasIsIgnored() bool`

HasIsIgnored returns a boolean if a field has been set.

### GetIsUnknown

`func (o *GetGlobalDependency200ResponseDependency) GetIsUnknown() bool`

GetIsUnknown returns the IsUnknown field if non-nil, zero value otherwise.

### GetIsUnknownOk

`func (o *GetGlobalDependency200ResponseDependency) GetIsUnknownOk() (*bool, bool)`

GetIsUnknownOk returns a tuple with the IsUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnknown

`func (o *GetGlobalDependency200ResponseDependency) SetIsUnknown(v bool)`

SetIsUnknown sets IsUnknown field to given value.

### HasIsUnknown

`func (o *GetGlobalDependency200ResponseDependency) HasIsUnknown() bool`

HasIsUnknown returns a boolean if a field has been set.

### GetIsSnippetConversion

`func (o *GetGlobalDependency200ResponseDependency) GetIsSnippetConversion() bool`

GetIsSnippetConversion returns the IsSnippetConversion field if non-nil, zero value otherwise.

### GetIsSnippetConversionOk

`func (o *GetGlobalDependency200ResponseDependency) GetIsSnippetConversionOk() (*bool, bool)`

GetIsSnippetConversionOk returns a tuple with the IsSnippetConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnippetConversion

`func (o *GetGlobalDependency200ResponseDependency) SetIsSnippetConversion(v bool)`

SetIsSnippetConversion sets IsSnippetConversion field to given value.

### HasIsSnippetConversion

`func (o *GetGlobalDependency200ResponseDependency) HasIsSnippetConversion() bool`

HasIsSnippetConversion returns a boolean if a field has been set.

### GetConclusions

`func (o *GetGlobalDependency200ResponseDependency) GetConclusions() GetProjectDependencies200ResponseDependenciesInnerConclusions`

GetConclusions returns the Conclusions field if non-nil, zero value otherwise.

### GetConclusionsOk

`func (o *GetGlobalDependency200ResponseDependency) GetConclusionsOk() (*GetProjectDependencies200ResponseDependenciesInnerConclusions, bool)`

GetConclusionsOk returns a tuple with the Conclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusions

`func (o *GetGlobalDependency200ResponseDependency) SetConclusions(v GetProjectDependencies200ResponseDependenciesInnerConclusions)`

SetConclusions sets Conclusions field to given value.

### HasConclusions

`func (o *GetGlobalDependency200ResponseDependency) HasConclusions() bool`

HasConclusions returns a boolean if a field has been set.

### GetLicenses

`func (o *GetGlobalDependency200ResponseDependency) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetGlobalDependency200ResponseDependency) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetGlobalDependency200ResponseDependency) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetGlobalDependency200ResponseDependency) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetDeclaredLicenses

`func (o *GetGlobalDependency200ResponseDependency) GetDeclaredLicenses() []string`

GetDeclaredLicenses returns the DeclaredLicenses field if non-nil, zero value otherwise.

### GetDeclaredLicensesOk

`func (o *GetGlobalDependency200ResponseDependency) GetDeclaredLicensesOk() (*[]string, bool)`

GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLicenses

`func (o *GetGlobalDependency200ResponseDependency) SetDeclaredLicenses(v []string)`

SetDeclaredLicenses sets DeclaredLicenses field to given value.

### HasDeclaredLicenses

`func (o *GetGlobalDependency200ResponseDependency) HasDeclaredLicenses() bool`

HasDeclaredLicenses returns a boolean if a field has been set.

### GetDiscoveredLicenses

`func (o *GetGlobalDependency200ResponseDependency) GetDiscoveredLicenses() []string`

GetDiscoveredLicenses returns the DiscoveredLicenses field if non-nil, zero value otherwise.

### GetDiscoveredLicensesOk

`func (o *GetGlobalDependency200ResponseDependency) GetDiscoveredLicensesOk() (*[]string, bool)`

GetDiscoveredLicensesOk returns a tuple with the DiscoveredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredLicenses

`func (o *GetGlobalDependency200ResponseDependency) SetDiscoveredLicenses(v []string)`

SetDiscoveredLicenses sets DiscoveredLicenses field to given value.

### HasDiscoveredLicenses

`func (o *GetGlobalDependency200ResponseDependency) HasDiscoveredLicenses() bool`

HasDiscoveredLicenses returns a boolean if a field has been set.

### GetLicenseGroups

`func (o *GetGlobalDependency200ResponseDependency) GetLicenseGroups() []GetProjectDependencies200ResponseDependenciesInnerLicenseGroupsInner`

GetLicenseGroups returns the LicenseGroups field if non-nil, zero value otherwise.

### GetLicenseGroupsOk

`func (o *GetGlobalDependency200ResponseDependency) GetLicenseGroupsOk() (*[]GetProjectDependencies200ResponseDependenciesInnerLicenseGroupsInner, bool)`

GetLicenseGroupsOk returns a tuple with the LicenseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseGroups

`func (o *GetGlobalDependency200ResponseDependency) SetLicenseGroups(v []GetProjectDependencies200ResponseDependenciesInnerLicenseGroupsInner)`

SetLicenseGroups sets LicenseGroups field to given value.

### HasLicenseGroups

`func (o *GetGlobalDependency200ResponseDependency) HasLicenseGroups() bool`

HasLicenseGroups returns a boolean if a field has been set.

### GetLabels

`func (o *GetGlobalDependency200ResponseDependency) GetLabels() []GetProjectDependencies200ResponseDependenciesInnerLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetGlobalDependency200ResponseDependency) GetLabelsOk() (*[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetGlobalDependency200ResponseDependency) SetLabels(v []GetProjectDependencies200ResponseDependenciesInnerLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetGlobalDependency200ResponseDependency) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDepth

`func (o *GetGlobalDependency200ResponseDependency) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetGlobalDependency200ResponseDependency) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetGlobalDependency200ResponseDependency) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetGlobalDependency200ResponseDependency) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetOriginPaths

`func (o *GetGlobalDependency200ResponseDependency) GetOriginPaths() []string`

GetOriginPaths returns the OriginPaths field if non-nil, zero value otherwise.

### GetOriginPathsOk

`func (o *GetGlobalDependency200ResponseDependency) GetOriginPathsOk() (*[]string, bool)`

GetOriginPathsOk returns a tuple with the OriginPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPaths

`func (o *GetGlobalDependency200ResponseDependency) SetOriginPaths(v []string)`

SetOriginPaths sets OriginPaths field to given value.

### HasOriginPaths

`func (o *GetGlobalDependency200ResponseDependency) HasOriginPaths() bool`

HasOriginPaths returns a boolean if a field has been set.

### GetStatus

`func (o *GetGlobalDependency200ResponseDependency) GetStatus() GetProjectDependencies200ResponseDependenciesInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGlobalDependency200ResponseDependency) GetStatusOk() (*GetProjectDependencies200ResponseDependenciesInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGlobalDependency200ResponseDependency) SetStatus(v GetProjectDependencies200ResponseDependenciesInnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGlobalDependency200ResponseDependency) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIssues

`func (o *GetGlobalDependency200ResponseDependency) GetIssues() []GetGlobalDependency200ResponseDependencyAllOfIssuesInner`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetGlobalDependency200ResponseDependency) GetIssuesOk() (*[]GetGlobalDependency200ResponseDependencyAllOfIssuesInner, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetGlobalDependency200ResponseDependency) SetIssues(v []GetGlobalDependency200ResponseDependencyAllOfIssuesInner)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetGlobalDependency200ResponseDependency) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRootProjects

`func (o *GetGlobalDependency200ResponseDependency) GetRootProjects() []GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner`

GetRootProjects returns the RootProjects field if non-nil, zero value otherwise.

### GetRootProjectsOk

`func (o *GetGlobalDependency200ResponseDependency) GetRootProjectsOk() (*[]GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner, bool)`

GetRootProjectsOk returns a tuple with the RootProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProjects

`func (o *GetGlobalDependency200ResponseDependency) SetRootProjects(v []GetGlobalDependency200ResponseDependencyAllOfRootProjectsInner)`

SetRootProjects sets RootProjects field to given value.

### HasRootProjects

`func (o *GetGlobalDependency200ResponseDependency) HasRootProjects() bool`

HasRootProjects returns a boolean if a field has been set.

### GetLayerDepth

`func (o *GetGlobalDependency200ResponseDependency) GetLayerDepth() float32`

GetLayerDepth returns the LayerDepth field if non-nil, zero value otherwise.

### GetLayerDepthOk

`func (o *GetGlobalDependency200ResponseDependency) GetLayerDepthOk() (*float32, bool)`

GetLayerDepthOk returns a tuple with the LayerDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerDepth

`func (o *GetGlobalDependency200ResponseDependency) SetLayerDepth(v float32)`

SetLayerDepth sets LayerDepth field to given value.

### HasLayerDepth

`func (o *GetGlobalDependency200ResponseDependency) HasLayerDepth() bool`

HasLayerDepth returns a boolean if a field has been set.

### GetCpes

`func (o *GetGlobalDependency200ResponseDependency) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *GetGlobalDependency200ResponseDependency) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *GetGlobalDependency200ResponseDependency) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *GetGlobalDependency200ResponseDependency) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetVendoredPaths

`func (o *GetGlobalDependency200ResponseDependency) GetVendoredPaths() []string`

GetVendoredPaths returns the VendoredPaths field if non-nil, zero value otherwise.

### GetVendoredPathsOk

`func (o *GetGlobalDependency200ResponseDependency) GetVendoredPathsOk() (*[]string, bool)`

GetVendoredPathsOk returns a tuple with the VendoredPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendoredPaths

`func (o *GetGlobalDependency200ResponseDependency) SetVendoredPaths(v []string)`

SetVendoredPaths sets VendoredPaths field to given value.

### HasVendoredPaths

`func (o *GetGlobalDependency200ResponseDependency) HasVendoredPaths() bool`

HasVendoredPaths returns a boolean if a field has been set.

### GetVersion

`func (o *GetGlobalDependency200ResponseDependency) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetGlobalDependency200ResponseDependency) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetGlobalDependency200ResponseDependency) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetGlobalDependency200ResponseDependency) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetAuthors

`func (o *GetGlobalDependency200ResponseDependency) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *GetGlobalDependency200ResponseDependency) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *GetGlobalDependency200ResponseDependency) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *GetGlobalDependency200ResponseDependency) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetDescription

`func (o *GetGlobalDependency200ResponseDependency) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGlobalDependency200ResponseDependency) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGlobalDependency200ResponseDependency) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGlobalDependency200ResponseDependency) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *GetGlobalDependency200ResponseDependency) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetGlobalDependency200ResponseDependency) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetGlobalDependency200ResponseDependency) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetGlobalDependency200ResponseDependency) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIndividualLicenses

`func (o *GetGlobalDependency200ResponseDependency) GetIndividualLicenses() []string`

GetIndividualLicenses returns the IndividualLicenses field if non-nil, zero value otherwise.

### GetIndividualLicensesOk

`func (o *GetGlobalDependency200ResponseDependency) GetIndividualLicensesOk() (*[]string, bool)`

GetIndividualLicensesOk returns a tuple with the IndividualLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndividualLicenses

`func (o *GetGlobalDependency200ResponseDependency) SetIndividualLicenses(v []string)`

SetIndividualLicenses sets IndividualLicenses field to given value.

### HasIndividualLicenses

`func (o *GetGlobalDependency200ResponseDependency) HasIndividualLicenses() bool`

HasIndividualLicenses returns a boolean if a field has been set.

### GetGroupedLicenses

`func (o *GetGlobalDependency200ResponseDependency) GetGroupedLicenses() []GetGlobalDependency200ResponseDependencyAllOfGroupedLicensesInner`

GetGroupedLicenses returns the GroupedLicenses field if non-nil, zero value otherwise.

### GetGroupedLicensesOk

`func (o *GetGlobalDependency200ResponseDependency) GetGroupedLicensesOk() (*[]GetGlobalDependency200ResponseDependencyAllOfGroupedLicensesInner, bool)`

GetGroupedLicensesOk returns a tuple with the GroupedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupedLicenses

`func (o *GetGlobalDependency200ResponseDependency) SetGroupedLicenses(v []GetGlobalDependency200ResponseDependencyAllOfGroupedLicensesInner)`

SetGroupedLicenses sets GroupedLicenses field to given value.

### HasGroupedLicenses

`func (o *GetGlobalDependency200ResponseDependency) HasGroupedLicenses() bool`

HasGroupedLicenses returns a boolean if a field has been set.

### GetSourceType

`func (o *GetGlobalDependency200ResponseDependency) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *GetGlobalDependency200ResponseDependency) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *GetGlobalDependency200ResponseDependency) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *GetGlobalDependency200ResponseDependency) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetLastPublishedDate

`func (o *GetGlobalDependency200ResponseDependency) GetLastPublishedDate() time.Time`

GetLastPublishedDate returns the LastPublishedDate field if non-nil, zero value otherwise.

### GetLastPublishedDateOk

`func (o *GetGlobalDependency200ResponseDependency) GetLastPublishedDateOk() (*time.Time, bool)`

GetLastPublishedDateOk returns a tuple with the LastPublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPublishedDate

`func (o *GetGlobalDependency200ResponseDependency) SetLastPublishedDate(v time.Time)`

SetLastPublishedDate sets LastPublishedDate field to given value.

### HasLastPublishedDate

`func (o *GetGlobalDependency200ResponseDependency) HasLastPublishedDate() bool`

HasLastPublishedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


