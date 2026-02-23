# UpdateProjectRequestIssueTrackerCustomFieldsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **string** | The Jira custom field ID | [optional] 
**DisplayName** | Pointer to **string** | Display name for the field in FOSSA UI | [optional] 
**IsRequired** | Pointer to **bool** | Whether this field is required when creating issues (accepts \&quot;true\&quot;/\&quot;false\&quot; strings which are converted to boolean) | [optional] 
**DefaultValue** | Pointer to **string** | Default value for the field | [optional] 

## Methods

### NewUpdateProjectRequestIssueTrackerCustomFieldsValue

`func NewUpdateProjectRequestIssueTrackerCustomFieldsValue() *UpdateProjectRequestIssueTrackerCustomFieldsValue`

NewUpdateProjectRequestIssueTrackerCustomFieldsValue instantiates a new UpdateProjectRequestIssueTrackerCustomFieldsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProjectRequestIssueTrackerCustomFieldsValueWithDefaults

`func NewUpdateProjectRequestIssueTrackerCustomFieldsValueWithDefaults() *UpdateProjectRequestIssueTrackerCustomFieldsValue`

NewUpdateProjectRequestIssueTrackerCustomFieldsValueWithDefaults instantiates a new UpdateProjectRequestIssueTrackerCustomFieldsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsRequired

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UpdateProjectRequestIssueTrackerCustomFieldsValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


