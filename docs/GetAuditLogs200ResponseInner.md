# GetAuditLogs200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique identifier for the audit log. | [optional] 
**ActingUserId** | Pointer to **NullableInt32** | The unique identifier for the user who performed the action. | [optional] 
**ActingUserEmail** | Pointer to **NullableString** | The email address of the user who performed the action. | [optional] 
**ActingUserName** | Pointer to **NullableString** | The name of the user who performed the action. | [optional] 
**ActingUserRole** | Pointer to **NullableString** | The role of the user who performed the action. | [optional] 
**OrganizationId** | Pointer to **NullableInt32** | The unique identifier for the organization. | [optional] 
**UserId** | Pointer to **NullableInt32** | The unique identifier for the user. | [optional] 
**TeamId** | Pointer to **NullableInt32** | The unique identifier for the team. | [optional] 
**BuildId** | Pointer to **NullableInt32** | The unique identifier for the build. | [optional] 
**DependencyId** | Pointer to **NullableString** | The unique identifier for the dependency. | [optional] 
**LicenseId** | Pointer to **NullableString** | The unique identifier for the license. | [optional] 
**PolicyId** | Pointer to **NullableInt32** | The unique identifier for the policy. | [optional] 
**ProjectId** | Pointer to **NullableString** | The unique identifier for the project. | [optional] 
**RuleId** | Pointer to **NullableInt32** | The unique identifier for the rule. | [optional] 
**Locator** | Pointer to **NullableString** | The locator for the dependency. | [optional] 
**RevisionLicenseId** | Pointer to **NullableInt32** | The unique identifier for the revision license. | [optional] 
**IssueId** | Pointer to **NullableInt32** | The unique identifier for the issue. | [optional] 
**Action** | Pointer to **string** | The action that was performed. | [optional] 
**Topic** | Pointer to **string** | The topic of the action. | [optional] 
**Name** | Pointer to **NullableString** | The name of the action. | [optional] 
**OldValue** | Pointer to **NullableString** | The old value before the change. | [optional] 
**NewValue** | Pointer to **NullableString** | The new value after the change. | [optional] 
**Description** | Pointer to **NullableString** | The description of the action. | [optional] 
**CreatedAt** | Pointer to **string** | The date and time the action was performed. | [optional] 

## Methods

### NewGetAuditLogs200ResponseInner

`func NewGetAuditLogs200ResponseInner() *GetAuditLogs200ResponseInner`

NewGetAuditLogs200ResponseInner instantiates a new GetAuditLogs200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogs200ResponseInnerWithDefaults

`func NewGetAuditLogs200ResponseInnerWithDefaults() *GetAuditLogs200ResponseInner`

NewGetAuditLogs200ResponseInnerWithDefaults instantiates a new GetAuditLogs200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAuditLogs200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAuditLogs200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAuditLogs200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAuditLogs200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActingUserId

`func (o *GetAuditLogs200ResponseInner) GetActingUserId() int32`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *GetAuditLogs200ResponseInner) GetActingUserIdOk() (*int32, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *GetAuditLogs200ResponseInner) SetActingUserId(v int32)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *GetAuditLogs200ResponseInner) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.

### SetActingUserIdNil

`func (o *GetAuditLogs200ResponseInner) SetActingUserIdNil(b bool)`

 SetActingUserIdNil sets the value for ActingUserId to be an explicit nil

### UnsetActingUserId
`func (o *GetAuditLogs200ResponseInner) UnsetActingUserId()`

UnsetActingUserId ensures that no value is present for ActingUserId, not even an explicit nil
### GetActingUserEmail

`func (o *GetAuditLogs200ResponseInner) GetActingUserEmail() string`

GetActingUserEmail returns the ActingUserEmail field if non-nil, zero value otherwise.

### GetActingUserEmailOk

`func (o *GetAuditLogs200ResponseInner) GetActingUserEmailOk() (*string, bool)`

GetActingUserEmailOk returns a tuple with the ActingUserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserEmail

`func (o *GetAuditLogs200ResponseInner) SetActingUserEmail(v string)`

SetActingUserEmail sets ActingUserEmail field to given value.

### HasActingUserEmail

`func (o *GetAuditLogs200ResponseInner) HasActingUserEmail() bool`

HasActingUserEmail returns a boolean if a field has been set.

### SetActingUserEmailNil

`func (o *GetAuditLogs200ResponseInner) SetActingUserEmailNil(b bool)`

 SetActingUserEmailNil sets the value for ActingUserEmail to be an explicit nil

### UnsetActingUserEmail
`func (o *GetAuditLogs200ResponseInner) UnsetActingUserEmail()`

UnsetActingUserEmail ensures that no value is present for ActingUserEmail, not even an explicit nil
### GetActingUserName

`func (o *GetAuditLogs200ResponseInner) GetActingUserName() string`

GetActingUserName returns the ActingUserName field if non-nil, zero value otherwise.

### GetActingUserNameOk

`func (o *GetAuditLogs200ResponseInner) GetActingUserNameOk() (*string, bool)`

GetActingUserNameOk returns a tuple with the ActingUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserName

`func (o *GetAuditLogs200ResponseInner) SetActingUserName(v string)`

SetActingUserName sets ActingUserName field to given value.

### HasActingUserName

`func (o *GetAuditLogs200ResponseInner) HasActingUserName() bool`

HasActingUserName returns a boolean if a field has been set.

### SetActingUserNameNil

`func (o *GetAuditLogs200ResponseInner) SetActingUserNameNil(b bool)`

 SetActingUserNameNil sets the value for ActingUserName to be an explicit nil

### UnsetActingUserName
`func (o *GetAuditLogs200ResponseInner) UnsetActingUserName()`

UnsetActingUserName ensures that no value is present for ActingUserName, not even an explicit nil
### GetActingUserRole

`func (o *GetAuditLogs200ResponseInner) GetActingUserRole() string`

GetActingUserRole returns the ActingUserRole field if non-nil, zero value otherwise.

### GetActingUserRoleOk

`func (o *GetAuditLogs200ResponseInner) GetActingUserRoleOk() (*string, bool)`

GetActingUserRoleOk returns a tuple with the ActingUserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserRole

`func (o *GetAuditLogs200ResponseInner) SetActingUserRole(v string)`

SetActingUserRole sets ActingUserRole field to given value.

### HasActingUserRole

`func (o *GetAuditLogs200ResponseInner) HasActingUserRole() bool`

HasActingUserRole returns a boolean if a field has been set.

### SetActingUserRoleNil

`func (o *GetAuditLogs200ResponseInner) SetActingUserRoleNil(b bool)`

 SetActingUserRoleNil sets the value for ActingUserRole to be an explicit nil

### UnsetActingUserRole
`func (o *GetAuditLogs200ResponseInner) UnsetActingUserRole()`

UnsetActingUserRole ensures that no value is present for ActingUserRole, not even an explicit nil
### GetOrganizationId

`func (o *GetAuditLogs200ResponseInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetAuditLogs200ResponseInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetAuditLogs200ResponseInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetAuditLogs200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *GetAuditLogs200ResponseInner) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *GetAuditLogs200ResponseInner) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetUserId

`func (o *GetAuditLogs200ResponseInner) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetAuditLogs200ResponseInner) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetAuditLogs200ResponseInner) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetAuditLogs200ResponseInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetAuditLogs200ResponseInner) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetAuditLogs200ResponseInner) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetTeamId

`func (o *GetAuditLogs200ResponseInner) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GetAuditLogs200ResponseInner) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GetAuditLogs200ResponseInner) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GetAuditLogs200ResponseInner) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### SetTeamIdNil

`func (o *GetAuditLogs200ResponseInner) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *GetAuditLogs200ResponseInner) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil
### GetBuildId

`func (o *GetAuditLogs200ResponseInner) GetBuildId() int32`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *GetAuditLogs200ResponseInner) GetBuildIdOk() (*int32, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *GetAuditLogs200ResponseInner) SetBuildId(v int32)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *GetAuditLogs200ResponseInner) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### SetBuildIdNil

`func (o *GetAuditLogs200ResponseInner) SetBuildIdNil(b bool)`

 SetBuildIdNil sets the value for BuildId to be an explicit nil

### UnsetBuildId
`func (o *GetAuditLogs200ResponseInner) UnsetBuildId()`

UnsetBuildId ensures that no value is present for BuildId, not even an explicit nil
### GetDependencyId

`func (o *GetAuditLogs200ResponseInner) GetDependencyId() string`

GetDependencyId returns the DependencyId field if non-nil, zero value otherwise.

### GetDependencyIdOk

`func (o *GetAuditLogs200ResponseInner) GetDependencyIdOk() (*string, bool)`

GetDependencyIdOk returns a tuple with the DependencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyId

`func (o *GetAuditLogs200ResponseInner) SetDependencyId(v string)`

SetDependencyId sets DependencyId field to given value.

### HasDependencyId

`func (o *GetAuditLogs200ResponseInner) HasDependencyId() bool`

HasDependencyId returns a boolean if a field has been set.

### SetDependencyIdNil

`func (o *GetAuditLogs200ResponseInner) SetDependencyIdNil(b bool)`

 SetDependencyIdNil sets the value for DependencyId to be an explicit nil

### UnsetDependencyId
`func (o *GetAuditLogs200ResponseInner) UnsetDependencyId()`

UnsetDependencyId ensures that no value is present for DependencyId, not even an explicit nil
### GetLicenseId

`func (o *GetAuditLogs200ResponseInner) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *GetAuditLogs200ResponseInner) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *GetAuditLogs200ResponseInner) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *GetAuditLogs200ResponseInner) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### SetLicenseIdNil

`func (o *GetAuditLogs200ResponseInner) SetLicenseIdNil(b bool)`

 SetLicenseIdNil sets the value for LicenseId to be an explicit nil

### UnsetLicenseId
`func (o *GetAuditLogs200ResponseInner) UnsetLicenseId()`

UnsetLicenseId ensures that no value is present for LicenseId, not even an explicit nil
### GetPolicyId

`func (o *GetAuditLogs200ResponseInner) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GetAuditLogs200ResponseInner) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GetAuditLogs200ResponseInner) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *GetAuditLogs200ResponseInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *GetAuditLogs200ResponseInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *GetAuditLogs200ResponseInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetProjectId

`func (o *GetAuditLogs200ResponseInner) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetAuditLogs200ResponseInner) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetAuditLogs200ResponseInner) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetAuditLogs200ResponseInner) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *GetAuditLogs200ResponseInner) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *GetAuditLogs200ResponseInner) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetRuleId

`func (o *GetAuditLogs200ResponseInner) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetAuditLogs200ResponseInner) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetAuditLogs200ResponseInner) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GetAuditLogs200ResponseInner) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *GetAuditLogs200ResponseInner) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *GetAuditLogs200ResponseInner) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetLocator

`func (o *GetAuditLogs200ResponseInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetAuditLogs200ResponseInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetAuditLogs200ResponseInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetAuditLogs200ResponseInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### SetLocatorNil

`func (o *GetAuditLogs200ResponseInner) SetLocatorNil(b bool)`

 SetLocatorNil sets the value for Locator to be an explicit nil

### UnsetLocator
`func (o *GetAuditLogs200ResponseInner) UnsetLocator()`

UnsetLocator ensures that no value is present for Locator, not even an explicit nil
### GetRevisionLicenseId

`func (o *GetAuditLogs200ResponseInner) GetRevisionLicenseId() int32`

GetRevisionLicenseId returns the RevisionLicenseId field if non-nil, zero value otherwise.

### GetRevisionLicenseIdOk

`func (o *GetAuditLogs200ResponseInner) GetRevisionLicenseIdOk() (*int32, bool)`

GetRevisionLicenseIdOk returns a tuple with the RevisionLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionLicenseId

`func (o *GetAuditLogs200ResponseInner) SetRevisionLicenseId(v int32)`

SetRevisionLicenseId sets RevisionLicenseId field to given value.

### HasRevisionLicenseId

`func (o *GetAuditLogs200ResponseInner) HasRevisionLicenseId() bool`

HasRevisionLicenseId returns a boolean if a field has been set.

### SetRevisionLicenseIdNil

`func (o *GetAuditLogs200ResponseInner) SetRevisionLicenseIdNil(b bool)`

 SetRevisionLicenseIdNil sets the value for RevisionLicenseId to be an explicit nil

### UnsetRevisionLicenseId
`func (o *GetAuditLogs200ResponseInner) UnsetRevisionLicenseId()`

UnsetRevisionLicenseId ensures that no value is present for RevisionLicenseId, not even an explicit nil
### GetIssueId

`func (o *GetAuditLogs200ResponseInner) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *GetAuditLogs200ResponseInner) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *GetAuditLogs200ResponseInner) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *GetAuditLogs200ResponseInner) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### SetIssueIdNil

`func (o *GetAuditLogs200ResponseInner) SetIssueIdNil(b bool)`

 SetIssueIdNil sets the value for IssueId to be an explicit nil

### UnsetIssueId
`func (o *GetAuditLogs200ResponseInner) UnsetIssueId()`

UnsetIssueId ensures that no value is present for IssueId, not even an explicit nil
### GetAction

`func (o *GetAuditLogs200ResponseInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAuditLogs200ResponseInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAuditLogs200ResponseInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetAuditLogs200ResponseInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetTopic

`func (o *GetAuditLogs200ResponseInner) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *GetAuditLogs200ResponseInner) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *GetAuditLogs200ResponseInner) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *GetAuditLogs200ResponseInner) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetName

`func (o *GetAuditLogs200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAuditLogs200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAuditLogs200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAuditLogs200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetAuditLogs200ResponseInner) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetAuditLogs200ResponseInner) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOldValue

`func (o *GetAuditLogs200ResponseInner) GetOldValue() string`

GetOldValue returns the OldValue field if non-nil, zero value otherwise.

### GetOldValueOk

`func (o *GetAuditLogs200ResponseInner) GetOldValueOk() (*string, bool)`

GetOldValueOk returns a tuple with the OldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldValue

`func (o *GetAuditLogs200ResponseInner) SetOldValue(v string)`

SetOldValue sets OldValue field to given value.

### HasOldValue

`func (o *GetAuditLogs200ResponseInner) HasOldValue() bool`

HasOldValue returns a boolean if a field has been set.

### SetOldValueNil

`func (o *GetAuditLogs200ResponseInner) SetOldValueNil(b bool)`

 SetOldValueNil sets the value for OldValue to be an explicit nil

### UnsetOldValue
`func (o *GetAuditLogs200ResponseInner) UnsetOldValue()`

UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
### GetNewValue

`func (o *GetAuditLogs200ResponseInner) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *GetAuditLogs200ResponseInner) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *GetAuditLogs200ResponseInner) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *GetAuditLogs200ResponseInner) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### SetNewValueNil

`func (o *GetAuditLogs200ResponseInner) SetNewValueNil(b bool)`

 SetNewValueNil sets the value for NewValue to be an explicit nil

### UnsetNewValue
`func (o *GetAuditLogs200ResponseInner) UnsetNewValue()`

UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
### GetDescription

`func (o *GetAuditLogs200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAuditLogs200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAuditLogs200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAuditLogs200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetAuditLogs200ResponseInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetAuditLogs200ResponseInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *GetAuditLogs200ResponseInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAuditLogs200ResponseInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAuditLogs200ResponseInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAuditLogs200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


