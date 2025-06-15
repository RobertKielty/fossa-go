# GetAllReleaseGroupTeams200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** |  | [optional] 
**Teams** | Pointer to [**[]GetAllReleaseGroupTeams200ResponseTeamsInner**](GetAllReleaseGroupTeams200ResponseTeamsInner.md) |  | [optional] 

## Methods

### NewGetAllReleaseGroupTeams200Response

`func NewGetAllReleaseGroupTeams200Response() *GetAllReleaseGroupTeams200Response`

NewGetAllReleaseGroupTeams200Response instantiates a new GetAllReleaseGroupTeams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllReleaseGroupTeams200ResponseWithDefaults

`func NewGetAllReleaseGroupTeams200ResponseWithDefaults() *GetAllReleaseGroupTeams200Response`

NewGetAllReleaseGroupTeams200ResponseWithDefaults instantiates a new GetAllReleaseGroupTeams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *GetAllReleaseGroupTeams200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetAllReleaseGroupTeams200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetAllReleaseGroupTeams200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetAllReleaseGroupTeams200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTeams

`func (o *GetAllReleaseGroupTeams200Response) GetTeams() []GetAllReleaseGroupTeams200ResponseTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetAllReleaseGroupTeams200Response) GetTeamsOk() (*[]GetAllReleaseGroupTeams200ResponseTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetAllReleaseGroupTeams200Response) SetTeams(v []GetAllReleaseGroupTeams200ResponseTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *GetAllReleaseGroupTeams200Response) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


