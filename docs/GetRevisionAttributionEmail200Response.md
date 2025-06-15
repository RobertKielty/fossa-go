# GetRevisionAttributionEmail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptNumber** | Pointer to **int32** | The number of attempts | [optional] 
**Id** | Pointer to **int32** | The ID of the task | [optional] 
**Task** | Pointer to **string** | The task to be performed | [optional] 
**Context** | Pointer to [**GetRevisionAttributionEmail200ResponseContext**](GetRevisionAttributionEmail200ResponseContext.md) |  | [optional] 
**MaxRetries** | Pointer to **int32** | The maximum number of retries | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** | The scheduled start time | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the task was last updated | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the task was created | [optional] 
**Started** | Pointer to **time.Time** | The time the task was started | [optional] 
**Finished** | Pointer to **time.Time** | The time the task was finished | [optional] 
**Pod** | Pointer to **string** | The pod the task is running on | [optional] 
**Status** | Pointer to **string** | The status of the task | [optional] 
**JobToken** | Pointer to **string** | The token of the job in the FOSSA Backend | [optional] 

## Methods

### NewGetRevisionAttributionEmail200Response

`func NewGetRevisionAttributionEmail200Response() *GetRevisionAttributionEmail200Response`

NewGetRevisionAttributionEmail200Response instantiates a new GetRevisionAttributionEmail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionAttributionEmail200ResponseWithDefaults

`func NewGetRevisionAttributionEmail200ResponseWithDefaults() *GetRevisionAttributionEmail200Response`

NewGetRevisionAttributionEmail200ResponseWithDefaults instantiates a new GetRevisionAttributionEmail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptNumber

`func (o *GetRevisionAttributionEmail200Response) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *GetRevisionAttributionEmail200Response) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *GetRevisionAttributionEmail200Response) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *GetRevisionAttributionEmail200Response) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetId

`func (o *GetRevisionAttributionEmail200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRevisionAttributionEmail200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRevisionAttributionEmail200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRevisionAttributionEmail200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *GetRevisionAttributionEmail200Response) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetRevisionAttributionEmail200Response) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetRevisionAttributionEmail200Response) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetRevisionAttributionEmail200Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetContext

`func (o *GetRevisionAttributionEmail200Response) GetContext() GetRevisionAttributionEmail200ResponseContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetRevisionAttributionEmail200Response) GetContextOk() (*GetRevisionAttributionEmail200ResponseContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetRevisionAttributionEmail200Response) SetContext(v GetRevisionAttributionEmail200ResponseContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetRevisionAttributionEmail200Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetRevisionAttributionEmail200Response) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetRevisionAttributionEmail200Response) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetRevisionAttributionEmail200Response) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetRevisionAttributionEmail200Response) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### GetScheduledStartTime

`func (o *GetRevisionAttributionEmail200Response) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *GetRevisionAttributionEmail200Response) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *GetRevisionAttributionEmail200Response) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *GetRevisionAttributionEmail200Response) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetRevisionAttributionEmail200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetRevisionAttributionEmail200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetRevisionAttributionEmail200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetRevisionAttributionEmail200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRevisionAttributionEmail200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRevisionAttributionEmail200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRevisionAttributionEmail200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRevisionAttributionEmail200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStarted

`func (o *GetRevisionAttributionEmail200Response) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetRevisionAttributionEmail200Response) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetRevisionAttributionEmail200Response) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetRevisionAttributionEmail200Response) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetFinished

`func (o *GetRevisionAttributionEmail200Response) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetRevisionAttributionEmail200Response) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetRevisionAttributionEmail200Response) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetRevisionAttributionEmail200Response) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetPod

`func (o *GetRevisionAttributionEmail200Response) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *GetRevisionAttributionEmail200Response) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *GetRevisionAttributionEmail200Response) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *GetRevisionAttributionEmail200Response) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetStatus

`func (o *GetRevisionAttributionEmail200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRevisionAttributionEmail200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRevisionAttributionEmail200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRevisionAttributionEmail200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetJobToken

`func (o *GetRevisionAttributionEmail200Response) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *GetRevisionAttributionEmail200Response) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *GetRevisionAttributionEmail200Response) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *GetRevisionAttributionEmail200Response) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


