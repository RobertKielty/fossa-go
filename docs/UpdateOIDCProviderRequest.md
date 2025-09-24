# UpdateOIDCProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Thumbprints** | **[]string** | Array of thumbprints for certificate validation | 

## Methods

### NewUpdateOIDCProviderRequest

`func NewUpdateOIDCProviderRequest(thumbprints []string, ) *UpdateOIDCProviderRequest`

NewUpdateOIDCProviderRequest instantiates a new UpdateOIDCProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOIDCProviderRequestWithDefaults

`func NewUpdateOIDCProviderRequestWithDefaults() *UpdateOIDCProviderRequest`

NewUpdateOIDCProviderRequestWithDefaults instantiates a new UpdateOIDCProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbprints

`func (o *UpdateOIDCProviderRequest) GetThumbprints() []string`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *UpdateOIDCProviderRequest) GetThumbprintsOk() (*[]string, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *UpdateOIDCProviderRequest) SetThumbprints(v []string)`

SetThumbprints sets Thumbprints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


