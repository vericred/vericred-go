package vericredclient

import (
)


type DrugPackage struct {
    // National Drug Code ID (Package)
    Id  string  `json:"id,omitempty"`
    // Package description
    Description  string  `json:"description,omitempty"`
}
