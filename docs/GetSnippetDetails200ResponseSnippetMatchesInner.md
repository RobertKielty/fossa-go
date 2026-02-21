# GetSnippetDetails200ResponseSnippetMatchesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The file path where the match was detected | 
**MatchPercentage** | **float32** | Match percentage for this specific match | 
**RejectionDetails** | Pointer to [**GetSnippets200ResponseResultsInnerRejectionDetails**](GetSnippets200ResponseResultsInnerRejectionDetails.md) |  | [optional] 

## Methods

### NewGetSnippetDetails200ResponseSnippetMatchesInner

`func NewGetSnippetDetails200ResponseSnippetMatchesInner(path string, matchPercentage float32, ) *GetSnippetDetails200ResponseSnippetMatchesInner`

NewGetSnippetDetails200ResponseSnippetMatchesInner instantiates a new GetSnippetDetails200ResponseSnippetMatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetDetails200ResponseSnippetMatchesInnerWithDefaults

`func NewGetSnippetDetails200ResponseSnippetMatchesInnerWithDefaults() *GetSnippetDetails200ResponseSnippetMatchesInner`

NewGetSnippetDetails200ResponseSnippetMatchesInnerWithDefaults instantiates a new GetSnippetDetails200ResponseSnippetMatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) SetPath(v string)`

SetPath sets Path field to given value.


### GetMatchPercentage

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetMatchPercentage() float32`

GetMatchPercentage returns the MatchPercentage field if non-nil, zero value otherwise.

### GetMatchPercentageOk

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetMatchPercentageOk() (*float32, bool)`

GetMatchPercentageOk returns a tuple with the MatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPercentage

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) SetMatchPercentage(v float32)`

SetMatchPercentage sets MatchPercentage field to given value.


### GetRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetRejectionDetails() GetSnippets200ResponseResultsInnerRejectionDetails`

GetRejectionDetails returns the RejectionDetails field if non-nil, zero value otherwise.

### GetRejectionDetailsOk

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) GetRejectionDetailsOk() (*GetSnippets200ResponseResultsInnerRejectionDetails, bool)`

GetRejectionDetailsOk returns a tuple with the RejectionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) SetRejectionDetails(v GetSnippets200ResponseResultsInnerRejectionDetails)`

SetRejectionDetails sets RejectionDetails field to given value.

### HasRejectionDetails

`func (o *GetSnippetDetails200ResponseSnippetMatchesInner) HasRejectionDetails() bool`

HasRejectionDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


