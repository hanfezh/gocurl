package main

import "fmt"
import "github.com/ufengzhu/gocurl/curl"

func main() {
	fmt.Printf("CURL version: %s\n", curl.Version())
	info := curl.VersionInfo()
	fmt.Printf("CURL version: %v\n", info)

	easy := curl.NewEasy()
	defer easy.Cleanup()

	easy.Setopt(curl.OPT_URL, "http://example.com")
	/* example.com is redirected, so we tell libcurl to follow redirection */
	easy.Setopt(curl.OPT_FOLLOWLOCATION, 1)
	easy.Setopt(curl.OPT_VERBOSE, 1)

	/* Perform the request, res will get the return code */
	err := easy.Perform()
	/* Check for errors */
	if err != nil {
		fmt.Printf("Perform failed: %v\n", err)
	}
}
