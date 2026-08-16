# GetFossabotStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** |  | 
**AppUrl** | **NullableString** |  | 
**CreditLevel** | **NullableString** |  | 

## Methods

### NewGetFossabotStatus200Response

`func NewGetFossabotStatus200Response(connected bool, appUrl NullableString, creditLevel NullableString, ) *GetFossabotStatus200Response`

NewGetFossabotStatus200Response instantiates a new GetFossabotStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFossabotStatus200ResponseWithDefaults

`func NewGetFossabotStatus200ResponseWithDefaults() *GetFossabotStatus200Response`

NewGetFossabotStatus200ResponseWithDefaults instantiates a new GetFossabotStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *GetFossabotStatus200Response) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GetFossabotStatus200Response) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GetFossabotStatus200Response) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetAppUrl

`func (o *GetFossabotStatus200Response) GetAppUrl() string`

GetAppUrl returns the AppUrl field if non-nil, zero value otherwise.

### GetAppUrlOk

`func (o *GetFossabotStatus200Response) GetAppUrlOk() (*string, bool)`

GetAppUrlOk returns a tuple with the AppUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUrl

`func (o *GetFossabotStatus200Response) SetAppUrl(v string)`

SetAppUrl sets AppUrl field to given value.


### SetAppUrlNil

`func (o *GetFossabotStatus200Response) SetAppUrlNil(b bool)`

 SetAppUrlNil sets the value for AppUrl to be an explicit nil

### UnsetAppUrl
`func (o *GetFossabotStatus200Response) UnsetAppUrl()`

UnsetAppUrl ensures that no value is present for AppUrl, not even an explicit nil
### GetCreditLevel

`func (o *GetFossabotStatus200Response) GetCreditLevel() string`

GetCreditLevel returns the CreditLevel field if non-nil, zero value otherwise.

### GetCreditLevelOk

`func (o *GetFossabotStatus200Response) GetCreditLevelOk() (*string, bool)`

GetCreditLevelOk returns a tuple with the CreditLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditLevel

`func (o *GetFossabotStatus200Response) SetCreditLevel(v string)`

SetCreditLevel sets CreditLevel field to given value.


### SetCreditLevelNil

`func (o *GetFossabotStatus200Response) SetCreditLevelNil(b bool)`

 SetCreditLevelNil sets the value for CreditLevel to be an explicit nil

### UnsetCreditLevel
`func (o *GetFossabotStatus200Response) UnsetCreditLevel()`

UnsetCreditLevel ensures that no value is present for CreditLevel, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


