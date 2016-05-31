package vericredclient

import (
)


type PlanSearchResponse struct {
    // Metadata for query
    Meta  Meta  `json:"meta,omitempty"`
    // Plan search results
    Plans  []Plan  `json:"plans,omitempty"`
    // null
    Coverages  []Drug  `json:"coverages,omitempty"`
}
