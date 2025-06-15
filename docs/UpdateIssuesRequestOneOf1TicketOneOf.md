# UpdateIssuesRequestOneOf1TicketOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**UniqueTickets** | Pointer to [**[]UpdateIssuesRequestOneOf1TicketOneOfUniqueTicketsInner**](UpdateIssuesRequestOneOf1TicketOneOfUniqueTicketsInner.md) |  | [optional] 
**TrackerId** | Pointer to **string** |  | [optional] 
**JiraId** | Pointer to **string** |  | [optional] 
**IssueType** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**CustomFields** | Pointer to [**[]UpdateIssuesRequestOneOf1TicketOneOfCustomFieldsInner**](UpdateIssuesRequestOneOf1TicketOneOfCustomFieldsInner.md) |  | [optional] 
**Components** | Pointer to [**[]UpdateIssuesRequestOneOf1TicketOneOfComponentsInner**](UpdateIssuesRequestOneOf1TicketOneOfComponentsInner.md) |  | [optional] 

## Methods

### NewUpdateIssuesRequestOneOf1TicketOneOf

`func NewUpdateIssuesRequestOneOf1TicketOneOf() *UpdateIssuesRequestOneOf1TicketOneOf`

NewUpdateIssuesRequestOneOf1TicketOneOf instantiates a new UpdateIssuesRequestOneOf1TicketOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIssuesRequestOneOf1TicketOneOfWithDefaults

`func NewUpdateIssuesRequestOneOf1TicketOneOfWithDefaults() *UpdateIssuesRequestOneOf1TicketOneOf`

NewUpdateIssuesRequestOneOf1TicketOneOfWithDefaults instantiates a new UpdateIssuesRequestOneOf1TicketOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUniqueTickets

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetUniqueTickets() []UpdateIssuesRequestOneOf1TicketOneOfUniqueTicketsInner`

GetUniqueTickets returns the UniqueTickets field if non-nil, zero value otherwise.

### GetUniqueTicketsOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetUniqueTicketsOk() (*[]UpdateIssuesRequestOneOf1TicketOneOfUniqueTicketsInner, bool)`

GetUniqueTicketsOk returns a tuple with the UniqueTickets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueTickets

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetUniqueTickets(v []UpdateIssuesRequestOneOf1TicketOneOfUniqueTicketsInner)`

SetUniqueTickets sets UniqueTickets field to given value.

### HasUniqueTickets

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasUniqueTickets() bool`

HasUniqueTickets returns a boolean if a field has been set.

### GetTrackerId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetTrackerId() string`

GetTrackerId returns the TrackerId field if non-nil, zero value otherwise.

### GetTrackerIdOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetTrackerIdOk() (*string, bool)`

GetTrackerIdOk returns a tuple with the TrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetTrackerId(v string)`

SetTrackerId sets TrackerId field to given value.

### HasTrackerId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasTrackerId() bool`

HasTrackerId returns a boolean if a field has been set.

### GetJiraId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetJiraId() string`

GetJiraId returns the JiraId field if non-nil, zero value otherwise.

### GetJiraIdOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetJiraIdOk() (*string, bool)`

GetJiraIdOk returns a tuple with the JiraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetJiraId(v string)`

SetJiraId sets JiraId field to given value.

### HasJiraId

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasJiraId() bool`

HasJiraId returns a boolean if a field has been set.

### GetIssueType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCustomFields

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetCustomFields() []UpdateIssuesRequestOneOf1TicketOneOfCustomFieldsInner`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetCustomFieldsOk() (*[]UpdateIssuesRequestOneOf1TicketOneOfCustomFieldsInner, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetCustomFields(v []UpdateIssuesRequestOneOf1TicketOneOfCustomFieldsInner)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetComponents

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetComponents() []UpdateIssuesRequestOneOf1TicketOneOfComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) GetComponentsOk() (*[]UpdateIssuesRequestOneOf1TicketOneOfComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) SetComponents(v []UpdateIssuesRequestOneOf1TicketOneOfComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *UpdateIssuesRequestOneOf1TicketOneOf) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


