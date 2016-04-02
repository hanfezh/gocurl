package main

import "fmt"
import "github.com/ufengzhu/gocurl/curl"

func main() {
	easy := curl.NewEasy()
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.OPT_URL, "http://example.com")
	/* example.com is redirected, so we tell libcurl to follow redirection */
	easy.EasySetopt(curl.OPT_FOLLOWLOCATION, 1)
	easy.EasySetopt(curl.OPT_VERBOSE, 1)

	/* Perform the request, res will get the return code */
	err := easy.EasyPerform()
	/* Check for errors */
	if err != nil {
		fmt.Printf("EasyPerform failed: %v\n", err)
	}
}
