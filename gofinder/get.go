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

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		logrus.WithError(err).Errorln("Error occurred creating request")
		return nil, err
	}

	// For some reason, people block this type of thing if we don't have a "real" user agent
	req.Header.Set("User-Agent", g.config.UserAgent)

	resp, err := g.client.Do(req)
	if err != nil {
		logrus.WithError(err).Errorln("Error occurred trying to reach %s", parsedURL.String())
		return nil, err
	}

	return resp, nil
}
