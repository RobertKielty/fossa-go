# GetRevisionDependenciesPost200ResponseInnerDependencyLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locator** | Pointer to **string** |  | [optional] 
**UnresolvedLocators** | Pointer to **[]string** | List of unresolved locators for this dependency | [optional] 
**Depth** | Pointer to **int32** | Depth of this dependency in the dependency tree | [optional] 
**Type** | Pointer to **string** | Type of dependency (deprecated, always \&quot;MEDIATED\&quot;) | [optional] 
**Tags** | Pointer to **[]string** | Tags for this dependency (deprecated, always empty) | [optional] 
**IsSubmodule** | Pointer to **bool** | Whether this is a submodule (deprecated, always false) | [optional] 
**Root** | Pointer to **string** | The root revision locator | [optional] 
**OriginPaths** | Pointer to **[]string** | Paths from the root to this dependency | [optional] 
**Manual** | Pointer to **bool** | Whether this dependency was manually added | [optional] 

## Methods

### NewGetRevisionDependenciesPost200ResponseInnerDependencyLock

`func NewGetRevisionDependenciesPost200ResponseInnerDependencyLock() *GetRevisionDependenciesPost200ResponseInnerDependencyLock`

NewGetRevisionDependenciesPost200ResponseInnerDependencyLock instantiates a new GetRevisionDependenciesPost200ResponseInnerDependencyLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionDependenciesPost200ResponseInnerDependencyLockWithDefaults

`func NewGetRevisionDependenciesPost200ResponseInnerDependencyLockWithDefaults() *GetRevisionDependenciesPost200ResponseInnerDependencyLock`

NewGetRevisionDependenciesPost200ResponseInnerDependencyLockWithDefaults instantiates a new GetRevisionDependenciesPost200ResponseInnerDependencyLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocator

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetUnresolvedLocators() []string`

GetUnresolvedLocators returns the UnresolvedLocators field if non-nil, zero value otherwise.

### GetUnresolvedLocatorsOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetUnresolvedLocatorsOk() (*[]string, bool)`

GetUnresolvedLocatorsOk returns a tuple with the UnresolvedLocators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetUnresolvedLocators(v []string)`

SetUnresolvedLocators sets UnresolvedLocators field to given value.

### HasUnresolvedLocators

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasUnresolvedLocators() bool`

HasUnresolvedLocators returns a boolean if a field has been set.

### GetDepth

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetType

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetIsSubmodule() bool`

GetIsSubmodule returns the IsSubmodule field if non-nil, zero value otherwise.

### GetIsSubmoduleOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetIsSubmoduleOk() (*bool, bool)`

GetIsSubmoduleOk returns a tuple with the IsSubmodule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetIsSubmodule(v bool)`

SetIsSubmodule sets IsSubmodule field to given value.

### HasIsSubmodule

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasIsSubmodule() bool`

HasIsSubmodule returns a boolean if a field has been set.

### GetRoot

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetRoot() string`

GetRoot returns the Root field if non-nil, zero value otherwise.

### GetRootOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetRootOk() (*string, bool)`

GetRootOk returns a tuple with the Root field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoot

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetRoot(v string)`

SetRoot sets Root field to given value.

### HasRoot

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasRoot() bool`

HasRoot returns a boolean if a field has been set.

### GetOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetOriginPaths() []string`

GetOriginPaths returns the OriginPaths field if non-nil, zero value otherwise.

### GetOriginPathsOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetOriginPathsOk() (*[]string, bool)`

GetOriginPathsOk returns a tuple with the OriginPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetOriginPaths(v []string)`

SetOriginPaths sets OriginPaths field to given value.

### HasOriginPaths

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasOriginPaths() bool`

HasOriginPaths returns a boolean if a field has been set.

### GetManual

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetManual() bool`

GetManual returns the Manual field if non-nil, zero value otherwise.

### GetManualOk

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) GetManualOk() (*bool, bool)`

GetManualOk returns a tuple with the Manual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManual

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) SetManual(v bool)`

SetManual sets Manual field to given value.

### HasManual

`func (o *GetRevisionDependenciesPost200ResponseInnerDependencyLock) HasManual() bool`

HasManual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


