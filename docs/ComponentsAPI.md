# \ComponentsAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Build**](ComponentsAPI.md#Build) | **Post** /components/build | 
[**GetSignedUrl**](ComponentsAPI.md#GetSignedUrl) | **Get** /components/signed_url | 



## Build

> Build(ctx).PackageSpec(packageSpec).Revision(revision).Dependency(dependency).Description(description).FileType(fileType).Branch(branch).JiraProjectKey(jiraProjectKey).Link(link).ProjectURL(projectURL).Policy(policy).PolicyId(policyId).Team(team).Title(title).ReleaseGroup(releaseGroup).ReleaseGroupRelease(releaseGroupRelease).Labels(labels).BuildRequest(buildRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	packageSpec := "packageSpec_example" // string | The unresolved spec (i.e. a URL) that can be passed to a fetcher and resolved to a package id. For example:, underscore (npm), https://github.com/fossas/fossa (git), nokogiri (gem), 
	revision := "revision_example" // string | The branch or revision of the component project being built.
	dependency := true // bool |  (optional)
	description := "description_example" // string | The project description for the archive (optional)
	fileType := "fileType_example" // string | The kind of component file to be built., If 'archive', the signed URL is for uplaoding a directory of source code., If 'sbom', the signed URL is for uploading an SBOM file (CycloneDX or SPDX).,  (optional)
	branch := "branch_example" // string |  (optional)
	jiraProjectKey := "jiraProjectKey_example" // string | The corresponding Jira project for this component (optional)
	link := "link_example" // string | A link to attach to this revision (optional)
	projectURL := "projectURL_example" // string | The URL of the project being uploaded (optional)
	policy := "policy_example" // string | The name of the policy for this build. (optional)
	policyId := float32(8.14) // float32 | The ID of the policy for this build. (optional)
	team := "team_example" // string | The name of the team connected to the project. (optional)
	title := "title_example" // string | The title of the component. (optional)
	releaseGroup := "releaseGroup_example" // string | The title of release group to add the project to. (optional)
	releaseGroupRelease := "releaseGroupRelease_example" // string | The title of release to add the project to. (optional)
	labels := []string{"Inner_example"} // []string | Set of labels to apply to a project. If the org or the project have too many labels, then the labels will be applied in order until limits are hit.  (optional)
	buildRequest := *openapiclient.NewBuildRequest() // BuildRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.Build(context.Background()).PackageSpec(packageSpec).Revision(revision).Dependency(dependency).Description(description).FileType(fileType).Branch(branch).JiraProjectKey(jiraProjectKey).Link(link).ProjectURL(projectURL).Policy(policy).PolicyId(policyId).Team(team).Title(title).ReleaseGroup(releaseGroup).ReleaseGroupRelease(releaseGroupRelease).Labels(labels).BuildRequest(buildRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.Build``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuildRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageSpec** | **string** | The unresolved spec (i.e. a URL) that can be passed to a fetcher and resolved to a package id. For example:, underscore (npm), https://github.com/fossas/fossa (git), nokogiri (gem),  | 
 **revision** | **string** | The branch or revision of the component project being built. | 
 **dependency** | **bool** |  | 
 **description** | **string** | The project description for the archive | 
 **fileType** | **string** | The kind of component file to be built., If &#39;archive&#39;, the signed URL is for uplaoding a directory of source code., If &#39;sbom&#39;, the signed URL is for uploading an SBOM file (CycloneDX or SPDX).,  | 
 **branch** | **string** |  | 
 **jiraProjectKey** | **string** | The corresponding Jira project for this component | 
 **link** | **string** | A link to attach to this revision | 
 **projectURL** | **string** | The URL of the project being uploaded | 
 **policy** | **string** | The name of the policy for this build. | 
 **policyId** | **float32** | The ID of the policy for this build. | 
 **team** | **string** | The name of the team connected to the project. | 
 **title** | **string** | The title of the component. | 
 **releaseGroup** | **string** | The title of release group to add the project to. | 
 **releaseGroupRelease** | **string** | The title of release to add the project to. | 
 **labels** | **[]string** | Set of labels to apply to a project. If the org or the project have too many labels, then the labels will be applied in order until limits are hit.  | 
 **buildRequest** | [**BuildRequest**](BuildRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignedUrl

> GetSignedUrl200Response GetSignedUrl(ctx).PackageSpec(packageSpec).Revision(revision).FileType(fileType).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/RobertKielty/fossa-go"
)

func main() {
	packageSpec := "packageSpec_example" // string | the package spec
	revision := "revision_example" // string | The branch or revision of the component project being built.
	fileType := "fileType_example" // string | The kind of file to be uploaded to the signed URL. If 'archive', the signed URL is for uplaoding a directory of source code. If 'sbom', the signed URL is for uploading an SBOM file (CycloneDX or SPDX).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetSignedUrl(context.Background()).PackageSpec(packageSpec).Revision(revision).FileType(fileType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetSignedUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSignedUrl`: GetSignedUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetSignedUrl`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignedUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageSpec** | **string** | the package spec | 
 **revision** | **string** | The branch or revision of the component project being built. | 
 **fileType** | **string** | The kind of file to be uploaded to the signed URL. If &#39;archive&#39;, the signed URL is for uplaoding a directory of source code. If &#39;sbom&#39;, the signed URL is for uploading an SBOM file (CycloneDX or SPDX).  | 

### Return type

[**GetSignedUrl200Response**](GetSignedUrl200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

