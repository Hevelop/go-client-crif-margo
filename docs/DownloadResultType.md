# # DownloadResultType
companies list with the Marketing List or datapackets requested

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalElements**| **int32** | Total number of elements the API should return without pagination  | [optional]
**TotalPages**| **int64** | Total number of pages available &#x60;totalPages &#x3D; ceil (totalElements / size)&#x60;  | [optional]
**Size**| **int64** | maximum number of elements inside a page (like input)  | [optional]
**Page**| **int64** | page number (like input)  | [optional]
**NumberOfElements**| **int64** | number of elements in this page &#x60;numberOfElements &lt;&#x3D; size&#x60;  | [optional]
**Content**| [**[]DownloadCompanyDataType**](DownloadCompanyDataType.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

