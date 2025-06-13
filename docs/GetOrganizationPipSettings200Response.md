# GetOrganizationPipSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]GetOrganizationPipSettings200ResponseRepositoriesInner**](GetOrganizationPipSettings200ResponseRepositoriesInner.md) | List of configured Pip Repositories | [optional] 

## Methods

### NewGetOrganizationPipSettings200Response

`func NewGetOrganizationPipSettings200Response() *GetOrganizationPipSettings200Response`

NewGetOrganizationPipSettings200Response instantiates a new GetOrganizationPipSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationPipSettings200ResponseWithDefaults

`func NewGetOrganizationPipSettings200ResponseWithDefaults() *GetOrganizationPipSettings200Response`

NewGetOrganizationPipSettings200ResponseWithDefaults instantiates a new GetOrganizationPipSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *GetOrganizationPipSettings200Response) GetRepositories() []GetOrganizationPipSettings200ResponseRepositoriesInner`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *GetOrganizationPipSettings200Response) GetRepositoriesOk() (*[]GetOrganizationPipSettings200ResponseRepositoriesInner, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *GetOrganizationPipSettings200Response) SetRepositories(v []GetOrganizationPipSettings200ResponseRepositoriesInner)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *GetOrganizationPipSettings200Response) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


