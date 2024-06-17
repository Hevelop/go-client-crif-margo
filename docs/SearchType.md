# # SearchType


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeText**| **string** |   | [optional]
**IgnoredPortfolios**| **[]string** |   | [optional]
**Filters**| [**FiltersType**](FiltersType.md) |   | [optional]
**Content**| [**SearchTypeContent**](SearchTypeContent.md) |   | [optional]
**Sort**| **string** | It defines the output sorting. Must be either ASC or DESC and the format must be: &#x60;fieldId,asc&#x60; or &#x60;fieldId,desc&#x60; where fieldId identifies the field for sorting.  | [optional]
**ContinueToken**| **string** | It manages the number of calls and it governs pagination. continueToken is provided after first call by Margo and after it must be insert in the body. The continueToken changes after each call and it must refresh in the body.  | [optional]
**Size**| **float32** | Zero-based number of the page to obtain. It must be higher than 0 and ≤ than 100  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

