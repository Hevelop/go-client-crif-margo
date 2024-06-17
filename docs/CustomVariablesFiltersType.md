# # CustomVariablesFiltersType
One portfolio&#39;s feature are the custom variables and their data type could be: numeric, date, boolean, string and domain.  Like &#x60;filters&#x60; object, also &#x60;customVariablesFilters&#x60; object contains a key for each data type.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numericfilters**| [**[]CustomVariablesFiltersTypeNumericfiltersInner**](CustomVariablesFiltersTypeNumericfiltersInner.md) | The array contains the custom variable with data type **numeric**.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: this object have to contain the range of value to be filtered.  | [optional]
**Datefilters**| [**[]CustomVariablesFiltersTypeDatefiltersInner**](CustomVariablesFiltersTypeDatefiltersInner.md) | The array contains the custom variable with data type **date**.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: this object have to contain the range of value to be filtered.  | [optional]
**Booleanfilters**| [**[]CustomVariablesFiltersTypeBooleanfiltersInner**](CustomVariablesFiltersTypeBooleanfiltersInner.md) | The array contains all fields with data type ***boolean***.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: the field have to contain a boolean value (true or false).  | [optional]
**Stringfilters**| [**[]CustomVariablesFiltersTypeStringfiltersInner**](CustomVariablesFiltersTypeStringfiltersInner.md) | The array contains all fields with data type ***string***.  The object must have the following structure: * variableIndex: it identifies the index associated at custom variable * value: the field have to contain a label.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

