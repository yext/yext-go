package yext

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	ACCOUNTS_PATH string = "accounts"
)

type ListOptions struct {
	Limit  int
	Offset int
}

type Client struct {
	Config *Config

	LocationService    *LocationService
	ListService        *ListService
	CustomFieldService *CustomFieldService
	FolderService      *FolderService
	CategoryService    *CategoryService
	UserService        *UserService
}

func NewClient(config *Config) *Client {
	c := &Client{Config: config}

	c.LocationService = &LocationService{client: c}
	c.ListService = &ListService{client: c}
	c.CustomFieldService = &CustomFieldService{client: c}
	c.FolderService = &FolderService{client: c}
	c.CategoryService = &CategoryService{client: c}
	c.UserService = &UserService{client: c}

	return c
}

// TODO: The account (e.g. /v2/account/me/locations) vs raw (e.g. /v2/categories)
// URL distiction is present in the API.  We currently have the NewXXX and DoXXX
// helpers split as well, but that probably isn't necessary.  We could have the
// services do their own URL path construction (possibly with some helpers in
// this file)

func (c *Client) accountRequestUrl(path string) string {
	return fmt.Sprintf("%s/%s/%s/%s", c.Config.BaseUrl, ACCOUNTS_PATH, c.Config.AccountId, path)
}

func (c *Client) rawRequestURL(path string) string {
	return fmt.Sprintf("%s/%s", c.Config.BaseUrl, path)
}

func (c *Client) NewRequest(method string, path string) (*http.Request, error) {
	return c.NewAccountRequestBody(method, path, nil)
}

func (c *Client) NewRootRequest(method string, path string) (*http.Request, error) {
	return c.NewRootRequestBody(method, path, nil)
}

func (c *Client) NewRequestJSON(method string, path string, obj interface{}) (*http.Request, error) {
	json, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return c.NewAccountRequestBody(method, path, json)
}

func (c *Client) NewRootRequestJSON(method string, path string, obj interface{}) (*http.Request, error) {
	json, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return c.NewRootRequestBody(method, path, json)
}

func (c *Client) NewRequestBody(method string, fullPath string, data []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, fullPath, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("api_key", c.Config.ApiKey)
	q.Add("v", c.Config.Version)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (c *Client) DoRequest(method string, path string, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, path)
	if err != nil {
		return nil, err
	}

	return c.Do(req, v)
}

func (c *Client) DoRootRequest(method string, path string, v interface{}) (*Response, error) {
	req, err := c.NewRootRequest(method, path)
	if err != nil {
		return nil, err
	}

	return c.Do(req, v)
}

func (c *Client) NewAccountRequestBody(method string, path string, data []byte) (*http.Request, error) {
	return c.NewRequestBody(method, c.accountRequestUrl(path), data)
}

func (c *Client) NewRootRequestBody(method string, path string, data []byte) (*http.Request, error) {
	return c.NewRequestBody(method, c.rawRequestURL(path), data)
}

func (c *Client) DoRequestJSON(method string, path string, obj interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequestJSON(method, path, obj)
	if err != nil {
		return nil, err
	}

	return c.Do(req, v)
}

func (c *Client) DoRootRequestJSON(method string, path string, obj interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRootRequestJSON(method, path, obj)
	if err != nil {
		return nil, err
	}

	return c.Do(req, v)
}

func getHeaderInt(r *http.Response, header string) int {
	v, _ := strconv.Atoi(r.Header.Get(header))
	return v
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	// drain and cache the request body
	originalRequestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var (
		resultError    error
		resultResponse *Response
		retries        = 3
		hitRateLimit   = false
	)

	if c.Config.RetryCount != nil {
		retries = *c.Config.RetryCount
	}

	for attempt := 0; attempt <= retries; attempt++ {
		resultError = nil
		resultResponse = nil
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

		resultResponse = &Response{
			RateLimitLimit:     getHeaderInt(resp, "Rate-Limit-Limit"),
			RateLimitRemaining: getHeaderInt(resp, "Rate-Limit-Remaining"),
			RateLimitReset:     getHeaderInt(resp, "Rate-Limit-Reset"),
		}

		decodeError := json.NewDecoder(resp.Body).Decode(resultResponse)

		var responseData []byte
		if resultResponse.ResponseRaw != nil {
			responseData = *resultResponse.ResponseRaw
		}
		resultResponse.ResponseRaw = nil

		if resp.StatusCode == http.StatusTooManyRequests && c.Config.RateLimitRetry {
			if !hitRateLimit {
				rateLimitWait := int64(resultResponse.RateLimitReset) - c.Config.Clock.Now().Unix()
				attempt--
				hitRateLimit = true
				if rateLimitWait > 0 {
					if c.Config.Logger != nil {
						c.Config.Logger.Log(fmt.Sprintf("rate limit hit, waiting for %d seconds", rateLimitWait))
					}
					c.Config.Clock.Sleep(time.Duration(rateLimitWait+1) * time.Second)
				}
				continue
			} else {
				if c.Config.Logger != nil {
					c.Config.Logger.Log("rate limit error persisted after waiting")
				}
			}
		}

		if resp.StatusCode >= 500 {
			if decodeError != nil {
				resultError = decodeError
				resultResponse = nil
			} else {
				resultError = resultResponse.Meta.Errors
			}
			continue
		}

		if v != nil {
			if w, ok := v.(io.Writer); ok {
				_, err = io.Copy(w, bytes.NewReader(responseData))
			} else {
				err = json.Unmarshal(responseData, &v)
				if err == nil {
					resultResponse.Response = v
				}
			}
		}

		if err != nil {
			return resultResponse, err
		} else if len(resultResponse.Meta.Errors) > 0 {
			return resultResponse, resultResponse.Meta.Errors
		} else {
			return resultResponse, nil
		}
	}
	return resultResponse, resultError
}

type listRetriever func(*ListOptions) (int, int, error)

// listHelper handles all the generic work of making paged requests up until
// we've recieved the last page of results.
func listHelper(lr listRetriever, limit int) error {
	var (
		opts                                         = &ListOptions{Limit: limit}
		found, firstReportedTotal, lastReportedTotal int
	)
	for {
		els, reportedtotal, err := lr(opts)
		if err != nil {
			return err
		}

		found += els

		if firstReportedTotal == 0 {
			firstReportedTotal = reportedtotal
		}
		lastReportedTotal = reportedtotal

		if reportedtotal <= opts.Offset+opts.Limit {
			break
		}

		opts.Offset += opts.Limit
	}

	// Safety check
	if firstReportedTotal != found || lastReportedTotal != found {
		return fmt.Errorf("got %d elements total, first response indicated %d, last response indicated %d", found, firstReportedTotal, lastReportedTotal)
	}

	return nil
}

func addListOptions(requrl string, opts *ListOptions) (string, error) {
	u, err := url.Parse(requrl)
	if err != nil {
		return "", nil
	}

	if opts == nil {
		return requrl, nil
	}

	q := u.Query()
	if opts.Limit != 0 {
		q.Add("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset != 0 {
		q.Add("offset", strconv.Itoa(opts.Offset))
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}
