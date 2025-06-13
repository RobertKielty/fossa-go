# GetProjectDependencies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependencies** | Pointer to [**[]GetProjectDependencies200ResponseDependenciesInner**](GetProjectDependencies200ResponseDependenciesInner.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

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

### GetCount

`func (o *GetProjectDependencies200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetProjectDependencies200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetProjectDependencies200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetProjectDependencies200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


