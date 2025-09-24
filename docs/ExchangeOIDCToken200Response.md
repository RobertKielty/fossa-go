# ExchangeOIDCToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | The ID of the user that the token is for | 
**ProviderId** | **int32** | The ID of the OIDC Provider | 
**Issuer** | **string** | The issuer of the token | 
**Subject** | **string** | The subject of the token | 
**Credential** | [**ExchangeOIDCToken200ResponseCredential**](ExchangeOIDCToken200ResponseCredential.md) |  | 

## Methods

### NewExchangeOIDCToken200Response

`func NewExchangeOIDCToken200Response(userId int32, providerId int32, issuer string, subject string, credential ExchangeOIDCToken200ResponseCredential, ) *ExchangeOIDCToken200Response`

NewExchangeOIDCToken200Response instantiates a new ExchangeOIDCToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeOIDCToken200ResponseWithDefaults

`func NewExchangeOIDCToken200ResponseWithDefaults() *ExchangeOIDCToken200Response`

NewExchangeOIDCToken200ResponseWithDefaults instantiates a new ExchangeOIDCToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ExchangeOIDCToken200Response) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExchangeOIDCToken200Response) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExchangeOIDCToken200Response) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetProviderId

`func (o *ExchangeOIDCToken200Response) GetProviderId() int32`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ExchangeOIDCToken200Response) GetProviderIdOk() (*int32, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ExchangeOIDCToken200Response) SetProviderId(v int32)`

SetProviderId sets ProviderId field to given value.


### GetIssuer

`func (o *ExchangeOIDCToken200Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ExchangeOIDCToken200Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ExchangeOIDCToken200Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSubject

`func (o *ExchangeOIDCToken200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExchangeOIDCToken200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExchangeOIDCToken200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetCredential

`func (o *ExchangeOIDCToken200Response) GetCredential() ExchangeOIDCToken200ResponseCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ExchangeOIDCToken200Response) GetCredentialOk() (*ExchangeOIDCToken200ResponseCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ExchangeOIDCToken200Response) SetCredential(v ExchangeOIDCToken200ResponseCredential)`

SetCredential sets Credential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


