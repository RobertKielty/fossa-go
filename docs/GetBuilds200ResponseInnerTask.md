# GetBuilds200ResponseInnerTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the task | [optional] 
**Started** | Pointer to **time.Time** | The date and time the task was started | [optional] 
**Finished** | Pointer to **time.Time** | The date and time the task was finished | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time the task was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the task was last updated | [optional] 
**AttemptNumber** | Pointer to **int32** | The number of attempts for the task | [optional] 
**MaxRetries** | Pointer to **int32** | The maximum number of retries for the task | [optional] 

## Methods

### NewGetBuilds200ResponseInnerTask

`func NewGetBuilds200ResponseInnerTask() *GetBuilds200ResponseInnerTask`

NewGetBuilds200ResponseInnerTask instantiates a new GetBuilds200ResponseInnerTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuilds200ResponseInnerTaskWithDefaults

`func NewGetBuilds200ResponseInnerTaskWithDefaults() *GetBuilds200ResponseInnerTask`

NewGetBuilds200ResponseInnerTaskWithDefaults instantiates a new GetBuilds200ResponseInnerTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetBuilds200ResponseInnerTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBuilds200ResponseInnerTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBuilds200ResponseInnerTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBuilds200ResponseInnerTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStarted

`func (o *GetBuilds200ResponseInnerTask) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetBuilds200ResponseInnerTask) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetBuilds200ResponseInnerTask) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetBuilds200ResponseInnerTask) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetFinished

`func (o *GetBuilds200ResponseInnerTask) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetBuilds200ResponseInnerTask) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetBuilds200ResponseInnerTask) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetBuilds200ResponseInnerTask) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetBuilds200ResponseInnerTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBuilds200ResponseInnerTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBuilds200ResponseInnerTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetBuilds200ResponseInnerTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetBuilds200ResponseInnerTask) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetBuilds200ResponseInnerTask) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetBuilds200ResponseInnerTask) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetBuilds200ResponseInnerTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *GetBuilds200ResponseInnerTask) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *GetBuilds200ResponseInnerTask) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *GetBuilds200ResponseInnerTask) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *GetBuilds200ResponseInnerTask) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetBuilds200ResponseInnerTask) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetBuilds200ResponseInnerTask) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetBuilds200ResponseInnerTask) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetBuilds200ResponseInnerTask) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


