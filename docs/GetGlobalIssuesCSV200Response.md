# GetGlobalIssuesCSV200Response

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
**ReportId** | Pointer to **string** | Only with &#x60;delivery&#x3D;async-download&#x60;. Once the task finishes, fetch the report from &#x60;https://&lt;host&gt;/attribution/{reportId}?download&#x3D;true&#x60; (not under &#x60;/api&#x60;), which responds with a redirect to a short-lived signed URL.  | [optional] 

## Methods

### NewGetGlobalIssuesCSV200Response

`func NewGetGlobalIssuesCSV200Response() *GetGlobalIssuesCSV200Response`

NewGetGlobalIssuesCSV200Response instantiates a new GetGlobalIssuesCSV200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGlobalIssuesCSV200ResponseWithDefaults

`func NewGetGlobalIssuesCSV200ResponseWithDefaults() *GetGlobalIssuesCSV200Response`

NewGetGlobalIssuesCSV200ResponseWithDefaults instantiates a new GetGlobalIssuesCSV200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGlobalIssuesCSV200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGlobalIssuesCSV200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGlobalIssuesCSV200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetGlobalIssuesCSV200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTask

`func (o *GetGlobalIssuesCSV200Response) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetGlobalIssuesCSV200Response) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetGlobalIssuesCSV200Response) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetGlobalIssuesCSV200Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetJobToken

`func (o *GetGlobalIssuesCSV200Response) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *GetGlobalIssuesCSV200Response) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *GetGlobalIssuesCSV200Response) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *GetGlobalIssuesCSV200Response) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.

### GetStatus

`func (o *GetGlobalIssuesCSV200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGlobalIssuesCSV200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGlobalIssuesCSV200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGlobalIssuesCSV200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAttemptNumber

`func (o *GetGlobalIssuesCSV200Response) GetAttemptNumber() int32`

GetAttemptNumber returns the AttemptNumber field if non-nil, zero value otherwise.

### GetAttemptNumberOk

`func (o *GetGlobalIssuesCSV200Response) GetAttemptNumberOk() (*int32, bool)`

GetAttemptNumberOk returns a tuple with the AttemptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptNumber

`func (o *GetGlobalIssuesCSV200Response) SetAttemptNumber(v int32)`

SetAttemptNumber sets AttemptNumber field to given value.

### HasAttemptNumber

`func (o *GetGlobalIssuesCSV200Response) HasAttemptNumber() bool`

HasAttemptNumber returns a boolean if a field has been set.

### GetMaxRetries

`func (o *GetGlobalIssuesCSV200Response) GetMaxRetries() int32`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *GetGlobalIssuesCSV200Response) GetMaxRetriesOk() (*int32, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *GetGlobalIssuesCSV200Response) SetMaxRetries(v int32)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *GetGlobalIssuesCSV200Response) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.

### SetMaxRetriesNil

`func (o *GetGlobalIssuesCSV200Response) SetMaxRetriesNil(b bool)`

 SetMaxRetriesNil sets the value for MaxRetries to be an explicit nil

### UnsetMaxRetries
`func (o *GetGlobalIssuesCSV200Response) UnsetMaxRetries()`

UnsetMaxRetries ensures that no value is present for MaxRetries, not even an explicit nil
### GetStarted

`func (o *GetGlobalIssuesCSV200Response) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetGlobalIssuesCSV200Response) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetGlobalIssuesCSV200Response) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetGlobalIssuesCSV200Response) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### SetStartedNil

`func (o *GetGlobalIssuesCSV200Response) SetStartedNil(b bool)`

 SetStartedNil sets the value for Started to be an explicit nil

### UnsetStarted
`func (o *GetGlobalIssuesCSV200Response) UnsetStarted()`

UnsetStarted ensures that no value is present for Started, not even an explicit nil
### GetFinished

`func (o *GetGlobalIssuesCSV200Response) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetGlobalIssuesCSV200Response) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetGlobalIssuesCSV200Response) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetGlobalIssuesCSV200Response) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### SetFinishedNil

`func (o *GetGlobalIssuesCSV200Response) SetFinishedNil(b bool)`

 SetFinishedNil sets the value for Finished to be an explicit nil

### UnsetFinished
`func (o *GetGlobalIssuesCSV200Response) UnsetFinished()`

UnsetFinished ensures that no value is present for Finished, not even an explicit nil
### GetScheduledStartTime

`func (o *GetGlobalIssuesCSV200Response) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *GetGlobalIssuesCSV200Response) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *GetGlobalIssuesCSV200Response) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *GetGlobalIssuesCSV200Response) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### SetScheduledStartTimeNil

`func (o *GetGlobalIssuesCSV200Response) SetScheduledStartTimeNil(b bool)`

 SetScheduledStartTimeNil sets the value for ScheduledStartTime to be an explicit nil

### UnsetScheduledStartTime
`func (o *GetGlobalIssuesCSV200Response) UnsetScheduledStartTime()`

UnsetScheduledStartTime ensures that no value is present for ScheduledStartTime, not even an explicit nil
### GetPod

`func (o *GetGlobalIssuesCSV200Response) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *GetGlobalIssuesCSV200Response) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *GetGlobalIssuesCSV200Response) SetPod(v string)`

SetPod sets Pod field to given value.

### HasPod

`func (o *GetGlobalIssuesCSV200Response) HasPod() bool`

HasPod returns a boolean if a field has been set.

### SetPodNil

`func (o *GetGlobalIssuesCSV200Response) SetPodNil(b bool)`

 SetPodNil sets the value for Pod to be an explicit nil

### UnsetPod
`func (o *GetGlobalIssuesCSV200Response) UnsetPod()`

UnsetPod ensures that no value is present for Pod, not even an explicit nil
### GetContext

`func (o *GetGlobalIssuesCSV200Response) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetGlobalIssuesCSV200Response) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetGlobalIssuesCSV200Response) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *GetGlobalIssuesCSV200Response) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetGlobalIssuesCSV200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetGlobalIssuesCSV200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetGlobalIssuesCSV200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetGlobalIssuesCSV200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetGlobalIssuesCSV200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetGlobalIssuesCSV200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetGlobalIssuesCSV200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetGlobalIssuesCSV200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReportId

`func (o *GetGlobalIssuesCSV200Response) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *GetGlobalIssuesCSV200Response) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *GetGlobalIssuesCSV200Response) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *GetGlobalIssuesCSV200Response) HasReportId() bool`

HasReportId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


