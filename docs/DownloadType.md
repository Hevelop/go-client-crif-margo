# # DownloadType
It defines the ouput requested in term of what is the marketing list  or data packets.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size**| **float32** | Zero-based number of the page to obtain  | [optional]
**ContinueToken**| **string** | It manages the number of calls and it governs pagination. continueToken is provided after first call by Margo and after it must be insert in the body. The continueToken changes after each call and it must refresh in the body.  | [optional]
**Filters**| [**FiltersType**](FiltersType.md) |   | [optional]
**PortfolioFilters**| [**PortfolioFiltersType**](PortfolioFiltersType.md) |   | [optional]
**Content**| [**DownloadTypeContent**](DownloadTypeContent.md) |   |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

