# ExchangeOIDCToken200ResponseCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The FOSSA token that can be used to authenticate with the FOSSA API | 
**IsPushOnly** | **bool** | Whether the FOSSA token is push-only | 
**Expiration** | **time.Time** | When the FOSSA token expires | 

## Methods

### NewExchangeOIDCToken200ResponseCredential

`func NewExchangeOIDCToken200ResponseCredential(token string, isPushOnly bool, expiration time.Time, ) *ExchangeOIDCToken200ResponseCredential`

NewExchangeOIDCToken200ResponseCredential instantiates a new ExchangeOIDCToken200ResponseCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeOIDCToken200ResponseCredentialWithDefaults

`func NewExchangeOIDCToken200ResponseCredentialWithDefaults() *ExchangeOIDCToken200ResponseCredential`

NewExchangeOIDCToken200ResponseCredentialWithDefaults instantiates a new ExchangeOIDCToken200ResponseCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ExchangeOIDCToken200ResponseCredential) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ExchangeOIDCToken200ResponseCredential) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ExchangeOIDCToken200ResponseCredential) SetToken(v string)`

SetToken sets Token field to given value.


### GetIsPushOnly

`func (o *ExchangeOIDCToken200ResponseCredential) GetIsPushOnly() bool`

GetIsPushOnly returns the IsPushOnly field if non-nil, zero value otherwise.

### GetIsPushOnlyOk

`func (o *ExchangeOIDCToken200ResponseCredential) GetIsPushOnlyOk() (*bool, bool)`

GetIsPushOnlyOk returns a tuple with the IsPushOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPushOnly

`func (o *ExchangeOIDCToken200ResponseCredential) SetIsPushOnly(v bool)`

SetIsPushOnly sets IsPushOnly field to given value.


### GetExpiration

`func (o *ExchangeOIDCToken200ResponseCredential) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ExchangeOIDCToken200ResponseCredential) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ExchangeOIDCToken200ResponseCredential) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


