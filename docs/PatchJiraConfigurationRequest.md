# PatchJiraConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**OrganizationId** | Pointer to **int32** |  | [optional] [readonly] 
**WebhookURL** | Pointer to **string** | Used by the Jira to FOSSA Webhook | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | Timestamp when Jira was created | [optional] [readonly] 
**Enabled** | Pointer to **bool** | On/Off state of the Jira site enabled in FOSSA | [optional] 
**Name** | Pointer to **string** | Display name of your Jira site inside of FOSSA | [optional] 
**BaseUrl** | Pointer to **NullableString** | Url of your Jira Site | [optional] 
**ResolvedStatuses** | Pointer to **[]string** |  | [optional] 
**ResolvedStatusesEnabled** | Pointer to **bool** | When true, incoming webhooks will ignore/unignore issues linked to the specific tickets in question if the status matches one of the statuses listed in resolved_statuses. | [optional] 
**Credentials** | Pointer to [**PatchJiraConfigurationRequestCredentials**](PatchJiraConfigurationRequestCredentials.md) |  | [optional] 
**Headers** | Pointer to **map[string]string** | HTTP headers to pass along when authenticating to the Jira site | [optional] 
**IssueTypes** | Pointer to **[]string** | Available issue types to use when exporting tickets | [optional] 
**Labels** | Pointer to **[]string** | Available labels to include when exporting tickets. Corresponds to a label in Jira | [optional] 
**JiraProjectIds** | Pointer to **[]string** | Available Jira Projects to export to from FOSSA | [optional] 
**CustomFields** | Pointer to [**map[string]PatchJiraConfigurationRequestCustomFieldsValue**](PatchJiraConfigurationRequestCustomFieldsValue.md) | a dictionary of custom fields | [optional] 
**DefaultLicensingProject** | Pointer to **NullableString** | The Jira Project to default to when exporting licensing issues | [optional] 
**DefaultSecurityProject** | Pointer to **NullableString** | The Jira Project to default to when exporting security issues | [optional] 
**DefaultUniqueTickets** | Pointer to **bool** | toggle to determine if each individual issue is 1:1 with a ticket upon creation | [optional] 

## Methods

### NewPatchJiraConfigurationRequest

`func NewPatchJiraConfigurationRequest() *PatchJiraConfigurationRequest`

NewPatchJiraConfigurationRequest instantiates a new PatchJiraConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJiraConfigurationRequestWithDefaults

`func NewPatchJiraConfigurationRequestWithDefaults() *PatchJiraConfigurationRequest`

NewPatchJiraConfigurationRequestWithDefaults instantiates a new PatchJiraConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchJiraConfigurationRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchJiraConfigurationRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchJiraConfigurationRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PatchJiraConfigurationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PatchJiraConfigurationRequest) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PatchJiraConfigurationRequest) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PatchJiraConfigurationRequest) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PatchJiraConfigurationRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetWebhookURL

`func (o *PatchJiraConfigurationRequest) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *PatchJiraConfigurationRequest) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *PatchJiraConfigurationRequest) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *PatchJiraConfigurationRequest) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PatchJiraConfigurationRequest) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PatchJiraConfigurationRequest) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PatchJiraConfigurationRequest) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PatchJiraConfigurationRequest) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchJiraConfigurationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchJiraConfigurationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchJiraConfigurationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchJiraConfigurationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *PatchJiraConfigurationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchJiraConfigurationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchJiraConfigurationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchJiraConfigurationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *PatchJiraConfigurationRequest) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *PatchJiraConfigurationRequest) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *PatchJiraConfigurationRequest) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *PatchJiraConfigurationRequest) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *PatchJiraConfigurationRequest) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *PatchJiraConfigurationRequest) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetResolvedStatuses

`func (o *PatchJiraConfigurationRequest) GetResolvedStatuses() []string`

GetResolvedStatuses returns the ResolvedStatuses field if non-nil, zero value otherwise.

### GetResolvedStatusesOk

`func (o *PatchJiraConfigurationRequest) GetResolvedStatusesOk() (*[]string, bool)`

GetResolvedStatusesOk returns a tuple with the ResolvedStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedStatuses

`func (o *PatchJiraConfigurationRequest) SetResolvedStatuses(v []string)`

SetResolvedStatuses sets ResolvedStatuses field to given value.

### HasResolvedStatuses

`func (o *PatchJiraConfigurationRequest) HasResolvedStatuses() bool`

HasResolvedStatuses returns a boolean if a field has been set.

### GetResolvedStatusesEnabled

`func (o *PatchJiraConfigurationRequest) GetResolvedStatusesEnabled() bool`

GetResolvedStatusesEnabled returns the ResolvedStatusesEnabled field if non-nil, zero value otherwise.

### GetResolvedStatusesEnabledOk

`func (o *PatchJiraConfigurationRequest) GetResolvedStatusesEnabledOk() (*bool, bool)`

GetResolvedStatusesEnabledOk returns a tuple with the ResolvedStatusesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedStatusesEnabled

`func (o *PatchJiraConfigurationRequest) SetResolvedStatusesEnabled(v bool)`

SetResolvedStatusesEnabled sets ResolvedStatusesEnabled field to given value.

### HasResolvedStatusesEnabled

`func (o *PatchJiraConfigurationRequest) HasResolvedStatusesEnabled() bool`

HasResolvedStatusesEnabled returns a boolean if a field has been set.

### GetCredentials

`func (o *PatchJiraConfigurationRequest) GetCredentials() PatchJiraConfigurationRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PatchJiraConfigurationRequest) GetCredentialsOk() (*PatchJiraConfigurationRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PatchJiraConfigurationRequest) SetCredentials(v PatchJiraConfigurationRequestCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PatchJiraConfigurationRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHeaders

`func (o *PatchJiraConfigurationRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatchJiraConfigurationRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatchJiraConfigurationRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatchJiraConfigurationRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetIssueTypes

`func (o *PatchJiraConfigurationRequest) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *PatchJiraConfigurationRequest) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *PatchJiraConfigurationRequest) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *PatchJiraConfigurationRequest) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### SetIssueTypesNil

`func (o *PatchJiraConfigurationRequest) SetIssueTypesNil(b bool)`

 SetIssueTypesNil sets the value for IssueTypes to be an explicit nil

### UnsetIssueTypes
`func (o *PatchJiraConfigurationRequest) UnsetIssueTypes()`

UnsetIssueTypes ensures that no value is present for IssueTypes, not even an explicit nil
### GetLabels

`func (o *PatchJiraConfigurationRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PatchJiraConfigurationRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PatchJiraConfigurationRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PatchJiraConfigurationRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *PatchJiraConfigurationRequest) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *PatchJiraConfigurationRequest) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetJiraProjectIds

`func (o *PatchJiraConfigurationRequest) GetJiraProjectIds() []string`

GetJiraProjectIds returns the JiraProjectIds field if non-nil, zero value otherwise.

### GetJiraProjectIdsOk

`func (o *PatchJiraConfigurationRequest) GetJiraProjectIdsOk() (*[]string, bool)`

GetJiraProjectIdsOk returns a tuple with the JiraProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraProjectIds

`func (o *PatchJiraConfigurationRequest) SetJiraProjectIds(v []string)`

SetJiraProjectIds sets JiraProjectIds field to given value.

### HasJiraProjectIds

`func (o *PatchJiraConfigurationRequest) HasJiraProjectIds() bool`

HasJiraProjectIds returns a boolean if a field has been set.

### SetJiraProjectIdsNil

`func (o *PatchJiraConfigurationRequest) SetJiraProjectIdsNil(b bool)`

 SetJiraProjectIdsNil sets the value for JiraProjectIds to be an explicit nil

### UnsetJiraProjectIds
`func (o *PatchJiraConfigurationRequest) UnsetJiraProjectIds()`

UnsetJiraProjectIds ensures that no value is present for JiraProjectIds, not even an explicit nil
### GetCustomFields

`func (o *PatchJiraConfigurationRequest) GetCustomFields() map[string]PatchJiraConfigurationRequestCustomFieldsValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchJiraConfigurationRequest) GetCustomFieldsOk() (*map[string]PatchJiraConfigurationRequestCustomFieldsValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchJiraConfigurationRequest) SetCustomFields(v map[string]PatchJiraConfigurationRequestCustomFieldsValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchJiraConfigurationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDefaultLicensingProject

`func (o *PatchJiraConfigurationRequest) GetDefaultLicensingProject() string`

GetDefaultLicensingProject returns the DefaultLicensingProject field if non-nil, zero value otherwise.

### GetDefaultLicensingProjectOk

`func (o *PatchJiraConfigurationRequest) GetDefaultLicensingProjectOk() (*string, bool)`

GetDefaultLicensingProjectOk returns a tuple with the DefaultLicensingProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicensingProject

`func (o *PatchJiraConfigurationRequest) SetDefaultLicensingProject(v string)`

SetDefaultLicensingProject sets DefaultLicensingProject field to given value.

### HasDefaultLicensingProject

`func (o *PatchJiraConfigurationRequest) HasDefaultLicensingProject() bool`

HasDefaultLicensingProject returns a boolean if a field has been set.

### SetDefaultLicensingProjectNil

`func (o *PatchJiraConfigurationRequest) SetDefaultLicensingProjectNil(b bool)`

 SetDefaultLicensingProjectNil sets the value for DefaultLicensingProject to be an explicit nil

### UnsetDefaultLicensingProject
`func (o *PatchJiraConfigurationRequest) UnsetDefaultLicensingProject()`

UnsetDefaultLicensingProject ensures that no value is present for DefaultLicensingProject, not even an explicit nil
### GetDefaultSecurityProject

`func (o *PatchJiraConfigurationRequest) GetDefaultSecurityProject() string`

GetDefaultSecurityProject returns the DefaultSecurityProject field if non-nil, zero value otherwise.

### GetDefaultSecurityProjectOk

`func (o *PatchJiraConfigurationRequest) GetDefaultSecurityProjectOk() (*string, bool)`

GetDefaultSecurityProjectOk returns a tuple with the DefaultSecurityProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityProject

`func (o *PatchJiraConfigurationRequest) SetDefaultSecurityProject(v string)`

SetDefaultSecurityProject sets DefaultSecurityProject field to given value.

### HasDefaultSecurityProject

`func (o *PatchJiraConfigurationRequest) HasDefaultSecurityProject() bool`

HasDefaultSecurityProject returns a boolean if a field has been set.

### SetDefaultSecurityProjectNil

`func (o *PatchJiraConfigurationRequest) SetDefaultSecurityProjectNil(b bool)`

 SetDefaultSecurityProjectNil sets the value for DefaultSecurityProject to be an explicit nil

### UnsetDefaultSecurityProject
`func (o *PatchJiraConfigurationRequest) UnsetDefaultSecurityProject()`

UnsetDefaultSecurityProject ensures that no value is present for DefaultSecurityProject, not even an explicit nil
### GetDefaultUniqueTickets

`func (o *PatchJiraConfigurationRequest) GetDefaultUniqueTickets() bool`

GetDefaultUniqueTickets returns the DefaultUniqueTickets field if non-nil, zero value otherwise.

### GetDefaultUniqueTicketsOk

`func (o *PatchJiraConfigurationRequest) GetDefaultUniqueTicketsOk() (*bool, bool)`

GetDefaultUniqueTicketsOk returns a tuple with the DefaultUniqueTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUniqueTickets

`func (o *PatchJiraConfigurationRequest) SetDefaultUniqueTickets(v bool)`

SetDefaultUniqueTickets sets DefaultUniqueTickets field to given value.

### HasDefaultUniqueTickets

`func (o *PatchJiraConfigurationRequest) HasDefaultUniqueTickets() bool`

HasDefaultUniqueTickets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


