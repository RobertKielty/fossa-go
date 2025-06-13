# GetIssues200ResponseIssuesOneOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**GetIssue200ResponseOneOfAllOfSource**](GetIssue200ResponseOneOfAllOfSource.md) |  | [optional] 
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

### NewGetIssues200ResponseIssuesOneOfInner

`func NewGetIssues200ResponseIssuesOneOfInner() *GetIssues200ResponseIssuesOneOfInner`

NewGetIssues200ResponseIssuesOneOfInner instantiates a new GetIssues200ResponseIssuesOneOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssues200ResponseIssuesOneOfInnerWithDefaults

`func NewGetIssues200ResponseIssuesOneOfInnerWithDefaults() *GetIssues200ResponseIssuesOneOfInner`

NewGetIssues200ResponseIssuesOneOfInnerWithDefaults instantiates a new GetIssues200ResponseIssuesOneOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssues200ResponseIssuesOneOfInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssues200ResponseIssuesOneOfInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssues200ResponseIssuesOneOfInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *GetIssues200ResponseIssuesOneOfInner) GetSource() GetIssue200ResponseOneOfAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetIssues200ResponseIssuesOneOfInner) SetSource(v GetIssue200ResponseOneOfAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetIssues200ResponseIssuesOneOfInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepths

`func (o *GetIssues200ResponseIssuesOneOfInner) GetDepths() GetIssue200ResponseOneOfAllOfDepths`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *GetIssues200ResponseIssuesOneOfInner) SetDepths(v GetIssue200ResponseOneOfAllOfDepths)`

SetDepths sets Depths field to given value.

### HasDepths

`func (o *GetIssues200ResponseIssuesOneOfInner) HasDepths() bool`

HasDepths returns a boolean if a field has been set.

### GetContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner) GetStatuses() GetIssueStatuses200ResponseIssues`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner) SetStatuses(v GetIssueStatuses200ResponseIssues)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProjects

`func (o *GetIssues200ResponseIssuesOneOfInner) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetProjectsOk() (*[]GetIssue200ResponseOneOfAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetIssues200ResponseIssuesOneOfInner) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetIssues200ResponseIssuesOneOfInner) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *GetIssues200ResponseIssuesOneOfInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssues200ResponseIssuesOneOfInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssues200ResponseIssuesOneOfInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVulnId

`func (o *GetIssues200ResponseIssuesOneOfInner) GetVulnId() string`

GetVulnId returns the VulnId field if non-nil, zero value otherwise.

### GetVulnIdOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetVulnIdOk() (*string, bool)`

GetVulnIdOk returns a tuple with the VulnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnId

`func (o *GetIssues200ResponseIssuesOneOfInner) SetVulnId(v string)`

SetVulnId sets VulnId field to given value.

### HasVulnId

`func (o *GetIssues200ResponseIssuesOneOfInner) HasVulnId() bool`

HasVulnId returns a boolean if a field has been set.

### GetTitle

`func (o *GetIssues200ResponseIssuesOneOfInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetIssues200ResponseIssuesOneOfInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetIssues200ResponseIssuesOneOfInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCve

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCvss() float32`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCvssOk() (*float32, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCvss(v float32)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvssVector

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetSeverity

`func (o *GetIssues200ResponseIssuesOneOfInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIssues200ResponseIssuesOneOfInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIssues200ResponseIssuesOneOfInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDetails

`func (o *GetIssues200ResponseIssuesOneOfInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetIssues200ResponseIssuesOneOfInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetIssues200ResponseIssuesOneOfInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRemediation

`func (o *GetIssues200ResponseIssuesOneOfInner) GetRemediation() GetIssue200ResponseOneOfAllOfRemediation`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetRemediationOk() (*GetIssue200ResponseOneOfAllOfRemediation, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *GetIssues200ResponseIssuesOneOfInner) SetRemediation(v GetIssue200ResponseOneOfAllOfRemediation)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *GetIssues200ResponseIssuesOneOfInner) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetMetrics

`func (o *GetIssues200ResponseIssuesOneOfInner) GetMetrics() []GetIssue200ResponseOneOfAllOfMetricsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetMetricsOk() (*[]GetIssue200ResponseOneOfAllOfMetricsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetIssues200ResponseIssuesOneOfInner) SetMetrics(v []GetIssue200ResponseOneOfAllOfMetricsInner)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetIssues200ResponseIssuesOneOfInner) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCveStatus

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCveStatus() string`

GetCveStatus returns the CveStatus field if non-nil, zero value otherwise.

### GetCveStatusOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCveStatusOk() (*string, bool)`

GetCveStatusOk returns a tuple with the CveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveStatus

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCveStatus(v string)`

SetCveStatus sets CveStatus field to given value.

### HasCveStatus

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCveStatus() bool`

HasCveStatus returns a boolean if a field has been set.

### GetCwes

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetCpes

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *GetIssues200ResponseIssuesOneOfInner) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *GetIssues200ResponseIssuesOneOfInner) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetPublished

`func (o *GetIssues200ResponseIssuesOneOfInner) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *GetIssues200ResponseIssuesOneOfInner) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *GetIssues200ResponseIssuesOneOfInner) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetAffectedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) GetAffectedVersionRanges() []string`

GetAffectedVersionRanges returns the AffectedVersionRanges field if non-nil, zero value otherwise.

### GetAffectedVersionRangesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetAffectedVersionRangesOk() (*[]string, bool)`

GetAffectedVersionRangesOk returns a tuple with the AffectedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) SetAffectedVersionRanges(v []string)`

SetAffectedVersionRanges sets AffectedVersionRanges field to given value.

### HasAffectedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) HasAffectedVersionRanges() bool`

HasAffectedVersionRanges returns a boolean if a field has been set.

### GetPatchedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) GetPatchedVersionRanges() []string`

GetPatchedVersionRanges returns the PatchedVersionRanges field if non-nil, zero value otherwise.

### GetPatchedVersionRangesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetPatchedVersionRangesOk() (*[]string, bool)`

GetPatchedVersionRangesOk returns a tuple with the PatchedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) SetPatchedVersionRanges(v []string)`

SetPatchedVersionRanges sets PatchedVersionRanges field to given value.

### HasPatchedVersionRanges

`func (o *GetIssues200ResponseIssuesOneOfInner) HasPatchedVersionRanges() bool`

HasPatchedVersionRanges returns a boolean if a field has been set.

### GetReferences

`func (o *GetIssues200ResponseIssuesOneOfInner) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetIssues200ResponseIssuesOneOfInner) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *GetIssues200ResponseIssuesOneOfInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetExploitability

`func (o *GetIssues200ResponseIssuesOneOfInner) GetExploitability() string`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetExploitabilityOk() (*string, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *GetIssues200ResponseIssuesOneOfInner) SetExploitability(v string)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *GetIssues200ResponseIssuesOneOfInner) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetEpss

`func (o *GetIssues200ResponseIssuesOneOfInner) GetEpss() GetIssue200ResponseOneOfAllOfEpss`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *GetIssues200ResponseIssuesOneOfInner) GetEpssOk() (*GetIssue200ResponseOneOfAllOfEpss, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *GetIssues200ResponseIssuesOneOfInner) SetEpss(v GetIssue200ResponseOneOfAllOfEpss)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *GetIssues200ResponseIssuesOneOfInner) HasEpss() bool`

HasEpss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


