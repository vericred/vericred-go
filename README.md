# Go API client for vericredclient

Vericred's API allows you to search for Health Plans that a specific doctor
accepts.

## Getting Started

Visit our [Developer Portal](https://developers.vericred.com) to
create an account.

Once you have created an account, you can create one Application for
Production and another for our Sandbox (select the appropriate Plan when
you create the Application).

## SDKs

Our API follows standard REST conventions, so you can use any HTTP client
to integrate with us. You will likely find it easier to use one of our
[autogenerated SDKs](https://github.com/vericred/?query=vericred-),
which we make available for several common programming languages.

## Authentication

To authenticate, pass the API Key you created in the Developer Portal as
a `Vericred-Api-Key` header.

`curl -H 'Vericred-Api-Key: YOUR_KEY' "https://api.vericred.com/providers?search_term=Foo&zip_code=11215"`

## Versioning

Vericred's API default to the latest version.  However, if you need a specific
version, you can request it with an `Accept-Version` header.

The current version is `v3`.  Previous versions are `v1` and `v2`.

`curl -H 'Vericred-Api-Key: YOUR_KEY' -H 'Accept-Version: v2' "https://api.vericred.com/providers?search_term=Foo&zip_code=11215"`

## Pagination

Endpoints that accept `page` and `per_page` parameters are paginated. They expose
four additional fields that contain data about your position in the response,
namely `Total`, `Per-Page`, `Link`, and `Page` as described in [RFC-5988](https://tools.ietf.org/html/rfc5988).

For example, to display 5 results per page and view the second page of a
`GET` to `/networks`, your final request would be `GET /networks?....page=2&per_page=5`.

## Sideloading

When we return multiple levels of an object graph (e.g. `Provider`s and their `State`s
we sideload the associated data.  In this example, we would provide an Array of
`State`s and a `state_id` for each provider.  This is done primarily to reduce the
payload size since many of the `Provider`s will share a `State`

```
{
  providers: [{ id: 1, state_id: 1}, { id: 2, state_id: 1 }],
  states: [{ id: 1, code: 'NY' }]
}
```

If you need the second level of the object graph, you can just match the
corresponding id.

## Selecting specific data

All endpoints allow you to specify which fields you would like to return.
This allows you to limit the response to contain only the data you need.

For example, let's take a request that returns the following JSON by default

```
{
  provider: {
    id: 1,
    name: 'John',
    phone: '1234567890',
    field_we_dont_care_about: 'value_we_dont_care_about'
  },
  states: [{
    id: 1,
    name: 'New York',
    code: 'NY',
    field_we_dont_care_about: 'value_we_dont_care_about'
  }]
}
```

To limit our results to only return the fields we care about, we specify the
`select` query string parameter for the corresponding fields in the JSON
document.

In this case, we want to select `name` and `phone` from the `provider` key,
so we would add the parameters `select=provider.name,provider.phone`.
We also want the `name` and `code` from the `states` key, so we would
add the parameters `select=states.name,staes.code`.  The id field of
each document is always returned whether or not it is requested.

Our final request would be `GET /providers/12345?select=provider.name,provider.phone,states.name,states.code`

The response would be

```
{
  provider: {
    id: 1,
    name: 'John',
    phone: '1234567890'
  },
  states: [{
    id: 1,
    name: 'New York',
    code: 'NY'
  }]
}
```

## Benefits summary format
Benefit cost-share strings are formatted to capture:
 * Network tiers
 * Compound or conditional cost-share
 * Limits on the cost-share
 * Benefit-specific maximum out-of-pocket costs

**Example #1**
As an example, we would represent [this Summary of Benefits &amp; Coverage](https://s3.amazonaws.com/vericred-data/SBC/2017/33602TX0780032.pdf) as:

* **Hospital stay facility fees**:
  - Network Provider: `$400 copay/admit plus 20% coinsurance`
  - Out-of-Network Provider: `$1,500 copay/admit plus 50% coinsurance`
  - Vericred's format for this benefit: `In-Network: $400 before deductible then 20% after deductible / Out-of-Network: $1,500 before deductible then 50% after deductible`

* **Rehabilitation services:**
  - Network Provider: `20% coinsurance`
  - Out-of-Network Provider: `50% coinsurance`
  - Limitations & Exceptions: `35 visit maximum per benefit period combined with Chiropractic care.`
  - Vericred's format for this benefit: `In-Network: 20% after deductible / Out-of-Network: 50% after deductible | limit: 35 visit(s) per Benefit Period`

**Example #2**
In [this other Summary of Benefits &amp; Coverage](https://s3.amazonaws.com/vericred-data/SBC/2017/40733CA0110568.pdf), the **specialty_drugs** cost-share has a maximum out-of-pocket for in-network pharmacies.
* **Specialty drugs:**
  - Network Provider: `40% coinsurance up to a $500 maximum for up to a 30 day supply`
  - Out-of-Network Provider `Not covered`
  - Vericred's format for this benefit: `In-Network: 40% after deductible, up to $500 per script / Out-of-Network: 100%`

**BNF**

Here's a description of the benefits summary string, represented as a context-free grammar:

```
root                      ::= coverage

coverage                  ::= (simple_coverage | tiered_coverage) (space pipe space coverage_limitation)?
tiered_coverage           ::= tier (space slash space tier)*
tier                      ::= tier_name colon space (tier_coverage | not_applicable)
tier_coverage             ::= simple_coverage (space (then | or | and) space simple_coverage)* tier_limitation?
simple_coverage           ::= (pre_coverage_limitation space)? coverage_amount (space post_coverage_limitation)? (comma? space coverage_condition)?
coverage_limitation       ::= "limit" colon space (((simple_coverage | simple_limitation) (semicolon space see_carrier_documentation)?) | see_carrier_documentation | waived_if_admitted)
waived_if_admitted        ::= ("copay" space)? "waived if admitted"
simple_limitation         ::= pre_coverage_limitation space "copay applies"
tier_name                 ::= "In-Network-Tier-2" | "Out-of-Network" | "In-Network"
tier_limitation           ::= comma space "up to" space (currency | (integer space time_unit plural?)) (space post_coverage_limitation)?
coverage_amount           ::= currency | unlimited | included | unknown | percentage | (digits space (treatment_unit | time_unit) plural?)
pre_coverage_limitation   ::= first space digits space time_unit plural?
post_coverage_limitation  ::= (((then space currency) | "per condition") space)? "per" space (treatment_unit | (integer space time_unit) | time_unit) plural?
coverage_condition        ::= ("before deductible" | "after deductible" | "penalty" | allowance | "in-state" | "out-of-state") (space allowance)?
allowance                 ::= upto_allowance | after_allowance
upto_allowance            ::= "up to" space (currency space)? "allowance"
after_allowance           ::= "after" space (currency space)? "allowance"
see_carrier_documentation ::= "see carrier documentation for more information"
unknown                   ::= "unknown"
unlimited                 ::= /[uU]nlimited/
included                  ::= /[iI]ncluded in [mM]edical/
time_unit                 ::= /[hH]our/ | (((/[cC]alendar/ | /[cC]ontract/) space)? /[yY]ear/) | /[mM]onth/ | /[dD]ay/ | /[wW]eek/ | /[vV]isit/ | /[lL]ifetime/ | ((((/[bB]enefit/ plural?) | /[eE]ligibility/) space)? /[pP]eriod/)
treatment_unit            ::= /[pP]erson/ | /[gG]roup/ | /[cC]ondition/ | /[sS]cript/ | /[vV]isit/ | /[eE]xam/ | /[iI]tem/ | /[sS]tay/ | /[tT]reatment/ | /[aA]dmission/ | /[eE]pisode/
comma                     ::= ","
colon                     ::= ":"
semicolon                 ::= ";"
pipe                      ::= "|"
slash                     ::= "/"
plural                    ::= "(s)" | "s"
then                      ::= "then" | ("," space) | space
or                        ::= "or"
and                       ::= "and"
not_applicable            ::= "Not Applicable" | "N/A" | "NA"
first                     ::= "first"
currency                  ::= "$" number
percentage                ::= number "%"
number                    ::= float | integer
float                     ::= digits "." digits
integer                   ::= /[0-9]/+ (comma_int | under_int)*
comma_int                 ::= ("," /[0-9]/*3) !"_"
under_int                 ::= ("_" /[0-9]/*3) !","
digits                    ::= /[0-9]/+ ("_" /[0-9]/+)*
space                     ::= /[ \t]/+
```



## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 0.0.7
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./vericredclient"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.vericred.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DrugCoveragesApi* | [**GetDrugCoverages**](docs/DrugCoveragesApi.md#getdrugcoverages) | **Get** /drug_packages/{ndc_package_code}/coverages | Search for DrugCoverages
*DrugPackagesApi* | [**ShowFormularyDrugPackageCoverage**](docs/DrugPackagesApi.md#showformularydrugpackagecoverage) | **Get** /formularies/{formulary_id}/drug_packages/{ndc_package_code} | Formulary Drug Package Search
*DrugsApi* | [**ListDrugs**](docs/DrugsApi.md#listdrugs) | **Get** /drugs | Drug Search
*NetworkSizesApi* | [**ListStateNetworkSizes**](docs/NetworkSizesApi.md#liststatenetworksizes) | **Get** /states/{state_id}/network_sizes | State Network Sizes
*NetworkSizesApi* | [**SearchNetworkSizes**](docs/NetworkSizesApi.md#searchnetworksizes) | **Post** /network_sizes/search | Network Sizes
*NetworksApi* | [**ListNetworks**](docs/NetworksApi.md#listnetworks) | **Get** /networks | Networks
*NetworksApi* | [**ShowNetwork**](docs/NetworksApi.md#shownetwork) | **Get** /networks/{id} | Network Details
*PlansApi* | [**FindPlans**](docs/PlansApi.md#findplans) | **Post** /plans/search | Find Plans
*PlansApi* | [**ShowPlan**](docs/PlansApi.md#showplan) | **Get** /plans/{id} | Show Plan
*ProvidersApi* | [**GetProvider**](docs/ProvidersApi.md#getprovider) | **Get** /providers/{npi} | Find a Provider
*ProvidersApi* | [**GetProviders**](docs/ProvidersApi.md#getproviders) | **Post** /providers/search | Find Providers
*ProvidersApi* | [**GetProviders_0**](docs/ProvidersApi.md#getproviders_0) | **Post** /providers/search/geocode | Find Providers
*ZipCountiesApi* | [**GetZipCounties**](docs/ZipCountiesApi.md#getzipcounties) | **Get** /zip_counties | Search for Zip Counties
*ZipCountiesApi* | [**ShowZipCounty**](docs/ZipCountiesApi.md#showzipcounty) | **Get** /zip_counties/{id} | Show an individual ZipCounty


## Documentation For Models

 - [Applicant](docs/Applicant.md)
 - [Base](docs/Base.md)
 - [Carrier](docs/Carrier.md)
 - [CarrierSubsidiary](docs/CarrierSubsidiary.md)
 - [County](docs/County.md)
 - [CountyBulk](docs/CountyBulk.md)
 - [Drug](docs/Drug.md)
 - [DrugCoverage](docs/DrugCoverage.md)
 - [DrugCoverageResponse](docs/DrugCoverageResponse.md)
 - [DrugPackage](docs/DrugPackage.md)
 - [DrugSearchResponse](docs/DrugSearchResponse.md)
 - [Formulary](docs/Formulary.md)
 - [FormularyDrugPackageResponse](docs/FormularyDrugPackageResponse.md)
 - [FormularyResponse](docs/FormularyResponse.md)
 - [Meta](docs/Meta.md)
 - [Network](docs/Network.md)
 - [NetworkDetails](docs/NetworkDetails.md)
 - [NetworkDetailsResponse](docs/NetworkDetailsResponse.md)
 - [NetworkSearchResponse](docs/NetworkSearchResponse.md)
 - [NetworkSize](docs/NetworkSize.md)
 - [Plan](docs/Plan.md)
 - [PlanCounty](docs/PlanCounty.md)
 - [PlanCountyBulk](docs/PlanCountyBulk.md)
 - [PlanSearchResponse](docs/PlanSearchResponse.md)
 - [PlanSearchResponseMeta](docs/PlanSearchResponseMeta.md)
 - [PlanSearchResult](docs/PlanSearchResult.md)
 - [PlanShowResponse](docs/PlanShowResponse.md)
 - [Provider](docs/Provider.md)
 - [ProviderDetails](docs/ProviderDetails.md)
 - [ProviderGeocode](docs/ProviderGeocode.md)
 - [ProviderShowResponse](docs/ProviderShowResponse.md)
 - [ProvidersGeocodeResponse](docs/ProvidersGeocodeResponse.md)
 - [ProvidersSearchResponse](docs/ProvidersSearchResponse.md)
 - [RatingArea](docs/RatingArea.md)
 - [RequestPlanFind](docs/RequestPlanFind.md)
 - [RequestPlanFindApplicant](docs/RequestPlanFindApplicant.md)
 - [RequestPlanFindDrugPackage](docs/RequestPlanFindDrugPackage.md)
 - [RequestPlanFindProvider](docs/RequestPlanFindProvider.md)
 - [RequestProvidersSearch](docs/RequestProvidersSearch.md)
 - [ServiceArea](docs/ServiceArea.md)
 - [ServiceAreaZipCounty](docs/ServiceAreaZipCounty.md)
 - [State](docs/State.md)
 - [StateNetworkSizeRequest](docs/StateNetworkSizeRequest.md)
 - [StateNetworkSizeResponse](docs/StateNetworkSizeResponse.md)
 - [ZipCode](docs/ZipCode.md)
 - [ZipCountiesResponse](docs/ZipCountiesResponse.md)
 - [ZipCounty](docs/ZipCounty.md)
 - [ZipCountyBulk](docs/ZipCountyBulk.md)
 - [ZipCountyResponse](docs/ZipCountyResponse.md)


## Documentation For Authorization


## Vericred-Api-Key

- **Type**: API key 
- **API key parameter name**: Vericred-Api-Key
- **Location**: HTTP header


## Author



