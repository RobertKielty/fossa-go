# CreateCustomRiskScoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Score** | **int32** | The custom risk score (0-100 scale) | 
**Reason** | Pointer to **string** | The reason for the custom risk score | [optional] 

## Methods

### NewCreateCustomRiskScoreRequest

`func NewCreateCustomRiskScoreRequest(score int32, ) *CreateCustomRiskScoreRequest`

NewCreateCustomRiskScoreRequest instantiates a new CreateCustomRiskScoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomRiskScoreRequestWithDefaults

`func NewCreateCustomRiskScoreRequestWithDefaults() *CreateCustomRiskScoreRequest`

NewCreateCustomRiskScoreRequestWithDefaults instantiates a new CreateCustomRiskScoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScore

`func (o *CreateCustomRiskScoreRequest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateCustomRiskScoreRequest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateCustomRiskScoreRequest) SetScore(v int32)`

SetScore sets Score field to given value.


### GetReason

`func (o *CreateCustomRiskScoreRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCustomRiskScoreRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCustomRiskScoreRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCustomRiskScoreRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


