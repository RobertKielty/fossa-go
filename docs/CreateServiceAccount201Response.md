# CreateServiceAccount201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The user id of the service account | 
**Username** | **string** | The username of the service account | 
**Email** | Pointer to **NullableString** | The email address of the service account (if provided) | [optional] 
**FullName** | Pointer to **NullableString** | The full name/description of the service account (if provided) | [optional] 
**FullApiToken** | Pointer to **string** | The full access API token (only present if hasFullApiToken was true) | [optional] 
**PushOnlyApiToken** | Pointer to **string** | The push-only API token (only present if hasPushOnlyApiToken was true) | [optional] 

## Methods

### NewCreateServiceAccount201Response

`func NewCreateServiceAccount201Response(id int32, username string, ) *CreateServiceAccount201Response`

NewCreateServiceAccount201Response instantiates a new CreateServiceAccount201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccount201ResponseWithDefaults

`func NewCreateServiceAccount201ResponseWithDefaults() *CreateServiceAccount201Response`

NewCreateServiceAccount201ResponseWithDefaults instantiates a new CreateServiceAccount201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateServiceAccount201Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateServiceAccount201Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateServiceAccount201Response) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *CreateServiceAccount201Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateServiceAccount201Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateServiceAccount201Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *CreateServiceAccount201Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateServiceAccount201Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateServiceAccount201Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateServiceAccount201Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CreateServiceAccount201Response) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CreateServiceAccount201Response) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *CreateServiceAccount201Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateServiceAccount201Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateServiceAccount201Response) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CreateServiceAccount201Response) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### SetFullNameNil

`func (o *CreateServiceAccount201Response) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *CreateServiceAccount201Response) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetFullApiToken

`func (o *CreateServiceAccount201Response) GetFullApiToken() string`

GetFullApiToken returns the FullApiToken field if non-nil, zero value otherwise.

### GetFullApiTokenOk

`func (o *CreateServiceAccount201Response) GetFullApiTokenOk() (*string, bool)`

GetFullApiTokenOk returns a tuple with the FullApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullApiToken

`func (o *CreateServiceAccount201Response) SetFullApiToken(v string)`

SetFullApiToken sets FullApiToken field to given value.

### HasFullApiToken

`func (o *CreateServiceAccount201Response) HasFullApiToken() bool`

HasFullApiToken returns a boolean if a field has been set.

### GetPushOnlyApiToken

`func (o *CreateServiceAccount201Response) GetPushOnlyApiToken() string`

GetPushOnlyApiToken returns the PushOnlyApiToken field if non-nil, zero value otherwise.

### GetPushOnlyApiTokenOk

`func (o *CreateServiceAccount201Response) GetPushOnlyApiTokenOk() (*string, bool)`

GetPushOnlyApiTokenOk returns a tuple with the PushOnlyApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushOnlyApiToken

`func (o *CreateServiceAccount201Response) SetPushOnlyApiToken(v string)`

SetPushOnlyApiToken sets PushOnlyApiToken field to given value.

### HasPushOnlyApiToken

`func (o *CreateServiceAccount201Response) HasPushOnlyApiToken() bool`

HasPushOnlyApiToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


