# GetRevisionComponentsPaths200ResponsePathsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of component path node (directory or file) | [optional] 
**Name** | Pointer to **string** | Name of file | [optional] 
**Path** | Pointer to **string** | Path to file | [optional] 
**Count** | Pointer to **float32** | Count of component matches in directory | [optional] 
**RevisionId** | Pointer to **string** | Dependency revision matched in file | [optional] 
**ComponentId** | Pointer to **string** | Component Id for dependency and file | [optional] 
**PackageName** | Pointer to **string** | Package name of Dependency matched in file | [optional] 
**Version** | Pointer to **string** | Version of Dependency matched in file | [optional] 
**RootProjectRevisionLocator** | Pointer to **string** | Revision of root Binary Decomposition project | [optional] 

## Methods

### NewGetRevisionComponentsPaths200ResponsePathsInner

`func NewGetRevisionComponentsPaths200ResponsePathsInner() *GetRevisionComponentsPaths200ResponsePathsInner`

NewGetRevisionComponentsPaths200ResponsePathsInner instantiates a new GetRevisionComponentsPaths200ResponsePathsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionComponentsPaths200ResponsePathsInnerWithDefaults

`func NewGetRevisionComponentsPaths200ResponsePathsInnerWithDefaults() *GetRevisionComponentsPaths200ResponsePathsInner`

NewGetRevisionComponentsPaths200ResponsePathsInnerWithDefaults instantiates a new GetRevisionComponentsPaths200ResponsePathsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetCount

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRevisionId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetComponentId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetPackageName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetVersion

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetRootProjectRevisionLocator

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetRootProjectRevisionLocator() string`

GetRootProjectRevisionLocator returns the RootProjectRevisionLocator field if non-nil, zero value otherwise.

### GetRootProjectRevisionLocatorOk

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) GetRootProjectRevisionLocatorOk() (*string, bool)`

GetRootProjectRevisionLocatorOk returns a tuple with the RootProjectRevisionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootProjectRevisionLocator

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) SetRootProjectRevisionLocator(v string)`

SetRootProjectRevisionLocator sets RootProjectRevisionLocator field to given value.

### HasRootProjectRevisionLocator

`func (o *GetRevisionComponentsPaths200ResponsePathsInner) HasRootProjectRevisionLocator() bool`

HasRootProjectRevisionLocator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


