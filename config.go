package yext

import (
	"net/http"
	"os"
	"time"

	"github.com/jonboulle/clockwork"
)

const (
	SandboxHost    string = "api-sandbox.yext.com"
	ProductionHost string = "api.yext.com"
	EUHost         string = "api.eu.yextapis.com"
	AccountId      string = "me"
	Version        string = "20180226"
)

type Config struct {
	HTTPClient *http.Client
	BaseUrl    string
	ApiKey     string
	AccountId  string
	Version    string

	RetryCount     *int
	RateLimitRetry bool

	Clock  clockwork.Clock
	Logger Logger
}

func NewConfig() *Config {
	return &Config{
		HTTPClient: http.DefaultClient,
		AccountId:  AccountId,
		Version:    Version,
		Clock:      clockwork.NewRealClock(),
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

func (c *Config) WithEUHost() *Config {
	return c.WithHost(EUHost)
}

func (c *Config) WithProductionHost() *Config {
	return c.WithHost(ProductionHost)
}

func (c *Config) WithApiKey(apikey string) *Config {
	c.ApiKey = apikey
	return c
}

func (c *Config) WithAccountId(id string) *Config {
	c.AccountId = id
	return c
}

func (c *Config) WithVersion(v string) *Config {
	c.Version = v
	return c
}

func (c *Config) WithTodaysVersion() *Config {
	c.Version = time.Now().Format("20060102")
	return c
}

func (c *Config) WithEnvCredentials() *Config {
	c = c.WithApiKey(os.ExpandEnv("$YEXT_API_KEY"))
	if os.ExpandEnv("$YEXT_API_ACCOUNTID") != "" {
		c = c.WithAccountId(os.ExpandEnv("$YEXT_API_ACCOUNTID"))
	}
	return c
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

func (c *Config) WithRateLimitRetry() *Config {
	c.RateLimitRetry = true
	return c
}

func (c *Config) WithMockClock() *Config {
	c.Clock = clockwork.NewFakeClock()
	return c
}
