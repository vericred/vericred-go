# RequestProvidersSearch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptsInsurance** | **bool** | Limit results to Providers who accept at least one insurance         plan.  Note that the inverse of this filter is not supported and         any value will evaluate to true | [optional] [default to null]
**HiosIds** | **[]string** | List of HIOS ids | [optional] [default to null]
**MinScore** | **float32** | Minimum search threshold to be included in the results | [optional] [default to null]
**Page** | **int32** | Page number | [optional] [default to null]
**PerPage** | **int32** | Number of records to return per page | [optional] [default to null]
**Radius** | **int32** | Radius (in miles) to use to limit results | [optional] [default to null]
**SearchTerm** | **string** | String to search by | [optional] [default to null]
**ZipCode** | **string** | Zip Code to search near | [optional] [default to null]
**Type_** | **string** | Either organization or individual | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


