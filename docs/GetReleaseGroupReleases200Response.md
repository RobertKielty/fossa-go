# GetReleaseGroupReleases200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]GetReleaseGroupById200ResponseAllOfReleasesInner**](GetReleaseGroupById200ResponseAllOfReleasesInner.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetReleaseGroupReleases200Response

`func NewGetReleaseGroupReleases200Response() *GetReleaseGroupReleases200Response`

NewGetReleaseGroupReleases200Response instantiates a new GetReleaseGroupReleases200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleases200ResponseWithDefaults

`func NewGetReleaseGroupReleases200ResponseWithDefaults() *GetReleaseGroupReleases200Response`

NewGetReleaseGroupReleases200ResponseWithDefaults instantiates a new GetReleaseGroupReleases200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *GetReleaseGroupReleases200Response) GetReleases() []GetReleaseGroupById200ResponseAllOfReleasesInner`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *GetReleaseGroupReleases200Response) GetReleasesOk() (*[]GetReleaseGroupById200ResponseAllOfReleasesInner, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *GetReleaseGroupReleases200Response) SetReleases(v []GetReleaseGroupById200ResponseAllOfReleasesInner)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *GetReleaseGroupReleases200Response) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetTotal

`func (o *GetReleaseGroupReleases200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetReleaseGroupReleases200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetReleaseGroupReleases200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetReleaseGroupReleases200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


