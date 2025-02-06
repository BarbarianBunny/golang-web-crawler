package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get("https://" + rawURL)
	if err != nil {
		return rawURL, err
	}
	if resp.StatusCode >= 400 {
		return rawURL, errors.New("status code error")
	}
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		return rawURL, errors.New("content-type is not text/html")
	}
	htmlBody, err := io.ReadAll(resp.Body)
	return string(htmlBody), err
}
