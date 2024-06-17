# # CompanyUpdateType


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MargoId**| **string** | The guid reported in this item is generated automatically by Margo system and it is the unique code that identifies the company in the portfolio.  |
**CompanyPortfolioStatus**| **string** |  for more information please, see Model/string.php  | [optional]
**Tags**| **string** | It&#39;s possible to insert one or more tag divided by comma. Max occurrence is 10 values.  | [optional]
**Amount**| **float32** | It identifies the amount billed to the customer.  | [optional]
**AreaManagerId**| **string** | It must be an Area manager user id defined for the organization.  | [optional]
**SalesId**| **string** | It must be a Sales user id defined for the organization.  | [optional]
**ProductId**| **string** | It must be a Product Id defined for the organization.  | [optional]
**IsEnriched**| **bool** | It indicates if a company has been enriched from the web users  | [optional]
**CustomVariables**| [**[]CustomVariableType**](CustomVariableType.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

