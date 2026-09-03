# GetOrganizationScimStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Whether SCIM provisioning is currently active for the organization. True when the organization has an enabled SCIM token; revoking the token turns this off. | 
**IsSsoConfigured** | **bool** | Whether the organization has an SSO method (SAML, LDAP, or Google) configured. SCIM provisions SSO-only users, so a SCIM token cannot be generated until this is true. | 

## Methods

### NewGetOrganizationScimStatus200Response

`func NewGetOrganizationScimStatus200Response(isEnabled bool, isSsoConfigured bool, ) *GetOrganizationScimStatus200Response`

NewGetOrganizationScimStatus200Response instantiates a new GetOrganizationScimStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationScimStatus200ResponseWithDefaults

`func NewGetOrganizationScimStatus200ResponseWithDefaults() *GetOrganizationScimStatus200Response`

NewGetOrganizationScimStatus200ResponseWithDefaults instantiates a new GetOrganizationScimStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *GetOrganizationScimStatus200Response) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *GetOrganizationScimStatus200Response) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *GetOrganizationScimStatus200Response) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetIsSsoConfigured

`func (o *GetOrganizationScimStatus200Response) GetIsSsoConfigured() bool`

GetIsSsoConfigured returns the IsSsoConfigured field if non-nil, zero value otherwise.

### GetIsSsoConfiguredOk

`func (o *GetOrganizationScimStatus200Response) GetIsSsoConfiguredOk() (*bool, bool)`

GetIsSsoConfiguredOk returns a tuple with the IsSsoConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoConfigured

`func (o *GetOrganizationScimStatus200Response) SetIsSsoConfigured(v bool)`

SetIsSsoConfigured sets IsSsoConfigured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


