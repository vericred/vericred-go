package vericredclient

import (
)


type CarrierSubsidiary struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // Subsidiary name
    Name  string  `json:"name,omitempty"`
    // Parent Carrier Name
    AlternateName  string  `json:"alternate_name,omitempty"`
}
