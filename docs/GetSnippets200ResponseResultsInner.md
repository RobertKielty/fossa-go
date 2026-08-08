# GetSnippets200ResponseResultsInner

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
**HighestMatchPercentage** | **float32** | The highest match percentage across all matches | 
**ReleaseDate** | Pointer to **time.Time** | Release date of the package (optional) | [optional] 
**HomeUrl** | Pointer to **string** | Homepage URL of the package (optional) | [optional] 
**CodeUrl** | Pointer to **string** | Source code URL of the package (optional) | [optional] 
**Licenses** | [**[]GetSnippets200ResponseResultsInnerLicensesInner**](GetSnippets200ResponseResultsInnerLicensesInner.md) | Array of licenses associated with the snippet | 
**IssueCounts** | [**GetSnippets200ResponseResultsInnerIssueCounts**](GetSnippets200ResponseResultsInnerIssueCounts.md) |  | 
**RejectionDetails** | Pointer to [**GetSnippets200ResponseResultsInnerRejectionDetails**](GetSnippets200ResponseResultsInnerRejectionDetails.md) |  | [optional] 
**Labels** | [**[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner**](GetProjectDependencies200ResponseDependenciesInnerLabelsInner.md) | Package labels assigned to this snippet | 
**IsVendored** | **bool** | Whether the snippet exists as a vendored dependency | 
**IsConverted** | **bool** | Whether the snippet has been converted to a vendored dependency | 

## Methods

### NewGetSnippets200ResponseResultsInner

`func NewGetSnippets200ResponseResultsInner(id string, packageId string, purl string, locator string, package_ string, version string, kind string, matchCount int32, highestMatchPercentage float32, licenses []GetSnippets200ResponseResultsInnerLicensesInner, issueCounts GetSnippets200ResponseResultsInnerIssueCounts, labels []GetProjectDependencies200ResponseDependenciesInnerLabelsInner, isVendored bool, isConverted bool, ) *GetSnippets200ResponseResultsInner`

NewGetSnippets200ResponseResultsInner instantiates a new GetSnippets200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippets200ResponseResultsInnerWithDefaults

`func NewGetSnippets200ResponseResultsInnerWithDefaults() *GetSnippets200ResponseResultsInner`

NewGetSnippets200ResponseResultsInnerWithDefaults instantiates a new GetSnippets200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSnippets200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSnippets200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSnippets200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetPackageId

`func (o *GetSnippets200ResponseResultsInner) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *GetSnippets200ResponseResultsInner) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *GetSnippets200ResponseResultsInner) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetPurl

`func (o *GetSnippets200ResponseResultsInner) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *GetSnippets200ResponseResultsInner) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *GetSnippets200ResponseResultsInner) SetPurl(v string)`

SetPurl sets Purl field to given value.


### GetLocator

`func (o *GetSnippets200ResponseResultsInner) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *GetSnippets200ResponseResultsInner) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *GetSnippets200ResponseResultsInner) SetLocator(v string)`

SetLocator sets Locator field to given value.


### GetPackage

`func (o *GetSnippets200ResponseResultsInner) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *GetSnippets200ResponseResultsInner) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *GetSnippets200ResponseResultsInner) SetPackage(v string)`

SetPackage sets Package field to given value.


### GetVersion

`func (o *GetSnippets200ResponseResultsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetSnippets200ResponseResultsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetSnippets200ResponseResultsInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetKind

`func (o *GetSnippets200ResponseResultsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetSnippets200ResponseResultsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetSnippets200ResponseResultsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMatchCount

`func (o *GetSnippets200ResponseResultsInner) GetMatchCount() int32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *GetSnippets200ResponseResultsInner) GetMatchCountOk() (*int32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *GetSnippets200ResponseResultsInner) SetMatchCount(v int32)`

SetMatchCount sets MatchCount field to given value.


### GetHighestMatchPercentage

`func (o *GetSnippets200ResponseResultsInner) GetHighestMatchPercentage() float32`

GetHighestMatchPercentage returns the HighestMatchPercentage field if non-nil, zero value otherwise.

### GetHighestMatchPercentageOk

`func (o *GetSnippets200ResponseResultsInner) GetHighestMatchPercentageOk() (*float32, bool)`

GetHighestMatchPercentageOk returns a tuple with the HighestMatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestMatchPercentage

`func (o *GetSnippets200ResponseResultsInner) SetHighestMatchPercentage(v float32)`

SetHighestMatchPercentage sets HighestMatchPercentage field to given value.


### GetReleaseDate

`func (o *GetSnippets200ResponseResultsInner) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetSnippets200ResponseResultsInner) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetSnippets200ResponseResultsInner) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetSnippets200ResponseResultsInner) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetHomeUrl

`func (o *GetSnippets200ResponseResultsInner) GetHomeUrl() string`

GetHomeUrl returns the HomeUrl field if non-nil, zero value otherwise.

### GetHomeUrlOk

`func (o *GetSnippets200ResponseResultsInner) GetHomeUrlOk() (*string, bool)`

GetHomeUrlOk returns a tuple with the HomeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeUrl

`func (o *GetSnippets200ResponseResultsInner) SetHomeUrl(v string)`

SetHomeUrl sets HomeUrl field to given value.

### HasHomeUrl

`func (o *GetSnippets200ResponseResultsInner) HasHomeUrl() bool`

HasHomeUrl returns a boolean if a field has been set.

### GetCodeUrl

`func (o *GetSnippets200ResponseResultsInner) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *GetSnippets200ResponseResultsInner) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *GetSnippets200ResponseResultsInner) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.

### HasCodeUrl

`func (o *GetSnippets200ResponseResultsInner) HasCodeUrl() bool`

HasCodeUrl returns a boolean if a field has been set.

### GetLicenses

`func (o *GetSnippets200ResponseResultsInner) GetLicenses() []GetSnippets200ResponseResultsInnerLicensesInner`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *GetSnippets200ResponseResultsInner) GetLicensesOk() (*[]GetSnippets200ResponseResultsInnerLicensesInner, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *GetSnippets200ResponseResultsInner) SetLicenses(v []GetSnippets200ResponseResultsInnerLicensesInner)`

SetLicenses sets Licenses field to given value.


### GetIssueCounts

`func (o *GetSnippets200ResponseResultsInner) GetIssueCounts() GetSnippets200ResponseResultsInnerIssueCounts`

GetIssueCounts returns the IssueCounts field if non-nil, zero value otherwise.

### GetIssueCountsOk

`func (o *GetSnippets200ResponseResultsInner) GetIssueCountsOk() (*GetSnippets200ResponseResultsInnerIssueCounts, bool)`

GetIssueCountsOk returns a tuple with the IssueCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCounts

`func (o *GetSnippets200ResponseResultsInner) SetIssueCounts(v GetSnippets200ResponseResultsInnerIssueCounts)`

SetIssueCounts sets IssueCounts field to given value.


### GetRejectionDetails

`func (o *GetSnippets200ResponseResultsInner) GetRejectionDetails() GetSnippets200ResponseResultsInnerRejectionDetails`

GetRejectionDetails returns the RejectionDetails field if non-nil, zero value otherwise.

### GetRejectionDetailsOk

`func (o *GetSnippets200ResponseResultsInner) GetRejectionDetailsOk() (*GetSnippets200ResponseResultsInnerRejectionDetails, bool)`

GetRejectionDetailsOk returns a tuple with the RejectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionDetails

`func (o *GetSnippets200ResponseResultsInner) SetRejectionDetails(v GetSnippets200ResponseResultsInnerRejectionDetails)`

SetRejectionDetails sets RejectionDetails field to given value.

### HasRejectionDetails

`func (o *GetSnippets200ResponseResultsInner) HasRejectionDetails() bool`

HasRejectionDetails returns a boolean if a field has been set.

### GetLabels

`func (o *GetSnippets200ResponseResultsInner) GetLabels() []GetProjectDependencies200ResponseDependenciesInnerLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetSnippets200ResponseResultsInner) GetLabelsOk() (*[]GetProjectDependencies200ResponseDependenciesInnerLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetSnippets200ResponseResultsInner) SetLabels(v []GetProjectDependencies200ResponseDependenciesInnerLabelsInner)`

SetLabels sets Labels field to given value.


### GetIsVendored

`func (o *GetSnippets200ResponseResultsInner) GetIsVendored() bool`

GetIsVendored returns the IsVendored field if non-nil, zero value otherwise.

### GetIsVendoredOk

`func (o *GetSnippets200ResponseResultsInner) GetIsVendoredOk() (*bool, bool)`

GetIsVendoredOk returns a tuple with the IsVendored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendored

`func (o *GetSnippets200ResponseResultsInner) SetIsVendored(v bool)`

SetIsVendored sets IsVendored field to given value.


### GetIsConverted

`func (o *GetSnippets200ResponseResultsInner) GetIsConverted() bool`

GetIsConverted returns the IsConverted field if non-nil, zero value otherwise.

### GetIsConvertedOk

`func (o *GetSnippets200ResponseResultsInner) GetIsConvertedOk() (*bool, bool)`

GetIsConvertedOk returns a tuple with the IsConverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConverted

`func (o *GetSnippets200ResponseResultsInner) SetIsConverted(v bool)`

SetIsConverted sets IsConverted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


