# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locator** | Pointer to **string** | Text ID that uniquely identifies a project | [optional] 
**Title** | Pointer to **string** | Name of the project | [optional] 
**Labels** | Pointer to [**[]GetOrganizationLabels200ResponseLabelsInner**](GetOrganizationLabels200ResponseLabelsInner.md) |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocator

`func (o *Project) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *Project) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *Project) SetLocator(v string)`

SetLocator sets Locator field to given value.

### HasLocator

`func (o *Project) HasLocator() bool`

HasLocator returns a boolean if a field has been set.

### GetTitle

`func (o *Project) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Project) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Project) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Project) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLabels

`func (o *Project) GetLabels() []GetOrganizationLabels200ResponseLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Project) GetLabelsOk() (*[]GetOrganizationLabels200ResponseLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Project) SetLabels(v []GetOrganizationLabels200ResponseLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Project) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


