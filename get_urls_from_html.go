package main

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

func getURLsFromHTML(htmlBody string, rawBaseURL string) ([]string, error) {
	urls := []string{}
	htmlNodes, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return urls, err
	}
	for n := range htmlNodes.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if a.Val != "" && a.Val[0] == '/' {
						a.Val = rawBaseURL + a.Val
					}
					urls = append(urls, a.Val)
				}
			}
		}
	}
	return urls, err
}
