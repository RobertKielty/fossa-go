# GetReleaseGroupReleaseById200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Release ID | [optional] 
**Title** | Pointer to **string** | Release title/version | [optional] 
**ProjectGroupId** | Pointer to **int32** | The release group which this belongs to | [optional] 
**DependencyCount** | Pointer to **int32** | The count of dependencies in this release | [optional] 
**LicenseCount** | Pointer to **int32** | The count of licenses in this release | [optional] 
**UnresolvedLicensingIssueCount** | Pointer to **int32** | The number of licensing issues in this release | [optional] 
**UnresolvedSecurityIssueCount** | Pointer to **int32** | The number of security issues in this release | [optional] 
**UnresolvedQualityIssueCount** | Pointer to **int32** | The number of quality issues in this release | [optional] 
**PublishedOnPortal** | Pointer to **string** | If this release has been published on an SBOM portal | [optional] 
**PublishedAt** | Pointer to **time.Time** | When the release was published to the portal | [optional] 
**ReportPath** | Pointer to **string** | Path to the SBOM report for this release | [optional] 
**PublishedLicenses** | Pointer to **[]string** | List of published license IDs | [optional] 
**Projects** | Pointer to [**[]GetReleaseGroupReleaseById200ResponseAllOfProjectsInner**](GetReleaseGroupReleaseById200ResponseAllOfProjectsInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroupReleaseById200Response

`func NewGetReleaseGroupReleaseById200Response() *GetReleaseGroupReleaseById200Response`

NewGetReleaseGroupReleaseById200Response instantiates a new GetReleaseGroupReleaseById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupReleaseById200ResponseWithDefaults

`func NewGetReleaseGroupReleaseById200ResponseWithDefaults() *GetReleaseGroupReleaseById200Response`

NewGetReleaseGroupReleaseById200ResponseWithDefaults instantiates a new GetReleaseGroupReleaseById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupReleaseById200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupReleaseById200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupReleaseById200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetReleaseGroupReleaseById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetReleaseGroupReleaseById200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupReleaseById200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupReleaseById200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetReleaseGroupReleaseById200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProjectGroupId

`func (o *GetReleaseGroupReleaseById200Response) GetProjectGroupId() int32`

GetProjectGroupId returns the ProjectGroupId field if non-nil, zero value otherwise.

### GetProjectGroupIdOk

`func (o *GetReleaseGroupReleaseById200Response) GetProjectGroupIdOk() (*int32, bool)`

GetProjectGroupIdOk returns a tuple with the ProjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupId

`func (o *GetReleaseGroupReleaseById200Response) SetProjectGroupId(v int32)`

SetProjectGroupId sets ProjectGroupId field to given value.

### HasProjectGroupId

`func (o *GetReleaseGroupReleaseById200Response) HasProjectGroupId() bool`

HasProjectGroupId returns a boolean if a field has been set.

### GetDependencyCount

`func (o *GetReleaseGroupReleaseById200Response) GetDependencyCount() int32`

GetDependencyCount returns the DependencyCount field if non-nil, zero value otherwise.

### GetDependencyCountOk

`func (o *GetReleaseGroupReleaseById200Response) GetDependencyCountOk() (*int32, bool)`

GetDependencyCountOk returns a tuple with the DependencyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyCount

`func (o *GetReleaseGroupReleaseById200Response) SetDependencyCount(v int32)`

SetDependencyCount sets DependencyCount field to given value.

### HasDependencyCount

`func (o *GetReleaseGroupReleaseById200Response) HasDependencyCount() bool`

HasDependencyCount returns a boolean if a field has been set.

### GetLicenseCount

`func (o *GetReleaseGroupReleaseById200Response) GetLicenseCount() int32`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *GetReleaseGroupReleaseById200Response) GetLicenseCountOk() (*int32, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *GetReleaseGroupReleaseById200Response) SetLicenseCount(v int32)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *GetReleaseGroupReleaseById200Response) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedLicensingIssueCount() int32`

GetUnresolvedLicensingIssueCount returns the UnresolvedLicensingIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedLicensingIssueCountOk

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedLicensingIssueCountOk() (*int32, bool)`

GetUnresolvedLicensingIssueCountOk returns a tuple with the UnresolvedLicensingIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupReleaseById200Response) SetUnresolvedLicensingIssueCount(v int32)`

SetUnresolvedLicensingIssueCount sets UnresolvedLicensingIssueCount field to given value.

### HasUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupReleaseById200Response) HasUnresolvedLicensingIssueCount() bool`

HasUnresolvedLicensingIssueCount returns a boolean if a field has been set.

### GetUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedSecurityIssueCount() int32`

GetUnresolvedSecurityIssueCount returns the UnresolvedSecurityIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedSecurityIssueCountOk

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedSecurityIssueCountOk() (*int32, bool)`

GetUnresolvedSecurityIssueCountOk returns a tuple with the UnresolvedSecurityIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) SetUnresolvedSecurityIssueCount(v int32)`

SetUnresolvedSecurityIssueCount sets UnresolvedSecurityIssueCount field to given value.

### HasUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) HasUnresolvedSecurityIssueCount() bool`

HasUnresolvedSecurityIssueCount returns a boolean if a field has been set.

### GetUnresolvedQualityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedQualityIssueCount() int32`

GetUnresolvedQualityIssueCount returns the UnresolvedQualityIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedQualityIssueCountOk

`func (o *GetReleaseGroupReleaseById200Response) GetUnresolvedQualityIssueCountOk() (*int32, bool)`

GetUnresolvedQualityIssueCountOk returns a tuple with the UnresolvedQualityIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedQualityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) SetUnresolvedQualityIssueCount(v int32)`

SetUnresolvedQualityIssueCount sets UnresolvedQualityIssueCount field to given value.

### HasUnresolvedQualityIssueCount

`func (o *GetReleaseGroupReleaseById200Response) HasUnresolvedQualityIssueCount() bool`

HasUnresolvedQualityIssueCount returns a boolean if a field has been set.

### GetPublishedOnPortal

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedOnPortal() string`

GetPublishedOnPortal returns the PublishedOnPortal field if non-nil, zero value otherwise.

### GetPublishedOnPortalOk

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedOnPortalOk() (*string, bool)`

GetPublishedOnPortalOk returns a tuple with the PublishedOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOnPortal

`func (o *GetReleaseGroupReleaseById200Response) SetPublishedOnPortal(v string)`

SetPublishedOnPortal sets PublishedOnPortal field to given value.

### HasPublishedOnPortal

`func (o *GetReleaseGroupReleaseById200Response) HasPublishedOnPortal() bool`

HasPublishedOnPortal returns a boolean if a field has been set.

### GetPublishedAt

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetReleaseGroupReleaseById200Response) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *GetReleaseGroupReleaseById200Response) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetReportPath

`func (o *GetReleaseGroupReleaseById200Response) GetReportPath() string`

GetReportPath returns the ReportPath field if non-nil, zero value otherwise.

### GetReportPathOk

`func (o *GetReleaseGroupReleaseById200Response) GetReportPathOk() (*string, bool)`

GetReportPathOk returns a tuple with the ReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPath

`func (o *GetReleaseGroupReleaseById200Response) SetReportPath(v string)`

SetReportPath sets ReportPath field to given value.

### HasReportPath

`func (o *GetReleaseGroupReleaseById200Response) HasReportPath() bool`

HasReportPath returns a boolean if a field has been set.

### GetPublishedLicenses

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedLicenses() []string`

GetPublishedLicenses returns the PublishedLicenses field if non-nil, zero value otherwise.

### GetPublishedLicensesOk

`func (o *GetReleaseGroupReleaseById200Response) GetPublishedLicensesOk() (*[]string, bool)`

GetPublishedLicensesOk returns a tuple with the PublishedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedLicenses

`func (o *GetReleaseGroupReleaseById200Response) SetPublishedLicenses(v []string)`

SetPublishedLicenses sets PublishedLicenses field to given value.

### HasPublishedLicenses

`func (o *GetReleaseGroupReleaseById200Response) HasPublishedLicenses() bool`

HasPublishedLicenses returns a boolean if a field has been set.

### GetProjects

`func (o *GetReleaseGroupReleaseById200Response) GetProjects() []GetReleaseGroupReleaseById200ResponseAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetReleaseGroupReleaseById200Response) GetProjectsOk() (*[]GetReleaseGroupReleaseById200ResponseAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetReleaseGroupReleaseById200Response) SetProjects(v []GetReleaseGroupReleaseById200ResponseAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetReleaseGroupReleaseById200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


