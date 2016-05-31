package vericredclient

import (
)


type RequestPlanFindProvider struct {
    // NPI of provider to search for
    Npi  int32  `json:"npi,omitempty"`
}
