package client

import (
	"bytes"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
)

const (
	BASE_API_URL         = "https://api.containership.io"
	BASE_API_VERSION     = "v1"
	CONTENT_TYPE_HEADER  = "Content-Type"
	ORGANIZATION_HEADER  = "X-ContainerShip-Cloud-Organization"
	API_KEY_HEADER       = "X-ContainerShip-Cloud-API-Key"
	DEFAULT_CONTENT_TYPE = "application/json"
	DEFAULT_METHOD       = "POST"
)

func ApiRequest(requestOpts *RequestOpts) (*ResponseOpts, error) {
	httpsClient := &http.Client{}

	url := fmt.Sprintf("%s/%s/%s/clusters/%s/proxy", BASE_API_URL, BASE_API_VERSION, requestOpts.Organization, requestOpts.ClusterId)

	jsonBody := simplejson.New()
	jsonBody.Set("method", requestOpts.Method)
	jsonBody.Set("url", requestOpts.Path)

	if requestOpts.Method == "POST" || requestOpts.Method == "PUT" {
		jsonBody.Set("data", requestOpts.Data)
	}

	if requestOpts.QueryString != nil {
		for k, v := range requestOpts.QueryString {
			jsonBody.SetPath([]string{"qs", k}, v)
		}
	}

	jsonBodyBytes, err := jsonBody.MarshalJSON()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(DEFAULT_METHOD, url, bytes.NewBuffer(jsonBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Add(ORGANIZATION_HEADER, requestOpts.Organization)
	req.Header.Add(API_KEY_HEADER, requestOpts.ApiKey)
	req.Header.Add(CONTENT_TYPE_HEADER, DEFAULT_CONTENT_TYPE)

	resp, err := httpsClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responseOpts := &ResponseOpts{
		StatusCode: resp.StatusCode,
	}

	json, err := simplejson.NewJson(body)
	if err == nil {
		responseOpts.Body = *json
	}

	return responseOpts, nil
}
