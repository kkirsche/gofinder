package gofinder

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// transportWrapper wraps http.Transport to allow us to record URLs which we
// are redirected through
type transportWrapper struct {
	*http.Transport
}

// RoundTrip executes a single HTTP transaction, returning a Response for the
// provided Request
func (t *transportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Use the default transport for this function, we only want to add this
	// function for logging purposes, not adjusting the actual transport mechanism.

	transport := t.Transport

	if transport == nil {
		transport = http.DefaultTransport.(*http.Transport)
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	logrus.WithFields(logrus.Fields{
		"Status Code": resp.StatusCode,
		"URL":         req.URL.String(),
	}).Debugln("Retrieved site")

	return resp, err
}
