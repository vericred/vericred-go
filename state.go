package vericredclient

import (
    "time"
)


type State struct {
    // Primary Key of State
    Id  int32  `json:"id,omitempty"`
    // Name of state
    Name  string  `json:"name,omitempty"`
    // 2 letter code for state
    Code  string  `json:"code,omitempty"`
    // National FIPs number
    FipsNumber  string  `json:"fips_number,omitempty"`
    // Last date this state is live for individuals
    LastDateForIndividual  time.Time  `json:"last_date_for_individual,omitempty"`
    // Last date this state is live for shop
    LastDateForShop  time.Time  `json:"last_date_for_shop,omitempty"`
    // Is this State available for businesses
    LiveForBusiness  bool  `json:"live_for_business,omitempty"`
    // Is this State available for individuals
    LiveForConsumers  bool  `json:"live_for_consumers,omitempty"`
}
