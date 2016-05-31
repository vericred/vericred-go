package vericredclient

import (
)


type RequestProvidersSearch struct {
    // Limit results to Providers who accept at least one insurance
        plan.  Note that the inverse of this filter is not supported and
        any value will evaluate to true
    AcceptsInsurance  bool  `json:"accepts_insurance,omitempty"`
    // List of HIOS ids
    HiosIds  []string  `json:"hios_ids,omitempty"`
    // Page number
    Page  int32  `json:"page,omitempty"`
    // Number of records to return per page
    PerPage  int32  `json:"per_page,omitempty"`
    // Radius (in miles) to use to limit results
    Radius  int32  `json:"radius,omitempty"`
    // String to search by
    SearchTerm  string  `json:"search_term,omitempty"`
    // Zip Code to search near
    ZipCode  string  `json:"zip_code,omitempty"`
    // Either organization or individual
    Type_  string  `json:"type,omitempty"`
}
