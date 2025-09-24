# AddReleaseGroupsToTeam200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Team ID | [optional] 
**ReleaseGroups** | Pointer to **[]int32** | Updated list of release group IDs assigned to the team | [optional] 

## Methods

### NewAddReleaseGroupsToTeam200Response

`func NewAddReleaseGroupsToTeam200Response() *AddReleaseGroupsToTeam200Response`

NewAddReleaseGroupsToTeam200Response instantiates a new AddReleaseGroupsToTeam200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddReleaseGroupsToTeam200ResponseWithDefaults

`func NewAddReleaseGroupsToTeam200ResponseWithDefaults() *AddReleaseGroupsToTeam200Response`

NewAddReleaseGroupsToTeam200ResponseWithDefaults instantiates a new AddReleaseGroupsToTeam200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddReleaseGroupsToTeam200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddReleaseGroupsToTeam200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddReleaseGroupsToTeam200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddReleaseGroupsToTeam200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReleaseGroups

`func (o *AddReleaseGroupsToTeam200Response) GetReleaseGroups() []int32`

GetReleaseGroups returns the ReleaseGroups field if non-nil, zero value otherwise.

### GetReleaseGroupsOk

`func (o *AddReleaseGroupsToTeam200Response) GetReleaseGroupsOk() (*[]int32, bool)`

GetReleaseGroupsOk returns a tuple with the ReleaseGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseGroups

`func (o *AddReleaseGroupsToTeam200Response) SetReleaseGroups(v []int32)`

SetReleaseGroups sets ReleaseGroups field to given value.

### HasReleaseGroups

`func (o *AddReleaseGroupsToTeam200Response) HasReleaseGroups() bool`

HasReleaseGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


