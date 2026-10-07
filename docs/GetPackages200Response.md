# GetPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetPackages200ResponseDataInner**](GetPackages200ResponseDataInner.md) | The page of packages matching the supplied filters. | 
**Count** | **int32** | The total number of packages matching the supplied filters across all pages. | 

## Methods

### NewGetPackages200Response

`func NewGetPackages200Response(data []GetPackages200ResponseDataInner, count int32, ) *GetPackages200Response`

NewGetPackages200Response instantiates a new GetPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPackages200ResponseWithDefaults

`func NewGetPackages200ResponseWithDefaults() *GetPackages200Response`

NewGetPackages200ResponseWithDefaults instantiates a new GetPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetPackages200Response) GetData() []GetPackages200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPackages200Response) GetDataOk() (*[]GetPackages200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPackages200Response) SetData(v []GetPackages200ResponseDataInner)`

SetData sets Data field to given value.


### GetCount

`func (o *GetPackages200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetPackages200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetPackages200Response) SetCount(v int32)`

SetCount sets Count field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


