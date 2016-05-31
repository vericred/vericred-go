package vericredclient

import (
)


type CountyBulk struct {
    // FIPs code for the county
    Id  string  `json:"id,omitempty"`
    // Name of the county
    Name  string  `json:"name,omitempty"`
    // State code
    StateId  string  `json:"state_id,omitempty"`
}
