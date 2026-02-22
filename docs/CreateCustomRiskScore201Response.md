# CreateCustomRiskScore201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueId** | **int32** | The ID of the issue | 
**Score** | **int32** | The custom risk score (0-100 scale) | 
**Scope** | **string** | The scope type of the custom risk score | 
**Reason** | Pointer to **string** | The reason for the custom risk score | [optional] 
**ScopeId** | **string** | The project locator or release group ID | 

## Methods

### NewCreateCustomRiskScore201Response

`func NewCreateCustomRiskScore201Response(issueId int32, score int32, scope string, scopeId string, ) *CreateCustomRiskScore201Response`

NewCreateCustomRiskScore201Response instantiates a new CreateCustomRiskScore201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomRiskScore201ResponseWithDefaults

`func NewCreateCustomRiskScore201ResponseWithDefaults() *CreateCustomRiskScore201Response`

NewCreateCustomRiskScore201ResponseWithDefaults instantiates a new CreateCustomRiskScore201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueId

`func (o *CreateCustomRiskScore201Response) GetIssueId() int32`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *CreateCustomRiskScore201Response) GetIssueIdOk() (*int32, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *CreateCustomRiskScore201Response) SetIssueId(v int32)`

SetIssueId sets IssueId field to given value.


### GetScore

`func (o *CreateCustomRiskScore201Response) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateCustomRiskScore201Response) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateCustomRiskScore201Response) SetScore(v int32)`

SetScore sets Score field to given value.


### GetScope

`func (o *CreateCustomRiskScore201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateCustomRiskScore201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateCustomRiskScore201Response) SetScope(v string)`

SetScope sets Scope field to given value.


### GetReason

`func (o *CreateCustomRiskScore201Response) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCustomRiskScore201Response) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCustomRiskScore201Response) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCustomRiskScore201Response) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetScopeId

`func (o *CreateCustomRiskScore201Response) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *CreateCustomRiskScore201Response) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *CreateCustomRiskScore201Response) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


