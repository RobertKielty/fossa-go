# UpdateProjectLabelsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]int32** | List of Label IDs that should be associated with this Project | [optional] 

## Methods

### NewUpdateProjectLabelsRequest

`func NewUpdateProjectLabelsRequest() *UpdateProjectLabelsRequest`

NewUpdateProjectLabelsRequest instantiates a new UpdateProjectLabelsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectLabelsRequestWithDefaults

`func NewUpdateProjectLabelsRequestWithDefaults() *UpdateProjectLabelsRequest`

NewUpdateProjectLabelsRequestWithDefaults instantiates a new UpdateProjectLabelsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *UpdateProjectLabelsRequest) GetLabels() []int32`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateProjectLabelsRequest) GetLabelsOk() (*[]int32, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateProjectLabelsRequest) SetLabels(v []int32)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateProjectLabelsRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


