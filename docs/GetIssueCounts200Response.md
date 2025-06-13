# GetIssueCounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counts** | [**GetIssueCounts200ResponseCounts**](GetIssueCounts200ResponseCounts.md) |  | 
**TotalProjects** | **int32** |  | 
**Licensing** | Pointer to [**GetIssueCounts200ResponseLicensing**](GetIssueCounts200ResponseLicensing.md) |  | [optional] 
**Vulnerability** | Pointer to [**GetIssueCounts200ResponseVulnerability**](GetIssueCounts200ResponseVulnerability.md) |  | [optional] 
**Quality** | Pointer to [**GetIssueCounts200ResponseQuality**](GetIssueCounts200ResponseQuality.md) |  | [optional] 

## Methods

### NewGetIssueCounts200Response

`func NewGetIssueCounts200Response(counts GetIssueCounts200ResponseCounts, totalProjects int32, ) *GetIssueCounts200Response`

NewGetIssueCounts200Response instantiates a new GetIssueCounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIssueCounts200ResponseWithDefaults

`func NewGetIssueCounts200ResponseWithDefaults() *GetIssueCounts200Response`

NewGetIssueCounts200ResponseWithDefaults instantiates a new GetIssueCounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounts

`func (o *GetIssueCounts200Response) GetCounts() GetIssueCounts200ResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetIssueCounts200Response) GetCountsOk() (*GetIssueCounts200ResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetIssueCounts200Response) SetCounts(v GetIssueCounts200ResponseCounts)`

SetCounts sets Counts field to given value.


### GetTotalProjects

`func (o *GetIssueCounts200Response) GetTotalProjects() int32`

GetTotalProjects returns the TotalProjects field if non-nil, zero value otherwise.

### GetTotalProjectsOk

`func (o *GetIssueCounts200Response) GetTotalProjectsOk() (*int32, bool)`

GetTotalProjectsOk returns a tuple with the TotalProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalProjects

`func (o *GetIssueCounts200Response) SetTotalProjects(v int32)`

SetTotalProjects sets TotalProjects field to given value.


### GetLicensing

`func (o *GetIssueCounts200Response) GetLicensing() GetIssueCounts200ResponseLicensing`

GetLicensing returns the Licensing field if non-nil, zero value otherwise.

### GetLicensingOk

`func (o *GetIssueCounts200Response) GetLicensingOk() (*GetIssueCounts200ResponseLicensing, bool)`

GetLicensingOk returns a tuple with the Licensing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensing

`func (o *GetIssueCounts200Response) SetLicensing(v GetIssueCounts200ResponseLicensing)`

SetLicensing sets Licensing field to given value.

### HasLicensing

`func (o *GetIssueCounts200Response) HasLicensing() bool`

HasLicensing returns a boolean if a field has been set.

### GetVulnerability

`func (o *GetIssueCounts200Response) GetVulnerability() GetIssueCounts200ResponseVulnerability`

GetVulnerability returns the Vulnerability field if non-nil, zero value otherwise.

### GetVulnerabilityOk

`func (o *GetIssueCounts200Response) GetVulnerabilityOk() (*GetIssueCounts200ResponseVulnerability, bool)`

GetVulnerabilityOk returns a tuple with the Vulnerability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerability

`func (o *GetIssueCounts200Response) SetVulnerability(v GetIssueCounts200ResponseVulnerability)`

SetVulnerability sets Vulnerability field to given value.

### HasVulnerability

`func (o *GetIssueCounts200Response) HasVulnerability() bool`

HasVulnerability returns a boolean if a field has been set.

### GetQuality

`func (o *GetIssueCounts200Response) GetQuality() GetIssueCounts200ResponseQuality`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *GetIssueCounts200Response) GetQualityOk() (*GetIssueCounts200ResponseQuality, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *GetIssueCounts200Response) SetQuality(v GetIssueCounts200ResponseQuality)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *GetIssueCounts200Response) HasQuality() bool`

HasQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


