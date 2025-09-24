# GetTeamGroupById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team group ID | [optional] 
**Name** | Pointer to **string** | Team group name | [optional] 
**DefaultRoleId** | Pointer to **int32** | Default role ID for members | [optional] 
**Teams** | Pointer to [**[]GetTeamGroups200ResponseInnerTeamsInner**](GetTeamGroups200ResponseInnerTeamsInner.md) | Teams within this team group | [optional] 
**Members** | Pointer to [**[]GetTeamGroupById200ResponseMembersInner**](GetTeamGroupById200ResponseMembersInner.md) | Members of this team group with user details | [optional] 

## Methods

### NewGetTeamGroupById200Response

`func NewGetTeamGroupById200Response() *GetTeamGroupById200Response`

NewGetTeamGroupById200Response instantiates a new GetTeamGroupById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamGroupById200ResponseWithDefaults

`func NewGetTeamGroupById200ResponseWithDefaults() *GetTeamGroupById200Response`

NewGetTeamGroupById200ResponseWithDefaults instantiates a new GetTeamGroupById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTeamGroupById200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamGroupById200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamGroupById200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTeamGroupById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetTeamGroupById200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTeamGroupById200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTeamGroupById200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTeamGroupById200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *GetTeamGroupById200Response) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *GetTeamGroupById200Response) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *GetTeamGroupById200Response) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *GetTeamGroupById200Response) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.

### GetTeams

`func (o *GetTeamGroupById200Response) GetTeams() []GetTeamGroups200ResponseInnerTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetTeamGroupById200Response) GetTeamsOk() (*[]GetTeamGroups200ResponseInnerTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetTeamGroupById200Response) SetTeams(v []GetTeamGroups200ResponseInnerTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GetTeamGroupById200Response) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetMembers

`func (o *GetTeamGroupById200Response) GetMembers() []GetTeamGroupById200ResponseMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetTeamGroupById200Response) GetMembersOk() (*[]GetTeamGroupById200ResponseMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetTeamGroupById200Response) SetMembers(v []GetTeamGroupById200ResponseMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetTeamGroupById200Response) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


