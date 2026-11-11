# GetRevisionDependenciesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Maximum number of dependencies to return. The value is clamped server-side to the range 25–100: any value below 25 is treated as 25, and any value above 100 is treated as 100.  | [optional] 
**Offset** | Pointer to **int32** | Number of dependencies to skip for pagination | [optional] 
**IncludeIgnored** | Pointer to **bool** | Whether to include ignored dependencies in the response | [optional] [default to false]
**IncludeHashData** | Pointer to **bool** | Whether to include hash and version data for dependencies | [optional] [default to false]
**IncludeLicenseText** | Pointer to **bool** | Whether to include full license text in the license information | [optional] [default to false]
**IncludeLocators** | Pointer to **[]string** | Array of locators to filter dependencies. Only dependencies matching these locators will be returned | [optional] 

## Methods

### NewGetRevisionDependenciesPostRequest

`func NewGetRevisionDependenciesPostRequest() *GetRevisionDependenciesPostRequest`

NewGetRevisionDependenciesPostRequest instantiates a new GetRevisionDependenciesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionDependenciesPostRequestWithDefaults

`func NewGetRevisionDependenciesPostRequestWithDefaults() *GetRevisionDependenciesPostRequest`

NewGetRevisionDependenciesPostRequestWithDefaults instantiates a new GetRevisionDependenciesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *GetRevisionDependenciesPostRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetRevisionDependenciesPostRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetRevisionDependenciesPostRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetRevisionDependenciesPostRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *GetRevisionDependenciesPostRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetRevisionDependenciesPostRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetRevisionDependenciesPostRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetRevisionDependenciesPostRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetIncludeIgnored

`func (o *GetRevisionDependenciesPostRequest) GetIncludeIgnored() bool`

GetIncludeIgnored returns the IncludeIgnored field if non-nil, zero value otherwise.

### GetIncludeIgnoredOk

`func (o *GetRevisionDependenciesPostRequest) GetIncludeIgnoredOk() (*bool, bool)`

GetIncludeIgnoredOk returns a tuple with the IncludeIgnored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeIgnored

`func (o *GetRevisionDependenciesPostRequest) SetIncludeIgnored(v bool)`

SetIncludeIgnored sets IncludeIgnored field to given value.

### HasIncludeIgnored

`func (o *GetRevisionDependenciesPostRequest) HasIncludeIgnored() bool`

HasIncludeIgnored returns a boolean if a field has been set.

### GetIncludeHashData

`func (o *GetRevisionDependenciesPostRequest) GetIncludeHashData() bool`

GetIncludeHashData returns the IncludeHashData field if non-nil, zero value otherwise.

### GetIncludeHashDataOk

`func (o *GetRevisionDependenciesPostRequest) GetIncludeHashDataOk() (*bool, bool)`

GetIncludeHashDataOk returns a tuple with the IncludeHashData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHashData

`func (o *GetRevisionDependenciesPostRequest) SetIncludeHashData(v bool)`

SetIncludeHashData sets IncludeHashData field to given value.

### HasIncludeHashData

`func (o *GetRevisionDependenciesPostRequest) HasIncludeHashData() bool`

HasIncludeHashData returns a boolean if a field has been set.

### GetIncludeLicenseText

`func (o *GetRevisionDependenciesPostRequest) GetIncludeLicenseText() bool`

GetIncludeLicenseText returns the IncludeLicenseText field if non-nil, zero value otherwise.

### GetIncludeLicenseTextOk

`func (o *GetRevisionDependenciesPostRequest) GetIncludeLicenseTextOk() (*bool, bool)`

GetIncludeLicenseTextOk returns a tuple with the IncludeLicenseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLicenseText

`func (o *GetRevisionDependenciesPostRequest) SetIncludeLicenseText(v bool)`

SetIncludeLicenseText sets IncludeLicenseText field to given value.

### HasIncludeLicenseText

`func (o *GetRevisionDependenciesPostRequest) HasIncludeLicenseText() bool`

HasIncludeLicenseText returns a boolean if a field has been set.

### GetIncludeLocators

`func (o *GetRevisionDependenciesPostRequest) GetIncludeLocators() []string`

GetIncludeLocators returns the IncludeLocators field if non-nil, zero value otherwise.

### GetIncludeLocatorsOk

`func (o *GetRevisionDependenciesPostRequest) GetIncludeLocatorsOk() (*[]string, bool)`

GetIncludeLocatorsOk returns a tuple with the IncludeLocators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeLocators

`func (o *GetRevisionDependenciesPostRequest) SetIncludeLocators(v []string)`

SetIncludeLocators sets IncludeLocators field to given value.

### HasIncludeLocators

`func (o *GetRevisionDependenciesPostRequest) HasIncludeLocators() bool`

HasIncludeLocators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


