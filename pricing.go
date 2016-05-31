package vericredclient

import (
    "time"
)


type Pricing struct {
    // Age of applicant
    Age  int32  `json:"age,omitempty"`
    // Effective date of plan
    EffectiveDate  time.Time  `json:"effective_date,omitempty"`
    // Plan expiration date
    ExpirationDate  time.Time  `json:"expiration_date,omitempty"`
    // Foreign key to plans
    PlanId  int32  `json:"plan_id,omitempty"`
    // Child-only premium
    PremiumChildOnly  float32  `json:"premium_child_only,omitempty"`
    // Family premium
    PremiumFamily  float32  `json:"premium_family,omitempty"`
    // Single-person premium
    PremiumSingle  float32  `json:"premium_single,omitempty"`
    // Single person including children premium
    PremiumSingleAndChildren  float32  `json:"premium_single_and_children,omitempty"`
    // Person with spouse premium
    PremiumSingleAndSpouse  float32  `json:"premium_single_and_spouse,omitempty"`
    // Premium for single smoker
    PremiumSingleSmoker  float32  `json:"premium_single_smoker,omitempty"`
    // Foreign key to rating areas
    RatingAreaId  string  `json:"rating_area_id,omitempty"`
}
