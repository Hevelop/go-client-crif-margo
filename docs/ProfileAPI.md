# Hevelop\CrifMargo\ProfileAPI

All URIs are relative to *https://api-uat.crif.com/margo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOfficeUsers**](ProfileAPI.md#GetOfficeUsers) | **Get** /offices/{officeId}/users | Get users list of the office
[**GetOffices**](ProfileAPI.md#GetOffices) | **Get** /offices | List all offices



## GetOfficeUsers

> UserArrayMetadataType GetOfficeUsers(ctx, officeId).Page(page).Size(size).Execute()

Get users list of the office



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-crif-margo"
)

func main() {
	officeId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | a unique identifier for a `Office`.
	page := int64(789) // int64 | Zero-based number of the page to obtain (optional) (default to 0)
	size := int64(789) // int64 | It manages the maximum number of elements inside a response. Maximum can't be more than 100. (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.GetOfficeUsers(context.Background(), officeId).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetOfficeUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOfficeUsers`: UserArrayMetadataType
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetOfficeUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**officeId** | **string** | a unique identifier for a &#x60;Office&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOfficeUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Zero-based number of the page to obtain | [default to 0]
 **size** | **int64** | It manages the maximum number of elements inside a response. Maximum can&#39;t be more than 100. | [default to 15]

### Return type

[**UserArrayMetadataType**](UserArrayMetadataType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffices

> OfficeArrayMetadataType GetOffices(ctx).Page(page).Size(size).Execute()

List all offices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-crif-margo"
)

func main() {
	page := int64(789) // int64 | Zero-based number of the page to obtain (optional) (default to 0)
	size := int64(789) // int64 | It manages the maximum number of elements inside a response. Maximum can't be more than 100. (optional) (default to 15)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProfileAPI.GetOffices(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.GetOffices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOffices`: OfficeArrayMetadataType
	fmt.Fprintf(os.Stdout, "Response from `ProfileAPI.GetOffices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOfficesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Zero-based number of the page to obtain | [default to 0]
 **size** | **int64** | It manages the maximum number of elements inside a response. Maximum can&#39;t be more than 100. | [default to 15]

### Return type

[**OfficeArrayMetadataType**](OfficeArrayMetadataType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

