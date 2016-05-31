package vericredclient

import (
)


type ZipCountyBulk struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // Foreign key for rating area
    RatingAreaId  string  `json:"rating_area_id,omitempty"`
    // Foreign key for county (fips code)
    CountyId  string  `json:"county_id,omitempty"`
    // Foreign key for zip code (zip code)
    ZipCodeId  string  `json:"zip_code_id,omitempty"`
}
