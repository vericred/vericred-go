package vericredclient

import (
    "strings"
    "fmt"
    "errors"
    "encoding/json"
)

type ProvidersApi struct {
    Configuration Configuration
}

func NewProvidersApi() *ProvidersApi{
    configuration := NewConfiguration()
    return &ProvidersApi {
        Configuration: *configuration,
    }
}

func NewProvidersApiWithBasePath(basePath string) *ProvidersApi{
    configuration := NewConfiguration()
    configuration.BasePath = basePath
    
    return &ProvidersApi {
        Configuration: *configuration,
    }
}

/**
 * Find a Provider
 * To retrieve a specific provider, just perform a GET using his NPI number
 * @param npi NPI number
 * @param vericredApiKey API Key
 * @return ProviderShowResponse
 */
func (a ProvidersApi) GetProvider (npi string, vericredApiKey string) (ProviderShowResponse, APIResponse, error) {

  var httpMethod = "Get"
 // create path and map variables
  path := a.Configuration.BasePath + "/providers/{npi}"
 path = strings.Replace(path, "{" + "npi" + "}", fmt.Sprintf("%v", npi), -1)

  // verify the required parameter 'npi' is set
  if &npi == nil {
      return *new(ProviderShowResponse), *NewAPIResponseWithError("400 - Bad Request"), errors.New("Missing required parameter 'npi' when calling ProvidersApi->GetProvider")
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


  var successPayload = new(ProviderShowResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
/**
 * Find Providers
 * All &#x60;Provider&#x60; searches require a &#x60;zip_code&#x60;, which we use for weighting
the search results to favor &#x60;Provider&#x60;s that are near the user.  For example,
we would want &quot;Dr. John Smith&quot; who is 5 miles away to appear before
&quot;Dr. John Smith&quot; who is 100 miles away.

The weighting also allows for non-exact matches.  In our prior example, we
would want &quot;Dr. Jon Smith&quot; who is 2 miles away to appear before the exact
match &quot;Dr. John Smith&quot; who is 100 miles away because it is more likely that
the user just entered an incorrect name.

The free text search also supports Specialty name search and &quot;body part&quot;
Specialty name search.  So, searching &quot;John Smith nose&quot; would return
&quot;Dr. John Smith&quot;, the ENT Specialist before &quot;Dr. John Smith&quot; the Internist.

 * @param body 
 * @return ProvidersSearchResponse
 */
func (a ProvidersApi) GetProviders (body RequestProvidersSearch) (ProvidersSearchResponse, APIResponse, error) {

  var httpMethod = "Post"
 // create path and map variables
  path := a.Configuration.BasePath + "/providers/search"


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

  var successPayload = new(ProvidersSearchResponse)
  httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)


  if err != nil {
    return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
  }

  err = json.Unmarshal(httpResponse.Body(), &successPayload)

  return *successPayload, *NewAPIResponse(httpResponse.RawResponse), err
}
