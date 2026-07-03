# GetIssues201Response

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

### NewGetIssues201Response

`func NewGetIssues201Response() *GetIssues201Response`

NewGetIssues201Response instantiates a new GetIssues201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssues201ResponseWithDefaults

`func NewGetIssues201ResponseWithDefaults() *GetIssues201Response`

NewGetIssues201ResponseWithDefaults instantiates a new GetIssues201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssues201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssues201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssues201Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssues201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *GetIssues201Response) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetIssues201Response) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetIssues201Response) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetIssues201Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetJobToken

`func (o *GetIssues201Response) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *GetIssues201Response) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *GetIssues201Response) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *GetIssues201Response) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.

### GetStatus

`func (o *GetIssues201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIssues201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIssues201Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIssues201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *GetIssues201Response) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *GetIssues201Response) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *GetIssues201Response) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *GetIssues201Response) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetIssues201Response) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetIssues201Response) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetIssues201Response) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetIssues201Response) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *GetIssues201Response) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *GetIssues201Response) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetStarted

`func (o *GetIssues201Response) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetIssues201Response) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetIssues201Response) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetIssues201Response) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### SetStartedNil

`func (o *GetIssues201Response) SetStartedNil(b bool)`

 SetStartedNil sets the value for Started to be an explicit nil

### UnsetStarted
`func (o *GetIssues201Response) UnsetStarted()`

UnsetStarted ensures that no value is present for Started, not even an explicit nil
### GetFinished

`func (o *GetIssues201Response) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetIssues201Response) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetIssues201Response) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetIssues201Response) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *GetIssues201Response) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *GetIssues201Response) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetScheduledStartTime

`func (o *GetIssues201Response) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *GetIssues201Response) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *GetIssues201Response) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *GetIssues201Response) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *GetIssues201Response) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *GetIssues201Response) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetPod

`func (o *GetIssues201Response) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *GetIssues201Response) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *GetIssues201Response) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *GetIssues201Response) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *GetIssues201Response) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *GetIssues201Response) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetContext

`func (o *GetIssues201Response) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetIssues201Response) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetIssues201Response) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *GetIssues201Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssues201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssues201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssues201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssues201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetIssues201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetIssues201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetIssues201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetIssues201Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


