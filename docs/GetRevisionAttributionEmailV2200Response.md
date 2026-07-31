# GetRevisionAttributionEmailV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier for the task | [optional] 
**Task** | Pointer to **string** | Name of the task being run | [optional] 
**JobToken** | Pointer to **string** | Token used to poll for the task&#39;s status and result | [optional] 
**Status** | Pointer to **string** | Current status of the task | [optional] 
**AttemptNumber** | Pointer to **int32** | Number of times the task has been attempted | [optional] 
**MaxRetries** | Pointer to **NullableInt32** | Maximum number of retries permitted for the task | [optional] 
**Started** | Pointer to **NullableTime** | When the task started running, or null if it has not started | [optional] 
**Finished** | Pointer to **NullableTime** | When the task finished, or null if it has not finished | [optional] 
**ScheduledStartTime** | Pointer to **NullableTime** | When the task is scheduled to start, or null | [optional] 
**Pod** | Pointer to **NullableString** | Identifier of the pod running the task, or null | [optional] 
**Context** | Pointer to **map[string]interface{}** | Task-specific context payload | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the task record was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | When the task record was last updated | [optional] 

## Methods

### NewGetRevisionAttributionEmailV2200Response

`func NewGetRevisionAttributionEmailV2200Response() *GetRevisionAttributionEmailV2200Response`

NewGetRevisionAttributionEmailV2200Response instantiates a new GetRevisionAttributionEmailV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionAttributionEmailV2200ResponseWithDefaults

`func NewGetRevisionAttributionEmailV2200ResponseWithDefaults() *GetRevisionAttributionEmailV2200Response`

NewGetRevisionAttributionEmailV2200ResponseWithDefaults instantiates a new GetRevisionAttributionEmailV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRevisionAttributionEmailV2200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRevisionAttributionEmailV2200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRevisionAttributionEmailV2200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRevisionAttributionEmailV2200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *GetRevisionAttributionEmailV2200Response) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetRevisionAttributionEmailV2200Response) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetRevisionAttributionEmailV2200Response) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetRevisionAttributionEmailV2200Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetJobToken

`func (o *GetRevisionAttributionEmailV2200Response) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *GetRevisionAttributionEmailV2200Response) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *GetRevisionAttributionEmailV2200Response) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *GetRevisionAttributionEmailV2200Response) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.

### GetStatus

`func (o *GetRevisionAttributionEmailV2200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRevisionAttributionEmailV2200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRevisionAttributionEmailV2200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRevisionAttributionEmailV2200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *GetRevisionAttributionEmailV2200Response) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *GetRevisionAttributionEmailV2200Response) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *GetRevisionAttributionEmailV2200Response) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *GetRevisionAttributionEmailV2200Response) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetRevisionAttributionEmailV2200Response) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetRevisionAttributionEmailV2200Response) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetRevisionAttributionEmailV2200Response) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetRevisionAttributionEmailV2200Response) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *GetRevisionAttributionEmailV2200Response) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *GetRevisionAttributionEmailV2200Response) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetStarted

`func (o *GetRevisionAttributionEmailV2200Response) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetRevisionAttributionEmailV2200Response) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetRevisionAttributionEmailV2200Response) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetRevisionAttributionEmailV2200Response) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### SetStartedNil

`func (o *GetRevisionAttributionEmailV2200Response) SetStartedNil(b bool)`

 SetStartedNil sets the value for Started to be an explicit nil

### UnsetStarted
`func (o *GetRevisionAttributionEmailV2200Response) UnsetStarted()`

UnsetStarted ensures that no value is present for Started, not even an explicit nil
### GetFinished

`func (o *GetRevisionAttributionEmailV2200Response) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetRevisionAttributionEmailV2200Response) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetRevisionAttributionEmailV2200Response) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetRevisionAttributionEmailV2200Response) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *GetRevisionAttributionEmailV2200Response) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *GetRevisionAttributionEmailV2200Response) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetScheduledStartTime

`func (o *GetRevisionAttributionEmailV2200Response) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *GetRevisionAttributionEmailV2200Response) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *GetRevisionAttributionEmailV2200Response) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *GetRevisionAttributionEmailV2200Response) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *GetRevisionAttributionEmailV2200Response) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *GetRevisionAttributionEmailV2200Response) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetPod

`func (o *GetRevisionAttributionEmailV2200Response) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *GetRevisionAttributionEmailV2200Response) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *GetRevisionAttributionEmailV2200Response) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *GetRevisionAttributionEmailV2200Response) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *GetRevisionAttributionEmailV2200Response) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *GetRevisionAttributionEmailV2200Response) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetContext

`func (o *GetRevisionAttributionEmailV2200Response) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetRevisionAttributionEmailV2200Response) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetRevisionAttributionEmailV2200Response) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *GetRevisionAttributionEmailV2200Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRevisionAttributionEmailV2200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRevisionAttributionEmailV2200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRevisionAttributionEmailV2200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRevisionAttributionEmailV2200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetRevisionAttributionEmailV2200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetRevisionAttributionEmailV2200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetRevisionAttributionEmailV2200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetRevisionAttributionEmailV2200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


