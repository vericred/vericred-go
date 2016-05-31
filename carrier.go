package vericredclient

import (
)


type Carrier struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // Name of the Carrier
    Name  string  `json:"name,omitempty"`
    // URL for the Carrier's logo
    LogoPath  string  `json:"logo_path,omitempty"`
}
