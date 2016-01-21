package client

import (
	"fmt"
)

func (c *ContainerShipCloudClient) GetApplication(clusterId string, appName string) (*ResponseOpts, error) {
	path := fmt.Sprintf("/applications/%s", appName)

	requestOpts := &RequestOpts{
		ApiKey:       c.Opts.ApiKey,
		ClusterId:    clusterId,
		Method:       "GET",
		Organization: c.Opts.Organization,
		Path:         path,
		QueryString:  nil,
		Data:         nil,
	}

	return ApiRequest(requestOpts)
}

func (c *ContainerShipCloudClient) CreateApplication(clusterId string, appName string, application map[string]interface{}) (*ResponseOpts, error) {
	path := fmt.Sprintf("/applications/%s", appName)

	requestOpts := &RequestOpts{
		ApiKey:       c.Opts.ApiKey,
		ClusterId:    clusterId,
		Method:       "POST",
		Organization: c.Opts.Organization,
		Path:         path,
		QueryString:  nil,
		Data:         application,
	}

	return ApiRequest(requestOpts)
}

func (c *ContainerShipCloudClient) UpdateApplication(clusterId string, appName string, application map[string]interface{}) (*ResponseOpts, error) {
	path := fmt.Sprintf("/applications/%s", appName)

	requestOpts := &RequestOpts{
		ApiKey:       c.Opts.ApiKey,
		ClusterId:    clusterId,
		Method:       "PUT",
		Organization: c.Opts.Organization,
		Path:         path,
		QueryString:  nil,
		Data:         application,
	}

	return ApiRequest(requestOpts)
}

func (c *ContainerShipCloudClient) DeleteApplication(clusterId string, appName string) (*ResponseOpts, error) {
	path := fmt.Sprintf("/applications/%s", appName)

	requestOpts := &RequestOpts{
		ApiKey:       c.Opts.ApiKey,
		ClusterId:    clusterId,
		Method:       "DELETE",
		Organization: c.Opts.Organization,
		Path:         path,
		QueryString:  nil,
		Data:         nil,
	}

	return ApiRequest(requestOpts)
}
