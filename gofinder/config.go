package gofinder

import "github.com/sirupsen/logrus"

// Config represents a configuration object for Gofinder
type Config struct {
	Title    bool
	Links    bool
	Comments bool
	Scripts  bool
}

// NewConfig creates the structure for us
func NewConfig(title, link, comment, script bool) *Config {
	logrus.WithFields(logrus.Fields{
		"titleEnabled":   title,
		"linkEnabled":    link,
		"commentEnabled": comment,
	}).Debugln("Building config")
	return &Config{
		Title:    title,
		Links:    link,
		Comments: comment,
		Scripts:  script,
	}
}
