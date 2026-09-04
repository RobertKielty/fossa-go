# GetFossabotDependencyUpgradePR200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**VcsUrl** | **NullableString** | PR URL on the VCS provider. Null until the PR exists. | 
**FossabotUrl** | **NullableString** | PR URL in the fossabot UI. Null until the PR exists. | 
**IsAnalysisDelayed** | **bool** | True while analysis is taking longer than expected or should be retried. | 
**JobId** | **NullableString** | Creation-job id to pass back on the status endpoint. | 

## Methods

### NewGetFossabotDependencyUpgradePR200Response

`func NewGetFossabotDependencyUpgradePR200Response(status string, vcsUrl NullableString, fossabotUrl NullableString, isAnalysisDelayed bool, jobId NullableString, ) *GetFossabotDependencyUpgradePR200Response`

NewGetFossabotDependencyUpgradePR200Response instantiates a new GetFossabotDependencyUpgradePR200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFossabotDependencyUpgradePR200ResponseWithDefaults

`func NewGetFossabotDependencyUpgradePR200ResponseWithDefaults() *GetFossabotDependencyUpgradePR200Response`

NewGetFossabotDependencyUpgradePR200ResponseWithDefaults instantiates a new GetFossabotDependencyUpgradePR200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetFossabotDependencyUpgradePR200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetFossabotDependencyUpgradePR200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetFossabotDependencyUpgradePR200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVcsUrl

`func (o *GetFossabotDependencyUpgradePR200Response) GetVcsUrl() string`

GetVcsUrl returns the VcsUrl field if non-nil, zero value otherwise.

### GetVcsUrlOk

`func (o *GetFossabotDependencyUpgradePR200Response) GetVcsUrlOk() (*string, bool)`

GetVcsUrlOk returns a tuple with the VcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUrl

`func (o *GetFossabotDependencyUpgradePR200Response) SetVcsUrl(v string)`

SetVcsUrl sets VcsUrl field to given value.


### SetVcsUrlNil

`func (o *GetFossabotDependencyUpgradePR200Response) SetVcsUrlNil(b bool)`

 SetVcsUrlNil sets the value for VcsUrl to be an explicit nil

### UnsetVcsUrl
`func (o *GetFossabotDependencyUpgradePR200Response) UnsetVcsUrl()`

UnsetVcsUrl ensures that no value is present for VcsUrl, not even an explicit nil
### GetFossabotUrl

`func (o *GetFossabotDependencyUpgradePR200Response) GetFossabotUrl() string`

GetFossabotUrl returns the FossabotUrl field if non-nil, zero value otherwise.

### GetFossabotUrlOk

`func (o *GetFossabotDependencyUpgradePR200Response) GetFossabotUrlOk() (*string, bool)`

GetFossabotUrlOk returns a tuple with the FossabotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFossabotUrl

`func (o *GetFossabotDependencyUpgradePR200Response) SetFossabotUrl(v string)`

SetFossabotUrl sets FossabotUrl field to given value.


### SetFossabotUrlNil

`func (o *GetFossabotDependencyUpgradePR200Response) SetFossabotUrlNil(b bool)`

 SetFossabotUrlNil sets the value for FossabotUrl to be an explicit nil

### UnsetFossabotUrl
`func (o *GetFossabotDependencyUpgradePR200Response) UnsetFossabotUrl()`

UnsetFossabotUrl ensures that no value is present for FossabotUrl, not even an explicit nil
### GetIsAnalysisDelayed

`func (o *GetFossabotDependencyUpgradePR200Response) GetIsAnalysisDelayed() bool`

GetIsAnalysisDelayed returns the IsAnalysisDelayed field if non-nil, zero value otherwise.

### GetIsAnalysisDelayedOk

`func (o *GetFossabotDependencyUpgradePR200Response) GetIsAnalysisDelayedOk() (*bool, bool)`

GetIsAnalysisDelayedOk returns a tuple with the IsAnalysisDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnalysisDelayed

`func (o *GetFossabotDependencyUpgradePR200Response) SetIsAnalysisDelayed(v bool)`

SetIsAnalysisDelayed sets IsAnalysisDelayed field to given value.


### GetJobId

`func (o *GetFossabotDependencyUpgradePR200Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GetFossabotDependencyUpgradePR200Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GetFossabotDependencyUpgradePR200Response) SetJobId(v string)`

SetJobId sets JobId field to given value.


### SetJobIdNil

`func (o *GetFossabotDependencyUpgradePR200Response) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *GetFossabotDependencyUpgradePR200Response) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


