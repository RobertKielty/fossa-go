# AddTeamsToTeamGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamIds** | **[]int32** | Array of team IDs to add to the team group | 

## Methods

### NewAddTeamsToTeamGroupRequest

`func NewAddTeamsToTeamGroupRequest(teamIds []int32, ) *AddTeamsToTeamGroupRequest`

NewAddTeamsToTeamGroupRequest instantiates a new AddTeamsToTeamGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTeamsToTeamGroupRequestWithDefaults

`func NewAddTeamsToTeamGroupRequestWithDefaults() *AddTeamsToTeamGroupRequest`

NewAddTeamsToTeamGroupRequestWithDefaults instantiates a new AddTeamsToTeamGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamIds

`func (o *AddTeamsToTeamGroupRequest) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *AddTeamsToTeamGroupRequest) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *AddTeamsToTeamGroupRequest) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


