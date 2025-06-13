# GetAllRoles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The role&#39;s unique identifier | [optional] 
**OrganizationId** | Pointer to **int32** | The organization&#39;s unique identifier | [optional] 
**IsCustom** | Pointer to **bool** | Whether the role is custom | [optional] 
**Scope** | Pointer to **string** | The role&#39;s scope | [optional] 
**Name** | Pointer to **string** | The role&#39;s name | [optional] 
**Description** | Pointer to **string** | The role&#39;s description | [optional] 
**CreatedAt** | Pointer to **time.Time** | The role&#39;s creation date | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The role&#39;s last update date | [optional] 
**Permissions** | Pointer to [**[]GetAllRoles200ResponseInnerPermissionsInner**](GetAllRoles200ResponseInnerPermissionsInner.md) |  | [optional] 

## Methods

### NewGetAllRoles200ResponseInner

`func NewGetAllRoles200ResponseInner() *GetAllRoles200ResponseInner`

NewGetAllRoles200ResponseInner instantiates a new GetAllRoles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllRoles200ResponseInnerWithDefaults

`func NewGetAllRoles200ResponseInnerWithDefaults() *GetAllRoles200ResponseInner`

NewGetAllRoles200ResponseInnerWithDefaults instantiates a new GetAllRoles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllRoles200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllRoles200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllRoles200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAllRoles200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetAllRoles200ResponseInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetAllRoles200ResponseInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetAllRoles200ResponseInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetAllRoles200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetIsCustom

`func (o *GetAllRoles200ResponseInner) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *GetAllRoles200ResponseInner) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *GetAllRoles200ResponseInner) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.

### HasIsCustom

`func (o *GetAllRoles200ResponseInner) HasIsCustom() bool`

HasIsCustom returns a boolean if a field has been set.

### GetScope

`func (o *GetAllRoles200ResponseInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetAllRoles200ResponseInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetAllRoles200ResponseInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetAllRoles200ResponseInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetName

`func (o *GetAllRoles200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllRoles200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllRoles200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllRoles200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetAllRoles200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAllRoles200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAllRoles200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAllRoles200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAllRoles200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAllRoles200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAllRoles200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAllRoles200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetAllRoles200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAllRoles200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAllRoles200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetAllRoles200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPermissions

`func (o *GetAllRoles200ResponseInner) GetPermissions() []GetAllRoles200ResponseInnerPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetAllRoles200ResponseInner) GetPermissionsOk() (*[]GetAllRoles200ResponseInnerPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetAllRoles200ResponseInner) SetPermissions(v []GetAllRoles200ResponseInnerPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetAllRoles200ResponseInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


