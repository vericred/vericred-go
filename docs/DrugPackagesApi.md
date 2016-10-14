# \DrugPackagesApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowFormularyDrugPackageCoverage**](DrugPackagesApi.md#ShowFormularyDrugPackageCoverage) | **Get** /formularies/{formulary_id}/drug_packages/{ndc_package_code} | Formulary Drug Package Search


# **ShowFormularyDrugPackageCoverage**
> FormularyDrugPackageResponse ShowFormularyDrugPackageCoverage($formularyId, $ndcPackageCode)

Formulary Drug Package Search

Search for coverage by Formulary and DrugPackage ID


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **formularyId** | **string**| ID of the Formulary in question | 
 **ndcPackageCode** | **string**| ID of the DrugPackage in question | 

### Return type

[**FormularyDrugPackageResponse**](FormularyDrugPackageResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

