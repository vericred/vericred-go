package vericredclient

import (
)


type RequestPlanFind struct {
    // Applicants for desired plans.
    Applicants  []RequestPlanFindApplicant  `json:"applicants,omitempty"`
    // Date of enrollment
    EnrollmentDate  string  `json:"enrollment_date,omitempty"`
    // National Drug Code Package Id
    DrugPackages  []DrugPackage  `json:"drug_packages,omitempty"`
    // County code to determine eligibility
    FipsCode  string  `json:"fips_code,omitempty"`
    // Total household income.
    HouseholdIncome  int32  `json:"household_income,omitempty"`
    // Number of people living in household.
    HouseholdSize  int32  `json:"household_size,omitempty"`
    // Type of plan to search for.
    Market  string  `json:"market,omitempty"`
    // List of providers to search for.
    Providers  []RequestPlanFindProvider  `json:"providers,omitempty"`
    // 5-digit zip code - this helps determine pricing.
    ZipCode  string  `json:"zip_code,omitempty"`
}
