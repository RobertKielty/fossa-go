# GetAllUsersV2200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The user&#39;s unique identifier | [optional] 
**Username** | Pointer to **NullableString** | The user&#39;s username | [optional] 
**Email** | Pointer to **NullableString** | The user&#39;s email address | [optional] 
**EmailVerified** | Pointer to **NullableBool** | Whether the user&#39;s email address has been verified | [optional] 
**Demo** | Pointer to **bool** | Whether the user is a demo user | [optional] 
**Super** | Pointer to **bool** | Whether the user is a super user | [optional] 
**Joined** | Pointer to **NullableTime** | The date the user joined the organization | [optional] 
**LastVisit** | Pointer to **NullableTime** | The date the user last visited the organization | [optional] 
**TermsAgreed** | Pointer to **NullableTime** | The date the user agreed to the organization&#39;s terms | [optional] 
**FullName** | Pointer to **NullableString** | The user&#39;s full name | [optional] 
**Phone** | Pointer to **NullableString** | The user&#39;s phone number | [optional] 
**Role** | Pointer to **NullableString** | The user&#39;s role in the organization | [optional] 
**OrganizationId** | Pointer to **int32** | The organization the user belongs to | [optional] 
**SsoOnly** | Pointer to **bool** | Whether the user is SSO only | [optional] 
**Enabled** | Pointer to **bool** | Whether the user is enabled | [optional] 
**HasSetPassword** | Pointer to **NullableBool** | Whether the user has set a password | [optional] 
**InstallAdmin** | Pointer to **NullableBool** | Whether the user is an install admin | [optional] 
**IsServiceAccount** | Pointer to **bool** | Whether the user is a service account | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date the user was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date the user was last updated | [optional] 
**UserRole** | Pointer to [**GetAllUsersV2200ResponseResultsInnerUserRole**](GetAllUsersV2200ResponseResultsInnerUserRole.md) |  | [optional] 
**Tokens** | Pointer to [**[]GetAllUsersV2200ResponseResultsInnerTokensInner**](GetAllUsersV2200ResponseResultsInnerTokensInner.md) |  | [optional] 
**TeamsCount** | Pointer to **int32** | Number of teams the user belongs to | [optional] 
**Organization** | Pointer to [**GetAllUsers200ResponseInnerOrganization**](GetAllUsers200ResponseInnerOrganization.md) |  | [optional] 
**Github** | Pointer to [**GetAllUsers200ResponseInnerGithub**](GetAllUsers200ResponseInnerGithub.md) |  | [optional] 
**BitbucketCloud** | Pointer to [**GetAllUsers200ResponseInnerBitbucketCloud**](GetAllUsers200ResponseInnerBitbucketCloud.md) |  | [optional] 

## Methods

### NewGetAllUsersV2200ResponseResultsInner

`func NewGetAllUsersV2200ResponseResultsInner() *GetAllUsersV2200ResponseResultsInner`

NewGetAllUsersV2200ResponseResultsInner instantiates a new GetAllUsersV2200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllUsersV2200ResponseResultsInnerWithDefaults

`func NewGetAllUsersV2200ResponseResultsInnerWithDefaults() *GetAllUsersV2200ResponseResultsInner`

NewGetAllUsersV2200ResponseResultsInnerWithDefaults instantiates a new GetAllUsersV2200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllUsersV2200ResponseResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllUsersV2200ResponseResultsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAllUsersV2200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *GetAllUsersV2200ResponseResultsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAllUsersV2200ResponseResultsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAllUsersV2200ResponseResultsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *GetAllUsersV2200ResponseResultsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAllUsersV2200ResponseResultsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAllUsersV2200ResponseResultsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *GetAllUsersV2200ResponseResultsInner) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *GetAllUsersV2200ResponseResultsInner) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *GetAllUsersV2200ResponseResultsInner) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetDemo

`func (o *GetAllUsersV2200ResponseResultsInner) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *GetAllUsersV2200ResponseResultsInner) SetDemo(v bool)`

SetDemo sets Demo field to given value.

### HasDemo

`func (o *GetAllUsersV2200ResponseResultsInner) HasDemo() bool`

HasDemo returns a boolean if a field has been set.

### GetSuper

`func (o *GetAllUsersV2200ResponseResultsInner) GetSuper() bool`

GetSuper returns the Super field if non-nil, zero value otherwise.

### GetSuperOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetSuperOk() (*bool, bool)`

GetSuperOk returns a tuple with the Super field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuper

`func (o *GetAllUsersV2200ResponseResultsInner) SetSuper(v bool)`

SetSuper sets Super field to given value.

### HasSuper

`func (o *GetAllUsersV2200ResponseResultsInner) HasSuper() bool`

HasSuper returns a boolean if a field has been set.

### GetJoined

`func (o *GetAllUsersV2200ResponseResultsInner) GetJoined() time.Time`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetJoinedOk() (*time.Time, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *GetAllUsersV2200ResponseResultsInner) SetJoined(v time.Time)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *GetAllUsersV2200ResponseResultsInner) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### SetJoinedNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetJoinedNil(b bool)`

 SetJoinedNil sets the value for Joined to be an explicit nil

### UnsetJoined
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetJoined()`

UnsetJoined ensures that no value is present for Joined, not even an explicit nil
### GetLastVisit

`func (o *GetAllUsersV2200ResponseResultsInner) GetLastVisit() time.Time`

GetLastVisit returns the LastVisit field if non-nil, zero value otherwise.

### GetLastVisitOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetLastVisitOk() (*time.Time, bool)`

GetLastVisitOk returns a tuple with the LastVisit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVisit

`func (o *GetAllUsersV2200ResponseResultsInner) SetLastVisit(v time.Time)`

SetLastVisit sets LastVisit field to given value.

### HasLastVisit

`func (o *GetAllUsersV2200ResponseResultsInner) HasLastVisit() bool`

HasLastVisit returns a boolean if a field has been set.

### SetLastVisitNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetLastVisitNil(b bool)`

 SetLastVisitNil sets the value for LastVisit to be an explicit nil

### UnsetLastVisit
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetLastVisit()`

UnsetLastVisit ensures that no value is present for LastVisit, not even an explicit nil
### GetTermsAgreed

`func (o *GetAllUsersV2200ResponseResultsInner) GetTermsAgreed() time.Time`

GetTermsAgreed returns the TermsAgreed field if non-nil, zero value otherwise.

### GetTermsAgreedOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetTermsAgreedOk() (*time.Time, bool)`

GetTermsAgreedOk returns a tuple with the TermsAgreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAgreed

`func (o *GetAllUsersV2200ResponseResultsInner) SetTermsAgreed(v time.Time)`

SetTermsAgreed sets TermsAgreed field to given value.

### HasTermsAgreed

`func (o *GetAllUsersV2200ResponseResultsInner) HasTermsAgreed() bool`

HasTermsAgreed returns a boolean if a field has been set.

### SetTermsAgreedNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetTermsAgreedNil(b bool)`

 SetTermsAgreedNil sets the value for TermsAgreed to be an explicit nil

### UnsetTermsAgreed
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetTermsAgreed()`

UnsetTermsAgreed ensures that no value is present for TermsAgreed, not even an explicit nil
### GetFullName

`func (o *GetAllUsersV2200ResponseResultsInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetAllUsersV2200ResponseResultsInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetAllUsersV2200ResponseResultsInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPhone

`func (o *GetAllUsersV2200ResponseResultsInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetAllUsersV2200ResponseResultsInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetAllUsersV2200ResponseResultsInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetRole

`func (o *GetAllUsersV2200ResponseResultsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetAllUsersV2200ResponseResultsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetAllUsersV2200ResponseResultsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetOrganizationId

`func (o *GetAllUsersV2200ResponseResultsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetAllUsersV2200ResponseResultsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetAllUsersV2200ResponseResultsInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSsoOnly

`func (o *GetAllUsersV2200ResponseResultsInner) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *GetAllUsersV2200ResponseResultsInner) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *GetAllUsersV2200ResponseResultsInner) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.

### GetEnabled

`func (o *GetAllUsersV2200ResponseResultsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAllUsersV2200ResponseResultsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetAllUsersV2200ResponseResultsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasSetPassword

`func (o *GetAllUsersV2200ResponseResultsInner) GetHasSetPassword() bool`

GetHasSetPassword returns the HasSetPassword field if non-nil, zero value otherwise.

### GetHasSetPasswordOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetHasSetPasswordOk() (*bool, bool)`

GetHasSetPasswordOk returns a tuple with the HasSetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSetPassword

`func (o *GetAllUsersV2200ResponseResultsInner) SetHasSetPassword(v bool)`

SetHasSetPassword sets HasSetPassword field to given value.

### HasHasSetPassword

`func (o *GetAllUsersV2200ResponseResultsInner) HasHasSetPassword() bool`

HasHasSetPassword returns a boolean if a field has been set.

### SetHasSetPasswordNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetHasSetPasswordNil(b bool)`

 SetHasSetPasswordNil sets the value for HasSetPassword to be an explicit nil

### UnsetHasSetPassword
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetHasSetPassword()`

UnsetHasSetPassword ensures that no value is present for HasSetPassword, not even an explicit nil
### GetInstallAdmin

`func (o *GetAllUsersV2200ResponseResultsInner) GetInstallAdmin() bool`

GetInstallAdmin returns the InstallAdmin field if non-nil, zero value otherwise.

### GetInstallAdminOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetInstallAdminOk() (*bool, bool)`

GetInstallAdminOk returns a tuple with the InstallAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAdmin

`func (o *GetAllUsersV2200ResponseResultsInner) SetInstallAdmin(v bool)`

SetInstallAdmin sets InstallAdmin field to given value.

### HasInstallAdmin

`func (o *GetAllUsersV2200ResponseResultsInner) HasInstallAdmin() bool`

HasInstallAdmin returns a boolean if a field has been set.

### SetInstallAdminNil

`func (o *GetAllUsersV2200ResponseResultsInner) SetInstallAdminNil(b bool)`

 SetInstallAdminNil sets the value for InstallAdmin to be an explicit nil

### UnsetInstallAdmin
`func (o *GetAllUsersV2200ResponseResultsInner) UnsetInstallAdmin()`

UnsetInstallAdmin ensures that no value is present for InstallAdmin, not even an explicit nil
### GetIsServiceAccount

`func (o *GetAllUsersV2200ResponseResultsInner) GetIsServiceAccount() bool`

GetIsServiceAccount returns the IsServiceAccount field if non-nil, zero value otherwise.

### GetIsServiceAccountOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetIsServiceAccountOk() (*bool, bool)`

GetIsServiceAccountOk returns a tuple with the IsServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsServiceAccount

`func (o *GetAllUsersV2200ResponseResultsInner) SetIsServiceAccount(v bool)`

SetIsServiceAccount sets IsServiceAccount field to given value.

### HasIsServiceAccount

`func (o *GetAllUsersV2200ResponseResultsInner) HasIsServiceAccount() bool`

HasIsServiceAccount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetAllUsersV2200ResponseResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserRole

`func (o *GetAllUsersV2200ResponseResultsInner) GetUserRole() GetAllUsersV2200ResponseResultsInnerUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetUserRoleOk() (*GetAllUsersV2200ResponseResultsInnerUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *GetAllUsersV2200ResponseResultsInner) SetUserRole(v GetAllUsersV2200ResponseResultsInnerUserRole)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *GetAllUsersV2200ResponseResultsInner) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetTokens

`func (o *GetAllUsersV2200ResponseResultsInner) GetTokens() []GetAllUsersV2200ResponseResultsInnerTokensInner`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetTokensOk() (*[]GetAllUsersV2200ResponseResultsInnerTokensInner, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GetAllUsersV2200ResponseResultsInner) SetTokens(v []GetAllUsersV2200ResponseResultsInnerTokensInner)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *GetAllUsersV2200ResponseResultsInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTeamsCount

`func (o *GetAllUsersV2200ResponseResultsInner) GetTeamsCount() int32`

GetTeamsCount returns the TeamsCount field if non-nil, zero value otherwise.

### GetTeamsCountOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetTeamsCountOk() (*int32, bool)`

GetTeamsCountOk returns a tuple with the TeamsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsCount

`func (o *GetAllUsersV2200ResponseResultsInner) SetTeamsCount(v int32)`

SetTeamsCount sets TeamsCount field to given value.

### HasTeamsCount

`func (o *GetAllUsersV2200ResponseResultsInner) HasTeamsCount() bool`

HasTeamsCount returns a boolean if a field has been set.

### GetOrganization

`func (o *GetAllUsersV2200ResponseResultsInner) GetOrganization() GetAllUsers200ResponseInnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetOrganizationOk() (*GetAllUsers200ResponseInnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetAllUsersV2200ResponseResultsInner) SetOrganization(v GetAllUsers200ResponseInnerOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetAllUsersV2200ResponseResultsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetGithub

`func (o *GetAllUsersV2200ResponseResultsInner) GetGithub() GetAllUsers200ResponseInnerGithub`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetGithubOk() (*GetAllUsers200ResponseInnerGithub, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GetAllUsersV2200ResponseResultsInner) SetGithub(v GetAllUsers200ResponseInnerGithub)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GetAllUsersV2200ResponseResultsInner) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetBitbucketCloud

`func (o *GetAllUsersV2200ResponseResultsInner) GetBitbucketCloud() GetAllUsers200ResponseInnerBitbucketCloud`

GetBitbucketCloud returns the BitbucketCloud field if non-nil, zero value otherwise.

### GetBitbucketCloudOk

`func (o *GetAllUsersV2200ResponseResultsInner) GetBitbucketCloudOk() (*GetAllUsers200ResponseInnerBitbucketCloud, bool)`

GetBitbucketCloudOk returns a tuple with the BitbucketCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketCloud

`func (o *GetAllUsersV2200ResponseResultsInner) SetBitbucketCloud(v GetAllUsers200ResponseInnerBitbucketCloud)`

SetBitbucketCloud sets BitbucketCloud field to given value.

### HasBitbucketCloud

`func (o *GetAllUsersV2200ResponseResultsInner) HasBitbucketCloud() bool`

HasBitbucketCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


