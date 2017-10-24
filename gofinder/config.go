package gofinder

import "github.com/sirupsen/logrus"

// Config represents a configuration object for Gofinder
type Config struct {
	Title     bool
	Links     bool
	Comments  bool
	Scripts   bool
	UserAgent string
}

// NewConfig creates the structure for us
func NewConfig(ua string, te, le, ce, se bool) *Config {
	logrus.WithFields(logrus.Fields{
		"userAgent":       ua,
		"titleEnabled":    te,
		"linksEnabled":    le,
		"commentsEnabled": ce,
		"scriptsEnabled":  se,
	}).Debugln("Building config")
	return &Config{
		Title:     te,
		Links:     le,
		Comments:  ce,
		Scripts:   se,
		UserAgent: ua,
	}
}
