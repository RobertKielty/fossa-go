# PatchJiraConfigurationRequestCustomFieldsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | Pointer to **string** | The corresponding Field ID in the Jira site for the given field | [optional] 
**DisplayName** | Pointer to **string** | Display name to use in FOSSA | [optional] 
**IsRequired** | Pointer to **string** | On/off switch to tell FOSSA if the field is required when exporting tickets | [optional] 
**DefaultValue** | Pointer to **string** | When provided, is the default value used when exporting a ticket | [optional] 

## Methods

### NewPatchJiraConfigurationRequestCustomFieldsValue

`func NewPatchJiraConfigurationRequestCustomFieldsValue() *PatchJiraConfigurationRequestCustomFieldsValue`

NewPatchJiraConfigurationRequestCustomFieldsValue instantiates a new PatchJiraConfigurationRequestCustomFieldsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchJiraConfigurationRequestCustomFieldsValueWithDefaults

`func NewPatchJiraConfigurationRequestCustomFieldsValueWithDefaults() *PatchJiraConfigurationRequestCustomFieldsValue`

NewPatchJiraConfigurationRequestCustomFieldsValueWithDefaults instantiates a new PatchJiraConfigurationRequestCustomFieldsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.

### HasFieldId

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) HasFieldId() bool`

HasFieldId returns a boolean if a field has been set.

### GetDisplayName

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsRequired

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetIsRequired() string`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetIsRequiredOk() (*string, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) SetIsRequired(v string)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetDefaultValue

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *PatchJiraConfigurationRequestCustomFieldsValue) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


