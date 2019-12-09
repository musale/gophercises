package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link - to an anchor tag href and text.
type Link struct {
	Href, Text string
}

// Parse - returns a slice of links from a HTML document.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := findLinks(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var l Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			l.Href = attr.Val
			break
		}
	}
	l.Text = linkText(n)
	return l
}

func linkText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var txt string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		txt += linkText(c)
	}
	return strings.Join(strings.Fields(txt), " ")
}

func findLinks(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinks(c)...)
	}
	return ret
}
