package vericredclient

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type NetworksApi struct {
    Configuration Configuration
}

func NewNetworksApi() *NetworksApi{
    configuration := NewConfiguration()
    return &NetworksApi {
        Configuration: *configuration,
    }
}

func NewNetworksApiWithBasePath(basePath string) *NetworksApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &NetworksApi {
        Configuration: *configuration,
    }
}

/**
 * Networks
 * A network is a list of the doctors, other health care providers,
and hospitals that a plan has contracted with to provide medical care to
its members. This endpoint is paginated.
 * @param carrierId Carrier HIOS Issuer ID
 * @param page Page of paginated response
 * @param perPage Responses per page
 * @return NetworkSearchResponse
 */
func (a NetworksApi) ListNetworks (carrierId string, page int32, perPage int32) (NetworkSearchResponse, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/networks"

  // verify the required parameter 'carrierId' is set
  if &carrierId == nil {
      return *new(NetworkSearchResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'carrierId' when calling NetworksApi->ListNetworks")
  }

  headerParams := make(map[string]string)
  queryParams := make(map[string]string)
  formParams := make(map[string]string)
  var postBody interface{}
  var fileName string
  var fileBytes []byte

  // authentication (Vericred-Api-Key) required
  
  // set key with prefix in header
  headerParams["Vericred-Api-Key"] = a.Configuration.GetAPIKeyWithPrefix("Vericred-Api-Key")
      

  // add default headers if any
  for key := range a.Configuration.DefaultHeader {
      headerParams[key] = a.Configuration.DefaultHeader[key]
  }
  
  queryParams["carrierId"] = a.Configuration.APIClient.ParameterToString(carrierId)
  queryParams["page"] = a.Configuration.APIClient.ParameterToString(page)
  queryParams["perPage"] = a.Configuration.APIClient.ParameterToString(perPage)

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


  var successPayload = new(NetworkSearchResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
