# GetReleaseGroupReleaseScans200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScannedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectGroupReleaseId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RevisionScans** | Pointer to [**[]GetReleaseGroupReleaseScans200ResponseInnerAllOfRevisionScansInner**](GetReleaseGroupReleaseScans200ResponseInnerAllOfRevisionScansInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseScans200ResponseInner

`func NewGetReleaseGroupReleaseScans200ResponseInner() *GetReleaseGroupReleaseScans200ResponseInner`

NewGetReleaseGroupReleaseScans200ResponseInner instantiates a new GetReleaseGroupReleaseScans200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseScans200ResponseInnerWithDefaults

`func NewGetReleaseGroupReleaseScans200ResponseInnerWithDefaults() *GetReleaseGroupReleaseScans200ResponseInner`

NewGetReleaseGroupReleaseScans200ResponseInnerWithDefaults instantiates a new GetReleaseGroupReleaseScans200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetProjectGroupReleaseId() int32`

GetProjectGroupReleaseId returns the ProjectGroupReleaseId field if non-nil, zero value otherwise.

### GetProjectGroupReleaseIdOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetProjectGroupReleaseIdOk() (*int32, bool)`

GetProjectGroupReleaseIdOk returns a tuple with the ProjectGroupReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetProjectGroupReleaseId(v int32)`

SetProjectGroupReleaseId sets ProjectGroupReleaseId field to given value.

### HasProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasProjectGroupReleaseId() bool`

HasProjectGroupReleaseId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRevisionScans

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetRevisionScans() []GetReleaseGroupReleaseScans200ResponseInnerAllOfRevisionScansInner`

GetRevisionScans returns the RevisionScans field if non-nil, zero value otherwise.

### GetRevisionScansOk

`func (o *GetReleaseGroupReleaseScans200ResponseInner) GetRevisionScansOk() (*[]GetReleaseGroupReleaseScans200ResponseInnerAllOfRevisionScansInner, bool)`

GetRevisionScansOk returns a tuple with the RevisionScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionScans

`func (o *GetReleaseGroupReleaseScans200ResponseInner) SetRevisionScans(v []GetReleaseGroupReleaseScans200ResponseInnerAllOfRevisionScansInner)`

SetRevisionScans sets RevisionScans field to given value.

### HasRevisionScans

`func (o *GetReleaseGroupReleaseScans200ResponseInner) HasRevisionScans() bool`

HasRevisionScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


