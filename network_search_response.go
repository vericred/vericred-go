package vericredclient

import (
)


type NetworkSearchResponse struct {
    // Metadata for query
    Meta  Meta  `json:"meta,omitempty"`
    // Networks that fit the requested criterion.
    Networks  []Network  `json:"networks,omitempty"`
}
