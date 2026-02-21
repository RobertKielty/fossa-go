# GetAllReportOptions200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier for the report option | 
**Name** | **string** | Name of the saved report option | 
**OrganizationId** | **int32** | ID of the organization that owns this report option | 
**Options** | [**GetAllReportOptions200ResponseResultsInnerOptions**](GetAllReportOptions200ResponseResultsInnerOptions.md) |  | 
**CreatedAt** | **time.Time** | Timestamp when the report option was created | 
**UpdatedAt** | **time.Time** | Timestamp when the report option was last updated | 

## Methods

### NewGetAllReportOptions200ResponseResultsInner

`func NewGetAllReportOptions200ResponseResultsInner(id int32, name string, organizationId int32, options GetAllReportOptions200ResponseResultsInnerOptions, createdAt time.Time, updatedAt time.Time, ) *GetAllReportOptions200ResponseResultsInner`

NewGetAllReportOptions200ResponseResultsInner instantiates a new GetAllReportOptions200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllReportOptions200ResponseResultsInnerWithDefaults

`func NewGetAllReportOptions200ResponseResultsInnerWithDefaults() *GetAllReportOptions200ResponseResultsInner`

NewGetAllReportOptions200ResponseResultsInnerWithDefaults instantiates a new GetAllReportOptions200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAllReportOptions200ResponseResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAllReportOptions200ResponseResultsInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *GetAllReportOptions200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllReportOptions200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *GetAllReportOptions200ResponseResultsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetAllReportOptions200ResponseResultsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOptions

`func (o *GetAllReportOptions200ResponseResultsInner) GetOptions() GetAllReportOptions200ResponseResultsInnerOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetOptionsOk() (*GetAllReportOptions200ResponseResultsInnerOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetAllReportOptions200ResponseResultsInner) SetOptions(v GetAllReportOptions200ResponseResultsInnerOptions)`

SetOptions sets Options field to given value.


### GetCreatedAt

`func (o *GetAllReportOptions200ResponseResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAllReportOptions200ResponseResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetAllReportOptions200ResponseResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAllReportOptions200ResponseResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAllReportOptions200ResponseResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


