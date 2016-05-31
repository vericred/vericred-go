package vericredclient

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type DrugsApi struct {
    Configuration Configuration
}

func NewDrugsApi() *DrugsApi{
    configuration := NewConfiguration()
    return &DrugsApi {
        Configuration: *configuration,
    }
}

func NewDrugsApiWithBasePath(basePath string) *DrugsApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &DrugsApi {
        Configuration: *configuration,
    }
}

/**
 * Search for DrugCoverages
 * Drug Coverages are the specific tier level, quantity limit, prior
authorization and step therapy for a given Drug/Plan combination. This endpoint
returns all DrugCoverages for a given Drug
 * @param ndcPackageCode NDC package code
 * @param audience Two-character state code
 * @param stateCode Two-character state code
 * @param vericredApiKey API Key
 * @return DrugCoverageResponse
 */
func (a DrugsApi) GetDrugCoverages (ndcPackageCode string, audience string, stateCode string, vericredApiKey string) (DrugCoverageResponse, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/drug_packages/{ndc_package_code}/coverages"
 path = strings.Replace(path, "{" + "ndc_package_code" + "}", fmt.Sprintf("%v", ndcPackageCode), -1)

  // verify the required parameter 'ndcPackageCode' is set
  if &ndcPackageCode == nil {
      return *new(DrugCoverageResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'ndcPackageCode' when calling DrugsApi->GetDrugCoverages")
  }
  // verify the required parameter 'audience' is set
  if &audience == nil {
      return *new(DrugCoverageResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'audience' when calling DrugsApi->GetDrugCoverages")
  }
  // verify the required parameter 'stateCode' is set
  if &stateCode == nil {
      return *new(DrugCoverageResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'stateCode' when calling DrugsApi->GetDrugCoverages")
  }

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
  
  queryParams["audience"] = a.Configuration.APIClient.ParameterToString(audience)
  queryParams["stateCode"] = a.Configuration.APIClient.ParameterToString(stateCode)

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
    // header params "Vericred-Api-Key"
    headerParams["Vericred-Api-Key"] = vericredApiKey


  var successPayload = new(DrugCoverageResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Drug Search
 * Search for drugs by proprietary name
 * @param searchTerm Full or partial proprietary name of drug
 * @param vericredApiKey API Key
 * @return DrugSearchResponse
 */
func (a DrugsApi) ListDrugs (searchTerm string, vericredApiKey string) (DrugSearchResponse, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/drugs"

  // verify the required parameter 'searchTerm' is set
  if &searchTerm == nil {
      return *new(DrugSearchResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'searchTerm' when calling DrugsApi->ListDrugs")
  }

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
  
  queryParams["searchTerm"] = a.Configuration.APIClient.ParameterToString(searchTerm)

  // to determine the Content-Type header
  localVarHttpContentTypes := []string {
      "application/json", 
  }
  //set Content-Type header
  localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
  if localVarHttpContentType != "" {    
      headerParams["Content-Type"] = localVarHttpContentType
  }
  // to determine the Accept header
  localVarHttpHeaderAccepts := []string {
      "application/json", 
  }
  //set Accept header
  localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
  if localVarHttpHeaderAccept != "" {  
      headerParams["Accept"] = localVarHttpHeaderAccept
  }
    // header params "Vericred-Api-Key"
    headerParams["Vericred-Api-Key"] = vericredApiKey


  var successPayload = new(DrugSearchResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
