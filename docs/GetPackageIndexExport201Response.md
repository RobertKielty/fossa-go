# GetPackageIndexExport201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to [**GetPackageIndexExport201ResponseTask**](GetPackageIndexExport201ResponseTask.md) |  | [optional] 
**Target** | Pointer to **string** | Email address where the report link will be sent | [optional] 

## Methods

### NewGetPackageIndexExport201Response

`func NewGetPackageIndexExport201Response() *GetPackageIndexExport201Response`

NewGetPackageIndexExport201Response instantiates a new GetPackageIndexExport201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPackageIndexExport201ResponseWithDefaults

`func NewGetPackageIndexExport201ResponseWithDefaults() *GetPackageIndexExport201Response`

NewGetPackageIndexExport201ResponseWithDefaults instantiates a new GetPackageIndexExport201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *GetPackageIndexExport201Response) GetTask() GetPackageIndexExport201ResponseTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetPackageIndexExport201Response) GetTaskOk() (*GetPackageIndexExport201ResponseTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetPackageIndexExport201Response) SetTask(v GetPackageIndexExport201ResponseTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetPackageIndexExport201Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetTarget

`func (o *GetPackageIndexExport201Response) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GetPackageIndexExport201Response) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GetPackageIndexExport201Response) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GetPackageIndexExport201Response) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


