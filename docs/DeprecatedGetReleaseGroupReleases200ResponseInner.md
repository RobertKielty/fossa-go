# DeprecatedGetReleaseGroupReleases200ResponseInner

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
**Projects** | Pointer to [**[]GetReleaseGroupById200ResponseAllOfProjectsInner**](GetReleaseGroupById200ResponseAllOfProjectsInner.md) |  | [optional] 

## Methods

### NewDeprecatedGetReleaseGroupReleases200ResponseInner

`func NewDeprecatedGetReleaseGroupReleases200ResponseInner() *DeprecatedGetReleaseGroupReleases200ResponseInner`

NewDeprecatedGetReleaseGroupReleases200ResponseInner instantiates a new DeprecatedGetReleaseGroupReleases200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedGetReleaseGroupReleases200ResponseInnerWithDefaults

`func NewDeprecatedGetReleaseGroupReleases200ResponseInnerWithDefaults() *DeprecatedGetReleaseGroupReleases200ResponseInner`

NewDeprecatedGetReleaseGroupReleases200ResponseInnerWithDefaults instantiates a new DeprecatedGetReleaseGroupReleases200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganizationId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSecurityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.

### HasSecurityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasSecurityPolicyId() bool`

HasSecurityPolicyId returns a boolean if a field has been set.

### SetSecurityPolicyIdNil

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetSecurityPolicyIdNil(b bool)`

 SetSecurityPolicyIdNil sets the value for SecurityPolicyId to be an explicit nil

### UnsetSecurityPolicyId
`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) UnsetSecurityPolicyId()`

UnsetSecurityPolicyId ensures that no value is present for SecurityPolicyId, not even an explicit nil
### GetQualityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.

### HasQualityPolicyId

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasQualityPolicyId() bool`

HasQualityPolicyId returns a boolean if a field has been set.

### SetQualityPolicyIdNil

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetQualityPolicyIdNil(b bool)`

 SetQualityPolicyIdNil sets the value for QualityPolicyId to be an explicit nil

### UnsetQualityPolicyId
`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) UnsetQualityPolicyId()`

UnsetQualityPolicyId ensures that no value is present for QualityPolicyId, not even an explicit nil
### GetPublicOnPortal

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetPublicOnPortal() bool`

GetPublicOnPortal returns the PublicOnPortal field if non-nil, zero value otherwise.

### GetPublicOnPortalOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetPublicOnPortalOk() (*bool, bool)`

GetPublicOnPortalOk returns a tuple with the PublicOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOnPortal

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetPublicOnPortal(v bool)`

SetPublicOnPortal sets PublicOnPortal field to given value.

### HasPublicOnPortal

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasPublicOnPortal() bool`

HasPublicOnPortal returns a boolean if a field has been set.

### GetReportCustomText

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetReportCustomText() string`

GetReportCustomText returns the ReportCustomText field if non-nil, zero value otherwise.

### GetReportCustomTextOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetReportCustomTextOk() (*string, bool)`

GetReportCustomTextOk returns a tuple with the ReportCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomText

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetReportCustomText(v string)`

SetReportCustomText sets ReportCustomText field to given value.

### HasReportCustomText

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasReportCustomText() bool`

HasReportCustomText returns a boolean if a field has been set.

### SetReportCustomTextNil

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetReportCustomTextNil(b bool)`

 SetReportCustomTextNil sets the value for ReportCustomText to be an explicit nil

### UnsetReportCustomText
`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) UnsetReportCustomText()`

UnsetReportCustomText ensures that no value is present for ReportCustomText, not even an explicit nil
### GetCreatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetProjects

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetProjects() []GetReleaseGroupById200ResponseAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) GetProjectsOk() (*[]GetReleaseGroupById200ResponseAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) SetProjects(v []GetReleaseGroupById200ResponseAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *DeprecatedGetReleaseGroupReleases200ResponseInner) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


