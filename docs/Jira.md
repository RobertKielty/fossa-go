# Jira

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

### NewJira

`func NewJira() *Jira`

NewJira instantiates a new Jira object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraWithDefaults

`func NewJiraWithDefaults() *Jira`

NewJiraWithDefaults instantiates a new Jira object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Jira) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Jira) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Jira) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Jira) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Jira) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Jira) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Jira) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Jira) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetWebhookURL

`func (o *Jira) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *Jira) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *Jira) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *Jira) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Jira) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Jira) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Jira) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Jira) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *Jira) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Jira) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Jira) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Jira) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *Jira) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Jira) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Jira) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Jira) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBaseUrl

`func (o *Jira) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *Jira) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *Jira) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *Jira) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *Jira) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *Jira) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetResolvedStatuses

`func (o *Jira) GetResolvedStatuses() []string`

GetResolvedStatuses returns the ResolvedStatuses field if non-nil, zero value otherwise.

### GetResolvedStatusesOk

`func (o *Jira) GetResolvedStatusesOk() (*[]string, bool)`

GetResolvedStatusesOk returns a tuple with the ResolvedStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedStatuses

`func (o *Jira) SetResolvedStatuses(v []string)`

SetResolvedStatuses sets ResolvedStatuses field to given value.

### HasResolvedStatuses

`func (o *Jira) HasResolvedStatuses() bool`

HasResolvedStatuses returns a boolean if a field has been set.

### GetResolvedStatusesEnabled

`func (o *Jira) GetResolvedStatusesEnabled() bool`

GetResolvedStatusesEnabled returns the ResolvedStatusesEnabled field if non-nil, zero value otherwise.

### GetResolvedStatusesEnabledOk

`func (o *Jira) GetResolvedStatusesEnabledOk() (*bool, bool)`

GetResolvedStatusesEnabledOk returns a tuple with the ResolvedStatusesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedStatusesEnabled

`func (o *Jira) SetResolvedStatusesEnabled(v bool)`

SetResolvedStatusesEnabled sets ResolvedStatusesEnabled field to given value.

### HasResolvedStatusesEnabled

`func (o *Jira) HasResolvedStatusesEnabled() bool`

HasResolvedStatusesEnabled returns a boolean if a field has been set.

### GetCredentials

`func (o *Jira) GetCredentials() PatchJiraConfigurationRequestCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *Jira) GetCredentialsOk() (*PatchJiraConfigurationRequestCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *Jira) SetCredentials(v PatchJiraConfigurationRequestCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *Jira) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetHeaders

`func (o *Jira) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Jira) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Jira) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Jira) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetIssueTypes

`func (o *Jira) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *Jira) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *Jira) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *Jira) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### SetIssueTypesNil

`func (o *Jira) SetIssueTypesNil(b bool)`

 SetIssueTypesNil sets the value for IssueTypes to be an explicit nil

### UnsetIssueTypes
`func (o *Jira) UnsetIssueTypes()`

UnsetIssueTypes ensures that no value is present for IssueTypes, not even an explicit nil
### GetLabels

`func (o *Jira) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Jira) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Jira) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Jira) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Jira) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Jira) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetJiraProjectIds

`func (o *Jira) GetJiraProjectIds() []string`

GetJiraProjectIds returns the JiraProjectIds field if non-nil, zero value otherwise.

### GetJiraProjectIdsOk

`func (o *Jira) GetJiraProjectIdsOk() (*[]string, bool)`

GetJiraProjectIdsOk returns a tuple with the JiraProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraProjectIds

`func (o *Jira) SetJiraProjectIds(v []string)`

SetJiraProjectIds sets JiraProjectIds field to given value.

### HasJiraProjectIds

`func (o *Jira) HasJiraProjectIds() bool`

HasJiraProjectIds returns a boolean if a field has been set.

### SetJiraProjectIdsNil

`func (o *Jira) SetJiraProjectIdsNil(b bool)`

 SetJiraProjectIdsNil sets the value for JiraProjectIds to be an explicit nil

### UnsetJiraProjectIds
`func (o *Jira) UnsetJiraProjectIds()`

UnsetJiraProjectIds ensures that no value is present for JiraProjectIds, not even an explicit nil
### GetCustomFields

`func (o *Jira) GetCustomFields() map[string]PatchJiraConfigurationRequestCustomFieldsValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Jira) GetCustomFieldsOk() (*map[string]PatchJiraConfigurationRequestCustomFieldsValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Jira) SetCustomFields(v map[string]PatchJiraConfigurationRequestCustomFieldsValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Jira) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetDefaultLicensingProject

`func (o *Jira) GetDefaultLicensingProject() string`

GetDefaultLicensingProject returns the DefaultLicensingProject field if non-nil, zero value otherwise.

### GetDefaultLicensingProjectOk

`func (o *Jira) GetDefaultLicensingProjectOk() (*string, bool)`

GetDefaultLicensingProjectOk returns a tuple with the DefaultLicensingProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLicensingProject

`func (o *Jira) SetDefaultLicensingProject(v string)`

SetDefaultLicensingProject sets DefaultLicensingProject field to given value.

### HasDefaultLicensingProject

`func (o *Jira) HasDefaultLicensingProject() bool`

HasDefaultLicensingProject returns a boolean if a field has been set.

### SetDefaultLicensingProjectNil

`func (o *Jira) SetDefaultLicensingProjectNil(b bool)`

 SetDefaultLicensingProjectNil sets the value for DefaultLicensingProject to be an explicit nil

### UnsetDefaultLicensingProject
`func (o *Jira) UnsetDefaultLicensingProject()`

UnsetDefaultLicensingProject ensures that no value is present for DefaultLicensingProject, not even an explicit nil
### GetDefaultSecurityProject

`func (o *Jira) GetDefaultSecurityProject() string`

GetDefaultSecurityProject returns the DefaultSecurityProject field if non-nil, zero value otherwise.

### GetDefaultSecurityProjectOk

`func (o *Jira) GetDefaultSecurityProjectOk() (*string, bool)`

GetDefaultSecurityProjectOk returns a tuple with the DefaultSecurityProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityProject

`func (o *Jira) SetDefaultSecurityProject(v string)`

SetDefaultSecurityProject sets DefaultSecurityProject field to given value.

### HasDefaultSecurityProject

`func (o *Jira) HasDefaultSecurityProject() bool`

HasDefaultSecurityProject returns a boolean if a field has been set.

### SetDefaultSecurityProjectNil

`func (o *Jira) SetDefaultSecurityProjectNil(b bool)`

 SetDefaultSecurityProjectNil sets the value for DefaultSecurityProject to be an explicit nil

### UnsetDefaultSecurityProject
`func (o *Jira) UnsetDefaultSecurityProject()`

UnsetDefaultSecurityProject ensures that no value is present for DefaultSecurityProject, not even an explicit nil
### GetDefaultUniqueTickets

`func (o *Jira) GetDefaultUniqueTickets() bool`

GetDefaultUniqueTickets returns the DefaultUniqueTickets field if non-nil, zero value otherwise.

### GetDefaultUniqueTicketsOk

`func (o *Jira) GetDefaultUniqueTicketsOk() (*bool, bool)`

GetDefaultUniqueTicketsOk returns a tuple with the DefaultUniqueTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUniqueTickets

`func (o *Jira) SetDefaultUniqueTickets(v bool)`

SetDefaultUniqueTickets sets DefaultUniqueTickets field to given value.

### HasDefaultUniqueTickets

`func (o *Jira) HasDefaultUniqueTickets() bool`

HasDefaultUniqueTickets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


