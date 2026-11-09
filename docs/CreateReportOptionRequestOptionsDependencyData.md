# CreateReportOptionRequestOptionsDependencyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | **bool** | For release group reports, show which projects the dependency is present in | 
**Authors** | **bool** | The authors of the dependency | 
**Description** | **bool** | The description of the dependency | 
**Homepage** | **bool** | The homepage of the dependency | 
**PackageManager** | **bool** | The package manager of the dependency | 
**DownloadUrl** | **bool** | The download URL of the dependency | 
**ConcludedLicenses** | **bool** | The concluded licenses of the dependency | 
**DeclaredLicenses** | **bool** | The declared licenses of the dependency | 
**DiscoveredLicenses** | **bool** | The discovered licenses of the dependency | 
**Copyrights** | **bool** | The copyrights of the dependency | 
**LicenseUrl** | **bool** | The license URL of the dependency | 
**LicenseFileMatches** | **bool** | The list of files that the license was discovered in | 
**IssueResolutionNotes** | **bool** | Issue resolution notes | 
**PackageLabels** | **bool** | Package labels | 
**DependencyPaths** | **bool** | The origin paths of the dependency (what files the dependency was defined/found in) | 
**FilePaths** | **bool** | The paths_to of the dependency (the chain of dependencies that brought it into the project) | 
**NoticeFiles** | **bool** | Notice files | 
**FullLicenseText** | **bool** | Full license text | 
**LicenseTextAppendix** | Pointer to **bool** | Move full license texts out of each dependency and into a single Licenses appendix at the end of the report | [optional] 

## Methods

### NewCreateReportOptionRequestOptionsDependencyData

`func NewCreateReportOptionRequestOptionsDependencyData(projects bool, authors bool, description bool, homepage bool, packageManager bool, downloadUrl bool, concludedLicenses bool, declaredLicenses bool, discoveredLicenses bool, copyrights bool, licenseUrl bool, licenseFileMatches bool, issueResolutionNotes bool, packageLabels bool, dependencyPaths bool, filePaths bool, noticeFiles bool, fullLicenseText bool, ) *CreateReportOptionRequestOptionsDependencyData`

NewCreateReportOptionRequestOptionsDependencyData instantiates a new CreateReportOptionRequestOptionsDependencyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportOptionRequestOptionsDependencyDataWithDefaults

`func NewCreateReportOptionRequestOptionsDependencyDataWithDefaults() *CreateReportOptionRequestOptionsDependencyData`

NewCreateReportOptionRequestOptionsDependencyDataWithDefaults instantiates a new CreateReportOptionRequestOptionsDependencyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *CreateReportOptionRequestOptionsDependencyData) GetProjects() bool`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetProjectsOk() (*bool, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CreateReportOptionRequestOptionsDependencyData) SetProjects(v bool)`

SetProjects sets Projects field to given value.


### GetAuthors

`func (o *CreateReportOptionRequestOptionsDependencyData) GetAuthors() bool`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetAuthorsOk() (*bool, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *CreateReportOptionRequestOptionsDependencyData) SetAuthors(v bool)`

SetAuthors sets Authors field to given value.


### GetDescription

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDescription() bool`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDescriptionOk() (*bool, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateReportOptionRequestOptionsDependencyData) SetDescription(v bool)`

SetDescription sets Description field to given value.


### GetHomepage

`func (o *CreateReportOptionRequestOptionsDependencyData) GetHomepage() bool`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetHomepageOk() (*bool, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *CreateReportOptionRequestOptionsDependencyData) SetHomepage(v bool)`

SetHomepage sets Homepage field to given value.


### GetPackageManager

`func (o *CreateReportOptionRequestOptionsDependencyData) GetPackageManager() bool`

GetPackageManager returns the PackageManager field if non-nil, zero value otherwise.

### GetPackageManagerOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetPackageManagerOk() (*bool, bool)`

GetPackageManagerOk returns a tuple with the PackageManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManager

`func (o *CreateReportOptionRequestOptionsDependencyData) SetPackageManager(v bool)`

SetPackageManager sets PackageManager field to given value.


### GetDownloadUrl

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDownloadUrl() bool`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDownloadUrlOk() (*bool, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *CreateReportOptionRequestOptionsDependencyData) SetDownloadUrl(v bool)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetConcludedLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) GetConcludedLicenses() bool`

GetConcludedLicenses returns the ConcludedLicenses field if non-nil, zero value otherwise.

### GetConcludedLicensesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetConcludedLicensesOk() (*bool, bool)`

GetConcludedLicensesOk returns a tuple with the ConcludedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcludedLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) SetConcludedLicenses(v bool)`

SetConcludedLicenses sets ConcludedLicenses field to given value.


### GetDeclaredLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDeclaredLicenses() bool`

GetDeclaredLicenses returns the DeclaredLicenses field if non-nil, zero value otherwise.

### GetDeclaredLicensesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDeclaredLicensesOk() (*bool, bool)`

GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) SetDeclaredLicenses(v bool)`

SetDeclaredLicenses sets DeclaredLicenses field to given value.


### GetDiscoveredLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDiscoveredLicenses() bool`

GetDiscoveredLicenses returns the DiscoveredLicenses field if non-nil, zero value otherwise.

### GetDiscoveredLicensesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDiscoveredLicensesOk() (*bool, bool)`

GetDiscoveredLicensesOk returns a tuple with the DiscoveredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredLicenses

`func (o *CreateReportOptionRequestOptionsDependencyData) SetDiscoveredLicenses(v bool)`

SetDiscoveredLicenses sets DiscoveredLicenses field to given value.


### GetCopyrights

`func (o *CreateReportOptionRequestOptionsDependencyData) GetCopyrights() bool`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetCopyrightsOk() (*bool, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *CreateReportOptionRequestOptionsDependencyData) SetCopyrights(v bool)`

SetCopyrights sets Copyrights field to given value.


### GetLicenseUrl

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseUrl() bool`

GetLicenseUrl returns the LicenseUrl field if non-nil, zero value otherwise.

### GetLicenseUrlOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseUrlOk() (*bool, bool)`

GetLicenseUrlOk returns a tuple with the LicenseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseUrl

`func (o *CreateReportOptionRequestOptionsDependencyData) SetLicenseUrl(v bool)`

SetLicenseUrl sets LicenseUrl field to given value.


### GetLicenseFileMatches

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseFileMatches() bool`

GetLicenseFileMatches returns the LicenseFileMatches field if non-nil, zero value otherwise.

### GetLicenseFileMatchesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseFileMatchesOk() (*bool, bool)`

GetLicenseFileMatchesOk returns a tuple with the LicenseFileMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFileMatches

`func (o *CreateReportOptionRequestOptionsDependencyData) SetLicenseFileMatches(v bool)`

SetLicenseFileMatches sets LicenseFileMatches field to given value.


### GetIssueResolutionNotes

`func (o *CreateReportOptionRequestOptionsDependencyData) GetIssueResolutionNotes() bool`

GetIssueResolutionNotes returns the IssueResolutionNotes field if non-nil, zero value otherwise.

### GetIssueResolutionNotesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetIssueResolutionNotesOk() (*bool, bool)`

GetIssueResolutionNotesOk returns a tuple with the IssueResolutionNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueResolutionNotes

`func (o *CreateReportOptionRequestOptionsDependencyData) SetIssueResolutionNotes(v bool)`

SetIssueResolutionNotes sets IssueResolutionNotes field to given value.


### GetPackageLabels

`func (o *CreateReportOptionRequestOptionsDependencyData) GetPackageLabels() bool`

GetPackageLabels returns the PackageLabels field if non-nil, zero value otherwise.

### GetPackageLabelsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetPackageLabelsOk() (*bool, bool)`

GetPackageLabelsOk returns a tuple with the PackageLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabels

`func (o *CreateReportOptionRequestOptionsDependencyData) SetPackageLabels(v bool)`

SetPackageLabels sets PackageLabels field to given value.


### GetDependencyPaths

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDependencyPaths() bool`

GetDependencyPaths returns the DependencyPaths field if non-nil, zero value otherwise.

### GetDependencyPathsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetDependencyPathsOk() (*bool, bool)`

GetDependencyPathsOk returns a tuple with the DependencyPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyPaths

`func (o *CreateReportOptionRequestOptionsDependencyData) SetDependencyPaths(v bool)`

SetDependencyPaths sets DependencyPaths field to given value.


### GetFilePaths

`func (o *CreateReportOptionRequestOptionsDependencyData) GetFilePaths() bool`

GetFilePaths returns the FilePaths field if non-nil, zero value otherwise.

### GetFilePathsOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetFilePathsOk() (*bool, bool)`

GetFilePathsOk returns a tuple with the FilePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePaths

`func (o *CreateReportOptionRequestOptionsDependencyData) SetFilePaths(v bool)`

SetFilePaths sets FilePaths field to given value.


### GetNoticeFiles

`func (o *CreateReportOptionRequestOptionsDependencyData) GetNoticeFiles() bool`

GetNoticeFiles returns the NoticeFiles field if non-nil, zero value otherwise.

### GetNoticeFilesOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetNoticeFilesOk() (*bool, bool)`

GetNoticeFilesOk returns a tuple with the NoticeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeFiles

`func (o *CreateReportOptionRequestOptionsDependencyData) SetNoticeFiles(v bool)`

SetNoticeFiles sets NoticeFiles field to given value.


### GetFullLicenseText

`func (o *CreateReportOptionRequestOptionsDependencyData) GetFullLicenseText() bool`

GetFullLicenseText returns the FullLicenseText field if non-nil, zero value otherwise.

### GetFullLicenseTextOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetFullLicenseTextOk() (*bool, bool)`

GetFullLicenseTextOk returns a tuple with the FullLicenseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullLicenseText

`func (o *CreateReportOptionRequestOptionsDependencyData) SetFullLicenseText(v bool)`

SetFullLicenseText sets FullLicenseText field to given value.


### GetLicenseTextAppendix

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseTextAppendix() bool`

GetLicenseTextAppendix returns the LicenseTextAppendix field if non-nil, zero value otherwise.

### GetLicenseTextAppendixOk

`func (o *CreateReportOptionRequestOptionsDependencyData) GetLicenseTextAppendixOk() (*bool, bool)`

GetLicenseTextAppendixOk returns a tuple with the LicenseTextAppendix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseTextAppendix

`func (o *CreateReportOptionRequestOptionsDependencyData) SetLicenseTextAppendix(v bool)`

SetLicenseTextAppendix sets LicenseTextAppendix field to given value.

### HasLicenseTextAppendix

`func (o *CreateReportOptionRequestOptionsDependencyData) HasLicenseTextAppendix() bool`

HasLicenseTextAppendix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


