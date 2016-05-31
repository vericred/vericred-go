package vericredclient

import (
)


type County struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // State FIPS code
    FipsCode  string  `json:"fips_code,omitempty"`
    // Human-readable name
    Name  string  `json:"name,omitempty"`
    // Two-character state code
    StateCode  string  `json:"state_code,omitempty"`
    // state relationship
    StateId  int32  `json:"state_id,omitempty"`
    // Is the state containing this county active for consumers?(deprecated in favor of last_date_for_individual)
    StateLive  bool  `json:"state_live,omitempty"`
    // Is the state containing this county active for business?(deprecated in favor of last_date_for_shop)
    StateLiveForBusiness  bool  `json:"state_live_for_business,omitempty"`
}
