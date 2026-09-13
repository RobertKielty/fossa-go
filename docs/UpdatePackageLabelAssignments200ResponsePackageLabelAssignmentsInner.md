# UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the package label assignment. | 
**CreatedAt** | **time.Time** | The date and time the package label assignment was created. | 
**UpdatedAt** | **time.Time** | The date and time the package label assignment was last updated. | 
**OrganizationId** | **int32** | The ID of the organization that owns the package label assignment. | 
**LabelId** | **int32** | The ID of the label that was assigned to the package. | 
**PackageId** | **string** | The ID of the package that the label was assigned to. | 
**PackageVersion** | Pointer to **string** | The version of the package that the label was assigned to or null if the label was assigned to all versions. | [optional] 
**Scope** | **string** | The scope of the package label assignment. | 
**ScopeId** | **string** | The ID of the scope that the label was assigned to or null if the label was assigned to all scopes. | 

## Methods

### NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner

`func NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner(id int32, createdAt time.Time, updatedAt time.Time, organizationId int32, labelId int32, packageId string, scope string, scopeId string, ) *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner`

NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner instantiates a new UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults

`func NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults() *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner`

NewUpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults instantiates a new UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganizationId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetLabelId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.


### GetPackageId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageVersion

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetScope

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


