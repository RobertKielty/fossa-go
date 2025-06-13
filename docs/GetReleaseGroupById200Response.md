# GetReleaseGroupById200Response

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
**Releases** | Pointer to [**[]GetReleaseGroupById200ResponseAllOfReleasesInner**](GetReleaseGroupById200ResponseAllOfReleasesInner.md) |  | [optional] 
**Projects** | Pointer to [**[]GetReleaseGroupById200ResponseAllOfProjectsInner**](GetReleaseGroupById200ResponseAllOfProjectsInner.md) |  | [optional] 
**Scans** | Pointer to [**[]GetReleaseGroupById200ResponseAllOfScansInner**](GetReleaseGroupById200ResponseAllOfScansInner.md) |  | [optional] 

## Methods

### NewGetReleaseGroupById200Response

`func NewGetReleaseGroupById200Response() *GetReleaseGroupById200Response`

NewGetReleaseGroupById200Response instantiates a new GetReleaseGroupById200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupById200ResponseWithDefaults

`func NewGetReleaseGroupById200ResponseWithDefaults() *GetReleaseGroupById200Response`

NewGetReleaseGroupById200ResponseWithDefaults instantiates a new GetReleaseGroupById200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupById200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupById200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupById200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetReleaseGroupById200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetReleaseGroupById200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupById200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupById200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetReleaseGroupById200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetReleaseGroupById200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetReleaseGroupById200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetReleaseGroupById200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetReleaseGroupById200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPolicyId

`func (o *GetReleaseGroupById200Response) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GetReleaseGroupById200Response) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GetReleaseGroupById200Response) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *GetReleaseGroupById200Response) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### SetPolicyIdNil

`func (o *GetReleaseGroupById200Response) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *GetReleaseGroupById200Response) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSecurityPolicyId

`func (o *GetReleaseGroupById200Response) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *GetReleaseGroupById200Response) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *GetReleaseGroupById200Response) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.

### HasSecurityPolicyId

`func (o *GetReleaseGroupById200Response) HasSecurityPolicyId() bool`

HasSecurityPolicyId returns a boolean if a field has been set.

### SetSecurityPolicyIdNil

`func (o *GetReleaseGroupById200Response) SetSecurityPolicyIdNil(b bool)`

 SetSecurityPolicyIdNil sets the value for SecurityPolicyId to be an explicit nil

### UnsetSecurityPolicyId
`func (o *GetReleaseGroupById200Response) UnsetSecurityPolicyId()`

UnsetSecurityPolicyId ensures that no value is present for SecurityPolicyId, not even an explicit nil
### GetQualityPolicyId

`func (o *GetReleaseGroupById200Response) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *GetReleaseGroupById200Response) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *GetReleaseGroupById200Response) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.

### HasQualityPolicyId

`func (o *GetReleaseGroupById200Response) HasQualityPolicyId() bool`

HasQualityPolicyId returns a boolean if a field has been set.

### SetQualityPolicyIdNil

`func (o *GetReleaseGroupById200Response) SetQualityPolicyIdNil(b bool)`

 SetQualityPolicyIdNil sets the value for QualityPolicyId to be an explicit nil

### UnsetQualityPolicyId
`func (o *GetReleaseGroupById200Response) UnsetQualityPolicyId()`

UnsetQualityPolicyId ensures that no value is present for QualityPolicyId, not even an explicit nil
### GetPublicOnPortal

`func (o *GetReleaseGroupById200Response) GetPublicOnPortal() bool`

GetPublicOnPortal returns the PublicOnPortal field if non-nil, zero value otherwise.

### GetPublicOnPortalOk

`func (o *GetReleaseGroupById200Response) GetPublicOnPortalOk() (*bool, bool)`

GetPublicOnPortalOk returns a tuple with the PublicOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOnPortal

`func (o *GetReleaseGroupById200Response) SetPublicOnPortal(v bool)`

SetPublicOnPortal sets PublicOnPortal field to given value.

### HasPublicOnPortal

`func (o *GetReleaseGroupById200Response) HasPublicOnPortal() bool`

HasPublicOnPortal returns a boolean if a field has been set.

### GetReportCustomText

`func (o *GetReleaseGroupById200Response) GetReportCustomText() string`

GetReportCustomText returns the ReportCustomText field if non-nil, zero value otherwise.

### GetReportCustomTextOk

`func (o *GetReleaseGroupById200Response) GetReportCustomTextOk() (*string, bool)`

GetReportCustomTextOk returns a tuple with the ReportCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomText

`func (o *GetReleaseGroupById200Response) SetReportCustomText(v string)`

SetReportCustomText sets ReportCustomText field to given value.

### HasReportCustomText

`func (o *GetReleaseGroupById200Response) HasReportCustomText() bool`

HasReportCustomText returns a boolean if a field has been set.

### SetReportCustomTextNil

`func (o *GetReleaseGroupById200Response) SetReportCustomTextNil(b bool)`

 SetReportCustomTextNil sets the value for ReportCustomText to be an explicit nil

### UnsetReportCustomText
`func (o *GetReleaseGroupById200Response) UnsetReportCustomText()`

UnsetReportCustomText ensures that no value is present for ReportCustomText, not even an explicit nil
### GetCreatedAt

`func (o *GetReleaseGroupById200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReleaseGroupById200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReleaseGroupById200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetReleaseGroupById200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetReleaseGroupById200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetReleaseGroupById200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetReleaseGroupById200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetReleaseGroupById200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetReleases

`func (o *GetReleaseGroupById200Response) GetReleases() []GetReleaseGroupById200ResponseAllOfReleasesInner`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *GetReleaseGroupById200Response) GetReleasesOk() (*[]GetReleaseGroupById200ResponseAllOfReleasesInner, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *GetReleaseGroupById200Response) SetReleases(v []GetReleaseGroupById200ResponseAllOfReleasesInner)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *GetReleaseGroupById200Response) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetProjects

`func (o *GetReleaseGroupById200Response) GetProjects() []GetReleaseGroupById200ResponseAllOfProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *GetReleaseGroupById200Response) GetProjectsOk() (*[]GetReleaseGroupById200ResponseAllOfProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *GetReleaseGroupById200Response) SetProjects(v []GetReleaseGroupById200ResponseAllOfProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *GetReleaseGroupById200Response) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetScans

`func (o *GetReleaseGroupById200Response) GetScans() []GetReleaseGroupById200ResponseAllOfScansInner`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *GetReleaseGroupById200Response) GetScansOk() (*[]GetReleaseGroupById200ResponseAllOfScansInner, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *GetReleaseGroupById200Response) SetScans(v []GetReleaseGroupById200ResponseAllOfScansInner)`

SetScans sets Scans field to given value.

### HasScans

`func (o *GetReleaseGroupById200Response) HasScans() bool`

HasScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


