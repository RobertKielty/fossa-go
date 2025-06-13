# GetReleaseGroupById200ResponseAllOfReleasesInner

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

## Methods

### NewGetReleaseGroupById200ResponseAllOfReleasesInner

`func NewGetReleaseGroupById200ResponseAllOfReleasesInner() *GetReleaseGroupById200ResponseAllOfReleasesInner`

NewGetReleaseGroupById200ResponseAllOfReleasesInner instantiates a new GetReleaseGroupById200ResponseAllOfReleasesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupById200ResponseAllOfReleasesInnerWithDefaults

`func NewGetReleaseGroupById200ResponseAllOfReleasesInnerWithDefaults() *GetReleaseGroupById200ResponseAllOfReleasesInner`

NewGetReleaseGroupById200ResponseAllOfReleasesInnerWithDefaults instantiates a new GetReleaseGroupById200ResponseAllOfReleasesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetProjectGroupId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetProjectGroupId() int32`

GetProjectGroupId returns the ProjectGroupId field if non-nil, zero value otherwise.

### GetProjectGroupIdOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetProjectGroupIdOk() (*int32, bool)`

GetProjectGroupIdOk returns a tuple with the ProjectGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectGroupId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetProjectGroupId(v int32)`

SetProjectGroupId sets ProjectGroupId field to given value.

### HasProjectGroupId

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasProjectGroupId() bool`

HasProjectGroupId returns a boolean if a field has been set.

### GetDependencyCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetDependencyCount() int32`

GetDependencyCount returns the DependencyCount field if non-nil, zero value otherwise.

### GetDependencyCountOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetDependencyCountOk() (*int32, bool)`

GetDependencyCountOk returns a tuple with the DependencyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetDependencyCount(v int32)`

SetDependencyCount sets DependencyCount field to given value.

### HasDependencyCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasDependencyCount() bool`

HasDependencyCount returns a boolean if a field has been set.

### GetLicenseCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetLicenseCount() int32`

GetLicenseCount returns the LicenseCount field if non-nil, zero value otherwise.

### GetLicenseCountOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetLicenseCountOk() (*int32, bool)`

GetLicenseCountOk returns a tuple with the LicenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetLicenseCount(v int32)`

SetLicenseCount sets LicenseCount field to given value.

### HasLicenseCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasLicenseCount() bool`

HasLicenseCount returns a boolean if a field has been set.

### GetUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedLicensingIssueCount() int32`

GetUnresolvedLicensingIssueCount returns the UnresolvedLicensingIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedLicensingIssueCountOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedLicensingIssueCountOk() (*int32, bool)`

GetUnresolvedLicensingIssueCountOk returns a tuple with the UnresolvedLicensingIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetUnresolvedLicensingIssueCount(v int32)`

SetUnresolvedLicensingIssueCount sets UnresolvedLicensingIssueCount field to given value.

### HasUnresolvedLicensingIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasUnresolvedLicensingIssueCount() bool`

HasUnresolvedLicensingIssueCount returns a boolean if a field has been set.

### GetUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedSecurityIssueCount() int32`

GetUnresolvedSecurityIssueCount returns the UnresolvedSecurityIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedSecurityIssueCountOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedSecurityIssueCountOk() (*int32, bool)`

GetUnresolvedSecurityIssueCountOk returns a tuple with the UnresolvedSecurityIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetUnresolvedSecurityIssueCount(v int32)`

SetUnresolvedSecurityIssueCount sets UnresolvedSecurityIssueCount field to given value.

### HasUnresolvedSecurityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasUnresolvedSecurityIssueCount() bool`

HasUnresolvedSecurityIssueCount returns a boolean if a field has been set.

### GetUnresolvedQualityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedQualityIssueCount() int32`

GetUnresolvedQualityIssueCount returns the UnresolvedQualityIssueCount field if non-nil, zero value otherwise.

### GetUnresolvedQualityIssueCountOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetUnresolvedQualityIssueCountOk() (*int32, bool)`

GetUnresolvedQualityIssueCountOk returns a tuple with the UnresolvedQualityIssueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolvedQualityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetUnresolvedQualityIssueCount(v int32)`

SetUnresolvedQualityIssueCount sets UnresolvedQualityIssueCount field to given value.

### HasUnresolvedQualityIssueCount

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasUnresolvedQualityIssueCount() bool`

HasUnresolvedQualityIssueCount returns a boolean if a field has been set.

### GetPublishedOnPortal

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedOnPortal() string`

GetPublishedOnPortal returns the PublishedOnPortal field if non-nil, zero value otherwise.

### GetPublishedOnPortalOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedOnPortalOk() (*string, bool)`

GetPublishedOnPortalOk returns a tuple with the PublishedOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOnPortal

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetPublishedOnPortal(v string)`

SetPublishedOnPortal sets PublishedOnPortal field to given value.

### HasPublishedOnPortal

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasPublishedOnPortal() bool`

HasPublishedOnPortal returns a boolean if a field has been set.

### GetPublishedAt

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetReportPath

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetReportPath() string`

GetReportPath returns the ReportPath field if non-nil, zero value otherwise.

### GetReportPathOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetReportPathOk() (*string, bool)`

GetReportPathOk returns a tuple with the ReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPath

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetReportPath(v string)`

SetReportPath sets ReportPath field to given value.

### HasReportPath

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasReportPath() bool`

HasReportPath returns a boolean if a field has been set.

### GetPublishedLicenses

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedLicenses() []string`

GetPublishedLicenses returns the PublishedLicenses field if non-nil, zero value otherwise.

### GetPublishedLicensesOk

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) GetPublishedLicensesOk() (*[]string, bool)`

GetPublishedLicensesOk returns a tuple with the PublishedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedLicenses

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) SetPublishedLicenses(v []string)`

SetPublishedLicenses sets PublishedLicenses field to given value.

### HasPublishedLicenses

`func (o *GetReleaseGroupById200ResponseAllOfReleasesInner) HasPublishedLicenses() bool`

HasPublishedLicenses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


