package main

import (
	"fmt"
	"os"
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
	fmt.Printf("starting crawl of: %v\n", baseURL)
	htmlBody, err := getHTML(baseURL)
	if err != nil {
		fmt.Printf("issue getting html: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("HTML:\n%v\n", htmlBody)

}
