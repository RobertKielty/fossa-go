# GetReleaseGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**ReleaseGroups** | Pointer to [**[]GetReleaseGroups200ResponseReleaseGroupsInner**](GetReleaseGroups200ResponseReleaseGroupsInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroups200Response

`func NewGetReleaseGroups200Response() *GetReleaseGroups200Response`

NewGetReleaseGroups200Response instantiates a new GetReleaseGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroups200ResponseWithDefaults

`func NewGetReleaseGroups200ResponseWithDefaults() *GetReleaseGroups200Response`

NewGetReleaseGroups200ResponseWithDefaults instantiates a new GetReleaseGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetReleaseGroups200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetReleaseGroups200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetReleaseGroups200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetReleaseGroups200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetReleaseGroups

`func (o *GetReleaseGroups200Response) GetReleaseGroups() []GetReleaseGroups200ResponseReleaseGroupsInner`

GetReleaseGroups returns the ReleaseGroups field if non-nil, zero value otherwise.

### GetReleaseGroupsOk

`func (o *GetReleaseGroups200Response) GetReleaseGroupsOk() (*[]GetReleaseGroups200ResponseReleaseGroupsInner, bool)`

GetReleaseGroupsOk returns a tuple with the ReleaseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroups

`func (o *GetReleaseGroups200Response) SetReleaseGroups(v []GetReleaseGroups200ResponseReleaseGroupsInner)`

SetReleaseGroups sets ReleaseGroups field to given value.

### HasReleaseGroups

`func (o *GetReleaseGroups200Response) HasReleaseGroups() bool`

HasReleaseGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


