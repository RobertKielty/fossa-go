# GetOrganizationSlackSettings200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The Slack channel to send updates to | [optional] 
**WebhookUrl** | Pointer to **string** | The configured webhook url in your Slack Application | [optional] 

## Methods

### NewGetOrganizationSlackSettings200ResponseInner

`func NewGetOrganizationSlackSettings200ResponseInner() *GetOrganizationSlackSettings200ResponseInner`

NewGetOrganizationSlackSettings200ResponseInner instantiates a new GetOrganizationSlackSettings200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSlackSettings200ResponseInnerWithDefaults

`func NewGetOrganizationSlackSettings200ResponseInnerWithDefaults() *GetOrganizationSlackSettings200ResponseInner`

NewGetOrganizationSlackSettings200ResponseInnerWithDefaults instantiates a new GetOrganizationSlackSettings200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *GetOrganizationSlackSettings200ResponseInner) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GetOrganizationSlackSettings200ResponseInner) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GetOrganizationSlackSettings200ResponseInner) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GetOrganizationSlackSettings200ResponseInner) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *GetOrganizationSlackSettings200ResponseInner) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GetOrganizationSlackSettings200ResponseInner) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GetOrganizationSlackSettings200ResponseInner) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *GetOrganizationSlackSettings200ResponseInner) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


