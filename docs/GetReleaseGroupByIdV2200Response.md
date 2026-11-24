# GetReleaseGroupByIdV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Title** | **string** |  | 
**OrganizationId** | **int32** |  | 
**PublicOnPortal** | **bool** |  | 
**PolicyId** | **NullableInt32** |  | 
**SecurityPolicyId** | **NullableInt32** |  | 
**QualityPolicyId** | **NullableInt32** |  | 
**ReportCustomText** | **NullableString** | The release group&#39;s report text, or the organization&#39;s default report text when the release group has none. | 
**DefaultSavedOptionLicensingReport** | **NullableInt32** |  | 
**DefaultSavedOptionSbom** | **NullableInt32** |  | 

## Methods

### NewGetReleaseGroupByIdV2200Response

`func NewGetReleaseGroupByIdV2200Response(id int32, title string, organizationId int32, publicOnPortal bool, policyId NullableInt32, securityPolicyId NullableInt32, qualityPolicyId NullableInt32, reportCustomText NullableString, defaultSavedOptionLicensingReport NullableInt32, defaultSavedOptionSbom NullableInt32, ) *GetReleaseGroupByIdV2200Response`

NewGetReleaseGroupByIdV2200Response instantiates a new GetReleaseGroupByIdV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReleaseGroupByIdV2200ResponseWithDefaults

`func NewGetReleaseGroupByIdV2200ResponseWithDefaults() *GetReleaseGroupByIdV2200Response`

NewGetReleaseGroupByIdV2200ResponseWithDefaults instantiates a new GetReleaseGroupByIdV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReleaseGroupByIdV2200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReleaseGroupByIdV2200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReleaseGroupByIdV2200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetTitle

`func (o *GetReleaseGroupByIdV2200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetReleaseGroupByIdV2200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetReleaseGroupByIdV2200Response) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetOrganizationId

`func (o *GetReleaseGroupByIdV2200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetReleaseGroupByIdV2200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetReleaseGroupByIdV2200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.


### GetPublicOnPortal

`func (o *GetReleaseGroupByIdV2200Response) GetPublicOnPortal() bool`

GetPublicOnPortal returns the PublicOnPortal field if non-nil, zero value otherwise.

### GetPublicOnPortalOk

`func (o *GetReleaseGroupByIdV2200Response) GetPublicOnPortalOk() (*bool, bool)`

GetPublicOnPortalOk returns a tuple with the PublicOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOnPortal

`func (o *GetReleaseGroupByIdV2200Response) SetPublicOnPortal(v bool)`

SetPublicOnPortal sets PublicOnPortal field to given value.


### GetPolicyId

`func (o *GetReleaseGroupByIdV2200Response) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GetReleaseGroupByIdV2200Response) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GetReleaseGroupByIdV2200Response) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.


### SetPolicyIdNil

`func (o *GetReleaseGroupByIdV2200Response) SetPolicyIdNil(b bool)`

 SetPolicyIdNil sets the value for PolicyId to be an explicit nil

### UnsetPolicyId
`func (o *GetReleaseGroupByIdV2200Response) UnsetPolicyId()`

UnsetPolicyId ensures that no value is present for PolicyId, not even an explicit nil
### GetSecurityPolicyId

`func (o *GetReleaseGroupByIdV2200Response) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *GetReleaseGroupByIdV2200Response) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *GetReleaseGroupByIdV2200Response) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.


### SetSecurityPolicyIdNil

`func (o *GetReleaseGroupByIdV2200Response) SetSecurityPolicyIdNil(b bool)`

 SetSecurityPolicyIdNil sets the value for SecurityPolicyId to be an explicit nil

### UnsetSecurityPolicyId
`func (o *GetReleaseGroupByIdV2200Response) UnsetSecurityPolicyId()`

UnsetSecurityPolicyId ensures that no value is present for SecurityPolicyId, not even an explicit nil
### GetQualityPolicyId

`func (o *GetReleaseGroupByIdV2200Response) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *GetReleaseGroupByIdV2200Response) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *GetReleaseGroupByIdV2200Response) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.


### SetQualityPolicyIdNil

`func (o *GetReleaseGroupByIdV2200Response) SetQualityPolicyIdNil(b bool)`

 SetQualityPolicyIdNil sets the value for QualityPolicyId to be an explicit nil

### UnsetQualityPolicyId
`func (o *GetReleaseGroupByIdV2200Response) UnsetQualityPolicyId()`

UnsetQualityPolicyId ensures that no value is present for QualityPolicyId, not even an explicit nil
### GetReportCustomText

`func (o *GetReleaseGroupByIdV2200Response) GetReportCustomText() string`

GetReportCustomText returns the ReportCustomText field if non-nil, zero value otherwise.

### GetReportCustomTextOk

`func (o *GetReleaseGroupByIdV2200Response) GetReportCustomTextOk() (*string, bool)`

GetReportCustomTextOk returns a tuple with the ReportCustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCustomText

`func (o *GetReleaseGroupByIdV2200Response) SetReportCustomText(v string)`

SetReportCustomText sets ReportCustomText field to given value.


### SetReportCustomTextNil

`func (o *GetReleaseGroupByIdV2200Response) SetReportCustomTextNil(b bool)`

 SetReportCustomTextNil sets the value for ReportCustomText to be an explicit nil

### UnsetReportCustomText
`func (o *GetReleaseGroupByIdV2200Response) UnsetReportCustomText()`

UnsetReportCustomText ensures that no value is present for ReportCustomText, not even an explicit nil
### GetDefaultSavedOptionLicensingReport

`func (o *GetReleaseGroupByIdV2200Response) GetDefaultSavedOptionLicensingReport() int32`

GetDefaultSavedOptionLicensingReport returns the DefaultSavedOptionLicensingReport field if non-nil, zero value otherwise.

### GetDefaultSavedOptionLicensingReportOk

`func (o *GetReleaseGroupByIdV2200Response) GetDefaultSavedOptionLicensingReportOk() (*int32, bool)`

GetDefaultSavedOptionLicensingReportOk returns a tuple with the DefaultSavedOptionLicensingReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSavedOptionLicensingReport

`func (o *GetReleaseGroupByIdV2200Response) SetDefaultSavedOptionLicensingReport(v int32)`

SetDefaultSavedOptionLicensingReport sets DefaultSavedOptionLicensingReport field to given value.


### SetDefaultSavedOptionLicensingReportNil

`func (o *GetReleaseGroupByIdV2200Response) SetDefaultSavedOptionLicensingReportNil(b bool)`

 SetDefaultSavedOptionLicensingReportNil sets the value for DefaultSavedOptionLicensingReport to be an explicit nil

### UnsetDefaultSavedOptionLicensingReport
`func (o *GetReleaseGroupByIdV2200Response) UnsetDefaultSavedOptionLicensingReport()`

UnsetDefaultSavedOptionLicensingReport ensures that no value is present for DefaultSavedOptionLicensingReport, not even an explicit nil
### GetDefaultSavedOptionSbom

`func (o *GetReleaseGroupByIdV2200Response) GetDefaultSavedOptionSbom() int32`

GetDefaultSavedOptionSbom returns the DefaultSavedOptionSbom field if non-nil, zero value otherwise.

### GetDefaultSavedOptionSbomOk

`func (o *GetReleaseGroupByIdV2200Response) GetDefaultSavedOptionSbomOk() (*int32, bool)`

GetDefaultSavedOptionSbomOk returns a tuple with the DefaultSavedOptionSbom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSavedOptionSbom

`func (o *GetReleaseGroupByIdV2200Response) SetDefaultSavedOptionSbom(v int32)`

SetDefaultSavedOptionSbom sets DefaultSavedOptionSbom field to given value.


### SetDefaultSavedOptionSbomNil

`func (o *GetReleaseGroupByIdV2200Response) SetDefaultSavedOptionSbomNil(b bool)`

 SetDefaultSavedOptionSbomNil sets the value for DefaultSavedOptionSbom to be an explicit nil

### UnsetDefaultSavedOptionSbom
`func (o *GetReleaseGroupByIdV2200Response) UnsetDefaultSavedOptionSbom()`

UnsetDefaultSavedOptionSbom ensures that no value is present for DefaultSavedOptionSbom, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


