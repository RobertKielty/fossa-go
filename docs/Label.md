# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of this label | [optional] 
**OrganizationId** | Pointer to **int32** | ID of the organization that this Label is associated with | [optional] 
**Label** | Pointer to **string** | Text that this Label represents | [optional] 
**Projects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewLabel

`func NewLabel() *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Label) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Label) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Label) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Label) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *Label) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Label) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Label) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *Label) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetLabel

`func (o *Label) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Label) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Label) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Label) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProjects

`func (o *Label) GetProjects() []string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *Label) GetProjectsOk() (*[]string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *Label) SetProjects(v []string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *Label) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


