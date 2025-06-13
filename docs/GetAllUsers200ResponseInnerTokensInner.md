# GetAllUsers200ResponseInnerTokensInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The token&#39;s unique identifier | [optional] 
**Name** | Pointer to **string** | The token&#39;s name | [optional] 
**IsDisabled** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date the token was last updated | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date the token was created | [optional] 
**Meta** | Pointer to [**GetAllUsers200ResponseInnerTokensInnerMeta**](GetAllUsers200ResponseInnerTokensInnerMeta.md) |  | [optional] 

## Methods

### NewGetAllUsers200ResponseInnerTokensInner

`func NewGetAllUsers200ResponseInnerTokensInner() *GetAllUsers200ResponseInnerTokensInner`

NewGetAllUsers200ResponseInnerTokensInner instantiates a new GetAllUsers200ResponseInnerTokensInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllUsers200ResponseInnerTokensInnerWithDefaults

`func NewGetAllUsers200ResponseInnerTokensInnerWithDefaults() *GetAllUsers200ResponseInnerTokensInner`

NewGetAllUsers200ResponseInnerTokensInnerWithDefaults instantiates a new GetAllUsers200ResponseInnerTokensInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllUsers200ResponseInnerTokensInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllUsers200ResponseInnerTokensInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAllUsers200ResponseInnerTokensInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetAllUsers200ResponseInnerTokensInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllUsers200ResponseInnerTokensInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllUsers200ResponseInnerTokensInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDisabled

`func (o *GetAllUsers200ResponseInnerTokensInner) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *GetAllUsers200ResponseInnerTokensInner) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *GetAllUsers200ResponseInnerTokensInner) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetAllUsers200ResponseInnerTokensInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMeta

`func (o *GetAllUsers200ResponseInnerTokensInner) GetMeta() GetAllUsers200ResponseInnerTokensInnerMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllUsers200ResponseInnerTokensInner) GetMetaOk() (*GetAllUsers200ResponseInnerTokensInnerMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllUsers200ResponseInnerTokensInner) SetMeta(v GetAllUsers200ResponseInnerTokensInnerMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAllUsers200ResponseInnerTokensInner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


