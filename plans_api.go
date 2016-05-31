package vericredclient

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type PlansApi struct {
    Configuration Configuration
}

func NewPlansApi() *PlansApi{
    configuration := NewConfiguration()
    return &PlansApi {
        Configuration: *configuration,
    }
}

func NewPlansApiWithBasePath(basePath string) *PlansApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &PlansApi {
        Configuration: *configuration,
    }
}

/**
 * Find Plans
 * ### Location Information

Searching for a set of plans requires a &#x60;zip_code&#x60; and &#x60;fips_code&#x60;
code.  These are used to determine pricing and availabity
of health plans.

Optionally, you may provide a list of Applicants or Providers

### Applicants

This is a list of people who will be covered by the plan.  We
use this list to calculate the premium.  You must include &#x60;age&#x60;
and can include &#x60;smoker&#x60;, which also factors into pricing in some
states.

Applicants *must* include an age.  If smoker is omitted, its value is assumed
to be false.

#### Multiple Applicants
To get pricing for multiple applicants, just append multiple sets
of data to the URL with the age and smoking status of each applicant
next to each other.

For example, given two applicants - one age 32 and a non-smoker and one
age 29 and a smoker, you could use the following request

&#x60;GET /plans?zip_code&#x3D;07451&amp;fips_code&#x3D;33025&amp;applicants[][age]&#x3D;32&amp;applicants[][age]&#x3D;29&amp;applicants[][smoker]&#x3D;true&#x60;

It would also be acceptible to include &#x60;applicants[][smoker]&#x3D;false&#x60; after the
first applicant&#39;s age.

### Providers

We identify Providers (Doctors) by their National Practitioner
Index number (NPI).  If you pass a list of Providers, keyed by
their NPI number, we will return a list of which Providers are
in and out of network for each plan returned.

For example, if we had two providers with the NPI numbers &#x60;12345&#x60; and &#x60;23456&#x60;
you would make the following request

&#x60;GET /plans?zip_code&#x3D;07451&amp;fips_code&#x3D;33025&amp;providers[][npi]&#x3D;12345&amp;providers[][npi]&#x3D;23456&#x60;

### Enrollment Date

To calculate plan pricing and availability, we default to the current date
as the enrollment date.  To specify a date in the future (or the past), pass
a string with the format &#x60;YYYY-MM-DD&#x60; in the &#x60;enrollment_date&#x60; parameter.

&#x60;GET /plans?zip_code&#x3D;07451&amp;fips_code&#x3D;33025&amp;enrollment_date&#x3D;2016-01-01&#x60;

### Subsidy

On-marketplace plans are eligible for a subsidy based on the
&#x60;household_size&#x60; and &#x60;household_income&#x60; of the applicants.  If you
pass those values, we will calculate the &#x60;subsidized_premium&#x60;
and return it for each plan.  If no values are provided, the
&#x60;subsidized_premium&#x60; will be the same as the &#x60;premium&#x60;

&#x60;GET /plans?zip_code&#x3D;07451&amp;fips_code&#x3D;33025&amp;household_size&#x3D;4&amp;household_income&#x3D;40000&#x60;

 * @param body 
 * @return PlanSearchResponse
 */
func (a PlansApi) FindPlans (body RequestPlanFind) (PlanSearchResponse, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/plans/search"


  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  
  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }

  // body params
  postBody = &body

  var successPayload = new(PlanSearchResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
