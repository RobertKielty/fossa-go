# UserInvitationCreateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **int32** |  | 
**Email** | **string** |  | 
**TokenId** | **int32** |  | 
**CreatedByUserId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUserInvitationCreateResult

`func NewUserInvitationCreateResult(organizationId int32, email string, tokenId int32, createdByUserId int32, createdAt time.Time, updatedAt time.Time, ) *UserInvitationCreateResult`

NewUserInvitationCreateResult instantiates a new UserInvitationCreateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInvitationCreateResultWithDefaults

`func NewUserInvitationCreateResultWithDefaults() *UserInvitationCreateResult`

NewUserInvitationCreateResultWithDefaults instantiates a new UserInvitationCreateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *UserInvitationCreateResult) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserInvitationCreateResult) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserInvitationCreateResult) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetEmail

`func (o *UserInvitationCreateResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInvitationCreateResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInvitationCreateResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTokenId

`func (o *UserInvitationCreateResult) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UserInvitationCreateResult) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UserInvitationCreateResult) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetCreatedByUserId

`func (o *UserInvitationCreateResult) GetCreatedByUserId() int32`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *UserInvitationCreateResult) GetCreatedByUserIdOk() (*int32, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *UserInvitationCreateResult) SetCreatedByUserId(v int32)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### GetCreatedAt

`func (o *UserInvitationCreateResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInvitationCreateResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInvitationCreateResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserInvitationCreateResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserInvitationCreateResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserInvitationCreateResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


