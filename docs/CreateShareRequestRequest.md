# CreateShareRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionId** | **string** | The locator for the revision to share | 
**SharedOrganizationId** | **int64** | The ID of the shared organization record | 

## Methods

### NewCreateShareRequestRequest

`func NewCreateShareRequestRequest(revisionId string, sharedOrganizationId int64, ) *CreateShareRequestRequest`

NewCreateShareRequestRequest instantiates a new CreateShareRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateShareRequestRequestWithDefaults

`func NewCreateShareRequestRequestWithDefaults() *CreateShareRequestRequest`

NewCreateShareRequestRequestWithDefaults instantiates a new CreateShareRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionId

`func (o *CreateShareRequestRequest) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *CreateShareRequestRequest) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *CreateShareRequestRequest) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetSharedOrganizationId

`func (o *CreateShareRequestRequest) GetSharedOrganizationId() int64`

GetSharedOrganizationId returns the SharedOrganizationId field if non-nil, zero value otherwise.

### GetSharedOrganizationIdOk

`func (o *CreateShareRequestRequest) GetSharedOrganizationIdOk() (*int64, bool)`

GetSharedOrganizationIdOk returns a tuple with the SharedOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOrganizationId

`func (o *CreateShareRequestRequest) SetSharedOrganizationId(v int64)`

SetSharedOrganizationId sets SharedOrganizationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


