# AddLicenseConclusion201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyRevisionLocator** | **string** | The locator of the dependency revision | 
**Scope** | **string** | The scope of the license conclusion | 
**ScopeId** | **string** | The ID of the scope (project locator, revision locator, release group ID, etc.) | 
**ConcludedLicenses** | **[]string** | The list of concluded licenses for this dependency | 
**CreatedAt** | **time.Time** | The date and time the conclusion was created | 
**UpdatedAt** | **time.Time** | The date and time the conclusion was last updated | 

## Methods

### NewAddLicenseConclusion201Response

`func NewAddLicenseConclusion201Response(dependencyRevisionLocator string, scope string, scopeId string, concludedLicenses []string, createdAt time.Time, updatedAt time.Time, ) *AddLicenseConclusion201Response`

NewAddLicenseConclusion201Response instantiates a new AddLicenseConclusion201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLicenseConclusion201ResponseWithDefaults

`func NewAddLicenseConclusion201ResponseWithDefaults() *AddLicenseConclusion201Response`

NewAddLicenseConclusion201ResponseWithDefaults instantiates a new AddLicenseConclusion201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyRevisionLocator

`func (o *AddLicenseConclusion201Response) GetDependencyRevisionLocator() string`

GetDependencyRevisionLocator returns the DependencyRevisionLocator field if non-nil, zero value otherwise.

### GetDependencyRevisionLocatorOk

`func (o *AddLicenseConclusion201Response) GetDependencyRevisionLocatorOk() (*string, bool)`

GetDependencyRevisionLocatorOk returns a tuple with the DependencyRevisionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyRevisionLocator

`func (o *AddLicenseConclusion201Response) SetDependencyRevisionLocator(v string)`

SetDependencyRevisionLocator sets DependencyRevisionLocator field to given value.


### GetScope

`func (o *AddLicenseConclusion201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddLicenseConclusion201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddLicenseConclusion201Response) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeId

`func (o *AddLicenseConclusion201Response) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *AddLicenseConclusion201Response) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *AddLicenseConclusion201Response) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.


### GetConcludedLicenses

`func (o *AddLicenseConclusion201Response) GetConcludedLicenses() []string`

GetConcludedLicenses returns the ConcludedLicenses field if non-nil, zero value otherwise.

### GetConcludedLicensesOk

`func (o *AddLicenseConclusion201Response) GetConcludedLicensesOk() (*[]string, bool)`

GetConcludedLicensesOk returns a tuple with the ConcludedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcludedLicenses

`func (o *AddLicenseConclusion201Response) SetConcludedLicenses(v []string)`

SetConcludedLicenses sets ConcludedLicenses field to given value.


### GetCreatedAt

`func (o *AddLicenseConclusion201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AddLicenseConclusion201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AddLicenseConclusion201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AddLicenseConclusion201Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AddLicenseConclusion201Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AddLicenseConclusion201Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


