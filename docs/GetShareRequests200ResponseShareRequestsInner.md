# GetShareRequests200ResponseShareRequestsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginRevisionId** | Pointer to **string** | The ID of the revision that was shared | [optional] 
**OriginProjectLocator** | Pointer to **string** | The locator of the project that was shared | [optional] 
**ProjectTitle** | Pointer to **string** | The title of the project that was shared | [optional] 
**SharedOrganizationId** | Pointer to **int64** | The ID of the shared organization record | [optional] 
**SharedAt** | Pointer to **time.Time** | The timestamp when the share was created | [optional] 
**ReceivingOrganizationName** | Pointer to **string** | The name of the organization that received the share | [optional] 

## Methods

### NewGetShareRequests200ResponseShareRequestsInner

`func NewGetShareRequests200ResponseShareRequestsInner() *GetShareRequests200ResponseShareRequestsInner`

NewGetShareRequests200ResponseShareRequestsInner instantiates a new GetShareRequests200ResponseShareRequestsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetShareRequests200ResponseShareRequestsInnerWithDefaults

`func NewGetShareRequests200ResponseShareRequestsInnerWithDefaults() *GetShareRequests200ResponseShareRequestsInner`

NewGetShareRequests200ResponseShareRequestsInnerWithDefaults instantiates a new GetShareRequests200ResponseShareRequestsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginRevisionId

`func (o *GetShareRequests200ResponseShareRequestsInner) GetOriginRevisionId() string`

GetOriginRevisionId returns the OriginRevisionId field if non-nil, zero value otherwise.

### GetOriginRevisionIdOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetOriginRevisionIdOk() (*string, bool)`

GetOriginRevisionIdOk returns a tuple with the OriginRevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginRevisionId

`func (o *GetShareRequests200ResponseShareRequestsInner) SetOriginRevisionId(v string)`

SetOriginRevisionId sets OriginRevisionId field to given value.

### HasOriginRevisionId

`func (o *GetShareRequests200ResponseShareRequestsInner) HasOriginRevisionId() bool`

HasOriginRevisionId returns a boolean if a field has been set.

### GetOriginProjectLocator

`func (o *GetShareRequests200ResponseShareRequestsInner) GetOriginProjectLocator() string`

GetOriginProjectLocator returns the OriginProjectLocator field if non-nil, zero value otherwise.

### GetOriginProjectLocatorOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetOriginProjectLocatorOk() (*string, bool)`

GetOriginProjectLocatorOk returns a tuple with the OriginProjectLocator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginProjectLocator

`func (o *GetShareRequests200ResponseShareRequestsInner) SetOriginProjectLocator(v string)`

SetOriginProjectLocator sets OriginProjectLocator field to given value.

### HasOriginProjectLocator

`func (o *GetShareRequests200ResponseShareRequestsInner) HasOriginProjectLocator() bool`

HasOriginProjectLocator returns a boolean if a field has been set.

### GetProjectTitle

`func (o *GetShareRequests200ResponseShareRequestsInner) GetProjectTitle() string`

GetProjectTitle returns the ProjectTitle field if non-nil, zero value otherwise.

### GetProjectTitleOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetProjectTitleOk() (*string, bool)`

GetProjectTitleOk returns a tuple with the ProjectTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTitle

`func (o *GetShareRequests200ResponseShareRequestsInner) SetProjectTitle(v string)`

SetProjectTitle sets ProjectTitle field to given value.

### HasProjectTitle

`func (o *GetShareRequests200ResponseShareRequestsInner) HasProjectTitle() bool`

HasProjectTitle returns a boolean if a field has been set.

### GetSharedOrganizationId

`func (o *GetShareRequests200ResponseShareRequestsInner) GetSharedOrganizationId() int64`

GetSharedOrganizationId returns the SharedOrganizationId field if non-nil, zero value otherwise.

### GetSharedOrganizationIdOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetSharedOrganizationIdOk() (*int64, bool)`

GetSharedOrganizationIdOk returns a tuple with the SharedOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedOrganizationId

`func (o *GetShareRequests200ResponseShareRequestsInner) SetSharedOrganizationId(v int64)`

SetSharedOrganizationId sets SharedOrganizationId field to given value.

### HasSharedOrganizationId

`func (o *GetShareRequests200ResponseShareRequestsInner) HasSharedOrganizationId() bool`

HasSharedOrganizationId returns a boolean if a field has been set.

### GetSharedAt

`func (o *GetShareRequests200ResponseShareRequestsInner) GetSharedAt() time.Time`

GetSharedAt returns the SharedAt field if non-nil, zero value otherwise.

### GetSharedAtOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetSharedAtOk() (*time.Time, bool)`

GetSharedAtOk returns a tuple with the SharedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAt

`func (o *GetShareRequests200ResponseShareRequestsInner) SetSharedAt(v time.Time)`

SetSharedAt sets SharedAt field to given value.

### HasSharedAt

`func (o *GetShareRequests200ResponseShareRequestsInner) HasSharedAt() bool`

HasSharedAt returns a boolean if a field has been set.

### GetReceivingOrganizationName

`func (o *GetShareRequests200ResponseShareRequestsInner) GetReceivingOrganizationName() string`

GetReceivingOrganizationName returns the ReceivingOrganizationName field if non-nil, zero value otherwise.

### GetReceivingOrganizationNameOk

`func (o *GetShareRequests200ResponseShareRequestsInner) GetReceivingOrganizationNameOk() (*string, bool)`

GetReceivingOrganizationNameOk returns a tuple with the ReceivingOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivingOrganizationName

`func (o *GetShareRequests200ResponseShareRequestsInner) SetReceivingOrganizationName(v string)`

SetReceivingOrganizationName sets ReceivingOrganizationName field to given value.

### HasReceivingOrganizationName

`func (o *GetShareRequests200ResponseShareRequestsInner) HasReceivingOrganizationName() bool`

HasReceivingOrganizationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


