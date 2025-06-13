# GetIssue200Response

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
**License** | Pointer to **string** |  | [optional] 
**QualityRule** | Pointer to [**GetIssue200ResponseOneOf2AllOfQualityRule**](GetIssue200ResponseOneOf2AllOfQualityRule.md) |  | [optional] 

## Methods

### NewGetIssue200Response

`func NewGetIssue200Response() *GetIssue200Response`

NewGetIssue200Response instantiates a new GetIssue200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseWithDefaults

`func NewGetIssue200ResponseWithDefaults() *GetIssue200Response`

NewGetIssue200ResponseWithDefaults instantiates a new GetIssue200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssue200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssue200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssue200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssue200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssue200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssue200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssue200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssue200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *GetIssue200Response) GetSource() GetIssue200ResponseOneOfAllOfSource1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetIssue200Response) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetIssue200Response) SetSource(v GetIssue200ResponseOneOfAllOfSource1)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetIssue200Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepths

`func (o *GetIssue200Response) GetDepths() GetIssue200ResponseOneOfAllOfDepths`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *GetIssue200Response) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *GetIssue200Response) SetDepths(v GetIssue200ResponseOneOfAllOfDepths)`

SetDepths sets Depths field to given value.

### HasDepths

`func (o *GetIssue200Response) HasDepths() bool`

HasDepths returns a boolean if a field has been set.

### GetContainerLayers

`func (o *GetIssue200Response) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *GetIssue200Response) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *GetIssue200Response) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *GetIssue200Response) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetStatuses

`func (o *GetIssue200Response) GetStatuses() GetIssueStatuses200ResponseIssues`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetIssue200Response) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetIssue200Response) SetStatuses(v GetIssueStatuses200ResponseIssues)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetIssue200Response) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProjects

`func (o *GetIssue200Response) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetIssue200Response) GetProjectsOk() (*[]GetIssue200ResponseOneOfAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetIssue200Response) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetIssue200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *GetIssue200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssue200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssue200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssue200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVulnId

`func (o *GetIssue200Response) GetVulnId() string`

GetVulnId returns the VulnId field if non-nil, zero value otherwise.

### GetVulnIdOk

`func (o *GetIssue200Response) GetVulnIdOk() (*string, bool)`

GetVulnIdOk returns a tuple with the VulnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnId

`func (o *GetIssue200Response) SetVulnId(v string)`

SetVulnId sets VulnId field to given value.

### HasVulnId

`func (o *GetIssue200Response) HasVulnId() bool`

HasVulnId returns a boolean if a field has been set.

### GetTitle

`func (o *GetIssue200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetIssue200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetIssue200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetIssue200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCve

`func (o *GetIssue200Response) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *GetIssue200Response) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *GetIssue200Response) SetCve(v string)`

SetCve sets Cve field to given value.

### HasCve

`func (o *GetIssue200Response) HasCve() bool`

HasCve returns a boolean if a field has been set.

### GetCvss

`func (o *GetIssue200Response) GetCvss() float32`

GetCvss returns the Cvss field if non-nil, zero value otherwise.

### GetCvssOk

`func (o *GetIssue200Response) GetCvssOk() (*float32, bool)`

GetCvssOk returns a tuple with the Cvss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvss

`func (o *GetIssue200Response) SetCvss(v float32)`

SetCvss sets Cvss field to given value.

### HasCvss

`func (o *GetIssue200Response) HasCvss() bool`

HasCvss returns a boolean if a field has been set.

### GetCvssVector

`func (o *GetIssue200Response) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *GetIssue200Response) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *GetIssue200Response) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *GetIssue200Response) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetSeverity

`func (o *GetIssue200Response) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIssue200Response) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIssue200Response) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIssue200Response) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetDetails

`func (o *GetIssue200Response) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetIssue200Response) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetIssue200Response) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetIssue200Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRemediation

`func (o *GetIssue200Response) GetRemediation() GetIssue200ResponseOneOfAllOfRemediation`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *GetIssue200Response) GetRemediationOk() (*GetIssue200ResponseOneOfAllOfRemediation, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *GetIssue200Response) SetRemediation(v GetIssue200ResponseOneOfAllOfRemediation)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *GetIssue200Response) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.

### GetMetrics

`func (o *GetIssue200Response) GetMetrics() []GetIssue200ResponseOneOfAllOfMetricsInner`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetIssue200Response) GetMetricsOk() (*[]GetIssue200ResponseOneOfAllOfMetricsInner, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetIssue200Response) SetMetrics(v []GetIssue200ResponseOneOfAllOfMetricsInner)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetIssue200Response) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCveStatus

`func (o *GetIssue200Response) GetCveStatus() string`

GetCveStatus returns the CveStatus field if non-nil, zero value otherwise.

### GetCveStatusOk

`func (o *GetIssue200Response) GetCveStatusOk() (*string, bool)`

GetCveStatusOk returns a tuple with the CveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveStatus

`func (o *GetIssue200Response) SetCveStatus(v string)`

SetCveStatus sets CveStatus field to given value.

### HasCveStatus

`func (o *GetIssue200Response) HasCveStatus() bool`

HasCveStatus returns a boolean if a field has been set.

### GetCwes

`func (o *GetIssue200Response) GetCwes() []string`

GetCwes returns the Cwes field if non-nil, zero value otherwise.

### GetCwesOk

`func (o *GetIssue200Response) GetCwesOk() (*[]string, bool)`

GetCwesOk returns a tuple with the Cwes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwes

`func (o *GetIssue200Response) SetCwes(v []string)`

SetCwes sets Cwes field to given value.

### HasCwes

`func (o *GetIssue200Response) HasCwes() bool`

HasCwes returns a boolean if a field has been set.

### GetCpes

`func (o *GetIssue200Response) GetCpes() []string`

GetCpes returns the Cpes field if non-nil, zero value otherwise.

### GetCpesOk

`func (o *GetIssue200Response) GetCpesOk() (*[]string, bool)`

GetCpesOk returns a tuple with the Cpes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpes

`func (o *GetIssue200Response) SetCpes(v []string)`

SetCpes sets Cpes field to given value.

### HasCpes

`func (o *GetIssue200Response) HasCpes() bool`

HasCpes returns a boolean if a field has been set.

### GetPublished

`func (o *GetIssue200Response) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *GetIssue200Response) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *GetIssue200Response) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *GetIssue200Response) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetAffectedVersionRanges

`func (o *GetIssue200Response) GetAffectedVersionRanges() []string`

GetAffectedVersionRanges returns the AffectedVersionRanges field if non-nil, zero value otherwise.

### GetAffectedVersionRangesOk

`func (o *GetIssue200Response) GetAffectedVersionRangesOk() (*[]string, bool)`

GetAffectedVersionRangesOk returns a tuple with the AffectedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedVersionRanges

`func (o *GetIssue200Response) SetAffectedVersionRanges(v []string)`

SetAffectedVersionRanges sets AffectedVersionRanges field to given value.

### HasAffectedVersionRanges

`func (o *GetIssue200Response) HasAffectedVersionRanges() bool`

HasAffectedVersionRanges returns a boolean if a field has been set.

### GetPatchedVersionRanges

`func (o *GetIssue200Response) GetPatchedVersionRanges() []string`

GetPatchedVersionRanges returns the PatchedVersionRanges field if non-nil, zero value otherwise.

### GetPatchedVersionRangesOk

`func (o *GetIssue200Response) GetPatchedVersionRangesOk() (*[]string, bool)`

GetPatchedVersionRangesOk returns a tuple with the PatchedVersionRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchedVersionRanges

`func (o *GetIssue200Response) SetPatchedVersionRanges(v []string)`

SetPatchedVersionRanges sets PatchedVersionRanges field to given value.

### HasPatchedVersionRanges

`func (o *GetIssue200Response) HasPatchedVersionRanges() bool`

HasPatchedVersionRanges returns a boolean if a field has been set.

### GetReferences

`func (o *GetIssue200Response) GetReferences() []string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *GetIssue200Response) GetReferencesOk() (*[]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *GetIssue200Response) SetReferences(v []string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *GetIssue200Response) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetExploitability

`func (o *GetIssue200Response) GetExploitability() string`

GetExploitability returns the Exploitability field if non-nil, zero value otherwise.

### GetExploitabilityOk

`func (o *GetIssue200Response) GetExploitabilityOk() (*string, bool)`

GetExploitabilityOk returns a tuple with the Exploitability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExploitability

`func (o *GetIssue200Response) SetExploitability(v string)`

SetExploitability sets Exploitability field to given value.

### HasExploitability

`func (o *GetIssue200Response) HasExploitability() bool`

HasExploitability returns a boolean if a field has been set.

### GetEpss

`func (o *GetIssue200Response) GetEpss() GetIssue200ResponseOneOfAllOfEpss`

GetEpss returns the Epss field if non-nil, zero value otherwise.

### GetEpssOk

`func (o *GetIssue200Response) GetEpssOk() (*GetIssue200ResponseOneOfAllOfEpss, bool)`

GetEpssOk returns a tuple with the Epss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpss

`func (o *GetIssue200Response) SetEpss(v GetIssue200ResponseOneOfAllOfEpss)`

SetEpss sets Epss field to given value.

### HasEpss

`func (o *GetIssue200Response) HasEpss() bool`

HasEpss returns a boolean if a field has been set.

### GetLicense

`func (o *GetIssue200Response) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetIssue200Response) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetIssue200Response) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *GetIssue200Response) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetQualityRule

`func (o *GetIssue200Response) GetQualityRule() GetIssue200ResponseOneOf2AllOfQualityRule`

GetQualityRule returns the QualityRule field if non-nil, zero value otherwise.

### GetQualityRuleOk

`func (o *GetIssue200Response) GetQualityRuleOk() (*GetIssue200ResponseOneOf2AllOfQualityRule, bool)`

GetQualityRuleOk returns a tuple with the QualityRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRule

`func (o *GetIssue200Response) SetQualityRule(v GetIssue200ResponseOneOf2AllOfQualityRule)`

SetQualityRule sets QualityRule field to given value.

### HasQualityRule

`func (o *GetIssue200Response) HasQualityRule() bool`

HasQualityRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


