package yext

import (
	"net/http"
	"os"
)

const (
	SandboxHost      string = "api-sandbox.yext.com"
	ProductionHost   string = "api.yext.com"
	DefaultAccountId string = "me"
	DefaultVParam    string = "20161101"
)

type Config struct {
	HTTPClient *http.Client
	BaseUrl    string
	ApiKey     string
	AccountId  string
	VParam     string

	RetryCount *int

	Logger Logger
}

func NewConfig() *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
		AccountId:  DefaultAccountId,
		VParam:     DefaultVParam,
	}
}

func NewDefaultConfig() *Config {
	return NewConfig().WithProductionHost().WithRetries(3)
}

func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}

func (c *Config) WithBaseUrl(baseUrl string) *Config {
	c.BaseUrl = baseUrl
	return c
}

func (c *Config) WithHost(host string) *Config {
	return c.WithBaseUrl("https://" + host + "/v2")
}

func (c *Config) WithSandboxHost() *Config {
	return c.WithHost(SandboxHost)
}

func (c *Config) WithProductionHost() *Config {
	return c.WithHost(ProductionHost)
}

func (c *Config) WithApiKey(apikey string) *Config {
	c.ApiKey = apikey
	return c
}

func (c *Config) WithAccountId(accountId string) *Config {
	c.AccountId = accountId
	return c
}

func (c *Config) WithVParam(vParam string) *Config {
	c.VParam = vParam
	return c
}

func (c *Config) WithEnvCredentials() *Config {
	return c.WithApiKey(os.ExpandEnv("$YEXT_API_KEY")).WithAccountId(os.ExpandEnv("$YEXT_API_CUSTOMERID"))
}

func (c *Config) WithRetries(r int) *Config {
	c.RetryCount = Int(r)
	return c
}

func (c *Config) WithLogger(l Logger) *Config {
	c.Logger = l
	return c
}

func (c *Config) WithStdLogger() *Config {
	c.Logger = NewStdLogger()
	return c
}
