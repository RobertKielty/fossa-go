# CreateSavedFilterRequestFilterOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Identification** | Pointer to **string** |  | [optional] 
**PackageManagers** | Pointer to **[]string** |  | [optional] 
**ProjectLabels** | Pointer to **[]string** |  | [optional] 
**Search** | Pointer to **string** |  | [optional] 
**Ticketed** | Pointer to **string** |  | [optional] 
**RevisionIds** | Pointer to **[]string** |  | [optional] 
**ContainerLayers** | Pointer to **[]string** |  | [optional] 
**FoundAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateSavedFilterRequestFilterOneOf

`func NewCreateSavedFilterRequestFilterOneOf() *CreateSavedFilterRequestFilterOneOf`

NewCreateSavedFilterRequestFilterOneOf instantiates a new CreateSavedFilterRequestFilterOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedFilterRequestFilterOneOfWithDefaults

`func NewCreateSavedFilterRequestFilterOneOfWithDefaults() *CreateSavedFilterRequestFilterOneOf`

NewCreateSavedFilterRequestFilterOneOfWithDefaults instantiates a new CreateSavedFilterRequestFilterOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateSavedFilterRequestFilterOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSavedFilterRequestFilterOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSavedFilterRequestFilterOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIdentification

`func (o *CreateSavedFilterRequestFilterOneOf) GetIdentification() string`

GetIdentification returns the Identification field if non-nil, zero value otherwise.

### GetIdentificationOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetIdentificationOk() (*string, bool)`

GetIdentificationOk returns a tuple with the Identification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentification

`func (o *CreateSavedFilterRequestFilterOneOf) SetIdentification(v string)`

SetIdentification sets Identification field to given value.

### HasIdentification

`func (o *CreateSavedFilterRequestFilterOneOf) HasIdentification() bool`

HasIdentification returns a boolean if a field has been set.

### GetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf) GetPackageManagers() []string`

GetPackageManagers returns the PackageManagers field if non-nil, zero value otherwise.

### GetPackageManagersOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetPackageManagersOk() (*[]string, bool)`

GetPackageManagersOk returns a tuple with the PackageManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf) SetPackageManagers(v []string)`

SetPackageManagers sets PackageManagers field to given value.

### HasPackageManagers

`func (o *CreateSavedFilterRequestFilterOneOf) HasPackageManagers() bool`

HasPackageManagers returns a boolean if a field has been set.

### GetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf) GetProjectLabels() []string`

GetProjectLabels returns the ProjectLabels field if non-nil, zero value otherwise.

### GetProjectLabelsOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetProjectLabelsOk() (*[]string, bool)`

GetProjectLabelsOk returns a tuple with the ProjectLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf) SetProjectLabels(v []string)`

SetProjectLabels sets ProjectLabels field to given value.

### HasProjectLabels

`func (o *CreateSavedFilterRequestFilterOneOf) HasProjectLabels() bool`

HasProjectLabels returns a boolean if a field has been set.

### GetSearch

`func (o *CreateSavedFilterRequestFilterOneOf) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *CreateSavedFilterRequestFilterOneOf) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *CreateSavedFilterRequestFilterOneOf) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf) GetTicketed() string`

GetTicketed returns the Ticketed field if non-nil, zero value otherwise.

### GetTicketedOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetTicketedOk() (*string, bool)`

GetTicketedOk returns a tuple with the Ticketed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketed

`func (o *CreateSavedFilterRequestFilterOneOf) SetTicketed(v string)`

SetTicketed sets Ticketed field to given value.

### HasTicketed

`func (o *CreateSavedFilterRequestFilterOneOf) HasTicketed() bool`

HasTicketed returns a boolean if a field has been set.

### GetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf) GetRevisionIds() []string`

GetRevisionIds returns the RevisionIds field if non-nil, zero value otherwise.

### GetRevisionIdsOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetRevisionIdsOk() (*[]string, bool)`

GetRevisionIdsOk returns a tuple with the RevisionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf) SetRevisionIds(v []string)`

SetRevisionIds sets RevisionIds field to given value.

### HasRevisionIds

`func (o *CreateSavedFilterRequestFilterOneOf) HasRevisionIds() bool`

HasRevisionIds returns a boolean if a field has been set.

### GetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf) GetContainerLayers() []string`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetContainerLayersOk() (*[]string, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf) SetContainerLayers(v []string)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *CreateSavedFilterRequestFilterOneOf) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf) GetFoundAfter() time.Time`

GetFoundAfter returns the FoundAfter field if non-nil, zero value otherwise.

### GetFoundAfterOk

`func (o *CreateSavedFilterRequestFilterOneOf) GetFoundAfterOk() (*time.Time, bool)`

GetFoundAfterOk returns a tuple with the FoundAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf) SetFoundAfter(v time.Time)`

SetFoundAfter sets FoundAfter field to given value.

### HasFoundAfter

`func (o *CreateSavedFilterRequestFilterOneOf) HasFoundAfter() bool`

HasFoundAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


