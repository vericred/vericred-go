# State

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Primary Key of State | [optional] [default to null]
**Name** | **string** | Name of state | [optional] [default to null]
**Code** | **string** | 2 letter code for state | [optional] [default to null]
**FipsNumber** | **string** | National FIPs number | [optional] [default to null]
**LastDateForIndividual** | [**time.Time**](time.Time.md) | Last date this state is live for individuals | [optional] [default to null]
**LastDateForShop** | [**time.Time**](time.Time.md) | Last date this state is live for shop | [optional] [default to null]
**LiveForBusiness** | **bool** | Is this State available for businesses | [optional] [default to null]
**LiveForConsumers** | **bool** | Is this State available for individuals | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


