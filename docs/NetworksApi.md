# \NetworksApi

All URIs are relative to *https://api.vericred.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworks**](NetworksApi.md#ListNetworks) | **Get** /networks | Networks
[**ShowNetwork**](NetworksApi.md#ShowNetwork) | **Get** /networks/{id} | Network Details


# **ListNetworks**
> NetworkSearchResponse ListNetworks($carrierId, $page, $perPage)

Networks

A network is a list of the doctors, other health care providers, and hospitals that a plan has contracted with to provide medical care to its members. This endpoint is paginated.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **carrierId** | **string**| Carrier HIOS Issuer ID | 
 **page** | **int32**| Page of paginated response | [optional] 
 **perPage** | **int32**| Responses per page | [optional] 

### Return type

[**NetworkSearchResponse**](NetworkSearchResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ShowNetwork**
> NetworkDetailsResponse ShowNetwork($id)

Network Details

A network is a list of the doctors, other health care providers, and hospitals that a plan has contracted with to provide medical care to its members.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int32**| Primary key of the network | 

### Return type

[**NetworkDetailsResponse**](NetworkDetailsResponse.md)

### Authorization

[Vericred-Api-Key](../README.md#Vericred-Api-Key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

