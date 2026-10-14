# GetProjectDependencies200ResponseDependenciesInnerLabelsInner

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
**Name** | **string** | The name of the label that was assigned to the package. | 

## Methods

### NewGetProjectDependencies200ResponseDependenciesInnerLabelsInner

`func NewGetProjectDependencies200ResponseDependenciesInnerLabelsInner(id int32, createdAt time.Time, updatedAt time.Time, organizationId int32, labelId int32, packageId string, scope string, scopeId string, name string, ) *GetProjectDependencies200ResponseDependenciesInnerLabelsInner`

NewGetProjectDependencies200ResponseDependenciesInnerLabelsInner instantiates a new GetProjectDependencies200ResponseDependenciesInnerLabelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectDependencies200ResponseDependenciesInnerLabelsInnerWithDefaults

`func NewGetProjectDependencies200ResponseDependenciesInnerLabelsInnerWithDefaults() *GetProjectDependencies200ResponseDependenciesInnerLabelsInner`

NewGetProjectDependencies200ResponseDependenciesInnerLabelsInnerWithDefaults instantiates a new GetProjectDependencies200ResponseDependenciesInnerLabelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganizationId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetLabelId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.


### GetPackageId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageVersion

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetScope

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetName

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetProjectDependencies200ResponseDependenciesInnerLabelsInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


