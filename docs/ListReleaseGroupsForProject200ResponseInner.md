# ListReleaseGroupsForProject200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseGroupId** | Pointer to **float32** |  | [optional] 
**ReleaseGroupTitle** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the release group was created | [optional] 
**Revisions** | Pointer to [**[]ListReleaseGroupsForProject200ResponseInnerRevisionsInner**](ListReleaseGroupsForProject200ResponseInnerRevisionsInner.md) |  | [optional] 

## Methods

### NewListReleaseGroupsForProject200ResponseInner

`func NewListReleaseGroupsForProject200ResponseInner() *ListReleaseGroupsForProject200ResponseInner`

NewListReleaseGroupsForProject200ResponseInner instantiates a new ListReleaseGroupsForProject200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReleaseGroupsForProject200ResponseInnerWithDefaults

`func NewListReleaseGroupsForProject200ResponseInnerWithDefaults() *ListReleaseGroupsForProject200ResponseInner`

NewListReleaseGroupsForProject200ResponseInnerWithDefaults instantiates a new ListReleaseGroupsForProject200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseGroupId

`func (o *ListReleaseGroupsForProject200ResponseInner) GetReleaseGroupId() float32`

GetReleaseGroupId returns the ReleaseGroupId field if non-nil, zero value otherwise.

### GetReleaseGroupIdOk

`func (o *ListReleaseGroupsForProject200ResponseInner) GetReleaseGroupIdOk() (*float32, bool)`

GetReleaseGroupIdOk returns a tuple with the ReleaseGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroupId

`func (o *ListReleaseGroupsForProject200ResponseInner) SetReleaseGroupId(v float32)`

SetReleaseGroupId sets ReleaseGroupId field to given value.

### HasReleaseGroupId

`func (o *ListReleaseGroupsForProject200ResponseInner) HasReleaseGroupId() bool`

HasReleaseGroupId returns a boolean if a field has been set.

### GetReleaseGroupTitle

`func (o *ListReleaseGroupsForProject200ResponseInner) GetReleaseGroupTitle() string`

GetReleaseGroupTitle returns the ReleaseGroupTitle field if non-nil, zero value otherwise.

### GetReleaseGroupTitleOk

`func (o *ListReleaseGroupsForProject200ResponseInner) GetReleaseGroupTitleOk() (*string, bool)`

GetReleaseGroupTitleOk returns a tuple with the ReleaseGroupTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroupTitle

`func (o *ListReleaseGroupsForProject200ResponseInner) SetReleaseGroupTitle(v string)`

SetReleaseGroupTitle sets ReleaseGroupTitle field to given value.

### HasReleaseGroupTitle

`func (o *ListReleaseGroupsForProject200ResponseInner) HasReleaseGroupTitle() bool`

HasReleaseGroupTitle returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListReleaseGroupsForProject200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListReleaseGroupsForProject200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListReleaseGroupsForProject200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListReleaseGroupsForProject200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRevisions

`func (o *ListReleaseGroupsForProject200ResponseInner) GetRevisions() []ListReleaseGroupsForProject200ResponseInnerRevisionsInner`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *ListReleaseGroupsForProject200ResponseInner) GetRevisionsOk() (*[]ListReleaseGroupsForProject200ResponseInnerRevisionsInner, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *ListReleaseGroupsForProject200ResponseInner) SetRevisions(v []ListReleaseGroupsForProject200ResponseInnerRevisionsInner)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *ListReleaseGroupsForProject200ResponseInner) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


