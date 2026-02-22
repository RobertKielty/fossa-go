# AddLicenseConclusionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DependencyRevisionLocator** | **string** | The locator of the dependency revision | 
**Scope** | [**AddLicenseConclusionRequestScope**](AddLicenseConclusionRequestScope.md) |  | 
**LicenseId** | **string** | The license ID to add or remove | 
**OriginId** | Pointer to **string** | Optional origin revision or release ID to trigger an issue scan | [optional] 

## Methods

### NewAddLicenseConclusionRequest

`func NewAddLicenseConclusionRequest(dependencyRevisionLocator string, scope AddLicenseConclusionRequestScope, licenseId string, ) *AddLicenseConclusionRequest`

NewAddLicenseConclusionRequest instantiates a new AddLicenseConclusionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLicenseConclusionRequestWithDefaults

`func NewAddLicenseConclusionRequestWithDefaults() *AddLicenseConclusionRequest`

NewAddLicenseConclusionRequestWithDefaults instantiates a new AddLicenseConclusionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependencyRevisionLocator

`func (o *AddLicenseConclusionRequest) GetDependencyRevisionLocator() string`

GetDependencyRevisionLocator returns the DependencyRevisionLocator field if non-nil, zero value otherwise.

### GetDependencyRevisionLocatorOk

`func (o *AddLicenseConclusionRequest) GetDependencyRevisionLocatorOk() (*string, bool)`

GetDependencyRevisionLocatorOk returns a tuple with the DependencyRevisionLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyRevisionLocator

`func (o *AddLicenseConclusionRequest) SetDependencyRevisionLocator(v string)`

SetDependencyRevisionLocator sets DependencyRevisionLocator field to given value.


### GetScope

`func (o *AddLicenseConclusionRequest) GetScope() AddLicenseConclusionRequestScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddLicenseConclusionRequest) GetScopeOk() (*AddLicenseConclusionRequestScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddLicenseConclusionRequest) SetScope(v AddLicenseConclusionRequestScope)`

SetScope sets Scope field to given value.


### GetLicenseId

`func (o *AddLicenseConclusionRequest) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *AddLicenseConclusionRequest) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *AddLicenseConclusionRequest) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetOriginId

`func (o *AddLicenseConclusionRequest) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *AddLicenseConclusionRequest) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *AddLicenseConclusionRequest) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.

### HasOriginId

`func (o *AddLicenseConclusionRequest) HasOriginId() bool`

HasOriginId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


