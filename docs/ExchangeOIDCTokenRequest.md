# ExchangeOIDCTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderId** | **int32** | The ID of the OIDC Provider that issued the token | 
**Username** | **string** | The username of the user that we want to log in as | 
**Token** | **string** | The JWT that was issued by the OIDC Provider | 
**ExpiresIn** | Pointer to **int32** | The desired validity duration of the generated FOSSA token, in seconds. Min 15 minutes (900s), Max 12 hours (43200s), Default 1 hour (3600s). | [optional] [default to 3600]
**IsPushOnly** | Pointer to **bool** | Whether the FOSSA token should be push-only. Defaults to true. | [optional] [default to true]

## Methods

### NewExchangeOIDCTokenRequest

`func NewExchangeOIDCTokenRequest(providerId int32, username string, token string, ) *ExchangeOIDCTokenRequest`

NewExchangeOIDCTokenRequest instantiates a new ExchangeOIDCTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeOIDCTokenRequestWithDefaults

`func NewExchangeOIDCTokenRequestWithDefaults() *ExchangeOIDCTokenRequest`

NewExchangeOIDCTokenRequestWithDefaults instantiates a new ExchangeOIDCTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderId

`func (o *ExchangeOIDCTokenRequest) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ExchangeOIDCTokenRequest) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ExchangeOIDCTokenRequest) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetUsername

`func (o *ExchangeOIDCTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ExchangeOIDCTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ExchangeOIDCTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetToken

`func (o *ExchangeOIDCTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ExchangeOIDCTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ExchangeOIDCTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiresIn

`func (o *ExchangeOIDCTokenRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ExchangeOIDCTokenRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ExchangeOIDCTokenRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ExchangeOIDCTokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetIsPushOnly

`func (o *ExchangeOIDCTokenRequest) GetIsPushOnly() bool`

GetIsPushOnly returns the IsPushOnly field if non-nil, zero value otherwise.

### GetIsPushOnlyOk

`func (o *ExchangeOIDCTokenRequest) GetIsPushOnlyOk() (*bool, bool)`

GetIsPushOnlyOk returns a tuple with the IsPushOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPushOnly

`func (o *ExchangeOIDCTokenRequest) SetIsPushOnly(v bool)`

SetIsPushOnly sets IsPushOnly field to given value.

### HasIsPushOnly

`func (o *ExchangeOIDCTokenRequest) HasIsPushOnly() bool`

HasIsPushOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


