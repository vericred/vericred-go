package vericredclient

import (
)


type Provider struct {
    // Is this provider accepting patients with a change of insurance?
    AcceptingChangeOfPayorPatients  bool  `json:"accepting_change_of_payor_patients,omitempty"`
    // Is this provider accepting new Medicaid patients?
    AcceptingMedicaidPatients  bool  `json:"accepting_medicaid_patients,omitempty"`
    // Is this provider accepting new Medicare patients?
    AcceptingMedicarePatients  bool  `json:"accepting_medicare_patients,omitempty"`
    // Is this provider accepting new patients with private insurance?
    AcceptingPrivatePatients  bool  `json:"accepting_private_patients,omitempty"`
    // Is this provider accepting new patients via referrals?
    AcceptingReferralPatients  bool  `json:"accepting_referral_patients,omitempty"`
    // City name (e.g. Springfield).
    City  string  `json:"city,omitempty"`
    // Primary email address to contact the provider.
    Email  string  `json:"email,omitempty"`
    // Provider's gender (M or F)
    Gender  string  `json:"gender,omitempty"`
    // Given name for the provider.
    FirstName  string  `json:"first_name,omitempty"`
    // List of HIOS ids for this provider
    HiosIds  []string  `json:"hios_ids,omitempty"`
    // National Provider Index (NPI) number
    Id  int32  `json:"id,omitempty"`
    // Family name for the provider.
    LastName  string  `json:"last_name,omitempty"`
    // Latitude of provider
    Latitude  float32  `json:"latitude,omitempty"`
    // Longitude of provider
    Longitude  float32  `json:"longitude,omitempty"`
    // Middle name for the provider.
    MiddleName  string  `json:"middle_name,omitempty"`
    // Array of network ids
    NetworkIds  []int32  `json:"network_ids,omitempty"`
    // Personal contact phone for the provider.
    PersonalPhone  string  `json:"personal_phone,omitempty"`
    // Office phone for the provider
    Phone  string  `json:"phone,omitempty"`
    // Preferred name for display (e.g. Dr. Francis White may prefer Dr. Frank White)
    PresentationName  string  `json:"presentation_name,omitempty"`
    // Name of the primary Specialty
    Specialty  string  `json:"specialty,omitempty"`
    // State code for the provider's address (e.g. NY).
    State  string  `json:"state,omitempty"`
    // Foreign key to States
    StateId  int32  `json:"state_id,omitempty"`
    // First line of the provider's street address.
    StreetLine1  string  `json:"street_line_1,omitempty"`
    // Second line of the provider's street address.
    StreetLine2  string  `json:"street_line_2,omitempty"`
    // Suffix for the provider's name (e.g. Jr)
    Suffix  string  `json:"suffix,omitempty"`
    // Professional title for the provider (e.g. Dr).
    Title  string  `json:"title,omitempty"`
    // Type of NPI number (individual provider vs organization).
    Type_  string  `json:"type,omitempty"`
    // Postal code for the provider's address (e.g. 11215)
    ZipCode  string  `json:"zip_code,omitempty"`
}
