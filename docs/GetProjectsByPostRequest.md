# GetProjectsByPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | Pointer to **string** | The category to order the results by and sort direction. | [optional] 
**Page** | Pointer to **int32** | The specific page of data to return. | [optional] [default to 1]
**Count** | Pointer to **int32** | The number of items to return in each page of results. | [optional] [default to 20]
**Title** | Pointer to **string** | Filter by project name. | [optional] 
**Type** | Pointer to **[]string** | Filter by project type. | [optional] 
**IsPublic** | Pointer to **bool** | Filter by project being public or private. | [optional] 
**Labels** | Pointer to **[]string** | Filter by project labels. | [optional] 
**TeamId** | Pointer to [**[]GetIssueCountsTeamIdParameterInner**](GetIssueCountsTeamIdParameterInner.md) | Filter by one or more team IDs. Providing \&quot;null\&quot; will return all unassigned projects. | [optional] 
**LatestScan** | Pointer to **int32** | Filter by last policy scan within N days. | [optional] 
**LastRevisionWithin** | Pointer to **int32** | Filter by last revision analyzed within N days. | [optional] 
**Locators** | Pointer to **[]string** | Filter by project locators (exact match). | [optional] 
**Url** | Pointer to **string** | Filter by a project&#39;s URL. | [optional] 
**IncludeSharedProjects** | Pointer to **bool** | Include shared projects. | [optional] 
**OnlyIncludeSharedProjects** | Pointer to **bool** | Only show projects that have been shared with other organizations. | [optional] 
**Inventory** | Pointer to **[]string** | Filter by by additional inventory types. | [optional] 

## Methods

### NewGetProjectsByPostRequest

`func NewGetProjectsByPostRequest() *GetProjectsByPostRequest`

NewGetProjectsByPostRequest instantiates a new GetProjectsByPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectsByPostRequestWithDefaults

`func NewGetProjectsByPostRequestWithDefaults() *GetProjectsByPostRequest`

NewGetProjectsByPostRequestWithDefaults instantiates a new GetProjectsByPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *GetProjectsByPostRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GetProjectsByPostRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GetProjectsByPostRequest) SetSort(v string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GetProjectsByPostRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetPage

`func (o *GetProjectsByPostRequest) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetProjectsByPostRequest) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetProjectsByPostRequest) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetProjectsByPostRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetCount

`func (o *GetProjectsByPostRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetProjectsByPostRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetProjectsByPostRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetProjectsByPostRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTitle

`func (o *GetProjectsByPostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetProjectsByPostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetProjectsByPostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetProjectsByPostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *GetProjectsByPostRequest) GetType() []string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProjectsByPostRequest) GetTypeOk() (*[]string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProjectsByPostRequest) SetType(v []string)`

SetType sets Type field to given value.

### HasType

`func (o *GetProjectsByPostRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetProjectsByPostRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetProjectsByPostRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetProjectsByPostRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetProjectsByPostRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLabels

`func (o *GetProjectsByPostRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetProjectsByPostRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetProjectsByPostRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetProjectsByPostRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTeamId

`func (o *GetProjectsByPostRequest) GetTeamId() []GetIssueCountsTeamIdParameterInner`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *GetProjectsByPostRequest) GetTeamIdOk() (*[]GetIssueCountsTeamIdParameterInner, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *GetProjectsByPostRequest) SetTeamId(v []GetIssueCountsTeamIdParameterInner)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *GetProjectsByPostRequest) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetLatestScan

`func (o *GetProjectsByPostRequest) GetLatestScan() int32`

GetLatestScan returns the LatestScan field if non-nil, zero value otherwise.

### GetLatestScanOk

`func (o *GetProjectsByPostRequest) GetLatestScanOk() (*int32, bool)`

GetLatestScanOk returns a tuple with the LatestScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestScan

`func (o *GetProjectsByPostRequest) SetLatestScan(v int32)`

SetLatestScan sets LatestScan field to given value.

### HasLatestScan

`func (o *GetProjectsByPostRequest) HasLatestScan() bool`

HasLatestScan returns a boolean if a field has been set.

### GetLastRevisionWithin

`func (o *GetProjectsByPostRequest) GetLastRevisionWithin() int32`

GetLastRevisionWithin returns the LastRevisionWithin field if non-nil, zero value otherwise.

### GetLastRevisionWithinOk

`func (o *GetProjectsByPostRequest) GetLastRevisionWithinOk() (*int32, bool)`

GetLastRevisionWithinOk returns a tuple with the LastRevisionWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRevisionWithin

`func (o *GetProjectsByPostRequest) SetLastRevisionWithin(v int32)`

SetLastRevisionWithin sets LastRevisionWithin field to given value.

### HasLastRevisionWithin

`func (o *GetProjectsByPostRequest) HasLastRevisionWithin() bool`

HasLastRevisionWithin returns a boolean if a field has been set.

### GetLocators

`func (o *GetProjectsByPostRequest) GetLocators() []string`

GetLocators returns the Locators field if non-nil, zero value otherwise.

### GetLocatorsOk

`func (o *GetProjectsByPostRequest) GetLocatorsOk() (*[]string, bool)`

GetLocatorsOk returns a tuple with the Locators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocators

`func (o *GetProjectsByPostRequest) SetLocators(v []string)`

SetLocators sets Locators field to given value.

### HasLocators

`func (o *GetProjectsByPostRequest) HasLocators() bool`

HasLocators returns a boolean if a field has been set.

### GetUrl

`func (o *GetProjectsByPostRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetProjectsByPostRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetProjectsByPostRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetProjectsByPostRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIncludeSharedProjects

`func (o *GetProjectsByPostRequest) GetIncludeSharedProjects() bool`

GetIncludeSharedProjects returns the IncludeSharedProjects field if non-nil, zero value otherwise.

### GetIncludeSharedProjectsOk

`func (o *GetProjectsByPostRequest) GetIncludeSharedProjectsOk() (*bool, bool)`

GetIncludeSharedProjectsOk returns a tuple with the IncludeSharedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSharedProjects

`func (o *GetProjectsByPostRequest) SetIncludeSharedProjects(v bool)`

SetIncludeSharedProjects sets IncludeSharedProjects field to given value.

### HasIncludeSharedProjects

`func (o *GetProjectsByPostRequest) HasIncludeSharedProjects() bool`

HasIncludeSharedProjects returns a boolean if a field has been set.

### GetOnlyIncludeSharedProjects

`func (o *GetProjectsByPostRequest) GetOnlyIncludeSharedProjects() bool`

GetOnlyIncludeSharedProjects returns the OnlyIncludeSharedProjects field if non-nil, zero value otherwise.

### GetOnlyIncludeSharedProjectsOk

`func (o *GetProjectsByPostRequest) GetOnlyIncludeSharedProjectsOk() (*bool, bool)`

GetOnlyIncludeSharedProjectsOk returns a tuple with the OnlyIncludeSharedProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIncludeSharedProjects

`func (o *GetProjectsByPostRequest) SetOnlyIncludeSharedProjects(v bool)`

SetOnlyIncludeSharedProjects sets OnlyIncludeSharedProjects field to given value.

### HasOnlyIncludeSharedProjects

`func (o *GetProjectsByPostRequest) HasOnlyIncludeSharedProjects() bool`

HasOnlyIncludeSharedProjects returns a boolean if a field has been set.

### GetInventory

`func (o *GetProjectsByPostRequest) GetInventory() []string`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *GetProjectsByPostRequest) GetInventoryOk() (*[]string, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *GetProjectsByPostRequest) SetInventory(v []string)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *GetProjectsByPostRequest) HasInventory() bool`

HasInventory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


