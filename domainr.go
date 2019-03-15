// This library is meant to interact with the Client API – https://domainr.com/docs/api
// It currently only supports their free endpoint: https://domainr.p.mashape.com
// refer to https://domainr.com/docs/api/auth for authentication

package domainr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// default endpoint to the domainr / Rapid API.
	defaultEndpoint = "https://domainr.p.mashape.com"

	apiVersion = "v2"
)

type Client struct {
	Config
	Search
}
type Config struct {
	APIEndpoint string
	AuthKey     string
	ClientID    string
	APIVersion  string
}
type CfgOptions func(*Client)

func NewClient(opts ...CfgOptions) *Client {
	c := &Client{}
	// sets the default endpoint to the RapidAPI/mashape endpoint
	// You can override this by using the commercial endpoint func or using the SetCfg func.
	c.APIEndpoint = defaultEndpoint
	c.APIVersion = apiVersion
	for _, opt := range opts {
		opt(c)
	}
	c.Search = &searchService{c}
	return c
}

// RapidAPI Key
func AuthKey(key string) CfgOptions {
	return func(d *Client) {
		d.AuthKey = key
	}
}

// for commercial API
func ClientID(clientID string) CfgOptions {
	return func(d *Client) {
		d.ClientID = clientID
	}
}

// for commercial API
func CommercialEndpoint() CfgOptions {
	return func(d *Client) {
		d.APIEndpoint = "https://api.domainr.com"
	}
}

func (d *Client) SetCfg(config Config) *Client {
	d.Config = config
	return d
}

// NewRequest creates an API request.
// The path is expected to be a relative path and will be resolved
// according to the BaseURL of the Client. Paths should always be specified without a preceding slash.
func (c *Client) newRequest(method, path string, payload interface{}) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s/%s", c.APIEndpoint, c.APIVersion, path)

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}

func (c *Client) get(path string, obj interface{}) (*http.Response, error) {
	req, err := c.newRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

func (c *Client) post(path string, payload, obj interface{}) (*http.Response, error) {
	req, err := c.newRequest("POST", path, payload)
	if err != nil {
		return nil, err
	}

	return c.Do(req, obj)
}

// Do sends an API request and returns the API response.
//
// The API response is JSON decoded and stored in the value pointed by obj,
// or returned as an error if an API error has occurred.
// If obj implements the io.Writer interface, the raw response body will be written to obj,
// without attempting to decode it.
func (c *Client) Do(req *http.Request, obj interface{}) (*http.Response, error) {
	var r = http.DefaultClient
	resp, err := r.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// If obj implements the io.Writer,
	// the response body is decoded into v.
	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(obj)
		}
	}

	return resp, err
}
