# CreateRoleRequestPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Resource type for the permission. Must be a valid &#x60;resourceType&#x60; returned by &#x60;GET /api/roles/all-permissions&#x60;, and the &#x60;resourceType&#x60;/&#x60;action&#x60; combination must be valid for the role&#39;s &#x60;scope&#x60;.  | 
**Action** | **string** | Action for the permission. Must be a valid &#x60;action&#x60; returned by &#x60;GET /api/roles/all-permissions&#x60;, and the &#x60;resourceType&#x60;/&#x60;action&#x60; combination must be valid for the role&#39;s &#x60;scope&#x60;.  | 

## Methods

### NewCreateRoleRequestPermissionsInner

`func NewCreateRoleRequestPermissionsInner(resourceType string, action string, ) *CreateRoleRequestPermissionsInner`

NewCreateRoleRequestPermissionsInner instantiates a new CreateRoleRequestPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoleRequestPermissionsInnerWithDefaults

`func NewCreateRoleRequestPermissionsInnerWithDefaults() *CreateRoleRequestPermissionsInner`

NewCreateRoleRequestPermissionsInnerWithDefaults instantiates a new CreateRoleRequestPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *CreateRoleRequestPermissionsInner) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateRoleRequestPermissionsInner) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateRoleRequestPermissionsInner) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetAction

`func (o *CreateRoleRequestPermissionsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateRoleRequestPermissionsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateRoleRequestPermissionsInner) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


