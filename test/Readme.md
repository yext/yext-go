This directory contains integration tests for the Yext API.  Since it makes calls against the production Yext API, it is intended to be run manually.  In order to run the tests,
you need an API key and need to expose it as an environment variable:

```
YEXT_API_KEY=[your api key] go run main.go
```
