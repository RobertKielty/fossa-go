# SlackWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **string** | The Slack channel to send updates to | [optional] 
**WebhookUrl** | Pointer to **string** | The configured webhook url in your Slack Application | [optional] 

## Methods

### NewSlackWebhook

`func NewSlackWebhook() *SlackWebhook`

NewSlackWebhook instantiates a new SlackWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackWebhookWithDefaults

`func NewSlackWebhookWithDefaults() *SlackWebhook`

NewSlackWebhookWithDefaults instantiates a new SlackWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *SlackWebhook) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *SlackWebhook) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *SlackWebhook) SetChannel(v string)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *SlackWebhook) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *SlackWebhook) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *SlackWebhook) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *SlackWebhook) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *SlackWebhook) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


