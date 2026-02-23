# UserInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int32** |  | 
**Email** | **string** |  | 
**TokenId** | **int32** |  | 
**CreatedByUserId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Token** | [**UserInvitationToken**](UserInvitationToken.md) |  | 

## Methods

### NewUserInvitation

`func NewUserInvitation(organizationId int32, email string, tokenId int32, createdByUserId int32, createdAt time.Time, updatedAt time.Time, token UserInvitationToken, ) *UserInvitation`

NewUserInvitation instantiates a new UserInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationWithDefaults

`func NewUserInvitationWithDefaults() *UserInvitation`

NewUserInvitationWithDefaults instantiates a new UserInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *UserInvitation) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserInvitation) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserInvitation) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetEmail

`func (o *UserInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTokenId

`func (o *UserInvitation) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UserInvitation) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UserInvitation) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetCreatedByUserId

`func (o *UserInvitation) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *UserInvitation) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *UserInvitation) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetCreatedAt

`func (o *UserInvitation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInvitation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInvitation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserInvitation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserInvitation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserInvitation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetToken

`func (o *UserInvitation) GetToken() UserInvitationToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserInvitation) GetTokenOk() (*UserInvitationToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserInvitation) SetToken(v UserInvitationToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


