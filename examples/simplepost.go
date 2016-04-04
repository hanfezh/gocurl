package main

import "fmt"
import "github.com/ufengzhu/gocurl"

func main() {
	postthis := "moo mooo moo moo"
	easy := gocurl.NewEasy()
	if easy == nil {
		fmt.Printf("NewEasy failed\n")
		return
	}
	defer easy.Cleanup()

	easy.Setopt(gocurl.OPT_URL, "http://example.com")
	easy.Setopt(gocurl.OPT_VERBOSE, 1)
	easy.Setopt(gocurl.OPT_POSTFIELDS, postthis)

	/* if we don't provide POSTFIELDSIZE, libcurl will strlen() by itself */
	easy.Setopt(gocurl.OPT_POSTFIELDSIZE, len(postthis))

	/* Perform the request, res will get the return code */
	err := easy.Perform()
	/* Check for errors */
	if err != nil {
		fmt.Printf("Perform failed: %v\n", err)
	}
}
