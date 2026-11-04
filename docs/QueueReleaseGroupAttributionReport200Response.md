# QueueReleaseGroupAttributionReport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **int32** |  | [optional] 
**ReportId** | Pointer to **string** | The report&#39;s ID. Once the task finishes, fetch the report from &#x60;https://&lt;host&gt;/attribution/{reportId}?download&#x3D;true&#x60; (not under &#x60;/api&#x60;), which responds with a redirect to a short-lived signed URL.  | [optional] 

## Methods

### NewQueueReleaseGroupAttributionReport200Response

`func NewQueueReleaseGroupAttributionReport200Response() *QueueReleaseGroupAttributionReport200Response`

NewQueueReleaseGroupAttributionReport200Response instantiates a new QueueReleaseGroupAttributionReport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueReleaseGroupAttributionReport200ResponseWithDefaults

`func NewQueueReleaseGroupAttributionReport200ResponseWithDefaults() *QueueReleaseGroupAttributionReport200Response`

NewQueueReleaseGroupAttributionReport200ResponseWithDefaults instantiates a new QueueReleaseGroupAttributionReport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *QueueReleaseGroupAttributionReport200Response) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *QueueReleaseGroupAttributionReport200Response) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *QueueReleaseGroupAttributionReport200Response) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *QueueReleaseGroupAttributionReport200Response) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetReportId

`func (o *QueueReleaseGroupAttributionReport200Response) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *QueueReleaseGroupAttributionReport200Response) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *QueueReleaseGroupAttributionReport200Response) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *QueueReleaseGroupAttributionReport200Response) HasReportId() bool`

HasReportId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


