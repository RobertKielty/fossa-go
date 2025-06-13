# GetIssueCounts200ResponseCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **int32** |  | [optional] 
**Ignored** | Pointer to **int32** |  | [optional] 
**Remediated** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetIssueCounts200ResponseCounts

`func NewGetIssueCounts200ResponseCounts() *GetIssueCounts200ResponseCounts`

NewGetIssueCounts200ResponseCounts instantiates a new GetIssueCounts200ResponseCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueCounts200ResponseCountsWithDefaults

`func NewGetIssueCounts200ResponseCountsWithDefaults() *GetIssueCounts200ResponseCounts`

NewGetIssueCounts200ResponseCountsWithDefaults instantiates a new GetIssueCounts200ResponseCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetIssueCounts200ResponseCounts) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetIssueCounts200ResponseCounts) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetIssueCounts200ResponseCounts) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetIssueCounts200ResponseCounts) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetActive

`func (o *GetIssueCounts200ResponseCounts) GetActive() int32`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetIssueCounts200ResponseCounts) GetActiveOk() (*int32, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetIssueCounts200ResponseCounts) SetActive(v int32)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetIssueCounts200ResponseCounts) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIgnored

`func (o *GetIssueCounts200ResponseCounts) GetIgnored() int32`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *GetIssueCounts200ResponseCounts) GetIgnoredOk() (*int32, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *GetIssueCounts200ResponseCounts) SetIgnored(v int32)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *GetIssueCounts200ResponseCounts) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetRemediated

`func (o *GetIssueCounts200ResponseCounts) GetRemediated() int32`

GetRemediated returns the Remediated field if non-nil, zero value otherwise.

### GetRemediatedOk

`func (o *GetIssueCounts200ResponseCounts) GetRemediatedOk() (*int32, bool)`

GetRemediatedOk returns a tuple with the Remediated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediated

`func (o *GetIssueCounts200ResponseCounts) SetRemediated(v int32)`

SetRemediated sets Remediated field to given value.

### HasRemediated

`func (o *GetIssueCounts200ResponseCounts) HasRemediated() bool`

HasRemediated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


