# CreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | Scope of the role | 
**Name** | **string** | Name of the role | 
**Description** | **string** | Description of the role | 
**Permissions** | Pointer to [**[]CreateRoleRequestPermissionsInner**](CreateRoleRequestPermissionsInner.md) | List of permissions for this role | [optional] 

## Methods

### NewCreateRoleRequest

`func NewCreateRoleRequest(scope string, name string, description string, ) *CreateRoleRequest`

NewCreateRoleRequest instantiates a new CreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestWithDefaults

`func NewCreateRoleRequestWithDefaults() *CreateRoleRequest`

NewCreateRoleRequestWithDefaults instantiates a new CreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CreateRoleRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateRoleRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateRoleRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetName

`func (o *CreateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRoleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPermissions

`func (o *CreateRoleRequest) GetPermissions() []CreateRoleRequestPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateRoleRequest) GetPermissionsOk() (*[]CreateRoleRequestPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateRoleRequest) SetPermissions(v []CreateRoleRequestPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *CreateRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


