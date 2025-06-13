# GetProjectExportIssues200ResponseQualityIssuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependency** | Pointer to **string** | The dependency that the issue is associated with.  | [optional] 
**Revision** | Pointer to **string** | The revision of the dependency that the issue is associated with.  | [optional] 
**Type** | Pointer to **string** | The type of issue.  | [optional] 
**FOSSA_URL** | Pointer to **string** | A link to the project in the FOSSA UI.  | [optional] 
**AffectedProject** | Pointer to **string** | The User Project that the issue is associated with.  | [optional] 
**Details** | Pointer to **string** | A description of the policy rule that triggered this issue (if available).  | [optional] 
**DependencyDepths** | Pointer to **string** | The depth(s) of the dependency in the dependency tree.  | [optional] 

## Methods

### NewGetProjectExportIssues200ResponseQualityIssuesInner

`func NewGetProjectExportIssues200ResponseQualityIssuesInner() *GetProjectExportIssues200ResponseQualityIssuesInner`

NewGetProjectExportIssues200ResponseQualityIssuesInner instantiates a new GetProjectExportIssues200ResponseQualityIssuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectExportIssues200ResponseQualityIssuesInnerWithDefaults

`func NewGetProjectExportIssues200ResponseQualityIssuesInnerWithDefaults() *GetProjectExportIssues200ResponseQualityIssuesInner`

NewGetProjectExportIssues200ResponseQualityIssuesInnerWithDefaults instantiates a new GetProjectExportIssues200ResponseQualityIssuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependency

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetRevision

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetType

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFOSSA_URL

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetFOSSA_URL() string`

GetFOSSA_URL returns the FOSSA_URL field if non-nil, zero value otherwise.

### GetFOSSA_URLOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetFOSSA_URLOk() (*string, bool)`

GetFOSSA_URLOk returns a tuple with the FOSSA_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFOSSA_URL

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetFOSSA_URL(v string)`

SetFOSSA_URL sets FOSSA_URL field to given value.

### HasFOSSA_URL

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasFOSSA_URL() bool`

HasFOSSA_URL returns a boolean if a field has been set.

### GetAffectedProject

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetAffectedProject() string`

GetAffectedProject returns the AffectedProject field if non-nil, zero value otherwise.

### GetAffectedProjectOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetAffectedProjectOk() (*string, bool)`

GetAffectedProjectOk returns a tuple with the AffectedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProject

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetAffectedProject(v string)`

SetAffectedProject sets AffectedProject field to given value.

### HasAffectedProject

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasAffectedProject() bool`

HasAffectedProject returns a boolean if a field has been set.

### GetDetails

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDependencyDepths

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDependencyDepths() string`

GetDependencyDepths returns the DependencyDepths field if non-nil, zero value otherwise.

### GetDependencyDepthsOk

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) GetDependencyDepthsOk() (*string, bool)`

GetDependencyDepthsOk returns a tuple with the DependencyDepths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDepths

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) SetDependencyDepths(v string)`

SetDependencyDepths sets DependencyDepths field to given value.

### HasDependencyDepths

`func (o *GetProjectExportIssues200ResponseQualityIssuesInner) HasDependencyDepths() bool`

HasDependencyDepths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


