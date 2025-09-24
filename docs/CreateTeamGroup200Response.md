# CreateTeamGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team group ID | [optional] 
**Name** | Pointer to **string** | Team group name | [optional] 
**DefaultRoleId** | Pointer to **int32** | Default role ID for members | [optional] 

## Methods

### NewCreateTeamGroup200Response

`func NewCreateTeamGroup200Response() *CreateTeamGroup200Response`

NewCreateTeamGroup200Response instantiates a new CreateTeamGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamGroup200ResponseWithDefaults

`func NewCreateTeamGroup200ResponseWithDefaults() *CreateTeamGroup200Response`

NewCreateTeamGroup200ResponseWithDefaults instantiates a new CreateTeamGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateTeamGroup200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateTeamGroup200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateTeamGroup200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateTeamGroup200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateTeamGroup200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTeamGroup200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTeamGroup200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTeamGroup200Response) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultRoleId

`func (o *CreateTeamGroup200Response) GetDefaultRoleId() int32`

GetDefaultRoleId returns the DefaultRoleId field if non-nil, zero value otherwise.

### GetDefaultRoleIdOk

`func (o *CreateTeamGroup200Response) GetDefaultRoleIdOk() (*int32, bool)`

GetDefaultRoleIdOk returns a tuple with the DefaultRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRoleId

`func (o *CreateTeamGroup200Response) SetDefaultRoleId(v int32)`

SetDefaultRoleId sets DefaultRoleId field to given value.

### HasDefaultRoleId

`func (o *CreateTeamGroup200Response) HasDefaultRoleId() bool`

HasDefaultRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


