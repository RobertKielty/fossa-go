# GetRevisionScans200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier for the scan | [optional] 
**ScannedAt** | Pointer to **time.Time** | Timestamp when the scan was performed | [optional] 
**ScanType** | Pointer to **string** | The scan type. This endpoint only returns project scans. | [optional] 
**RevisionId** | Pointer to **string** | The locator of the revision the scan belongs to | [optional] 
**OrganizationId** | Pointer to **int32** | The organization that owns the scan | [optional] 
**ReleaseScanId** | Pointer to **int32** | Associated release scan identifier, if any | [optional] 
**LicensingPolicyVersionId** | Pointer to **int32** | The licensing policy version used during this scan | [optional] 
**SecurityPolicyVersionId** | Pointer to **int32** | The security policy version used during this scan | [optional] 
**QualityPolicyVersionId** | Pointer to **int32** | The quality policy version used during this scan | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetRevisionScans200ResponseResultsInner

`func NewGetRevisionScans200ResponseResultsInner() *GetRevisionScans200ResponseResultsInner`

NewGetRevisionScans200ResponseResultsInner instantiates a new GetRevisionScans200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevisionScans200ResponseResultsInnerWithDefaults

`func NewGetRevisionScans200ResponseResultsInnerWithDefaults() *GetRevisionScans200ResponseResultsInner`

NewGetRevisionScans200ResponseResultsInnerWithDefaults instantiates a new GetRevisionScans200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRevisionScans200ResponseResultsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRevisionScans200ResponseResultsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetRevisionScans200ResponseResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScannedAt

`func (o *GetRevisionScans200ResponseResultsInner) GetScannedAt() time.Time`

GetScannedAt returns the ScannedAt field if non-nil, zero value otherwise.

### GetScannedAtOk

`func (o *GetRevisionScans200ResponseResultsInner) GetScannedAtOk() (*time.Time, bool)`

GetScannedAtOk returns a tuple with the ScannedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScannedAt

`func (o *GetRevisionScans200ResponseResultsInner) SetScannedAt(v time.Time)`

SetScannedAt sets ScannedAt field to given value.

### HasScannedAt

`func (o *GetRevisionScans200ResponseResultsInner) HasScannedAt() bool`

HasScannedAt returns a boolean if a field has been set.

### GetScanType

`func (o *GetRevisionScans200ResponseResultsInner) GetScanType() string`

GetScanType returns the ScanType field if non-nil, zero value otherwise.

### GetScanTypeOk

`func (o *GetRevisionScans200ResponseResultsInner) GetScanTypeOk() (*string, bool)`

GetScanTypeOk returns a tuple with the ScanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanType

`func (o *GetRevisionScans200ResponseResultsInner) SetScanType(v string)`

SetScanType sets ScanType field to given value.

### HasScanType

`func (o *GetRevisionScans200ResponseResultsInner) HasScanType() bool`

HasScanType returns a boolean if a field has been set.

### GetRevisionId

`func (o *GetRevisionScans200ResponseResultsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetRevisionScans200ResponseResultsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetRevisionScans200ResponseResultsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetOrganizationId

`func (o *GetRevisionScans200ResponseResultsInner) GetOrganizationId() int32`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetOrganizationIdOk() (*int32, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GetRevisionScans200ResponseResultsInner) SetOrganizationId(v int32)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GetRevisionScans200ResponseResultsInner) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetReleaseScanId

`func (o *GetRevisionScans200ResponseResultsInner) GetReleaseScanId() int32`

GetReleaseScanId returns the ReleaseScanId field if non-nil, zero value otherwise.

### GetReleaseScanIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetReleaseScanIdOk() (*int32, bool)`

GetReleaseScanIdOk returns a tuple with the ReleaseScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseScanId

`func (o *GetRevisionScans200ResponseResultsInner) SetReleaseScanId(v int32)`

SetReleaseScanId sets ReleaseScanId field to given value.

### HasReleaseScanId

`func (o *GetRevisionScans200ResponseResultsInner) HasReleaseScanId() bool`

HasReleaseScanId returns a boolean if a field has been set.

### GetLicensingPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) GetLicensingPolicyVersionId() int32`

GetLicensingPolicyVersionId returns the LicensingPolicyVersionId field if non-nil, zero value otherwise.

### GetLicensingPolicyVersionIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetLicensingPolicyVersionIdOk() (*int32, bool)`

GetLicensingPolicyVersionIdOk returns a tuple with the LicensingPolicyVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) SetLicensingPolicyVersionId(v int32)`

SetLicensingPolicyVersionId sets LicensingPolicyVersionId field to given value.

### HasLicensingPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) HasLicensingPolicyVersionId() bool`

HasLicensingPolicyVersionId returns a boolean if a field has been set.

### GetSecurityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) GetSecurityPolicyVersionId() int32`

GetSecurityPolicyVersionId returns the SecurityPolicyVersionId field if non-nil, zero value otherwise.

### GetSecurityPolicyVersionIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetSecurityPolicyVersionIdOk() (*int32, bool)`

GetSecurityPolicyVersionIdOk returns a tuple with the SecurityPolicyVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) SetSecurityPolicyVersionId(v int32)`

SetSecurityPolicyVersionId sets SecurityPolicyVersionId field to given value.

### HasSecurityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) HasSecurityPolicyVersionId() bool`

HasSecurityPolicyVersionId returns a boolean if a field has been set.

### GetQualityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) GetQualityPolicyVersionId() int32`

GetQualityPolicyVersionId returns the QualityPolicyVersionId field if non-nil, zero value otherwise.

### GetQualityPolicyVersionIdOk

`func (o *GetRevisionScans200ResponseResultsInner) GetQualityPolicyVersionIdOk() (*int32, bool)`

GetQualityPolicyVersionIdOk returns a tuple with the QualityPolicyVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) SetQualityPolicyVersionId(v int32)`

SetQualityPolicyVersionId sets QualityPolicyVersionId field to given value.

### HasQualityPolicyVersionId

`func (o *GetRevisionScans200ResponseResultsInner) HasQualityPolicyVersionId() bool`

HasQualityPolicyVersionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetRevisionScans200ResponseResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetRevisionScans200ResponseResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetRevisionScans200ResponseResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetRevisionScans200ResponseResultsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetRevisionScans200ResponseResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetRevisionScans200ResponseResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetRevisionScans200ResponseResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetRevisionScans200ResponseResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


