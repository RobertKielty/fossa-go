# CreateRevisionAttributionPublicReportV2202ResponseTask

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

### NewCreateRevisionAttributionPublicReportV2202ResponseTask

`func NewCreateRevisionAttributionPublicReportV2202ResponseTask() *CreateRevisionAttributionPublicReportV2202ResponseTask`

NewCreateRevisionAttributionPublicReportV2202ResponseTask instantiates a new CreateRevisionAttributionPublicReportV2202ResponseTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRevisionAttributionPublicReportV2202ResponseTaskWithDefaults

`func NewCreateRevisionAttributionPublicReportV2202ResponseTaskWithDefaults() *CreateRevisionAttributionPublicReportV2202ResponseTask`

NewCreateRevisionAttributionPublicReportV2202ResponseTaskWithDefaults instantiates a new CreateRevisionAttributionPublicReportV2202ResponseTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetJobToken

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.

### GetStatus

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetMaxRetries

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetStarted

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### SetStartedNil

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetStartedNil(b bool)`

 SetStartedNil sets the value for Started to be an explicit nil

### UnsetStarted
`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) UnsetStarted()`

UnsetStarted ensures that no value is present for Started, not even an explicit nil
### GetFinished

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetScheduledStartTime

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetPod

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetContext

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateRevisionAttributionPublicReportV2202ResponseTask) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


