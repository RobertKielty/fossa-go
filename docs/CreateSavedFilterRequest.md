# CreateSavedFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**CreateSavedFilterRequestFilter**](CreateSavedFilterRequestFilter.md) |  | [optional] 
**Sort** | Pointer to **string** | Can be any of the sort options related to issue or revision. | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSavedFilterRequest

`func NewCreateSavedFilterRequest() *CreateSavedFilterRequest`

NewCreateSavedFilterRequest instantiates a new CreateSavedFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedFilterRequestWithDefaults

`func NewCreateSavedFilterRequestWithDefaults() *CreateSavedFilterRequest`

NewCreateSavedFilterRequestWithDefaults instantiates a new CreateSavedFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CreateSavedFilterRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateSavedFilterRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateSavedFilterRequest) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateSavedFilterRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetName

`func (o *CreateSavedFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSavedFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSavedFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateSavedFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilter

`func (o *CreateSavedFilterRequest) GetFilter() CreateSavedFilterRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CreateSavedFilterRequest) GetFilterOk() (*CreateSavedFilterRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CreateSavedFilterRequest) SetFilter(v CreateSavedFilterRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CreateSavedFilterRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *CreateSavedFilterRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CreateSavedFilterRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CreateSavedFilterRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CreateSavedFilterRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroup

`func (o *CreateSavedFilterRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CreateSavedFilterRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CreateSavedFilterRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CreateSavedFilterRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


