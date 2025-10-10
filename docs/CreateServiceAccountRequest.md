# CreateServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | The username for the service account | 
**Email** | Pointer to **string** | Optional email address for the service account | [optional] 
**FullName** | Pointer to **string** | Optional full name/description for the service account | [optional] 
**OrgRoleId** | Pointer to **int32** | ID of the organization role to assign to the service account. At least one of orgRoleId or team must be provided. | [optional] 
**Team** | Pointer to [**CreateServiceAccountRequestTeam**](CreateServiceAccountRequestTeam.md) |  | [optional] 
**HasPushOnlyApiToken** | Pointer to **bool** | Whether to create a push-only API token for the service account | [optional] [default to false]
**HasFullApiToken** | Pointer to **bool** | Whether to create a full access API token for the service account | [optional] [default to false]

## Methods

### NewCreateServiceAccountRequest

`func NewCreateServiceAccountRequest(username string, ) *CreateServiceAccountRequest`

NewCreateServiceAccountRequest instantiates a new CreateServiceAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountRequestWithDefaults

`func NewCreateServiceAccountRequestWithDefaults() *CreateServiceAccountRequest`

NewCreateServiceAccountRequestWithDefaults instantiates a new CreateServiceAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateServiceAccountRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateServiceAccountRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateServiceAccountRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *CreateServiceAccountRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateServiceAccountRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateServiceAccountRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateServiceAccountRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *CreateServiceAccountRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateServiceAccountRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateServiceAccountRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CreateServiceAccountRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetOrgRoleId

`func (o *CreateServiceAccountRequest) GetOrgRoleId() int32`

GetOrgRoleId returns the OrgRoleId field if non-nil, zero value otherwise.

### GetOrgRoleIdOk

`func (o *CreateServiceAccountRequest) GetOrgRoleIdOk() (*int32, bool)`

GetOrgRoleIdOk returns a tuple with the OrgRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRoleId

`func (o *CreateServiceAccountRequest) SetOrgRoleId(v int32)`

SetOrgRoleId sets OrgRoleId field to given value.

### HasOrgRoleId

`func (o *CreateServiceAccountRequest) HasOrgRoleId() bool`

HasOrgRoleId returns a boolean if a field has been set.

### GetTeam

`func (o *CreateServiceAccountRequest) GetTeam() CreateServiceAccountRequestTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *CreateServiceAccountRequest) GetTeamOk() (*CreateServiceAccountRequestTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *CreateServiceAccountRequest) SetTeam(v CreateServiceAccountRequestTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *CreateServiceAccountRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetHasPushOnlyApiToken

`func (o *CreateServiceAccountRequest) GetHasPushOnlyApiToken() bool`

GetHasPushOnlyApiToken returns the HasPushOnlyApiToken field if non-nil, zero value otherwise.

### GetHasPushOnlyApiTokenOk

`func (o *CreateServiceAccountRequest) GetHasPushOnlyApiTokenOk() (*bool, bool)`

GetHasPushOnlyApiTokenOk returns a tuple with the HasPushOnlyApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPushOnlyApiToken

`func (o *CreateServiceAccountRequest) SetHasPushOnlyApiToken(v bool)`

SetHasPushOnlyApiToken sets HasPushOnlyApiToken field to given value.

### HasHasPushOnlyApiToken

`func (o *CreateServiceAccountRequest) HasHasPushOnlyApiToken() bool`

HasHasPushOnlyApiToken returns a boolean if a field has been set.

### GetHasFullApiToken

`func (o *CreateServiceAccountRequest) GetHasFullApiToken() bool`

GetHasFullApiToken returns the HasFullApiToken field if non-nil, zero value otherwise.

### GetHasFullApiTokenOk

`func (o *CreateServiceAccountRequest) GetHasFullApiTokenOk() (*bool, bool)`

GetHasFullApiTokenOk returns a tuple with the HasFullApiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFullApiToken

`func (o *CreateServiceAccountRequest) SetHasFullApiToken(v bool)`

SetHasFullApiToken sets HasFullApiToken field to given value.

### HasHasFullApiToken

`func (o *CreateServiceAccountRequest) HasHasFullApiToken() bool`

HasHasFullApiToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


