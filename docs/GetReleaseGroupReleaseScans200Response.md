# GetReleaseGroupReleaseScans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScannedAt** | Pointer to **time.Time** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectGroupReleaseId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RevisionScans** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseScans200Response

`func NewGetReleaseGroupReleaseScans200Response() *GetReleaseGroupReleaseScans200Response`

NewGetReleaseGroupReleaseScans200Response instantiates a new GetReleaseGroupReleaseScans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseScans200ResponseWithDefaults

`func NewGetReleaseGroupReleaseScans200ResponseWithDefaults() *GetReleaseGroupReleaseScans200Response`

NewGetReleaseGroupReleaseScans200ResponseWithDefaults instantiates a new GetReleaseGroupReleaseScans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupReleaseScans200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupReleaseScans200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupReleaseScans200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetReleaseGroupReleaseScans200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetReleaseGroupReleaseScans200Response) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetReleaseGroupReleaseScans200Response) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetReleaseGroupReleaseScans200Response) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetReleaseGroupReleaseScans200Response) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetReleaseGroupReleaseScans200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetReleaseGroupReleaseScans200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetReleaseGroupReleaseScans200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetReleaseGroupReleaseScans200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200Response) GetProjectGroupReleaseId() int32`

GetProjectGroupReleaseId returns the ProjectGroupReleaseId field if non-nil, zero value otherwise.

### GetProjectGroupReleaseIdOk

`func (o *GetReleaseGroupReleaseScans200Response) GetProjectGroupReleaseIdOk() (*int32, bool)`

GetProjectGroupReleaseIdOk returns a tuple with the ProjectGroupReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200Response) SetProjectGroupReleaseId(v int32)`

SetProjectGroupReleaseId sets ProjectGroupReleaseId field to given value.

### HasProjectGroupReleaseId

`func (o *GetReleaseGroupReleaseScans200Response) HasProjectGroupReleaseId() bool`

HasProjectGroupReleaseId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetReleaseGroupReleaseScans200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReleaseGroupReleaseScans200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReleaseGroupReleaseScans200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetReleaseGroupReleaseScans200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReleaseGroupReleaseScans200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReleaseGroupReleaseScans200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReleaseGroupReleaseScans200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReleaseGroupReleaseScans200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRevisionScans

`func (o *GetReleaseGroupReleaseScans200Response) GetRevisionScans() []interface{}`

GetRevisionScans returns the RevisionScans field if non-nil, zero value otherwise.

### GetRevisionScansOk

`func (o *GetReleaseGroupReleaseScans200Response) GetRevisionScansOk() (*[]interface{}, bool)`

GetRevisionScansOk returns a tuple with the RevisionScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionScans

`func (o *GetReleaseGroupReleaseScans200Response) SetRevisionScans(v []interface{})`

SetRevisionScans sets RevisionScans field to given value.

### HasRevisionScans

`func (o *GetReleaseGroupReleaseScans200Response) HasRevisionScans() bool`

HasRevisionScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


