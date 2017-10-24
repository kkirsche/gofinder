package gofinder

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

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
				logrus.Println(fmt.Sprintf("COMMENT FOUND:\t%s", string(c.Text())))
			}
		case html.TextToken:
			if depth > 0 {
				if eltype == "TITLE" {
					if g.config.Title {
						logrus.Println(fmt.Sprintf("TITLE FOUND:\t%s", string(c.Text())))
					}
				}
			}
		case html.StartTagToken:
			t := c.Token()
			isAnchor := t.Data == "a"
			isTitle := t.Data == "title"
			if isAnchor {
				if g.config.Links {
					for _, a := range t.Attr {
						if a.Key == "href" {
							logrus.Debugln("Anchor Found")
							logrus.Println(fmt.Sprintf("LINK FOUND:\t%s", a.Val))
							break
						}
					}
				}
			} else if isTitle {
				logrus.Debugln("Title Found")
				eltype = "TITLE"
				depth++
			}
		case html.EndTagToken:
			t := c.Token()

			isTitle := t.Data == "title"
			if isTitle {
				depth--
			}
		}
	}
}
