# UpdateOrganizationSamlSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryPoint** | **string** | Identity Provider Single Sign On URL | 
**Cert** | **string** | X.509 Certificate from the identity provider | 
**Audience** | **string** | Audience URI / SP Entity ID | 
**OrgRoleManagement** | Pointer to **string** | Whether Organization Roles should be managed by FOSSA, the Identity Provider, or both | [optional] [default to "fossa"]
**TeamRoleManagement** | Pointer to **string** | Whether Teams and Team Roles should be managed by FOSSA, the Identity Provider, or both | [optional] [default to "fossa"]
**CreateMissingTeams** | Pointer to **bool** | Whether to automatically create teams that are specified in SAML attributes but don&#39;t exist in the organization yet | [optional] [default to true]

## Methods

### NewUpdateOrganizationSamlSettingsRequest

`func NewUpdateOrganizationSamlSettingsRequest(entryPoint string, cert string, audience string, ) *UpdateOrganizationSamlSettingsRequest`

NewUpdateOrganizationSamlSettingsRequest instantiates a new UpdateOrganizationSamlSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSamlSettingsRequestWithDefaults

`func NewUpdateOrganizationSamlSettingsRequestWithDefaults() *UpdateOrganizationSamlSettingsRequest`

NewUpdateOrganizationSamlSettingsRequestWithDefaults instantiates a new UpdateOrganizationSamlSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryPoint

`func (o *UpdateOrganizationSamlSettingsRequest) GetEntryPoint() string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetEntryPointOk() (*string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *UpdateOrganizationSamlSettingsRequest) SetEntryPoint(v string)`

SetEntryPoint sets EntryPoint field to given value.


### GetCert

`func (o *UpdateOrganizationSamlSettingsRequest) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UpdateOrganizationSamlSettingsRequest) SetCert(v string)`

SetCert sets Cert field to given value.


### GetAudience

`func (o *UpdateOrganizationSamlSettingsRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateOrganizationSamlSettingsRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetOrgRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) GetOrgRoleManagement() string`

GetOrgRoleManagement returns the OrgRoleManagement field if non-nil, zero value otherwise.

### GetOrgRoleManagementOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetOrgRoleManagementOk() (*string, bool)`

GetOrgRoleManagementOk returns a tuple with the OrgRoleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) SetOrgRoleManagement(v string)`

SetOrgRoleManagement sets OrgRoleManagement field to given value.

### HasOrgRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) HasOrgRoleManagement() bool`

HasOrgRoleManagement returns a boolean if a field has been set.

### GetTeamRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) GetTeamRoleManagement() string`

GetTeamRoleManagement returns the TeamRoleManagement field if non-nil, zero value otherwise.

### GetTeamRoleManagementOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetTeamRoleManagementOk() (*string, bool)`

GetTeamRoleManagementOk returns a tuple with the TeamRoleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) SetTeamRoleManagement(v string)`

SetTeamRoleManagement sets TeamRoleManagement field to given value.

### HasTeamRoleManagement

`func (o *UpdateOrganizationSamlSettingsRequest) HasTeamRoleManagement() bool`

HasTeamRoleManagement returns a boolean if a field has been set.

### GetCreateMissingTeams

`func (o *UpdateOrganizationSamlSettingsRequest) GetCreateMissingTeams() bool`

GetCreateMissingTeams returns the CreateMissingTeams field if non-nil, zero value otherwise.

### GetCreateMissingTeamsOk

`func (o *UpdateOrganizationSamlSettingsRequest) GetCreateMissingTeamsOk() (*bool, bool)`

GetCreateMissingTeamsOk returns a tuple with the CreateMissingTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMissingTeams

`func (o *UpdateOrganizationSamlSettingsRequest) SetCreateMissingTeams(v bool)`

SetCreateMissingTeams sets CreateMissingTeams field to given value.

### HasCreateMissingTeams

`func (o *UpdateOrganizationSamlSettingsRequest) HasCreateMissingTeams() bool`

HasCreateMissingTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


