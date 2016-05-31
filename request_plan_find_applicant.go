package vericredclient

import (
)


type RequestPlanFindApplicant struct {
    // Age of applicant to search for
    Age  int32  `json:"age,omitempty"`
}
