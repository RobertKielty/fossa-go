# \PackageObservabilityAPI

All URIs are relative to *https://app.fossa.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPackageIndexExport**](PackageObservabilityAPI.md#GetPackageIndexExport) | **Get** /packages/report | 
[**GetPackages**](PackageObservabilityAPI.md#GetPackages) | **Get** /packages | 
[**GetPackagesPackageLocators**](PackageObservabilityAPI.md#GetPackagesPackageLocators) | **Get** /packages/package-locators | 
[**GetPackagesPackageManagers**](PackageObservabilityAPI.md#GetPackagesPackageManagers) | **Get** /packages/package-managers | 
[**GetPackagesPackageSummary**](PackageObservabilityAPI.md#GetPackagesPackageSummary) | **Get** /packages/package-summary | 



## GetPackageIndexExport

> GetPackageIndexExport201Response GetPackageIndexExport(ctx).Fetchers(fetchers).PackageName(packageName).Depth(depth).Labels(labels).ProjectName(projectName).Sources(sources).Visibility(visibility).BlockType(blockType).Cve(cve).Cwes(cwes).Locators(locators).FixTypes(fixTypes).Severities(severities).TeamIds(teamIds).Execute()





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
	fetchers := []string{"Fetchers_example"} // []string | Filter packages to those from the specified fetchers.  For example, fetchers[0]=npm&fetchers[1]=apk will match all NPM packages and all APK packages. (optional)
	packageName := "packageName_example" // string | Filter results to only packages with the specified name.  Supports partial matches.  For example \"foo\" will match \"foo\", \"foobar\", and \"foo-bar\". (optional)
	depth := []string{"Depth_example"} // []string | Filter results to only include packages which are direct or transitive dependencies of your projects. (optional)
	labels := []string{"Inner_example"} // []string | Filter packages to those belonging to your projects with the specified labels. (optional)
	projectName := "projectName_example" // string | Filter packages to only one of your specific projects.  Exact match only. (optional)
	sources := []string{"Inner_example"} // []string | Filter packages to those belonging to your projects from the specified set of sources (optional)
	visibility := []string{"Visibility_example"} // []string | Filter results to your projects which are public or private (optional)
	blockType := "blockType_example" // string | Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization (optional)
	cve := "cve_example" // string | Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers (optional)
	cwes := []string{"Inner_example"} // []string | Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers (optional)
	locators := []string{"Inner_example"} // []string | Filter packages to those with specific package locators (exact match only) (optional)
	fixTypes := []string{"FixTypes_example"} // []string | Filter packages to those with vulnerabilities that either have or do not have a fix available (optional)
	severities := []string{"Severities_example"} // []string | Filter packages by severity levels of issues (optional)
	teamIds := []float32{float32(123)} // []float32 | Filter packages to just those owned by the specified teams.  Specify the string \"null\" to filter packages that are not owned by any team. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageObservabilityAPI.GetPackageIndexExport(context.Background()).Fetchers(fetchers).PackageName(packageName).Depth(depth).Labels(labels).ProjectName(projectName).Sources(sources).Visibility(visibility).BlockType(blockType).Cve(cve).Cwes(cwes).Locators(locators).FixTypes(fixTypes).Severities(severities).TeamIds(teamIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageObservabilityAPI.GetPackageIndexExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackageIndexExport`: GetPackageIndexExport201Response
	fmt.Fprintf(os.Stdout, "Response from `PackageObservabilityAPI.GetPackageIndexExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackageIndexExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchers** | **[]string** | Filter packages to those from the specified fetchers.  For example, fetchers[0]&#x3D;npm&amp;fetchers[1]&#x3D;apk will match all NPM packages and all APK packages. | 
 **packageName** | **string** | Filter results to only packages with the specified name.  Supports partial matches.  For example \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;foo-bar\&quot;. | 
 **depth** | **[]string** | Filter results to only include packages which are direct or transitive dependencies of your projects. | 
 **labels** | **[]string** | Filter packages to those belonging to your projects with the specified labels. | 
 **projectName** | **string** | Filter packages to only one of your specific projects.  Exact match only. | 
 **sources** | **[]string** | Filter packages to those belonging to your projects from the specified set of sources | 
 **visibility** | **[]string** | Filter results to your projects which are public or private | 
 **blockType** | **string** | Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization | 
 **cve** | **string** | Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers | 
 **cwes** | **[]string** | Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers | 
 **locators** | **[]string** | Filter packages to those with specific package locators (exact match only) | 
 **fixTypes** | **[]string** | Filter packages to those with vulnerabilities that either have or do not have a fix available | 
 **severities** | **[]string** | Filter packages by severity levels of issues | 
 **teamIds** | **[]float32** | Filter packages to just those owned by the specified teams.  Specify the string \&quot;null\&quot; to filter packages that are not owned by any team. | 

### Return type

[**GetPackageIndexExport201Response**](GetPackageIndexExport201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackages

> GetPackageIndexExport201Response GetPackages(ctx).Fetchers(fetchers).PackageName(packageName).Depth(depth).Labels(labels).ProjectName(projectName).Sources(sources).Visibility(visibility).BlockType(blockType).Cve(cve).Cwes(cwes).FixTypes(fixTypes).Severities(severities).TeamIds(teamIds).Locators(locators).Page(page).Count(count).Sort(sort).Execute()





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
	fetchers := []string{"Fetchers_example"} // []string | Filter packages to those from the specified fetchers.  For example, fetchers[0]=npm&fetchers[1]=apk will match all NPM packages and all APK packages. (optional)
	packageName := "packageName_example" // string | Filter results to only packages with the specified name.  Supports partial matches.  For example \"foo\" will match \"foo\", \"foobar\", and \"foo-bar\". (optional)
	depth := []string{"Depth_example"} // []string | Filter results to only include packages which are direct or transitive dependencies of your projects. (optional)
	labels := []string{"Inner_example"} // []string | Filter packages to those belonging to your projects with the specified labels. (optional)
	projectName := "projectName_example" // string | Filter packages to only one of your specific projects.  Exact match only. (optional)
	sources := []string{"Inner_example"} // []string | Filter packages to those belonging to your projects from the specified set of sources (optional)
	visibility := []string{"Visibility_example"} // []string | Filter results to your projects which are public or private (optional)
	blockType := "blockType_example" // string | Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization (optional)
	cve := "cve_example" // string | Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers (optional)
	cwes := []string{"Inner_example"} // []string | Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers (optional)
	fixTypes := []string{"FixTypes_example"} // []string | Filter packages to those with vulnerabilities that either have or do not have a fix available (optional)
	severities := []string{"Severities_example"} // []string | Filter packages by severity levels of issues (optional)
	teamIds := []float32{float32(123)} // []float32 | Filter packages to just those owned by the specified teams.  Specify the string \"null\" to filter packages that are not owned by any team. (optional)
	locators := []string{"Inner_example"} // []string | Filter packages by the specified dependency project locators without revision / version data.  Locators are unique identifiers for packages in the FOSSA system. Exact matches only. (optional)
	page := int32(56) // int32 | The page number to retrieve (optional) (default to 1)
	count := int32(56) // int32 | The number of results to return per page (optional)
	sort := "sort_example" // string | The field to sort by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageObservabilityAPI.GetPackages(context.Background()).Fetchers(fetchers).PackageName(packageName).Depth(depth).Labels(labels).ProjectName(projectName).Sources(sources).Visibility(visibility).BlockType(blockType).Cve(cve).Cwes(cwes).FixTypes(fixTypes).Severities(severities).TeamIds(teamIds).Locators(locators).Page(page).Count(count).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageObservabilityAPI.GetPackages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackages`: GetPackageIndexExport201Response
	fmt.Fprintf(os.Stdout, "Response from `PackageObservabilityAPI.GetPackages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchers** | **[]string** | Filter packages to those from the specified fetchers.  For example, fetchers[0]&#x3D;npm&amp;fetchers[1]&#x3D;apk will match all NPM packages and all APK packages. | 
 **packageName** | **string** | Filter results to only packages with the specified name.  Supports partial matches.  For example \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;foo-bar\&quot;. | 
 **depth** | **[]string** | Filter results to only include packages which are direct or transitive dependencies of your projects. | 
 **labels** | **[]string** | Filter packages to those belonging to your projects with the specified labels. | 
 **projectName** | **string** | Filter packages to only one of your specific projects.  Exact match only. | 
 **sources** | **[]string** | Filter packages to those belonging to your projects from the specified set of sources | 
 **visibility** | **[]string** | Filter results to your projects which are public or private | 
 **blockType** | **string** | Filter packages to include only packages that do or do not have packages as dependencies which are blocked by your organization | 
 **cve** | **string** | Filter packages to those with vulnerabilities that have specific Common Vulnerabilities and Exposures (CVE) identifiers | 
 **cwes** | **[]string** | Filter packages to those with vulnerabilities that have specific Common Weakness Enumeration (CWE) identifiers | 
 **fixTypes** | **[]string** | Filter packages to those with vulnerabilities that either have or do not have a fix available | 
 **severities** | **[]string** | Filter packages by severity levels of issues | 
 **teamIds** | **[]float32** | Filter packages to just those owned by the specified teams.  Specify the string \&quot;null\&quot; to filter packages that are not owned by any team. | 
 **locators** | **[]string** | Filter packages by the specified dependency project locators without revision / version data.  Locators are unique identifiers for packages in the FOSSA system. Exact matches only. | 
 **page** | **int32** | The page number to retrieve | [default to 1]
 **count** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 

### Return type

[**GetPackageIndexExport201Response**](GetPackageIndexExport201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesPackageLocators

> GetPackagesPackageLocators200Response GetPackagesPackageLocators(ctx).PackageLocator(packageLocator).Count(count).Execute()





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
	packageLocator := "packageLocator_example" // string | Filter locators to those partially matching the specified locator.  For example, \"foo\" will match \"foo\", \"foobar\", and \"bazfoobar\". (optional)
	count := int32(56) // int32 | The number of results to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageObservabilityAPI.GetPackagesPackageLocators(context.Background()).PackageLocator(packageLocator).Count(count).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageObservabilityAPI.GetPackagesPackageLocators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesPackageLocators`: GetPackagesPackageLocators200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageObservabilityAPI.GetPackagesPackageLocators`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesPackageLocatorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageLocator** | **string** | Filter locators to those partially matching the specified locator.  For example, \&quot;foo\&quot; will match \&quot;foo\&quot;, \&quot;foobar\&quot;, and \&quot;bazfoobar\&quot;. | 
 **count** | **int32** | The number of results to return. | 

### Return type

[**GetPackagesPackageLocators200Response**](GetPackagesPackageLocators200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesPackageManagers

> []string GetPackagesPackageManagers(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageObservabilityAPI.GetPackagesPackageManagers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageObservabilityAPI.GetPackagesPackageManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesPackageManagers`: []string
	fmt.Fprintf(os.Stdout, "Response from `PackageObservabilityAPI.GetPackagesPackageManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesPackageManagersRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPackagesPackageSummary

> GetPackagesPackageSummary200Response GetPackagesPackageSummary(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PackageObservabilityAPI.GetPackagesPackageSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PackageObservabilityAPI.GetPackagesPackageSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPackagesPackageSummary`: GetPackagesPackageSummary200Response
	fmt.Fprintf(os.Stdout, "Response from `PackageObservabilityAPI.GetPackagesPackageSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPackagesPackageSummaryRequest struct via the builder pattern


### Return type

[**GetPackagesPackageSummary200Response**](GetPackagesPackageSummary200Response.md)

### Authorization

[ApiToken](../README.md#ApiToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

