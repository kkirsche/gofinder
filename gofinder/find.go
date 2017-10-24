package gofinder

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const title = "title"
const script = "script"

// Find is used to find comment nodes within a page
func (g *GoFinder) Find(resp *http.Response) error {
	c := html.NewTokenizer(resp.Body)
	depth := 0
	eltype := "unknown"
	found := false

	if !g.config.Comments || !g.config.Links || !g.config.Scripts || !g.config.Title {
		found = true
		logrus.Warnln("Not searching for any items...")
	}

	for {
		tt := c.Next()
		switch tt {
		case html.ErrorToken:
			if !found {
				logrus.WithField("url", resp.Request.URL.String()).Warnln("NO ITEMS FOUND")
			}
			return c.Err()
		case html.CommentToken:
			if g.config.Comments {
				found = true
				stxt := strings.TrimSpace(string(c.Text()))
				logrus.WithFields(logrus.Fields{
					"url":     resp.Request.URL.String(),
					"comment": stxt,
				}).Println(fmt.Sprintf("COMMENT FOUND:\t<!-- %s -->", stxt))
			}
		case html.TextToken:
			if depth > 0 {
				stxt := strings.TrimSpace(string(c.Text()))
				if stxt != "" {
					if eltype == title {
						if g.config.Title {
							found = true
							logrus.WithFields(logrus.Fields{
								"url":   resp.Request.URL.String(),
								"title": stxt,
							}).Println(fmt.Sprintf("TITLE FOUND:\t<title>%s</title>", stxt))
						}
					}
					if eltype == script {
						if g.config.Scripts {
							found = true
							logrus.WithFields(logrus.Fields{
								"url":          resp.Request.URL.String(),
								"inlineScript": stxt,
							}).Println(fmt.Sprintf("INLINE SCRIPT FOUND:\t%s", stxt))
						}
					}
				}
			}
		case html.StartTagToken:
			t := c.Token()
			eltype = t.Data
			isAnchor := t.Data == "a"
			isTitle := t.Data == title
			isScript := t.Data == script
			if isAnchor {
				if g.config.Links {
					for _, a := range t.Attr {
						if a.Key == "href" {
							found = true
							logrus.WithFields(logrus.Fields{
								"url":      resp.Request.URL.String(),
								"linkHREF": a.Val,
							}).Println(fmt.Sprintf("LINK FOUND:\t%s", a.Val))
							break
						}
					}
				}
			} else if isTitle || isScript {
				if g.config.Scripts && isScript {
					for _, a := range t.Attr {
						if a.Key == "src" {
							found = true
							logrus.WithFields(logrus.Fields{
								"url":            resp.Request.URL.String(),
								"externalScript": a.Val,
							}).Println(fmt.Sprintf("EXTERNAL SCRIPT FOUND:\t%s", a.Val))
							break
						}
					}
				}
				depth++
			}
		case html.EndTagToken:
			t := c.Token()

			isTitle := t.Data == title
			isScript := t.Data == script
			if isTitle || isScript {
				depth--
			}
		}
	}
}
