package vericredclient

import (
)


type ZipCountyResponse struct {
    // Counties that exist in the provided zip prefix.
    Counties  []County  `json:"counties,omitempty"`
    // States that exist in the provided zip prefix.
    States  []State  `json:"states,omitempty"`
    // ZipCodes that exist in the provided zip prefix.
    ZipCodes  []ZipCode  `json:"zip_codes,omitempty"`
    // ZipCounty data
    ZipCounty  ZipCounty  `json:"zip_county,omitempty"`
}
