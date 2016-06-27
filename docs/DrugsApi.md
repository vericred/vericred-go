# \DrugsApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDrugCoverages**](DrugsApi.md#GetDrugCoverages) | **Get** /drug_packages/{ndc_package_code}/coverages | Search for DrugCoverages
[**ListDrugs**](DrugsApi.md#ListDrugs) | **Get** /drugs | Drug Search


# **GetDrugCoverages**
> DrugCoverageResponse GetDrugCoverages($ndcPackageCode, $audience, $stateCode)

Search for DrugCoverages

Drug Coverages are the specific tier level, quantity limit, prior
authorization and step therapy for a given Drug/Plan combination. This endpoint
returns all DrugCoverages for a given Drug


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ndcPackageCode** | **string**| NDC package code | 
 **audience** | **string**| Two-character state code | 
 **stateCode** | **string**| Two-character state code | 

### Return type

[**DrugCoverageResponse**](DrugCoverageResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDrugs**
> DrugSearchResponse ListDrugs($searchTerm)

Drug Search

Search for drugs by proprietary name


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **string**| Full or partial proprietary name of drug | 

### Return type

[**DrugSearchResponse**](DrugSearchResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

