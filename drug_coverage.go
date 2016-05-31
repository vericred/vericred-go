package vericredclient

import (
)


type DrugCoverage struct {
    // Health Insurance Oversight System id
    PlanId  string  `json:"plan_id,omitempty"`
    // NDC package code
    DrugPackageId  string  `json:"drug_package_id,omitempty"`
    // Tier Name
    Tier  string  `json:"tier,omitempty"`
    // Quantity limit exists
    QuantityLimit  bool  `json:"quantity_limit,omitempty"`
    // Prior authorization required
    PriorAuthorization  bool  `json:"prior_authorization,omitempty"`
    // Step Treatment required
    StepTherapy  bool  `json:"step_therapy,omitempty"`
}
