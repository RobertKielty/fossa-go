# UpdateReportOptionByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the report option (1-80 characters). Must be unique within the organization. | [optional] 
**Options** | Pointer to [**UpdateReportOptionByIdRequestOptions**](UpdateReportOptionByIdRequestOptions.md) |  | [optional] 

## Methods

### NewUpdateReportOptionByIdRequest

`func NewUpdateReportOptionByIdRequest() *UpdateReportOptionByIdRequest`

NewUpdateReportOptionByIdRequest instantiates a new UpdateReportOptionByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReportOptionByIdRequestWithDefaults

`func NewUpdateReportOptionByIdRequestWithDefaults() *UpdateReportOptionByIdRequest`

NewUpdateReportOptionByIdRequestWithDefaults instantiates a new UpdateReportOptionByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateReportOptionByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReportOptionByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReportOptionByIdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateReportOptionByIdRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateReportOptionByIdRequest) GetOptions() UpdateReportOptionByIdRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateReportOptionByIdRequest) GetOptionsOk() (*UpdateReportOptionByIdRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateReportOptionByIdRequest) SetOptions(v UpdateReportOptionByIdRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateReportOptionByIdRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


