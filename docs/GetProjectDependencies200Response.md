# GetProjectDependencies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInner**](GetProjectDependencies200ResponseDependenciesInner.md) |  | [optional] 
**Total** | Pointer to **int32** | The total number of dependencies matching the filters, across all pages. | [optional] 

## Methods

### NewGetProjectDependencies200Response

`func NewGetProjectDependencies200Response() *GetProjectDependencies200Response`

NewGetProjectDependencies200Response instantiates a new GetProjectDependencies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectDependencies200ResponseWithDefaults

`func NewGetProjectDependencies200ResponseWithDefaults() *GetProjectDependencies200Response`

NewGetProjectDependencies200ResponseWithDefaults instantiates a new GetProjectDependencies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencies

`func (o *GetProjectDependencies200Response) GetDependencies() []GetProjectDependencies200ResponseDependenciesInner`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *GetProjectDependencies200Response) GetDependenciesOk() (*[]GetProjectDependencies200ResponseDependenciesInner, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *GetProjectDependencies200Response) SetDependencies(v []GetProjectDependencies200ResponseDependenciesInner)`

SetDependencies sets Dependencies field to given value.

### HasDependencies

`func (o *GetProjectDependencies200Response) HasDependencies() bool`

HasDependencies returns a boolean if a field has been set.

### GetTotal

`func (o *GetProjectDependencies200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetProjectDependencies200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetProjectDependencies200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetProjectDependencies200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


