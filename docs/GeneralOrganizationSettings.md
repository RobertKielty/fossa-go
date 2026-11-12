# GeneralOrganizationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]GetOrganizationGeneralSettings200ResponseLabelsInner**](GetOrganizationGeneralSettings200ResponseLabelsInner.md) |  | [optional] 
**PackageLabels** | Pointer to [**[]GetOrganizationGeneralSettings200ResponsePackageLabelsInner**](GetOrganizationGeneralSettings200ResponsePackageLabelsInner.md) | The organization&#39;s package labels, each with the count of packages it is assigned to. | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DefaultRoleId** | Pointer to **int32** |  | [optional] 
**DependencySignatures** | Pointer to **string** |  | [optional] 
**DisableNonCustomTeamUserRoles** | Pointer to **bool** |  | [optional] 
**SnippetSourceCodeRetentionDays** | Pointer to **int32** | Number of days source code from snippet matches is retained. Must be between 1 and 30. | [optional] 
**LicenseConcludedEnabled** | Pointer to **bool** | Whether the license-concluded workflow is enabled for the organization. | [optional] 

## Methods

### NewGeneralOrganizationSettings

`func NewGeneralOrganizationSettings() *GeneralOrganizationSettings`

NewGeneralOrganizationSettings instantiates a new GeneralOrganizationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralOrganizationSettingsWithDefaults

`func NewGeneralOrganizationSettingsWithDefaults() *GeneralOrganizationSettings`

NewGeneralOrganizationSettingsWithDefaults instantiates a new GeneralOrganizationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *GeneralOrganizationSettings) GetLabels() []GetOrganizationGeneralSettings200ResponseLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GeneralOrganizationSettings) GetLabelsOk() (*[]GetOrganizationGeneralSettings200ResponseLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GeneralOrganizationSettings) SetLabels(v []GetOrganizationGeneralSettings200ResponseLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GeneralOrganizationSettings) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetPackageLabels

`func (o *GeneralOrganizationSettings) GetPackageLabels() []GetOrganizationGeneralSettings200ResponsePackageLabelsInner`

GetPackageLabels returns the PackageLabels field if non-nil, zero value otherwise.

### GetPackageLabelsOk

`func (o *GeneralOrganizationSettings) GetPackageLabelsOk() (*[]GetOrganizationGeneralSettings200ResponsePackageLabelsInner, bool)`

GetPackageLabelsOk returns a tuple with the PackageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabels

`func (o *GeneralOrganizationSettings) SetPackageLabels(v []GetOrganizationGeneralSettings200ResponsePackageLabelsInner)`

SetPackageLabels sets PackageLabels field to given value.

### HasPackageLabels

`func (o *GeneralOrganizationSettings) HasPackageLabels() bool`

HasPackageLabels returns a boolean if a field has been set.

### GetTitle

`func (o *GeneralOrganizationSettings) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GeneralOrganizationSettings) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GeneralOrganizationSettings) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GeneralOrganizationSettings) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetEmail

`func (o *GeneralOrganizationSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GeneralOrganizationSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GeneralOrganizationSettings) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GeneralOrganizationSettings) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *GeneralOrganizationSettings) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *GeneralOrganizationSettings) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *GeneralOrganizationSettings) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *GeneralOrganizationSettings) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetDependencySignatures

`func (o *GeneralOrganizationSettings) GetDependencySignatures() string`

GetDependencySignatures returns the DependencySignatures field if non-nil, zero value otherwise.

### GetDependencySignaturesOk

`func (o *GeneralOrganizationSettings) GetDependencySignaturesOk() (*string, bool)`

GetDependencySignaturesOk returns a tuple with the DependencySignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencySignatures

`func (o *GeneralOrganizationSettings) SetDependencySignatures(v string)`

SetDependencySignatures sets DependencySignatures field to given value.

### HasDependencySignatures

`func (o *GeneralOrganizationSettings) HasDependencySignatures() bool`

HasDependencySignatures returns a boolean if a field has been set.

### GetDisableNonCustomTeamUserRoles

`func (o *GeneralOrganizationSettings) GetDisableNonCustomTeamUserRoles() bool`

GetDisableNonCustomTeamUserRoles returns the DisableNonCustomTeamUserRoles field if non-nil, zero value otherwise.

### GetDisableNonCustomTeamUserRolesOk

`func (o *GeneralOrganizationSettings) GetDisableNonCustomTeamUserRolesOk() (*bool, bool)`

GetDisableNonCustomTeamUserRolesOk returns a tuple with the DisableNonCustomTeamUserRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNonCustomTeamUserRoles

`func (o *GeneralOrganizationSettings) SetDisableNonCustomTeamUserRoles(v bool)`

SetDisableNonCustomTeamUserRoles sets DisableNonCustomTeamUserRoles field to given value.

### HasDisableNonCustomTeamUserRoles

`func (o *GeneralOrganizationSettings) HasDisableNonCustomTeamUserRoles() bool`

HasDisableNonCustomTeamUserRoles returns a boolean if a field has been set.

### GetSnippetSourceCodeRetentionDays

`func (o *GeneralOrganizationSettings) GetSnippetSourceCodeRetentionDays() int32`

GetSnippetSourceCodeRetentionDays returns the SnippetSourceCodeRetentionDays field if non-nil, zero value otherwise.

### GetSnippetSourceCodeRetentionDaysOk

`func (o *GeneralOrganizationSettings) GetSnippetSourceCodeRetentionDaysOk() (*int32, bool)`

GetSnippetSourceCodeRetentionDaysOk returns a tuple with the SnippetSourceCodeRetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippetSourceCodeRetentionDays

`func (o *GeneralOrganizationSettings) SetSnippetSourceCodeRetentionDays(v int32)`

SetSnippetSourceCodeRetentionDays sets SnippetSourceCodeRetentionDays field to given value.

### HasSnippetSourceCodeRetentionDays

`func (o *GeneralOrganizationSettings) HasSnippetSourceCodeRetentionDays() bool`

HasSnippetSourceCodeRetentionDays returns a boolean if a field has been set.

### GetLicenseConcludedEnabled

`func (o *GeneralOrganizationSettings) GetLicenseConcludedEnabled() bool`

GetLicenseConcludedEnabled returns the LicenseConcludedEnabled field if non-nil, zero value otherwise.

### GetLicenseConcludedEnabledOk

`func (o *GeneralOrganizationSettings) GetLicenseConcludedEnabledOk() (*bool, bool)`

GetLicenseConcludedEnabledOk returns a tuple with the LicenseConcludedEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseConcludedEnabled

`func (o *GeneralOrganizationSettings) SetLicenseConcludedEnabled(v bool)`

SetLicenseConcludedEnabled sets LicenseConcludedEnabled field to given value.

### HasLicenseConcludedEnabled

`func (o *GeneralOrganizationSettings) HasLicenseConcludedEnabled() bool`

HasLicenseConcludedEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


