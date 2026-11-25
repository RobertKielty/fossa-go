# CreateReportOption201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Unique identifier for the report option | 
**Name** | **string** | Name of the saved report option | 
**OrganizationId** | **int32** | ID of the organization that owns this report option | 
**Options** | [**CreateReportOptionRequestOptions**](CreateReportOptionRequestOptions.md) |  | 
**CreatedAt** | **time.Time** | Timestamp when the report option was created | 
**UpdatedAt** | **time.Time** | Timestamp when the report option was last updated | 

## Methods

### NewCreateReportOption201Response

`func NewCreateReportOption201Response(id int32, name string, organizationId int32, options CreateReportOptionRequestOptions, createdAt time.Time, updatedAt time.Time, ) *CreateReportOption201Response`

NewCreateReportOption201Response instantiates a new CreateReportOption201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportOption201ResponseWithDefaults

`func NewCreateReportOption201ResponseWithDefaults() *CreateReportOption201Response`

NewCreateReportOption201ResponseWithDefaults instantiates a new CreateReportOption201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReportOption201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReportOption201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReportOption201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CreateReportOption201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReportOption201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReportOption201Response) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *CreateReportOption201Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateReportOption201Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateReportOption201Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetOptions

`func (o *CreateReportOption201Response) GetOptions() CreateReportOptionRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateReportOption201Response) GetOptionsOk() (*CreateReportOptionRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateReportOption201Response) SetOptions(v CreateReportOptionRequestOptions)`

SetOptions sets Options field to given value.


### GetCreatedAt

`func (o *CreateReportOption201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateReportOption201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateReportOption201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateReportOption201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateReportOption201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateReportOption201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


