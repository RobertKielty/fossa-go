# EditingGeneralOrganizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** |  | [optional] 
**PackageLabels** | Pointer to **[]string** | Package label names to set on the organization. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DefaultRoleId** | Pointer to **int32** |  | [optional] 
**DependencySignatures** | Pointer to **string** |  | [optional] 
**DisableNonCustomTeamUserRoles** | Pointer to **bool** |  | [optional] 
**SnippetSourceCodeRetentionDays** | Pointer to **int32** | Number of days to retain source code from snippet matches. Must be between 1 and 30, or a 400 is returned. | [optional] 
**LicenseConcludedEnabled** | Pointer to **bool** | Whether the license-concluded workflow is enabled for the organization. | [optional] 

## Methods

### NewEditingGeneralOrganizationSettings

`func NewEditingGeneralOrganizationSettings() *EditingGeneralOrganizationSettings`

NewEditingGeneralOrganizationSettings instantiates a new EditingGeneralOrganizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditingGeneralOrganizationSettingsWithDefaults

`func NewEditingGeneralOrganizationSettingsWithDefaults() *EditingGeneralOrganizationSettings`

NewEditingGeneralOrganizationSettingsWithDefaults instantiates a new EditingGeneralOrganizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *EditingGeneralOrganizationSettings) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EditingGeneralOrganizationSettings) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EditingGeneralOrganizationSettings) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EditingGeneralOrganizationSettings) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPackageLabels

`func (o *EditingGeneralOrganizationSettings) GetPackageLabels() []string`

GetPackageLabels returns the PackageLabels field if non-nil, zero value otherwise.

### GetPackageLabelsOk

`func (o *EditingGeneralOrganizationSettings) GetPackageLabelsOk() (*[]string, bool)`

GetPackageLabelsOk returns a tuple with the PackageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabels

`func (o *EditingGeneralOrganizationSettings) SetPackageLabels(v []string)`

SetPackageLabels sets PackageLabels field to given value.

### HasPackageLabels

`func (o *EditingGeneralOrganizationSettings) HasPackageLabels() bool`

HasPackageLabels returns a boolean if a field has been set.

### GetTitle

`func (o *EditingGeneralOrganizationSettings) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EditingGeneralOrganizationSettings) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EditingGeneralOrganizationSettings) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EditingGeneralOrganizationSettings) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEmail

`func (o *EditingGeneralOrganizationSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *EditingGeneralOrganizationSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *EditingGeneralOrganizationSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *EditingGeneralOrganizationSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *EditingGeneralOrganizationSettings) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *EditingGeneralOrganizationSettings) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *EditingGeneralOrganizationSettings) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *EditingGeneralOrganizationSettings) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDependencySignatures

`func (o *EditingGeneralOrganizationSettings) GetDependencySignatures() string`

GetDependencySignatures returns the DependencySignatures field if non-nil, zero value otherwise.

### GetDependencySignaturesOk

`func (o *EditingGeneralOrganizationSettings) GetDependencySignaturesOk() (*string, bool)`

GetDependencySignaturesOk returns a tuple with the DependencySignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencySignatures

`func (o *EditingGeneralOrganizationSettings) SetDependencySignatures(v string)`

SetDependencySignatures sets DependencySignatures field to given value.

### HasDependencySignatures

`func (o *EditingGeneralOrganizationSettings) HasDependencySignatures() bool`

HasDependencySignatures returns a boolean if a field has been set.

### GetDisableNonCustomTeamUserRoles

`func (o *EditingGeneralOrganizationSettings) GetDisableNonCustomTeamUserRoles() bool`

GetDisableNonCustomTeamUserRoles returns the DisableNonCustomTeamUserRoles field if non-nil, zero value otherwise.

### GetDisableNonCustomTeamUserRolesOk

`func (o *EditingGeneralOrganizationSettings) GetDisableNonCustomTeamUserRolesOk() (*bool, bool)`

GetDisableNonCustomTeamUserRolesOk returns a tuple with the DisableNonCustomTeamUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNonCustomTeamUserRoles

`func (o *EditingGeneralOrganizationSettings) SetDisableNonCustomTeamUserRoles(v bool)`

SetDisableNonCustomTeamUserRoles sets DisableNonCustomTeamUserRoles field to given value.

### HasDisableNonCustomTeamUserRoles

`func (o *EditingGeneralOrganizationSettings) HasDisableNonCustomTeamUserRoles() bool`

HasDisableNonCustomTeamUserRoles returns a boolean if a field has been set.

### GetSnippetSourceCodeRetentionDays

`func (o *EditingGeneralOrganizationSettings) GetSnippetSourceCodeRetentionDays() int32`

GetSnippetSourceCodeRetentionDays returns the SnippetSourceCodeRetentionDays field if non-nil, zero value otherwise.

### GetSnippetSourceCodeRetentionDaysOk

`func (o *EditingGeneralOrganizationSettings) GetSnippetSourceCodeRetentionDaysOk() (*int32, bool)`

GetSnippetSourceCodeRetentionDaysOk returns a tuple with the SnippetSourceCodeRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetSourceCodeRetentionDays

`func (o *EditingGeneralOrganizationSettings) SetSnippetSourceCodeRetentionDays(v int32)`

SetSnippetSourceCodeRetentionDays sets SnippetSourceCodeRetentionDays field to given value.

### HasSnippetSourceCodeRetentionDays

`func (o *EditingGeneralOrganizationSettings) HasSnippetSourceCodeRetentionDays() bool`

HasSnippetSourceCodeRetentionDays returns a boolean if a field has been set.

### GetLicenseConcludedEnabled

`func (o *EditingGeneralOrganizationSettings) GetLicenseConcludedEnabled() bool`

GetLicenseConcludedEnabled returns the LicenseConcludedEnabled field if non-nil, zero value otherwise.

### GetLicenseConcludedEnabledOk

`func (o *EditingGeneralOrganizationSettings) GetLicenseConcludedEnabledOk() (*bool, bool)`

GetLicenseConcludedEnabledOk returns a tuple with the LicenseConcludedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConcludedEnabled

`func (o *EditingGeneralOrganizationSettings) SetLicenseConcludedEnabled(v bool)`

SetLicenseConcludedEnabled sets LicenseConcludedEnabled field to given value.

### HasLicenseConcludedEnabled

`func (o *EditingGeneralOrganizationSettings) HasLicenseConcludedEnabled() bool`

HasLicenseConcludedEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


