# GetSnippets200ResponseResultsInnerRejectionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RejectedAt** | **time.Time** | Timestamp when the snippet was rejected | 
**RejectedBy** | Pointer to **string** | User who rejected the snippet (optional) | [optional] 

## Methods

### NewGetSnippets200ResponseResultsInnerRejectionDetails

`func NewGetSnippets200ResponseResultsInnerRejectionDetails(rejectedAt time.Time, ) *GetSnippets200ResponseResultsInnerRejectionDetails`

NewGetSnippets200ResponseResultsInnerRejectionDetails instantiates a new GetSnippets200ResponseResultsInnerRejectionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippets200ResponseResultsInnerRejectionDetailsWithDefaults

`func NewGetSnippets200ResponseResultsInnerRejectionDetailsWithDefaults() *GetSnippets200ResponseResultsInnerRejectionDetails`

NewGetSnippets200ResponseResultsInnerRejectionDetailsWithDefaults instantiates a new GetSnippets200ResponseResultsInnerRejectionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRejectedAt

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) GetRejectedAt() time.Time`

GetRejectedAt returns the RejectedAt field if non-nil, zero value otherwise.

### GetRejectedAtOk

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) GetRejectedAtOk() (*time.Time, bool)`

GetRejectedAtOk returns a tuple with the RejectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedAt

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) SetRejectedAt(v time.Time)`

SetRejectedAt sets RejectedAt field to given value.


### GetRejectedBy

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) GetRejectedBy() string`

GetRejectedBy returns the RejectedBy field if non-nil, zero value otherwise.

### GetRejectedByOk

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) GetRejectedByOk() (*string, bool)`

GetRejectedByOk returns a tuple with the RejectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedBy

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) SetRejectedBy(v string)`

SetRejectedBy sets RejectedBy field to given value.

### HasRejectedBy

`func (o *GetSnippets200ResponseResultsInnerRejectionDetails) HasRejectedBy() bool`

HasRejectedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


