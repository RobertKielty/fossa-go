# GetOrganizationNugetSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sources** | Pointer to [**[]GetOrganizationNugetSettings200ResponseSourcesInner**](GetOrganizationNugetSettings200ResponseSourcesInner.md) | List of configured Nuget Feeds | [optional] 

## Methods

### NewGetOrganizationNugetSettings200Response

`func NewGetOrganizationNugetSettings200Response() *GetOrganizationNugetSettings200Response`

NewGetOrganizationNugetSettings200Response instantiates a new GetOrganizationNugetSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationNugetSettings200ResponseWithDefaults

`func NewGetOrganizationNugetSettings200ResponseWithDefaults() *GetOrganizationNugetSettings200Response`

NewGetOrganizationNugetSettings200ResponseWithDefaults instantiates a new GetOrganizationNugetSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSources

`func (o *GetOrganizationNugetSettings200Response) GetSources() []GetOrganizationNugetSettings200ResponseSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetOrganizationNugetSettings200Response) GetSourcesOk() (*[]GetOrganizationNugetSettings200ResponseSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetOrganizationNugetSettings200Response) SetSources(v []GetOrganizationNugetSettings200ResponseSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetOrganizationNugetSettings200Response) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


