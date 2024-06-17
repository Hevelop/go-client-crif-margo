# # CompanyResultType


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MargoId**| **string** | The guid reported in this item is generated automatically by Margo system and it is the unique code that identifies the company in the portfolio.  |
**Result**| **string** | The value reported indicates if the system accepted the company information. * Success: ok * Ceased: company is ceased and the system rejected it * Unmatched: company is not present * Invalid: company is not valid * Validated: company is formally correct * Error:  the company information contains some error. The error  details are reported in the &#x60;errors&#x60; item. for more information please, see Model/string.php  |
**Errors**| [**[]ErrorType**](ErrorType.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

