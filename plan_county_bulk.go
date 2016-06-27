package vericredclient

import (
)


type PlanCountyBulk struct {
    // Foreign key to plan
    PlanId  int32  `json:"plan_id,omitempty"`
    // Foreign key to county
    CountyId  int32  `json:"county_id,omitempty"`
}
