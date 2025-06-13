# GetAllUsers200ResponseInner

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
**CreatedAt** | Pointer to **time.Time** | The date the user was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date the user was last updated | [optional] 
**UserRole** | Pointer to [**GetAllUsers200ResponseInnerUserRole**](GetAllUsers200ResponseInnerUserRole.md) |  | [optional] 
**Tokens** | Pointer to [**[]GetAllUsers200ResponseInnerTokensInner**](GetAllUsers200ResponseInnerTokensInner.md) |  | [optional] 
**TeamUsers** | Pointer to [**[]GetAllUsers200ResponseInnerTeamUsersInner**](GetAllUsers200ResponseInnerTeamUsersInner.md) |  | [optional] 
**Organization** | Pointer to [**GetAllUsers200ResponseInnerOrganization**](GetAllUsers200ResponseInnerOrganization.md) |  | [optional] 
**Github** | Pointer to [**GetAllUsers200ResponseInnerGithub**](GetAllUsers200ResponseInnerGithub.md) |  | [optional] 
**BitbucketCloud** | Pointer to [**GetAllUsers200ResponseInnerBitbucketCloud**](GetAllUsers200ResponseInnerBitbucketCloud.md) |  | [optional] 

## Methods

### NewGetAllUsers200ResponseInner

`func NewGetAllUsers200ResponseInner() *GetAllUsers200ResponseInner`

NewGetAllUsers200ResponseInner instantiates a new GetAllUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllUsers200ResponseInnerWithDefaults

`func NewGetAllUsers200ResponseInnerWithDefaults() *GetAllUsers200ResponseInner`

NewGetAllUsers200ResponseInnerWithDefaults instantiates a new GetAllUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllUsers200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllUsers200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllUsers200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAllUsers200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *GetAllUsers200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAllUsers200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAllUsers200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAllUsers200ResponseInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetAllUsers200ResponseInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetAllUsers200ResponseInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetEmail

`func (o *GetAllUsers200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAllUsers200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAllUsers200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAllUsers200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetAllUsers200ResponseInner) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetAllUsers200ResponseInner) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailVerified

`func (o *GetAllUsers200ResponseInner) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *GetAllUsers200ResponseInner) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *GetAllUsers200ResponseInner) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *GetAllUsers200ResponseInner) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### SetEmailVerifiedNil

`func (o *GetAllUsers200ResponseInner) SetEmailVerifiedNil(b bool)`

 SetEmailVerifiedNil sets the value for EmailVerified to be an explicit nil

### UnsetEmailVerified
`func (o *GetAllUsers200ResponseInner) UnsetEmailVerified()`

UnsetEmailVerified ensures that no value is present for EmailVerified, not even an explicit nil
### GetDemo

`func (o *GetAllUsers200ResponseInner) GetDemo() bool`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *GetAllUsers200ResponseInner) GetDemoOk() (*bool, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *GetAllUsers200ResponseInner) SetDemo(v bool)`

SetDemo sets Demo field to given value.

### HasDemo

`func (o *GetAllUsers200ResponseInner) HasDemo() bool`

HasDemo returns a boolean if a field has been set.

### GetSuper

`func (o *GetAllUsers200ResponseInner) GetSuper() bool`

GetSuper returns the Super field if non-nil, zero value otherwise.

### GetSuperOk

`func (o *GetAllUsers200ResponseInner) GetSuperOk() (*bool, bool)`

GetSuperOk returns a tuple with the Super field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuper

`func (o *GetAllUsers200ResponseInner) SetSuper(v bool)`

SetSuper sets Super field to given value.

### HasSuper

`func (o *GetAllUsers200ResponseInner) HasSuper() bool`

HasSuper returns a boolean if a field has been set.

### GetJoined

`func (o *GetAllUsers200ResponseInner) GetJoined() time.Time`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *GetAllUsers200ResponseInner) GetJoinedOk() (*time.Time, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *GetAllUsers200ResponseInner) SetJoined(v time.Time)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *GetAllUsers200ResponseInner) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### SetJoinedNil

`func (o *GetAllUsers200ResponseInner) SetJoinedNil(b bool)`

 SetJoinedNil sets the value for Joined to be an explicit nil

### UnsetJoined
`func (o *GetAllUsers200ResponseInner) UnsetJoined()`

UnsetJoined ensures that no value is present for Joined, not even an explicit nil
### GetLastVisit

`func (o *GetAllUsers200ResponseInner) GetLastVisit() time.Time`

GetLastVisit returns the LastVisit field if non-nil, zero value otherwise.

### GetLastVisitOk

`func (o *GetAllUsers200ResponseInner) GetLastVisitOk() (*time.Time, bool)`

GetLastVisitOk returns a tuple with the LastVisit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVisit

`func (o *GetAllUsers200ResponseInner) SetLastVisit(v time.Time)`

SetLastVisit sets LastVisit field to given value.

### HasLastVisit

`func (o *GetAllUsers200ResponseInner) HasLastVisit() bool`

HasLastVisit returns a boolean if a field has been set.

### SetLastVisitNil

`func (o *GetAllUsers200ResponseInner) SetLastVisitNil(b bool)`

 SetLastVisitNil sets the value for LastVisit to be an explicit nil

### UnsetLastVisit
`func (o *GetAllUsers200ResponseInner) UnsetLastVisit()`

UnsetLastVisit ensures that no value is present for LastVisit, not even an explicit nil
### GetTermsAgreed

`func (o *GetAllUsers200ResponseInner) GetTermsAgreed() time.Time`

GetTermsAgreed returns the TermsAgreed field if non-nil, zero value otherwise.

### GetTermsAgreedOk

`func (o *GetAllUsers200ResponseInner) GetTermsAgreedOk() (*time.Time, bool)`

GetTermsAgreedOk returns a tuple with the TermsAgreed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsAgreed

`func (o *GetAllUsers200ResponseInner) SetTermsAgreed(v time.Time)`

SetTermsAgreed sets TermsAgreed field to given value.

### HasTermsAgreed

`func (o *GetAllUsers200ResponseInner) HasTermsAgreed() bool`

HasTermsAgreed returns a boolean if a field has been set.

### SetTermsAgreedNil

`func (o *GetAllUsers200ResponseInner) SetTermsAgreedNil(b bool)`

 SetTermsAgreedNil sets the value for TermsAgreed to be an explicit nil

### UnsetTermsAgreed
`func (o *GetAllUsers200ResponseInner) UnsetTermsAgreed()`

UnsetTermsAgreed ensures that no value is present for TermsAgreed, not even an explicit nil
### GetFullName

`func (o *GetAllUsers200ResponseInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *GetAllUsers200ResponseInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *GetAllUsers200ResponseInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *GetAllUsers200ResponseInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *GetAllUsers200ResponseInner) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *GetAllUsers200ResponseInner) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetPhone

`func (o *GetAllUsers200ResponseInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetAllUsers200ResponseInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetAllUsers200ResponseInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetAllUsers200ResponseInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *GetAllUsers200ResponseInner) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *GetAllUsers200ResponseInner) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetRole

`func (o *GetAllUsers200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetAllUsers200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetAllUsers200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetAllUsers200ResponseInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *GetAllUsers200ResponseInner) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *GetAllUsers200ResponseInner) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetOrganizationId

`func (o *GetAllUsers200ResponseInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetAllUsers200ResponseInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetAllUsers200ResponseInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetAllUsers200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetSsoOnly

`func (o *GetAllUsers200ResponseInner) GetSsoOnly() bool`

GetSsoOnly returns the SsoOnly field if non-nil, zero value otherwise.

### GetSsoOnlyOk

`func (o *GetAllUsers200ResponseInner) GetSsoOnlyOk() (*bool, bool)`

GetSsoOnlyOk returns a tuple with the SsoOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoOnly

`func (o *GetAllUsers200ResponseInner) SetSsoOnly(v bool)`

SetSsoOnly sets SsoOnly field to given value.

### HasSsoOnly

`func (o *GetAllUsers200ResponseInner) HasSsoOnly() bool`

HasSsoOnly returns a boolean if a field has been set.

### GetEnabled

`func (o *GetAllUsers200ResponseInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAllUsers200ResponseInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAllUsers200ResponseInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetAllUsers200ResponseInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHasSetPassword

`func (o *GetAllUsers200ResponseInner) GetHasSetPassword() bool`

GetHasSetPassword returns the HasSetPassword field if non-nil, zero value otherwise.

### GetHasSetPasswordOk

`func (o *GetAllUsers200ResponseInner) GetHasSetPasswordOk() (*bool, bool)`

GetHasSetPasswordOk returns a tuple with the HasSetPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSetPassword

`func (o *GetAllUsers200ResponseInner) SetHasSetPassword(v bool)`

SetHasSetPassword sets HasSetPassword field to given value.

### HasHasSetPassword

`func (o *GetAllUsers200ResponseInner) HasHasSetPassword() bool`

HasHasSetPassword returns a boolean if a field has been set.

### SetHasSetPasswordNil

`func (o *GetAllUsers200ResponseInner) SetHasSetPasswordNil(b bool)`

 SetHasSetPasswordNil sets the value for HasSetPassword to be an explicit nil

### UnsetHasSetPassword
`func (o *GetAllUsers200ResponseInner) UnsetHasSetPassword()`

UnsetHasSetPassword ensures that no value is present for HasSetPassword, not even an explicit nil
### GetInstallAdmin

`func (o *GetAllUsers200ResponseInner) GetInstallAdmin() bool`

GetInstallAdmin returns the InstallAdmin field if non-nil, zero value otherwise.

### GetInstallAdminOk

`func (o *GetAllUsers200ResponseInner) GetInstallAdminOk() (*bool, bool)`

GetInstallAdminOk returns a tuple with the InstallAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAdmin

`func (o *GetAllUsers200ResponseInner) SetInstallAdmin(v bool)`

SetInstallAdmin sets InstallAdmin field to given value.

### HasInstallAdmin

`func (o *GetAllUsers200ResponseInner) HasInstallAdmin() bool`

HasInstallAdmin returns a boolean if a field has been set.

### SetInstallAdminNil

`func (o *GetAllUsers200ResponseInner) SetInstallAdminNil(b bool)`

 SetInstallAdminNil sets the value for InstallAdmin to be an explicit nil

### UnsetInstallAdmin
`func (o *GetAllUsers200ResponseInner) UnsetInstallAdmin()`

UnsetInstallAdmin ensures that no value is present for InstallAdmin, not even an explicit nil
### GetCreatedAt

`func (o *GetAllUsers200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAllUsers200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAllUsers200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAllUsers200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetAllUsers200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAllUsers200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAllUsers200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetAllUsers200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserRole

`func (o *GetAllUsers200ResponseInner) GetUserRole() GetAllUsers200ResponseInnerUserRole`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *GetAllUsers200ResponseInner) GetUserRoleOk() (*GetAllUsers200ResponseInnerUserRole, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *GetAllUsers200ResponseInner) SetUserRole(v GetAllUsers200ResponseInnerUserRole)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *GetAllUsers200ResponseInner) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetTokens

`func (o *GetAllUsers200ResponseInner) GetTokens() []GetAllUsers200ResponseInnerTokensInner`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *GetAllUsers200ResponseInner) GetTokensOk() (*[]GetAllUsers200ResponseInnerTokensInner, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *GetAllUsers200ResponseInner) SetTokens(v []GetAllUsers200ResponseInnerTokensInner)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *GetAllUsers200ResponseInner) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTeamUsers

`func (o *GetAllUsers200ResponseInner) GetTeamUsers() []GetAllUsers200ResponseInnerTeamUsersInner`

GetTeamUsers returns the TeamUsers field if non-nil, zero value otherwise.

### GetTeamUsersOk

`func (o *GetAllUsers200ResponseInner) GetTeamUsersOk() (*[]GetAllUsers200ResponseInnerTeamUsersInner, bool)`

GetTeamUsersOk returns a tuple with the TeamUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamUsers

`func (o *GetAllUsers200ResponseInner) SetTeamUsers(v []GetAllUsers200ResponseInnerTeamUsersInner)`

SetTeamUsers sets TeamUsers field to given value.

### HasTeamUsers

`func (o *GetAllUsers200ResponseInner) HasTeamUsers() bool`

HasTeamUsers returns a boolean if a field has been set.

### GetOrganization

`func (o *GetAllUsers200ResponseInner) GetOrganization() GetAllUsers200ResponseInnerOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *GetAllUsers200ResponseInner) GetOrganizationOk() (*GetAllUsers200ResponseInnerOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *GetAllUsers200ResponseInner) SetOrganization(v GetAllUsers200ResponseInnerOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *GetAllUsers200ResponseInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetGithub

`func (o *GetAllUsers200ResponseInner) GetGithub() GetAllUsers200ResponseInnerGithub`

GetGithub returns the Github field if non-nil, zero value otherwise.

### GetGithubOk

`func (o *GetAllUsers200ResponseInner) GetGithubOk() (*GetAllUsers200ResponseInnerGithub, bool)`

GetGithubOk returns a tuple with the Github field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithub

`func (o *GetAllUsers200ResponseInner) SetGithub(v GetAllUsers200ResponseInnerGithub)`

SetGithub sets Github field to given value.

### HasGithub

`func (o *GetAllUsers200ResponseInner) HasGithub() bool`

HasGithub returns a boolean if a field has been set.

### GetBitbucketCloud

`func (o *GetAllUsers200ResponseInner) GetBitbucketCloud() GetAllUsers200ResponseInnerBitbucketCloud`

GetBitbucketCloud returns the BitbucketCloud field if non-nil, zero value otherwise.

### GetBitbucketCloudOk

`func (o *GetAllUsers200ResponseInner) GetBitbucketCloudOk() (*GetAllUsers200ResponseInnerBitbucketCloud, bool)`

GetBitbucketCloudOk returns a tuple with the BitbucketCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitbucketCloud

`func (o *GetAllUsers200ResponseInner) SetBitbucketCloud(v GetAllUsers200ResponseInnerBitbucketCloud)`

SetBitbucketCloud sets BitbucketCloud field to given value.

### HasBitbucketCloud

`func (o *GetAllUsers200ResponseInner) HasBitbucketCloud() bool`

HasBitbucketCloud returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


