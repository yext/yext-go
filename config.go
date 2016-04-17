package yext

import (
	"net/http"
	"os"
)

const (
	SandboxHost    string = "api-sandbox.yext.com"
	ProductionHost string = "api.yext.com"
)

type Config struct {
	HTTPClient *http.Client
	BaseUrl    string

	Username   string
	Password   string
	CustomerId string

	RetryCount *int

	Logger Logger
}

func NewConfig() *Config {
	return &Config{}
}

func NewDefaultConfig() *Config {
	return NewConfig().WithHTTPClient(http.DefaultClient).WithProductionHost().WithRetries(3)
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
	return c.WithBaseUrl("https://" + host + "/v1")
}

func (c *Config) WithSandboxHost() *Config {
	return c.WithHost(SandboxHost)
}

func (c *Config) WithProductionHost() *Config {
	return c.WithHost(ProductionHost)
}

func (c *Config) WithCredentials(username, password, customerid string) *Config {
	c.Username = username
	c.Password = password
	c.CustomerId = customerid
	return c
}

func (c *Config) WithEnvCredentials() *Config {
	return c.WithCredentials(
		os.ExpandEnv("$YEXT_API_USER"),
		os.ExpandEnv("$YEXT_API_PASS"),
		os.ExpandEnv("$YEXT_API_CUSTOMERID"))
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
