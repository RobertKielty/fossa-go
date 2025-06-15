# GetIssue200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**GetIssue200ResponseOneOfAllOfSource1**](GetIssue200ResponseOneOfAllOfSource1.md) |  | [optional] 
**Depths** | Pointer to [**GetIssue200ResponseOneOfAllOfDepths**](GetIssue200ResponseOneOfAllOfDepths.md) |  | [optional] 
**ContainerLayers** | Pointer to [**GetIssue200ResponseOneOfAllOfContainerLayers**](GetIssue200ResponseOneOfAllOfContainerLayers.md) |  | [optional] 
**Statuses** | Pointer to [**GetIssueStatuses200ResponseIssues**](GetIssueStatuses200ResponseIssues.md) |  | [optional] 
**Projects** | Pointer to [**[]GetIssue200ResponseOneOfAllOfProjectsInner**](GetIssue200ResponseOneOfAllOfProjectsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VulnId** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Cve** | Pointer to **string** |  | [optional] 
**Cvss** | Pointer to **float32** |  | [optional] 
**CvssVector** | Pointer to **string** | The CVSS Vector for the vuln. Prefers the V3 vector if it exists, and falls back to the V2 vector. Is null if neither vector exists. | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Remediation** | Pointer to [**GetIssue200ResponseOneOfAllOfRemediation**](GetIssue200ResponseOneOfAllOfRemediation.md) |  | [optional] 
**Metrics** | Pointer to [**[]GetIssue200ResponseOneOfAllOfMetricsInner**](GetIssue200ResponseOneOfAllOfMetricsInner.md) |  | [optional] 
**CveStatus** | Pointer to **string** |  | [optional] 
**Cwes** | Pointer to **[]string** |  | [optional] 
**Cpes** | Pointer to **[]string** |  | [optional] 
**Published** | Pointer to **time.Time** |  | [optional] 
**AffectedVersionRanges** | Pointer to **[]string** |  | [optional] 
**PatchedVersionRanges** | Pointer to **[]string** |  | [optional] 
**References** | Pointer to **[]string** |  | [optional] 
**Exploitability** | Pointer to **string** |  | [optional] 
**Epss** | Pointer to [**GetIssue200ResponseOneOfAllOfEpss**](GetIssue200ResponseOneOfAllOfEpss.md) |  | [optional] 

## Methods

### NewGetIssue200ResponseOneOf

`func NewGetIssue200ResponseOneOf() *GetIssue200ResponseOneOf`

NewGetIssue200ResponseOneOf instantiates a new GetIssue200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseOneOfWithDefaults

`func NewGetIssue200ResponseOneOfWithDefaults() *GetIssue200ResponseOneOf`

NewGetIssue200ResponseOneOfWithDefaults instantiates a new GetIssue200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssue200ResponseOneOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssue200ResponseOneOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssue200ResponseOneOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssue200ResponseOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssue200ResponseOneOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssue200ResponseOneOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssue200ResponseOneOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssue200ResponseOneOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *GetIssue200ResponseOneOf) GetSource() GetIssue200ResponseOneOfAllOfSource1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetIssue200ResponseOneOf) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetIssue200ResponseOneOf) SetSource(v GetIssue200ResponseOneOfAllOfSource1)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetIssue200ResponseOneOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepths

`func (o *GetIssue200ResponseOneOf) GetDepths() GetIssue200ResponseOneOfAllOfDepths`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *GetIssue200ResponseOneOf) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *GetIssue200ResponseOneOf) SetDepths(v GetIssue200ResponseOneOfAllOfDepths)`

SetDepths sets Depths field to given value.

### HasDepths

`func (o *GetIssue200ResponseOneOf) HasDepths() bool`

HasDepths returns a boolean if a field has been set.

### GetContainerLayers

`func (o *GetIssue200ResponseOneOf) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *GetIssue200ResponseOneOf) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *GetIssue200ResponseOneOf) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *GetIssue200ResponseOneOf) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetStatuses

`func (o *GetIssue200ResponseOneOf) GetStatuses() GetIssueStatuses200ResponseIssues`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetIssue200ResponseOneOf) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetIssue200ResponseOneOf) SetStatuses(v GetIssueStatuses200ResponseIssues)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetIssue200ResponseOneOf) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProjects

`func (o *GetIssue200ResponseOneOf) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetIssue200ResponseOneOf) GetProjectsOk() (*[]GetIssue200ResponseOneOfAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetIssue200ResponseOneOf) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetIssue200ResponseOneOf) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *GetIssue200ResponseOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssue200ResponseOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssue200ResponseOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssue200ResponseOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVulnId

`func (o *GetIssue200ResponseOneOf) GetVulnId() string`

GetVulnId returns the VulnId field if non-nil, zero value otherwise.

### GetVulnIdOk

`func (o *GetIssue200ResponseOneOf) GetVulnIdOk() (*string, bool)`

GetVulnIdOk returns a tuple with the VulnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnId

`func (o *GetIssue200ResponseOneOf) SetVulnId(v string)`

SetVulnId sets VulnId field to given value.

### HasVulnId

`func (o *GetIssue200ResponseOneOf) HasVulnId() bool`

HasVulnId returns a boolean if a field has been set.

### GetTitle

`func (o *GetIssue200ResponseOneOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetIssue200ResponseOneOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetIssue200ResponseOneOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetIssue200ResponseOneOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCve

`func (o *GetIssue200ResponseOneOf) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *GetIssue200ResponseOneOf) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *GetIssue200ResponseOneOf) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *GetIssue200ResponseOneOf) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *GetIssue200ResponseOneOf) GetCvss() float32`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *GetIssue200ResponseOneOf) GetCvssOk() (*float32, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *GetIssue200ResponseOneOf) SetCvss(v float32)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *GetIssue200ResponseOneOf) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvssVector

`func (o *GetIssue200ResponseOneOf) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *GetIssue200ResponseOneOf) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *GetIssue200ResponseOneOf) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *GetIssue200ResponseOneOf) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetSeverity

`func (o *GetIssue200ResponseOneOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIssue200ResponseOneOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIssue200ResponseOneOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIssue200ResponseOneOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDetails

`func (o *GetIssue200ResponseOneOf) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetIssue200ResponseOneOf) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetIssue200ResponseOneOf) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetIssue200ResponseOneOf) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRemediation

`func (o *GetIssue200ResponseOneOf) GetRemediation() GetIssue200ResponseOneOfAllOfRemediation`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *GetIssue200ResponseOneOf) GetRemediationOk() (*GetIssue200ResponseOneOfAllOfRemediation, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *GetIssue200ResponseOneOf) SetRemediation(v GetIssue200ResponseOneOfAllOfRemediation)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *GetIssue200ResponseOneOf) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetMetrics

`func (o *GetIssue200ResponseOneOf) GetMetrics() []GetIssue200ResponseOneOfAllOfMetricsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetIssue200ResponseOneOf) GetMetricsOk() (*[]GetIssue200ResponseOneOfAllOfMetricsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetIssue200ResponseOneOf) SetMetrics(v []GetIssue200ResponseOneOfAllOfMetricsInner)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetIssue200ResponseOneOf) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCveStatus

`func (o *GetIssue200ResponseOneOf) GetCveStatus() string`

GetCveStatus returns the CveStatus field if non-nil, zero value otherwise.

### GetCveStatusOk

`func (o *GetIssue200ResponseOneOf) GetCveStatusOk() (*string, bool)`

GetCveStatusOk returns a tuple with the CveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveStatus

`func (o *GetIssue200ResponseOneOf) SetCveStatus(v string)`

SetCveStatus sets CveStatus field to given value.

### HasCveStatus

`func (o *GetIssue200ResponseOneOf) HasCveStatus() bool`

HasCveStatus returns a boolean if a field has been set.

### GetCwes

`func (o *GetIssue200ResponseOneOf) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *GetIssue200ResponseOneOf) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *GetIssue200ResponseOneOf) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *GetIssue200ResponseOneOf) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetCpes

`func (o *GetIssue200ResponseOneOf) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *GetIssue200ResponseOneOf) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *GetIssue200ResponseOneOf) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *GetIssue200ResponseOneOf) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetPublished

`func (o *GetIssue200ResponseOneOf) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *GetIssue200ResponseOneOf) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *GetIssue200ResponseOneOf) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *GetIssue200ResponseOneOf) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetAffectedVersionRanges

`func (o *GetIssue200ResponseOneOf) GetAffectedVersionRanges() []string`

GetAffectedVersionRanges returns the AffectedVersionRanges field if non-nil, zero value otherwise.

### GetAffectedVersionRangesOk

`func (o *GetIssue200ResponseOneOf) GetAffectedVersionRangesOk() (*[]string, bool)`

GetAffectedVersionRangesOk returns a tuple with the AffectedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersionRanges

`func (o *GetIssue200ResponseOneOf) SetAffectedVersionRanges(v []string)`

SetAffectedVersionRanges sets AffectedVersionRanges field to given value.

### HasAffectedVersionRanges

`func (o *GetIssue200ResponseOneOf) HasAffectedVersionRanges() bool`

HasAffectedVersionRanges returns a boolean if a field has been set.

### GetPatchedVersionRanges

`func (o *GetIssue200ResponseOneOf) GetPatchedVersionRanges() []string`

GetPatchedVersionRanges returns the PatchedVersionRanges field if non-nil, zero value otherwise.

### GetPatchedVersionRangesOk

`func (o *GetIssue200ResponseOneOf) GetPatchedVersionRangesOk() (*[]string, bool)`

GetPatchedVersionRangesOk returns a tuple with the PatchedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchedVersionRanges

`func (o *GetIssue200ResponseOneOf) SetPatchedVersionRanges(v []string)`

SetPatchedVersionRanges sets PatchedVersionRanges field to given value.

### HasPatchedVersionRanges

`func (o *GetIssue200ResponseOneOf) HasPatchedVersionRanges() bool`

HasPatchedVersionRanges returns a boolean if a field has been set.

### GetReferences

`func (o *GetIssue200ResponseOneOf) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetIssue200ResponseOneOf) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetIssue200ResponseOneOf) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *GetIssue200ResponseOneOf) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetExploitability

`func (o *GetIssue200ResponseOneOf) GetExploitability() string`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *GetIssue200ResponseOneOf) GetExploitabilityOk() (*string, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *GetIssue200ResponseOneOf) SetExploitability(v string)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *GetIssue200ResponseOneOf) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetEpss

`func (o *GetIssue200ResponseOneOf) GetEpss() GetIssue200ResponseOneOfAllOfEpss`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *GetIssue200ResponseOneOf) GetEpssOk() (*GetIssue200ResponseOneOfAllOfEpss, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *GetIssue200ResponseOneOf) SetEpss(v GetIssue200ResponseOneOfAllOfEpss)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *GetIssue200ResponseOneOf) HasEpss() bool`

HasEpss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


