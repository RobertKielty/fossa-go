# GetProjectDependencies200ResponseDependenciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locator** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**IsManual** | Pointer to **bool** |  | [optional] 
**IsIgnored** | Pointer to **bool** |  | [optional] 
**IsUnknown** | Pointer to **bool** |  | [optional] 
**Licenses** | Pointer to **[]string** |  | [optional] 
**DeclaredLicenses** | Pointer to **[]string** |  | [optional] 
**Depth** | Pointer to **int32** |  | [optional] 
**OriginPaths** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to [**GetProjectDependencies200ResponseDependenciesInnerStatus**](GetProjectDependencies200ResponseDependenciesInnerStatus.md) |  | [optional] 
**Issues** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInnerIssuesInner**](GetProjectDependencies200ResponseDependenciesInnerIssuesInner.md) |  | [optional] 
**RootProjects** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner**](GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner.md) |  | [optional] 
**LayerDepth** | Pointer to **float32** |  | [optional] 
**Cpes** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewGetProjectDependencies200ResponseDependenciesInner

`func NewGetProjectDependencies200ResponseDependenciesInner() *GetProjectDependencies200ResponseDependenciesInner`

NewGetProjectDependencies200ResponseDependenciesInner instantiates a new GetProjectDependencies200ResponseDependenciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectDependencies200ResponseDependenciesInnerWithDefaults

`func NewGetProjectDependencies200ResponseDependenciesInnerWithDefaults() *GetProjectDependencies200ResponseDependenciesInner`

NewGetProjectDependencies200ResponseDependenciesInnerWithDefaults instantiates a new GetProjectDependencies200ResponseDependenciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocator

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetTitle

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIsManual

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsIgnored

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsIgnored() bool`

GetIsIgnored returns the IsIgnored field if non-nil, zero value otherwise.

### GetIsIgnoredOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsIgnoredOk() (*bool, bool)`

GetIsIgnoredOk returns a tuple with the IsIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIgnored

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsIgnored(v bool)`

SetIsIgnored sets IsIgnored field to given value.

### HasIsIgnored

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsIgnored() bool`

HasIsIgnored returns a boolean if a field has been set.

### GetIsUnknown

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsUnknown() bool`

GetIsUnknown returns the IsUnknown field if non-nil, zero value otherwise.

### GetIsUnknownOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIsUnknownOk() (*bool, bool)`

GetIsUnknownOk returns a tuple with the IsUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnknown

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetIsUnknown(v bool)`

SetIsUnknown sets IsUnknown field to given value.

### HasIsUnknown

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasIsUnknown() bool`

HasIsUnknown returns a boolean if a field has been set.

### GetLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLicenses() []string`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLicensesOk() (*[]string, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetLicenses(v []string)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetDeclaredLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetDeclaredLicenses() []string`

GetDeclaredLicenses returns the DeclaredLicenses field if non-nil, zero value otherwise.

### GetDeclaredLicensesOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetDeclaredLicensesOk() (*[]string, bool)`

GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetDeclaredLicenses(v []string)`

SetDeclaredLicenses sets DeclaredLicenses field to given value.

### HasDeclaredLicenses

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasDeclaredLicenses() bool`

HasDeclaredLicenses returns a boolean if a field has been set.

### GetDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetOriginPaths

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetOriginPaths() []string`

GetOriginPaths returns the OriginPaths field if non-nil, zero value otherwise.

### GetOriginPathsOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetOriginPathsOk() (*[]string, bool)`

GetOriginPathsOk returns a tuple with the OriginPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPaths

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetOriginPaths(v []string)`

SetOriginPaths sets OriginPaths field to given value.

### HasOriginPaths

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasOriginPaths() bool`

HasOriginPaths returns a boolean if a field has been set.

### GetStatus

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetStatus() GetProjectDependencies200ResponseDependenciesInnerStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetStatusOk() (*GetProjectDependencies200ResponseDependenciesInnerStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetStatus(v GetProjectDependencies200ResponseDependenciesInnerStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIssues

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIssues() []GetProjectDependencies200ResponseDependenciesInnerIssuesInner`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetIssuesOk() (*[]GetProjectDependencies200ResponseDependenciesInnerIssuesInner, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetIssues(v []GetProjectDependencies200ResponseDependenciesInnerIssuesInner)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRootProjects

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetRootProjects() []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner`

GetRootProjects returns the RootProjects field if non-nil, zero value otherwise.

### GetRootProjectsOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetRootProjectsOk() (*[]GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner, bool)`

GetRootProjectsOk returns a tuple with the RootProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProjects

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetRootProjects(v []GetProjectDependencies200ResponseDependenciesInnerRootProjectsInner)`

SetRootProjects sets RootProjects field to given value.

### HasRootProjects

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasRootProjects() bool`

HasRootProjects returns a boolean if a field has been set.

### GetLayerDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLayerDepth() float32`

GetLayerDepth returns the LayerDepth field if non-nil, zero value otherwise.

### GetLayerDepthOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetLayerDepthOk() (*float32, bool)`

GetLayerDepthOk returns a tuple with the LayerDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayerDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetLayerDepth(v float32)`

SetLayerDepth sets LayerDepth field to given value.

### HasLayerDepth

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasLayerDepth() bool`

HasLayerDepth returns a boolean if a field has been set.

### GetCpes

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetVersion

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetProjectDependencies200ResponseDependenciesInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetProjectDependencies200ResponseDependenciesInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetProjectDependencies200ResponseDependenciesInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


