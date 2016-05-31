package vericredclient

import (
)


type DrugCoverageResponse struct {
    // Metadata for query
    Meta  Meta  `json:"meta,omitempty"`
    // DrugCoverage search results
    DrugCoverages  []DrugCoverage  `json:"drug_coverages,omitempty"`
    // Drug
    Drugs  []Drug  `json:"drugs,omitempty"`
    // Drug Packages
    DrugPackages  []DrugPackage  `json:"drug_packages,omitempty"`
}
