package main

import (
	"errors"
	"net/url"
	"strings"
)

func normalizeURL(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return str, err
	}
	if u.Scheme == "" {
		u, err = url.Parse("https://" + str)
	}
	if u.Host == "" {
		return str, errors.New("no host")
	}
	if !strings.Contains(u.Host, ".") {
		return str, errors.New("invalid host")
	}
	for u.Path != "" && u.Path[len(u.Path)-1] == '/' {
		u.Path = u.Path[0 : len(u.Path)-1]
	}
	strippedURL := u.Host + u.Path
	// println(stripped_url)
	return strippedURL, err
}
