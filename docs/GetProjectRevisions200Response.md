# GetProjectRevisions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to [**map[string][]GetProjectRevisions200ResponseBranchValueInner**](array.md) | map of branches to their Revisions | [optional] 
**Tag** | Pointer to [**map[string][]GetProjectRevisions200ResponseBranchValueInner**](array.md) | map of tags to their Revisions | [optional] 

## Methods

### NewGetProjectRevisions200Response

`func NewGetProjectRevisions200Response() *GetProjectRevisions200Response`

NewGetProjectRevisions200Response instantiates a new GetProjectRevisions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectRevisions200ResponseWithDefaults

`func NewGetProjectRevisions200ResponseWithDefaults() *GetProjectRevisions200Response`

NewGetProjectRevisions200ResponseWithDefaults instantiates a new GetProjectRevisions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *GetProjectRevisions200Response) GetBranch() map[string][]GetProjectRevisions200ResponseBranchValueInner`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GetProjectRevisions200Response) GetBranchOk() (*map[string][]GetProjectRevisions200ResponseBranchValueInner, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GetProjectRevisions200Response) SetBranch(v map[string][]GetProjectRevisions200ResponseBranchValueInner)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GetProjectRevisions200Response) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetTag

`func (o *GetProjectRevisions200Response) GetTag() map[string][]GetProjectRevisions200ResponseBranchValueInner`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GetProjectRevisions200Response) GetTagOk() (*map[string][]GetProjectRevisions200ResponseBranchValueInner, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GetProjectRevisions200Response) SetTag(v map[string][]GetProjectRevisions200ResponseBranchValueInner)`

SetTag sets Tag field to given value.

### HasTag

`func (o *GetProjectRevisions200Response) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


