# CreateSavedFilterRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Identification** | Pointer to **string** |  | [optional] 
**PackageManagers** | Pointer to **[]string** |  | [optional] 
**ProjectLabels** | Pointer to **[]string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Ticketed** | Pointer to **string** |  | [optional] 
**RevisionIds** | Pointer to **[]string** |  | [optional] 
**ContainerLayers** | Pointer to **[]string** |  | [optional] 
**FoundAfter** | Pointer to **time.Time** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**HasFix** | Pointer to **string** |  | [optional] 
**UpgradeDistance** | Pointer to **string** |  | [optional] 
**ExploitMaturity** | Pointer to **string** |  | [optional] 
**IgnoreReason** | Pointer to **string** | Provided reason for ignoring or resolving a security issue. &#39;Fixed&#39; and &#39;Under_investigation&#39; map to VEX statuses with the same names. All other values map to the VEX status &#39;Not Affected&#39;. This value appears in the vulnerabilities.analysis.justification field of CycloneDX SBOM reports.  | [optional] 
**Reachability** | Pointer to **string** |  | [optional] 
**Cwes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateSavedFilterRequestFilter

`func NewCreateSavedFilterRequestFilter() *CreateSavedFilterRequestFilter`

NewCreateSavedFilterRequestFilter instantiates a new CreateSavedFilterRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedFilterRequestFilterWithDefaults

`func NewCreateSavedFilterRequestFilterWithDefaults() *CreateSavedFilterRequestFilter`

NewCreateSavedFilterRequestFilterWithDefaults instantiates a new CreateSavedFilterRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateSavedFilterRequestFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSavedFilterRequestFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSavedFilterRequestFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSavedFilterRequestFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdentification

`func (o *CreateSavedFilterRequestFilter) GetIdentification() string`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *CreateSavedFilterRequestFilter) GetIdentificationOk() (*string, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *CreateSavedFilterRequestFilter) SetIdentification(v string)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *CreateSavedFilterRequestFilter) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetPackageManagers

`func (o *CreateSavedFilterRequestFilter) GetPackageManagers() []string`

GetPackageManagers returns the PackageManagers field if non-nil, zero value otherwise.

### GetPackageManagersOk

`func (o *CreateSavedFilterRequestFilter) GetPackageManagersOk() (*[]string, bool)`

GetPackageManagersOk returns a tuple with the PackageManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagers

`func (o *CreateSavedFilterRequestFilter) SetPackageManagers(v []string)`

SetPackageManagers sets PackageManagers field to given value.

### HasPackageManagers

`func (o *CreateSavedFilterRequestFilter) HasPackageManagers() bool`

HasPackageManagers returns a boolean if a field has been set.

### GetProjectLabels

`func (o *CreateSavedFilterRequestFilter) GetProjectLabels() []string`

GetProjectLabels returns the ProjectLabels field if non-nil, zero value otherwise.

### GetProjectLabelsOk

`func (o *CreateSavedFilterRequestFilter) GetProjectLabelsOk() (*[]string, bool)`

GetProjectLabelsOk returns a tuple with the ProjectLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabels

`func (o *CreateSavedFilterRequestFilter) SetProjectLabels(v []string)`

SetProjectLabels sets ProjectLabels field to given value.

### HasProjectLabels

`func (o *CreateSavedFilterRequestFilter) HasProjectLabels() bool`

HasProjectLabels returns a boolean if a field has been set.

### GetSearch

`func (o *CreateSavedFilterRequestFilter) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreateSavedFilterRequestFilter) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreateSavedFilterRequestFilter) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreateSavedFilterRequestFilter) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTicketed

`func (o *CreateSavedFilterRequestFilter) GetTicketed() string`

GetTicketed returns the Ticketed field if non-nil, zero value otherwise.

### GetTicketedOk

`func (o *CreateSavedFilterRequestFilter) GetTicketedOk() (*string, bool)`

GetTicketedOk returns a tuple with the Ticketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketed

`func (o *CreateSavedFilterRequestFilter) SetTicketed(v string)`

SetTicketed sets Ticketed field to given value.

### HasTicketed

`func (o *CreateSavedFilterRequestFilter) HasTicketed() bool`

HasTicketed returns a boolean if a field has been set.

### GetRevisionIds

`func (o *CreateSavedFilterRequestFilter) GetRevisionIds() []string`

GetRevisionIds returns the RevisionIds field if non-nil, zero value otherwise.

### GetRevisionIdsOk

`func (o *CreateSavedFilterRequestFilter) GetRevisionIdsOk() (*[]string, bool)`

GetRevisionIdsOk returns a tuple with the RevisionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionIds

`func (o *CreateSavedFilterRequestFilter) SetRevisionIds(v []string)`

SetRevisionIds sets RevisionIds field to given value.

### HasRevisionIds

`func (o *CreateSavedFilterRequestFilter) HasRevisionIds() bool`

HasRevisionIds returns a boolean if a field has been set.

### GetContainerLayers

`func (o *CreateSavedFilterRequestFilter) GetContainerLayers() []string`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *CreateSavedFilterRequestFilter) GetContainerLayersOk() (*[]string, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *CreateSavedFilterRequestFilter) SetContainerLayers(v []string)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *CreateSavedFilterRequestFilter) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetFoundAfter

`func (o *CreateSavedFilterRequestFilter) GetFoundAfter() time.Time`

GetFoundAfter returns the FoundAfter field if non-nil, zero value otherwise.

### GetFoundAfterOk

`func (o *CreateSavedFilterRequestFilter) GetFoundAfterOk() (*time.Time, bool)`

GetFoundAfterOk returns a tuple with the FoundAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAfter

`func (o *CreateSavedFilterRequestFilter) SetFoundAfter(v time.Time)`

SetFoundAfter sets FoundAfter field to given value.

### HasFoundAfter

`func (o *CreateSavedFilterRequestFilter) HasFoundAfter() bool`

HasFoundAfter returns a boolean if a field has been set.

### GetSeverity

`func (o *CreateSavedFilterRequestFilter) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *CreateSavedFilterRequestFilter) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *CreateSavedFilterRequestFilter) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *CreateSavedFilterRequestFilter) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetHasFix

`func (o *CreateSavedFilterRequestFilter) GetHasFix() string`

GetHasFix returns the HasFix field if non-nil, zero value otherwise.

### GetHasFixOk

`func (o *CreateSavedFilterRequestFilter) GetHasFixOk() (*string, bool)`

GetHasFixOk returns a tuple with the HasFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFix

`func (o *CreateSavedFilterRequestFilter) SetHasFix(v string)`

SetHasFix sets HasFix field to given value.

### HasHasFix

`func (o *CreateSavedFilterRequestFilter) HasHasFix() bool`

HasHasFix returns a boolean if a field has been set.

### GetUpgradeDistance

`func (o *CreateSavedFilterRequestFilter) GetUpgradeDistance() string`

GetUpgradeDistance returns the UpgradeDistance field if non-nil, zero value otherwise.

### GetUpgradeDistanceOk

`func (o *CreateSavedFilterRequestFilter) GetUpgradeDistanceOk() (*string, bool)`

GetUpgradeDistanceOk returns a tuple with the UpgradeDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeDistance

`func (o *CreateSavedFilterRequestFilter) SetUpgradeDistance(v string)`

SetUpgradeDistance sets UpgradeDistance field to given value.

### HasUpgradeDistance

`func (o *CreateSavedFilterRequestFilter) HasUpgradeDistance() bool`

HasUpgradeDistance returns a boolean if a field has been set.

### GetExploitMaturity

`func (o *CreateSavedFilterRequestFilter) GetExploitMaturity() string`

GetExploitMaturity returns the ExploitMaturity field if non-nil, zero value otherwise.

### GetExploitMaturityOk

`func (o *CreateSavedFilterRequestFilter) GetExploitMaturityOk() (*string, bool)`

GetExploitMaturityOk returns a tuple with the ExploitMaturity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitMaturity

`func (o *CreateSavedFilterRequestFilter) SetExploitMaturity(v string)`

SetExploitMaturity sets ExploitMaturity field to given value.

### HasExploitMaturity

`func (o *CreateSavedFilterRequestFilter) HasExploitMaturity() bool`

HasExploitMaturity returns a boolean if a field has been set.

### GetIgnoreReason

`func (o *CreateSavedFilterRequestFilter) GetIgnoreReason() string`

GetIgnoreReason returns the IgnoreReason field if non-nil, zero value otherwise.

### GetIgnoreReasonOk

`func (o *CreateSavedFilterRequestFilter) GetIgnoreReasonOk() (*string, bool)`

GetIgnoreReasonOk returns a tuple with the IgnoreReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreReason

`func (o *CreateSavedFilterRequestFilter) SetIgnoreReason(v string)`

SetIgnoreReason sets IgnoreReason field to given value.

### HasIgnoreReason

`func (o *CreateSavedFilterRequestFilter) HasIgnoreReason() bool`

HasIgnoreReason returns a boolean if a field has been set.

### GetReachability

`func (o *CreateSavedFilterRequestFilter) GetReachability() string`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *CreateSavedFilterRequestFilter) GetReachabilityOk() (*string, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *CreateSavedFilterRequestFilter) SetReachability(v string)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *CreateSavedFilterRequestFilter) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetCwes

`func (o *CreateSavedFilterRequestFilter) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *CreateSavedFilterRequestFilter) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *CreateSavedFilterRequestFilter) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *CreateSavedFilterRequestFilter) HasCwes() bool`

HasCwes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


