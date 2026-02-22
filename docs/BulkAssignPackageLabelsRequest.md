# BulkAssignPackageLabelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocators** | **[]string** | Array of revision locators to assign the label to. Must include the version (e.g., &#39;npm+lodash$4.17.21&#39;). | 
**LabelId** | **int32** | The ID of the label to assign to all packages. | 
**Scope** | **string** | The scope of the package label assignment. | 
**ScopeId** | Pointer to **string** | The ID of the scope to assign labels to. Required if scope is &#39;project&#39; or &#39;revision&#39;. | [optional] 
**ShouldUseSpecificVersion** | Pointer to **bool** | If true, labels will apply to the specific versions provided in the packageLocators. If false, labels will apply to all versions of the packages (ignoring the version part of the locator). Defaults to true.  | [optional] [default to true]

## Methods

### NewBulkAssignPackageLabelsRequest

`func NewBulkAssignPackageLabelsRequest(packageLocators []string, labelId int32, scope string, ) *BulkAssignPackageLabelsRequest`

NewBulkAssignPackageLabelsRequest instantiates a new BulkAssignPackageLabelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkAssignPackageLabelsRequestWithDefaults

`func NewBulkAssignPackageLabelsRequestWithDefaults() *BulkAssignPackageLabelsRequest`

NewBulkAssignPackageLabelsRequestWithDefaults instantiates a new BulkAssignPackageLabelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocators

`func (o *BulkAssignPackageLabelsRequest) GetPackageLocators() []string`

GetPackageLocators returns the PackageLocators field if non-nil, zero value otherwise.

### GetPackageLocatorsOk

`func (o *BulkAssignPackageLabelsRequest) GetPackageLocatorsOk() (*[]string, bool)`

GetPackageLocatorsOk returns a tuple with the PackageLocators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocators

`func (o *BulkAssignPackageLabelsRequest) SetPackageLocators(v []string)`

SetPackageLocators sets PackageLocators field to given value.


### GetLabelId

`func (o *BulkAssignPackageLabelsRequest) GetLabelId() int32`

GetLabelId returns the LabelId field if non-nil, zero value otherwise.

### GetLabelIdOk

`func (o *BulkAssignPackageLabelsRequest) GetLabelIdOk() (*int32, bool)`

GetLabelIdOk returns a tuple with the LabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelId

`func (o *BulkAssignPackageLabelsRequest) SetLabelId(v int32)`

SetLabelId sets LabelId field to given value.


### GetScope

`func (o *BulkAssignPackageLabelsRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BulkAssignPackageLabelsRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BulkAssignPackageLabelsRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *BulkAssignPackageLabelsRequest) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *BulkAssignPackageLabelsRequest) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *BulkAssignPackageLabelsRequest) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *BulkAssignPackageLabelsRequest) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetShouldUseSpecificVersion

`func (o *BulkAssignPackageLabelsRequest) GetShouldUseSpecificVersion() bool`

GetShouldUseSpecificVersion returns the ShouldUseSpecificVersion field if non-nil, zero value otherwise.

### GetShouldUseSpecificVersionOk

`func (o *BulkAssignPackageLabelsRequest) GetShouldUseSpecificVersionOk() (*bool, bool)`

GetShouldUseSpecificVersionOk returns a tuple with the ShouldUseSpecificVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldUseSpecificVersion

`func (o *BulkAssignPackageLabelsRequest) SetShouldUseSpecificVersion(v bool)`

SetShouldUseSpecificVersion sets ShouldUseSpecificVersion field to given value.

### HasShouldUseSpecificVersion

`func (o *BulkAssignPackageLabelsRequest) HasShouldUseSpecificVersion() bool`

HasShouldUseSpecificVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


