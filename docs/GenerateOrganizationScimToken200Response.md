# GenerateOrganizationScimToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The generated SCIM bearer token, for the identity provider to authenticate against the &#x60;/api/scim&#x60; endpoints. FOSSA stores only a hash of it, so this response is the only time it is ever returned — it cannot be retrieved again afterward. | 

## Methods

### NewGenerateOrganizationScimToken200Response

`func NewGenerateOrganizationScimToken200Response(token string, ) *GenerateOrganizationScimToken200Response`

NewGenerateOrganizationScimToken200Response instantiates a new GenerateOrganizationScimToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOrganizationScimToken200ResponseWithDefaults

`func NewGenerateOrganizationScimToken200ResponseWithDefaults() *GenerateOrganizationScimToken200Response`

NewGenerateOrganizationScimToken200ResponseWithDefaults instantiates a new GenerateOrganizationScimToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *GenerateOrganizationScimToken200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GenerateOrganizationScimToken200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GenerateOrganizationScimToken200Response) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


