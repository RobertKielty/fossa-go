# RejectSnippetsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The path to filter snippets by | 
**Ids** | Pointer to **[]string** | Filter by specific snippet IDs | [optional] 
**PackageIds** | Pointer to **[]string** | Filter by specific snippet package IDs | [optional] 
**Search** | Pointer to **string** | Search term for filtering snippets | [optional] 
**RejectionStatus** | Pointer to **[]string** | Filter by rejection status | [optional] 
**PackageLabels** | Pointer to **[]string** | Filter by package labels | [optional] 
**VendoredMatch** | Pointer to **[]string** | Filter by vendored/converted match status | [optional] 

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

### GetPackageIds

`func (o *RejectSnippetsRequest) GetPackageIds() []string`

GetPackageIds returns the PackageIds field if non-nil, zero value otherwise.

### GetPackageIdsOk

`func (o *RejectSnippetsRequest) GetPackageIdsOk() (*[]string, bool)`

GetPackageIdsOk returns a tuple with the PackageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageIds

`func (o *RejectSnippetsRequest) SetPackageIds(v []string)`

SetPackageIds sets PackageIds field to given value.

### HasPackageIds

`func (o *RejectSnippetsRequest) HasPackageIds() bool`

HasPackageIds returns a boolean if a field has been set.

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

### GetRejectionStatus

`func (o *RejectSnippetsRequest) GetRejectionStatus() []string`

GetRejectionStatus returns the RejectionStatus field if non-nil, zero value otherwise.

### GetRejectionStatusOk

`func (o *RejectSnippetsRequest) GetRejectionStatusOk() (*[]string, bool)`

GetRejectionStatusOk returns a tuple with the RejectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionStatus

`func (o *RejectSnippetsRequest) SetRejectionStatus(v []string)`

SetRejectionStatus sets RejectionStatus field to given value.

### HasRejectionStatus

`func (o *RejectSnippetsRequest) HasRejectionStatus() bool`

HasRejectionStatus returns a boolean if a field has been set.

### GetPackageLabels

`func (o *RejectSnippetsRequest) GetPackageLabels() []string`

GetPackageLabels returns the PackageLabels field if non-nil, zero value otherwise.

### GetPackageLabelsOk

`func (o *RejectSnippetsRequest) GetPackageLabelsOk() (*[]string, bool)`

GetPackageLabelsOk returns a tuple with the PackageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabels

`func (o *RejectSnippetsRequest) SetPackageLabels(v []string)`

SetPackageLabels sets PackageLabels field to given value.

### HasPackageLabels

`func (o *RejectSnippetsRequest) HasPackageLabels() bool`

HasPackageLabels returns a boolean if a field has been set.

### GetVendoredMatch

`func (o *RejectSnippetsRequest) GetVendoredMatch() []string`

GetVendoredMatch returns the VendoredMatch field if non-nil, zero value otherwise.

### GetVendoredMatchOk

`func (o *RejectSnippetsRequest) GetVendoredMatchOk() (*[]string, bool)`

GetVendoredMatchOk returns a tuple with the VendoredMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendoredMatch

`func (o *RejectSnippetsRequest) SetVendoredMatch(v []string)`

SetVendoredMatch sets VendoredMatch field to given value.

### HasVendoredMatch

`func (o *RejectSnippetsRequest) HasVendoredMatch() bool`

HasVendoredMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


