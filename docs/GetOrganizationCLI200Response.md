# GetOrganizationCLI200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | Pointer to **int32** | Unique organization identifier | [optional] 
**UsesSAML** | Pointer to **bool** | True if organization uses SAML, otherwise False. | [optional] 
**SupportsCliLicenseScanning** | Pointer to **bool** | True if organization supports CLI license scanning (native), otherwise False. | [optional] 
**SupportsAnalyzedRevisionsQuery** | Pointer to **bool** | True if organization supports analyzed revision query, otherwise False. | [optional] 
**SupportsDependenciesCachePolling** | Pointer to **bool** | True if organization supports dependencies cache polling, otherwise False. | [optional] 
**SupportsIssueDiffs** | Pointer to **bool** | True if organization supports issue diffing, otherwise False. | [optional] 
**DefaultVendoredDependencyScanType** | Pointer to **string** | Default vendor scanning preferred by the organization. | [optional] 
**SupportsNativeContainerScans** | Pointer to **bool** | True if organization supports native container scanning, otherwise False. | [optional] 
**RequireFullFileUploads** | Pointer to **bool** | True if the organization wants CLI-side license scans to upload full file contents instead of just the match string | [optional] 
**SupportsPathDependency** | Pointer to **bool** | True if the organization supports path dependencies. Otherwise False. | [optional] 
**SupportsFirstPartyScans** | Pointer to **bool** | True if the FOSSA instance supports first-party scans. Otherwise False. | [optional] 
**DefaultToFirstPartyScans** | Pointer to **bool** | True if the organization defaults to first-party scans. Otherwise False. | [optional] 
**CustomLicenseScanConfigs** | Pointer to [**[]GetOrganizationCLI200ResponseCustomLicenseScanConfigsInner**](GetOrganizationCLI200ResponseCustomLicenseScanConfigsInner.md) | Configuration for custom license scans | [optional] 

## Methods

### NewGetOrganizationCLI200Response

`func NewGetOrganizationCLI200Response() *GetOrganizationCLI200Response`

NewGetOrganizationCLI200Response instantiates a new GetOrganizationCLI200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationCLI200ResponseWithDefaults

`func NewGetOrganizationCLI200ResponseWithDefaults() *GetOrganizationCLI200Response`

NewGetOrganizationCLI200ResponseWithDefaults instantiates a new GetOrganizationCLI200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *GetOrganizationCLI200Response) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetOrganizationCLI200Response) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetOrganizationCLI200Response) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetOrganizationCLI200Response) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetUsesSAML

`func (o *GetOrganizationCLI200Response) GetUsesSAML() bool`

GetUsesSAML returns the UsesSAML field if non-nil, zero value otherwise.

### GetUsesSAMLOk

`func (o *GetOrganizationCLI200Response) GetUsesSAMLOk() (*bool, bool)`

GetUsesSAMLOk returns a tuple with the UsesSAML field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesSAML

`func (o *GetOrganizationCLI200Response) SetUsesSAML(v bool)`

SetUsesSAML sets UsesSAML field to given value.

### HasUsesSAML

`func (o *GetOrganizationCLI200Response) HasUsesSAML() bool`

HasUsesSAML returns a boolean if a field has been set.

### GetSupportsCliLicenseScanning

`func (o *GetOrganizationCLI200Response) GetSupportsCliLicenseScanning() bool`

GetSupportsCliLicenseScanning returns the SupportsCliLicenseScanning field if non-nil, zero value otherwise.

### GetSupportsCliLicenseScanningOk

`func (o *GetOrganizationCLI200Response) GetSupportsCliLicenseScanningOk() (*bool, bool)`

GetSupportsCliLicenseScanningOk returns a tuple with the SupportsCliLicenseScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCliLicenseScanning

`func (o *GetOrganizationCLI200Response) SetSupportsCliLicenseScanning(v bool)`

SetSupportsCliLicenseScanning sets SupportsCliLicenseScanning field to given value.

### HasSupportsCliLicenseScanning

`func (o *GetOrganizationCLI200Response) HasSupportsCliLicenseScanning() bool`

HasSupportsCliLicenseScanning returns a boolean if a field has been set.

### GetSupportsAnalyzedRevisionsQuery

`func (o *GetOrganizationCLI200Response) GetSupportsAnalyzedRevisionsQuery() bool`

GetSupportsAnalyzedRevisionsQuery returns the SupportsAnalyzedRevisionsQuery field if non-nil, zero value otherwise.

### GetSupportsAnalyzedRevisionsQueryOk

`func (o *GetOrganizationCLI200Response) GetSupportsAnalyzedRevisionsQueryOk() (*bool, bool)`

GetSupportsAnalyzedRevisionsQueryOk returns a tuple with the SupportsAnalyzedRevisionsQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAnalyzedRevisionsQuery

`func (o *GetOrganizationCLI200Response) SetSupportsAnalyzedRevisionsQuery(v bool)`

SetSupportsAnalyzedRevisionsQuery sets SupportsAnalyzedRevisionsQuery field to given value.

### HasSupportsAnalyzedRevisionsQuery

`func (o *GetOrganizationCLI200Response) HasSupportsAnalyzedRevisionsQuery() bool`

HasSupportsAnalyzedRevisionsQuery returns a boolean if a field has been set.

### GetSupportsDependenciesCachePolling

`func (o *GetOrganizationCLI200Response) GetSupportsDependenciesCachePolling() bool`

GetSupportsDependenciesCachePolling returns the SupportsDependenciesCachePolling field if non-nil, zero value otherwise.

### GetSupportsDependenciesCachePollingOk

`func (o *GetOrganizationCLI200Response) GetSupportsDependenciesCachePollingOk() (*bool, bool)`

GetSupportsDependenciesCachePollingOk returns a tuple with the SupportsDependenciesCachePolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDependenciesCachePolling

`func (o *GetOrganizationCLI200Response) SetSupportsDependenciesCachePolling(v bool)`

SetSupportsDependenciesCachePolling sets SupportsDependenciesCachePolling field to given value.

### HasSupportsDependenciesCachePolling

`func (o *GetOrganizationCLI200Response) HasSupportsDependenciesCachePolling() bool`

HasSupportsDependenciesCachePolling returns a boolean if a field has been set.

### GetSupportsIssueDiffs

`func (o *GetOrganizationCLI200Response) GetSupportsIssueDiffs() bool`

GetSupportsIssueDiffs returns the SupportsIssueDiffs field if non-nil, zero value otherwise.

### GetSupportsIssueDiffsOk

`func (o *GetOrganizationCLI200Response) GetSupportsIssueDiffsOk() (*bool, bool)`

GetSupportsIssueDiffsOk returns a tuple with the SupportsIssueDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIssueDiffs

`func (o *GetOrganizationCLI200Response) SetSupportsIssueDiffs(v bool)`

SetSupportsIssueDiffs sets SupportsIssueDiffs field to given value.

### HasSupportsIssueDiffs

`func (o *GetOrganizationCLI200Response) HasSupportsIssueDiffs() bool`

HasSupportsIssueDiffs returns a boolean if a field has been set.

### GetDefaultVendoredDependencyScanType

`func (o *GetOrganizationCLI200Response) GetDefaultVendoredDependencyScanType() string`

GetDefaultVendoredDependencyScanType returns the DefaultVendoredDependencyScanType field if non-nil, zero value otherwise.

### GetDefaultVendoredDependencyScanTypeOk

`func (o *GetOrganizationCLI200Response) GetDefaultVendoredDependencyScanTypeOk() (*string, bool)`

GetDefaultVendoredDependencyScanTypeOk returns a tuple with the DefaultVendoredDependencyScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVendoredDependencyScanType

`func (o *GetOrganizationCLI200Response) SetDefaultVendoredDependencyScanType(v string)`

SetDefaultVendoredDependencyScanType sets DefaultVendoredDependencyScanType field to given value.

### HasDefaultVendoredDependencyScanType

`func (o *GetOrganizationCLI200Response) HasDefaultVendoredDependencyScanType() bool`

HasDefaultVendoredDependencyScanType returns a boolean if a field has been set.

### GetSupportsNativeContainerScans

`func (o *GetOrganizationCLI200Response) GetSupportsNativeContainerScans() bool`

GetSupportsNativeContainerScans returns the SupportsNativeContainerScans field if non-nil, zero value otherwise.

### GetSupportsNativeContainerScansOk

`func (o *GetOrganizationCLI200Response) GetSupportsNativeContainerScansOk() (*bool, bool)`

GetSupportsNativeContainerScansOk returns a tuple with the SupportsNativeContainerScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNativeContainerScans

`func (o *GetOrganizationCLI200Response) SetSupportsNativeContainerScans(v bool)`

SetSupportsNativeContainerScans sets SupportsNativeContainerScans field to given value.

### HasSupportsNativeContainerScans

`func (o *GetOrganizationCLI200Response) HasSupportsNativeContainerScans() bool`

HasSupportsNativeContainerScans returns a boolean if a field has been set.

### GetRequireFullFileUploads

`func (o *GetOrganizationCLI200Response) GetRequireFullFileUploads() bool`

GetRequireFullFileUploads returns the RequireFullFileUploads field if non-nil, zero value otherwise.

### GetRequireFullFileUploadsOk

`func (o *GetOrganizationCLI200Response) GetRequireFullFileUploadsOk() (*bool, bool)`

GetRequireFullFileUploadsOk returns a tuple with the RequireFullFileUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireFullFileUploads

`func (o *GetOrganizationCLI200Response) SetRequireFullFileUploads(v bool)`

SetRequireFullFileUploads sets RequireFullFileUploads field to given value.

### HasRequireFullFileUploads

`func (o *GetOrganizationCLI200Response) HasRequireFullFileUploads() bool`

HasRequireFullFileUploads returns a boolean if a field has been set.

### GetSupportsPathDependency

`func (o *GetOrganizationCLI200Response) GetSupportsPathDependency() bool`

GetSupportsPathDependency returns the SupportsPathDependency field if non-nil, zero value otherwise.

### GetSupportsPathDependencyOk

`func (o *GetOrganizationCLI200Response) GetSupportsPathDependencyOk() (*bool, bool)`

GetSupportsPathDependencyOk returns a tuple with the SupportsPathDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsPathDependency

`func (o *GetOrganizationCLI200Response) SetSupportsPathDependency(v bool)`

SetSupportsPathDependency sets SupportsPathDependency field to given value.

### HasSupportsPathDependency

`func (o *GetOrganizationCLI200Response) HasSupportsPathDependency() bool`

HasSupportsPathDependency returns a boolean if a field has been set.

### GetSupportsFirstPartyScans

`func (o *GetOrganizationCLI200Response) GetSupportsFirstPartyScans() bool`

GetSupportsFirstPartyScans returns the SupportsFirstPartyScans field if non-nil, zero value otherwise.

### GetSupportsFirstPartyScansOk

`func (o *GetOrganizationCLI200Response) GetSupportsFirstPartyScansOk() (*bool, bool)`

GetSupportsFirstPartyScansOk returns a tuple with the SupportsFirstPartyScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsFirstPartyScans

`func (o *GetOrganizationCLI200Response) SetSupportsFirstPartyScans(v bool)`

SetSupportsFirstPartyScans sets SupportsFirstPartyScans field to given value.

### HasSupportsFirstPartyScans

`func (o *GetOrganizationCLI200Response) HasSupportsFirstPartyScans() bool`

HasSupportsFirstPartyScans returns a boolean if a field has been set.

### GetDefaultToFirstPartyScans

`func (o *GetOrganizationCLI200Response) GetDefaultToFirstPartyScans() bool`

GetDefaultToFirstPartyScans returns the DefaultToFirstPartyScans field if non-nil, zero value otherwise.

### GetDefaultToFirstPartyScansOk

`func (o *GetOrganizationCLI200Response) GetDefaultToFirstPartyScansOk() (*bool, bool)`

GetDefaultToFirstPartyScansOk returns a tuple with the DefaultToFirstPartyScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultToFirstPartyScans

`func (o *GetOrganizationCLI200Response) SetDefaultToFirstPartyScans(v bool)`

SetDefaultToFirstPartyScans sets DefaultToFirstPartyScans field to given value.

### HasDefaultToFirstPartyScans

`func (o *GetOrganizationCLI200Response) HasDefaultToFirstPartyScans() bool`

HasDefaultToFirstPartyScans returns a boolean if a field has been set.

### GetCustomLicenseScanConfigs

`func (o *GetOrganizationCLI200Response) GetCustomLicenseScanConfigs() []GetOrganizationCLI200ResponseCustomLicenseScanConfigsInner`

GetCustomLicenseScanConfigs returns the CustomLicenseScanConfigs field if non-nil, zero value otherwise.

### GetCustomLicenseScanConfigsOk

`func (o *GetOrganizationCLI200Response) GetCustomLicenseScanConfigsOk() (*[]GetOrganizationCLI200ResponseCustomLicenseScanConfigsInner, bool)`

GetCustomLicenseScanConfigsOk returns a tuple with the CustomLicenseScanConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLicenseScanConfigs

`func (o *GetOrganizationCLI200Response) SetCustomLicenseScanConfigs(v []GetOrganizationCLI200ResponseCustomLicenseScanConfigsInner)`

SetCustomLicenseScanConfigs sets CustomLicenseScanConfigs field to given value.

### HasCustomLicenseScanConfigs

`func (o *GetOrganizationCLI200Response) HasCustomLicenseScanConfigs() bool`

HasCustomLicenseScanConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


