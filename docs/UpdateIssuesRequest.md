# UpdateIssuesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** | This is a free-form field for users to provide additional context for the exception. This value appears in the vulnerabilities.analysis.detail field in CycloneDX SBOM reports  | [optional] 
**Reason** | Pointer to **string** | Provided reason for ignoring or resolving a security issue. &#39;Fixed&#39; and &#39;Under_investigation&#39; map to VEX statuses with the same names. All other values map to the VEX status &#39;Not Affected&#39;. This value appears in the vulnerabilities.analysis.justification field of CycloneDX SBOM reports.  | [optional] 
**Ticket** | Pointer to [**UpdateIssuesRequestOneOf1Ticket**](UpdateIssuesRequestOneOf1Ticket.md) |  | [optional] 
**PackageScope** | Pointer to **string** |  | [optional] 
**IgnoreScope** | Pointer to **string** |  | [optional] 
**ExpiresAfter** | Pointer to **string** | The UTC date after which the exception will expire. If not provided, the exception will never expire.  | [optional] 
**LicenseId** | Pointer to **string** | Set a license ID to create exceptions for a specific license ID.  | [optional] 

## Methods

### NewUpdateIssuesRequest

`func NewUpdateIssuesRequest() *UpdateIssuesRequest`

NewUpdateIssuesRequest instantiates a new UpdateIssuesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIssuesRequestWithDefaults

`func NewUpdateIssuesRequestWithDefaults() *UpdateIssuesRequest`

NewUpdateIssuesRequestWithDefaults instantiates a new UpdateIssuesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateIssuesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateIssuesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateIssuesRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateIssuesRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateIssuesRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateIssuesRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateIssuesRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateIssuesRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetReason

`func (o *UpdateIssuesRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateIssuesRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateIssuesRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateIssuesRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTicket

`func (o *UpdateIssuesRequest) GetTicket() UpdateIssuesRequestOneOf1Ticket`

GetTicket returns the Ticket field if non-nil, zero value otherwise.

### GetTicketOk

`func (o *UpdateIssuesRequest) GetTicketOk() (*UpdateIssuesRequestOneOf1Ticket, bool)`

GetTicketOk returns a tuple with the Ticket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicket

`func (o *UpdateIssuesRequest) SetTicket(v UpdateIssuesRequestOneOf1Ticket)`

SetTicket sets Ticket field to given value.

### HasTicket

`func (o *UpdateIssuesRequest) HasTicket() bool`

HasTicket returns a boolean if a field has been set.

### GetPackageScope

`func (o *UpdateIssuesRequest) GetPackageScope() string`

GetPackageScope returns the PackageScope field if non-nil, zero value otherwise.

### GetPackageScopeOk

`func (o *UpdateIssuesRequest) GetPackageScopeOk() (*string, bool)`

GetPackageScopeOk returns a tuple with the PackageScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageScope

`func (o *UpdateIssuesRequest) SetPackageScope(v string)`

SetPackageScope sets PackageScope field to given value.

### HasPackageScope

`func (o *UpdateIssuesRequest) HasPackageScope() bool`

HasPackageScope returns a boolean if a field has been set.

### GetIgnoreScope

`func (o *UpdateIssuesRequest) GetIgnoreScope() string`

GetIgnoreScope returns the IgnoreScope field if non-nil, zero value otherwise.

### GetIgnoreScopeOk

`func (o *UpdateIssuesRequest) GetIgnoreScopeOk() (*string, bool)`

GetIgnoreScopeOk returns a tuple with the IgnoreScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreScope

`func (o *UpdateIssuesRequest) SetIgnoreScope(v string)`

SetIgnoreScope sets IgnoreScope field to given value.

### HasIgnoreScope

`func (o *UpdateIssuesRequest) HasIgnoreScope() bool`

HasIgnoreScope returns a boolean if a field has been set.

### GetExpiresAfter

`func (o *UpdateIssuesRequest) GetExpiresAfter() string`

GetExpiresAfter returns the ExpiresAfter field if non-nil, zero value otherwise.

### GetExpiresAfterOk

`func (o *UpdateIssuesRequest) GetExpiresAfterOk() (*string, bool)`

GetExpiresAfterOk returns a tuple with the ExpiresAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAfter

`func (o *UpdateIssuesRequest) SetExpiresAfter(v string)`

SetExpiresAfter sets ExpiresAfter field to given value.

### HasExpiresAfter

`func (o *UpdateIssuesRequest) HasExpiresAfter() bool`

HasExpiresAfter returns a boolean if a field has been set.

### GetLicenseId

`func (o *UpdateIssuesRequest) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *UpdateIssuesRequest) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *UpdateIssuesRequest) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *UpdateIssuesRequest) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


