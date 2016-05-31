package vericredclient

import (
    "time"
)


type Applicant struct {
    // Primary key
    Id  int32  `json:"id,omitempty"`
    // Date of Birth
    Dob  time.Time  `json:"dob,omitempty"`
    // Member token
    MemberId  string  `json:"member_id,omitempty"`
    // Full name of the Applicant
    Name  string  `json:"name,omitempty"`
    // Relationship of the Applicant to the Member
    Relationship  string  `json:"relationship,omitempty"`
    // Does the Applicant smoke?
    Smoker  bool  `json:"smoker,omitempty"`
    // Applicant's Social Security Number
    Ssn  string  `json:"ssn,omitempty"`
}
