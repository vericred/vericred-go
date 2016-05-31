package vericredclient

import (
)


type ZipCountiesResponse struct {
    // Counties that exist in the provided zip prefix.
    Counties  []County  `json:"counties,omitempty"`
    // States that exist in the provided zip prefix.
    States  []State  `json:"states,omitempty"`
    // ZipCounties that exist in the provided zip prefix.
    ZipCounties  []ZipCounty  `json:"zip_counties,omitempty"`
    // ZipCodes that exist in the provided zip prefix.
    ZipCodes  []ZipCode  `json:"zip_codes,omitempty"`
}
