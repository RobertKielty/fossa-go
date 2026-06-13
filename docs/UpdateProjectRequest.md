# UpdateProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Display name of the project | [optional] 
**Description** | Pointer to **string** | Detailed description of the project | [optional] 
**Url** | Pointer to **string** | Project homepage URL | [optional] 
**Notes** | Pointer to **string** | Internal notes about the project | [optional] 
**Public** | Pointer to **bool** | Whether the project is publicly accessible. Requires MakePublic permission. Public projects can be viewed by anyone with the link.  | [optional] 
**ScmUrl** | Pointer to **string** | Source control management URL for the project repository | [optional] 
**ManualScmUrl** | Pointer to **bool** | Whether the SCM URL was manually provided by the user | [optional] 
**VcsHost** | Pointer to **string** | VCS hosting provider | [optional] 
**DefaultBranch** | Pointer to **string** | The default branch to analyze for this project | [optional] 
**TrackingBranches** | Pointer to **[]string** | Branches to actively track and analyze. Tags are automatically filtered out. Branches in this list are automatically removed from &#x60;hidden_branches&#x60;.  | [optional] 
**HiddenBranches** | Pointer to **[]string** | Branches to hide from the UI. Tags are automatically filtered out. Branches in this list are automatically removed from &#x60;tracking_branches&#x60;.  | [optional] 
**PolicyId** | Pointer to **int32** | ID of the licensing policy to apply to this project. Requires SetPolicy permission for LICENSING policy type. | [optional] 
**SecurityPolicyId** | Pointer to **int32** | ID of the security policy to apply to this project. Requires SetPolicy permission for SECURITY policy type. | [optional] 
**QualityPolicyId** | Pointer to **int32** | ID of the quality policy to apply to this project. Requires SetPolicy permission for QUALITY policy type. | [optional] 
**SbomPolicyId** | Pointer to **int32** | ID of the SBOM policy to apply to this project. Requires SetPolicy permission for SBOM policy type. | [optional] 
**PoliciesApproveMultilicense** | Pointer to **bool** | Whether to automatically approve dependencies with multiple licenses if any license is approved | [optional] 
**LicensingIssueScanningEnabled** | Pointer to **bool** | Enable or disable licensing issue scanning for this project | [optional] 
**SecurityIssueScanningEnabled** | Pointer to **bool** | Enable or disable security vulnerability scanning for this project. Can only be modified if the organization has security features enabled.  | [optional] 
**QualityIssueScanningEnabled** | Pointer to **bool** | Enable or disable quality issue scanning for this project. Can only be modified if the organization has quality features enabled.  | [optional] 
**SbomAnalysisEnabled** | Pointer to **bool** | Enable or disable SBOM policy analysis for this project. Requires SetPolicy permission for SBOM policy type. | [optional] 
**SnippetLicensingIssueScanningEnabled** | Pointer to **bool** | Enable or disable snippet licensing issue scanning for this project | [optional] 
**SnippetSecurityIssueScanningEnabled** | Pointer to **bool** | Enable or disable snippet security issue scanning for this project | [optional] 
**SnippetAutoRejectMatchPercentageThreshold** | Pointer to **int32** | Match percentage at (or above) which snippet matches are automatically rejected. Must be an integer between 1 and 100, or &#x60;null&#x60; to disable auto-rejection.  | [optional] 
**VendoredLicensingIssueScanningEnabled** | Pointer to **bool** | Enable or disable licensing issue scanning for vendored dependencies in this project | [optional] 
**VendoredSecurityIssueScanningEnabled** | Pointer to **bool** | Enable or disable security issue scanning for vendored dependencies in this project | [optional] 
**VendoredQualityIssueScanningEnabled** | Pointer to **bool** | Enable or disable quality issue scanning for vendored dependencies in this project | [optional] 
**QuickImportVendoredDetectionEnabled** | Pointer to **bool** | Enable or disable vendored dependency detection during quick imports for this project. Can only be modified if the organization has quick import vendored detection and vendored dependency detection features enabled.  | [optional] 
**QuickImportSnippetDetectionEnabled** | Pointer to **bool** | Enable or disable snippet detection during quick imports for this project. Can only be modified if the organization has quick import snippet detection and snippet detection features enabled.  | [optional] 
**DefaultSavedOptionLicensingReport** | Pointer to **int32** | ID of the saved report option to use as the default for licensing reports. Must reference a report option belonging to your organization, or &#x60;null&#x60; to clear it.  | [optional] 
**DefaultSavedOptionSbom** | Pointer to **int32** | ID of the saved report option to use as the default for SBOM reports. Must reference a report option belonging to your organization, or &#x60;null&#x60; to clear it.  | [optional] 
**InvalidCredential** | Pointer to **bool** | Whether the credential associated with this project is currently invalid. | [optional] 
**LicensingStatusCheckEnabled** | Pointer to **bool** | Enable or disable licensing issue CI/CD status checks | [optional] 
**SecurityStatusCheckEnabled** | Pointer to **bool** | Enable or disable security issue CI/CD status checks | [optional] 
**QualityStatusCheckEnabled** | Pointer to **bool** | Enable or disable quality issue CI/CD status checks | [optional] 
**ExcludeBaseLayerIssuesLicensing** | Pointer to **bool** | Exclude licensing issues found in container base layers | [optional] 
**ExcludeBaseLayerIssuesSecurity** | Pointer to **bool** | Exclude security issues found in container base layers | [optional] 
**ExcludeBaseLayerIssuesQuality** | Pointer to **bool** | Exclude quality issues found in container base layers | [optional] 
**IntegrationhookTimeout** | Pointer to **int32** | Timeout in seconds for integration hooks (e.g., GitHub status checks) | [optional] 
**IntegrationhookFailState** | Pointer to **string** | Status to report when a hook times out or fails | [optional] 
**IssueTrackerUrl** | Pointer to **string** | URL of the external issue tracker (e.g., Jira, GitHub Issues) | [optional] 
**IssueTrackerType** | Pointer to **string** | Type of issue tracker | [optional] 
**IssueTrackerLabels** | Pointer to **[]string** | Labels to automatically apply to issues created in the tracker | [optional] 
**IssueTrackerIssueTypes** | Pointer to **[]string** | Jira issue types available for this project | [optional] 
**IssueTrackerProjectIds** | Pointer to **[]string** | Jira project IDs associated with this project | [optional] 
**IssueTrackerCustomFields** | Pointer to [**map[string]UpdateProjectRequestIssueTrackerCustomFieldsValue**](UpdateProjectRequestIssueTrackerCustomFieldsValue.md) | Custom Jira fields configuration. The object keys are Jira field IDs, and values contain field metadata. The &#x60;isRequired&#x60; field accepts both boolean and stringified boolean values (\&quot;true\&quot;/\&quot;false\&quot;).  | [optional] 
**UseGlobalTrackerSettings** | Pointer to **bool** | Whether to use organization-level issue tracker settings instead of project-specific settings | [optional] 
**TransitiveExcludes** | Pointer to **[]string** | List of dependency locators to exclude from analysis. Removing items from this array is logged as \&quot;un-ignoring\&quot; dependencies. Format: \&quot;fetcher+package$revision\&quot; (e.g., \&quot;npm+lodash$4.17.21\&quot;)  | [optional] 
**ReportCustomText** | Pointer to **string** | Custom text to include in attribution reports for this project | [optional] 
**BomColumnSettings** | Pointer to **[]string** | Columns to display in the Bill of Materials (BOM) report. Available options: All, Name, Version, Type, License, DirectLicense, DirectLicenseOrigin, DeepLicense, DeepLicenseOrigin, Description, Homepage, PrimaryLanguage, SourceLocation, ReleasePublishDate, OriginId, Tags, ComponentComment  | [optional] 
**BomPublicId** | Pointer to **string** | Public identifier for accessing the project&#39;s attribution report | [optional] 
**Labels** | Pointer to **[]int32** | Array of label IDs to associate with this project. This replaces all existing labels. Labels must exist in the organization before being assigned.  | [optional] 
**Filters** | Pointer to [**UpdateProjectRequestFilters**](UpdateProjectRequestFilters.md) |  | [optional] 
**Notifications** | Pointer to [**[]UpdateProjectRequestNotificationsInner**](UpdateProjectRequestNotificationsInner.md) | Array of notification configurations for this project. Empty notification objects are automatically filtered out. Changes are processed asynchronously alongside the main update.  | [optional] 

## Methods

### NewUpdateProjectRequest

`func NewUpdateProjectRequest() *UpdateProjectRequest`

NewUpdateProjectRequest instantiates a new UpdateProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestWithDefaults

`func NewUpdateProjectRequestWithDefaults() *UpdateProjectRequest`

NewUpdateProjectRequestWithDefaults instantiates a new UpdateProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateProjectRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateProjectRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateProjectRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateProjectRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateProjectRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateProjectRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateProjectRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateProjectRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateProjectRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateProjectRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateProjectRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateProjectRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPublic

`func (o *UpdateProjectRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *UpdateProjectRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *UpdateProjectRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *UpdateProjectRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetScmUrl

`func (o *UpdateProjectRequest) GetScmUrl() string`

GetScmUrl returns the ScmUrl field if non-nil, zero value otherwise.

### GetScmUrlOk

`func (o *UpdateProjectRequest) GetScmUrlOk() (*string, bool)`

GetScmUrlOk returns a tuple with the ScmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScmUrl

`func (o *UpdateProjectRequest) SetScmUrl(v string)`

SetScmUrl sets ScmUrl field to given value.

### HasScmUrl

`func (o *UpdateProjectRequest) HasScmUrl() bool`

HasScmUrl returns a boolean if a field has been set.

### GetManualScmUrl

`func (o *UpdateProjectRequest) GetManualScmUrl() bool`

GetManualScmUrl returns the ManualScmUrl field if non-nil, zero value otherwise.

### GetManualScmUrlOk

`func (o *UpdateProjectRequest) GetManualScmUrlOk() (*bool, bool)`

GetManualScmUrlOk returns a tuple with the ManualScmUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualScmUrl

`func (o *UpdateProjectRequest) SetManualScmUrl(v bool)`

SetManualScmUrl sets ManualScmUrl field to given value.

### HasManualScmUrl

`func (o *UpdateProjectRequest) HasManualScmUrl() bool`

HasManualScmUrl returns a boolean if a field has been set.

### GetVcsHost

`func (o *UpdateProjectRequest) GetVcsHost() string`

GetVcsHost returns the VcsHost field if non-nil, zero value otherwise.

### GetVcsHostOk

`func (o *UpdateProjectRequest) GetVcsHostOk() (*string, bool)`

GetVcsHostOk returns a tuple with the VcsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsHost

`func (o *UpdateProjectRequest) SetVcsHost(v string)`

SetVcsHost sets VcsHost field to given value.

### HasVcsHost

`func (o *UpdateProjectRequest) HasVcsHost() bool`

HasVcsHost returns a boolean if a field has been set.

### GetDefaultBranch

`func (o *UpdateProjectRequest) GetDefaultBranch() string`

GetDefaultBranch returns the DefaultBranch field if non-nil, zero value otherwise.

### GetDefaultBranchOk

`func (o *UpdateProjectRequest) GetDefaultBranchOk() (*string, bool)`

GetDefaultBranchOk returns a tuple with the DefaultBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranch

`func (o *UpdateProjectRequest) SetDefaultBranch(v string)`

SetDefaultBranch sets DefaultBranch field to given value.

### HasDefaultBranch

`func (o *UpdateProjectRequest) HasDefaultBranch() bool`

HasDefaultBranch returns a boolean if a field has been set.

### GetTrackingBranches

`func (o *UpdateProjectRequest) GetTrackingBranches() []string`

GetTrackingBranches returns the TrackingBranches field if non-nil, zero value otherwise.

### GetTrackingBranchesOk

`func (o *UpdateProjectRequest) GetTrackingBranchesOk() (*[]string, bool)`

GetTrackingBranchesOk returns a tuple with the TrackingBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingBranches

`func (o *UpdateProjectRequest) SetTrackingBranches(v []string)`

SetTrackingBranches sets TrackingBranches field to given value.

### HasTrackingBranches

`func (o *UpdateProjectRequest) HasTrackingBranches() bool`

HasTrackingBranches returns a boolean if a field has been set.

### GetHiddenBranches

`func (o *UpdateProjectRequest) GetHiddenBranches() []string`

GetHiddenBranches returns the HiddenBranches field if non-nil, zero value otherwise.

### GetHiddenBranchesOk

`func (o *UpdateProjectRequest) GetHiddenBranchesOk() (*[]string, bool)`

GetHiddenBranchesOk returns a tuple with the HiddenBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenBranches

`func (o *UpdateProjectRequest) SetHiddenBranches(v []string)`

SetHiddenBranches sets HiddenBranches field to given value.

### HasHiddenBranches

`func (o *UpdateProjectRequest) HasHiddenBranches() bool`

HasHiddenBranches returns a boolean if a field has been set.

### GetPolicyId

`func (o *UpdateProjectRequest) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *UpdateProjectRequest) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *UpdateProjectRequest) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *UpdateProjectRequest) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetSecurityPolicyId

`func (o *UpdateProjectRequest) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *UpdateProjectRequest) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *UpdateProjectRequest) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.

### HasSecurityPolicyId

`func (o *UpdateProjectRequest) HasSecurityPolicyId() bool`

HasSecurityPolicyId returns a boolean if a field has been set.

### GetQualityPolicyId

`func (o *UpdateProjectRequest) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *UpdateProjectRequest) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *UpdateProjectRequest) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.

### HasQualityPolicyId

`func (o *UpdateProjectRequest) HasQualityPolicyId() bool`

HasQualityPolicyId returns a boolean if a field has been set.

### GetSbomPolicyId

`func (o *UpdateProjectRequest) GetSbomPolicyId() int32`

GetSbomPolicyId returns the SbomPolicyId field if non-nil, zero value otherwise.

### GetSbomPolicyIdOk

`func (o *UpdateProjectRequest) GetSbomPolicyIdOk() (*int32, bool)`

GetSbomPolicyIdOk returns a tuple with the SbomPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbomPolicyId

`func (o *UpdateProjectRequest) SetSbomPolicyId(v int32)`

SetSbomPolicyId sets SbomPolicyId field to given value.

### HasSbomPolicyId

`func (o *UpdateProjectRequest) HasSbomPolicyId() bool`

HasSbomPolicyId returns a boolean if a field has been set.

### GetPoliciesApproveMultilicense

`func (o *UpdateProjectRequest) GetPoliciesApproveMultilicense() bool`

GetPoliciesApproveMultilicense returns the PoliciesApproveMultilicense field if non-nil, zero value otherwise.

### GetPoliciesApproveMultilicenseOk

`func (o *UpdateProjectRequest) GetPoliciesApproveMultilicenseOk() (*bool, bool)`

GetPoliciesApproveMultilicenseOk returns a tuple with the PoliciesApproveMultilicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoliciesApproveMultilicense

`func (o *UpdateProjectRequest) SetPoliciesApproveMultilicense(v bool)`

SetPoliciesApproveMultilicense sets PoliciesApproveMultilicense field to given value.

### HasPoliciesApproveMultilicense

`func (o *UpdateProjectRequest) HasPoliciesApproveMultilicense() bool`

HasPoliciesApproveMultilicense returns a boolean if a field has been set.

### GetLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) GetLicensingIssueScanningEnabled() bool`

GetLicensingIssueScanningEnabled returns the LicensingIssueScanningEnabled field if non-nil, zero value otherwise.

### GetLicensingIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetLicensingIssueScanningEnabledOk() (*bool, bool)`

GetLicensingIssueScanningEnabledOk returns a tuple with the LicensingIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) SetLicensingIssueScanningEnabled(v bool)`

SetLicensingIssueScanningEnabled sets LicensingIssueScanningEnabled field to given value.

### HasLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) HasLicensingIssueScanningEnabled() bool`

HasLicensingIssueScanningEnabled returns a boolean if a field has been set.

### GetSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) GetSecurityIssueScanningEnabled() bool`

GetSecurityIssueScanningEnabled returns the SecurityIssueScanningEnabled field if non-nil, zero value otherwise.

### GetSecurityIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetSecurityIssueScanningEnabledOk() (*bool, bool)`

GetSecurityIssueScanningEnabledOk returns a tuple with the SecurityIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) SetSecurityIssueScanningEnabled(v bool)`

SetSecurityIssueScanningEnabled sets SecurityIssueScanningEnabled field to given value.

### HasSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) HasSecurityIssueScanningEnabled() bool`

HasSecurityIssueScanningEnabled returns a boolean if a field has been set.

### GetQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) GetQualityIssueScanningEnabled() bool`

GetQualityIssueScanningEnabled returns the QualityIssueScanningEnabled field if non-nil, zero value otherwise.

### GetQualityIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetQualityIssueScanningEnabledOk() (*bool, bool)`

GetQualityIssueScanningEnabledOk returns a tuple with the QualityIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) SetQualityIssueScanningEnabled(v bool)`

SetQualityIssueScanningEnabled sets QualityIssueScanningEnabled field to given value.

### HasQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) HasQualityIssueScanningEnabled() bool`

HasQualityIssueScanningEnabled returns a boolean if a field has been set.

### GetSbomAnalysisEnabled

`func (o *UpdateProjectRequest) GetSbomAnalysisEnabled() bool`

GetSbomAnalysisEnabled returns the SbomAnalysisEnabled field if non-nil, zero value otherwise.

### GetSbomAnalysisEnabledOk

`func (o *UpdateProjectRequest) GetSbomAnalysisEnabledOk() (*bool, bool)`

GetSbomAnalysisEnabledOk returns a tuple with the SbomAnalysisEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbomAnalysisEnabled

`func (o *UpdateProjectRequest) SetSbomAnalysisEnabled(v bool)`

SetSbomAnalysisEnabled sets SbomAnalysisEnabled field to given value.

### HasSbomAnalysisEnabled

`func (o *UpdateProjectRequest) HasSbomAnalysisEnabled() bool`

HasSbomAnalysisEnabled returns a boolean if a field has been set.

### GetSnippetLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) GetSnippetLicensingIssueScanningEnabled() bool`

GetSnippetLicensingIssueScanningEnabled returns the SnippetLicensingIssueScanningEnabled field if non-nil, zero value otherwise.

### GetSnippetLicensingIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetSnippetLicensingIssueScanningEnabledOk() (*bool, bool)`

GetSnippetLicensingIssueScanningEnabledOk returns a tuple with the SnippetLicensingIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) SetSnippetLicensingIssueScanningEnabled(v bool)`

SetSnippetLicensingIssueScanningEnabled sets SnippetLicensingIssueScanningEnabled field to given value.

### HasSnippetLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) HasSnippetLicensingIssueScanningEnabled() bool`

HasSnippetLicensingIssueScanningEnabled returns a boolean if a field has been set.

### GetSnippetSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) GetSnippetSecurityIssueScanningEnabled() bool`

GetSnippetSecurityIssueScanningEnabled returns the SnippetSecurityIssueScanningEnabled field if non-nil, zero value otherwise.

### GetSnippetSecurityIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetSnippetSecurityIssueScanningEnabledOk() (*bool, bool)`

GetSnippetSecurityIssueScanningEnabledOk returns a tuple with the SnippetSecurityIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) SetSnippetSecurityIssueScanningEnabled(v bool)`

SetSnippetSecurityIssueScanningEnabled sets SnippetSecurityIssueScanningEnabled field to given value.

### HasSnippetSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) HasSnippetSecurityIssueScanningEnabled() bool`

HasSnippetSecurityIssueScanningEnabled returns a boolean if a field has been set.

### GetSnippetAutoRejectMatchPercentageThreshold

`func (o *UpdateProjectRequest) GetSnippetAutoRejectMatchPercentageThreshold() int32`

GetSnippetAutoRejectMatchPercentageThreshold returns the SnippetAutoRejectMatchPercentageThreshold field if non-nil, zero value otherwise.

### GetSnippetAutoRejectMatchPercentageThresholdOk

`func (o *UpdateProjectRequest) GetSnippetAutoRejectMatchPercentageThresholdOk() (*int32, bool)`

GetSnippetAutoRejectMatchPercentageThresholdOk returns a tuple with the SnippetAutoRejectMatchPercentageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetAutoRejectMatchPercentageThreshold

`func (o *UpdateProjectRequest) SetSnippetAutoRejectMatchPercentageThreshold(v int32)`

SetSnippetAutoRejectMatchPercentageThreshold sets SnippetAutoRejectMatchPercentageThreshold field to given value.

### HasSnippetAutoRejectMatchPercentageThreshold

`func (o *UpdateProjectRequest) HasSnippetAutoRejectMatchPercentageThreshold() bool`

HasSnippetAutoRejectMatchPercentageThreshold returns a boolean if a field has been set.

### GetVendoredLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) GetVendoredLicensingIssueScanningEnabled() bool`

GetVendoredLicensingIssueScanningEnabled returns the VendoredLicensingIssueScanningEnabled field if non-nil, zero value otherwise.

### GetVendoredLicensingIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetVendoredLicensingIssueScanningEnabledOk() (*bool, bool)`

GetVendoredLicensingIssueScanningEnabledOk returns a tuple with the VendoredLicensingIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendoredLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) SetVendoredLicensingIssueScanningEnabled(v bool)`

SetVendoredLicensingIssueScanningEnabled sets VendoredLicensingIssueScanningEnabled field to given value.

### HasVendoredLicensingIssueScanningEnabled

`func (o *UpdateProjectRequest) HasVendoredLicensingIssueScanningEnabled() bool`

HasVendoredLicensingIssueScanningEnabled returns a boolean if a field has been set.

### GetVendoredSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) GetVendoredSecurityIssueScanningEnabled() bool`

GetVendoredSecurityIssueScanningEnabled returns the VendoredSecurityIssueScanningEnabled field if non-nil, zero value otherwise.

### GetVendoredSecurityIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetVendoredSecurityIssueScanningEnabledOk() (*bool, bool)`

GetVendoredSecurityIssueScanningEnabledOk returns a tuple with the VendoredSecurityIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendoredSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) SetVendoredSecurityIssueScanningEnabled(v bool)`

SetVendoredSecurityIssueScanningEnabled sets VendoredSecurityIssueScanningEnabled field to given value.

### HasVendoredSecurityIssueScanningEnabled

`func (o *UpdateProjectRequest) HasVendoredSecurityIssueScanningEnabled() bool`

HasVendoredSecurityIssueScanningEnabled returns a boolean if a field has been set.

### GetVendoredQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) GetVendoredQualityIssueScanningEnabled() bool`

GetVendoredQualityIssueScanningEnabled returns the VendoredQualityIssueScanningEnabled field if non-nil, zero value otherwise.

### GetVendoredQualityIssueScanningEnabledOk

`func (o *UpdateProjectRequest) GetVendoredQualityIssueScanningEnabledOk() (*bool, bool)`

GetVendoredQualityIssueScanningEnabledOk returns a tuple with the VendoredQualityIssueScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendoredQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) SetVendoredQualityIssueScanningEnabled(v bool)`

SetVendoredQualityIssueScanningEnabled sets VendoredQualityIssueScanningEnabled field to given value.

### HasVendoredQualityIssueScanningEnabled

`func (o *UpdateProjectRequest) HasVendoredQualityIssueScanningEnabled() bool`

HasVendoredQualityIssueScanningEnabled returns a boolean if a field has been set.

### GetQuickImportVendoredDetectionEnabled

`func (o *UpdateProjectRequest) GetQuickImportVendoredDetectionEnabled() bool`

GetQuickImportVendoredDetectionEnabled returns the QuickImportVendoredDetectionEnabled field if non-nil, zero value otherwise.

### GetQuickImportVendoredDetectionEnabledOk

`func (o *UpdateProjectRequest) GetQuickImportVendoredDetectionEnabledOk() (*bool, bool)`

GetQuickImportVendoredDetectionEnabledOk returns a tuple with the QuickImportVendoredDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickImportVendoredDetectionEnabled

`func (o *UpdateProjectRequest) SetQuickImportVendoredDetectionEnabled(v bool)`

SetQuickImportVendoredDetectionEnabled sets QuickImportVendoredDetectionEnabled field to given value.

### HasQuickImportVendoredDetectionEnabled

`func (o *UpdateProjectRequest) HasQuickImportVendoredDetectionEnabled() bool`

HasQuickImportVendoredDetectionEnabled returns a boolean if a field has been set.

### GetQuickImportSnippetDetectionEnabled

`func (o *UpdateProjectRequest) GetQuickImportSnippetDetectionEnabled() bool`

GetQuickImportSnippetDetectionEnabled returns the QuickImportSnippetDetectionEnabled field if non-nil, zero value otherwise.

### GetQuickImportSnippetDetectionEnabledOk

`func (o *UpdateProjectRequest) GetQuickImportSnippetDetectionEnabledOk() (*bool, bool)`

GetQuickImportSnippetDetectionEnabledOk returns a tuple with the QuickImportSnippetDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickImportSnippetDetectionEnabled

`func (o *UpdateProjectRequest) SetQuickImportSnippetDetectionEnabled(v bool)`

SetQuickImportSnippetDetectionEnabled sets QuickImportSnippetDetectionEnabled field to given value.

### HasQuickImportSnippetDetectionEnabled

`func (o *UpdateProjectRequest) HasQuickImportSnippetDetectionEnabled() bool`

HasQuickImportSnippetDetectionEnabled returns a boolean if a field has been set.

### GetDefaultSavedOptionLicensingReport

`func (o *UpdateProjectRequest) GetDefaultSavedOptionLicensingReport() int32`

GetDefaultSavedOptionLicensingReport returns the DefaultSavedOptionLicensingReport field if non-nil, zero value otherwise.

### GetDefaultSavedOptionLicensingReportOk

`func (o *UpdateProjectRequest) GetDefaultSavedOptionLicensingReportOk() (*int32, bool)`

GetDefaultSavedOptionLicensingReportOk returns a tuple with the DefaultSavedOptionLicensingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSavedOptionLicensingReport

`func (o *UpdateProjectRequest) SetDefaultSavedOptionLicensingReport(v int32)`

SetDefaultSavedOptionLicensingReport sets DefaultSavedOptionLicensingReport field to given value.

### HasDefaultSavedOptionLicensingReport

`func (o *UpdateProjectRequest) HasDefaultSavedOptionLicensingReport() bool`

HasDefaultSavedOptionLicensingReport returns a boolean if a field has been set.

### GetDefaultSavedOptionSbom

`func (o *UpdateProjectRequest) GetDefaultSavedOptionSbom() int32`

GetDefaultSavedOptionSbom returns the DefaultSavedOptionSbom field if non-nil, zero value otherwise.

### GetDefaultSavedOptionSbomOk

`func (o *UpdateProjectRequest) GetDefaultSavedOptionSbomOk() (*int32, bool)`

GetDefaultSavedOptionSbomOk returns a tuple with the DefaultSavedOptionSbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSavedOptionSbom

`func (o *UpdateProjectRequest) SetDefaultSavedOptionSbom(v int32)`

SetDefaultSavedOptionSbom sets DefaultSavedOptionSbom field to given value.

### HasDefaultSavedOptionSbom

`func (o *UpdateProjectRequest) HasDefaultSavedOptionSbom() bool`

HasDefaultSavedOptionSbom returns a boolean if a field has been set.

### GetInvalidCredential

`func (o *UpdateProjectRequest) GetInvalidCredential() bool`

GetInvalidCredential returns the InvalidCredential field if non-nil, zero value otherwise.

### GetInvalidCredentialOk

`func (o *UpdateProjectRequest) GetInvalidCredentialOk() (*bool, bool)`

GetInvalidCredentialOk returns a tuple with the InvalidCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCredential

`func (o *UpdateProjectRequest) SetInvalidCredential(v bool)`

SetInvalidCredential sets InvalidCredential field to given value.

### HasInvalidCredential

`func (o *UpdateProjectRequest) HasInvalidCredential() bool`

HasInvalidCredential returns a boolean if a field has been set.

### GetLicensingStatusCheckEnabled

`func (o *UpdateProjectRequest) GetLicensingStatusCheckEnabled() bool`

GetLicensingStatusCheckEnabled returns the LicensingStatusCheckEnabled field if non-nil, zero value otherwise.

### GetLicensingStatusCheckEnabledOk

`func (o *UpdateProjectRequest) GetLicensingStatusCheckEnabledOk() (*bool, bool)`

GetLicensingStatusCheckEnabledOk returns a tuple with the LicensingStatusCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingStatusCheckEnabled

`func (o *UpdateProjectRequest) SetLicensingStatusCheckEnabled(v bool)`

SetLicensingStatusCheckEnabled sets LicensingStatusCheckEnabled field to given value.

### HasLicensingStatusCheckEnabled

`func (o *UpdateProjectRequest) HasLicensingStatusCheckEnabled() bool`

HasLicensingStatusCheckEnabled returns a boolean if a field has been set.

### GetSecurityStatusCheckEnabled

`func (o *UpdateProjectRequest) GetSecurityStatusCheckEnabled() bool`

GetSecurityStatusCheckEnabled returns the SecurityStatusCheckEnabled field if non-nil, zero value otherwise.

### GetSecurityStatusCheckEnabledOk

`func (o *UpdateProjectRequest) GetSecurityStatusCheckEnabledOk() (*bool, bool)`

GetSecurityStatusCheckEnabledOk returns a tuple with the SecurityStatusCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityStatusCheckEnabled

`func (o *UpdateProjectRequest) SetSecurityStatusCheckEnabled(v bool)`

SetSecurityStatusCheckEnabled sets SecurityStatusCheckEnabled field to given value.

### HasSecurityStatusCheckEnabled

`func (o *UpdateProjectRequest) HasSecurityStatusCheckEnabled() bool`

HasSecurityStatusCheckEnabled returns a boolean if a field has been set.

### GetQualityStatusCheckEnabled

`func (o *UpdateProjectRequest) GetQualityStatusCheckEnabled() bool`

GetQualityStatusCheckEnabled returns the QualityStatusCheckEnabled field if non-nil, zero value otherwise.

### GetQualityStatusCheckEnabledOk

`func (o *UpdateProjectRequest) GetQualityStatusCheckEnabledOk() (*bool, bool)`

GetQualityStatusCheckEnabledOk returns a tuple with the QualityStatusCheckEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityStatusCheckEnabled

`func (o *UpdateProjectRequest) SetQualityStatusCheckEnabled(v bool)`

SetQualityStatusCheckEnabled sets QualityStatusCheckEnabled field to given value.

### HasQualityStatusCheckEnabled

`func (o *UpdateProjectRequest) HasQualityStatusCheckEnabled() bool`

HasQualityStatusCheckEnabled returns a boolean if a field has been set.

### GetExcludeBaseLayerIssuesLicensing

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesLicensing() bool`

GetExcludeBaseLayerIssuesLicensing returns the ExcludeBaseLayerIssuesLicensing field if non-nil, zero value otherwise.

### GetExcludeBaseLayerIssuesLicensingOk

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesLicensingOk() (*bool, bool)`

GetExcludeBaseLayerIssuesLicensingOk returns a tuple with the ExcludeBaseLayerIssuesLicensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseLayerIssuesLicensing

`func (o *UpdateProjectRequest) SetExcludeBaseLayerIssuesLicensing(v bool)`

SetExcludeBaseLayerIssuesLicensing sets ExcludeBaseLayerIssuesLicensing field to given value.

### HasExcludeBaseLayerIssuesLicensing

`func (o *UpdateProjectRequest) HasExcludeBaseLayerIssuesLicensing() bool`

HasExcludeBaseLayerIssuesLicensing returns a boolean if a field has been set.

### GetExcludeBaseLayerIssuesSecurity

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesSecurity() bool`

GetExcludeBaseLayerIssuesSecurity returns the ExcludeBaseLayerIssuesSecurity field if non-nil, zero value otherwise.

### GetExcludeBaseLayerIssuesSecurityOk

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesSecurityOk() (*bool, bool)`

GetExcludeBaseLayerIssuesSecurityOk returns a tuple with the ExcludeBaseLayerIssuesSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseLayerIssuesSecurity

`func (o *UpdateProjectRequest) SetExcludeBaseLayerIssuesSecurity(v bool)`

SetExcludeBaseLayerIssuesSecurity sets ExcludeBaseLayerIssuesSecurity field to given value.

### HasExcludeBaseLayerIssuesSecurity

`func (o *UpdateProjectRequest) HasExcludeBaseLayerIssuesSecurity() bool`

HasExcludeBaseLayerIssuesSecurity returns a boolean if a field has been set.

### GetExcludeBaseLayerIssuesQuality

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesQuality() bool`

GetExcludeBaseLayerIssuesQuality returns the ExcludeBaseLayerIssuesQuality field if non-nil, zero value otherwise.

### GetExcludeBaseLayerIssuesQualityOk

`func (o *UpdateProjectRequest) GetExcludeBaseLayerIssuesQualityOk() (*bool, bool)`

GetExcludeBaseLayerIssuesQualityOk returns a tuple with the ExcludeBaseLayerIssuesQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeBaseLayerIssuesQuality

`func (o *UpdateProjectRequest) SetExcludeBaseLayerIssuesQuality(v bool)`

SetExcludeBaseLayerIssuesQuality sets ExcludeBaseLayerIssuesQuality field to given value.

### HasExcludeBaseLayerIssuesQuality

`func (o *UpdateProjectRequest) HasExcludeBaseLayerIssuesQuality() bool`

HasExcludeBaseLayerIssuesQuality returns a boolean if a field has been set.

### GetIntegrationhookTimeout

`func (o *UpdateProjectRequest) GetIntegrationhookTimeout() int32`

GetIntegrationhookTimeout returns the IntegrationhookTimeout field if non-nil, zero value otherwise.

### GetIntegrationhookTimeoutOk

`func (o *UpdateProjectRequest) GetIntegrationhookTimeoutOk() (*int32, bool)`

GetIntegrationhookTimeoutOk returns a tuple with the IntegrationhookTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationhookTimeout

`func (o *UpdateProjectRequest) SetIntegrationhookTimeout(v int32)`

SetIntegrationhookTimeout sets IntegrationhookTimeout field to given value.

### HasIntegrationhookTimeout

`func (o *UpdateProjectRequest) HasIntegrationhookTimeout() bool`

HasIntegrationhookTimeout returns a boolean if a field has been set.

### GetIntegrationhookFailState

`func (o *UpdateProjectRequest) GetIntegrationhookFailState() string`

GetIntegrationhookFailState returns the IntegrationhookFailState field if non-nil, zero value otherwise.

### GetIntegrationhookFailStateOk

`func (o *UpdateProjectRequest) GetIntegrationhookFailStateOk() (*string, bool)`

GetIntegrationhookFailStateOk returns a tuple with the IntegrationhookFailState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationhookFailState

`func (o *UpdateProjectRequest) SetIntegrationhookFailState(v string)`

SetIntegrationhookFailState sets IntegrationhookFailState field to given value.

### HasIntegrationhookFailState

`func (o *UpdateProjectRequest) HasIntegrationhookFailState() bool`

HasIntegrationhookFailState returns a boolean if a field has been set.

### GetIssueTrackerUrl

`func (o *UpdateProjectRequest) GetIssueTrackerUrl() string`

GetIssueTrackerUrl returns the IssueTrackerUrl field if non-nil, zero value otherwise.

### GetIssueTrackerUrlOk

`func (o *UpdateProjectRequest) GetIssueTrackerUrlOk() (*string, bool)`

GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerUrl

`func (o *UpdateProjectRequest) SetIssueTrackerUrl(v string)`

SetIssueTrackerUrl sets IssueTrackerUrl field to given value.

### HasIssueTrackerUrl

`func (o *UpdateProjectRequest) HasIssueTrackerUrl() bool`

HasIssueTrackerUrl returns a boolean if a field has been set.

### GetIssueTrackerType

`func (o *UpdateProjectRequest) GetIssueTrackerType() string`

GetIssueTrackerType returns the IssueTrackerType field if non-nil, zero value otherwise.

### GetIssueTrackerTypeOk

`func (o *UpdateProjectRequest) GetIssueTrackerTypeOk() (*string, bool)`

GetIssueTrackerTypeOk returns a tuple with the IssueTrackerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerType

`func (o *UpdateProjectRequest) SetIssueTrackerType(v string)`

SetIssueTrackerType sets IssueTrackerType field to given value.

### HasIssueTrackerType

`func (o *UpdateProjectRequest) HasIssueTrackerType() bool`

HasIssueTrackerType returns a boolean if a field has been set.

### GetIssueTrackerLabels

`func (o *UpdateProjectRequest) GetIssueTrackerLabels() []string`

GetIssueTrackerLabels returns the IssueTrackerLabels field if non-nil, zero value otherwise.

### GetIssueTrackerLabelsOk

`func (o *UpdateProjectRequest) GetIssueTrackerLabelsOk() (*[]string, bool)`

GetIssueTrackerLabelsOk returns a tuple with the IssueTrackerLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerLabels

`func (o *UpdateProjectRequest) SetIssueTrackerLabels(v []string)`

SetIssueTrackerLabels sets IssueTrackerLabels field to given value.

### HasIssueTrackerLabels

`func (o *UpdateProjectRequest) HasIssueTrackerLabels() bool`

HasIssueTrackerLabels returns a boolean if a field has been set.

### GetIssueTrackerIssueTypes

`func (o *UpdateProjectRequest) GetIssueTrackerIssueTypes() []string`

GetIssueTrackerIssueTypes returns the IssueTrackerIssueTypes field if non-nil, zero value otherwise.

### GetIssueTrackerIssueTypesOk

`func (o *UpdateProjectRequest) GetIssueTrackerIssueTypesOk() (*[]string, bool)`

GetIssueTrackerIssueTypesOk returns a tuple with the IssueTrackerIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerIssueTypes

`func (o *UpdateProjectRequest) SetIssueTrackerIssueTypes(v []string)`

SetIssueTrackerIssueTypes sets IssueTrackerIssueTypes field to given value.

### HasIssueTrackerIssueTypes

`func (o *UpdateProjectRequest) HasIssueTrackerIssueTypes() bool`

HasIssueTrackerIssueTypes returns a boolean if a field has been set.

### GetIssueTrackerProjectIds

`func (o *UpdateProjectRequest) GetIssueTrackerProjectIds() []string`

GetIssueTrackerProjectIds returns the IssueTrackerProjectIds field if non-nil, zero value otherwise.

### GetIssueTrackerProjectIdsOk

`func (o *UpdateProjectRequest) GetIssueTrackerProjectIdsOk() (*[]string, bool)`

GetIssueTrackerProjectIdsOk returns a tuple with the IssueTrackerProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerProjectIds

`func (o *UpdateProjectRequest) SetIssueTrackerProjectIds(v []string)`

SetIssueTrackerProjectIds sets IssueTrackerProjectIds field to given value.

### HasIssueTrackerProjectIds

`func (o *UpdateProjectRequest) HasIssueTrackerProjectIds() bool`

HasIssueTrackerProjectIds returns a boolean if a field has been set.

### GetIssueTrackerCustomFields

`func (o *UpdateProjectRequest) GetIssueTrackerCustomFields() map[string]UpdateProjectRequestIssueTrackerCustomFieldsValue`

GetIssueTrackerCustomFields returns the IssueTrackerCustomFields field if non-nil, zero value otherwise.

### GetIssueTrackerCustomFieldsOk

`func (o *UpdateProjectRequest) GetIssueTrackerCustomFieldsOk() (*map[string]UpdateProjectRequestIssueTrackerCustomFieldsValue, bool)`

GetIssueTrackerCustomFieldsOk returns a tuple with the IssueTrackerCustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerCustomFields

`func (o *UpdateProjectRequest) SetIssueTrackerCustomFields(v map[string]UpdateProjectRequestIssueTrackerCustomFieldsValue)`

SetIssueTrackerCustomFields sets IssueTrackerCustomFields field to given value.

### HasIssueTrackerCustomFields

`func (o *UpdateProjectRequest) HasIssueTrackerCustomFields() bool`

HasIssueTrackerCustomFields returns a boolean if a field has been set.

### GetUseGlobalTrackerSettings

`func (o *UpdateProjectRequest) GetUseGlobalTrackerSettings() bool`

GetUseGlobalTrackerSettings returns the UseGlobalTrackerSettings field if non-nil, zero value otherwise.

### GetUseGlobalTrackerSettingsOk

`func (o *UpdateProjectRequest) GetUseGlobalTrackerSettingsOk() (*bool, bool)`

GetUseGlobalTrackerSettingsOk returns a tuple with the UseGlobalTrackerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalTrackerSettings

`func (o *UpdateProjectRequest) SetUseGlobalTrackerSettings(v bool)`

SetUseGlobalTrackerSettings sets UseGlobalTrackerSettings field to given value.

### HasUseGlobalTrackerSettings

`func (o *UpdateProjectRequest) HasUseGlobalTrackerSettings() bool`

HasUseGlobalTrackerSettings returns a boolean if a field has been set.

### GetTransitiveExcludes

`func (o *UpdateProjectRequest) GetTransitiveExcludes() []string`

GetTransitiveExcludes returns the TransitiveExcludes field if non-nil, zero value otherwise.

### GetTransitiveExcludesOk

`func (o *UpdateProjectRequest) GetTransitiveExcludesOk() (*[]string, bool)`

GetTransitiveExcludesOk returns a tuple with the TransitiveExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitiveExcludes

`func (o *UpdateProjectRequest) SetTransitiveExcludes(v []string)`

SetTransitiveExcludes sets TransitiveExcludes field to given value.

### HasTransitiveExcludes

`func (o *UpdateProjectRequest) HasTransitiveExcludes() bool`

HasTransitiveExcludes returns a boolean if a field has been set.

### GetReportCustomText

`func (o *UpdateProjectRequest) GetReportCustomText() string`

GetReportCustomText returns the ReportCustomText field if non-nil, zero value otherwise.

### GetReportCustomTextOk

`func (o *UpdateProjectRequest) GetReportCustomTextOk() (*string, bool)`

GetReportCustomTextOk returns a tuple with the ReportCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomText

`func (o *UpdateProjectRequest) SetReportCustomText(v string)`

SetReportCustomText sets ReportCustomText field to given value.

### HasReportCustomText

`func (o *UpdateProjectRequest) HasReportCustomText() bool`

HasReportCustomText returns a boolean if a field has been set.

### GetBomColumnSettings

`func (o *UpdateProjectRequest) GetBomColumnSettings() []string`

GetBomColumnSettings returns the BomColumnSettings field if non-nil, zero value otherwise.

### GetBomColumnSettingsOk

`func (o *UpdateProjectRequest) GetBomColumnSettingsOk() (*[]string, bool)`

GetBomColumnSettingsOk returns a tuple with the BomColumnSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomColumnSettings

`func (o *UpdateProjectRequest) SetBomColumnSettings(v []string)`

SetBomColumnSettings sets BomColumnSettings field to given value.

### HasBomColumnSettings

`func (o *UpdateProjectRequest) HasBomColumnSettings() bool`

HasBomColumnSettings returns a boolean if a field has been set.

### GetBomPublicId

`func (o *UpdateProjectRequest) GetBomPublicId() string`

GetBomPublicId returns the BomPublicId field if non-nil, zero value otherwise.

### GetBomPublicIdOk

`func (o *UpdateProjectRequest) GetBomPublicIdOk() (*string, bool)`

GetBomPublicIdOk returns a tuple with the BomPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBomPublicId

`func (o *UpdateProjectRequest) SetBomPublicId(v string)`

SetBomPublicId sets BomPublicId field to given value.

### HasBomPublicId

`func (o *UpdateProjectRequest) HasBomPublicId() bool`

HasBomPublicId returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateProjectRequest) GetLabels() []int32`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateProjectRequest) GetLabelsOk() (*[]int32, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateProjectRequest) SetLabels(v []int32)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateProjectRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFilters

`func (o *UpdateProjectRequest) GetFilters() UpdateProjectRequestFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateProjectRequest) GetFiltersOk() (*UpdateProjectRequestFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateProjectRequest) SetFilters(v UpdateProjectRequestFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateProjectRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetNotifications

`func (o *UpdateProjectRequest) GetNotifications() []UpdateProjectRequestNotificationsInner`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *UpdateProjectRequest) GetNotificationsOk() (*[]UpdateProjectRequestNotificationsInner, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *UpdateProjectRequest) SetNotifications(v []UpdateProjectRequestNotificationsInner)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *UpdateProjectRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


