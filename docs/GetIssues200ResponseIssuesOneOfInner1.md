# GetIssues200ResponseIssuesOneOfInner1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**GetIssue200ResponseOneOfAllOfSource**](GetIssue200ResponseOneOfAllOfSource.md) |  | [optional] 
**Depths** | Pointer to [**GetIssue200ResponseOneOfAllOfDepths**](GetIssue200ResponseOneOfAllOfDepths.md) |  | [optional] 
**ContainerLayers** | Pointer to [**GetIssue200ResponseOneOfAllOfContainerLayers**](GetIssue200ResponseOneOfAllOfContainerLayers.md) |  | [optional] 
**Statuses** | Pointer to [**GetIssueStatuses200ResponseIssues**](GetIssueStatuses200ResponseIssues.md) |  | [optional] 
**Projects** | Pointer to [**[]GetIssue200ResponseOneOfAllOfProjectsInner**](GetIssue200ResponseOneOfAllOfProjectsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewGetIssues200ResponseIssuesOneOfInner1

`func NewGetIssues200ResponseIssuesOneOfInner1() *GetIssues200ResponseIssuesOneOfInner1`

NewGetIssues200ResponseIssuesOneOfInner1 instantiates a new GetIssues200ResponseIssuesOneOfInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssues200ResponseIssuesOneOfInner1WithDefaults

`func NewGetIssues200ResponseIssuesOneOfInner1WithDefaults() *GetIssues200ResponseIssuesOneOfInner1`

NewGetIssues200ResponseIssuesOneOfInner1WithDefaults instantiates a new GetIssues200ResponseIssuesOneOfInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetSource() GetIssue200ResponseOneOfAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetSource(v GetIssue200ResponseOneOfAllOfSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepths

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetDepths() GetIssue200ResponseOneOfAllOfDepths`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetDepths(v GetIssue200ResponseOneOfAllOfDepths)`

SetDepths sets Depths field to given value.

### HasDepths

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasDepths() bool`

HasDepths returns a boolean if a field has been set.

### GetContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetStatuses() GetIssueStatuses200ResponseIssues`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetStatuses(v GetIssueStatuses200ResponseIssues)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProjects

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetProjectsOk() (*[]GetIssue200ResponseOneOfAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLicense

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetIssues200ResponseIssuesOneOfInner1) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetIssues200ResponseIssuesOneOfInner1) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *GetIssues200ResponseIssuesOneOfInner1) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


