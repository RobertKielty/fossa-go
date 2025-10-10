# GetAssignableRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignableOrgRoles** | Pointer to **[]int32** | Array of role IDs that can be assigned for organization-scoped roles | [optional] 
**AssignableTeamRoles** | Pointer to **map[string][]int32** | Object mapping team IDs to arrays of role IDs that can be assigned for those teams | [optional] 

## Methods

### NewGetAssignableRoles200Response

`func NewGetAssignableRoles200Response() *GetAssignableRoles200Response`

NewGetAssignableRoles200Response instantiates a new GetAssignableRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssignableRoles200ResponseWithDefaults

`func NewGetAssignableRoles200ResponseWithDefaults() *GetAssignableRoles200Response`

NewGetAssignableRoles200ResponseWithDefaults instantiates a new GetAssignableRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignableOrgRoles

`func (o *GetAssignableRoles200Response) GetAssignableOrgRoles() []int32`

GetAssignableOrgRoles returns the AssignableOrgRoles field if non-nil, zero value otherwise.

### GetAssignableOrgRolesOk

`func (o *GetAssignableRoles200Response) GetAssignableOrgRolesOk() (*[]int32, bool)`

GetAssignableOrgRolesOk returns a tuple with the AssignableOrgRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableOrgRoles

`func (o *GetAssignableRoles200Response) SetAssignableOrgRoles(v []int32)`

SetAssignableOrgRoles sets AssignableOrgRoles field to given value.

### HasAssignableOrgRoles

`func (o *GetAssignableRoles200Response) HasAssignableOrgRoles() bool`

HasAssignableOrgRoles returns a boolean if a field has been set.

### GetAssignableTeamRoles

`func (o *GetAssignableRoles200Response) GetAssignableTeamRoles() map[string][]int32`

GetAssignableTeamRoles returns the AssignableTeamRoles field if non-nil, zero value otherwise.

### GetAssignableTeamRolesOk

`func (o *GetAssignableRoles200Response) GetAssignableTeamRolesOk() (*map[string][]int32, bool)`

GetAssignableTeamRolesOk returns a tuple with the AssignableTeamRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignableTeamRoles

`func (o *GetAssignableRoles200Response) SetAssignableTeamRoles(v map[string][]int32)`

SetAssignableTeamRoles sets AssignableTeamRoles field to given value.

### HasAssignableTeamRoles

`func (o *GetAssignableRoles200Response) HasAssignableTeamRoles() bool`

HasAssignableTeamRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


