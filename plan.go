/* 
 * Vericred API
 *
 * Vericred's API allows you to search for Health Plans that a specific doctor
accepts.

## Getting Started

Visit our [Developer Portal](https://vericred.3scale.net) to
create an account.

Once you have created an account, you can create one Application for
Production and another for our Sandbox (select the appropriate Plan when
you create the Application).

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


 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vericredclient

type Plan struct {

	// Does the plan provide dental coverage for adults?
	AdultDental bool `json:"adult_dental,omitempty"`

	// True if the plan allows dependents up to age 29
	Age29Rider bool `json:"age29_rider,omitempty"`

	// Benefits string for ambulance coverage
	Ambulance string `json:"ambulance,omitempty"`

	// Link to the summary of benefits and coverage (SBC) document.
	BenefitsSummaryUrl string `json:"benefits_summary_url,omitempty"`

	// Link to a location to purchase the plan.
	BuyLink string `json:"buy_link,omitempty"`

	// Name of the insurance carrier
	CarrierName string `json:"carrier_name,omitempty"`

	// Does the plan provide dental coverage for children?
	ChildDental bool `json:"child_dental,omitempty"`

	// Child eyewear benefits summary
	ChildEyewear string `json:"child_eyewear,omitempty"`

	// Child eye exam benefits summary
	ChildEyeExam string `json:"child_eye_exam,omitempty"`

	// Phone number to contact the insurance carrier
	CustomerServicePhoneNumber string `json:"customer_service_phone_number,omitempty"`

	// Benefits summary for durable medical equipment
	DurableMedicalEquipment string `json:"durable_medical_equipment,omitempty"`

	// Diagnostic tests benefit summary
	DiagnosticTest string `json:"diagnostic_test,omitempty"`

	// Alternate name for the Plan
	DisplayName string `json:"display_name,omitempty"`

	// True if plan does not cover domestic partners
	DpRider bool `json:"dp_rider,omitempty"`

	// Link to the summary of drug benefits for the plan
	DrugFormularyUrl string `json:"drug_formulary_url,omitempty"`

	// Effective date of coverage.
	EffectiveDate string `json:"effective_date,omitempty"`

	// Expiration date of coverage.
	ExpirationDate string `json:"expiration_date,omitempty"`

	// Description of costs when visiting the ER
	EmergencyRoom string `json:"emergency_room,omitempty"`

	// Deductible for drugs when a family is on the plan.
	FamilyDrugDeductible string `json:"family_drug_deductible,omitempty"`

	// Maximum out-of-pocket for drugs when a family is on the plan
	FamilyDrugMoop string `json:"family_drug_moop,omitempty"`

	// Deductible when a family is on the plan
	FamilyMedicalDeductible string `json:"family_medical_deductible,omitempty"`

	// Maximum out-of-pocket when a family is on the plan
	FamilyMedicalMoop string `json:"family_medical_moop,omitempty"`

	// True if plan does not cover family planning
	FpRider bool `json:"fp_rider,omitempty"`

	// Cost for generic drugs
	GenericDrugs string `json:"generic_drugs,omitempty"`

	// Habilitation services benefits summary
	HabilitationServices string `json:"habilitation_services,omitempty"`

	// 
	HiosIssuerId string `json:"hios_issuer_id,omitempty"`

	// Home health care benefits summary
	HomeHealthCare string `json:"home_health_care,omitempty"`

	// Hospice service benefits summary
	HospiceService string `json:"hospice_service,omitempty"`

	// Is the plan HSA eligible?
	HsaEligible bool `json:"hsa_eligible,omitempty"`

	// Government-issued HIOS plan ID
	Id string `json:"id,omitempty"`

	// Benefits summary for imaging coverage
	Imaging string `json:"imaging,omitempty"`

	// List of NPI numbers for Providers passed in who accept this Plan
	InNetworkIds []int32 `json:"in_network_ids,omitempty"`

	// Deductible for drugs when an individual is on the plan
	IndividualDrugDeductible string `json:"individual_drug_deductible,omitempty"`

	// Maximum out-of-pocket for drugs when an individual is on the plan
	IndividualDrugMoop string `json:"individual_drug_moop,omitempty"`

	// Deductible when an individual is on the plan
	IndividualMedicalDeductible string `json:"individual_medical_deductible,omitempty"`

	// Maximum out-of-pocket when an individual is on the plan
	IndividualMedicalMoop string `json:"individual_medical_moop,omitempty"`

	// Inpatient birth benefits summary
	InpatientBirth string `json:"inpatient_birth,omitempty"`

	// Cost under the plan for an inpatient facility
	InpatientFacility string `json:"inpatient_facility,omitempty"`

	// Inpatient mental helath benefits summary
	InpatientMentalHealth string `json:"inpatient_mental_health,omitempty"`

	// Cost under the plan for an inpatient physician
	InpatientPhysician string `json:"inpatient_physician,omitempty"`

	// Inpatient substance abuse benefits summary
	InpatientSubstance string `json:"inpatient_substance,omitempty"`

	// Plan metal grouping (e.g. platinum, gold, silver, etc)
	Level string `json:"level,omitempty"`

	// Link to a copy of the insurance carrier's logo
	LogoUrl string `json:"logo_url,omitempty"`

	// Marketing name of the plan
	Name string `json:"name,omitempty"`

	// Total number of Providers in network
	NetworkSize int32 `json:"network_size,omitempty"`

	// Cost under the plan for non-preferred brand drugs
	NonPreferredBrandDrugs string `json:"non_preferred_brand_drugs,omitempty"`

	// Is the plan on-market?
	OnMarket bool `json:"on_market,omitempty"`

	// Is the plan off-market?
	OffMarket bool `json:"off_market,omitempty"`

	// Does this plan provide any out of network coverage?
	OutOfNetworkCoverage bool `json:"out_of_network_coverage,omitempty"`

	// List of NPI numbers for Providers passed in who do not accept this Plan
	OutOfNetworkIds []int32 `json:"out_of_network_ids,omitempty"`

	// Benefits summary for outpatient facility coverage
	OutpatientFacility string `json:"outpatient_facility,omitempty"`

	// Benefits summary for outpatient mental health coverage
	OutpatientMentalHealth string `json:"outpatient_mental_health,omitempty"`

	// Benefits summary for outpatient physician coverage
	OutpatientPhysician string `json:"outpatient_physician,omitempty"`

	// Outpatient substance abuse benefits summary
	OutpatientSubstance string `json:"outpatient_substance,omitempty"`

	// Market in which the plan is offered (on_marketplace, shop, etc)
	PlanMarket string `json:"plan_market,omitempty"`

	// Category of the plan (e.g. EPO, HMO, PPO, POS, Indemnity)
	PlanType string `json:"plan_type,omitempty"`

	// Cost under the plan for perferred brand drugs
	PreferredBrandDrugs string `json:"preferred_brand_drugs,omitempty"`

	// Inpatient substance abuse benefits summary
	PrenatalPostnatalCare string `json:"prenatal_postnatal_care,omitempty"`

	// Benefits summary for preventative care
	PreventativeCare string `json:"preventative_care,omitempty"`

	// Cumulative premium amount after subsidy
	PremiumSubsidized float32 `json:"premium_subsidized,omitempty"`

	// Cumulative premium amount
	Premium float32 `json:"premium,omitempty"`

	// Source of the base pricing data
	PremiumSource string `json:"premium_source,omitempty"`

	// Cost under the plan to visit a Primary Care Physician
	PrimaryCarePhysician string `json:"primary_care_physician,omitempty"`

	// Benefits summary for rehabilitation services
	RehabilitationServices string `json:"rehabilitation_services,omitempty"`

	// Foreign key for service area
	ServiceAreaId string `json:"service_area_id,omitempty"`

	// Benefits summary for skilled nursing services
	SkilledNursing string `json:"skilled_nursing,omitempty"`

	// Cost under the plan to visit a specialist
	Specialist string `json:"specialist,omitempty"`

	// Cost under the plan for specialty drugs
	SpecialtyDrugs string `json:"specialty_drugs,omitempty"`

	// Benefits summary for urgent care
	UrgentCare string `json:"urgent_care,omitempty"`
}
