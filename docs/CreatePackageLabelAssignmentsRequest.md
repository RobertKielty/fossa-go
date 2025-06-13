# CreatePackageLabelAssignmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **string** | The ID of the package to assign labels to. | 
**PackageVersion** | Pointer to **string** | The version of the package to assign labels to or blank for all versions. | [optional] 
**Scope** | Pointer to **string** | The scope of the package label assignment. | [optional] 
**ScopeId** | Pointer to **string** | The ID of the scope to assign labels to. | [optional] 
**LabelIds** | Pointer to **[]int32** | The IDs of the labels to assign to the package. | [optional] 

## Methods

### NewCreatePackageLabelAssignmentsRequest

`func NewCreatePackageLabelAssignmentsRequest(packageId string, ) *CreatePackageLabelAssignmentsRequest`

NewCreatePackageLabelAssignmentsRequest instantiates a new CreatePackageLabelAssignmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePackageLabelAssignmentsRequestWithDefaults

`func NewCreatePackageLabelAssignmentsRequestWithDefaults() *CreatePackageLabelAssignmentsRequest`

NewCreatePackageLabelAssignmentsRequestWithDefaults instantiates a new CreatePackageLabelAssignmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *CreatePackageLabelAssignmentsRequest) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *CreatePackageLabelAssignmentsRequest) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *CreatePackageLabelAssignmentsRequest) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPackageVersion

`func (o *CreatePackageLabelAssignmentsRequest) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *CreatePackageLabelAssignmentsRequest) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *CreatePackageLabelAssignmentsRequest) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *CreatePackageLabelAssignmentsRequest) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetScope

`func (o *CreatePackageLabelAssignmentsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreatePackageLabelAssignmentsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreatePackageLabelAssignmentsRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreatePackageLabelAssignmentsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopeId

`func (o *CreatePackageLabelAssignmentsRequest) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreatePackageLabelAssignmentsRequest) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreatePackageLabelAssignmentsRequest) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *CreatePackageLabelAssignmentsRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetLabelIds

`func (o *CreatePackageLabelAssignmentsRequest) GetLabelIds() []int32`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *CreatePackageLabelAssignmentsRequest) GetLabelIdsOk() (*[]int32, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *CreatePackageLabelAssignmentsRequest) SetLabelIds(v []int32)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *CreatePackageLabelAssignmentsRequest) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


