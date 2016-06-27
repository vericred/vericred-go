package vericredclient

import (
)


type Plan struct {
    // Does the plan provide dental coverage for adults?
    AdultDental  bool  `json:"adult_dental,omitempty"`
    // 
    Age29Rider  bool  `json:"age29_rider,omitempty"`
    // Benefits string for ambulance coverage
    Ambulance  string  `json:"ambulance,omitempty"`
    // Link to the summary of benefits and coverage (SBC) document.
    BenefitsSummaryUrl  string  `json:"benefits_summary_url,omitempty"`
    // Link to a location to purchase the plan.
    BuyLink  string  `json:"buy_link,omitempty"`
    // Name of the insurance carrier
    CarrierName  string  `json:"carrier_name,omitempty"`
    // Does the plan provide dental coverage for children?
    ChildDental  bool  `json:"child_dental,omitempty"`
    // Child eyewear benefits summary
    ChildEyewear  string  `json:"child_eyewear,omitempty"`
    // Child eye exam benefits summary
    ChildEyeExam  string  `json:"child_eye_exam,omitempty"`
    // Phone number to contact the insurance carrier
    CustomerServicePhoneNumber  string  `json:"customer_service_phone_number,omitempty"`
    // Benefits summary for durable medical equipment
    DurableMedicalEquipment  string  `json:"durable_medical_equipment,omitempty"`
    // Diagnostic tests benefit summary
    DiagnosticTest  string  `json:"diagnostic_test,omitempty"`
    // Is this a domestic plan?
    DpRider  bool  `json:"dp_rider,omitempty"`
    // Link to the summary of drug benefits for the plan
    DrugFormularyUrl  string  `json:"drug_formulary_url,omitempty"`
    // Effective date of coverage.
    EffectiveDate  string  `json:"effective_date,omitempty"`
    // Expiration date of coverage.
    ExpirationDate  string  `json:"expiration_date,omitempty"`
    // Description of costs when visiting the ER
    EmergencyRoom  string  `json:"emergency_room,omitempty"`
    // Deductible for drugs when a family is on the plan.
    FamilyDrugDeductible  string  `json:"family_drug_deductible,omitempty"`
    // Maximum out-of-pocket for drugs when a family is on the plan
    FamilyDrugMoop  string  `json:"family_drug_moop,omitempty"`
    // Deductible when a family is on the plan
    FamilyMedicalDeductible  string  `json:"family_medical_deductible,omitempty"`
    // Maximum out-of-pocket when a family is on the plan
    FamilyMedicalMoop  string  `json:"family_medical_moop,omitempty"`
    // Is this a family plan?
    FpRider  bool  `json:"fp_rider,omitempty"`
    // Cost for generic drugs
    GenericDrugs  string  `json:"generic_drugs,omitempty"`
    // Habilitation services benefits summary
    HabilitationServices  string  `json:"habilitation_services,omitempty"`
    // 
    HiosIssuerId  string  `json:"hios_issuer_id,omitempty"`
    // Home health care benefits summary
    HomeHealthCare  string  `json:"home_health_care,omitempty"`
    // Hospice service benefits summary
    HospiceService  string  `json:"hospice_service,omitempty"`
    // Government-issued HIOS plan ID
    Id  string  `json:"id,omitempty"`
    // Benefits summary for imaging coverage
    Imaging  string  `json:"imaging,omitempty"`
    // List of NPI numbers for Providers passed in who accept this Plan
    InNetworkIds  []int32  `json:"in_network_ids,omitempty"`
    // Deductible for drugs when an individual is on the plan
    IndividualDrugDeductible  string  `json:"individual_drug_deductible,omitempty"`
    // Maximum out-of-pocket for drugs when an individual is on the plan
    IndividualDrugMoop  string  `json:"individual_drug_moop,omitempty"`
    // Deductible when an individual is on the plan
    IndividualMedicalDeductible  string  `json:"individual_medical_deductible,omitempty"`
    // Maximum out-of-pocket when an individual is on the plan
    IndividualMedicalMoop  string  `json:"individual_medical_moop,omitempty"`
    // Inpatient birth benefits summary
    InpatientBirth  string  `json:"inpatient_birth,omitempty"`
    // Cost under the plan for an inpatient facility
    InpatientFacility  string  `json:"inpatient_facility,omitempty"`
    // Inpatient mental helath benefits summary
    InpatientMentalHealth  string  `json:"inpatient_mental_health,omitempty"`
    // Cost under the plan for an inpatient physician
    InpatientPhysician  string  `json:"inpatient_physician,omitempty"`
    // Inpatient substance abuse benefits summary
    InpatientSubstance  string  `json:"inpatient_substance,omitempty"`
    // Plan metal grouping (e.g. platinum, gold, silver, etc)
    Level  string  `json:"level,omitempty"`
    // Link to a copy of the insurance carrier's logo
    LogoUrl  string  `json:"logo_url,omitempty"`
    // Marketing name of the plan
    Name  string  `json:"name,omitempty"`
    // Total number of Providers in network
    NetworkSize  int32  `json:"network_size,omitempty"`
    // Cost under the plan for non-preferred brand drugs
    NonPreferredBrandDrugs  string  `json:"non_preferred_brand_drugs,omitempty"`
    // Is the plan on-market?
    OnMarket  bool  `json:"on_market,omitempty"`
    // Is the plan off-market?
    OffMarket  bool  `json:"off_market,omitempty"`
    // Does this plan provide any out of network coverage?
    OutOfNetworkCoverage  bool  `json:"out_of_network_coverage,omitempty"`
    // List of NPI numbers for Providers passed in who do not accept this Plan
    OutOfNetworkIds  []int32  `json:"out_of_network_ids,omitempty"`
    // Benefits summary for outpatient facility coverage
    OutpatientFacility  string  `json:"outpatient_facility,omitempty"`
    // Benefits summary for outpatient mental health coverage
    OutpatientMentalHealth  string  `json:"outpatient_mental_health,omitempty"`
    // Benefits summary for outpatient physician coverage
    OutpatientPhysician  string  `json:"outpatient_physician,omitempty"`
    // Outpatient substance abuse benefits summary
    OutpatientSubstance  string  `json:"outpatient_substance,omitempty"`
    // Market in which the plan is offered (on_marketplace, shop, etc)
    PlanMarket  string  `json:"plan_market,omitempty"`
    // Category of the plan (e.g. EPO, HMO, PPO, POS, Indemnity)
    PlanType  string  `json:"plan_type,omitempty"`
    // Cost under the plan for perferred brand drugs
    PreferredBrandDrugs  string  `json:"preferred_brand_drugs,omitempty"`
    // Inpatient substance abuse benefits summary
    PrenatalPostnatalCare  string  `json:"prenatal_postnatal_care,omitempty"`
    // Benefits summary for preventative care
    PreventativeCare  string  `json:"preventative_care,omitempty"`
    // Cumulative premium amount after subsidy
    PremiumSubsidized  float32  `json:"premium_subsidized,omitempty"`
    // Cumulative premium amount
    Premium  float32  `json:"premium,omitempty"`
    // Cost under the plan to visit a Primary Care Physician
    PrimaryCarePhysician  string  `json:"primary_care_physician,omitempty"`
    // Benefits summary for rehabilitation services
    RehabilitationServices  string  `json:"rehabilitation_services,omitempty"`
    // Benefits summary for skilled nursing services
    SkilledNursing  string  `json:"skilled_nursing,omitempty"`
    // Cost under the plan to visit a specialist
    Specialist  string  `json:"specialist,omitempty"`
    // Cost under the plan for specialty drugs
    SpecialtyDrugs  string  `json:"specialty_drugs,omitempty"`
    // Benefits summary for urgent care
    UrgentCare  string  `json:"urgent_care,omitempty"`
}
