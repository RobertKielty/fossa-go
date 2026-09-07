# ListFossabotDependencyUpgradePRs200ResponsePageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartCursor** | **NullableString** |  | 
**EndCursor** | **NullableString** |  | 
**HasNextPage** | **bool** | For forward requests, whether another forward slice exists. Always false for backward requests.  | 
**HasPreviousPage** | **bool** | For backward requests, whether another backward slice exists. Always false for forward requests.  | 

## Methods

### NewListFossabotDependencyUpgradePRs200ResponsePageInfo

`func NewListFossabotDependencyUpgradePRs200ResponsePageInfo(startCursor NullableString, endCursor NullableString, hasNextPage bool, hasPreviousPage bool, ) *ListFossabotDependencyUpgradePRs200ResponsePageInfo`

NewListFossabotDependencyUpgradePRs200ResponsePageInfo instantiates a new ListFossabotDependencyUpgradePRs200ResponsePageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFossabotDependencyUpgradePRs200ResponsePageInfoWithDefaults

`func NewListFossabotDependencyUpgradePRs200ResponsePageInfoWithDefaults() *ListFossabotDependencyUpgradePRs200ResponsePageInfo`

NewListFossabotDependencyUpgradePRs200ResponsePageInfoWithDefaults instantiates a new ListFossabotDependencyUpgradePRs200ResponsePageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartCursor

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetStartCursor() string`

GetStartCursor returns the StartCursor field if non-nil, zero value otherwise.

### GetStartCursorOk

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetStartCursorOk() (*string, bool)`

GetStartCursorOk returns a tuple with the StartCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCursor

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetStartCursor(v string)`

SetStartCursor sets StartCursor field to given value.


### SetStartCursorNil

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetStartCursorNil(b bool)`

 SetStartCursorNil sets the value for StartCursor to be an explicit nil

### UnsetStartCursor
`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) UnsetStartCursor()`

UnsetStartCursor ensures that no value is present for StartCursor, not even an explicit nil
### GetEndCursor

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetEndCursor() string`

GetEndCursor returns the EndCursor field if non-nil, zero value otherwise.

### GetEndCursorOk

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetEndCursorOk() (*string, bool)`

GetEndCursorOk returns a tuple with the EndCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCursor

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetEndCursor(v string)`

SetEndCursor sets EndCursor field to given value.


### SetEndCursorNil

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetEndCursorNil(b bool)`

 SetEndCursorNil sets the value for EndCursor to be an explicit nil

### UnsetEndCursor
`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) UnsetEndCursor()`

UnsetEndCursor ensures that no value is present for EndCursor, not even an explicit nil
### GetHasNextPage

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.


### GetHasPreviousPage

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetHasPreviousPage() bool`

GetHasPreviousPage returns the HasPreviousPage field if non-nil, zero value otherwise.

### GetHasPreviousPageOk

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) GetHasPreviousPageOk() (*bool, bool)`

GetHasPreviousPageOk returns a tuple with the HasPreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPreviousPage

`func (o *ListFossabotDependencyUpgradePRs200ResponsePageInfo) SetHasPreviousPage(v bool)`

SetHasPreviousPage sets HasPreviousPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


