# GetRevisionAttributionJSON200ResponseDirectDependenciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Package** | Pointer to **string** | The package name | [optional] 
**Source** | Pointer to **string** | The source of the package | [optional] 
**Version** | Pointer to **string** | The version of the package | [optional] 
**Hash** | Pointer to **string** | The hash of the package | [optional] 
**Authors** | Pointer to **[]string** | The authors of the package | [optional] 
**Description** | Pointer to **string** | The description of the package | [optional] 
**Licenses** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner.md) | The licenses of the package | [optional] 
**OtherLicenses** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner.md) | The other licenses of the package | [optional] 
**ProjectUrl** | Pointer to **string** | The project URL | [optional] 
**DependencyPaths** | Pointer to **[]string** | The dependency paths | [optional] 
**Notes** | Pointer to **[]string** | The notes of the package | [optional] 
**DownloadUrl** | Pointer to **string** | The download URL | [optional] 
**IsGolang** | Pointer to **bool** | Whether the package is Golang | [optional] 
**Title** | Pointer to **string** | The title of the package | [optional] 
**NoticeFiles** | Pointer to [**[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner**](GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner.md) | The notice file matches of the package | [optional] 
**PackageLabels** | Pointer to **[]string** | All applicable package labels assigned to the package, including globally applied labels, project applied labels, and revision applied labels that match this package in this revision. | [optional] 

## Methods

### NewGetRevisionAttributionJSON200ResponseDirectDependenciesInner

`func NewGetRevisionAttributionJSON200ResponseDirectDependenciesInner() *GetRevisionAttributionJSON200ResponseDirectDependenciesInner`

NewGetRevisionAttributionJSON200ResponseDirectDependenciesInner instantiates a new GetRevisionAttributionJSON200ResponseDirectDependenciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerWithDefaults

`func NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerWithDefaults() *GetRevisionAttributionJSON200ResponseDirectDependenciesInner`

NewGetRevisionAttributionJSON200ResponseDirectDependenciesInnerWithDefaults instantiates a new GetRevisionAttributionJSON200ResponseDirectDependenciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackage

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSource

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVersion

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHash

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetAuthors

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.

### HasAuthors

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasAuthors() bool`

HasAuthors returns a boolean if a field has been set.

### GetDescription

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetLicenses() []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetLicensesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetLicenses(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetOtherLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetOtherLicenses() []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner`

GetOtherLicenses returns the OtherLicenses field if non-nil, zero value otherwise.

### GetOtherLicensesOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetOtherLicensesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner, bool)`

GetOtherLicensesOk returns a tuple with the OtherLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetOtherLicenses(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerLicensesInner)`

SetOtherLicenses sets OtherLicenses field to given value.

### HasOtherLicenses

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasOtherLicenses() bool`

HasOtherLicenses returns a boolean if a field has been set.

### GetProjectUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetProjectUrl() string`

GetProjectUrl returns the ProjectUrl field if non-nil, zero value otherwise.

### GetProjectUrlOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetProjectUrlOk() (*string, bool)`

GetProjectUrlOk returns a tuple with the ProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetProjectUrl(v string)`

SetProjectUrl sets ProjectUrl field to given value.

### HasProjectUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasProjectUrl() bool`

HasProjectUrl returns a boolean if a field has been set.

### GetDependencyPaths

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDependencyPaths() []string`

GetDependencyPaths returns the DependencyPaths field if non-nil, zero value otherwise.

### GetDependencyPathsOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDependencyPathsOk() (*[]string, bool)`

GetDependencyPathsOk returns a tuple with the DependencyPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyPaths

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetDependencyPaths(v []string)`

SetDependencyPaths sets DependencyPaths field to given value.

### HasDependencyPaths

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasDependencyPaths() bool`

HasDependencyPaths returns a boolean if a field has been set.

### GetNotes

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetNotes(v []string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDownloadUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.

### HasDownloadUrl

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasDownloadUrl() bool`

HasDownloadUrl returns a boolean if a field has been set.

### GetIsGolang

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetIsGolang() bool`

GetIsGolang returns the IsGolang field if non-nil, zero value otherwise.

### GetIsGolangOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetIsGolangOk() (*bool, bool)`

GetIsGolangOk returns a tuple with the IsGolang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGolang

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetIsGolang(v bool)`

SetIsGolang sets IsGolang field to given value.

### HasIsGolang

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasIsGolang() bool`

HasIsGolang returns a boolean if a field has been set.

### GetTitle

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetNoticeFiles

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetNoticeFiles() []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner`

GetNoticeFiles returns the NoticeFiles field if non-nil, zero value otherwise.

### GetNoticeFilesOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetNoticeFilesOk() (*[]GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner, bool)`

GetNoticeFilesOk returns a tuple with the NoticeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeFiles

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetNoticeFiles(v []GetRevisionAttributionJSON200ResponseDirectDependenciesInnerNoticeFilesInner)`

SetNoticeFiles sets NoticeFiles field to given value.

### HasNoticeFiles

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasNoticeFiles() bool`

HasNoticeFiles returns a boolean if a field has been set.

### GetPackageLabels

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetPackageLabels() []string`

GetPackageLabels returns the PackageLabels field if non-nil, zero value otherwise.

### GetPackageLabelsOk

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) GetPackageLabelsOk() (*[]string, bool)`

GetPackageLabelsOk returns a tuple with the PackageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabels

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) SetPackageLabels(v []string)`

SetPackageLabels sets PackageLabels field to given value.

### HasPackageLabels

`func (o *GetRevisionAttributionJSON200ResponseDirectDependenciesInner) HasPackageLabels() bool`

HasPackageLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


