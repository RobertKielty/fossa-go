# GetSnippets200ResponseResultsInnerLabelsInner

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

### NewGetSnippets200ResponseResultsInnerLabelsInner

`func NewGetSnippets200ResponseResultsInnerLabelsInner(id int32, createdAt time.Time, updatedAt time.Time, organizationId int32, labelId int32, packageId string, scope string, scopeId string, name string, ) *GetSnippets200ResponseResultsInnerLabelsInner`

NewGetSnippets200ResponseResultsInnerLabelsInner instantiates a new GetSnippets200ResponseResultsInnerLabelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippets200ResponseResultsInnerLabelsInnerWithDefaults

`func NewGetSnippets200ResponseResultsInnerLabelsInnerWithDefaults() *GetSnippets200ResponseResultsInnerLabelsInner`

NewGetSnippets200ResponseResultsInnerLabelsInnerWithDefaults instantiates a new GetSnippets200ResponseResultsInnerLabelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOrganizationId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetLabelId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.


### GetPackageId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageVersion

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetScope

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetName

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSnippets200ResponseResultsInnerLabelsInner) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


