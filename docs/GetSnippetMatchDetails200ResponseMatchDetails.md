# GetSnippetMatchDetails200ResponseMatchDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The file path where the match was detected | 
**MatchPercentage** | **float32** | Percentage match confidence for this specific path | 
**ReferenceCode** | [**[]GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner**](GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner.md) | Code lines from the reference (third-party) source | 
**DetectedCode** | [**[]GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner**](GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner.md) | Code lines from the detected source in the user&#39;s project | 

## Methods

### NewGetSnippetMatchDetails200ResponseMatchDetails

`func NewGetSnippetMatchDetails200ResponseMatchDetails(path string, matchPercentage float32, referenceCode []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner, detectedCode []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner, ) *GetSnippetMatchDetails200ResponseMatchDetails`

NewGetSnippetMatchDetails200ResponseMatchDetails instantiates a new GetSnippetMatchDetails200ResponseMatchDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetMatchDetails200ResponseMatchDetailsWithDefaults

`func NewGetSnippetMatchDetails200ResponseMatchDetailsWithDefaults() *GetSnippetMatchDetails200ResponseMatchDetails`

NewGetSnippetMatchDetails200ResponseMatchDetailsWithDefaults instantiates a new GetSnippetMatchDetails200ResponseMatchDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) SetPath(v string)`

SetPath sets Path field to given value.


### GetMatchPercentage

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetMatchPercentage() float32`

GetMatchPercentage returns the MatchPercentage field if non-nil, zero value otherwise.

### GetMatchPercentageOk

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetMatchPercentageOk() (*float32, bool)`

GetMatchPercentageOk returns a tuple with the MatchPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPercentage

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) SetMatchPercentage(v float32)`

SetMatchPercentage sets MatchPercentage field to given value.


### GetReferenceCode

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetReferenceCode() []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner`

GetReferenceCode returns the ReferenceCode field if non-nil, zero value otherwise.

### GetReferenceCodeOk

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetReferenceCodeOk() (*[]GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner, bool)`

GetReferenceCodeOk returns a tuple with the ReferenceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceCode

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) SetReferenceCode(v []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner)`

SetReferenceCode sets ReferenceCode field to given value.


### GetDetectedCode

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetDetectedCode() []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner`

GetDetectedCode returns the DetectedCode field if non-nil, zero value otherwise.

### GetDetectedCodeOk

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) GetDetectedCodeOk() (*[]GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner, bool)`

GetDetectedCodeOk returns a tuple with the DetectedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedCode

`func (o *GetSnippetMatchDetails200ResponseMatchDetails) SetDetectedCode(v []GetSnippetMatchDetails200ResponseMatchDetailsAllOfReferenceCodeInner)`

SetDetectedCode sets DetectedCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


