package vericredclient

import (
)


type DrugSearchResponse struct {
    // Metadata for query
    Meta  Meta  `json:"meta,omitempty"`
    // Drugs found in query
    Drugs  []Drug  `json:"drugs,omitempty"`
    // DrugPackages
    DrugPackages  []DrugPackage  `json:"drug_packages,omitempty"`
}
