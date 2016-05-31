package vericredclient

import (
)


type Network struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // Carrier name
    Name  string  `json:"name,omitempty"`
}
