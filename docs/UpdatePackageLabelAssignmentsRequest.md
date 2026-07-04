# UpdatePackageLabelAssignmentsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **string** | The ID of the package whose label assignments are being reconciled. | 
**Scope** | **string** | The scope of the package label assignment. | 
**ScopeId** | Pointer to **string** | The ID of the scope to reconcile assignments for. | [optional] 
**NewLabelIds** | **map[string][]int32** | The desired label assignments, keyed by package version (or &#x60;all&#x60; for assignments that apply to every version). Each value is the list of label IDs that should be assigned for that version.  | 

## Methods

### NewUpdatePackageLabelAssignmentsRequest

`func NewUpdatePackageLabelAssignmentsRequest(packageId string, scope string, newLabelIds map[string][]int32, ) *UpdatePackageLabelAssignmentsRequest`

NewUpdatePackageLabelAssignmentsRequest instantiates a new UpdatePackageLabelAssignmentsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePackageLabelAssignmentsRequestWithDefaults

`func NewUpdatePackageLabelAssignmentsRequestWithDefaults() *UpdatePackageLabelAssignmentsRequest`

NewUpdatePackageLabelAssignmentsRequestWithDefaults instantiates a new UpdatePackageLabelAssignmentsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *UpdatePackageLabelAssignmentsRequest) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *UpdatePackageLabelAssignmentsRequest) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *UpdatePackageLabelAssignmentsRequest) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetScope

`func (o *UpdatePackageLabelAssignmentsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdatePackageLabelAssignmentsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdatePackageLabelAssignmentsRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *UpdatePackageLabelAssignmentsRequest) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *UpdatePackageLabelAssignmentsRequest) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *UpdatePackageLabelAssignmentsRequest) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *UpdatePackageLabelAssignmentsRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetNewLabelIds

`func (o *UpdatePackageLabelAssignmentsRequest) GetNewLabelIds() map[string][]int32`

GetNewLabelIds returns the NewLabelIds field if non-nil, zero value otherwise.

### GetNewLabelIdsOk

`func (o *UpdatePackageLabelAssignmentsRequest) GetNewLabelIdsOk() (*map[string][]int32, bool)`

GetNewLabelIdsOk returns a tuple with the NewLabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewLabelIds

`func (o *UpdatePackageLabelAssignmentsRequest) SetNewLabelIds(v map[string][]int32)`

SetNewLabelIds sets NewLabelIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


