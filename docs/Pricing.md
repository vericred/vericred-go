# Pricing

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **int32** | Age of applicant | [optional] [default to null]
**EffectiveDate** | [**time.Time**](time.Time.md) | Effective date of plan | [optional] [default to null]
**ExpirationDate** | [**time.Time**](time.Time.md) | Plan expiration date | [optional] [default to null]
**PlanId** | **int32** | Foreign key to plans | [optional] [default to null]
**PremiumChildOnly** | **float32** | Child-only premium | [optional] [default to null]
**PremiumFamily** | **float32** | Family premium | [optional] [default to null]
**PremiumSingle** | **float32** | Single-person premium | [optional] [default to null]
**PremiumSingleAndChildren** | **float32** | Single person including children premium | [optional] [default to null]
**PremiumSingleAndSpouse** | **float32** | Person with spouse premium | [optional] [default to null]
**PremiumSingleSmoker** | **float32** | Premium for single smoker | [optional] [default to null]
**RatingAreaId** | **string** | Foreign key to rating areas | [optional] [default to null]
**PremiumSource** | **string** | Where was this pricing data extracted from? | [optional] [default to null]
**UpdatedAt** | **string** | Time when pricing was last updated | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


