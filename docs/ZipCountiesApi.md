# \ZipCountiesApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZipCounties**](ZipCountiesApi.md#GetZipCounties) | **Get** /zip_counties | Search for Zip Counties


# **GetZipCounties**
> ZipCountyResponse GetZipCounties($zipPrefix, $vericredApiKey)

Search for Zip Counties

Our `Plan` endpoints require a zip code and a fips (county) code.  This is because plan pricing requires both of these elements.  Users are unlikely to know their fips code, so we provide this endpoint to look up a `ZipCounty` by zip code and return both the selected zip and fips codes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zipPrefix** | **string**| Partial five-digit Zip | 
 **vericredApiKey** | **string**| API Key | [optional] 

### Return type

[**ZipCountyResponse**](ZipCountyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

