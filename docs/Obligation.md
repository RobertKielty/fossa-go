# Obligation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the obligation. | [optional] 
**License** | Pointer to **string** | The license associated with the obligation. | [optional] 
**Revisions** | Pointer to **map[string][]string** | A dictionary where the key is the parent locator and the value is an array of locators in which the obligation was found. | [optional] 

## Methods

### NewObligation

`func NewObligation() *Obligation`

NewObligation instantiates a new Obligation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObligationWithDefaults

`func NewObligationWithDefaults() *Obligation`

NewObligationWithDefaults instantiates a new Obligation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Obligation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Obligation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Obligation) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Obligation) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLicense

`func (o *Obligation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Obligation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Obligation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Obligation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRevisions

`func (o *Obligation) GetRevisions() map[string][]string`

GetRevisions returns the Revisions field if non-nil, zero value otherwise.

### GetRevisionsOk

`func (o *Obligation) GetRevisionsOk() (*map[string][]string, bool)`

GetRevisionsOk returns a tuple with the Revisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisions

`func (o *Obligation) SetRevisions(v map[string][]string)`

SetRevisions sets Revisions field to given value.

### HasRevisions

`func (o *Obligation) HasRevisions() bool`

HasRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


