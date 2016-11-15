package yext

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	CUSTOMERS_PATH string = "accounts"
)

var ResourceNotFound = errors.New("Resource not found")

type Client struct {
	Config *Config

	LocationService    *LocationService
	ECLService         *ECLService
	CustomFieldService *CustomFieldService
	FolderService      *FolderService
	CategoryService    *CategoryService
	LicenseService     *LicenseService
	UserService        *UserService
}

func NewClient(config *Config) *Client {
	c := &Client{Config: config}

	c.LocationService = &LocationService{client: c}
	c.ECLService = &ECLService{client: c}
	c.CustomFieldService = &CustomFieldService{client: c}
	c.FolderService = &FolderService{client: c}
	c.CategoryService = &CategoryService{client: c}
	c.LicenseService = &LicenseService{client: c}
	c.UserService = &UserService{client: c}

	return c
}

func (c *Client) customerRequestUrl(path string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.Config.BaseUrl, CUSTOMERS_PATH, c.Config.AccountId, path)
}

func (c *Client) rawRequestURL(path string) string {
	return fmt.Sprintf("%s/%s", c.Config.BaseUrl, path)
}

func (c *Client) NewRequest(method string, path string) (*http.Request, error) {
	return c.NewCustomerRequestBody(method, path, nil)
}

func (c *Client) NewRawRequest(method string, path string) (*http.Request, error) {
	return c.NewRawRequestBody(method, path, nil)
}

func (c *Client) NewRequestJSON(method string, path string, obj interface{}) (*http.Request, error) {
	json, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return c.NewCustomerRequestBody(method, path, json)
}

func (c *Client) NewRawRequestJSON(method string, path string, obj interface{}) (*http.Request, error) {
	json, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return c.NewRawRequestBody(method, path, json)
}

func (c *Client) NewRequestBody(method string, fullPath string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("api_key", c.Config.ApiKey)
	q.Add("v", c.Config.VParam)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *Client) DoRequest(method string, path string, v interface{}) error {
	req, err := c.NewRequest(method, path)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

func (c *Client) DoRawRequest(method string, path string, v interface{}) error {
	req, err := c.NewRawRequest(method, path)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

func (c *Client) NewCustomerRequestBody(method string, path string, data []byte) (*http.Request, error) {
	return c.NewRequestBody(method, c.customerRequestUrl(path), data)
}

func (c *Client) NewRawRequestBody(method string, path string, data []byte) (*http.Request, error) {
	return c.NewRequestBody(method, c.rawRequestURL(path), data)
}

func (c *Client) DoRequestJSON(method string, path string, obj interface{}, v interface{}) error {
	req, err := c.NewRequestJSON(method, path, obj)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

func (c *Client) DoRawRequestJSON(method string, path string, obj interface{}, v interface{}) error {
	req, err := c.NewRawRequestJSON(method, path, obj)
	if err != nil {
		return err
	}

	return c.Do(req, v)
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	// drain and cache the request body
	originalRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var (
		resultError error
		retries     = 3
	)

	if c.Config.RetryCount != nil {
		retries = *c.Config.RetryCount
	}

	for attempt := 0; attempt <= retries; attempt++ {
		resultError = nil
		time.Sleep(DefaultBackoffPolicy.Duration(attempt))

		// Rehydrate the request body since it might have been drained by the previous attempt
		req.Body = ioutil.NopCloser(bytes.NewBuffer(originalRequestBody))

		if c.Config.Logger != nil {
			c.Config.Logger.Log(fmt.Printf("%+v", req))
		}

		resp, err := c.Config.HTTPClient.Do(req)
		if err != nil {
			resultError = err
			continue
		}

		defer resp.Body.Close()

		if retryable, err := CheckResponseError(resp); err != nil {
			resultError = err
			if retryable {
				continue
			} else {
				break
			}
		}

		if v != nil {
			var fullResponse interface{}
			fullResponse = &Response{Body: v}
			if w, ok := fullResponse.(io.Writer); ok {
				io.Copy(w, resp.Body)
			} else {
				resultError = json.NewDecoder(resp.Body).Decode(fullResponse)
			}
		}

		if resultError == nil {
			return nil
		}
	}
	return resultError
}

func CheckResponseError(res *http.Response) (bool, error) {
	if sc := res.StatusCode; 200 <= sc && sc <= 299 {
		return true, nil
	} else if sc == 404 {
		return false, ResourceNotFound
	} else {
		retryable := !(400 <= sc && sc <= 499)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return retryable, err
		}

		// errorResponse := &Response{HTTPResponse: res}
		errorResponse := &Response{}
		if err := json.Unmarshal(data, errorResponse); err != nil {
			return retryable, errors.New(fmt.Sprintf("unable to unmarshal error from: %s : %s", string(data), err))
		}
		return retryable, errorResponse
	}
}
