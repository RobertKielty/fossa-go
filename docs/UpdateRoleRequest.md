# UpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the role | [optional] 
**Description** | Pointer to **string** | New description for the role | [optional] 
**Permissions** | Pointer to [**[]CreateRoleRequestPermissionsInner**](CreateRoleRequestPermissionsInner.md) | Updated list of permissions for this role | [optional] 

## Methods

### NewUpdateRoleRequest

`func NewUpdateRoleRequest() *UpdateRoleRequest`

NewUpdateRoleRequest instantiates a new UpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleRequestWithDefaults

`func NewUpdateRoleRequestWithDefaults() *UpdateRoleRequest`

NewUpdateRoleRequestWithDefaults instantiates a new UpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRoleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRoleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRoleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRoleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRoleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRoleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRoleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRoleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateRoleRequest) GetPermissions() []CreateRoleRequestPermissionsInner`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateRoleRequest) GetPermissionsOk() (*[]CreateRoleRequestPermissionsInner, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateRoleRequest) SetPermissions(v []CreateRoleRequestPermissionsInner)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateRoleRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


