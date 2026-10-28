# LicenseDisputeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedBy** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**LicenseId** | **string** |  | 
**DisputedIssueId** | **int32** |  | 
**Reason** | **string** |  | 
**Comment** | **string** |  | 
**ResolvedAt** | **time.Time** |  | 
**ResolvedBy** | **int32** |  | 

## Methods

### NewLicenseDisputeResponse

`func NewLicenseDisputeResponse(id int32, createdBy int32, createdAt time.Time, licenseId string, disputedIssueId int32, reason string, comment string, resolvedAt time.Time, resolvedBy int32, ) *LicenseDisputeResponse`

NewLicenseDisputeResponse instantiates a new LicenseDisputeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseDisputeResponseWithDefaults

`func NewLicenseDisputeResponseWithDefaults() *LicenseDisputeResponse`

NewLicenseDisputeResponseWithDefaults instantiates a new LicenseDisputeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LicenseDisputeResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LicenseDisputeResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LicenseDisputeResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *LicenseDisputeResponse) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LicenseDisputeResponse) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LicenseDisputeResponse) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *LicenseDisputeResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LicenseDisputeResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LicenseDisputeResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLicenseId

`func (o *LicenseDisputeResponse) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *LicenseDisputeResponse) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *LicenseDisputeResponse) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.


### GetDisputedIssueId

`func (o *LicenseDisputeResponse) GetDisputedIssueId() int32`

GetDisputedIssueId returns the DisputedIssueId field if non-nil, zero value otherwise.

### GetDisputedIssueIdOk

`func (o *LicenseDisputeResponse) GetDisputedIssueIdOk() (*int32, bool)`

GetDisputedIssueIdOk returns a tuple with the DisputedIssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputedIssueId

`func (o *LicenseDisputeResponse) SetDisputedIssueId(v int32)`

SetDisputedIssueId sets DisputedIssueId field to given value.


### GetReason

`func (o *LicenseDisputeResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LicenseDisputeResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LicenseDisputeResponse) SetReason(v string)`

SetReason sets Reason field to given value.


### GetComment

`func (o *LicenseDisputeResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *LicenseDisputeResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *LicenseDisputeResponse) SetComment(v string)`

SetComment sets Comment field to given value.


### GetResolvedAt

`func (o *LicenseDisputeResponse) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *LicenseDisputeResponse) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *LicenseDisputeResponse) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.


### GetResolvedBy

`func (o *LicenseDisputeResponse) GetResolvedBy() int32`

GetResolvedBy returns the ResolvedBy field if non-nil, zero value otherwise.

### GetResolvedByOk

`func (o *LicenseDisputeResponse) GetResolvedByOk() (*int32, bool)`

GetResolvedByOk returns a tuple with the ResolvedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedBy

`func (o *LicenseDisputeResponse) SetResolvedBy(v int32)`

SetResolvedBy sets ResolvedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


