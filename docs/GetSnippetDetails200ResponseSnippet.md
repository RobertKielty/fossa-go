# GetSnippetDetails200ResponseSnippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the snippet | 
**PackageId** | **string** | Unique identifier for the snippet package | 
**Purl** | **string** | Package URL (purl) for the snippet | 
**Locator** | **string** | The locator for the snippet (optional) | 
**Package** | **string** | Name of the package containing the snippet | 
**Version** | **string** | Version of the package | 
**Kind** | **string** | Type of snippet detection (snippet&#x3D;partial match, file&#x3D;100% match) | 
**MatchCount** | **int32** | Total number of matches for this snippet | 
**Matches** | [**[]GetSnippetDetails200ResponseSnippetMatchesInner**](GetSnippetDetails200ResponseSnippetMatchesInner.md) | Array of path matches where this snippet was detected | 
**HighestMatchPercentage** | **float32** | The highest match percentage across all matches | 
**ReleaseDate** | Pointer to **time.Time** | Release date of the package (optional) | [optional] 
**HomeUrl** | Pointer to **string** | Homepage URL of the package (optional) | [optional] 
**CodeUrl** | Pointer to **string** | Source code URL of the package (optional) | [optional] 
**Licenses** | [**[]GetSnippets200ResponseResultsInnerLicensesInner**](GetSnippets200ResponseResultsInnerLicensesInner.md) | Array of licenses associated with the snippet | 
**IssueCounts** | [**GetSnippets200ResponseResultsInnerIssueCounts**](GetSnippets200ResponseResultsInnerIssueCounts.md) |  | 
**RejectionDetails** | Pointer to [**GetSnippets200ResponseResultsInnerRejectionDetails**](GetSnippets200ResponseResultsInnerRejectionDetails.md) |  | [optional] 
**Labels** | [**[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner**](GetProjectDependencies200ResponseDependenciesInnerLabelsInner.md) | Package labels assigned to this snippet | 
**OtherVersions** | Pointer to [**[]GetSnippetDetails200ResponseSnippetOtherVersionsInner**](GetSnippetDetails200ResponseSnippetOtherVersionsInner.md) | Other versions of the package where this snippet was detected | [optional] 
**IsVendored** | **bool** | Whether the snippet exists as a vendored dependency | 
**IsConverted** | **bool** | Whether the snippet has been converted to a vendored dependency | 

## Methods

### NewGetSnippetDetails200ResponseSnippet

`func NewGetSnippetDetails200ResponseSnippet(id string, packageId string, purl string, locator string, package_ string, version string, kind string, matchCount int32, matches []GetSnippetDetails200ResponseSnippetMatchesInner, highestMatchPercentage float32, licenses []GetSnippets200ResponseResultsInnerLicensesInner, issueCounts GetSnippets200ResponseResultsInnerIssueCounts, labels []GetProjectDependencies200ResponseDependenciesInnerLabelsInner, isVendored bool, isConverted bool, ) *GetSnippetDetails200ResponseSnippet`

NewGetSnippetDetails200ResponseSnippet instantiates a new GetSnippetDetails200ResponseSnippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetDetails200ResponseSnippetWithDefaults

`func NewGetSnippetDetails200ResponseSnippetWithDefaults() *GetSnippetDetails200ResponseSnippet`

NewGetSnippetDetails200ResponseSnippetWithDefaults instantiates a new GetSnippetDetails200ResponseSnippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSnippetDetails200ResponseSnippet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSnippetDetails200ResponseSnippet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSnippetDetails200ResponseSnippet) SetId(v string)`

SetId sets Id field to given value.


### GetPackageId

`func (o *GetSnippetDetails200ResponseSnippet) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetSnippetDetails200ResponseSnippet) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetSnippetDetails200ResponseSnippet) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPurl

`func (o *GetSnippetDetails200ResponseSnippet) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *GetSnippetDetails200ResponseSnippet) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *GetSnippetDetails200ResponseSnippet) SetPurl(v string)`

SetPurl sets Purl field to given value.


### GetLocator

`func (o *GetSnippetDetails200ResponseSnippet) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetSnippetDetails200ResponseSnippet) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetSnippetDetails200ResponseSnippet) SetLocator(v string)`

SetLocator sets Locator field to given value.


### GetPackage

`func (o *GetSnippetDetails200ResponseSnippet) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *GetSnippetDetails200ResponseSnippet) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *GetSnippetDetails200ResponseSnippet) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetVersion

`func (o *GetSnippetDetails200ResponseSnippet) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSnippetDetails200ResponseSnippet) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSnippetDetails200ResponseSnippet) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetKind

`func (o *GetSnippetDetails200ResponseSnippet) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetSnippetDetails200ResponseSnippet) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetSnippetDetails200ResponseSnippet) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMatchCount

`func (o *GetSnippetDetails200ResponseSnippet) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *GetSnippetDetails200ResponseSnippet) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *GetSnippetDetails200ResponseSnippet) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.


### GetMatches

`func (o *GetSnippetDetails200ResponseSnippet) GetMatches() []GetSnippetDetails200ResponseSnippetMatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *GetSnippetDetails200ResponseSnippet) GetMatchesOk() (*[]GetSnippetDetails200ResponseSnippetMatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *GetSnippetDetails200ResponseSnippet) SetMatches(v []GetSnippetDetails200ResponseSnippetMatchesInner)`

SetMatches sets Matches field to given value.


### GetHighestMatchPercentage

`func (o *GetSnippetDetails200ResponseSnippet) GetHighestMatchPercentage() float32`

GetHighestMatchPercentage returns the HighestMatchPercentage field if non-nil, zero value otherwise.

### GetHighestMatchPercentageOk

`func (o *GetSnippetDetails200ResponseSnippet) GetHighestMatchPercentageOk() (*float32, bool)`

GetHighestMatchPercentageOk returns a tuple with the HighestMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestMatchPercentage

`func (o *GetSnippetDetails200ResponseSnippet) SetHighestMatchPercentage(v float32)`

SetHighestMatchPercentage sets HighestMatchPercentage field to given value.


### GetReleaseDate

`func (o *GetSnippetDetails200ResponseSnippet) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetSnippetDetails200ResponseSnippet) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetSnippetDetails200ResponseSnippet) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetSnippetDetails200ResponseSnippet) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetHomeUrl

`func (o *GetSnippetDetails200ResponseSnippet) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *GetSnippetDetails200ResponseSnippet) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *GetSnippetDetails200ResponseSnippet) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *GetSnippetDetails200ResponseSnippet) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCodeUrl

`func (o *GetSnippetDetails200ResponseSnippet) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *GetSnippetDetails200ResponseSnippet) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *GetSnippetDetails200ResponseSnippet) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *GetSnippetDetails200ResponseSnippet) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetLicenses

`func (o *GetSnippetDetails200ResponseSnippet) GetLicenses() []GetSnippets200ResponseResultsInnerLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetSnippetDetails200ResponseSnippet) GetLicensesOk() (*[]GetSnippets200ResponseResultsInnerLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetSnippetDetails200ResponseSnippet) SetLicenses(v []GetSnippets200ResponseResultsInnerLicensesInner)`

SetLicenses sets Licenses field to given value.


### GetIssueCounts

`func (o *GetSnippetDetails200ResponseSnippet) GetIssueCounts() GetSnippets200ResponseResultsInnerIssueCounts`

GetIssueCounts returns the IssueCounts field if non-nil, zero value otherwise.

### GetIssueCountsOk

`func (o *GetSnippetDetails200ResponseSnippet) GetIssueCountsOk() (*GetSnippets200ResponseResultsInnerIssueCounts, bool)`

GetIssueCountsOk returns a tuple with the IssueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCounts

`func (o *GetSnippetDetails200ResponseSnippet) SetIssueCounts(v GetSnippets200ResponseResultsInnerIssueCounts)`

SetIssueCounts sets IssueCounts field to given value.


### GetRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippet) GetRejectionDetails() GetSnippets200ResponseResultsInnerRejectionDetails`

GetRejectionDetails returns the RejectionDetails field if non-nil, zero value otherwise.

### GetRejectionDetailsOk

`func (o *GetSnippetDetails200ResponseSnippet) GetRejectionDetailsOk() (*GetSnippets200ResponseResultsInnerRejectionDetails, bool)`

GetRejectionDetailsOk returns a tuple with the RejectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippet) SetRejectionDetails(v GetSnippets200ResponseResultsInnerRejectionDetails)`

SetRejectionDetails sets RejectionDetails field to given value.

### HasRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippet) HasRejectionDetails() bool`

HasRejectionDetails returns a boolean if a field has been set.

### GetLabels

`func (o *GetSnippetDetails200ResponseSnippet) GetLabels() []GetProjectDependencies200ResponseDependenciesInnerLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetSnippetDetails200ResponseSnippet) GetLabelsOk() (*[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetSnippetDetails200ResponseSnippet) SetLabels(v []GetProjectDependencies200ResponseDependenciesInnerLabelsInner)`

SetLabels sets Labels field to given value.


### GetOtherVersions

`func (o *GetSnippetDetails200ResponseSnippet) GetOtherVersions() []GetSnippetDetails200ResponseSnippetOtherVersionsInner`

GetOtherVersions returns the OtherVersions field if non-nil, zero value otherwise.

### GetOtherVersionsOk

`func (o *GetSnippetDetails200ResponseSnippet) GetOtherVersionsOk() (*[]GetSnippetDetails200ResponseSnippetOtherVersionsInner, bool)`

GetOtherVersionsOk returns a tuple with the OtherVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherVersions

`func (o *GetSnippetDetails200ResponseSnippet) SetOtherVersions(v []GetSnippetDetails200ResponseSnippetOtherVersionsInner)`

SetOtherVersions sets OtherVersions field to given value.

### HasOtherVersions

`func (o *GetSnippetDetails200ResponseSnippet) HasOtherVersions() bool`

HasOtherVersions returns a boolean if a field has been set.

### GetIsVendored

`func (o *GetSnippetDetails200ResponseSnippet) GetIsVendored() bool`

GetIsVendored returns the IsVendored field if non-nil, zero value otherwise.

### GetIsVendoredOk

`func (o *GetSnippetDetails200ResponseSnippet) GetIsVendoredOk() (*bool, bool)`

GetIsVendoredOk returns a tuple with the IsVendored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendored

`func (o *GetSnippetDetails200ResponseSnippet) SetIsVendored(v bool)`

SetIsVendored sets IsVendored field to given value.


### GetIsConverted

`func (o *GetSnippetDetails200ResponseSnippet) GetIsConverted() bool`

GetIsConverted returns the IsConverted field if non-nil, zero value otherwise.

### GetIsConvertedOk

`func (o *GetSnippetDetails200ResponseSnippet) GetIsConvertedOk() (*bool, bool)`

GetIsConvertedOk returns a tuple with the IsConverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConverted

`func (o *GetSnippetDetails200ResponseSnippet) SetIsConverted(v bool)`

SetIsConverted sets IsConverted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


