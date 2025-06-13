# CreateReleaseGroup200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **int32** |  | [optional] 
**PolicyId** | Pointer to **NullableInt32** |  | [optional] 
**SecurityPolicyId** | Pointer to **NullableInt32** |  | [optional] 
**QualityPolicyId** | Pointer to **NullableInt32** |  | [optional] 
**PublicOnPortal** | Pointer to **bool** |  | [optional] 
**ReportCustomText** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateReleaseGroup200Response

`func NewCreateReleaseGroup200Response() *CreateReleaseGroup200Response`

NewCreateReleaseGroup200Response instantiates a new CreateReleaseGroup200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseGroup200ResponseWithDefaults

`func NewCreateReleaseGroup200ResponseWithDefaults() *CreateReleaseGroup200Response`

NewCreateReleaseGroup200ResponseWithDefaults instantiates a new CreateReleaseGroup200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReleaseGroup200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReleaseGroup200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReleaseGroup200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateReleaseGroup200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateReleaseGroup200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateReleaseGroup200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateReleaseGroup200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateReleaseGroup200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganizationId

`func (o *CreateReleaseGroup200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateReleaseGroup200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateReleaseGroup200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *CreateReleaseGroup200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPolicyId

`func (o *CreateReleaseGroup200Response) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CreateReleaseGroup200Response) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CreateReleaseGroup200Response) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *CreateReleaseGroup200Response) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *CreateReleaseGroup200Response) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *CreateReleaseGroup200Response) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSecurityPolicyId

`func (o *CreateReleaseGroup200Response) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *CreateReleaseGroup200Response) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *CreateReleaseGroup200Response) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.

### HasSecurityPolicyId

`func (o *CreateReleaseGroup200Response) HasSecurityPolicyId() bool`

HasSecurityPolicyId returns a boolean if a field has been set.

### SetSecurityPolicyIdNil

`func (o *CreateReleaseGroup200Response) SetSecurityPolicyIdNil(b bool)`

 SetSecurityPolicyIdNil sets the value for SecurityPolicyId to be an explicit nil

### UnsetSecurityPolicyId
`func (o *CreateReleaseGroup200Response) UnsetSecurityPolicyId()`

UnsetSecurityPolicyId ensures that no value is present for SecurityPolicyId, not even an explicit nil
### GetQualityPolicyId

`func (o *CreateReleaseGroup200Response) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *CreateReleaseGroup200Response) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *CreateReleaseGroup200Response) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.

### HasQualityPolicyId

`func (o *CreateReleaseGroup200Response) HasQualityPolicyId() bool`

HasQualityPolicyId returns a boolean if a field has been set.

### SetQualityPolicyIdNil

`func (o *CreateReleaseGroup200Response) SetQualityPolicyIdNil(b bool)`

 SetQualityPolicyIdNil sets the value for QualityPolicyId to be an explicit nil

### UnsetQualityPolicyId
`func (o *CreateReleaseGroup200Response) UnsetQualityPolicyId()`

UnsetQualityPolicyId ensures that no value is present for QualityPolicyId, not even an explicit nil
### GetPublicOnPortal

`func (o *CreateReleaseGroup200Response) GetPublicOnPortal() bool`

GetPublicOnPortal returns the PublicOnPortal field if non-nil, zero value otherwise.

### GetPublicOnPortalOk

`func (o *CreateReleaseGroup200Response) GetPublicOnPortalOk() (*bool, bool)`

GetPublicOnPortalOk returns a tuple with the PublicOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOnPortal

`func (o *CreateReleaseGroup200Response) SetPublicOnPortal(v bool)`

SetPublicOnPortal sets PublicOnPortal field to given value.

### HasPublicOnPortal

`func (o *CreateReleaseGroup200Response) HasPublicOnPortal() bool`

HasPublicOnPortal returns a boolean if a field has been set.

### GetReportCustomText

`func (o *CreateReleaseGroup200Response) GetReportCustomText() string`

GetReportCustomText returns the ReportCustomText field if non-nil, zero value otherwise.

### GetReportCustomTextOk

`func (o *CreateReleaseGroup200Response) GetReportCustomTextOk() (*string, bool)`

GetReportCustomTextOk returns a tuple with the ReportCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomText

`func (o *CreateReleaseGroup200Response) SetReportCustomText(v string)`

SetReportCustomText sets ReportCustomText field to given value.

### HasReportCustomText

`func (o *CreateReleaseGroup200Response) HasReportCustomText() bool`

HasReportCustomText returns a boolean if a field has been set.

### SetReportCustomTextNil

`func (o *CreateReleaseGroup200Response) SetReportCustomTextNil(b bool)`

 SetReportCustomTextNil sets the value for ReportCustomText to be an explicit nil

### UnsetReportCustomText
`func (o *CreateReleaseGroup200Response) UnsetReportCustomText()`

UnsetReportCustomText ensures that no value is present for ReportCustomText, not even an explicit nil
### GetCreatedAt

`func (o *CreateReleaseGroup200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateReleaseGroup200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateReleaseGroup200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateReleaseGroup200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateReleaseGroup200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateReleaseGroup200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateReleaseGroup200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateReleaseGroup200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


