package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="foo">bar</a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTMO document and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return []Link{}, err
	}
	links := []Link{}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			text := n.FirstChild.Data
			text = strings.TrimSpace(text)
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					l := Link{Href: attr.Val, Text: text}
					links = append(links, l)
				}
			}
		}
	}
	return links, nil
}
