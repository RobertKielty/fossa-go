# RejectSnippetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path to filter snippets by | 
**Ids** | Pointer to **[]string** | Filter by specific snippet IDs | [optional] 
**Search** | Pointer to **string** | Search term for filtering snippets | [optional] 
**Confidence** | Pointer to **[]string** | Filter by confidence levels | [optional] 

## Methods

### NewRejectSnippetsRequest

`func NewRejectSnippetsRequest(path string, ) *RejectSnippetsRequest`

NewRejectSnippetsRequest instantiates a new RejectSnippetsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectSnippetsRequestWithDefaults

`func NewRejectSnippetsRequestWithDefaults() *RejectSnippetsRequest`

NewRejectSnippetsRequestWithDefaults instantiates a new RejectSnippetsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *RejectSnippetsRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RejectSnippetsRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RejectSnippetsRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetIds

`func (o *RejectSnippetsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RejectSnippetsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RejectSnippetsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *RejectSnippetsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSearch

`func (o *RejectSnippetsRequest) GetSearch() string`

GetSearch returns the Search field if non-nil, zero value otherwise.

### GetSearchOk

`func (o *RejectSnippetsRequest) GetSearchOk() (*string, bool)`

GetSearchOk returns a tuple with the Search field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearch

`func (o *RejectSnippetsRequest) SetSearch(v string)`

SetSearch sets Search field to given value.

### HasSearch

`func (o *RejectSnippetsRequest) HasSearch() bool`

HasSearch returns a boolean if a field has been set.

### GetConfidence

`func (o *RejectSnippetsRequest) GetConfidence() []string`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *RejectSnippetsRequest) GetConfidenceOk() (*[]string, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *RejectSnippetsRequest) SetConfidence(v []string)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *RejectSnippetsRequest) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


