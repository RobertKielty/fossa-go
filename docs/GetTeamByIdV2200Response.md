# GetTeamByIdV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The team&#39;s unique identifier | [optional] 
**OrganizationId** | Pointer to **int32** | The organization&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The team&#39;s name | [optional] 
**DefaultRoleId** | Pointer to **int32** | The default role&#39;s unique identifier | [optional] 
**AutoAddUsers** | Pointer to **bool** | Whether the team auto-adds users | [optional] 
**UniqueIdentifier** | Pointer to **string** | The team&#39;s unique identifier | [optional] 
**CreatedAt** | Pointer to **time.Time** | The team&#39;s creation date | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The team&#39;s last update date | [optional] 
**TeamType** | Pointer to **string** | Whether it&#39;s a normal team or a Team Group. | [optional] 
**TeamMembersCount** | Pointer to **int32** | The number of members in the team | [optional] 
**TeamReleaseGroupsCount** | Pointer to **int32** | The number of release groups associated with the team | [optional] 
**TeamProjectsCount** | Pointer to **int32** | The number of projects associated with the team | [optional] 

## Methods

### NewGetTeamByIdV2200Response

`func NewGetTeamByIdV2200Response() *GetTeamByIdV2200Response`

NewGetTeamByIdV2200Response instantiates a new GetTeamByIdV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamByIdV2200ResponseWithDefaults

`func NewGetTeamByIdV2200ResponseWithDefaults() *GetTeamByIdV2200Response`

NewGetTeamByIdV2200ResponseWithDefaults instantiates a new GetTeamByIdV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTeamByIdV2200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamByIdV2200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamByIdV2200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTeamByIdV2200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetTeamByIdV2200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetTeamByIdV2200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetTeamByIdV2200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetTeamByIdV2200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *GetTeamByIdV2200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTeamByIdV2200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTeamByIdV2200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTeamByIdV2200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *GetTeamByIdV2200Response) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *GetTeamByIdV2200Response) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *GetTeamByIdV2200Response) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *GetTeamByIdV2200Response) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetAutoAddUsers

`func (o *GetTeamByIdV2200Response) GetAutoAddUsers() bool`

GetAutoAddUsers returns the AutoAddUsers field if non-nil, zero value otherwise.

### GetAutoAddUsersOk

`func (o *GetTeamByIdV2200Response) GetAutoAddUsersOk() (*bool, bool)`

GetAutoAddUsersOk returns a tuple with the AutoAddUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddUsers

`func (o *GetTeamByIdV2200Response) SetAutoAddUsers(v bool)`

SetAutoAddUsers sets AutoAddUsers field to given value.

### HasAutoAddUsers

`func (o *GetTeamByIdV2200Response) HasAutoAddUsers() bool`

HasAutoAddUsers returns a boolean if a field has been set.

### GetUniqueIdentifier

`func (o *GetTeamByIdV2200Response) GetUniqueIdentifier() string`

GetUniqueIdentifier returns the UniqueIdentifier field if non-nil, zero value otherwise.

### GetUniqueIdentifierOk

`func (o *GetTeamByIdV2200Response) GetUniqueIdentifierOk() (*string, bool)`

GetUniqueIdentifierOk returns a tuple with the UniqueIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueIdentifier

`func (o *GetTeamByIdV2200Response) SetUniqueIdentifier(v string)`

SetUniqueIdentifier sets UniqueIdentifier field to given value.

### HasUniqueIdentifier

`func (o *GetTeamByIdV2200Response) HasUniqueIdentifier() bool`

HasUniqueIdentifier returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetTeamByIdV2200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTeamByIdV2200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTeamByIdV2200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetTeamByIdV2200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetTeamByIdV2200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetTeamByIdV2200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetTeamByIdV2200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetTeamByIdV2200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTeamType

`func (o *GetTeamByIdV2200Response) GetTeamType() string`

GetTeamType returns the TeamType field if non-nil, zero value otherwise.

### GetTeamTypeOk

`func (o *GetTeamByIdV2200Response) GetTeamTypeOk() (*string, bool)`

GetTeamTypeOk returns a tuple with the TeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamType

`func (o *GetTeamByIdV2200Response) SetTeamType(v string)`

SetTeamType sets TeamType field to given value.

### HasTeamType

`func (o *GetTeamByIdV2200Response) HasTeamType() bool`

HasTeamType returns a boolean if a field has been set.

### GetTeamMembersCount

`func (o *GetTeamByIdV2200Response) GetTeamMembersCount() int32`

GetTeamMembersCount returns the TeamMembersCount field if non-nil, zero value otherwise.

### GetTeamMembersCountOk

`func (o *GetTeamByIdV2200Response) GetTeamMembersCountOk() (*int32, bool)`

GetTeamMembersCountOk returns a tuple with the TeamMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMembersCount

`func (o *GetTeamByIdV2200Response) SetTeamMembersCount(v int32)`

SetTeamMembersCount sets TeamMembersCount field to given value.

### HasTeamMembersCount

`func (o *GetTeamByIdV2200Response) HasTeamMembersCount() bool`

HasTeamMembersCount returns a boolean if a field has been set.

### GetTeamReleaseGroupsCount

`func (o *GetTeamByIdV2200Response) GetTeamReleaseGroupsCount() int32`

GetTeamReleaseGroupsCount returns the TeamReleaseGroupsCount field if non-nil, zero value otherwise.

### GetTeamReleaseGroupsCountOk

`func (o *GetTeamByIdV2200Response) GetTeamReleaseGroupsCountOk() (*int32, bool)`

GetTeamReleaseGroupsCountOk returns a tuple with the TeamReleaseGroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamReleaseGroupsCount

`func (o *GetTeamByIdV2200Response) SetTeamReleaseGroupsCount(v int32)`

SetTeamReleaseGroupsCount sets TeamReleaseGroupsCount field to given value.

### HasTeamReleaseGroupsCount

`func (o *GetTeamByIdV2200Response) HasTeamReleaseGroupsCount() bool`

HasTeamReleaseGroupsCount returns a boolean if a field has been set.

### GetTeamProjectsCount

`func (o *GetTeamByIdV2200Response) GetTeamProjectsCount() int32`

GetTeamProjectsCount returns the TeamProjectsCount field if non-nil, zero value otherwise.

### GetTeamProjectsCountOk

`func (o *GetTeamByIdV2200Response) GetTeamProjectsCountOk() (*int32, bool)`

GetTeamProjectsCountOk returns a tuple with the TeamProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamProjectsCount

`func (o *GetTeamByIdV2200Response) SetTeamProjectsCount(v int32)`

SetTeamProjectsCount sets TeamProjectsCount field to given value.

### HasTeamProjectsCount

`func (o *GetTeamByIdV2200Response) HasTeamProjectsCount() bool`

HasTeamProjectsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


