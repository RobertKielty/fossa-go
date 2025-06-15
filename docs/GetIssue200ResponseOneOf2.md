# GetIssue200ResponseOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Source** | Pointer to [**GetIssue200ResponseOneOfAllOfSource1**](GetIssue200ResponseOneOfAllOfSource1.md) |  | [optional] 
**Depths** | Pointer to [**GetIssue200ResponseOneOfAllOfDepths**](GetIssue200ResponseOneOfAllOfDepths.md) |  | [optional] 
**ContainerLayers** | Pointer to [**GetIssue200ResponseOneOfAllOfContainerLayers**](GetIssue200ResponseOneOfAllOfContainerLayers.md) |  | [optional] 
**Statuses** | Pointer to [**GetIssueStatuses200ResponseIssues**](GetIssueStatuses200ResponseIssues.md) |  | [optional] 
**Projects** | Pointer to [**[]GetIssue200ResponseOneOfAllOfProjectsInner**](GetIssue200ResponseOneOfAllOfProjectsInner.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**QualityRule** | Pointer to [**GetIssue200ResponseOneOf2AllOfQualityRule**](GetIssue200ResponseOneOf2AllOfQualityRule.md) |  | [optional] 

## Methods

### NewGetIssue200ResponseOneOf2

`func NewGetIssue200ResponseOneOf2() *GetIssue200ResponseOneOf2`

NewGetIssue200ResponseOneOf2 instantiates a new GetIssue200ResponseOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseOneOf2WithDefaults

`func NewGetIssue200ResponseOneOf2WithDefaults() *GetIssue200ResponseOneOf2`

NewGetIssue200ResponseOneOf2WithDefaults instantiates a new GetIssue200ResponseOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIssue200ResponseOneOf2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIssue200ResponseOneOf2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIssue200ResponseOneOf2) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetIssue200ResponseOneOf2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssue200ResponseOneOf2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssue200ResponseOneOf2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssue200ResponseOneOf2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssue200ResponseOneOf2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetSource

`func (o *GetIssue200ResponseOneOf2) GetSource() GetIssue200ResponseOneOfAllOfSource1`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetIssue200ResponseOneOf2) GetSourceOk() (*GetIssue200ResponseOneOfAllOfSource1, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetIssue200ResponseOneOf2) SetSource(v GetIssue200ResponseOneOfAllOfSource1)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetIssue200ResponseOneOf2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDepths

`func (o *GetIssue200ResponseOneOf2) GetDepths() GetIssue200ResponseOneOfAllOfDepths`

GetDepths returns the Depths field if non-nil, zero value otherwise.

### GetDepthsOk

`func (o *GetIssue200ResponseOneOf2) GetDepthsOk() (*GetIssue200ResponseOneOfAllOfDepths, bool)`

GetDepthsOk returns a tuple with the Depths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepths

`func (o *GetIssue200ResponseOneOf2) SetDepths(v GetIssue200ResponseOneOfAllOfDepths)`

SetDepths sets Depths field to given value.

### HasDepths

`func (o *GetIssue200ResponseOneOf2) HasDepths() bool`

HasDepths returns a boolean if a field has been set.

### GetContainerLayers

`func (o *GetIssue200ResponseOneOf2) GetContainerLayers() GetIssue200ResponseOneOfAllOfContainerLayers`

GetContainerLayers returns the ContainerLayers field if non-nil, zero value otherwise.

### GetContainerLayersOk

`func (o *GetIssue200ResponseOneOf2) GetContainerLayersOk() (*GetIssue200ResponseOneOfAllOfContainerLayers, bool)`

GetContainerLayersOk returns a tuple with the ContainerLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerLayers

`func (o *GetIssue200ResponseOneOf2) SetContainerLayers(v GetIssue200ResponseOneOfAllOfContainerLayers)`

SetContainerLayers sets ContainerLayers field to given value.

### HasContainerLayers

`func (o *GetIssue200ResponseOneOf2) HasContainerLayers() bool`

HasContainerLayers returns a boolean if a field has been set.

### GetStatuses

`func (o *GetIssue200ResponseOneOf2) GetStatuses() GetIssueStatuses200ResponseIssues`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *GetIssue200ResponseOneOf2) GetStatusesOk() (*GetIssueStatuses200ResponseIssues, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *GetIssue200ResponseOneOf2) SetStatuses(v GetIssueStatuses200ResponseIssues)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *GetIssue200ResponseOneOf2) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetProjects

`func (o *GetIssue200ResponseOneOf2) GetProjects() []GetIssue200ResponseOneOfAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetIssue200ResponseOneOf2) GetProjectsOk() (*[]GetIssue200ResponseOneOfAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetIssue200ResponseOneOf2) SetProjects(v []GetIssue200ResponseOneOfAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetIssue200ResponseOneOf2) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetType

`func (o *GetIssue200ResponseOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssue200ResponseOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssue200ResponseOneOf2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssue200ResponseOneOf2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQualityRule

`func (o *GetIssue200ResponseOneOf2) GetQualityRule() GetIssue200ResponseOneOf2AllOfQualityRule`

GetQualityRule returns the QualityRule field if non-nil, zero value otherwise.

### GetQualityRuleOk

`func (o *GetIssue200ResponseOneOf2) GetQualityRuleOk() (*GetIssue200ResponseOneOf2AllOfQualityRule, bool)`

GetQualityRuleOk returns a tuple with the QualityRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRule

`func (o *GetIssue200ResponseOneOf2) SetQualityRule(v GetIssue200ResponseOneOf2AllOfQualityRule)`

SetQualityRule sets QualityRule field to given value.

### HasQualityRule

`func (o *GetIssue200ResponseOneOf2) HasQualityRule() bool`

HasQualityRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


