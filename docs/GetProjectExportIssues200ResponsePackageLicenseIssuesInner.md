# GetProjectExportIssues200ResponsePackageLicenseIssuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dependency** | Pointer to **string** | The dependency that the issue is associated with.  | [optional] 
**Revision** | Pointer to **string** | The revision of the dependency that the issue is associated with.  | [optional] 
**FlaggedByPolicy** | Pointer to **string** | The policy that triggered the issue.  | [optional] 
**IssueType** | Pointer to **string** | The type of issue.  | [optional] 
**License** | Pointer to **string** | The license that triggered the issue.  | [optional] 
**AffectedProject** | Pointer to **string** | The User Project that the issue is associated with.  | [optional] 
**FOSSA_URL** | Pointer to **string** | A link to the project in the FOSSA UI.  | [optional] 
**Details** | Pointer to **string** | A description of the policy rule that triggered this issue (if available).  | [optional] 
**DependencyDepths** | Pointer to **string** | The depth(s) of the dependency in the dependency tree.  | [optional] 

## Methods

### NewGetProjectExportIssues200ResponsePackageLicenseIssuesInner

`func NewGetProjectExportIssues200ResponsePackageLicenseIssuesInner() *GetProjectExportIssues200ResponsePackageLicenseIssuesInner`

NewGetProjectExportIssues200ResponsePackageLicenseIssuesInner instantiates a new GetProjectExportIssues200ResponsePackageLicenseIssuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectExportIssues200ResponsePackageLicenseIssuesInnerWithDefaults

`func NewGetProjectExportIssues200ResponsePackageLicenseIssuesInnerWithDefaults() *GetProjectExportIssues200ResponsePackageLicenseIssuesInner`

NewGetProjectExportIssues200ResponsePackageLicenseIssuesInnerWithDefaults instantiates a new GetProjectExportIssues200ResponsePackageLicenseIssuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDependency

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDependency() string`

GetDependency returns the Dependency field if non-nil, zero value otherwise.

### GetDependencyOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDependencyOk() (*string, bool)`

GetDependencyOk returns a tuple with the Dependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependency

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetDependency(v string)`

SetDependency sets Dependency field to given value.

### HasDependency

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasDependency() bool`

HasDependency returns a boolean if a field has been set.

### GetRevision

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetFlaggedByPolicy

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetFlaggedByPolicy() string`

GetFlaggedByPolicy returns the FlaggedByPolicy field if non-nil, zero value otherwise.

### GetFlaggedByPolicyOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetFlaggedByPolicyOk() (*string, bool)`

GetFlaggedByPolicyOk returns a tuple with the FlaggedByPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedByPolicy

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetFlaggedByPolicy(v string)`

SetFlaggedByPolicy sets FlaggedByPolicy field to given value.

### HasFlaggedByPolicy

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasFlaggedByPolicy() bool`

HasFlaggedByPolicy returns a boolean if a field has been set.

### GetIssueType

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetLicense

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetAffectedProject

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetAffectedProject() string`

GetAffectedProject returns the AffectedProject field if non-nil, zero value otherwise.

### GetAffectedProjectOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetAffectedProjectOk() (*string, bool)`

GetAffectedProjectOk returns a tuple with the AffectedProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProject

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetAffectedProject(v string)`

SetAffectedProject sets AffectedProject field to given value.

### HasAffectedProject

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasAffectedProject() bool`

HasAffectedProject returns a boolean if a field has been set.

### GetFOSSA_URL

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetFOSSA_URL() string`

GetFOSSA_URL returns the FOSSA_URL field if non-nil, zero value otherwise.

### GetFOSSA_URLOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetFOSSA_URLOk() (*string, bool)`

GetFOSSA_URLOk returns a tuple with the FOSSA_URL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFOSSA_URL

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetFOSSA_URL(v string)`

SetFOSSA_URL sets FOSSA_URL field to given value.

### HasFOSSA_URL

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasFOSSA_URL() bool`

HasFOSSA_URL returns a boolean if a field has been set.

### GetDetails

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDependencyDepths

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDependencyDepths() string`

GetDependencyDepths returns the DependencyDepths field if non-nil, zero value otherwise.

### GetDependencyDepthsOk

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) GetDependencyDepthsOk() (*string, bool)`

GetDependencyDepthsOk returns a tuple with the DependencyDepths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyDepths

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) SetDependencyDepths(v string)`

SetDependencyDepths sets DependencyDepths field to given value.

### HasDependencyDepths

`func (o *GetProjectExportIssues200ResponsePackageLicenseIssuesInner) HasDependencyDepths() bool`

HasDependencyDepths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


