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

	for {
		tt := c.Next()
		switch tt {
		case html.ErrorToken:
			return c.Err()
		case html.CommentToken:
			if g.config.Comments {
				logrus.Println(fmt.Sprintf("COMMENT FOUND:\t<!-- %s -->", string(c.Text())))
			}
		case html.TextToken:
			if depth > 0 {
				stxt := strings.TrimSpace(string(c.Text()))
				if stxt != "" {
					if eltype == title {
						if g.config.Title {
							logrus.Println(fmt.Sprintf("TITLE FOUND:\t<title>%s</title>", string(c.Text())))
						}
					}
					if eltype == script {
						if g.config.Scripts {
							logrus.Println(fmt.Sprintf("INLINE SCRIPT FOUND:\t%s", stxt))
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
							logrus.Println(fmt.Sprintf("LINK FOUND:\t%s", a.Val))
							break
						}
					}
				}
			} else if isTitle || isScript {
				if g.config.Scripts && isScript {
					for _, a := range t.Attr {
						if a.Key == "src" {
							logrus.Println(fmt.Sprintf("EXTERNAL SCRIPT FOUND:\t%s", a.Val))
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
