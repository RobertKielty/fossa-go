# GetBuilds200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID for the build | [optional] 
**Locator** | Pointer to **string** | The revision locator for the build | [optional] 
**OwnerId** | Pointer to **NullableInt32** | The ID of the user who owns the build | [optional] 
**Error** | Pointer to **NullableString** | The error message for the build | [optional] 
**Warnings** | Pointer to **[]string** | The warning messages for the build | [optional] 
**Provided** | Pointer to **bool** | Is the build from a CLI upload | [optional] 
**TaskId** | Pointer to **NullableInt32** | The ID of the task for the build | [optional] 
**CreatedAt** | Pointer to **time.Time** | The date and time the build was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The date and time the build was last updated | [optional] 
**CliVersionId** | Pointer to **int32** | The ID of the CLI version for the build | [optional] 
**Revision** | Pointer to [**GetBuilds200ResponseInnerRevision**](GetBuilds200ResponseInnerRevision.md) |  | [optional] 
**Task** | Pointer to [**GetBuilds200ResponseInnerTask**](GetBuilds200ResponseInnerTask.md) |  | [optional] 

## Methods

### NewGetBuilds200ResponseInner

`func NewGetBuilds200ResponseInner() *GetBuilds200ResponseInner`

NewGetBuilds200ResponseInner instantiates a new GetBuilds200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuilds200ResponseInnerWithDefaults

`func NewGetBuilds200ResponseInnerWithDefaults() *GetBuilds200ResponseInner`

NewGetBuilds200ResponseInnerWithDefaults instantiates a new GetBuilds200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBuilds200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBuilds200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBuilds200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetBuilds200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocator

`func (o *GetBuilds200ResponseInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetBuilds200ResponseInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetBuilds200ResponseInner) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *GetBuilds200ResponseInner) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetOwnerId

`func (o *GetBuilds200ResponseInner) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetBuilds200ResponseInner) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetBuilds200ResponseInner) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GetBuilds200ResponseInner) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *GetBuilds200ResponseInner) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *GetBuilds200ResponseInner) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetError

`func (o *GetBuilds200ResponseInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetBuilds200ResponseInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetBuilds200ResponseInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetBuilds200ResponseInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetBuilds200ResponseInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetBuilds200ResponseInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetWarnings

`func (o *GetBuilds200ResponseInner) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GetBuilds200ResponseInner) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GetBuilds200ResponseInner) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GetBuilds200ResponseInner) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetProvided

`func (o *GetBuilds200ResponseInner) GetProvided() bool`

GetProvided returns the Provided field if non-nil, zero value otherwise.

### GetProvidedOk

`func (o *GetBuilds200ResponseInner) GetProvidedOk() (*bool, bool)`

GetProvidedOk returns a tuple with the Provided field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvided

`func (o *GetBuilds200ResponseInner) SetProvided(v bool)`

SetProvided sets Provided field to given value.

### HasProvided

`func (o *GetBuilds200ResponseInner) HasProvided() bool`

HasProvided returns a boolean if a field has been set.

### GetTaskId

`func (o *GetBuilds200ResponseInner) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *GetBuilds200ResponseInner) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *GetBuilds200ResponseInner) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *GetBuilds200ResponseInner) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *GetBuilds200ResponseInner) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *GetBuilds200ResponseInner) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetCreatedAt

`func (o *GetBuilds200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBuilds200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBuilds200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetBuilds200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetBuilds200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetBuilds200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetBuilds200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetBuilds200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCliVersionId

`func (o *GetBuilds200ResponseInner) GetCliVersionId() int32`

GetCliVersionId returns the CliVersionId field if non-nil, zero value otherwise.

### GetCliVersionIdOk

`func (o *GetBuilds200ResponseInner) GetCliVersionIdOk() (*int32, bool)`

GetCliVersionIdOk returns a tuple with the CliVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliVersionId

`func (o *GetBuilds200ResponseInner) SetCliVersionId(v int32)`

SetCliVersionId sets CliVersionId field to given value.

### HasCliVersionId

`func (o *GetBuilds200ResponseInner) HasCliVersionId() bool`

HasCliVersionId returns a boolean if a field has been set.

### GetRevision

`func (o *GetBuilds200ResponseInner) GetRevision() GetBuilds200ResponseInnerRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetBuilds200ResponseInner) GetRevisionOk() (*GetBuilds200ResponseInnerRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetBuilds200ResponseInner) SetRevision(v GetBuilds200ResponseInnerRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GetBuilds200ResponseInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetTask

`func (o *GetBuilds200ResponseInner) GetTask() GetBuilds200ResponseInnerTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *GetBuilds200ResponseInner) GetTaskOk() (*GetBuilds200ResponseInnerTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *GetBuilds200ResponseInner) SetTask(v GetBuilds200ResponseInnerTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *GetBuilds200ResponseInner) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


