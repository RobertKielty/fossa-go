# GetPackages200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locator** | **string** | The package locator (without revision / version data). | 
**Title** | **string** | The human-readable package title. | 
**NumberOfProjects** | **int32** | The number of your projects that depend on this package. | 
**BlockedVersionCount** | **int32** | The number of versions of this package that are blocked by your organization. | 

## Methods

### NewGetPackages200ResponseDataInner

`func NewGetPackages200ResponseDataInner(locator string, title string, numberOfProjects int32, blockedVersionCount int32, ) *GetPackages200ResponseDataInner`

NewGetPackages200ResponseDataInner instantiates a new GetPackages200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPackages200ResponseDataInnerWithDefaults

`func NewGetPackages200ResponseDataInnerWithDefaults() *GetPackages200ResponseDataInner`

NewGetPackages200ResponseDataInnerWithDefaults instantiates a new GetPackages200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocator

`func (o *GetPackages200ResponseDataInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetPackages200ResponseDataInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetPackages200ResponseDataInner) SetLocator(v string)`

SetLocator sets Locator field to given value.


### GetTitle

`func (o *GetPackages200ResponseDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetPackages200ResponseDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetPackages200ResponseDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetNumberOfProjects

`func (o *GetPackages200ResponseDataInner) GetNumberOfProjects() int32`

GetNumberOfProjects returns the NumberOfProjects field if non-nil, zero value otherwise.

### GetNumberOfProjectsOk

`func (o *GetPackages200ResponseDataInner) GetNumberOfProjectsOk() (*int32, bool)`

GetNumberOfProjectsOk returns a tuple with the NumberOfProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfProjects

`func (o *GetPackages200ResponseDataInner) SetNumberOfProjects(v int32)`

SetNumberOfProjects sets NumberOfProjects field to given value.


### GetBlockedVersionCount

`func (o *GetPackages200ResponseDataInner) GetBlockedVersionCount() int32`

GetBlockedVersionCount returns the BlockedVersionCount field if non-nil, zero value otherwise.

### GetBlockedVersionCountOk

`func (o *GetPackages200ResponseDataInner) GetBlockedVersionCountOk() (*int32, bool)`

GetBlockedVersionCountOk returns a tuple with the BlockedVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVersionCount

`func (o *GetPackages200ResponseDataInner) SetBlockedVersionCount(v int32)`

SetBlockedVersionCount sets BlockedVersionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


