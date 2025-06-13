# MavenRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** | UUID of the Maven repo (For FOSSA internal usage) | [optional] 
**Id** | Pointer to **string** | User configured ID (corresponds with a Maven Server if credentials are necessary) | [optional] 
**Url** | Pointer to **string** | URL of the Maven Repository | [optional] 

## Methods

### NewMavenRepository

`func NewMavenRepository() *MavenRepository`

NewMavenRepository instantiates a new MavenRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenRepositoryWithDefaults

`func NewMavenRepositoryWithDefaults() *MavenRepository`

NewMavenRepositoryWithDefaults instantiates a new MavenRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MavenRepository) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MavenRepository) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MavenRepository) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *MavenRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetId

`func (o *MavenRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MavenRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MavenRepository) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MavenRepository) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *MavenRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MavenRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MavenRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MavenRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


