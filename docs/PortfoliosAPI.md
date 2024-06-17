# Hevelop\CrifMargo\PortfoliosAPI

All URIs are relative to *https://api-uat.crif.com/margo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePortfolio**](PortfoliosAPI.md#DeletePortfolio) | **Delete** /portfolios/{portfolioId} | Delete a portfolio
[**DeleteSharingPortfolio**](PortfoliosAPI.md#DeleteSharingPortfolio) | **Delete** /portfolios/{portfolioId}/share | Delete Sharing a Portfolio
[**GetPortfolio**](PortfoliosAPI.md#GetPortfolio) | **Get** /portfolios/{portfolioId} | Get a portfolio
[**GetPortfolios**](PortfoliosAPI.md#GetPortfolios) | **Get** /portfolios | List all portfolios
[**PortfolioSharesDelete**](PortfoliosAPI.md#PortfolioSharesDelete) | **Post** /portfolios/{portfolioId}/shares/delete | Delete Portfolio shares
[**PortfolioSharingList**](PortfoliosAPI.md#PortfolioSharingList) | **Get** /portfolios/{portfolioId}/share | Share a Portfolio
[**PostAddPortfolioCompanies**](PortfoliosAPI.md#PostAddPortfolioCompanies) | **Post** /portfolios/{portfolioId}/companies/bulk-add | Add companies to portfolio
[**PostDeletePortfolioCompanies**](PortfoliosAPI.md#PostDeletePortfolioCompanies) | **Post** /portfolios/{portfolioId}/companies/bulk-delete | Delete companies from portfolio
[**PostPortfolioDownload**](PortfoliosAPI.md#PostPortfolioDownload) | **Post** /portfolios/{portfolioId}/download | Download portfolio
[**PostPortfolios**](PortfoliosAPI.md#PostPortfolios) | **Post** /portfolios | Create portfolio
[**PostSimilarity**](PortfoliosAPI.md#PostSimilarity) | **Post** /portfolios/{portfolioId}/create-similarity | Generate prospect based on own portfolio
[**PostUpdatePortfolioCompanies**](PortfoliosAPI.md#PostUpdatePortfolioCompanies) | **Post** /portfolios/{portfolioId}/companies/bulk-update | Update companies to portfolio
[**SharePortfolio**](PortfoliosAPI.md#SharePortfolio) | **Post** /portfolios/{portfolioId}/share | Share a Portfolio



## DeletePortfolio

> DeletePortfolio(ctx, portfolioId).Execute()

Delete a portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortfoliosAPI.DeletePortfolio(context.Background(), portfolioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.DeletePortfolio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharingPortfolio

> DeleteSharingPortfolio(ctx, portfolioId).Execute()

Delete Sharing a Portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortfoliosAPI.DeleteSharingPortfolio(context.Background(), portfolioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.DeleteSharingPortfolio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharingPortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolio

> PortfolioType GetPortfolio(ctx, portfolioId).Execute()

Get a portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.GetPortfolio(context.Background(), portfolioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.GetPortfolio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolio`: PortfolioType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.GetPortfolio`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortfolioType**](PortfolioType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPortfolios

> PortfolioSummaryArrayMetadataType GetPortfolios(ctx).Page(page).Size(size).Execute()

List all portfolios



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
	resp, r, err := apiClient.PortfoliosAPI.GetPortfolios(context.Background()).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.GetPortfolios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPortfolios`: PortfolioSummaryArrayMetadataType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.GetPortfolios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPortfoliosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int64** | Zero-based number of the page to obtain | [default to 0]
 **size** | **int64** | It manages the maximum number of elements inside a response. Maximum can&#39;t be more than 100. | [default to 15]

### Return type

[**PortfolioSummaryArrayMetadataType**](PortfolioSummaryArrayMetadataType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioSharesDelete

> PortfolioSharesDelete(ctx, portfolioId).OfficesList(officesList).Execute()

Delete Portfolio shares



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	officesList := *openapiclient.NewOfficesList([]string{"d069e399-894b-46be-857b-fbccd497d9b3"}) // OfficesList | Office ids for which is asking to delete the portfolio share (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortfoliosAPI.PortfolioSharesDelete(context.Background(), portfolioId).OfficesList(officesList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PortfolioSharesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioSharesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **officesList** | [**OfficesList**](OfficesList.md) | Office ids for which is asking to delete the portfolio share | 

### Return type

 (empty response body)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioSharingList

> []AccessRight PortfolioSharingList(ctx, portfolioId).Execute()

Share a Portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PortfolioSharingList(context.Background(), portfolioId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PortfolioSharingList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioSharingList`: []AccessRight
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PortfolioSharingList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioSharingListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AccessRight**](AccessRight.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddPortfolioCompanies

> []CompanyResultType PostAddPortfolioCompanies(ctx, portfolioId).CompanyArrayType(companyArrayType).Execute()

Add companies to portfolio

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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	companyArrayType := *openapiclient.NewCompanyArrayType() // CompanyArrayType | Add one or more companies to an existing portfolio.  `companies` list needs to have maximum 100 objects per call and each object belonging to a list must contain the company's information to upload in the portfolio.  The service is able to upload the headquarter or the branch data. The data upload depends by the values reported in each object: * if the object contains only the **identificationCode** the system upload only the headquarter data. * if the object contains the **identificationCode** and the **companyUnitId** the system upload only the branch data.  Values reported in the `value` items (customVariable array) shall respect the data type defined in the **POST /portfolios**.  `margoId` and `isEnriched` are read-only fields and must not be provided in input. `margoId` is generated by Margo system during the company upload in order to identify a company inside portfolio. `isEnriched` indicate if a company is enriched or not.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostAddPortfolioCompanies(context.Background(), portfolioId).CompanyArrayType(companyArrayType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostAddPortfolioCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAddPortfolioCompanies`: []CompanyResultType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostAddPortfolioCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddPortfolioCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyArrayType** | [**CompanyArrayType**](CompanyArrayType.md) | Add one or more companies to an existing portfolio.  &#x60;companies&#x60; list needs to have maximum 100 objects per call and each object belonging to a list must contain the company&#39;s information to upload in the portfolio.  The service is able to upload the headquarter or the branch data. The data upload depends by the values reported in each object: * if the object contains only the **identificationCode** the system upload only the headquarter data. * if the object contains the **identificationCode** and the **companyUnitId** the system upload only the branch data.  Values reported in the &#x60;value&#x60; items (customVariable array) shall respect the data type defined in the **POST /portfolios**.  &#x60;margoId&#x60; and &#x60;isEnriched&#x60; are read-only fields and must not be provided in input. &#x60;margoId&#x60; is generated by Margo system during the company upload in order to identify a company inside portfolio. &#x60;isEnriched&#x60; indicate if a company is enriched or not. | 

### Return type

[**[]CompanyResultType**](CompanyResultType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDeletePortfolioCompanies

> []CompanyResultType PostDeletePortfolioCompanies(ctx, portfolioId).CompanyMargoIdArrayType(companyMargoIdArrayType).Execute()

Delete companies from portfolio

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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	companyMargoIdArrayType := *openapiclient.NewCompanyMargoIdArrayType() // CompanyMargoIdArrayType | Based on portfolioId parameter Margo delete a subset of companies belongs to the related portfolio.  The json body to provide is an array with `margoId` associated an each company belongs to the portfolio.  The margoIds are available with the method **POST /portfolios/{portfolioId}/download** (data packet: **portfolioInfo**).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostDeletePortfolioCompanies(context.Background(), portfolioId).CompanyMargoIdArrayType(companyMargoIdArrayType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostDeletePortfolioCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDeletePortfolioCompanies`: []CompanyResultType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostDeletePortfolioCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDeletePortfolioCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyMargoIdArrayType** | [**CompanyMargoIdArrayType**](CompanyMargoIdArrayType.md) | Based on portfolioId parameter Margo delete a subset of companies belongs to the related portfolio.  The json body to provide is an array with &#x60;margoId&#x60; associated an each company belongs to the portfolio.  The margoIds are available with the method **POST /portfolios/{portfolioId}/download** (data packet: **portfolioInfo**). | 

### Return type

[**[]CompanyResultType**](CompanyResultType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPortfolioDownload

> DownloadResultType PostPortfolioDownload(ctx, portfolioId).AcceptLanguage(acceptLanguage).DownloadType(downloadType).Execute()

Download portfolio

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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	acceptLanguage := "acceptLanguage_example" // string | Set here the language you want in output (optional) (default to "en-GB")
	downloadType := *openapiclient.NewDownloadType(openapiclient.DownloadType_content{DataPacketArrayType: openapiclient.NewDataPacketArrayType()}) // DownloadType | This method allows to download the business information of all companies content in the portfolio or of a subset of it.  Margo retrives the data based on two values: * The parameter `portfolioId`: the guid is generated automatically by the system after the portfolio creation and it identifies uniquely the portfolio in the system. The `portfolioId` of the created portfolios are available using the method **GET /portfolios**.    * The value specified in the object `content` reported in the request body. The key must be one of `dataPacketList` or `marketingList`. The first one is an array and each value must be the dataPacketId  (for more details of the data packet availables go to the page **https://developer-cms.crifnet.com:8080/apis/margo/_**). `marketingList` enables the download of the company data with a fixed layout and it must be a string with a specific value (e.g. **marketing**).  Moreover the method returns the business information for the whole portfolio indicated or for a subset of it.            Margo retrives data of a subset of whole portfolio throught the query parameters specify inside the request body. There are two type of parameters in the request body: * `filters`: it contains all filter parameters linked to the companies' business information. Parameters are grouped based on the data type: **numericFilters**, **dateFilters**, **booleanFilters**, **stringFilters**, **domainFilters**. * `portfolioFilters`: it contains all portfolio's filter parameters associated to the companies that it contains. Based on the definition in **POST /portfolios**, custom variables are grouped based on data type: **numericFilters**, **dateFilters**, **stringFilters**.  Another important parameter that handle API's response is **continueToken**: it manages the number of calls based on the total number of companies in the portfolio and the pagination reported in the body of this method. The continueToken is provided after first call by Margo and it must be inserted in the body of next request. Moreover, it must be updated in the request body with the token returned by the previous response as it changes so as to identify each paginated response.  **Use case**:The CRM system needs to download the data packet **address** for a subset of the whole portfolio and it should consider only the companies localized in the province of Rome and with employees in the range 1-50. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostPortfolioDownload(context.Background(), portfolioId).AcceptLanguage(acceptLanguage).DownloadType(downloadType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostPortfolioDownload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortfolioDownload`: DownloadResultType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostPortfolioDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPortfolioDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** | Set here the language you want in output | [default to &quot;en-GB&quot;]
 **downloadType** | [**DownloadType**](DownloadType.md) | This method allows to download the business information of all companies content in the portfolio or of a subset of it.  Margo retrives the data based on two values: * The parameter &#x60;portfolioId&#x60;: the guid is generated automatically by the system after the portfolio creation and it identifies uniquely the portfolio in the system. The &#x60;portfolioId&#x60; of the created portfolios are available using the method **GET /portfolios**.    * The value specified in the object &#x60;content&#x60; reported in the request body. The key must be one of &#x60;dataPacketList&#x60; or &#x60;marketingList&#x60;. The first one is an array and each value must be the dataPacketId  (for more details of the data packet availables go to the page **https://developer-cms.crifnet.com:8080/apis/margo/_**). &#x60;marketingList&#x60; enables the download of the company data with a fixed layout and it must be a string with a specific value (e.g. **marketing**).  Moreover the method returns the business information for the whole portfolio indicated or for a subset of it.            Margo retrives data of a subset of whole portfolio throught the query parameters specify inside the request body. There are two type of parameters in the request body: * &#x60;filters&#x60;: it contains all filter parameters linked to the companies&#39; business information. Parameters are grouped based on the data type: **numericFilters**, **dateFilters**, **booleanFilters**, **stringFilters**, **domainFilters**. * &#x60;portfolioFilters&#x60;: it contains all portfolio&#39;s filter parameters associated to the companies that it contains. Based on the definition in **POST /portfolios**, custom variables are grouped based on data type: **numericFilters**, **dateFilters**, **stringFilters**.  Another important parameter that handle API&#39;s response is **continueToken**: it manages the number of calls based on the total number of companies in the portfolio and the pagination reported in the body of this method. The continueToken is provided after first call by Margo and it must be inserted in the body of next request. Moreover, it must be updated in the request body with the token returned by the previous response as it changes so as to identify each paginated response.  **Use case**:The CRM system needs to download the data packet **address** for a subset of the whole portfolio and it should consider only the companies localized in the province of Rome and with employees in the range 1-50. | 

### Return type

[**DownloadResultType**](DownloadResultType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPortfolios

> PortfolioSummaryType PostPortfolios(ctx).PortfolioCreationType(portfolioCreationType).Execute()

Create portfolio

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
	portfolioCreationType := *openapiclient.NewPortfolioCreationType() // PortfolioCreationType | This method allows to create a `Portfolio`.   The JSON body could contain the following fields: * name: it contains the name of the portfolio (the field must be unique and it's mandatory). * customVariables: this array shall contain the information for custom variables present in the portfolio (the field is not mandatory).  Information for the portfolios' custom variables: * There can be at max 6 custom variables inside the portfolio and the number is handled by the `index` item. The values accepted are from 0 to 5.  * Custom variable `name` must be unique in the portfolio. * The `type` must be one of the following: String, Number, Date.  This method allow only the portfolio creation. The addition of companies is handled by a different method: **POST /portfolios/{portfolioId}/companies**.  **Use case:** The CRM system needs to create a portfolio in Margo whit the own customer base. It would identifies this portfolio with name OwnCustomerBase and he would to create two custom variables: * Customer segment: it describes the customer segment class according to revenue volume. * Quantity purchased: it determines the quantity of product sold to the customer.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostPortfolios(context.Background()).PortfolioCreationType(portfolioCreationType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostPortfolios``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPortfolios`: PortfolioSummaryType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostPortfolios`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostPortfoliosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **portfolioCreationType** | [**PortfolioCreationType**](PortfolioCreationType.md) | This method allows to create a &#x60;Portfolio&#x60;.   The JSON body could contain the following fields: * name: it contains the name of the portfolio (the field must be unique and it&#39;s mandatory). * customVariables: this array shall contain the information for custom variables present in the portfolio (the field is not mandatory).  Information for the portfolios&#39; custom variables: * There can be at max 6 custom variables inside the portfolio and the number is handled by the &#x60;index&#x60; item. The values accepted are from 0 to 5.  * Custom variable &#x60;name&#x60; must be unique in the portfolio. * The &#x60;type&#x60; must be one of the following: String, Number, Date.  This method allow only the portfolio creation. The addition of companies is handled by a different method: **POST /portfolios/{portfolioId}/companies**.  **Use case:** The CRM system needs to create a portfolio in Margo whit the own customer base. It would identifies this portfolio with name OwnCustomerBase and he would to create two custom variables: * Customer segment: it describes the customer segment class according to revenue volume. * Quantity purchased: it determines the quantity of product sold to the customer. | 

### Return type

[**PortfolioSummaryType**](PortfolioSummaryType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSimilarity

> SimilarityResultType PostSimilarity(ctx, portfolioId).CompanyPortfolioStatus(companyPortfolioStatus).Execute()

Generate prospect based on own portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	companyPortfolioStatus := "companyPortfolioStatus_example" // string | It indicates one of company portfolio status available in own portfolio.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostSimilarity(context.Background(), portfolioId).CompanyPortfolioStatus(companyPortfolioStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostSimilarity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSimilarity`: SimilarityResultType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostSimilarity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSimilarityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyPortfolioStatus** | **string** | It indicates one of company portfolio status available in own portfolio. | 

### Return type

[**SimilarityResultType**](SimilarityResultType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUpdatePortfolioCompanies

> []CompanyResultType PostUpdatePortfolioCompanies(ctx, portfolioId).CompanyUpdateArrayType(companyUpdateArrayType).Execute()

Update companies to portfolio

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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	companyUpdateArrayType := *openapiclient.NewCompanyUpdateArrayType() // CompanyUpdateArrayType | The service updates one or more companies of an existing portfolio. The information that is possible to update are only the portfolio information: companyPortfolioStatus, tags, amount, areaManagerId, salesId, productId and customVariables’ value.  The key field in order to allow the update of company data in the portfolio is the `margoId` - this information is available for each company belongs to the portfolio using the method **POST /portfolios/{portfolioId}/download** (data packet: **portfolioInfo**).  `companies` list needs to have maximum 100 objects per call and each object belonging to a list must contain, for each company, at least the fileds: margoId, identificationCode and companyPortfolioStatus.  `isEnriched` is read-only fields and must not be provided in input. These items are generated by Margo system  and it indicate if a company is enrinched or not.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfoliosAPI.PostUpdatePortfolioCompanies(context.Background(), portfolioId).CompanyUpdateArrayType(companyUpdateArrayType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.PostUpdatePortfolioCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostUpdatePortfolioCompanies`: []CompanyResultType
	fmt.Fprintf(os.Stdout, "Response from `PortfoliosAPI.PostUpdatePortfolioCompanies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostUpdatePortfolioCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyUpdateArrayType** | [**CompanyUpdateArrayType**](CompanyUpdateArrayType.md) | The service updates one or more companies of an existing portfolio. The information that is possible to update are only the portfolio information: companyPortfolioStatus, tags, amount, areaManagerId, salesId, productId and customVariables’ value.  The key field in order to allow the update of company data in the portfolio is the &#x60;margoId&#x60; - this information is available for each company belongs to the portfolio using the method **POST /portfolios/{portfolioId}/download** (data packet: **portfolioInfo**).  &#x60;companies&#x60; list needs to have maximum 100 objects per call and each object belonging to a list must contain, for each company, at least the fileds: margoId, identificationCode and companyPortfolioStatus.  &#x60;isEnriched&#x60; is read-only fields and must not be provided in input. These items are generated by Margo system  and it indicate if a company is enrinched or not. | 

### Return type

[**[]CompanyResultType**](CompanyResultType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharePortfolio

> SharePortfolio(ctx, portfolioId).AccessRight(accessRight).Execute()

Share a Portfolio



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
	portfolioId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | It is a `Portfolio`'s unique identifier. It must contains the portfolio's identifier indicated in  the method **GET /portfolios**.
	accessRight := *openapiclient.NewAccessRight("AccessRight_example", []string{"d069e399-894b-46be-857b-fbccd497d9b3"}) // AccessRight | `accessRight` defines which operation can be done by web users for portfolio received. Values can be one of: - **Read**: it doesn’t allow to edit the API portfolio to web users but they can see inside the platform in read only way - **Write**: it allows to see and edit the API’s portfolio to all web users (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PortfoliosAPI.SharePortfolio(context.Background(), portfolioId).AccessRight(accessRight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfoliosAPI.SharePortfolio``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**portfolioId** | **string** | It is a &#x60;Portfolio&#x60;&#39;s unique identifier. It must contains the portfolio&#39;s identifier indicated in  the method **GET /portfolios**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSharePortfolioRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRight** | [**AccessRight**](AccessRight.md) | &#x60;accessRight&#x60; defines which operation can be done by web users for portfolio received. Values can be one of: - **Read**: it doesn’t allow to edit the API portfolio to web users but they can see inside the platform in read only way - **Write**: it allows to see and edit the API’s portfolio to all web users | 

### Return type

 (empty response body)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

