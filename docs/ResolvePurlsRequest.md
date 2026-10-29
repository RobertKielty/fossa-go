# ResolvePurlsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purls** | **[]string** | The PURLs (package URLs) to resolve. Between 1 and 100 entries. | 

## Methods

### NewResolvePurlsRequest

`func NewResolvePurlsRequest(purls []string, ) *ResolvePurlsRequest`

NewResolvePurlsRequest instantiates a new ResolvePurlsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResolvePurlsRequestWithDefaults

`func NewResolvePurlsRequestWithDefaults() *ResolvePurlsRequest`

NewResolvePurlsRequestWithDefaults instantiates a new ResolvePurlsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurls

`func (o *ResolvePurlsRequest) GetPurls() []string`

GetPurls returns the Purls field if non-nil, zero value otherwise.

### GetPurlsOk

`func (o *ResolvePurlsRequest) GetPurlsOk() (*[]string, bool)`

GetPurlsOk returns a tuple with the Purls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurls

`func (o *ResolvePurlsRequest) SetPurls(v []string)`

SetPurls sets Purls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


