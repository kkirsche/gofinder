package gofinder

import (
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

// Get is used to retrieve a URL and ensure that the URL we request has a
// configured scheme.
func (g *GoFinder) Get(urlstr string) (*http.Response, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to parse URL")
		return nil, err
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	resp, err := g.client.Get(parsedURL.String())
	if err != nil {
		logrus.WithError(err).Errorln("Error occurred trying to reach %s", parsedURL.String())
		return nil, err
	}

	return resp, nil
}
