# QueueReleaseGroupAttributionReportV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | The job token of the queued report-generation task | 
**ReportId** | **string** | The report&#39;s ID. Once the task finishes, fetch the report from &#x60;https://&lt;host&gt;/attribution/{reportId}?download&#x3D;true&#x60; (not under &#x60;/api&#x60;), which responds with a redirect to a short-lived signed URL.  | 

## Methods

### NewQueueReleaseGroupAttributionReportV2200Response

`func NewQueueReleaseGroupAttributionReportV2200Response(taskId string, reportId string, ) *QueueReleaseGroupAttributionReportV2200Response`

NewQueueReleaseGroupAttributionReportV2200Response instantiates a new QueueReleaseGroupAttributionReportV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueReleaseGroupAttributionReportV2200ResponseWithDefaults

`func NewQueueReleaseGroupAttributionReportV2200ResponseWithDefaults() *QueueReleaseGroupAttributionReportV2200Response`

NewQueueReleaseGroupAttributionReportV2200ResponseWithDefaults instantiates a new QueueReleaseGroupAttributionReportV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *QueueReleaseGroupAttributionReportV2200Response) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *QueueReleaseGroupAttributionReportV2200Response) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *QueueReleaseGroupAttributionReportV2200Response) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetReportId

`func (o *QueueReleaseGroupAttributionReportV2200Response) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *QueueReleaseGroupAttributionReportV2200Response) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *QueueReleaseGroupAttributionReportV2200Response) SetReportId(v string)`

SetReportId sets ReportId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


