# Hevelop\CrifMargo\ProspectingAPI

All URIs are relative to *https://api-uat.crif.com/margo/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSearch**](ProspectingAPI.md#PostSearch) | **Post** /prospecting/search | Search companies



## PostSearch

> SearchDataArrayMetadataType PostSearch(ctx).AcceptLanguage(acceptLanguage).SearchType(searchType).Execute()

Search companies

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
	searchType := *openapiclient.NewSearchType() // SearchType | This method allows to search a list of prospect setting a series of filters and it could download companies' data packets.  It's possible to search companies based on the value reported in the body: * `freeText`: search field that enables to find a company by name or identifier. * `ignoredPortfolios`: the array contains `portfoloIds` associated by the system to the portfolios (information are available using the method **GET /portfolios**). The method excludes from the response all the companies belong to the portfolios specifiend in the array. * `filters`: it contains all filter parameters linked to the companies' business information. Parameters are grouped based on the data type: **numericFilters**, **dateFilters**, **booleanFilters**, **stringFilters**, **domainFilters**.  How can you perform a search? * if you search a specific company (e.g.: CRIF S.p.A.), you can insert in the body **only** the item `freeText`. * if you would like to search for a list of companies based on a subset of parameters, it's useful to insert in the body the items `ignoredPortfolios` and `filters` in order to perform the query (`freeText` might not present in the body).    In case you are going to use filter parameters provide by **POST /portfolios/{portfolioId}/create-similarity**, you should insert in `filters` the parameters provided by similarity method and `ignoredPortfolios` parameter should contain the **portfolioId** of the starting portfolio.   Data reported in the response depends by the presence of `content` item: * if the body doesn't contain it, the system provides a subset of information necessary to identify companies provided. * if the item is present, the system provides the data required based on the parameters: `dataPacketList` or `marketingList` (take a look: `content` must contain one of `dataPacketList` or `marketingList`). The first one is an array and each value must be the dataPacketId (for more details of the data packet availables go to the page **https://developer-cms.crifnet.com:8080/apis/margo/_**). `marketingList` enables the download of the company data with a fixed layout and it must be a string with a specific value (e.g. **marketing**).  An important parameter that handles the API response is the **continueToken**: it manages the number of calls based on the total number of companies searched and the pagination reported in the body of this method. The continueToken is provided after first call by Margo and it must be inserted in the body of next request. Moreover, it must be updated in the request body with the token returned by the previous response as it changes so as to identify each paginated response.  Finally, you could sorting the result by company name, turnover and  employee using the item `sort`. The format to use is `fieldId,(asc|desc)` and the fields are: * company name: **companyName** * turnover: **turnover** * employees: **employee**  Multiple sort criteria are not supported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProspectingAPI.PostSearch(context.Background()).AcceptLanguage(acceptLanguage).SearchType(searchType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProspectingAPI.PostSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSearch`: SearchDataArrayMetadataType
	fmt.Fprintf(os.Stdout, "Response from `ProspectingAPI.PostSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acceptLanguage** | **string** | Set here the language you want in output | [default to &quot;en-GB&quot;]
 **searchType** | [**SearchType**](SearchType.md) | This method allows to search a list of prospect setting a series of filters and it could download companies&#39; data packets.  It&#39;s possible to search companies based on the value reported in the body: * &#x60;freeText&#x60;: search field that enables to find a company by name or identifier. * &#x60;ignoredPortfolios&#x60;: the array contains &#x60;portfoloIds&#x60; associated by the system to the portfolios (information are available using the method **GET /portfolios**). The method excludes from the response all the companies belong to the portfolios specifiend in the array. * &#x60;filters&#x60;: it contains all filter parameters linked to the companies&#39; business information. Parameters are grouped based on the data type: **numericFilters**, **dateFilters**, **booleanFilters**, **stringFilters**, **domainFilters**.  How can you perform a search? * if you search a specific company (e.g.: CRIF S.p.A.), you can insert in the body **only** the item &#x60;freeText&#x60;. * if you would like to search for a list of companies based on a subset of parameters, it&#39;s useful to insert in the body the items &#x60;ignoredPortfolios&#x60; and &#x60;filters&#x60; in order to perform the query (&#x60;freeText&#x60; might not present in the body).    In case you are going to use filter parameters provide by **POST /portfolios/{portfolioId}/create-similarity**, you should insert in &#x60;filters&#x60; the parameters provided by similarity method and &#x60;ignoredPortfolios&#x60; parameter should contain the **portfolioId** of the starting portfolio.   Data reported in the response depends by the presence of &#x60;content&#x60; item: * if the body doesn&#39;t contain it, the system provides a subset of information necessary to identify companies provided. * if the item is present, the system provides the data required based on the parameters: &#x60;dataPacketList&#x60; or &#x60;marketingList&#x60; (take a look: &#x60;content&#x60; must contain one of &#x60;dataPacketList&#x60; or &#x60;marketingList&#x60;). The first one is an array and each value must be the dataPacketId (for more details of the data packet availables go to the page **https://developer-cms.crifnet.com:8080/apis/margo/_**). &#x60;marketingList&#x60; enables the download of the company data with a fixed layout and it must be a string with a specific value (e.g. **marketing**).  An important parameter that handles the API response is the **continueToken**: it manages the number of calls based on the total number of companies searched and the pagination reported in the body of this method. The continueToken is provided after first call by Margo and it must be inserted in the body of next request. Moreover, it must be updated in the request body with the token returned by the previous response as it changes so as to identify each paginated response.  Finally, you could sorting the result by company name, turnover and  employee using the item &#x60;sort&#x60;. The format to use is &#x60;fieldId,(asc|desc)&#x60; and the fields are: * company name: **companyName** * turnover: **turnover** * employees: **employee**  Multiple sort criteria are not supported. | 

### Return type

[**SearchDataArrayMetadataType**](SearchDataArrayMetadataType.md)

### Authorization

[Password](../README.md#Password)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

