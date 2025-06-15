# FossaApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier associated with the error | [optional] 
**Code** | Pointer to **int32** | fossa specific error code | [optional] 
**Message** | Pointer to **string** | message associated with this error | [optional] 
**Name** | Pointer to **string** | name of the error | [optional] 
**HttpStatusCode** | Pointer to **int32** | http status code number | [optional] 

## Methods

### NewFossaApiError

`func NewFossaApiError() *FossaApiError`

NewFossaApiError instantiates a new FossaApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFossaApiErrorWithDefaults

`func NewFossaApiErrorWithDefaults() *FossaApiError`

NewFossaApiErrorWithDefaults instantiates a new FossaApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *FossaApiError) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *FossaApiError) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *FossaApiError) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *FossaApiError) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCode

`func (o *FossaApiError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FossaApiError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FossaApiError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *FossaApiError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *FossaApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FossaApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FossaApiError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FossaApiError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetName

`func (o *FossaApiError) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FossaApiError) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FossaApiError) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FossaApiError) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *FossaApiError) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *FossaApiError) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *FossaApiError) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *FossaApiError) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


