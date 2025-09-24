# GetSnippetPaths200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]GetSnippetPaths200ResponsePathsInner**](GetSnippetPaths200ResponsePathsInner.md) | Array of file and directory paths where snippets were detected | 

## Methods

### NewGetSnippetPaths200Response

`func NewGetSnippetPaths200Response(paths []GetSnippetPaths200ResponsePathsInner, ) *GetSnippetPaths200Response`

NewGetSnippetPaths200Response instantiates a new GetSnippetPaths200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSnippetPaths200ResponseWithDefaults

`func NewGetSnippetPaths200ResponseWithDefaults() *GetSnippetPaths200Response`

NewGetSnippetPaths200ResponseWithDefaults instantiates a new GetSnippetPaths200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *GetSnippetPaths200Response) GetPaths() []GetSnippetPaths200ResponsePathsInner`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *GetSnippetPaths200Response) GetPathsOk() (*[]GetSnippetPaths200ResponsePathsInner, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *GetSnippetPaths200Response) SetPaths(v []GetSnippetPaths200ResponsePathsInner)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


