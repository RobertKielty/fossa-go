# ListSavedFilters200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**map[string]Json**](json.md) | A JSON object representing the specific filter criteria. | [optional] 
**Sort** | Pointer to **string** | Can be any of the sort options related to issue or revision. | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewListSavedFilters200ResponseInner

`func NewListSavedFilters200ResponseInner() *ListSavedFilters200ResponseInner`

NewListSavedFilters200ResponseInner instantiates a new ListSavedFilters200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSavedFilters200ResponseInnerWithDefaults

`func NewListSavedFilters200ResponseInnerWithDefaults() *ListSavedFilters200ResponseInner`

NewListSavedFilters200ResponseInnerWithDefaults instantiates a new ListSavedFilters200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSavedFilters200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSavedFilters200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSavedFilters200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListSavedFilters200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSavedFilters200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSavedFilters200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSavedFilters200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSavedFilters200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategory

`func (o *ListSavedFilters200ResponseInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListSavedFilters200ResponseInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListSavedFilters200ResponseInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListSavedFilters200ResponseInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetFilter

`func (o *ListSavedFilters200ResponseInner) GetFilter() map[string]Json`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ListSavedFilters200ResponseInner) GetFilterOk() (*map[string]Json, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ListSavedFilters200ResponseInner) SetFilter(v map[string]Json)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ListSavedFilters200ResponseInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *ListSavedFilters200ResponseInner) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ListSavedFilters200ResponseInner) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ListSavedFilters200ResponseInner) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ListSavedFilters200ResponseInner) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroup

`func (o *ListSavedFilters200ResponseInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListSavedFilters200ResponseInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListSavedFilters200ResponseInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListSavedFilters200ResponseInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


