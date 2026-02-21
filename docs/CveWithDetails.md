# CveWithDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | **string** | CVE identifier (e.g., CVE-2021-44228) | 
**Description** | Pointer to **string** | Detailed description of the vulnerability | [optional] 

## Methods

### NewCveWithDetails

`func NewCveWithDetails(cve string, ) *CveWithDetails`

NewCveWithDetails instantiates a new CveWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCveWithDetailsWithDefaults

`func NewCveWithDetailsWithDefaults() *CveWithDetails`

NewCveWithDetailsWithDefaults instantiates a new CveWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCve

`func (o *CveWithDetails) GetCve() string`

GetCve returns the Cve field if non-nil, zero value otherwise.

### GetCveOk

`func (o *CveWithDetails) GetCveOk() (*string, bool)`

GetCveOk returns a tuple with the Cve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCve

`func (o *CveWithDetails) SetCve(v string)`

SetCve sets Cve field to given value.


### GetDescription

`func (o *CveWithDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CveWithDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CveWithDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CveWithDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


