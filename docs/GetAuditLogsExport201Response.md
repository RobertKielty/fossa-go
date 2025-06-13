# GetAuditLogsExport201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to [**GetAuditLogsExport201ResponseTask**](GetAuditLogsExport201ResponseTask.md) |  | [optional] 
**Target** | Pointer to **string** | Email address where the export link will be sent | [optional] 

## Methods

### NewGetAuditLogsExport201Response

`func NewGetAuditLogsExport201Response() *GetAuditLogsExport201Response`

NewGetAuditLogsExport201Response instantiates a new GetAuditLogsExport201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogsExport201ResponseWithDefaults

`func NewGetAuditLogsExport201ResponseWithDefaults() *GetAuditLogsExport201Response`

NewGetAuditLogsExport201ResponseWithDefaults instantiates a new GetAuditLogsExport201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *GetAuditLogsExport201Response) GetTask() GetAuditLogsExport201ResponseTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetAuditLogsExport201Response) GetTaskOk() (*GetAuditLogsExport201ResponseTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetAuditLogsExport201Response) SetTask(v GetAuditLogsExport201ResponseTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetAuditLogsExport201Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTarget

`func (o *GetAuditLogsExport201Response) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetAuditLogsExport201Response) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetAuditLogsExport201Response) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetAuditLogsExport201Response) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


