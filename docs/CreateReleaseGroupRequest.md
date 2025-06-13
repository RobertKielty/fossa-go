# CreateReleaseGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**LicensingPolicyId** | Pointer to **int32** |  | [optional] 
**SecurityPolicyId** | Pointer to **int32** |  | [optional] 
**QualityPolicyId** | Pointer to **int32** |  | [optional] 
**PublicOnPortal** | Pointer to **bool** |  | [optional] 
**IssueTrackerType** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to **[]int32** |  | [optional] 
**Release** | [**CreateReleaseGroupRequestRelease**](CreateReleaseGroupRequestRelease.md) |  | 

## Methods

### NewCreateReleaseGroupRequest

`func NewCreateReleaseGroupRequest(title string, release CreateReleaseGroupRequestRelease, ) *CreateReleaseGroupRequest`

NewCreateReleaseGroupRequest instantiates a new CreateReleaseGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReleaseGroupRequestWithDefaults

`func NewCreateReleaseGroupRequestWithDefaults() *CreateReleaseGroupRequest`

NewCreateReleaseGroupRequestWithDefaults instantiates a new CreateReleaseGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateReleaseGroupRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateReleaseGroupRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateReleaseGroupRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLicensingPolicyId

`func (o *CreateReleaseGroupRequest) GetLicensingPolicyId() int32`

GetLicensingPolicyId returns the LicensingPolicyId field if non-nil, zero value otherwise.

### GetLicensingPolicyIdOk

`func (o *CreateReleaseGroupRequest) GetLicensingPolicyIdOk() (*int32, bool)`

GetLicensingPolicyIdOk returns a tuple with the LicensingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingPolicyId

`func (o *CreateReleaseGroupRequest) SetLicensingPolicyId(v int32)`

SetLicensingPolicyId sets LicensingPolicyId field to given value.

### HasLicensingPolicyId

`func (o *CreateReleaseGroupRequest) HasLicensingPolicyId() bool`

HasLicensingPolicyId returns a boolean if a field has been set.

### GetSecurityPolicyId

`func (o *CreateReleaseGroupRequest) GetSecurityPolicyId() int32`

GetSecurityPolicyId returns the SecurityPolicyId field if non-nil, zero value otherwise.

### GetSecurityPolicyIdOk

`func (o *CreateReleaseGroupRequest) GetSecurityPolicyIdOk() (*int32, bool)`

GetSecurityPolicyIdOk returns a tuple with the SecurityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyId

`func (o *CreateReleaseGroupRequest) SetSecurityPolicyId(v int32)`

SetSecurityPolicyId sets SecurityPolicyId field to given value.

### HasSecurityPolicyId

`func (o *CreateReleaseGroupRequest) HasSecurityPolicyId() bool`

HasSecurityPolicyId returns a boolean if a field has been set.

### GetQualityPolicyId

`func (o *CreateReleaseGroupRequest) GetQualityPolicyId() int32`

GetQualityPolicyId returns the QualityPolicyId field if non-nil, zero value otherwise.

### GetQualityPolicyIdOk

`func (o *CreateReleaseGroupRequest) GetQualityPolicyIdOk() (*int32, bool)`

GetQualityPolicyIdOk returns a tuple with the QualityPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyId

`func (o *CreateReleaseGroupRequest) SetQualityPolicyId(v int32)`

SetQualityPolicyId sets QualityPolicyId field to given value.

### HasQualityPolicyId

`func (o *CreateReleaseGroupRequest) HasQualityPolicyId() bool`

HasQualityPolicyId returns a boolean if a field has been set.

### GetPublicOnPortal

`func (o *CreateReleaseGroupRequest) GetPublicOnPortal() bool`

GetPublicOnPortal returns the PublicOnPortal field if non-nil, zero value otherwise.

### GetPublicOnPortalOk

`func (o *CreateReleaseGroupRequest) GetPublicOnPortalOk() (*bool, bool)`

GetPublicOnPortalOk returns a tuple with the PublicOnPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicOnPortal

`func (o *CreateReleaseGroupRequest) SetPublicOnPortal(v bool)`

SetPublicOnPortal sets PublicOnPortal field to given value.

### HasPublicOnPortal

`func (o *CreateReleaseGroupRequest) HasPublicOnPortal() bool`

HasPublicOnPortal returns a boolean if a field has been set.

### GetIssueTrackerType

`func (o *CreateReleaseGroupRequest) GetIssueTrackerType() string`

GetIssueTrackerType returns the IssueTrackerType field if non-nil, zero value otherwise.

### GetIssueTrackerTypeOk

`func (o *CreateReleaseGroupRequest) GetIssueTrackerTypeOk() (*string, bool)`

GetIssueTrackerTypeOk returns a tuple with the IssueTrackerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerType

`func (o *CreateReleaseGroupRequest) SetIssueTrackerType(v string)`

SetIssueTrackerType sets IssueTrackerType field to given value.

### HasIssueTrackerType

`func (o *CreateReleaseGroupRequest) HasIssueTrackerType() bool`

HasIssueTrackerType returns a boolean if a field has been set.

### GetTeams

`func (o *CreateReleaseGroupRequest) GetTeams() []int32`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *CreateReleaseGroupRequest) GetTeamsOk() (*[]int32, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *CreateReleaseGroupRequest) SetTeams(v []int32)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *CreateReleaseGroupRequest) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetRelease

`func (o *CreateReleaseGroupRequest) GetRelease() CreateReleaseGroupRequestRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *CreateReleaseGroupRequest) GetReleaseOk() (*CreateReleaseGroupRequestRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *CreateReleaseGroupRequest) SetRelease(v CreateReleaseGroupRequestRelease)`

SetRelease sets Release field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


