package vericredclient

import (
)


type Drug struct {
    // National Drug Code ID
    Id  string  `json:"id,omitempty"`
    // Proprietary name of drug
    ProprietaryName  string  `json:"proprietary_name,omitempty"`
    // Non-proprietary name of drug
    NonProprietaryName  string  `json:"non_proprietary_name,omitempty"`
    // Array of drug package Ids
    DrugPackageIds  []string  `json:"drug_package_ids,omitempty"`
}
