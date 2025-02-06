package main

import (
	"fmt"
	"strings"
)

func crawlPage(rawBaseURL string, rawCurrentURL string, pages map[string]int) map[string]int {
	if !strings.Contains(rawCurrentURL, rawBaseURL) {
		return pages
	}

	// Normalize URL
	currentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Failed to normalize URL: %v error: %v\n", rawCurrentURL, err)
		return pages
	}

	// Update Pages
	pageCount, pageExists := pages[currentURL]
	if pageExists {
		pages[currentURL] = pageCount + 1
		return pages
	}

	pages[currentURL] = 1

	// Get HTML
	htmlBody, err := getHTML(currentURL)
	if err != nil {
		fmt.Printf("Failed to get HTML for URL: %v error: %v\n", currentURL, err)
		return pages
	}

	// Get URLS
	urls, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("Failed to get URLs from html from url: %v error: %v\n", currentURL, err)
		return pages
	}

	// Using the range keyword
	for _, url := range urls {
		pages = crawlPage(rawBaseURL, url, pages)
	}

	return pages
}
