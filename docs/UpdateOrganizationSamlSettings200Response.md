# UpdateOrganizationSamlSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**EntryPoint** | Pointer to **string** |  | [optional] 
**Cert** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **string** |  | [optional] 
**OrgRoleManagement** | Pointer to **string** |  | [optional] 
**TeamRoleManagement** | Pointer to **string** |  | [optional] 
**CreateMissingTeams** | Pointer to **bool** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**LoginUrl** | Pointer to **string** |  | [optional] 
**AudienceUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateOrganizationSamlSettings200Response

`func NewUpdateOrganizationSamlSettings200Response() *UpdateOrganizationSamlSettings200Response`

NewUpdateOrganizationSamlSettings200Response instantiates a new UpdateOrganizationSamlSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSamlSettings200ResponseWithDefaults

`func NewUpdateOrganizationSamlSettings200ResponseWithDefaults() *UpdateOrganizationSamlSettings200Response`

NewUpdateOrganizationSamlSettings200ResponseWithDefaults instantiates a new UpdateOrganizationSamlSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrganizationSamlSettings200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationSamlSettings200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationSamlSettings200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationSamlSettings200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UpdateOrganizationSamlSettings200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdateOrganizationSamlSettings200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdateOrganizationSamlSettings200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UpdateOrganizationSamlSettings200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetEntryPoint

`func (o *UpdateOrganizationSamlSettings200Response) GetEntryPoint() string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *UpdateOrganizationSamlSettings200Response) GetEntryPointOk() (*string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *UpdateOrganizationSamlSettings200Response) SetEntryPoint(v string)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *UpdateOrganizationSamlSettings200Response) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### GetCert

`func (o *UpdateOrganizationSamlSettings200Response) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UpdateOrganizationSamlSettings200Response) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UpdateOrganizationSamlSettings200Response) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *UpdateOrganizationSamlSettings200Response) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetAudience

`func (o *UpdateOrganizationSamlSettings200Response) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *UpdateOrganizationSamlSettings200Response) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *UpdateOrganizationSamlSettings200Response) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *UpdateOrganizationSamlSettings200Response) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetOrgRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) GetOrgRoleManagement() string`

GetOrgRoleManagement returns the OrgRoleManagement field if non-nil, zero value otherwise.

### GetOrgRoleManagementOk

`func (o *UpdateOrganizationSamlSettings200Response) GetOrgRoleManagementOk() (*string, bool)`

GetOrgRoleManagementOk returns a tuple with the OrgRoleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) SetOrgRoleManagement(v string)`

SetOrgRoleManagement sets OrgRoleManagement field to given value.

### HasOrgRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) HasOrgRoleManagement() bool`

HasOrgRoleManagement returns a boolean if a field has been set.

### GetTeamRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) GetTeamRoleManagement() string`

GetTeamRoleManagement returns the TeamRoleManagement field if non-nil, zero value otherwise.

### GetTeamRoleManagementOk

`func (o *UpdateOrganizationSamlSettings200Response) GetTeamRoleManagementOk() (*string, bool)`

GetTeamRoleManagementOk returns a tuple with the TeamRoleManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) SetTeamRoleManagement(v string)`

SetTeamRoleManagement sets TeamRoleManagement field to given value.

### HasTeamRoleManagement

`func (o *UpdateOrganizationSamlSettings200Response) HasTeamRoleManagement() bool`

HasTeamRoleManagement returns a boolean if a field has been set.

### GetCreateMissingTeams

`func (o *UpdateOrganizationSamlSettings200Response) GetCreateMissingTeams() bool`

GetCreateMissingTeams returns the CreateMissingTeams field if non-nil, zero value otherwise.

### GetCreateMissingTeamsOk

`func (o *UpdateOrganizationSamlSettings200Response) GetCreateMissingTeamsOk() (*bool, bool)`

GetCreateMissingTeamsOk returns a tuple with the CreateMissingTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMissingTeams

`func (o *UpdateOrganizationSamlSettings200Response) SetCreateMissingTeams(v bool)`

SetCreateMissingTeams sets CreateMissingTeams field to given value.

### HasCreateMissingTeams

`func (o *UpdateOrganizationSamlSettings200Response) HasCreateMissingTeams() bool`

HasCreateMissingTeams returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *UpdateOrganizationSamlSettings200Response) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *UpdateOrganizationSamlSettings200Response) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *UpdateOrganizationSamlSettings200Response) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *UpdateOrganizationSamlSettings200Response) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetLoginUrl

`func (o *UpdateOrganizationSamlSettings200Response) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *UpdateOrganizationSamlSettings200Response) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *UpdateOrganizationSamlSettings200Response) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *UpdateOrganizationSamlSettings200Response) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetAudienceUrl

`func (o *UpdateOrganizationSamlSettings200Response) GetAudienceUrl() string`

GetAudienceUrl returns the AudienceUrl field if non-nil, zero value otherwise.

### GetAudienceUrlOk

`func (o *UpdateOrganizationSamlSettings200Response) GetAudienceUrlOk() (*string, bool)`

GetAudienceUrlOk returns a tuple with the AudienceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceUrl

`func (o *UpdateOrganizationSamlSettings200Response) SetAudienceUrl(v string)`

SetAudienceUrl sets AudienceUrl field to given value.

### HasAudienceUrl

`func (o *UpdateOrganizationSamlSettings200Response) HasAudienceUrl() bool`

HasAudienceUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UpdateOrganizationSamlSettings200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateOrganizationSamlSettings200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateOrganizationSamlSettings200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UpdateOrganizationSamlSettings200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UpdateOrganizationSamlSettings200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateOrganizationSamlSettings200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateOrganizationSamlSettings200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateOrganizationSamlSettings200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


