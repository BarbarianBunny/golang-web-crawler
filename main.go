package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
)

func main() {
	lenArgs := len(os.Args)
	if lenArgs == 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if lenArgs > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	baseURL, err := normalizeURL(baseURL)
	if err != nil {
		fmt.Println("normalizing baseURL failed")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", baseURL)
	pages := crawlPage(baseURL, baseURL, make(map[string]int))

	fmt.Println("URLs:")
	for _, url := range slices.Sorted(maps.Keys(pages)) {
		fmt.Printf("%4d - %v\n", pages[url], url)
	}
}
