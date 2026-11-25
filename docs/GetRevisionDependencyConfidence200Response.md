# GetRevisionDependencyConfidence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confidences** | **map[string]string** | A map keyed by dependency locator. Each value is the confidence string for that Binary Component dependency. | 

## Methods

### NewGetRevisionDependencyConfidence200Response

`func NewGetRevisionDependencyConfidence200Response(confidences map[string]string, ) *GetRevisionDependencyConfidence200Response`

NewGetRevisionDependencyConfidence200Response instantiates a new GetRevisionDependencyConfidence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionDependencyConfidence200ResponseWithDefaults

`func NewGetRevisionDependencyConfidence200ResponseWithDefaults() *GetRevisionDependencyConfidence200Response`

NewGetRevisionDependencyConfidence200ResponseWithDefaults instantiates a new GetRevisionDependencyConfidence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidences

`func (o *GetRevisionDependencyConfidence200Response) GetConfidences() map[string]string`

GetConfidences returns the Confidences field if non-nil, zero value otherwise.

### GetConfidencesOk

`func (o *GetRevisionDependencyConfidence200Response) GetConfidencesOk() (*map[string]string, bool)`

GetConfidencesOk returns a tuple with the Confidences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidences

`func (o *GetRevisionDependencyConfidence200Response) SetConfidences(v map[string]string)`

SetConfidences sets Confidences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


