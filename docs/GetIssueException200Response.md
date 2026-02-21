# GetIssueException200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IgnoreScope** | Pointer to **string** |  | [optional] 
**PackageScope** | Pointer to **string** |  | [optional] 
**IssueCategory** | Pointer to **string** |  | [optional] 
**ExceptionUserProjectId** | Pointer to **string** |  | [optional] 
**ExceptionUserReleaseGroupId** | Pointer to **int32** |  | [optional] 
**ExceptionUserPolicyId** | Pointer to **int32** |  | [optional] 
**ProjectTitle** | Pointer to **string** |  | [optional] 
**DependencyProjectLocator** | Pointer to **string** |  | [optional] 
**DependencyRevisionLocator** | Pointer to **string** |  | [optional] 
**DependencyTitle** | Pointer to **string** |  | [optional] 
**ExceptionTitle** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**PolicyTitle** | Pointer to **string** |  | [optional] 
**PackageLabel** | Pointer to **string** |  | [optional] 
**ExpiresAfter** | Pointer to **string** |  | [optional] 
**IsExpired** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIssueException200Response

`func NewGetIssueException200Response() *GetIssueException200Response`

NewGetIssueException200Response instantiates a new GetIssueException200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueException200ResponseWithDefaults

`func NewGetIssueException200ResponseWithDefaults() *GetIssueException200Response`

NewGetIssueException200ResponseWithDefaults instantiates a new GetIssueException200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssueException200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssueException200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssueException200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssueException200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIgnoreScope

`func (o *GetIssueException200Response) GetIgnoreScope() string`

GetIgnoreScope returns the IgnoreScope field if non-nil, zero value otherwise.

### GetIgnoreScopeOk

`func (o *GetIssueException200Response) GetIgnoreScopeOk() (*string, bool)`

GetIgnoreScopeOk returns a tuple with the IgnoreScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreScope

`func (o *GetIssueException200Response) SetIgnoreScope(v string)`

SetIgnoreScope sets IgnoreScope field to given value.

### HasIgnoreScope

`func (o *GetIssueException200Response) HasIgnoreScope() bool`

HasIgnoreScope returns a boolean if a field has been set.

### GetPackageScope

`func (o *GetIssueException200Response) GetPackageScope() string`

GetPackageScope returns the PackageScope field if non-nil, zero value otherwise.

### GetPackageScopeOk

`func (o *GetIssueException200Response) GetPackageScopeOk() (*string, bool)`

GetPackageScopeOk returns a tuple with the PackageScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageScope

`func (o *GetIssueException200Response) SetPackageScope(v string)`

SetPackageScope sets PackageScope field to given value.

### HasPackageScope

`func (o *GetIssueException200Response) HasPackageScope() bool`

HasPackageScope returns a boolean if a field has been set.

### GetIssueCategory

`func (o *GetIssueException200Response) GetIssueCategory() string`

GetIssueCategory returns the IssueCategory field if non-nil, zero value otherwise.

### GetIssueCategoryOk

`func (o *GetIssueException200Response) GetIssueCategoryOk() (*string, bool)`

GetIssueCategoryOk returns a tuple with the IssueCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCategory

`func (o *GetIssueException200Response) SetIssueCategory(v string)`

SetIssueCategory sets IssueCategory field to given value.

### HasIssueCategory

`func (o *GetIssueException200Response) HasIssueCategory() bool`

HasIssueCategory returns a boolean if a field has been set.

### GetExceptionUserProjectId

`func (o *GetIssueException200Response) GetExceptionUserProjectId() string`

GetExceptionUserProjectId returns the ExceptionUserProjectId field if non-nil, zero value otherwise.

### GetExceptionUserProjectIdOk

`func (o *GetIssueException200Response) GetExceptionUserProjectIdOk() (*string, bool)`

GetExceptionUserProjectIdOk returns a tuple with the ExceptionUserProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionUserProjectId

`func (o *GetIssueException200Response) SetExceptionUserProjectId(v string)`

SetExceptionUserProjectId sets ExceptionUserProjectId field to given value.

### HasExceptionUserProjectId

`func (o *GetIssueException200Response) HasExceptionUserProjectId() bool`

HasExceptionUserProjectId returns a boolean if a field has been set.

### GetExceptionUserReleaseGroupId

`func (o *GetIssueException200Response) GetExceptionUserReleaseGroupId() int32`

GetExceptionUserReleaseGroupId returns the ExceptionUserReleaseGroupId field if non-nil, zero value otherwise.

### GetExceptionUserReleaseGroupIdOk

`func (o *GetIssueException200Response) GetExceptionUserReleaseGroupIdOk() (*int32, bool)`

GetExceptionUserReleaseGroupIdOk returns a tuple with the ExceptionUserReleaseGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionUserReleaseGroupId

`func (o *GetIssueException200Response) SetExceptionUserReleaseGroupId(v int32)`

SetExceptionUserReleaseGroupId sets ExceptionUserReleaseGroupId field to given value.

### HasExceptionUserReleaseGroupId

`func (o *GetIssueException200Response) HasExceptionUserReleaseGroupId() bool`

HasExceptionUserReleaseGroupId returns a boolean if a field has been set.

### GetExceptionUserPolicyId

`func (o *GetIssueException200Response) GetExceptionUserPolicyId() int32`

GetExceptionUserPolicyId returns the ExceptionUserPolicyId field if non-nil, zero value otherwise.

### GetExceptionUserPolicyIdOk

`func (o *GetIssueException200Response) GetExceptionUserPolicyIdOk() (*int32, bool)`

GetExceptionUserPolicyIdOk returns a tuple with the ExceptionUserPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionUserPolicyId

`func (o *GetIssueException200Response) SetExceptionUserPolicyId(v int32)`

SetExceptionUserPolicyId sets ExceptionUserPolicyId field to given value.

### HasExceptionUserPolicyId

`func (o *GetIssueException200Response) HasExceptionUserPolicyId() bool`

HasExceptionUserPolicyId returns a boolean if a field has been set.

### GetProjectTitle

`func (o *GetIssueException200Response) GetProjectTitle() string`

GetProjectTitle returns the ProjectTitle field if non-nil, zero value otherwise.

### GetProjectTitleOk

`func (o *GetIssueException200Response) GetProjectTitleOk() (*string, bool)`

GetProjectTitleOk returns a tuple with the ProjectTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTitle

`func (o *GetIssueException200Response) SetProjectTitle(v string)`

SetProjectTitle sets ProjectTitle field to given value.

### HasProjectTitle

`func (o *GetIssueException200Response) HasProjectTitle() bool`

HasProjectTitle returns a boolean if a field has been set.

### GetDependencyProjectLocator

`func (o *GetIssueException200Response) GetDependencyProjectLocator() string`

GetDependencyProjectLocator returns the DependencyProjectLocator field if non-nil, zero value otherwise.

### GetDependencyProjectLocatorOk

`func (o *GetIssueException200Response) GetDependencyProjectLocatorOk() (*string, bool)`

GetDependencyProjectLocatorOk returns a tuple with the DependencyProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyProjectLocator

`func (o *GetIssueException200Response) SetDependencyProjectLocator(v string)`

SetDependencyProjectLocator sets DependencyProjectLocator field to given value.

### HasDependencyProjectLocator

`func (o *GetIssueException200Response) HasDependencyProjectLocator() bool`

HasDependencyProjectLocator returns a boolean if a field has been set.

### GetDependencyRevisionLocator

`func (o *GetIssueException200Response) GetDependencyRevisionLocator() string`

GetDependencyRevisionLocator returns the DependencyRevisionLocator field if non-nil, zero value otherwise.

### GetDependencyRevisionLocatorOk

`func (o *GetIssueException200Response) GetDependencyRevisionLocatorOk() (*string, bool)`

GetDependencyRevisionLocatorOk returns a tuple with the DependencyRevisionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyRevisionLocator

`func (o *GetIssueException200Response) SetDependencyRevisionLocator(v string)`

SetDependencyRevisionLocator sets DependencyRevisionLocator field to given value.

### HasDependencyRevisionLocator

`func (o *GetIssueException200Response) HasDependencyRevisionLocator() bool`

HasDependencyRevisionLocator returns a boolean if a field has been set.

### GetDependencyTitle

`func (o *GetIssueException200Response) GetDependencyTitle() string`

GetDependencyTitle returns the DependencyTitle field if non-nil, zero value otherwise.

### GetDependencyTitleOk

`func (o *GetIssueException200Response) GetDependencyTitleOk() (*string, bool)`

GetDependencyTitleOk returns a tuple with the DependencyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyTitle

`func (o *GetIssueException200Response) SetDependencyTitle(v string)`

SetDependencyTitle sets DependencyTitle field to given value.

### HasDependencyTitle

`func (o *GetIssueException200Response) HasDependencyTitle() bool`

HasDependencyTitle returns a boolean if a field has been set.

### GetExceptionTitle

`func (o *GetIssueException200Response) GetExceptionTitle() string`

GetExceptionTitle returns the ExceptionTitle field if non-nil, zero value otherwise.

### GetExceptionTitleOk

`func (o *GetIssueException200Response) GetExceptionTitleOk() (*string, bool)`

GetExceptionTitleOk returns a tuple with the ExceptionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionTitle

`func (o *GetIssueException200Response) SetExceptionTitle(v string)`

SetExceptionTitle sets ExceptionTitle field to given value.

### HasExceptionTitle

`func (o *GetIssueException200Response) HasExceptionTitle() bool`

HasExceptionTitle returns a boolean if a field has been set.

### GetNote

`func (o *GetIssueException200Response) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetIssueException200Response) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetIssueException200Response) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GetIssueException200Response) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetIssueException200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetIssueException200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetIssueException200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetIssueException200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetReason

`func (o *GetIssueException200Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetIssueException200Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetIssueException200Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetIssueException200Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPolicyTitle

`func (o *GetIssueException200Response) GetPolicyTitle() string`

GetPolicyTitle returns the PolicyTitle field if non-nil, zero value otherwise.

### GetPolicyTitleOk

`func (o *GetIssueException200Response) GetPolicyTitleOk() (*string, bool)`

GetPolicyTitleOk returns a tuple with the PolicyTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyTitle

`func (o *GetIssueException200Response) SetPolicyTitle(v string)`

SetPolicyTitle sets PolicyTitle field to given value.

### HasPolicyTitle

`func (o *GetIssueException200Response) HasPolicyTitle() bool`

HasPolicyTitle returns a boolean if a field has been set.

### GetPackageLabel

`func (o *GetIssueException200Response) GetPackageLabel() string`

GetPackageLabel returns the PackageLabel field if non-nil, zero value otherwise.

### GetPackageLabelOk

`func (o *GetIssueException200Response) GetPackageLabelOk() (*string, bool)`

GetPackageLabelOk returns a tuple with the PackageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabel

`func (o *GetIssueException200Response) SetPackageLabel(v string)`

SetPackageLabel sets PackageLabel field to given value.

### HasPackageLabel

`func (o *GetIssueException200Response) HasPackageLabel() bool`

HasPackageLabel returns a boolean if a field has been set.

### GetExpiresAfter

`func (o *GetIssueException200Response) GetExpiresAfter() string`

GetExpiresAfter returns the ExpiresAfter field if non-nil, zero value otherwise.

### GetExpiresAfterOk

`func (o *GetIssueException200Response) GetExpiresAfterOk() (*string, bool)`

GetExpiresAfterOk returns a tuple with the ExpiresAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfter

`func (o *GetIssueException200Response) SetExpiresAfter(v string)`

SetExpiresAfter sets ExpiresAfter field to given value.

### HasExpiresAfter

`func (o *GetIssueException200Response) HasExpiresAfter() bool`

HasExpiresAfter returns a boolean if a field has been set.

### GetIsExpired

`func (o *GetIssueException200Response) GetIsExpired() bool`

GetIsExpired returns the IsExpired field if non-nil, zero value otherwise.

### GetIsExpiredOk

`func (o *GetIssueException200Response) GetIsExpiredOk() (*bool, bool)`

GetIsExpiredOk returns a tuple with the IsExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExpired

`func (o *GetIssueException200Response) SetIsExpired(v bool)`

SetIsExpired sets IsExpired field to given value.

### HasIsExpired

`func (o *GetIssueException200Response) HasIsExpired() bool`

HasIsExpired returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssueException200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssueException200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssueException200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssueException200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetIssueException200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetIssueException200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetIssueException200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetIssueException200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


