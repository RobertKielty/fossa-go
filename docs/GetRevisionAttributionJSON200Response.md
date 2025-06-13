# GetRevisionAttributionJSON200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**GetRevisionAttributionJSON200ResponseProject**](GetRevisionAttributionJSON200ResponseProject.md) |  | [optional] 
**DirectDependencies** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInner.md) |  | [optional] 
**DeepDependencies** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInner.md) |  | [optional] 
**Licenses** | Pointer to **map[string]interface{}** | A record of license names to their text | [optional] 
**CopyrightsByLicense** | Pointer to **map[string]interface{}** | A record of license names to their copyrights | [optional] 
**NoticeFiles** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner.md) | A list of notice file matches | [optional] 

## Methods

### NewGetRevisionAttributionJSON200Response

`func NewGetRevisionAttributionJSON200Response() *GetRevisionAttributionJSON200Response`

NewGetRevisionAttributionJSON200Response instantiates a new GetRevisionAttributionJSON200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionAttributionJSON200ResponseWithDefaults

`func NewGetRevisionAttributionJSON200ResponseWithDefaults() *GetRevisionAttributionJSON200Response`

NewGetRevisionAttributionJSON200ResponseWithDefaults instantiates a new GetRevisionAttributionJSON200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *GetRevisionAttributionJSON200Response) GetProject() GetRevisionAttributionJSON200ResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *GetRevisionAttributionJSON200Response) GetProjectOk() (*GetRevisionAttributionJSON200ResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *GetRevisionAttributionJSON200Response) SetProject(v GetRevisionAttributionJSON200ResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *GetRevisionAttributionJSON200Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetDirectDependencies

`func (o *GetRevisionAttributionJSON200Response) GetDirectDependencies() []GetRevisionAttributionJSON200ResponseDirectDependenciesInner`

GetDirectDependencies returns the DirectDependencies field if non-nil, zero value otherwise.

### GetDirectDependenciesOk

`func (o *GetRevisionAttributionJSON200Response) GetDirectDependenciesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInner, bool)`

GetDirectDependenciesOk returns a tuple with the DirectDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependencies

`func (o *GetRevisionAttributionJSON200Response) SetDirectDependencies(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInner)`

SetDirectDependencies sets DirectDependencies field to given value.

### HasDirectDependencies

`func (o *GetRevisionAttributionJSON200Response) HasDirectDependencies() bool`

HasDirectDependencies returns a boolean if a field has been set.

### GetDeepDependencies

`func (o *GetRevisionAttributionJSON200Response) GetDeepDependencies() []GetRevisionAttributionJSON200ResponseDirectDependenciesInner`

GetDeepDependencies returns the DeepDependencies field if non-nil, zero value otherwise.

### GetDeepDependenciesOk

`func (o *GetRevisionAttributionJSON200Response) GetDeepDependenciesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInner, bool)`

GetDeepDependenciesOk returns a tuple with the DeepDependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepDependencies

`func (o *GetRevisionAttributionJSON200Response) SetDeepDependencies(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInner)`

SetDeepDependencies sets DeepDependencies field to given value.

### HasDeepDependencies

`func (o *GetRevisionAttributionJSON200Response) HasDeepDependencies() bool`

HasDeepDependencies returns a boolean if a field has been set.

### GetLicenses

`func (o *GetRevisionAttributionJSON200Response) GetLicenses() map[string]interface{}`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetRevisionAttributionJSON200Response) GetLicensesOk() (*map[string]interface{}, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetRevisionAttributionJSON200Response) SetLicenses(v map[string]interface{})`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetRevisionAttributionJSON200Response) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetCopyrightsByLicense

`func (o *GetRevisionAttributionJSON200Response) GetCopyrightsByLicense() map[string]interface{}`

GetCopyrightsByLicense returns the CopyrightsByLicense field if non-nil, zero value otherwise.

### GetCopyrightsByLicenseOk

`func (o *GetRevisionAttributionJSON200Response) GetCopyrightsByLicenseOk() (*map[string]interface{}, bool)`

GetCopyrightsByLicenseOk returns a tuple with the CopyrightsByLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightsByLicense

`func (o *GetRevisionAttributionJSON200Response) SetCopyrightsByLicense(v map[string]interface{})`

SetCopyrightsByLicense sets CopyrightsByLicense field to given value.

### HasCopyrightsByLicense

`func (o *GetRevisionAttributionJSON200Response) HasCopyrightsByLicense() bool`

HasCopyrightsByLicense returns a boolean if a field has been set.

### GetNoticeFiles

`func (o *GetRevisionAttributionJSON200Response) GetNoticeFiles() []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner`

GetNoticeFiles returns the NoticeFiles field if non-nil, zero value otherwise.

### GetNoticeFilesOk

`func (o *GetRevisionAttributionJSON200Response) GetNoticeFilesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner, bool)`

GetNoticeFilesOk returns a tuple with the NoticeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeFiles

`func (o *GetRevisionAttributionJSON200Response) SetNoticeFiles(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner)`

SetNoticeFiles sets NoticeFiles field to given value.

### HasNoticeFiles

`func (o *GetRevisionAttributionJSON200Response) HasNoticeFiles() bool`

HasNoticeFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


