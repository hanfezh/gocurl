package main

import "fmt"
import "github.com/ufengzhu/gocurl/curl"

func main() {
	postthis := "moo mooo moo moo"
	easy := curl.NewEasy()
	if easy == nil {
		fmt.Printf("NewEasy failed\n")
		return
	}
	defer easy.EasyCleanup()

	easy.EasySetopt(curl.OPT_URL, "http://example.com")
	easy.EasySetopt(curl.OPT_VERBOSE, 1)
	easy.EasySetopt(curl.OPT_POSTFIELDS, postthis)

	/* if we don't provide POSTFIELDSIZE, libcurl will strlen() by itself */
	easy.EasySetopt(curl.OPT_POSTFIELDSIZE, len(postthis))

	/* Perform the request, res will get the return code */
	err := easy.EasyPerform()
	/* Check for errors */
	if err != nil {
		fmt.Printf("EasyPerform failed: %v\n", err)
	}
}
