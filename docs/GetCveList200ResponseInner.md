# GetCveList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | **string** | CVE identifier (e.g., CVE-2021-44228) | 
**Description** | Pointer to **string** | Detailed description of the vulnerability | [optional] 

## Methods

### NewGetCveList200ResponseInner

`func NewGetCveList200ResponseInner(cve string, ) *GetCveList200ResponseInner`

NewGetCveList200ResponseInner instantiates a new GetCveList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCveList200ResponseInnerWithDefaults

`func NewGetCveList200ResponseInnerWithDefaults() *GetCveList200ResponseInner`

NewGetCveList200ResponseInnerWithDefaults instantiates a new GetCveList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *GetCveList200ResponseInner) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *GetCveList200ResponseInner) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *GetCveList200ResponseInner) SetCve(v string)`

SetCve sets Cve field to given value.


### GetDescription

`func (o *GetCveList200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCveList200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCveList200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCveList200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


