# CreateSavedFilterRequestFilterOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | Pointer to **string** |  | [optional] 
**HasFix** | Pointer to **string** |  | [optional] 
**UpgradeDistance** | Pointer to **string** |  | [optional] 
**ExploitMaturity** | Pointer to **string** |  | [optional] 
**IgnoreReason** | Pointer to **string** | Provided reason for ignoring or resolving a security issue. &#39;Fixed&#39; and &#39;Under_investigation&#39; map to VEX statuses with the same names. All other values map to the VEX status &#39;Not Affected&#39;. This value appears in the vulnerabilities.analysis.justification field of CycloneDX SBOM reports.  | [optional] 
**Reachability** | Pointer to **string** |  | [optional] 
**Cwes** | Pointer to **[]string** |  | [optional] 
**PackageManagers** | Pointer to **[]string** |  | [optional] 
**ProjectLabels** | Pointer to **[]string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Ticketed** | Pointer to **string** |  | [optional] 
**RevisionIds** | Pointer to **[]string** |  | [optional] 
**ContainerLayers** | Pointer to **[]string** |  | [optional] 
**FoundAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateSavedFilterRequestFilterOneOf1

`func NewCreateSavedFilterRequestFilterOneOf1() *CreateSavedFilterRequestFilterOneOf1`

NewCreateSavedFilterRequestFilterOneOf1 instantiates a new CreateSavedFilterRequestFilterOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedFilterRequestFilterOneOf1WithDefaults

`func NewCreateSavedFilterRequestFilterOneOf1WithDefaults() *CreateSavedFilterRequestFilterOneOf1`

NewCreateSavedFilterRequestFilterOneOf1WithDefaults instantiates a new CreateSavedFilterRequestFilterOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *CreateSavedFilterRequestFilterOneOf1) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateSavedFilterRequestFilterOneOf1) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CreateSavedFilterRequestFilterOneOf1) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetHasFix

`func (o *CreateSavedFilterRequestFilterOneOf1) GetHasFix() string`

GetHasFix returns the HasFix field if non-nil, zero value otherwise.

### GetHasFixOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetHasFixOk() (*string, bool)`

GetHasFixOk returns a tuple with the HasFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFix

`func (o *CreateSavedFilterRequestFilterOneOf1) SetHasFix(v string)`

SetHasFix sets HasFix field to given value.

### HasHasFix

`func (o *CreateSavedFilterRequestFilterOneOf1) HasHasFix() bool`

HasHasFix returns a boolean if a field has been set.

### GetUpgradeDistance

`func (o *CreateSavedFilterRequestFilterOneOf1) GetUpgradeDistance() string`

GetUpgradeDistance returns the UpgradeDistance field if non-nil, zero value otherwise.

### GetUpgradeDistanceOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetUpgradeDistanceOk() (*string, bool)`

GetUpgradeDistanceOk returns a tuple with the UpgradeDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDistance

`func (o *CreateSavedFilterRequestFilterOneOf1) SetUpgradeDistance(v string)`

SetUpgradeDistance sets UpgradeDistance field to given value.

### HasUpgradeDistance

`func (o *CreateSavedFilterRequestFilterOneOf1) HasUpgradeDistance() bool`

HasUpgradeDistance returns a boolean if a field has been set.

### GetExploitMaturity

`func (o *CreateSavedFilterRequestFilterOneOf1) GetExploitMaturity() string`

GetExploitMaturity returns the ExploitMaturity field if non-nil, zero value otherwise.

### GetExploitMaturityOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetExploitMaturityOk() (*string, bool)`

GetExploitMaturityOk returns a tuple with the ExploitMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitMaturity

`func (o *CreateSavedFilterRequestFilterOneOf1) SetExploitMaturity(v string)`

SetExploitMaturity sets ExploitMaturity field to given value.

### HasExploitMaturity

`func (o *CreateSavedFilterRequestFilterOneOf1) HasExploitMaturity() bool`

HasExploitMaturity returns a boolean if a field has been set.

### GetIgnoreReason

`func (o *CreateSavedFilterRequestFilterOneOf1) GetIgnoreReason() string`

GetIgnoreReason returns the IgnoreReason field if non-nil, zero value otherwise.

### GetIgnoreReasonOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetIgnoreReasonOk() (*string, bool)`

GetIgnoreReasonOk returns a tuple with the IgnoreReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreReason

`func (o *CreateSavedFilterRequestFilterOneOf1) SetIgnoreReason(v string)`

SetIgnoreReason sets IgnoreReason field to given value.

### HasIgnoreReason

`func (o *CreateSavedFilterRequestFilterOneOf1) HasIgnoreReason() bool`

HasIgnoreReason returns a boolean if a field has been set.

### GetReachability

`func (o *CreateSavedFilterRequestFilterOneOf1) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *CreateSavedFilterRequestFilterOneOf1) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *CreateSavedFilterRequestFilterOneOf1) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetCwes

`func (o *CreateSavedFilterRequestFilterOneOf1) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *CreateSavedFilterRequestFilterOneOf1) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *CreateSavedFilterRequestFilterOneOf1) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf1) GetPackageManagers() []string`

GetPackageManagers returns the PackageManagers field if non-nil, zero value otherwise.

### GetPackageManagersOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetPackageManagersOk() (*[]string, bool)`

GetPackageManagersOk returns a tuple with the PackageManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf1) SetPackageManagers(v []string)`

SetPackageManagers sets PackageManagers field to given value.

### HasPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf1) HasPackageManagers() bool`

HasPackageManagers returns a boolean if a field has been set.

### GetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf1) GetProjectLabels() []string`

GetProjectLabels returns the ProjectLabels field if non-nil, zero value otherwise.

### GetProjectLabelsOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetProjectLabelsOk() (*[]string, bool)`

GetProjectLabelsOk returns a tuple with the ProjectLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf1) SetProjectLabels(v []string)`

SetProjectLabels sets ProjectLabels field to given value.

### HasProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf1) HasProjectLabels() bool`

HasProjectLabels returns a boolean if a field has been set.

### GetSearch

`func (o *CreateSavedFilterRequestFilterOneOf1) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreateSavedFilterRequestFilterOneOf1) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreateSavedFilterRequestFilterOneOf1) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf1) GetTicketed() string`

GetTicketed returns the Ticketed field if non-nil, zero value otherwise.

### GetTicketedOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetTicketedOk() (*string, bool)`

GetTicketedOk returns a tuple with the Ticketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf1) SetTicketed(v string)`

SetTicketed sets Ticketed field to given value.

### HasTicketed

`func (o *CreateSavedFilterRequestFilterOneOf1) HasTicketed() bool`

HasTicketed returns a boolean if a field has been set.

### GetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf1) GetRevisionIds() []string`

GetRevisionIds returns the RevisionIds field if non-nil, zero value otherwise.

### GetRevisionIdsOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetRevisionIdsOk() (*[]string, bool)`

GetRevisionIdsOk returns a tuple with the RevisionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf1) SetRevisionIds(v []string)`

SetRevisionIds sets RevisionIds field to given value.

### HasRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf1) HasRevisionIds() bool`

HasRevisionIds returns a boolean if a field has been set.

### GetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf1) GetContainerLayers() []string`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetContainerLayersOk() (*[]string, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf1) SetContainerLayers(v []string)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf1) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf1) GetFoundAfter() time.Time`

GetFoundAfter returns the FoundAfter field if non-nil, zero value otherwise.

### GetFoundAfterOk

`func (o *CreateSavedFilterRequestFilterOneOf1) GetFoundAfterOk() (*time.Time, bool)`

GetFoundAfterOk returns a tuple with the FoundAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf1) SetFoundAfter(v time.Time)`

SetFoundAfter sets FoundAfter field to given value.

### HasFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf1) HasFoundAfter() bool`

HasFoundAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


