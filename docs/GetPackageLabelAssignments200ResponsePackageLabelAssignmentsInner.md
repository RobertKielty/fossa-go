# GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner

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
**PackageLabel** | [**GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel**](GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel.md) |  | 

## Methods

### NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner

`func NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner(id int32, createdAt time.Time, updatedAt time.Time, organizationId int32, labelId int32, packageId string, scope string, scopeId string, packageLabel GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel, ) *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner`

NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner instantiates a new GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults

`func NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults() *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner`

NewGetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerWithDefaults instantiates a new GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganizationId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetLabelId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.


### GetPackageId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageVersion

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetScope

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetPackageLabel

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageLabel() GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel`

GetPackageLabel returns the PackageLabel field if non-nil, zero value otherwise.

### GetPackageLabelOk

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) GetPackageLabelOk() (*GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel, bool)`

GetPackageLabelOk returns a tuple with the PackageLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabel

`func (o *GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInner) SetPackageLabel(v GetPackageLabelAssignments200ResponsePackageLabelAssignmentsInnerAllOfPackageLabel)`

SetPackageLabel sets PackageLabel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


