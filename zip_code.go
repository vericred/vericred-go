package vericredclient

import (
)


type ZipCode struct {
    // 5 digit code (e.g. 11215)
    Code  string  `json:"code,omitempty"`
    // Primary key
    Id  int32  `json:"id,omitempty"`
}
