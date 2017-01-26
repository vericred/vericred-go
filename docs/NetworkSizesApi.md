# \NetworkSizesApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListStateNetworkSizes**](NetworkSizesApi.md#ListStateNetworkSizes) | **Get** /states/{state_id}/network_sizes | State Network Sizes
[**SearchNetworkSizes**](NetworkSizesApi.md#SearchNetworkSizes) | **Post** /network_sizes/search | Network Sizes


# **ListStateNetworkSizes**
> StateNetworkSizeResponse ListStateNetworkSizes($stateId, $page, $perPage)

State Network Sizes

The number of in-network Providers for each network in a given state. This data is recalculated nightly.  The endpoint is paginated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateId** | **string**| State code | 
 **page** | **int32**| Page of paginated response | [optional] 
 **perPage** | **int32**| Responses per page | [optional] 

### Return type

[**StateNetworkSizeResponse**](StateNetworkSizeResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchNetworkSizes**
> StateNetworkSizeResponse SearchNetworkSizes($body)

Network Sizes

The number of in-network Providers for each network/state combination provided. This data is recalculated nightly.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StateNetworkSizeRequest**](StateNetworkSizeRequest.md)|  | 

### Return type

[**StateNetworkSizeResponse**](StateNetworkSizeResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

