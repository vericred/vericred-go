# \ZipCountiesApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetZipCounties**](ZipCountiesApi.md#GetZipCounties) | **Get** /zip_counties | Search for Zip Counties
[**ShowZipCounty**](ZipCountiesApi.md#ShowZipCounty) | **Get** /zip_counties/{id} | Show an individual ZipCounty


# **GetZipCounties**
> ZipCountiesResponse GetZipCounties($zipPrefix)

Search for Zip Counties

Our `Plan` endpoints require a zip code and a fips (county) code.  This is because plan pricing requires both of these elements.  Users are unlikely to know their fips code, so we provide this endpoint to look up a `ZipCounty` by zip code and return both the selected zip and fips codes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zipPrefix** | **string**| Partial five-digit Zip | 

### Return type

[**ZipCountiesResponse**](ZipCountiesResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowZipCounty**
> ZipCountyResponse ShowZipCounty($id)

Show an individual ZipCounty

Our `Plan` endpoints require a zip code and a fips (county) code.  This is because plan pricing requires both of these elements.  Users are unlikely to know their fips code, so we provide this endpoint to returns the details for a `ZipCounty` by zip code and return both the selected zip and fips codes.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Unique ID for ZipCounty | 

### Return type

[**ZipCountyResponse**](ZipCountyResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

