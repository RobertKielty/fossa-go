# UpdateSavedFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**CreateSavedFilterRequestFilter**](CreateSavedFilterRequestFilter.md) |  | [optional] 
**Sort** | Pointer to **string** | Can be any of the sort options related to issue or revision. | [optional] 
**Group** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSavedFilterRequest

`func NewUpdateSavedFilterRequest() *UpdateSavedFilterRequest`

NewUpdateSavedFilterRequest instantiates a new UpdateSavedFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSavedFilterRequestWithDefaults

`func NewUpdateSavedFilterRequestWithDefaults() *UpdateSavedFilterRequest`

NewUpdateSavedFilterRequestWithDefaults instantiates a new UpdateSavedFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSavedFilterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSavedFilterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSavedFilterRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSavedFilterRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilter

`func (o *UpdateSavedFilterRequest) GetFilter() CreateSavedFilterRequestFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UpdateSavedFilterRequest) GetFilterOk() (*CreateSavedFilterRequestFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UpdateSavedFilterRequest) SetFilter(v CreateSavedFilterRequestFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *UpdateSavedFilterRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *UpdateSavedFilterRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *UpdateSavedFilterRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *UpdateSavedFilterRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *UpdateSavedFilterRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateSavedFilterRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateSavedFilterRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateSavedFilterRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateSavedFilterRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


