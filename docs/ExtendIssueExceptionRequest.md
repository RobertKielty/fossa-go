# ExtendIssueExceptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAfter** | **string** | Expiration date in YYYY-MM-DD format, or null to never expire | 

## Methods

### NewExtendIssueExceptionRequest

`func NewExtendIssueExceptionRequest(expiresAfter string, ) *ExtendIssueExceptionRequest`

NewExtendIssueExceptionRequest instantiates a new ExtendIssueExceptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendIssueExceptionRequestWithDefaults

`func NewExtendIssueExceptionRequestWithDefaults() *ExtendIssueExceptionRequest`

NewExtendIssueExceptionRequestWithDefaults instantiates a new ExtendIssueExceptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAfter

`func (o *ExtendIssueExceptionRequest) GetExpiresAfter() string`

GetExpiresAfter returns the ExpiresAfter field if non-nil, zero value otherwise.

### GetExpiresAfterOk

`func (o *ExtendIssueExceptionRequest) GetExpiresAfterOk() (*string, bool)`

GetExpiresAfterOk returns a tuple with the ExpiresAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfter

`func (o *ExtendIssueExceptionRequest) SetExpiresAfter(v string)`

SetExpiresAfter sets ExpiresAfter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


