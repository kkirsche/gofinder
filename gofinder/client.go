package gofinder

import (
	"net/http"
	"time"
)

// GoFinder is the default client used to retrieve and search for content
type GoFinder struct {
	client *http.Client
	config *Config
}

// NewClient is used to create a new GoFinder client
func NewClient(config *Config) *GoFinder {
	t := &transportWrapper{
		Transport: http.DefaultTransport.(*http.Transport),
	}

	client := &GoFinder{
		client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: t,
		},
		config: config,
	}

	return client
}
