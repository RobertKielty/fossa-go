# GetIssue200ResponseOneOfAllOfRemediation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartialFix** | Pointer to **string** |  | [optional] 
**PartialFixDistance** | Pointer to **string** |  | [optional] 
**CompleteFix** | Pointer to **string** |  | [optional] 
**CompleteFixDistance** | Pointer to **string** |  | [optional] 
**PartialFixResolves** | Pointer to [**[]GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner**](GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner.md) | Vulnerabilities on this dependency revision that upgrading to partialFix resolves. Absent when the fix could not be calculated. | [optional] 
**CompleteFixResolves** | Pointer to [**[]GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner**](GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner.md) | Vulnerabilities on this dependency revision that upgrading to completeFix resolves. Absent when the fix could not be calculated. | [optional] 

## Methods

### NewGetIssue200ResponseOneOfAllOfRemediation

`func NewGetIssue200ResponseOneOfAllOfRemediation() *GetIssue200ResponseOneOfAllOfRemediation`

NewGetIssue200ResponseOneOfAllOfRemediation instantiates a new GetIssue200ResponseOneOfAllOfRemediation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssue200ResponseOneOfAllOfRemediationWithDefaults

`func NewGetIssue200ResponseOneOfAllOfRemediationWithDefaults() *GetIssue200ResponseOneOfAllOfRemediation`

NewGetIssue200ResponseOneOfAllOfRemediationWithDefaults instantiates a new GetIssue200ResponseOneOfAllOfRemediation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartialFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFix() string`

GetPartialFix returns the PartialFix field if non-nil, zero value otherwise.

### GetPartialFixOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFixOk() (*string, bool)`

GetPartialFixOk returns a tuple with the PartialFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetPartialFix(v string)`

SetPartialFix sets PartialFix field to given value.

### HasPartialFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasPartialFix() bool`

HasPartialFix returns a boolean if a field has been set.

### GetPartialFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFixDistance() string`

GetPartialFixDistance returns the PartialFixDistance field if non-nil, zero value otherwise.

### GetPartialFixDistanceOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFixDistanceOk() (*string, bool)`

GetPartialFixDistanceOk returns a tuple with the PartialFixDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetPartialFixDistance(v string)`

SetPartialFixDistance sets PartialFixDistance field to given value.

### HasPartialFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasPartialFixDistance() bool`

HasPartialFixDistance returns a boolean if a field has been set.

### GetCompleteFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFix() string`

GetCompleteFix returns the CompleteFix field if non-nil, zero value otherwise.

### GetCompleteFixOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFixOk() (*string, bool)`

GetCompleteFixOk returns a tuple with the CompleteFix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetCompleteFix(v string)`

SetCompleteFix sets CompleteFix field to given value.

### HasCompleteFix

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasCompleteFix() bool`

HasCompleteFix returns a boolean if a field has been set.

### GetCompleteFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFixDistance() string`

GetCompleteFixDistance returns the CompleteFixDistance field if non-nil, zero value otherwise.

### GetCompleteFixDistanceOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFixDistanceOk() (*string, bool)`

GetCompleteFixDistanceOk returns a tuple with the CompleteFixDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetCompleteFixDistance(v string)`

SetCompleteFixDistance sets CompleteFixDistance field to given value.

### HasCompleteFixDistance

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasCompleteFixDistance() bool`

HasCompleteFixDistance returns a boolean if a field has been set.

### GetPartialFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFixResolves() []GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner`

GetPartialFixResolves returns the PartialFixResolves field if non-nil, zero value otherwise.

### GetPartialFixResolvesOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetPartialFixResolvesOk() (*[]GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner, bool)`

GetPartialFixResolvesOk returns a tuple with the PartialFixResolves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetPartialFixResolves(v []GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner)`

SetPartialFixResolves sets PartialFixResolves field to given value.

### HasPartialFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasPartialFixResolves() bool`

HasPartialFixResolves returns a boolean if a field has been set.

### GetCompleteFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFixResolves() []GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner`

GetCompleteFixResolves returns the CompleteFixResolves field if non-nil, zero value otherwise.

### GetCompleteFixResolvesOk

`func (o *GetIssue200ResponseOneOfAllOfRemediation) GetCompleteFixResolvesOk() (*[]GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner, bool)`

GetCompleteFixResolvesOk returns a tuple with the CompleteFixResolves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) SetCompleteFixResolves(v []GetIssue200ResponseOneOfAllOfRemediationPartialFixResolvesInner)`

SetCompleteFixResolves sets CompleteFixResolves field to given value.

### HasCompleteFixResolves

`func (o *GetIssue200ResponseOneOfAllOfRemediation) HasCompleteFixResolves() bool`

HasCompleteFixResolves returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


