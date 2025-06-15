# DeletedJira

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | ID of the Jira site that was requested to be deleted | [optional] 
**Deleted** | Pointer to **bool** | whether the Jira sire was successfully deleted or not | [optional] 

## Methods

### NewDeletedJira

`func NewDeletedJira() *DeletedJira`

NewDeletedJira instantiates a new DeletedJira object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedJiraWithDefaults

`func NewDeletedJiraWithDefaults() *DeletedJira`

NewDeletedJiraWithDefaults instantiates a new DeletedJira object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeletedJira) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedJira) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedJira) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedJira) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeleted

`func (o *DeletedJira) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeletedJira) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeletedJira) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeletedJira) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


