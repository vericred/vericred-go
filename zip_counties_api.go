package vericredclient

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type ZipCountiesApi struct {
    Configuration Configuration
}

func NewZipCountiesApi() *ZipCountiesApi{
    configuration := NewConfiguration()
    return &ZipCountiesApi {
        Configuration: *configuration,
    }
}

func NewZipCountiesApiWithBasePath(basePath string) *ZipCountiesApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &ZipCountiesApi {
        Configuration: *configuration,
    }
}

/**
 * Search for Zip Counties
 * Our &#x60;Plan&#x60; endpoints require a zip code and a fips (county) code.  This is because plan pricing requires both of these elements.  Users are unlikely to know their fips code, so we provide this endpoint to look up a &#x60;ZipCounty&#x60; by zip code and return both the selected zip and fips codes.
 * @param zipPrefix Partial five-digit Zip
 * @param vericredApiKey API Key
 * @return ZipCountyResponse
 */
func (a ZipCountiesApi) GetZipCounties (zipPrefix string, vericredApiKey string) (ZipCountyResponse, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/zip_counties"

  // verify the required parameter 'zipPrefix' is set
  if &zipPrefix == nil {
      return *new(ZipCountyResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'zipPrefix' when calling ZipCountiesApi->GetZipCounties")
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
  
  queryParams["zipPrefix"] = a.Configuration.APIClient.ParameterToString(zipPrefix)

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


  var successPayload = new(ZipCountyResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
