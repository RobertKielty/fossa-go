# GetCustomLicenses200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyProjectLocator** | **string** | The locator of the dependency project that this custom license applies to | 
**DependencyRevisionLocator** | **string** | The specific revision locator of the dependency | 
**ProjectLocator** | **string** | The locator of the user project that contains this dependency | 
**ProjectRevision** | **string** | The revision locator of the user project | 
**Title** | **string** | The title/name of the dependency project | 
**Description** | **string** | Description of the dependency project | 
**Url** | **string** | URL for the dependency project | 
**LicenseTitle** | **string** | The custom license title/name | 
**LicenseId** | **string** | The license identifier, always &#39;custom-license&#39; for custom licenses | 
**Copyright** | **string** | Copyright notice for this license | 
**Text** | **string** | The full text of the license | 
**CorrectedAt** | **time.Time** | Timestamp when the license correction was created | 

## Methods

### NewGetCustomLicenses200ResponseResultsInner

`func NewGetCustomLicenses200ResponseResultsInner(dependencyProjectLocator string, dependencyRevisionLocator string, projectLocator string, projectRevision string, title string, description string, url string, licenseTitle string, licenseId string, copyright string, text string, correctedAt time.Time, ) *GetCustomLicenses200ResponseResultsInner`

NewGetCustomLicenses200ResponseResultsInner instantiates a new GetCustomLicenses200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomLicenses200ResponseResultsInnerWithDefaults

`func NewGetCustomLicenses200ResponseResultsInnerWithDefaults() *GetCustomLicenses200ResponseResultsInner`

NewGetCustomLicenses200ResponseResultsInnerWithDefaults instantiates a new GetCustomLicenses200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyProjectLocator

`func (o *GetCustomLicenses200ResponseResultsInner) GetDependencyProjectLocator() string`

GetDependencyProjectLocator returns the DependencyProjectLocator field if non-nil, zero value otherwise.

### GetDependencyProjectLocatorOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetDependencyProjectLocatorOk() (*string, bool)`

GetDependencyProjectLocatorOk returns a tuple with the DependencyProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyProjectLocator

`func (o *GetCustomLicenses200ResponseResultsInner) SetDependencyProjectLocator(v string)`

SetDependencyProjectLocator sets DependencyProjectLocator field to given value.


### GetDependencyRevisionLocator

`func (o *GetCustomLicenses200ResponseResultsInner) GetDependencyRevisionLocator() string`

GetDependencyRevisionLocator returns the DependencyRevisionLocator field if non-nil, zero value otherwise.

### GetDependencyRevisionLocatorOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetDependencyRevisionLocatorOk() (*string, bool)`

GetDependencyRevisionLocatorOk returns a tuple with the DependencyRevisionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyRevisionLocator

`func (o *GetCustomLicenses200ResponseResultsInner) SetDependencyRevisionLocator(v string)`

SetDependencyRevisionLocator sets DependencyRevisionLocator field to given value.


### GetProjectLocator

`func (o *GetCustomLicenses200ResponseResultsInner) GetProjectLocator() string`

GetProjectLocator returns the ProjectLocator field if non-nil, zero value otherwise.

### GetProjectLocatorOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetProjectLocatorOk() (*string, bool)`

GetProjectLocatorOk returns a tuple with the ProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLocator

`func (o *GetCustomLicenses200ResponseResultsInner) SetProjectLocator(v string)`

SetProjectLocator sets ProjectLocator field to given value.


### GetProjectRevision

`func (o *GetCustomLicenses200ResponseResultsInner) GetProjectRevision() string`

GetProjectRevision returns the ProjectRevision field if non-nil, zero value otherwise.

### GetProjectRevisionOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetProjectRevisionOk() (*string, bool)`

GetProjectRevisionOk returns a tuple with the ProjectRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectRevision

`func (o *GetCustomLicenses200ResponseResultsInner) SetProjectRevision(v string)`

SetProjectRevision sets ProjectRevision field to given value.


### GetTitle

`func (o *GetCustomLicenses200ResponseResultsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetCustomLicenses200ResponseResultsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *GetCustomLicenses200ResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCustomLicenses200ResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *GetCustomLicenses200ResponseResultsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetCustomLicenses200ResponseResultsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLicenseTitle

`func (o *GetCustomLicenses200ResponseResultsInner) GetLicenseTitle() string`

GetLicenseTitle returns the LicenseTitle field if non-nil, zero value otherwise.

### GetLicenseTitleOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetLicenseTitleOk() (*string, bool)`

GetLicenseTitleOk returns a tuple with the LicenseTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTitle

`func (o *GetCustomLicenses200ResponseResultsInner) SetLicenseTitle(v string)`

SetLicenseTitle sets LicenseTitle field to given value.


### GetLicenseId

`func (o *GetCustomLicenses200ResponseResultsInner) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *GetCustomLicenses200ResponseResultsInner) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetCopyright

`func (o *GetCustomLicenses200ResponseResultsInner) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *GetCustomLicenses200ResponseResultsInner) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.


### GetText

`func (o *GetCustomLicenses200ResponseResultsInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GetCustomLicenses200ResponseResultsInner) SetText(v string)`

SetText sets Text field to given value.


### GetCorrectedAt

`func (o *GetCustomLicenses200ResponseResultsInner) GetCorrectedAt() time.Time`

GetCorrectedAt returns the CorrectedAt field if non-nil, zero value otherwise.

### GetCorrectedAtOk

`func (o *GetCustomLicenses200ResponseResultsInner) GetCorrectedAtOk() (*time.Time, bool)`

GetCorrectedAtOk returns a tuple with the CorrectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrectedAt

`func (o *GetCustomLicenses200ResponseResultsInner) SetCorrectedAt(v time.Time)`

SetCorrectedAt sets CorrectedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


