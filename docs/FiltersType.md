# # FiltersType
It must contain all filter parameters linked to the companies&#39; business information.  All filter parameters are available to the following page: https://developer.crif.com/apis/margo/.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numericfilters**| [**[]FiltersTypeNumericfiltersInner**](FiltersTypeNumericfiltersInner.md) | The array contains all fields with data type **numeric**.  The object must have the following structure: * filterId: it identifies the name of the field used as a filter * properties: this object have to contain the range of value to be filtered.  | [optional]
**Datefilters**| [**[]FiltersTypeDatefiltersInner**](FiltersTypeDatefiltersInner.md) | The array contains all fields with data type ***date***.  The object must have the following structure: * filterId: it identifies the name of the field used as a filter * properties: this object have to contain the range of date to be filtered.  | [optional]
**Booleanfilters**| [**[]FiltersTypeBooleanfiltersInner**](FiltersTypeBooleanfiltersInner.md) | The array contains all fields with data type ***boolean***.  The object must have the following structure: * filterId: it identifies the name of the field used as a filter * value: the field have to contain a boolean value (true or false).  | [optional]
**Stringfilters**| [**[]FiltersTypeStringfiltersInner**](FiltersTypeStringfiltersInner.md) | The array contains all fields with data type ***string***.  The object must have the following structure: * filterId: it identifies the name of the field used as a filter * value: the field have to contain a label.  | [optional]
**Domainfilters**| [**[]FiltersTypeDomainfiltersInner**](FiltersTypeDomainfiltersInner.md) | The array contains all fields with data type ***domain***.  The object must have the following structure: * filterId: it identifies the name of the field used as a filter * codes: the field have to contain a label.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

