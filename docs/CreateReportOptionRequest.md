# CreateReportOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the report option (1-80 characters). Must be unique within the organization. | 
**Options** | [**CreateReportOptionRequestOptions**](CreateReportOptionRequestOptions.md) |  | 

## Methods

### NewCreateReportOptionRequest

`func NewCreateReportOptionRequest(name string, options CreateReportOptionRequestOptions, ) *CreateReportOptionRequest`

NewCreateReportOptionRequest instantiates a new CreateReportOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReportOptionRequestWithDefaults

`func NewCreateReportOptionRequestWithDefaults() *CreateReportOptionRequest`

NewCreateReportOptionRequestWithDefaults instantiates a new CreateReportOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateReportOptionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateReportOptionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateReportOptionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOptions

`func (o *CreateReportOptionRequest) GetOptions() CreateReportOptionRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateReportOptionRequest) GetOptionsOk() (*CreateReportOptionRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateReportOptionRequest) SetOptions(v CreateReportOptionRequestOptions)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


