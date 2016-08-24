# RequestPlanFind

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applicants** | [**[]RequestPlanFindApplicant**](RequestPlanFindApplicant.md) | Applicants for desired plans. | [optional] [default to null]
**EnrollmentDate** | **string** | Date of enrollment | [optional] [default to null]
**DrugPackages** | [**[]RequestPlanFindDrugPackage**](RequestPlanFindDrugPackage.md) | National Drug Code Package Id | [optional] [default to null]
**FipsCode** | **string** | County code to determine eligibility | [optional] [default to null]
**HouseholdIncome** | **int32** | Total household income. | [optional] [default to null]
**HouseholdSize** | **int32** | Number of people living in household. | [optional] [default to null]
**Ids** | **[]int32** | List of plan IDs to filter by | [optional] [default to null]
**Market** | **string** | Type of plan to search for. | [optional] [default to null]
**Providers** | [**[]RequestPlanFindProvider**](RequestPlanFindProvider.md) | List of providers to search for. | [optional] [default to null]
**Page** | **int32** | Selected page of paginated response. | [optional] [default to null]
**PerPage** | **int32** | Results per page of response. | [optional] [default to null]
**Sort** | **string** | Sort responses by plan field. | [optional] [default to null]
**ZipCode** | **string** | 5-digit zip code - this helps determine pricing. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


