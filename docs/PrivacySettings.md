# PrivacySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultProjectPrivacy** | Pointer to **string** | the default privacy setting for new projects | [optional] 

## Methods

### NewPrivacySettings

`func NewPrivacySettings() *PrivacySettings`

NewPrivacySettings instantiates a new PrivacySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacySettingsWithDefaults

`func NewPrivacySettingsWithDefaults() *PrivacySettings`

NewPrivacySettingsWithDefaults instantiates a new PrivacySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultProjectPrivacy

`func (o *PrivacySettings) GetDefaultProjectPrivacy() string`

GetDefaultProjectPrivacy returns the DefaultProjectPrivacy field if non-nil, zero value otherwise.

### GetDefaultProjectPrivacyOk

`func (o *PrivacySettings) GetDefaultProjectPrivacyOk() (*string, bool)`

GetDefaultProjectPrivacyOk returns a tuple with the DefaultProjectPrivacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectPrivacy

`func (o *PrivacySettings) SetDefaultProjectPrivacy(v string)`

SetDefaultProjectPrivacy sets DefaultProjectPrivacy field to given value.

### HasDefaultProjectPrivacy

`func (o *PrivacySettings) HasDefaultProjectPrivacy() bool`

HasDefaultProjectPrivacy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


