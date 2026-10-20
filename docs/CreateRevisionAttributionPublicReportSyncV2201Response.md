# CreateRevisionAttributionPublicReportSyncV2201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The report ID | [optional] 
**Uuid** | Pointer to **string** | The public UUID used in the report&#39;s URL | [optional] 
**Url** | Pointer to **string** | The S3 URL of the generated report | [optional] 
**Public** | Pointer to **bool** | Whether the report is publicly accessible | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**RevisionId** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateRevisionAttributionPublicReportSyncV2201Response

`func NewCreateRevisionAttributionPublicReportSyncV2201Response() *CreateRevisionAttributionPublicReportSyncV2201Response`

NewCreateRevisionAttributionPublicReportSyncV2201Response instantiates a new CreateRevisionAttributionPublicReportSyncV2201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRevisionAttributionPublicReportSyncV2201ResponseWithDefaults

`func NewCreateRevisionAttributionPublicReportSyncV2201ResponseWithDefaults() *CreateRevisionAttributionPublicReportSyncV2201Response`

NewCreateRevisionAttributionPublicReportSyncV2201ResponseWithDefaults instantiates a new CreateRevisionAttributionPublicReportSyncV2201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetUrl

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPublic

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRevisionId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *CreateRevisionAttributionPublicReportSyncV2201Response) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


