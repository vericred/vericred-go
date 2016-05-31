package vericredclient

import (
)


type ZipCounty struct {
    // ID of the parent County in Vericred's API
    CountyId  int32  `json:"county_id,omitempty"`
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // ID of the parent Zip Code in Vericred's API
    ZipCodeId  int32  `json:"zip_code_id,omitempty"`
}
