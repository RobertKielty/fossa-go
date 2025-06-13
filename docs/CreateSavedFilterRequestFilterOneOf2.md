# CreateSavedFilterRequestFilterOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Ticketed** | Pointer to **string** |  | [optional] 
**PackageManagers** | Pointer to **[]string** |  | [optional] 
**ProjectLabels** | Pointer to **[]string** |  | [optional] 
**RevisionIds** | Pointer to **[]string** |  | [optional] 
**ContainerLayers** | Pointer to **[]string** |  | [optional] 
**FoundAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateSavedFilterRequestFilterOneOf2

`func NewCreateSavedFilterRequestFilterOneOf2() *CreateSavedFilterRequestFilterOneOf2`

NewCreateSavedFilterRequestFilterOneOf2 instantiates a new CreateSavedFilterRequestFilterOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedFilterRequestFilterOneOf2WithDefaults

`func NewCreateSavedFilterRequestFilterOneOf2WithDefaults() *CreateSavedFilterRequestFilterOneOf2`

NewCreateSavedFilterRequestFilterOneOf2WithDefaults instantiates a new CreateSavedFilterRequestFilterOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateSavedFilterRequestFilterOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSavedFilterRequestFilterOneOf2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSavedFilterRequestFilterOneOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSearch

`func (o *CreateSavedFilterRequestFilterOneOf2) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreateSavedFilterRequestFilterOneOf2) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreateSavedFilterRequestFilterOneOf2) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf2) GetTicketed() string`

GetTicketed returns the Ticketed field if non-nil, zero value otherwise.

### GetTicketedOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetTicketedOk() (*string, bool)`

GetTicketedOk returns a tuple with the Ticketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf2) SetTicketed(v string)`

SetTicketed sets Ticketed field to given value.

### HasTicketed

`func (o *CreateSavedFilterRequestFilterOneOf2) HasTicketed() bool`

HasTicketed returns a boolean if a field has been set.

### GetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf2) GetPackageManagers() []string`

GetPackageManagers returns the PackageManagers field if non-nil, zero value otherwise.

### GetPackageManagersOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetPackageManagersOk() (*[]string, bool)`

GetPackageManagersOk returns a tuple with the PackageManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf2) SetPackageManagers(v []string)`

SetPackageManagers sets PackageManagers field to given value.

### HasPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf2) HasPackageManagers() bool`

HasPackageManagers returns a boolean if a field has been set.

### GetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf2) GetProjectLabels() []string`

GetProjectLabels returns the ProjectLabels field if non-nil, zero value otherwise.

### GetProjectLabelsOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetProjectLabelsOk() (*[]string, bool)`

GetProjectLabelsOk returns a tuple with the ProjectLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf2) SetProjectLabels(v []string)`

SetProjectLabels sets ProjectLabels field to given value.

### HasProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf2) HasProjectLabels() bool`

HasProjectLabels returns a boolean if a field has been set.

### GetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf2) GetRevisionIds() []string`

GetRevisionIds returns the RevisionIds field if non-nil, zero value otherwise.

### GetRevisionIdsOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetRevisionIdsOk() (*[]string, bool)`

GetRevisionIdsOk returns a tuple with the RevisionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf2) SetRevisionIds(v []string)`

SetRevisionIds sets RevisionIds field to given value.

### HasRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf2) HasRevisionIds() bool`

HasRevisionIds returns a boolean if a field has been set.

### GetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf2) GetContainerLayers() []string`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetContainerLayersOk() (*[]string, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf2) SetContainerLayers(v []string)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf2) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf2) GetFoundAfter() time.Time`

GetFoundAfter returns the FoundAfter field if non-nil, zero value otherwise.

### GetFoundAfterOk

`func (o *CreateSavedFilterRequestFilterOneOf2) GetFoundAfterOk() (*time.Time, bool)`

GetFoundAfterOk returns a tuple with the FoundAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf2) SetFoundAfter(v time.Time)`

SetFoundAfter sets FoundAfter field to given value.

### HasFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf2) HasFoundAfter() bool`

HasFoundAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


