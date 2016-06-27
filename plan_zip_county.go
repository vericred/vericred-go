package vericredclient

import (
)


type PlanZipCounty struct {
    // Foreign key to plan
    PlanId  int32  `json:"plan_id,omitempty"`
    // Foreign key to county
    CountyId  int32  `json:"county_id,omitempty"`
    // Foreign key to zip code
    ZipCodeId  int32  `json:"zip_code_id,omitempty"`
}
