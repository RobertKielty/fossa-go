# BuildRequestArchives

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageSpec** | Pointer to **string** | The unresolved spec (i.e. a URL) that can be passed to a fetcher and resolved to a package id. For example: underscore (npm) https://github.com/fossas/fossa (git) nokogiri (gem)  | [optional] 
**Revision** | Pointer to **string** | The branch or revision of the component being uploaded. | [optional] 
**Description** | Pointer to **string** | The description of the archive. | [optional] 
**ProjectURL** | Pointer to **string** | The homepage of the archive. | [optional] 

## Methods

### NewBuildRequestArchives

`func NewBuildRequestArchives() *BuildRequestArchives`

NewBuildRequestArchives instantiates a new BuildRequestArchives object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildRequestArchivesWithDefaults

`func NewBuildRequestArchivesWithDefaults() *BuildRequestArchives`

NewBuildRequestArchivesWithDefaults instantiates a new BuildRequestArchives object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageSpec

`func (o *BuildRequestArchives) GetPackageSpec() string`

GetPackageSpec returns the PackageSpec field if non-nil, zero value otherwise.

### GetPackageSpecOk

`func (o *BuildRequestArchives) GetPackageSpecOk() (*string, bool)`

GetPackageSpecOk returns a tuple with the PackageSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageSpec

`func (o *BuildRequestArchives) SetPackageSpec(v string)`

SetPackageSpec sets PackageSpec field to given value.

### HasPackageSpec

`func (o *BuildRequestArchives) HasPackageSpec() bool`

HasPackageSpec returns a boolean if a field has been set.

### GetRevision

`func (o *BuildRequestArchives) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *BuildRequestArchives) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *BuildRequestArchives) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *BuildRequestArchives) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetDescription

`func (o *BuildRequestArchives) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BuildRequestArchives) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BuildRequestArchives) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BuildRequestArchives) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectURL

`func (o *BuildRequestArchives) GetProjectURL() string`

GetProjectURL returns the ProjectURL field if non-nil, zero value otherwise.

### GetProjectURLOk

`func (o *BuildRequestArchives) GetProjectURLOk() (*string, bool)`

GetProjectURLOk returns a tuple with the ProjectURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectURL

`func (o *BuildRequestArchives) SetProjectURL(v string)`

SetProjectURL sets ProjectURL field to given value.

### HasProjectURL

`func (o *BuildRequestArchives) HasProjectURL() bool`

HasProjectURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


