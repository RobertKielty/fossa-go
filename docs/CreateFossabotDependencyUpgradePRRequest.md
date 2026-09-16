# CreateFossabotDependencyUpgradePRRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fix** | Pointer to **string** | Which remediation target to use. Defaults to complete. | [optional] 
**ProjectLocator** | Pointer to **string** | Required only when the issue affects more than one project. | [optional] 

## Methods

### NewCreateFossabotDependencyUpgradePRRequest

`func NewCreateFossabotDependencyUpgradePRRequest() *CreateFossabotDependencyUpgradePRRequest`

NewCreateFossabotDependencyUpgradePRRequest instantiates a new CreateFossabotDependencyUpgradePRRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFossabotDependencyUpgradePRRequestWithDefaults

`func NewCreateFossabotDependencyUpgradePRRequestWithDefaults() *CreateFossabotDependencyUpgradePRRequest`

NewCreateFossabotDependencyUpgradePRRequestWithDefaults instantiates a new CreateFossabotDependencyUpgradePRRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFix

`func (o *CreateFossabotDependencyUpgradePRRequest) GetFix() string`

GetFix returns the Fix field if non-nil, zero value otherwise.

### GetFixOk

`func (o *CreateFossabotDependencyUpgradePRRequest) GetFixOk() (*string, bool)`

GetFixOk returns a tuple with the Fix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFix

`func (o *CreateFossabotDependencyUpgradePRRequest) SetFix(v string)`

SetFix sets Fix field to given value.

### HasFix

`func (o *CreateFossabotDependencyUpgradePRRequest) HasFix() bool`

HasFix returns a boolean if a field has been set.

### GetProjectLocator

`func (o *CreateFossabotDependencyUpgradePRRequest) GetProjectLocator() string`

GetProjectLocator returns the ProjectLocator field if non-nil, zero value otherwise.

### GetProjectLocatorOk

`func (o *CreateFossabotDependencyUpgradePRRequest) GetProjectLocatorOk() (*string, bool)`

GetProjectLocatorOk returns a tuple with the ProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLocator

`func (o *CreateFossabotDependencyUpgradePRRequest) SetProjectLocator(v string)`

SetProjectLocator sets ProjectLocator field to given value.

### HasProjectLocator

`func (o *CreateFossabotDependencyUpgradePRRequest) HasProjectLocator() bool`

HasProjectLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


