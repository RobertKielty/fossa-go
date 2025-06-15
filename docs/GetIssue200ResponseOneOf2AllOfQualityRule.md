# GetIssue200ResponseOneOf2AllOfQualityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**DependencyOutdatedType** | Pointer to **NullableString** |  | [optional] 
**DependencyOutdatedVersionPart** | Pointer to **NullableString** |  | [optional] 
**DependencyOutdatedVersionDiff** | Pointer to **NullableString** |  | [optional] 
**DependencyFilterName** | Pointer to **NullableString** |  | [optional] 
**DependencyFilterVersionType** | Pointer to **NullableString** |  | [optional] 
**DependencyFilterStartVersion** | Pointer to **string** |  | [optional] 
**DependencyFilterEndVersion** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**PackageProjectLocator** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetIssue200ResponseOneOf2AllOfQualityRule

`func NewGetIssue200ResponseOneOf2AllOfQualityRule() *GetIssue200ResponseOneOf2AllOfQualityRule`

NewGetIssue200ResponseOneOf2AllOfQualityRule instantiates a new GetIssue200ResponseOneOf2AllOfQualityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseOneOf2AllOfQualityRuleWithDefaults

`func NewGetIssue200ResponseOneOf2AllOfQualityRuleWithDefaults() *GetIssue200ResponseOneOf2AllOfQualityRule`

NewGetIssue200ResponseOneOf2AllOfQualityRuleWithDefaults instantiates a new GetIssue200ResponseOneOf2AllOfQualityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetRuleId() int32`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetRuleIdOk() (*int32, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetRuleId(v int32)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNotes

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDependencyOutdatedType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedType() string`

GetDependencyOutdatedType returns the DependencyOutdatedType field if non-nil, zero value otherwise.

### GetDependencyOutdatedTypeOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedTypeOk() (*string, bool)`

GetDependencyOutdatedTypeOk returns a tuple with the DependencyOutdatedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyOutdatedType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedType(v string)`

SetDependencyOutdatedType sets DependencyOutdatedType field to given value.

### HasDependencyOutdatedType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyOutdatedType() bool`

HasDependencyOutdatedType returns a boolean if a field has been set.

### SetDependencyOutdatedTypeNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedTypeNil(b bool)`

 SetDependencyOutdatedTypeNil sets the value for DependencyOutdatedType to be an explicit nil

### UnsetDependencyOutdatedType
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyOutdatedType()`

UnsetDependencyOutdatedType ensures that no value is present for DependencyOutdatedType, not even an explicit nil
### GetDependencyOutdatedVersionPart

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedVersionPart() string`

GetDependencyOutdatedVersionPart returns the DependencyOutdatedVersionPart field if non-nil, zero value otherwise.

### GetDependencyOutdatedVersionPartOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedVersionPartOk() (*string, bool)`

GetDependencyOutdatedVersionPartOk returns a tuple with the DependencyOutdatedVersionPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyOutdatedVersionPart

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedVersionPart(v string)`

SetDependencyOutdatedVersionPart sets DependencyOutdatedVersionPart field to given value.

### HasDependencyOutdatedVersionPart

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyOutdatedVersionPart() bool`

HasDependencyOutdatedVersionPart returns a boolean if a field has been set.

### SetDependencyOutdatedVersionPartNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedVersionPartNil(b bool)`

 SetDependencyOutdatedVersionPartNil sets the value for DependencyOutdatedVersionPart to be an explicit nil

### UnsetDependencyOutdatedVersionPart
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyOutdatedVersionPart()`

UnsetDependencyOutdatedVersionPart ensures that no value is present for DependencyOutdatedVersionPart, not even an explicit nil
### GetDependencyOutdatedVersionDiff

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedVersionDiff() string`

GetDependencyOutdatedVersionDiff returns the DependencyOutdatedVersionDiff field if non-nil, zero value otherwise.

### GetDependencyOutdatedVersionDiffOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyOutdatedVersionDiffOk() (*string, bool)`

GetDependencyOutdatedVersionDiffOk returns a tuple with the DependencyOutdatedVersionDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyOutdatedVersionDiff

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedVersionDiff(v string)`

SetDependencyOutdatedVersionDiff sets DependencyOutdatedVersionDiff field to given value.

### HasDependencyOutdatedVersionDiff

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyOutdatedVersionDiff() bool`

HasDependencyOutdatedVersionDiff returns a boolean if a field has been set.

### SetDependencyOutdatedVersionDiffNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyOutdatedVersionDiffNil(b bool)`

 SetDependencyOutdatedVersionDiffNil sets the value for DependencyOutdatedVersionDiff to be an explicit nil

### UnsetDependencyOutdatedVersionDiff
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyOutdatedVersionDiff()`

UnsetDependencyOutdatedVersionDiff ensures that no value is present for DependencyOutdatedVersionDiff, not even an explicit nil
### GetDependencyFilterName

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterName() string`

GetDependencyFilterName returns the DependencyFilterName field if non-nil, zero value otherwise.

### GetDependencyFilterNameOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterNameOk() (*string, bool)`

GetDependencyFilterNameOk returns a tuple with the DependencyFilterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyFilterName

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterName(v string)`

SetDependencyFilterName sets DependencyFilterName field to given value.

### HasDependencyFilterName

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyFilterName() bool`

HasDependencyFilterName returns a boolean if a field has been set.

### SetDependencyFilterNameNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterNameNil(b bool)`

 SetDependencyFilterNameNil sets the value for DependencyFilterName to be an explicit nil

### UnsetDependencyFilterName
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyFilterName()`

UnsetDependencyFilterName ensures that no value is present for DependencyFilterName, not even an explicit nil
### GetDependencyFilterVersionType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterVersionType() string`

GetDependencyFilterVersionType returns the DependencyFilterVersionType field if non-nil, zero value otherwise.

### GetDependencyFilterVersionTypeOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterVersionTypeOk() (*string, bool)`

GetDependencyFilterVersionTypeOk returns a tuple with the DependencyFilterVersionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyFilterVersionType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterVersionType(v string)`

SetDependencyFilterVersionType sets DependencyFilterVersionType field to given value.

### HasDependencyFilterVersionType

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyFilterVersionType() bool`

HasDependencyFilterVersionType returns a boolean if a field has been set.

### SetDependencyFilterVersionTypeNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterVersionTypeNil(b bool)`

 SetDependencyFilterVersionTypeNil sets the value for DependencyFilterVersionType to be an explicit nil

### UnsetDependencyFilterVersionType
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyFilterVersionType()`

UnsetDependencyFilterVersionType ensures that no value is present for DependencyFilterVersionType, not even an explicit nil
### GetDependencyFilterStartVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterStartVersion() string`

GetDependencyFilterStartVersion returns the DependencyFilterStartVersion field if non-nil, zero value otherwise.

### GetDependencyFilterStartVersionOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterStartVersionOk() (*string, bool)`

GetDependencyFilterStartVersionOk returns a tuple with the DependencyFilterStartVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyFilterStartVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterStartVersion(v string)`

SetDependencyFilterStartVersion sets DependencyFilterStartVersion field to given value.

### HasDependencyFilterStartVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyFilterStartVersion() bool`

HasDependencyFilterStartVersion returns a boolean if a field has been set.

### GetDependencyFilterEndVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterEndVersion() string`

GetDependencyFilterEndVersion returns the DependencyFilterEndVersion field if non-nil, zero value otherwise.

### GetDependencyFilterEndVersionOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetDependencyFilterEndVersionOk() (*string, bool)`

GetDependencyFilterEndVersionOk returns a tuple with the DependencyFilterEndVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyFilterEndVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterEndVersion(v string)`

SetDependencyFilterEndVersion sets DependencyFilterEndVersion field to given value.

### HasDependencyFilterEndVersion

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasDependencyFilterEndVersion() bool`

HasDependencyFilterEndVersion returns a boolean if a field has been set.

### SetDependencyFilterEndVersionNil

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetDependencyFilterEndVersionNil(b bool)`

 SetDependencyFilterEndVersionNil sets the value for DependencyFilterEndVersion to be an explicit nil

### UnsetDependencyFilterEndVersion
`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) UnsetDependencyFilterEndVersion()`

UnsetDependencyFilterEndVersion ensures that no value is present for DependencyFilterEndVersion, not even an explicit nil
### GetEnabled

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPackageProjectLocator

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetPackageProjectLocator() string`

GetPackageProjectLocator returns the PackageProjectLocator field if non-nil, zero value otherwise.

### GetPackageProjectLocatorOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetPackageProjectLocatorOk() (*string, bool)`

GetPackageProjectLocatorOk returns a tuple with the PackageProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageProjectLocator

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetPackageProjectLocator(v string)`

SetPackageProjectLocator sets PackageProjectLocator field to given value.

### HasPackageProjectLocator

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasPackageProjectLocator() bool`

HasPackageProjectLocator returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetIssue200ResponseOneOf2AllOfQualityRule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


