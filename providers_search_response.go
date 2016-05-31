package vericredclient

import (
)


type ProvidersSearchResponse struct {
    // Metadata for query
    Meta  Meta  `json:"meta,omitempty"`
    // Providers that fit the requested criterion.
    Providers  []Provider  `json:"providers,omitempty"`
    // States that fit the requested criterion.
    States  []State  `json:"states,omitempty"`
}
