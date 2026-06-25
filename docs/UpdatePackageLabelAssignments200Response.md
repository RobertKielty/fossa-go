# UpdatePackageLabelAssignments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLabelAssignments** | [**[]UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner**](UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner.md) | The package label assignments that were written. Newly created assignments only -- existing (duplicate) assignments are not returned. | 

## Methods

### NewUpdatePackageLabelAssignments200Response

`func NewUpdatePackageLabelAssignments200Response(packageLabelAssignments []UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner, ) *UpdatePackageLabelAssignments200Response`

NewUpdatePackageLabelAssignments200Response instantiates a new UpdatePackageLabelAssignments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePackageLabelAssignments200ResponseWithDefaults

`func NewUpdatePackageLabelAssignments200ResponseWithDefaults() *UpdatePackageLabelAssignments200Response`

NewUpdatePackageLabelAssignments200ResponseWithDefaults instantiates a new UpdatePackageLabelAssignments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLabelAssignments

`func (o *UpdatePackageLabelAssignments200Response) GetPackageLabelAssignments() []UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner`

GetPackageLabelAssignments returns the PackageLabelAssignments field if non-nil, zero value otherwise.

### GetPackageLabelAssignmentsOk

`func (o *UpdatePackageLabelAssignments200Response) GetPackageLabelAssignmentsOk() (*[]UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner, bool)`

GetPackageLabelAssignmentsOk returns a tuple with the PackageLabelAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLabelAssignments

`func (o *UpdatePackageLabelAssignments200Response) SetPackageLabelAssignments(v []UpdatePackageLabelAssignments200ResponsePackageLabelAssignmentsInner)`

SetPackageLabelAssignments sets PackageLabelAssignments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


