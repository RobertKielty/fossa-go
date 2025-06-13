# GetAuditLogsExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | Filter audit logs to those created on or after this date (YYYY-MM-DD format) | 
**EndDate** | **string** | Filter audit logs to those created on or before this date (YYYY-MM-DD format) | 
**ActingUserIds** | Pointer to **[]string** | Filter audit logs to those created by the given user IDs | [optional] 
**Actions** | Pointer to **[]string** | Filter audit logs to those with the given actions | [optional] 
**Topics** | Pointer to **[]string** | Filter audit logs to those with the given topics | [optional] 
**TopicActions** | Pointer to **[]string** | Filter audit logs to those with the given topic and action | [optional] 

## Methods

### NewGetAuditLogsExportRequest

`func NewGetAuditLogsExportRequest(startDate string, endDate string, ) *GetAuditLogsExportRequest`

NewGetAuditLogsExportRequest instantiates a new GetAuditLogsExportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuditLogsExportRequestWithDefaults

`func NewGetAuditLogsExportRequestWithDefaults() *GetAuditLogsExportRequest`

NewGetAuditLogsExportRequestWithDefaults instantiates a new GetAuditLogsExportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *GetAuditLogsExportRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetAuditLogsExportRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetAuditLogsExportRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *GetAuditLogsExportRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetAuditLogsExportRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetAuditLogsExportRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetActingUserIds

`func (o *GetAuditLogsExportRequest) GetActingUserIds() []string`

GetActingUserIds returns the ActingUserIds field if non-nil, zero value otherwise.

### GetActingUserIdsOk

`func (o *GetAuditLogsExportRequest) GetActingUserIdsOk() (*[]string, bool)`

GetActingUserIdsOk returns a tuple with the ActingUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserIds

`func (o *GetAuditLogsExportRequest) SetActingUserIds(v []string)`

SetActingUserIds sets ActingUserIds field to given value.

### HasActingUserIds

`func (o *GetAuditLogsExportRequest) HasActingUserIds() bool`

HasActingUserIds returns a boolean if a field has been set.

### GetActions

`func (o *GetAuditLogsExportRequest) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetAuditLogsExportRequest) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetAuditLogsExportRequest) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetAuditLogsExportRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetTopics

`func (o *GetAuditLogsExportRequest) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *GetAuditLogsExportRequest) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *GetAuditLogsExportRequest) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *GetAuditLogsExportRequest) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetTopicActions

`func (o *GetAuditLogsExportRequest) GetTopicActions() []string`

GetTopicActions returns the TopicActions field if non-nil, zero value otherwise.

### GetTopicActionsOk

`func (o *GetAuditLogsExportRequest) GetTopicActionsOk() (*[]string, bool)`

GetTopicActionsOk returns a tuple with the TopicActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicActions

`func (o *GetAuditLogsExportRequest) SetTopicActions(v []string)`

SetTopicActions sets TopicActions field to given value.

### HasTopicActions

`func (o *GetAuditLogsExportRequest) HasTopicActions() bool`

HasTopicActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


