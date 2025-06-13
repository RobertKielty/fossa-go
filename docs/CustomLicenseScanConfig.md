# CustomLicenseScanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | custom license scan name | [optional] 
**MatchCriteria** | Pointer to **string** | regular expression for the custom license scan | [optional] 

## Methods

### NewCustomLicenseScanConfig

`func NewCustomLicenseScanConfig() *CustomLicenseScanConfig`

NewCustomLicenseScanConfig instantiates a new CustomLicenseScanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomLicenseScanConfigWithDefaults

`func NewCustomLicenseScanConfigWithDefaults() *CustomLicenseScanConfig`

NewCustomLicenseScanConfigWithDefaults instantiates a new CustomLicenseScanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CustomLicenseScanConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomLicenseScanConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomLicenseScanConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomLicenseScanConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatchCriteria

`func (o *CustomLicenseScanConfig) GetMatchCriteria() string`

GetMatchCriteria returns the MatchCriteria field if non-nil, zero value otherwise.

### GetMatchCriteriaOk

`func (o *CustomLicenseScanConfig) GetMatchCriteriaOk() (*string, bool)`

GetMatchCriteriaOk returns a tuple with the MatchCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCriteria

`func (o *CustomLicenseScanConfig) SetMatchCriteria(v string)`

SetMatchCriteria sets MatchCriteria field to given value.

### HasMatchCriteria

`func (o *CustomLicenseScanConfig) HasMatchCriteria() bool`

HasMatchCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


