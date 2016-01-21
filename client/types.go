package client

import (
	"github.com/bitly/go-simplejson"
)

type ClientOpts struct {
	ApiKey       string
	Organization string
}

type RequestOpts struct {
	ApiKey       string
	ClusterId    string
	Method       string
	Organization string
	Path         string
	QueryString  map[string]string
	Data         map[string]interface{}
}

type ResponseOpts struct {
	Body       simplejson.Json
	StatusCode int
}
