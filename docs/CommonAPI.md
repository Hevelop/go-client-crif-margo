# Hevelop\CrifMargo\CommonAPI

All URIs are relative to *https://api-uat.crif.com/margo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDomain**](CommonAPI.md#GetDomain) | **Get** /domains/{domainType} | Get values of the domainType indicated
[**GetDomains**](CommonAPI.md#GetDomains) | **Get** /domains | List all domains



## GetDomain

> DomainArrayMetadataType GetDomain(ctx, domainType).Page(page).Size(size).AcceptLanguage(acceptLanguage).Execute()

Get values of the domainType indicated



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
	domainType := "domainType_example" // string | the domain type a unique identifier for a `Domain`.
	page := int64(789) // int64 | Zero-based number of the page to obtain (optional) (default to 0)
	size := int64(789) // int64 | It manages the maximum number of elements inside a response. Maximum can't be more than 100. (optional) (default to 15)
	acceptLanguage := "acceptLanguage_example" // string | Set here the language you want in output (optional) (default to "en-GB")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.GetDomain(context.Background(), domainType).Page(page).Size(size).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.GetDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomain`: DomainArrayMetadataType
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.GetDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainType** | **string** | the domain type a unique identifier for a &#x60;Domain&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int64** | Zero-based number of the page to obtain | [default to 0]
 **size** | **int64** | It manages the maximum number of elements inside a response. Maximum can&#39;t be more than 100. | [default to 15]
 **acceptLanguage** | **string** | Set here the language you want in output | [default to &quot;en-GB&quot;]

### Return type

[**DomainArrayMetadataType**](DomainArrayMetadataType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomains

> []DomainsType GetDomains(ctx).AcceptLanguage(acceptLanguage).Execute()

List all domains



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
	acceptLanguage := "acceptLanguage_example" // string | Set here the language you want in output (optional) (default to "en-GB")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommonAPI.GetDomains(context.Background()).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommonAPI.GetDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomains`: []DomainsType
	fmt.Fprintf(os.Stdout, "Response from `CommonAPI.GetDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Set here the language you want in output | [default to &quot;en-GB&quot;]

### Return type

[**[]DomainsType**](DomainsType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

