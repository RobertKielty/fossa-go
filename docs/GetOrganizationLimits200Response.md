# GetOrganizationLimits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Usage** | Pointer to **int32** | The current usage of the resource. | [optional] 
**Max** | Pointer to **int32** | The contractual limit of the resource. | [optional] 
**Unlimited** | Pointer to **bool** | Whether or not the organization has an unlimited cap on the resource. | [optional] 

## Methods

### NewGetOrganizationLimits200Response

`func NewGetOrganizationLimits200Response() *GetOrganizationLimits200Response`

NewGetOrganizationLimits200Response instantiates a new GetOrganizationLimits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationLimits200ResponseWithDefaults

`func NewGetOrganizationLimits200ResponseWithDefaults() *GetOrganizationLimits200Response`

NewGetOrganizationLimits200ResponseWithDefaults instantiates a new GetOrganizationLimits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsage

`func (o *GetOrganizationLimits200Response) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *GetOrganizationLimits200Response) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *GetOrganizationLimits200Response) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *GetOrganizationLimits200Response) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetMax

`func (o *GetOrganizationLimits200Response) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetOrganizationLimits200Response) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetOrganizationLimits200Response) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GetOrganizationLimits200Response) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetUnlimited

`func (o *GetOrganizationLimits200Response) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *GetOrganizationLimits200Response) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *GetOrganizationLimits200Response) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.

### HasUnlimited

`func (o *GetOrganizationLimits200Response) HasUnlimited() bool`

HasUnlimited returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


